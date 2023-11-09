package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}

	sum = Add(2, 2, 2)
	expected = 6

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}

	sum = Add()
	expected = 0

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
