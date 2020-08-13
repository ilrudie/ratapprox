package ratapprox

import "testing"

/*func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}*/

func TestFarey(t *testing.T) {
	want := Fraction{1, 2}
	if got := farey(Fraction{0, 1}, Fraction{1, 1}); got != want {
		t.Errorf("farey() = %q, want %q", got, want)
	}
	want = Fraction{1, 3}
	if got := farey(Fraction{0, 1}, Fraction{1, 2}); got != want {
		t.Errorf("farey() = %q, want %q", got, want)
	}
}

func TestDivFraction(t *testing.T) {
	want := .5
	if got := (Fraction{1, 2}.divFraction()); got != want {
		t.Errorf("divFraction() = %f, want %f", got, want)
	}
	want = .125
	if got := (Fraction{1, 8}.divFraction()); got != want {
		t.Errorf("divFraction() = %f, want %f", got, want)
	}
}

func TestRatApprox(t *testing.T) {
	want := Fraction{1, 2}
	if got := RatApprox(.5); got != want {
		t.Errorf("farey() = %q, want %q", got, want)
	}
	want = Fraction{1, 8}
	if got := RatApprox(.125); got != want {
		t.Errorf("farey() = %q, want %q", got, want)
	}
	want = Fraction{19, 20}
	if got := RatApprox(.95); got != want {
		t.Errorf("farey() = %q, want %q", got, want)
	}
	want = Fraction{15, 16}
	if got := RatApprox(.9375); got != want {
		t.Errorf("farey() = %q, want %q", got, want)
	}
}
