package main

import "errors"

type Dictionary map[string]string

type DictionaryError string

var (
	ErrorNotFound       = errors.New("could not find the word you were looking for")
	ErrorWordExists     = errors.New("cannot add word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (dictionary Dictionary) Add(key string, value string) error {
	_, err := dictionary.Search(key)

	switch err {
	case ErrorNotFound:
		dictionary[key] = value
	case nil:
		return ErrorWordExists
	default:
		return err
	}
	return nil
}

func (dictionary Dictionary) Update(key string, value string) error {
	_, err := dictionary.Search(key)

	switch err {
	case ErrorNotFound:
		return ErrWordDoesNotExist
	case nil:
		dictionary[key] = value
	default:
		return err
	}
	return nil
}

func (e DictionaryError) Error() string {
	return string(e)
}
