package myDict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("not found")
var errExists = errors.New("already exists key")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound

}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	if err == errNotFound {
		d[key] = value
		return nil
	}
	return errExists
}

func (d Dictionary) Update(key string, value string) error {
	err := d.Add(key, value)
	if err != errExists {
		d[key] = value
		return nil
	}
	return err
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
