# echo-go-templates

This project provides some template helpers with the Go [IOFS](https://pkg.go.dev/io/fs) and [HTML templates](https://pkg.go.dev/html/template) packages for use with [echo](https://echo.labstack.com).

# Usage

The following code example sets up views with a common layout, and includes for a header and footer. The names of files in the pages directory are used to render the templates.

```go
	e := echo.New()

	render := templates.New()

	err := render.AddWithLayoutAndIncludes(views.Content, "layout.html", "includes/*.html", "pages/*.html")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load layout")
	}

	e.Renderer = render
```

This expects a directory structure such as:

```
views
  layout.html
  includes
    footer.html
    header.html
  views.go
  pages
    index.html
```

Note take a look at the test view templates in [test/views](test/views/)

In this structure `views.go` uses the embed feature in Go to include all the templates in the Go binary.

```go
package views

import "embed"

//go:embed pages/* includes/* layout.html
var Content embed.FS
```

To render the index page in this hierarchy.

```go
    return c.Render(http.StatusOK, "index.html", nil)
```

# Links

* https://francoposa.io/resources/golang/golang-templates-1/
* https://philipptanlak.com/web-frontends-in-go/

# License

This application is released under Apache 2.0 license and is copyright [Mark Wolfe](https://www.wolfe.id.au).
