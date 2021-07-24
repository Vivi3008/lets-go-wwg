package multiplicacao

import "testing"

func TestMultiplique(t *testing.T) {
	testes := []struct {
		nome     string
		x        int
		y        int
		esperado int
	}{
		{
			nome:     "7 x 5 = 35",
			x:        7,
			y:        5,
			esperado: 35,
		},
		{
			nome:     "3 x 20 = 60",
			x:        3,
			y:        20,
			esperado: 60,
		},
	}

	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			obtive := Multiplique(teste.x, teste.y)
			if obtive != teste.esperado {
				t.Errorf("espero '%d', mas obtive '%d'", teste.esperado, obtive)
			}
		})
	}

	/* t.Run("7 vezes 5 é 35", func(t *testing.T) {
		obtive := Multiplique(7, 5)
		espero := 35
		if obtive != espero {
			t.Errorf("Espero %d; mas obtive %d", espero, obtive)
		}
	})

	t.Run("3 vezes 20 é 60", func(t *testing.T) {
		obtive := Multiplique(3, 20)
		espero := 60
		if obtive != espero {
			t.Errorf("Espero %d; mas obtive %d", espero, obtive)
		}
	})
	*/
}
