package calculadora

import (
	"context"
	"fmt"
	"math"
	"testing"

	"github.com/cucumber/godog"
)

// estado guarda os dois números e o resultado da operação entre os passos
// de um mesmo cenário. Vai dentro do context.Context, do jeito que o godog espera.
type estado struct {
	a, b      float64
	resultado float64
	erro      error
}

type estadoCtxKey struct{}

func euTenhoOsNumeros(ctx context.Context, a, b float64) (context.Context, error) {
	return context.WithValue(ctx, estadoCtxKey{}, &estado{
		a: a,
		b: b,
	}), nil
}

func pegarEstado(ctx context.Context) (*estado, error) {
	e, ok := ctx.Value(estadoCtxKey{}).(*estado)
	if !ok {
		return nil, fmt.Errorf("nenhum número foi definido ainda")
	}

	return e, nil
}

func euSomoOsNumeros(ctx context.Context) error {
	e, err := pegarEstado(ctx)
	if err != nil {
		return err
	}

	e.resultado = Somar(e.a, e.b)

	return nil
}

func euSubtraioOsNumeros(ctx context.Context) error {
	e, err := pegarEstado(ctx)
	if err != nil {
		return err
	}

	e.resultado = Subtrair(e.a, e.b)

	return nil
}

func euMultiplicoOsNumeros(ctx context.Context) error {
	e, err := pegarEstado(ctx)
	if err != nil {
		return err
	}

	e.resultado = Multiplicar(e.a, e.b)

	return nil
}

func euDividoOsNumeros(ctx context.Context) error {
	e, err := pegarEstado(ctx)
	if err != nil {
		return err
	}

	e.resultado, e.erro = Dividir(e.a, e.b)

	return nil
}

func oResultadoDeveSer(ctx context.Context, esperado float64) error {
	e, err := pegarEstado(ctx)
	if err != nil {
		return err
	}

	if e.erro != nil {
		return fmt.Errorf("a operação devolveu um erro inesperado: %w", e.erro)
	}

	const tolerancia = 0.000001

	if math.Abs(e.resultado-esperado) > tolerancia {
		return fmt.Errorf(
			"resultado = %v; esperado %v",
			e.resultado,
			esperado,
		)
	}

	return nil
}

func deveOcorrerUmErroDeDivisaoPorZero(ctx context.Context) error {
	e, err := pegarEstado(ctx)
	if err != nil {
		return err
	}

	if e.erro == nil {
		return fmt.Errorf(
			"esperava um erro de divisão por zero, mas nenhum erro ocorreu",
		)
	}

	return nil
}

func InitializeScenario(sc *godog.ScenarioContext) {
	// O padrão aceita números positivos, negativos e decimais.
	sc.Given(
		`^que tenho os números (-?\d+(?:\.\d+)?) e (-?\d+(?:\.\d+)?)$`,
		euTenhoOsNumeros,
	)

	sc.When(`^eu somo os números$`, euSomoOsNumeros)
	sc.When(`^eu subtraio os números$`, euSubtraioOsNumeros)
	sc.When(`^eu multiplico os números$`, euMultiplicoOsNumeros)
	sc.When(`^eu divido os números$`, euDividoOsNumeros)

	sc.Then(
		`^o resultado deve ser (-?\d+(?:\.\d+)?)$`,
		oResultadoDeveSer,
	)

	sc.Then(
		`^deve ocorrer um erro de divisão por zero$`,
		deveOcorrerUmErroDeDivisaoPorZero,
	)
}

// TestFeatures é o ponto de entrada que o "go test" enxerga. Por baixo dos
// panos, ele manda o godog ler os arquivos .feature e rodar cada cenário,
// chamando os passos registrados em InitializeScenario.
func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("um ou mais cenários do BDD falharam")
	}
}
