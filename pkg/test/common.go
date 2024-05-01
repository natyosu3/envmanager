package test

import (
	"fmt"
	"os"
	"path/filepath"
)


func SearchGoMod(dir string) (string, error) {
	modPath := filepath.Join(dir, "go.mod")
	_, err := os.Stat(modPath)
	if err == nil {
		// go.mod ファイルが見つかった場合、そのディレクトリのパスを返す
		return filepath.ToSlash(dir), nil
	}
	// 親ディレクトリへのパスを取得
	parentDir := filepath.Dir(dir)
	// 親ディレクトリがルートディレクトリである場合や
	// 親ディレクトリと指定されたディレクトリが同じ場合は、
	// go.mod ファイルが見つからなかったことを示すエラーを返す
	if parentDir == dir || parentDir == "" {
		return "", fmt.Errorf("go.mod ファイルが見つかりませんでした")
	}
	// 親ディレクトリに対して再帰的に探索を行う
	return SearchGoMod(parentDir)
}