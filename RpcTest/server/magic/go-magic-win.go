package magic

import "fmt"

type Service struct {
}

type Args struct {
	A, B int
}

type Result string

func (t *Service) Multiply(req *Args, resp *Result) error {
	GetFileType("333")
	InitMagic()
	return nil
}

func InitMagic() error {

	return nil
}

// GetFileType 获取文件类型
// 需要头文件:  #include <magic.h>
// 返回值: image/jpeg
func GetFileType(file string) string {
	fmt.Println("调用成功")
	return ""
}
