# cn

Package for connection to DB, standart checking errors and load JSON configuration file.

Installation

$ go get github.com/loct7/cn

Example:
```go
func main() {
  var cfg cn.Configuration
	cfg.Load("")

  var dbWork cn.DatabaseWork
	dbWork.Init(&cfg)
	defer dbWork.DB.Close()
  }
  ```
