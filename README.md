# Movie API

Movie API Application is an backend app providing movie data to the client.

## Gin Web Framework

Gin is a web framework written in Go (Golang) like Express in Node.js. [Visit official website here](https://github.com/gin-gonic/gin).

## Installation

1. Initialize a project with go module

```
go mod init <project-name>
```

Then, it will create `go.mod` file in your project directory.

2. Install Gin web framework

```go
go get github.com/gin-gonic/gin
```

After that, `go.sum` file will be created. It contain all libraries you include in your project. Now, you can import gin web framework in your main file.

You can use `go build` after `go mod` instead of `go get` every packages. But first, you need to import packages in the files before `go build`

3. Add `server.go` to start server with gin

```go
package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func main() {
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "OK",
		})
	})

	server.Run(":8080")
}
```

4. Run server

```
go run server.go
```

## How to remove packages

```
go get <package_name>@none
```

## Gomon

Gomon aim to reproduce the behavior of nodemon for Go. It can re-build everytime when the code is changed.

```
go get -u github.com/TonPC64/gomon

gomon .
```
