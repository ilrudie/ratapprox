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
