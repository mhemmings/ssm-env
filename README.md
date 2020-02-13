# ssm-env

A Go package to set an environment using AWS SSM.

## Usage

Assuming you have parameters set in SSM using:

`$ aws ssm put-parameter --name=/ssm/path/FOO --value=secretfoo --type SecureString`

`$ aws ssm put-parameter --name=/ssm/path/BAR --value=secretbar --type SecureString`


```go
package main

import (
  "fmt"
  "os"

  "github.com/mhemmings/ssm-env"
)

func main() {
  err := ssm.Parse("/ssm/path")
  if err != nil {
    panic(err)
  }

  // Now the environment has been setup.

  foo := os.Getenv("FOO")
  bar := os.Getenv("BAR")

  fmt.Println(foo) // Prints: "secretfoo"
  fmt.Println(bar) // Prints: "secretbar"
}
```
