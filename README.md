# Leetcode
Leetcode solves math problems easily.


# Install
```shell
cd your-project

go get -u github.com/gitslagga/leetcode
```


# Usage

```go
package main

import (
	solution "github.com/gitslagga/leetcode/1.two-sum"
	"log"
	"reflect"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	expectResult := []int{0, 1}
	result := solution.TwoSum(nums, target)
	if reflect.DeepEqual(result, expectResult) == false {
		log.Fatalf("return wrong result, expect:%d, got:%d", expectResult, result)
	}
}
```


# Leetcode Team
Welcome to pay attention to my WeChat public account: the game of technology.

![](https://slagga.top/images/media-platform.png)
