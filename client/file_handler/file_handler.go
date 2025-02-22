package file_handler

import (
	"bufio"
	"errors"
	"os"

	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/client/interfaces"
)

type fileHandler struct {
	fileName string
	file     *os.File
	writer   *bufio.Writer
}

func NewFileHandler(fileName string) interfaces.FileHandlerInterface {
	return &fileHandler{
		fileName: fileName,
	}
}

func (f *fileHandler) open() error {
	if f.fileName == "" {
		return errors.New("invalid file params")
	}
	file, err := os.OpenFile(f.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	f.file = file
	writer := bufio.NewWriter(f.file)
	f.writer = writer
	return nil
}

func (f *fileHandler) Write(text string) error {
	var err error
	if f.file == nil {
		err = f.open()
		if err != nil {
			return err
		}
	}
	defer f.close()

	if f.writer == nil {
		return errors.New("invalid writer")
	}

	_, err = f.writer.WriteString(text + "\n")
	if err != nil {
		return errors.New("error writing new line: " + f.fileName)
	}
	f.writer.Flush()

	return nil
}

func (f *fileHandler) close() error {
	if f.file != nil {
		err := f.file.Close()
		f.file = nil
		f.writer = nil
		return err
	}
	return nil
}
