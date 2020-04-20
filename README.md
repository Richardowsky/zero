# Substrings

## Prerequisites:
You need to install - [go 1.13](https://golang.org/dl/)

## How to test:
1. Clone this repo: `git clone https://github.com/Richardowsky/zero.git`
2. `make test`

## How to use:
```go
package example

import zero "zero/src"

func example()  {
	
var n = 4
var array = []int{1, 1, 2, 2}
	
	zero.CheckZero(n, array)
}

```