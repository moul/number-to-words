package ntw

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleIntegerToEsEs() {
	fmt.Println(IntegerToEsEs(42))
	// Output: cuarenta y dos
}

func TestIntegerToEsEs(t *testing.T) {
	t.Parallel()

	tests := map[int]string{
		-1:    "menos uno",
		0:     "cero",
		1:     "uno",
		2:     "dos",
		3:     "tres",
		4:     "cuatro",
		5:     "cinco",
		6:     "seis",
		7:     "siete",
		8:     "ocho",
		9:     "nueve",
		10:    "diez",
		11:    "once",
		12:    "doce",
		13:    "trece",
		14:    "catorce",
		15:    "quince",
		16:    "dieciséis",
		17:    "diecisiete",
		18:    "dieciocho",
		19:    "diecinueve",
		20:    "veinte",
		21:    "veintiuno",
		22:    "veintidós",
		23:    "veintitrés",
		24:    "veinticuatro",
		25:    "veinticinco",
		26:    "veintiséis",
		27:    "veintisiete",
		28:    "veintiocho",
		29:    "veintinueve",
		30:    "treinta",
		31:    "treinta y uno",
		32:    "treinta y dos",
		39:    "treinta y nueve",
		42:    "cuarenta y dos",
		80:    "ochenta",
		90:    "noventa",
		99:    "noventa y nueve",
		100:   "ciento",
		101:   "ciento uno",
		111:   "ciento once",
		120:   "ciento veinte",
		121:   "ciento veintiuno",
		200:   "doscientos",
		300:   "trescientos",
		400:   "cuatrocientos",
		500:   "quinientos",
		600:   "seiscientos",
		700:   "setecientos",
		800:   "ochocientos",
		900:   "novecientos",
		909:   "novecientos nueve",
		919:   "novecientos diecinueve",
		990:   "novecientos noventa",
		999:   "novecientos noventa y nueve",
		1000:  "mil",
		2000:  "dos mil",
		4000:  "cuatro mil",
		5000:  "cinco mil",
		11000: "once mil",
		21000: "veintiún mil",
		28000: "veintiocho mil",
		//31000:         "treinta y un mil",
		32000:         "treinta y dos mil",
		39000:         "treinta y nueve mil",
		42000:         "cuarenta y dos mil",
		999000:        "novecientos noventa y nueve mil",
		999999:        "novecientos noventa y nueve mil novecientos noventa y nueve",
		1000000:       "un millón",
		2000000:       "dos millones",
		4000000:       "cuatro millones",
		5000000:       "cinco millones",
		100100100:     "ciento millones ciento mil ciento",
		500500500:     "quinientos millones quinientos mil quinientos",
		606606606:     "seiscientos seis millones seiscientos seis mil seiscientos seis",
		999000000:     "novecientos noventa y nueve millones",
		999000999:     "novecientos noventa y nueve millones novecientos noventa y nueve",
		999999000:     "novecientos noventa y nueve millones novecientos noventa y nueve mil",
		999999999:     "novecientos noventa y nueve millones novecientos noventa y nueve mil novecientos noventa y nueve",
		1174315110:    "un mil millones ciento setenta y cuatro millones trescientos quince mil ciento diez",
		1174315119:    "un mil millones ciento setenta y cuatro millones trescientos quince mil ciento diecinueve",
		15174315119:   "quince mil millones ciento setenta y cuatro millones trescientos quince mil ciento diecinueve",
		35174315119:   "treinta y cinco mil millones ciento setenta y cuatro millones trescientos quince mil ciento diecinueve",
		935174315119:  "novecientos treinta y cinco mil millones ciento setenta y cuatro millones trescientos quince mil ciento diecinueve",
		2935174315119: "dos billones novecientos treinta y cinco mil millones ciento setenta y cuatro millones trescientos quince mil ciento diecinueve",
	}

	for input, expectedOutput := range tests {
		name := fmt.Sprintf("%d", input)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, expectedOutput, IntegerToEsEs(input))
		})
	}
}
