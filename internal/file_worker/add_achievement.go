package fileworker

import (
	"os"

	"github.com/pkg/errors"
)

type AddAchievementIn struct {
	AchievementId uint64
	UserId        uint64
	FileContent   []byte
}

func (f *FileWorker) AddAchievement(params AddAchievementIn) error {

	generateAchievementUrlIn := GenerateAchievementUrlIn{
		AchievementId: params.AchievementId,
		UserId:        params.UserId,
	}
	dir, fileName := f.GenerateAchievementUrl(generateAchievementUrlIn)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "unable to create dir: [os.MkdirAll(%v, %v)]", dir, os.ModePerm)
	}

	err = os.WriteFile(dir+fileName, params.FileContent, 0644)
	if err != nil {
		return errors.Wrap(err, "unable to create file: [os.WriteFile]")
	}
	return nil
}
