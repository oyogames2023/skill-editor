package storage

import (
	"github.com/vrischmann/userdir"
	"os"
	"path"
)

type LocalStorage struct {
	ConfigPath string
}

func NewLocalStorage(filename string) *LocalStorage {
	return &LocalStorage{
		ConfigPath: path.Join(userdir.GetConfigHome(), "SkillEditor", filename),
	}
}

func (l *LocalStorage) Load() ([]byte, error) {
	data, err := os.ReadFile(l.ConfigPath)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (l *LocalStorage) Store(data []byte) error {
	dir := path.Dir(l.ConfigPath)
	if err := ensureDirExists(dir); err != nil {
		return err
	}
	if err := os.WriteFile(l.ConfigPath, data, 0775); err != nil {
		return err
	}
	return nil
}

func ensureDirExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err = os.Mkdir(path, 0775); err != nil {
			return err
		}
	}
	return nil
}
