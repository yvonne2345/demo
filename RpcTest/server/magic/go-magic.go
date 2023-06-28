//go:build linux
// +build linux

package magic

import (
	"fmt"
	"github.com/rakyll/magicmime"
	"netvine.com/firewall/server/global"
)

func InitMagic() error {
	if err := magicmime.Open(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_SYMLINK | magicmime.MAGIC_ERROR); err != nil {
		return err
	}
	return nil
}

func UnInitMagic() {
	magicmime.Close()
}

// GetFileType 获取文件类型
// 需要头文件:  #include <magic.h>
// 返回值: image/jpeg
func GetFileType(file string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("文件类型检测捕获异常", err)
		}
	}()
	mimetype, err := magicmime.TypeByFile(file)
	if err != nil {
		global.NETVINE_LOG.Info("获取文件类型失败，文件路径为：" + file)
		return ""
	}
	return mimetype
}
