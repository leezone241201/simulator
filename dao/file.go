package dao

type FileDB interface {
	CreateFile() error
	DeleteFile() error
}

type DefaultFileDB struct {
}
