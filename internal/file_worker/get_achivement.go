package fileworker

import (
	"io"
	"io/fs"
	"os"

	"github.com/pkg/errors"
)

type GetAchievementIn struct {
	UserId        uint64
	AchievementId uint64
}

func (f *FileWorker) GetAchievement(params GetAchievementIn) ([]byte, fs.FileInfo, error) {
	dir, fileName := f.GenerateAchievementUrl(GenerateAchievementUrlIn{
		UserId:        params.UserId,
		AchievementId: params.AchievementId,
	})
	file, err := os.Open(dir + fileName)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unable to open file: [os.Open(%s)]", dir+fileName)
	}
	defer file.Close()

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
