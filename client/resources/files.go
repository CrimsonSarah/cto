package resources

import "os"

// Used to differentiate between relative resource paths and OS file
// paths.
type ResPath string

// Turns a relative resource path like "shaders/cards/frag.glsl" into the
// file path where that is supposed to be.
func GameFilePath(path ResPath) string {
	// TODO: Prepend the installation path somehow.
	// Currently resource files will be looked for in the CWD
	return string(path)
}

// Reads bytes from a file and sticks a \0 at the end.
func ReadNullTerminated(path ResPath) ([]uint8, error) {
	filePath := GameFilePath(path)
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()

	if err != nil {
		return nil, err
	}

	size := stat.Size()

	// +1 for null terminator
	buffer := make([]uint8, size+1)

	bytesLeft := size

	for bytesLeft > 0 {
		bytesRead, err := file.Read(buffer)

		if err != nil {
			return nil, err
		}

		bytesLeft -= int64(bytesRead)
	}

	// Null-terminate
	buffer[size] = 0

	return buffer, nil
}
