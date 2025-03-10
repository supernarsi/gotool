package gotool

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// FileExists 检查文件是否存在
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirExists 检查目录是否存在
func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// CreateDirIfNotExist 如果目录不存在则创建
func CreateDirIfNotExist(path string) error {
	if !DirExists(path) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

// ReadFileToString 安全地读取文件内容到字符串
func ReadFileToString(filename string) (string, error) {
	if !FileExists(filename) {
		return "", fmt.Errorf("file not found: %s", filename)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// WriteStringToFile 安全地将字符串写入文件
func WriteStringToFile(filename, content string) error {
	dir := filepath.Dir(filename)
	if err := CreateDirIfNotExist(dir); err != nil {
		return err
	}

	return os.WriteFile(filename, []byte(content), 0644)
}

// AppendStringToFile 将字符串追加到文件末尾
func AppendStringToFile(filename, content string) error {
	dir := filepath.Dir(filename)
	if err := CreateDirIfNotExist(dir); err != nil {
		return err
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}

// ReadLines 读取文件的所有行
func ReadLines(filename string) ([]string, error) {
	if !FileExists(filename) {
		return nil, fmt.Errorf("file not found: %s", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// WriteLines 将字符串切片写入文件，每行一个字符串
func WriteLines(filename string, lines []string) error {
	dir := filepath.Dir(filename)
	if err := CreateDirIfNotExist(dir); err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

// CopyFile 复制文件
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destDir := filepath.Dir(dst)
	if err := CreateDirIfNotExist(destDir); err != nil {
		return err
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// GetFileSize 获取文件大小（字节数）
func GetFileSize(filename string) (int64, error) {
	if !FileExists(filename) {
		return 0, fmt.Errorf("file not found: %s", filename)
	}

	info, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}

	return info.Size(), nil
}

// ListFilesInDir 列出目录中的所有文件（不包括子目录）
func ListFilesInDir(dirPath string) ([]string, error) {
	if !DirExists(dirPath) {
		return nil, fmt.Errorf("directory not found: %s", dirPath)
	}

	var files []string
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, filepath.Join(dirPath, entry.Name()))
		}
	}

	return files, nil
}

// WalkDir 递归遍历目录，对每个文件执行回调函数
func WalkDir(dirPath string, callback func(path string, info os.FileInfo) error) error {
	if !DirExists(dirPath) {
		return fmt.Errorf("directory not found: %s", dirPath)
	}

	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return callback(path, info)
		}
		return nil
	})
}

// FindFilesByExt 查找指定目录下所有具有特定扩展名的文件
func FindFilesByExt(dirPath string, ext string) ([]string, error) {
	if !DirExists(dirPath) {
		return nil, fmt.Errorf("directory not found: %s", dirPath)
	}

	ext = strings.ToLower(ext)
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	var files []string
	err := WalkDir(dirPath, func(path string, info os.FileInfo) error {
		if strings.ToLower(filepath.Ext(path)) == ext {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// ReadJSONFile 从JSON文件读取数据到结构体
func ReadJSONFile(filename string, v interface{}) error {
	if !FileExists(filename) {
		return fmt.Errorf("file not found: %s", filename)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

// WriteJSONFile 将结构体数据写入JSON文件
func WriteJSONFile(filename string, v interface{}, pretty bool) error {
	dir := filepath.Dir(filename)
	if err := CreateDirIfNotExist(dir); err != nil {
		return err
	}

	var data []byte
	var err error

	if pretty {
		data, err = json.MarshalIndent(v, "", "  ")
	} else {
		data, err = json.Marshal(v)
	}

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// IsEmptyDir 检查目录是否为空
func IsEmptyDir(dirPath string) (bool, error) {
	if !DirExists(dirPath) {
		return false, fmt.Errorf("directory not found: %s", dirPath)
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return false, err
	}

	return len(entries) == 0, nil
}

// TempFile 创建临时文件并返回其文件句柄
func TempFile(dir, pattern string) (*os.File, error) {
	if dir != "" && !DirExists(dir) {
		if err := CreateDirIfNotExist(dir); err != nil {
			return nil, err
		}
	}

	return os.CreateTemp(dir, pattern)
}

// SafeRemove 安全删除文件或空目录
func SafeRemove(path string) error {
	if !FileExists(path) && !DirExists(path) {
		return nil // 文件或目录不存在，视为已删除
	}

	// 如果是目录，确保为空
	if DirExists(path) {
		isEmpty, err := IsEmptyDir(path)
		if err != nil {
			return err
		}
		if !isEmpty {
			return errors.New("cannot remove non-empty directory")
		}
	}

	return os.Remove(path)
}