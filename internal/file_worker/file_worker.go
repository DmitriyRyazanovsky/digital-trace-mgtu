package fileworker

import (
	"io"
	"io/fs"
	"mgtu/digital-trace/main-backend-service/internal/features"
	"mgtu/digital-trace/main-backend-service/internal/features/logging"
	"os"

	"github.com/pkg/errors"
)

type FileWorker struct {
	log logging.Logger

	achievementDirPath string
	userDirPath        string
}

func NewFileWorker(log logging.Logger) *FileWorker {
	return &FileWorker{
		log:                log,
		achievementDirPath: "achievement",
		userDirPath:        "user",
	}
}

func (fw *FileWorker) write(filePath string, content []byte) error {
	_, err := features.CreateFileIfNotExists(filePath)
	if err != nil {
		err = errors.Wrap(err, "[features.CreateFileIfNotExists(filePath)]")
		return err
	}

	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		return errors.Wrap(err, "unable to create file: [os.WriteFile]")
	}

	return nil
}

func (fw *FileWorker) read(filePath string) ([]byte, fs.FileInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unable to open file: [os.Open(%s)]", filePath)
	}

	fileStat, err := file.Stat()
	if err != nil {
		return nil, nil, errors.Wrap(err, "[file.Stat()]")
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, errors.Wrap(err, "[io.ReadAll(file)]")
	}

	return fileBytes, fileStat, nil
}
