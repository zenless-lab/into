package into

func ErrorToString(err error) (string, error) {
	return err.Error(), nil
}
