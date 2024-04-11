package fileworker

import (
	"os"

	"github.com/pkg/errors"
)

type GetAchievementUrlIn struct {
	UserId        uint64
	AchievementId uint64
}

func (f *FileWorker) GetAchievementUrl(params *GetAchievementUrlIn) (string, error) {
	dir, fileName := f.GenerateAchievementUrl(GenerateAchievementUrlIn{
		UserId:        params.UserId,
		AchievementId: params.AchievementId,
	})

	if _, err := os.Stat(dir + fileName); err != nil {
		if os.IsNotExist(err) {
			return "", errors.Wrapf(err, "unable to find file: [os.Stat(%s)]", dir+fileName)
		}

		return "", errors.Wrapf(err, "unable to check file: [os.Stat(%s)]", dir+fileName)
	}

	return dir + fileName, nil
}
