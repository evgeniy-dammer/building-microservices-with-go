package files

import (
	"golang.org/x/xerrors"
	"io"
	"os"
	"path/filepath"
)

// Local is an implementation of the Storage interface which works with the local disk on the current machine
type Local struct {
	maxFileSize int
	basePath    string
}

// NewLocal creates a new Local filesystem with the given base path basePath is the base directory to save
// files to maxSize is the max number of bytes that a file can be
func NewLocal(basePath string) (*Local, error) {
	p, err := filepath.Abs(basePath)

	if err != nil {
		return nil, err
	}

	return &Local{basePath: p}, nil
}

// fullPath returns the absolute path
func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basePath, path)
}

// Save the contents of the Writer to the given path. Path is a relative path, basePath will be appended
func (l *Local) Save(path string, content io.Reader) error {
	// get the full path of the file
	fullPath := l.fullPath(path)

	// get the directory and make sure it exists
	directory := filepath.Dir(fullPath)
	err := os.Mkdir(directory, os.ModePerm)

	if err != nil {
		return xerrors.Errorf("Unable to create directory: %w", err)
	}

	// if the file exists - delete it
	_, err = os.Stat(fullPath)
	if err == nil {
		err = os.Remove(fullPath)
		if err != nil {
			return xerrors.Errorf("Unable to delete file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		// if this is anything other than not exists
		return xerrors.Errorf("Unable to get file info: %w", err)
	}

	// create a new file at the path
	file, err := os.Create(fullPath)

	if err != nil {
		return xerrors.Errorf("Unable to create file: %w", err)
	}

	defer file.Close()

	// write the content to the file but not greater than max bytes
	_, err = io.Copy(file, content)

	if err != nil {
		return xerrors.Errorf("Unable to write to file: %w", err)
	}

	return nil
}

// Get opens the file at the given path and returns a Reader
func (l *Local) Get(path string) (*os.File, error) {
	// get the full path of the file
	fullPath := l.fullPath(path)

	// open the file
	file, err := os.Open(fullPath)

	if err != nil {
		return nil, xerrors.Errorf("Unable to open file: %w", err)
	}
	return file, nil
}
