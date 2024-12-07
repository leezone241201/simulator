package service

import (
	"io"

	"github/leezone/simulator/dao"
	"github/leezone/simulator/svc"
)

func AddServerTmplate(file io.Reader, paths string) error {
	// TODO,根据file内容读取文件信息,填充厂商等信息
	var template dao.ServerTemplate
	template.FilePaths = paths
	template.Comment = "测试备注信息"
	template.DeviceType = "服务器"
	template.Model = "CPU 5900x"
	template.Vendor = "SuperMicro"
	template.SoftVersion = "v1.0"
	return svc.Svc.ServerTemplateDB.CreateServerTemplate(&template)
}
