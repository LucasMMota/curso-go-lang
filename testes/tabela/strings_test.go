package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	t.Parallel()
	testes := []struct { // define a struct
		texto    string
		parte    string
		esperado int
	}{ // passa um data set
		{"Lucas é show", "Lucas", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"leonardo", "o", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte) // retorna a posicao da palavra na string

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
