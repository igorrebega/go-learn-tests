package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("real", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("not real", func(t *testing.T) {
		_, err := dictionary.Search("test2")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)
		if err != nil {
			t.Fatal("Should be no error, got:", err)
		}

		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		word := "test"
		definition := "def"

		dictionary := Dictionary{word: definition}
		newDefinition := "def2"

		err := dictionary.Update(word, newDefinition)
		if err != nil {
			t.Fatal("Should be no error, got:", err)
		}

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update("new word", "definition")

		assertError(t, err, ErrWordDoesntExist)
	})
}

func TestDelete(t *testing.T) {
	word := "word"
	dictionary := Dictionary{word: "test"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if err != ErrNotFound {
		t.Fatalf("Expected %s to be deleted", word)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
