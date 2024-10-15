Thiết lập dự án go  
Video 1

1. Tạo một dự án mới:

- go mod init ten_du_an

2. Xây dựng thư mục

- cmd
  - cli
  - cronjob
  - server - chứa thư mục khởi tạo dự án.
- internal - mã cục bộ của dự án
  - controller
  - models
  - service
  - repo - dto
  - router
  - middleware
  - initialize - khởi tạo database
- pkg
  - utils
  - setting
  - logger
- docs - lưu trữ tài liệu nội bộ
- tests
- third_party - nếu có
- scripts - các file bash chạy kịch bản
- migrations - lưu trữ các file sql để khởi tạo db
- response - tạo ra các file quản trị những thông tin trả về
- global - lưu trữ file dùng chung toàn dự án
- configs - cài đặt cho dự án

- .gitignore
- go.mod file thông tin dự án

---

video 2
Lựa chọn gin là framework.

1. Chuẩn bị

- Tải thư viện: https://github.com/gin-gonic/gin
- go get -u github.com/gin-gonic/gin

sau đó go sẽ tự động tải về một loạt các thư viện, có thể kiểm tra ở file go.mod

go tự động tạo cho ta file go.sum đơn giản là ghi log nhật ký tải thôi.

Quay lại trang github của git, copy tạm một demo của gin.

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  // r.Run(":9999") sửa port ở đây
}
```

Tìm hiểu một chút:

- r := gin.Default() giữ control và click vào ta sẽ thấy hàm này khởi tạo sẵn cho ta logger và recovery.

Khởi tạo một handler:

```go

func main() {
  r := gin.Default()
  r.GET("/ping", Pong)
  r.Run()
}

//Tạo một handler

func Pong (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

```

Tạo nhóm api

```go

func main() {
  r := gin.Default()

	//tạo nhóm api
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", Pong)
		v1.GET("/hello", Pong)
		v1.POST("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.PUT("/ping", Pong)
		v1.DELETE("/ping", Pong)
		v1.HEAD("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
	}

  r.GET("/ping", Pong)
  r.Run()
}
```

Tạo ra file routers.go trong thư mục routers. Chuyển routers qua bên file đó. Lưu ý vấn đề: chỉ chữ hoa mới được export

```go
//routers/routers.go
package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", Pong)
		v1.GET("/hello", Pong)
		v1.POST("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.PUT("/ping", Pong)
		v1.DELETE("/ping", Pong)
		v1.HEAD("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
	}
	return r
}

func Pong (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"users":[]string{
			"1",
			"2",
			"3",
		 },
	})
}


//cmd/server/main.go
package main

import (
	"go-tip/internal/routers"
)

func main() {
  r := routers.NewRouter()

  r.Run()
}

```

Tạo file tests/test.http để tạo một http resquest hoặc sử dụng cmd: curl http://localhost:8080/v1/ping

nếu tạo file thì nhớ cài extension http rest, file có nội dung như sau:

```http
GET  http://localhost:8080/v1/ping
```

Video 3: Hướng dẫn xây dựng api theo mô hình mvc

Yêu cầu: phải hiểu con trỏ, struct. (Học kỹ sau này)
controller -> service -> repo -> models -> db
![mvc](docs/mvc.png)
