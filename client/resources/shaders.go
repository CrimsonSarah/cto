package resources

func ReadShader(path ResPath) (string, error) {
	contents, err := ReadNullTerminated(path)

	if err != nil {
		return "", err
	}

	return string(contents), nil
}
