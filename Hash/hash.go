package Hash

// Hash - Mapeia a tabela hash e retorna o valor da key escolhida
func Hash(key string) float64 {

	tabela := map[string]float64{
		"maçã":    2.99,
		"leite":   4.99,
		"abacate": 3.99,
	}

	return tabela[key]
}
