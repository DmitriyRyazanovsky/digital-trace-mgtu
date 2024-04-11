package fileworker

import (
	"fmt"
	"io/fs"

	"github.com/pkg/errors"
)

const (
	userAvatarPath         = "./user/%d/avatar/avatar"
	userAvatarPathCapacity = uint64(1)
)

func (fw *FileWorker) GenUserAvatarPath(userId uint64) string {
	return fmt.Sprintf(userAvatarPath, userId)
}

func (fw *FileWorker) UserAvatarWrite(userId uint64, content []byte) error {
	err := fw.write(fw.GenUserAvatarPath(userId), content)
	if err != nil {
		err = errors.Wrap(err, "[fw.write(fw.genUserAvatarPath(userId, prefix), content, userAvatarPathCapacity)]")
		return err
	}

	return nil
}

func (fw *FileWorker) UserAvatarRead(userId uint64, prefix string) ([]byte, fs.FileInfo, error) {
	content, info, err := fw.read(fw.GenUserAvatarPath(userId))
	if err != nil {
		err = errors.Wrap(err, "[fw.read(fw.genUserAvatarPath(userId, prefix))]")
		return nil, nil, err
	}

	return content, info, nil
}
