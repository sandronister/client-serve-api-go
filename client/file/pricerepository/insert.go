package pricerepository

func (r *repository) Insert(filename string, content []byte) error {
	err := r.saveFile(filename, content)
	if err != nil {
		return err
	}

	return nil
}
