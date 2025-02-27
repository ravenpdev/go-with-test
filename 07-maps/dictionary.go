package maps

// An interesting propery of maps is that you can modify them without passing as an address to it
// (e.g &myMap)
// This may make them feel like a "reference type"
// A map value is a pointer to a runtime.hmap structure

// A gotcha with maps is that hey can be an il value. A nil map behaves like an
// empty map when reading, but attempts to write to a nil map will cause a runtime panic.

// Therefor, you should never initialize a nil map variable:
// var m map[string]string

// Instead, you can initialize an empty map or use the make keyword to create a map for you:
// var dictionary = map[string]string{}
// or
// var dictionary = make(map[string]string)
// Both approaches create an emtpy hash map and point dictionary at it. Which ensures that you
// will never get a runtime panic.

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

type Dictionary map[string]string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }
