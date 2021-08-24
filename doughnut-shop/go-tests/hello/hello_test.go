package hello

import "testing"

func TestSayHello(t *testing.T) {

	assertion := func(got string, want string, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("with a valid name the SayHello says hello to people", func(t *testing.T) {
		got := sayHello("Siddiq", "English")
		want := "Hi Siddiq!"
		assertion(got, want, t)
	})

	t.Run("with empty name, the sayHello function will return HI!", func(t *testing.T) {
		got := sayHello("", "")
		want := "Hi!"
		assertion(got, want, t)
	})
	t.Run("Hi in Spanish", func(t *testing.T) {
		got := sayHello("Siddiq", "Spanish")
		want := "Hola Siddiq!"
		assertion(got, want, t)
	})

	t.Run("Hi in French", func(t *testing.T) {
		got := sayHello("Siddiq", "French")
		want := "Bonjour Siddiq!"
		assertion(got, want, t)
	})

}
