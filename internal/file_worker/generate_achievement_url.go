package fileworker

import "fmt"

type GenerateAchievementUrlIn struct {
	AchievementId uint64
	UserId        uint64
}

func (f *FileWorker) GenerateAchievementUrl(params GenerateAchievementUrlIn) (string, string) {
	dir := fmt.Sprintf("./%s/%d/%s/", f.userDirPath, params.UserId, f.achievementDirPath)
	fileName := fmt.Sprintf("%d.png", params.AchievementId)
	return dir, fileName
}
