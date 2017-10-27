package ntw

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIntegerToPortuguese(t *testing.T) {
	Convey("Testing IntegerToPortuguese()", t, FailureContinues, func() {
		testing := map[int]string{
			0:          "zero",
			1:          "um",
			2:          "dois",
			3:          "três",
			4:          "quatro",
			5:          "cinco",
			6:          "seis",
			7:          "sete",
			8:          "oito",
			9:          "nove",
			10:         "dez",
			11:         "onze",
			12:         "doze",
			13:         "treze",
			14:         "catorze",
			15:         "quinze",
			16:         "dezasseis",
			17:         "dezasete",
			18:         "dezoito",
			19:         "dezanove",
			20:         "vinte",
			21:         "vinte e um",
			22:         "vinte e dois",
			23:         "vinte e três",
			24:         "vinte e quatro",
			25:         "vinte e cinco",
			26:         "vinte e seis",
			27:         "vinte e sete",
			28:         "vinte e oito",
			29:         "vinte e nove",
			30:         "trinta",
			31:         "trinta e um",
			32:         "trinta e dois",
			39:         "trinta e nove",
			42:         "quarenta e dois",
			80:         "oitenta",
			90:         "noventa",
			99:         "noventa e nove",
			100:        "cem",
			101:        "cento e um",
			111:        "cento e onze",
			120:        "cento e vinte",
			121:        "cento e vinte e um",
			200:        "duzentos",
			300:        "trezentos",
			400:        "quatrocentos",
			500:        "quinhentos",
			600:        "seiscentos",
			700:        "setecentos",
			800:        "oitocentos",
			900:        "novecentos",
			909:        "novecentos e nove",
			919:        "novecentos e dezanove",
			990:        "novecentos e noventa",
			999:        "novecentos e noventa e nove",
			1000:       "mil",
			1337:       "mil trezentos e trinta e sete",
			2000:       "dois mil",
			4000:       "quatro mil",
			5000:       "cinco mil",
			11000:      "onze mil",
			21000:      "vinte e um mil",
			28000:      "vinte e oito mil",
			31000:      "trinta e um mil",
			32000:      "trinta e dois mil",
			39000:      "trinta e nove mil",
			42000:      "quarenta e dois mil",
			999000:     "novecentos e noventa e nove mil",
			999999:     "novecentos e noventa e nove mil novecentos e noventa e nove",
			1000000:    "um milhão",
			2000000:    "dois milhões",
			4000000:    "quatro milhões",
			5000000:    "cinco milhões",
			100100100:  "cem milhões cem mil e cem",
			500500500:  "quinhentos milhões quinhentos mil e quinhentos",
			500500501:  "quinhentos milhões quinhentos mil quinhentos e um",
			606606606:  "seiscentos e seis milhões seiscentos e seis mil seiscentos e seis",
			999000000:  "novecentos e noventa e nove milhões",
			999000999:  "novecentos e noventa e nove milhões novecentos e noventa e nove",
			999999000:  "novecentos e noventa e nove milhões novecentos e noventa e nove mil",
			999999999:  "novecentos e noventa e nove milhões novecentos e noventa e nove mil novecentos e noventa e nove",
			1174315110: "mil milhões cento e setenta e quatro milhões trezentos e quinze mil cento e dez",
			1174315119: "mil milhões cento e setenta e quatro milhões trezentos e quinze mil cento e dezanove",

			1234567890:    "mil milhões duzentos e trinta e quatro milhões quinhentos e sessenta e sete mil oitocentos e noventa",
			15174315119:   "quinze mil milhões cento e setenta e quatro milhões trezentos e quinze mil cento e dezanove",
			35174315119:   "trinta e cinco mil milhões cento e setenta e quatro milhões trezentos e quinze mil cento e dezanove",
			935174315119:  "novecentos e trinta e cinco mil milhões cento e setenta e quatro milhões trezentos e quinze mil cento e dezanove",
			1000000000000: "um bilião",

			1935174315119: "um bilião novecentos e trinta e cinco mil milhões cento e setenta e quatro milhões trezentos e quinze mil cento e dezanove",
			2935174315119: "dois bilhões novecentos e trinta e cinco mil milhões cento e setenta e quatro milhões trezentos e quinze mil cento e dezanove",
		}
		for input, expectedOutput := range testing {
			So(IntegerToPortuguesePT(input), ShouldEqual, expectedOutput)
		}

		// testing negative values
		So(IntegerToPortuguesePT(-1), ShouldEqual, "menos um")
	})
}
