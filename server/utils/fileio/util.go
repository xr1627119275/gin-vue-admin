package fileio

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLines 逐行读取文件，并返回一个字符串数组
func ReadLines(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("读取文件时出错: %w", err)
	}
	return lines, nil
}

// WriteString 将字符串写入文件，覆盖原文件内容
func WriteString(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("无法创建文件: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("写入文件时出错: %w", err)
	}
	return nil
}
