package conta

func CalcularConta(
	valor float64,
	taxaServico float64,
	pessoas int,
) float64 {
	total := valor + (valor * taxaServico / 100)

	return total / float64(pessoas)
}
