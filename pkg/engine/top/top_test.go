package top

import (
	_ "envmanager/pkg/general/test"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"os"
	"path/filepath"
)

func searchGoMod(dir string) (string, error) {
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
	return searchGoMod(parentDir)
}


var testcase = []struct {
	name string
	session http.Cookie
	expected int
}{
	{
		name: "正常系",
		session: http.Cookie{
			Name:  "session",
			Value: "session",
		},
		expected: http.StatusOK,
	},
	{
		name: "異常系",
		session: http.Cookie{},
		expected: http.StatusOK,
	},
}


func TestIndexGet(t *testing.T) {
	r := gin.New()
	rootPath, _ := searchGoMod(func() string { dir, _ := os.Getwd(); return dir }())

	r.LoadHTMLGlob(rootPath + "/web/templates/*/*.html")
	r.GET("/", indexGet)

	for _, tt := range testcase {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			req := httptest.NewRequest("GET", "/", nil)
			req.AddCookie(&tt.session)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			assert.Equal(tt.expected, w.Code)
		})
	}
}
