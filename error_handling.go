package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var res Resource
	for {
		res, err = opener()
		if _, ok := err.(TransientError); ok {
			continue
		}
		break
	}
	if err != nil {
		return err
	}
	defer res.Close()
	defer func() {
		if r := recover(); r != nil {
			if frobError, ok := r.(FrobError); ok {
				res.Defrob(frobError.defrobTag)
			}
			err = r.(error)
		}
	}()
	res.Frob(input)
	return err
}
