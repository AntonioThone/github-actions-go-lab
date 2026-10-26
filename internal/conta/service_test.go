package conta

import "testing"

func TestCalculcarConta (t *testing.T) {

	resultado := CalcularConta(
		10000,
		10,
		5,
	)

	if resultado != 2200 {
		t.Errorf(
			"esperado 2200, recebido %.2f", resultado,
		)
	}
}
