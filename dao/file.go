package dao

import (
	"time"

	"gorm.io/gorm"

	"github/leezone/simulator/apistruct"
	"github/leezone/simulator/common/logger"
)

// 可以删除文件夹和文件
type File struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdateAt  time.Time
	Path      string
}

type FileDB interface {
	AutoMigrate()
	CreateFile(*File) error
	DeleteFile([]int) error
	FindFileList(map[string]interface{}, apistruct.Page) ([]File, int64, error)
}

func NewFileDB(db *gorm.DB) FileDB {
	return &DefaultFileDB{
		db: db,
	}
}

type DefaultFileDB struct {
	db *gorm.DB
}

func (d *DefaultFileDB) AutoMigrate() {
	d.db.AutoMigrate(&File{})
}

func (d *DefaultFileDB) CreateFile(file *File) error {
	err := d.db.Create(file).Error
	if err != nil {
		logger.Logger.ErrorWithStack("create file error", err, file)
	}
	return err
}

func (d *DefaultFileDB) DeleteFile(ids []int) error {
	err := d.db.Where("id IN ?", ids).Delete(&File{}).Error
	if err != nil {
		logger.Logger.ErrorWithStack("delete files error", err, ids)
	}
	return err
}

func (d *DefaultFileDB) FindFileList(condition map[string]interface{}, page apistruct.Page) ([]File, int64, error) {
	var files []File
	var total int64
	db := d.db.Where("1=1").Where(condition).Count(&total)
	if page.PageCount != 0 || page.PageNum != 0 {
		db = db.Offset(page.PageNum * page.PageCount).Limit(page.PageCount)
	}
	err := db.Find(&files).Error
	if err != nil {
		logger.Logger.ErrorWithStack("find files error", err, condition)
	}
	return files, total, err
}
