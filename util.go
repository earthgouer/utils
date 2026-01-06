package utils

import (
	"embed"
)

//go:embed internal/assets/data.zip
var zipFile embed.FS // 声明一个嵌入的文件系统变量

func main() {
	// 现在你可以直接读取 zipFile 这个变量，它包含了 data.zip 的内容
	zipFile.ReadFile("remote-verify.zip.000")
	zipFile.ReadFile("remote-verify.zip.001")
	zipFile.ReadFile("remote-verify.zip.002")
	zipFile.ReadFile("remote-verify.zip.003")
	zipFile.ReadFile("remote-verify.zip.004")

}
