package top

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"os"
	"path/filepath"
	"envmanager/pkg/model"
	"encoding/json"
    "github.com/stretchr/testify/mock"
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


type SessionMock struct {
    mock.Mock
}

func (s *SessionMock) GetSession(c *gin.Context, name string) []byte {
    args := s.Called(c, name)
    return args.Get(0).([]byte)
}

func TestIndexGet1(t *testing.T) {
    gin.SetMode(gin.DebugMode)

    // Create a session mock
    sessionMock := new(SessionMock)

    // Create a gin context with a buffer to record the response
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Mock the session data
    sessionData := model.Session_model{Logined: true}
    sessionDataBytes, _ := json.Marshal(sessionData)
    sessionMock.On("GetSession", c, "session").Return(sessionDataBytes)

	c.SetCookie("session", "session", 0, "/", "", false, true)

    // Call the function with the mocked context and session
    indexGet(c)

    // Check the response
    assert.Equal(t, http.StatusOK, w.Code)

    var response map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    assert.Equal(t, sessionData.Logined, response["IsAuthenticated"])
}


func TestIndexGet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	rootPath, _ := searchGoMod(func() string { dir, _ := os.Getwd(); return dir }())
	
	r.LoadHTMLGlob(rootPath + "/web/templates/*/*.html")
	r.GET("/", indexGet)

	// クッキーをセット
	cookie := &http.Cookie{
		Name:  "session",
		Value: "session",
	}
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(cookie)


	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}