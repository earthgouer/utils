package utils

import (
    "embed"
    "log"
)

//go:embed internal/assets/data.zip
var zipFile embed.FS // 声明一个嵌入的文件系统变量

func main() {
    // 现在你可以直接读取 zipFile 这个变量，它包含了 data.zip 的内容
    data, err := zipFile.ReadFile("remote-verify.zip")
    if err != nil {
        log.Fatal(err)
    }
    // 然后用 archive/zip 包处理 data 字节切片...
    log.Printf("ZIP file size: %d bytes\n", len(data))
}
