package pricerepository

func (r *repository) Insert(filename, content string) error {
	err := r.saveFile(filename, content)
	if err != nil {
		return err
	}

	return nil
}
