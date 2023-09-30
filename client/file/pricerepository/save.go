package pricerepository

func (r *repository) saveFile(filename string, content string) error {
	if r.fileExists(filename) {
		return r.appendFile(filename, content)
	}

	return r.writeFile(filename, content)
}
