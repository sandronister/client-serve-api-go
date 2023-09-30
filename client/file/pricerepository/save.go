package pricerepository

func (r *repository) saveFile(filename string, content []byte) error {
	if r.fileExists(filename) {
		return r.appendFile(filename, content)
	}

	return r.writeFile(filename, content)
}
