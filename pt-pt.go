package ntw

var portugueseMegasSingular = []string{"", "mil", "milhão", "mil milhões", "bilião"}
var portugueseMegasPlural = []string{"", "mil", "milhões", "mil milhões", "bilhões"}
var portugueseUnitsAdjectives = []string{"", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove"}
var portugueseUnits = []string{"", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove"}
var portugueseHundreds = []string{"", "cem", "duzentos", "trezentos", "quatrocentos", "quinhentos", "seiscentos", "setecentos", "oitocentos", "novecentos", "cento"}
var portugueseTens = []string{"", "dez", "vinte", "trinta", "quarenta", "cinquenta", "sessenta", "setenta", "oitenta", "noventa"}
var portugueseTeens = []string{"dez", "onze", "doze", "treze", "catorze", "quinze", "dezasseis", "dezasete", "dezoito", "dezanove"}

// IntegerToportuguese converts an integer to portuguese words
func IntegerToPortuguese_PT(input int) string {
	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "menos")
		input *= -1
	}

}
