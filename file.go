package helpers

import (
	"path/filepath"
	"strings"
)

func GetFileNameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}
