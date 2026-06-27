# Dashboard for Groker recommender system

[![build](https://github.com/linkerlin/dashboard/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/linkerlin/dashboard/actions/workflows/build.yml)

An admin dashboard for the Groker recommender system derived from [shards-dashboard-vue](https://github.com/DesignRevision/shards-dashboard-vue).

![](assets/dashboard.png)

## Quick Start

1. Install Node 18+ and `pnpm`.
2. Install dependencies by running `pnpm install --registry=https://registry.npmjs.org/`.
3. Run `pnpm serve` to start the local development server.

> - The build might fail if you are using newer versions of Node.
> - [Node Version Manager](http://nvm.sh/) is recommended for managing multiple Node versions on a single machine.

## Usage

The dashboard ships prebuilt assets embedded via `go:embed`. Add it as a dependency:

```bash
go get github.com/linkerlin/dashboard
```

Import it (side-effect) so the assets are registered into the statik filesystem:

```go
import (
  "github.com/rakyll/statik/fs"

  _ "github.com/linkerlin/dashboard"
)

  // ...

  statikFS, err := fs.New()
  if err != nil {
    log.Fatal(err)
  }

  // Serve the contents over HTTP.
  http.Handle("/", http.FileServer(statikFS))
  http.ListenAndServe(":8080", nil)
```
