package templates_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	templates "github.com/wolfeidau/echo-go-templates"
	"github.com/wolfeidau/echo-go-templates/test/views"
)

func Test_CustomFuncs_AddWithLayout(t *testing.T) {
	assert := require.New(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	rec := httptest.NewRecorder()

	render := templates.NewWithTemplateFuncs(template.FuncMap{
		"getTime2": func() string {
			return time.Now().Format("15:04:05")
		},
	})

	err := render.AddWithLayout(views.Content, "layout2.html", "pages2/*.html")
	assert.NoError(err)

	output := bytes.NewBufferString("")

	c := e.NewContext(req, rec)

	err = render.Render(output, "index2.html", nil, c)
	assert.NoError(err)

	assert.Regexp(`layout index \d{2}:\d{2}:\d{2} `, output.String())
	assert.Equal(200, rec.Result().StatusCode)
}
func Test_AddWithLayout(t *testing.T) {
	assert := require.New(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	rec := httptest.NewRecorder()

	render := templates.New()

	err := render.AddWithLayout(views.Content, "layout2.html", "pages/*.html")
	assert.NoError(err)

	output := bytes.NewBufferString("")

	c := e.NewContext(req, rec)

	err = render.Render(output, "index.html", nil, c)
	assert.NoError(err)

	assert.Regexp(`layout index \d{2}:\d{2}:\d{2} `, output.String())
	assert.Equal(200, rec.Result().StatusCode)
}

func Test_AddWithLayoutAndIncludes(t *testing.T) {
	assert := require.New(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	rec := httptest.NewRecorder()

	render := templates.New()

	err := render.AddWithLayoutAndIncludes(views.Content, "layout.html", "includes/*.html", "pages/*.html")
	assert.NoError(err)

	output := bytes.NewBufferString("")

	c := e.NewContext(req, rec)

	err = render.Render(output, "index.html", nil, c)
	assert.NoError(err)

	assert.Regexp(`header layout index \d{2}:\d{2}:\d{2} footer`, output.String())
	assert.Equal(200, rec.Result().StatusCode)
}

func Test_Add(t *testing.T) {
	assert := require.New(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	rec := httptest.NewRecorder()

	render := templates.New()

	err := render.Add(views.Content, "fragments/*.html")
	assert.NoError(err)

	output := bytes.NewBufferString("")

	c := e.NewContext(req, rec)

	err = render.Render(output, "data.html", nil, c)
	assert.NoError(err)

	assert.Equal("data", output.String())
	assert.Equal(200, rec.Result().StatusCode)
}
