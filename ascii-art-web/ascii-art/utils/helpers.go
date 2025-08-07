package utils

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// константные SHA‑256 оригинальных файлов
var validHashes = map[string]string{
	"standard":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"shadow":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	"thinkertoy": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
}

// IsValidBannerFile вычисляет SHA‑256 баннер‑файла и
// возвращает true, если он совпадает с эталонным хэшем.
func IsValidBannerFile(path string) (bool, error) {
	// Извлекаем имя файла без расширения: standard.txt → standard
	_, file := filepath.Split(path)
	bannerName := strings.TrimSuffix(file, filepath.Ext(file))

	expectedHash, ok := validHashes[bannerName]
	if !ok {
		// Хэш не задан ‑ проверка не требуется
		return true, nil
	}

	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		return false, err
	}

	actualHash := hex.EncodeToString(hasher.Sum(nil))
	return actualHash == expectedHash, nil
}

// GetContentBanner читает баннер-файл по указанному пути.
// Возвращает срез строк (каждая строка файла – отдельный элемент).
func GetContentBanner(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err // ошибка открытия файла
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err // ошибка чтения
	}

	return lines, nil
}

func IsValidAsciiInput(input string) bool {
	for _, char := range input {
		if char != '\n' && (char < 32 || char > 126) {
			return false
		}
	}
	return true
}
