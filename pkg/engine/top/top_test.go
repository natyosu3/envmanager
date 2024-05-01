package top

import (
	_ "envmanager/pkg/general/test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"envmanager/pkg/test"
	"os"
)




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
	rootPath, _ := test.SearchGoMod(func() string { dir, _ := os.Getwd(); return dir }())

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
