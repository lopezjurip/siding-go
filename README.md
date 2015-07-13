# siding-go

## Installation

```shell
go install github.com/mrpatiwi/siding-go
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/mrpatiwi/siding-go"
)

func main() {
	s := siding.Siding{Username: "username", Password: "password"}
	var courseID uint = 2000

	resp, err := s.Announcements(courseID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	html, err := siding.ReadResponse(resp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(html)
}

```
