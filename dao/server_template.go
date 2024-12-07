package dao

import (
	"gorm.io/gorm"

	"github/leezone/simulator/apistruct"
	"github/leezone/simulator/common/logger"
)

type ServerTemplate struct {
	ID          uint   `gorm:"primaryKey"`
	DeviceType  string // 设备类型:交换机 服务器
	SoftVersion string // 软件版本号
	Vendor      string // 厂商
	Model       string // 型号
	FilePaths   string // 对应模板文件列表
	Comment     string // 备注
}

type ServerTemplateDB interface {
	AutoMigrate()
	CreateServerTemplate(*ServerTemplate) error
	DeleteServerTemplates([]uint) error
	FindServerTemplateById(uint) (ServerTemplate, error)
	FindServerTemplateList(map[string]interface{}, apistruct.Page) ([]ServerTemplate, int64, error)
}

func NewServerTemplateDB(db *gorm.DB) ServerTemplateDB {
	return &DefaultServerTemplateDB{
		db: db,
	}
}

type DefaultServerTemplateDB struct {
	db *gorm.DB
}

func (d *DefaultServerTemplateDB) AutoMigrate() {
	d.db.AutoMigrate(&ServerTemplate{})
}

func (d *DefaultServerTemplateDB) CreateServerTemplate(template *ServerTemplate) error {
	err := d.db.Create(template).Error
	if err != nil {
		logger.Logger.ErrorWithStack("create template error", err, template)
	}
	return err
}

func (d *DefaultServerTemplateDB) DeleteServerTemplates(ids []uint) error {
	err := d.db.Where("id IN ?", ids).Delete(&ServerTemplate{}).Error
	if err != nil {
		logger.Logger.ErrorWithStack("detele templates error", err, ids)
	}
	return err
}

func (d *DefaultServerTemplateDB) FindServerTemplateById(id uint) (ServerTemplate, error) {
	var template ServerTemplate
	err := d.db.Where("id = ?", id).Take(&template).Error
	if err != nil {
		logger.Logger.ErrorWithStack("detele templates error", err, id)
	}
	return template, err
}

func (d *DefaultServerTemplateDB) FindServerTemplateList(condition map[string]interface{}, page apistruct.Page) ([]ServerTemplate, int64, error) {
	var templates []ServerTemplate
	var total int64
	db := d.db.Where("1=1").Where(condition).Count(&total)
	if page.PageCount != 0 || page.PageNum != 0 {
		db = db.Offset(page.PageNum * page.PageCount).Limit(page.PageCount)
	}
	err := db.Find(&templates).Error
	if err != nil {
		logger.Logger.ErrorWithStack("find templates error", err, condition)
	}
	return templates, total, err
}
