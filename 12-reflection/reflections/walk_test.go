package reflections

import (
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Raven", 32},
			[]string{"Raven"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !slices.Equal(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Kristine"}
			aChannel <- Profile{32, "Raven"}

			defer close(aChannel)
		}()

		var got []string
		want := []string{"Kristine", "Raven"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Kristine"}, Profile{32, "Raven"}
		}

		var got []string
		want := []string{"Kristine", "Raven"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

// func TestWalk(t *testing.T) {
// 	expected := "Chris"
// 	var got []string

// 	x := struct {
// 		Name string
// 	}{expected}

// 	walk(x, func(input string) {
// 		got = append(got, input)
// 	})

// 	if len(got) != 1 {
// 		t.Errorf("wrong number of function calls, got %d, want %d", len(got), 1)
// 	}

// 	if got[0] != expected {
// 		t.Errorf("got %q, want %q", got[0], expected)
// 	}
// }
