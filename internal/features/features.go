package features

import (
	"crypto/rand"
	"os"
	"path/filepath"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func ConvertInToPqInt64Array(value []*int64) *pq.Int64Array {
	arr := pq.Int64Array{}
	for _, v := range value {
		arr = append(arr, *v)
	}
	return &arr
}

func CreateDirectoriesIfNotExist(filePath string) error {
	dir := filepath.Dir(filePath)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			err = errors.Wrap(err, "[os.MkdirAll(dir, 0755)]")
			return err
		}
	} else {
		err = errors.Wrap(err, "[os.Stat(dir)]")
		return err
	}
	return nil
}

func GenerateRandomBytes(keySize int) ([]byte, error) {
	key := make([]byte, keySize)

	n, err := rand.Read(key)
	if err != nil {
		return []byte{}, errors.Wrap(err, "[rand.Read(key)]")
	}
	if n != keySize {
		return []byte{}, errors.New("[n != keySize]")
	}

	return key, nil
}

func CreateFileIfNotExists(filePath string) (bool, error) {
	// Проверяем существование файла
	var exist bool = true
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exist = false
		// Создаем директорию, если она не существует
		err := os.MkdirAll(filepath.Dir(filePath), 0755)
		if err != nil {
			return exist, err
		}

		// Создаем файл
		file, err := os.Create(filePath)
		if err != nil {
			return exist, err
		}
		defer file.Close()
	}

	return exist, nil
}

func WriteContentInFile(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadContentFromFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "[os.ReadFile(filename)]")
	}

	return data, nil
}

func CountFilesInDirectoryFile(filePath string) (int, error) {
	dir := filepath.Dir(filePath)
	fileCount := 0

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && path != filePath {
			fileCount++
		}
		return nil
	})

	return fileCount, err
}
