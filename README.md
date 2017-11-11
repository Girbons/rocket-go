# Rocket League Stats

[![Go Report Card](https://goreportcard.com/badge/github.com/Girbons/rocket-go)](https://goreportcard.com/report/github.com/Girbons/rocket-go)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

![RLS Logo](/img/rls_partner_horizontal_large.png)

## Setup

```
go get github.com/Girbons/rocket
```

## Usage

Get API_KEY [here](https://developers.rocketleaguestats.com/)

### Example

```go
package main

import (
    "fmt"

    "github.com/Girbons/rocket"
)

func main() {
    client := rocket.NewClient("API_KEY")
    response := client.Player(userId string, platform id int)
    fmt.Println(response.Content)
}
```

## Contribuiting

Feel free to submit Pull Request
