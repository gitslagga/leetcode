# Leetcode
Leetcode solves math problems easily.


# Usage

```go
package main

import (
	solution "github.com/gitslagga/leetcode/twosum"
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

Join the group chat and brush the questions together.

![](https://slagga.oss-cn-shanghai.aliyuncs.com/IMG_2693.JPG)

Welcome to pay attention to my WeChat public account: the game of technology.

![](https://slagga.oss-cn-shanghai.aliyuncs.com/qrcode_for_gh_e8a4b84039c3_258.jpg)