// Package erratum implements simple routing to handle error
package erratum

// Use perform operations on resource and returns error on failure
func Use(ro ResourceOpener, input string) (err error) {
	res, err := ro()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		res, err = ro()
	}
	defer res.Close()

	defer func() {
		if r := recover(); r != nil {
			if fe, ok := r.(FrobError); ok {
				res.Defrob(fe.defrobTag)
			}
			err = r.(error)
		}
	}()

	res.Frob(input)

	return nil
}
