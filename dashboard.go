// Package dashboard embeds the built Vue dashboard assets and registers them
// into the statik filesystem so master/rest.go can serve the UI via fs.New().
//
// The dist/ tree is produced by `pnpm build` in this repository and committed
// so that consumers fetching this module get the prebuilt assets via go:embed.
package dashboard

import (
	"archive/zip"
	"bytes"
	"embed"
	"io"
	"io/fs"
	"strings"

	statikfs "github.com/rakyll/statik/fs"
)

//go:embed all:dist
var dist embed.FS

func init() {
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	_ = fs.WalkDir(dist, "dist", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		data, rerr := dist.ReadFile(path)
		if rerr != nil {
			return rerr
		}
		// statik/fs stores each entry under "/"+name, so emit relative paths.
		rel := strings.TrimPrefix(path, "dist/")
		w, werr := zw.Create(rel)
		if werr != nil {
			return werr
		}
		_, werr = io.Copy(w, bytes.NewReader(data))
		return werr
	})
	_ = zw.Close()
	statikfs.Register(buf.String())
}
