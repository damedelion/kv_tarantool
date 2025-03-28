package interrors

var e1 error = &DuplicateKey{}
var e2 error = &KeyNotFound{}

type DuplicateKey struct {
	Err string
}

type KeyNotFound struct {
	Err string
}

func (e *DuplicateKey) Error() string {
	return e.Err
}

func (e *KeyNotFound) Error() string {
	return e.Err
}
