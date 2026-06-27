// Package dashboard registers the web dashboard assets into the statik
// filesystem so that master/rest.go can serve them via fs.New().
//
// ponytail: registers a placeholder index.html generated at init time. Replace
// by running `statik -src=dist` over a built Vue bundle when the real UI is
// needed; fs.New() succeeds with this minimal zip, so all API tests pass.
package dashboard

import (
	"archive/zip"
	"bytes"

	"github.com/rakyll/statik/fs"
)

const indexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>Groker Dashboard</title>
</head>
<body>
<h1>Groker Dashboard</h1>
<p>Dashboard web assets are not built. See the dashboard repository.</p>
</body>
</html>
`

func init() {
	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	f, err := w.Create("index.html")
	if err != nil {
		panic("dashboard: create index.html in zip: " + err.Error())
	}
	if _, err := f.Write([]byte(indexHTML)); err != nil {
		panic("dashboard: write index.html: " + err.Error())
	}
	if err := w.Close(); err != nil {
		panic("dashboard: close zip: " + err.Error())
	}
	fs.Register(buf.String())
}
