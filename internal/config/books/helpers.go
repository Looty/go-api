package books

import (
	"errors"
	"strconv"
)

func GetBookInstanceById(id string) (*Book, error) {

	idToInt, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	if idToInt <= 0 {
		return nil, errors.New("book id must be higher than 0")
	}

	for i, b := range Bookvar {
		if b.ID == id {
			return &Bookvar[i], nil
		}
	}

	return nil, errors.New("book not found")
}
