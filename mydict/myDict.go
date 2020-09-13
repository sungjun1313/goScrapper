package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

//Dictionary는 정의된 map을 가르키는 alias 역활을 한다,

var (
	errNotFound   = errors.New("Not Fount")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Can't update non-existing word")
	errCantDelete = errors.New("Can't delete non-existing word")
)

//Search for as word
func (d Dictionary) Search(word string) (string, error) {
	//value는 해당하는 값, exists에서는 찾았는 지 여부에 따른 bool을 담는다,
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

//Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	/*
		if err == errNotFound {
			d[word] = def
		}else if err == nil {
			return errWordExists
		}
	*/
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

//Update a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

//Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		//dictionary 안에 해당 키가 있으면 지우고 없으면 아무것도 하지 않는다. 리턴값은 없다.
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
