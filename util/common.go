package util

import (
	"fmt"
	"ginorm/config"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/net/html/charset"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetAbsPath(relPath string) string {
	if config.Env.Platform == config.Windows {
		return config.Env.RootDir + strings.Replace(relPath, "\\", "/", -1)
	} else {
		return config.Env.RootDir + strings.Replace(relPath, "\\", "/", -1)
	}
}

// GenerateTraceID 生成一个随机的 Trace ID
func GenerateTraceID() string {
	traceID, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("failed to generate uuid: %v", err)
	}
	return traceID.String()
}

// GetTraceID 从 context 中获取 Trace ID
func GetTraceID(c *gin.Context) string {
	if traceID, ok := c.Value("trace_id").(string); ok {
		return traceID
	}
	return "unknown"
}

func ReadBodyWithCharset(resp *http.Response) (string, error) {
	reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// FileExists 检查文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// CreateFileWithDirs 创建文件及其父目录（如果不存在）
func CreateFileWithDirs(filePath string) error {
	// 获取文件所在目录路径
	dir := filepath.Dir(filePath)

	// 创建父目录（如果不存在）
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	// 创建文件（如果不存在）
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("创建文件失败: %w", err)
	}
	defer file.Close()

	return nil
}
