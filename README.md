# nike

`nike` is a simple library to run functions in parallel and return the accumulated errors, if any. It will create as many workers as there are CPUs at runtime.

## Usage example

```golang
package main

import "github.com/cloudops/nike"

func main() {
    fns := []func() error {
        func() error {
            //do something here
            return fmt.Errorf("uninmplemented yet #1")
        },
        func() error {
            //do something here
            return fmt.Errorf("uninmplemented yet #2")
        },
    }

    errs := nike.JustDoIt(fns)
    if len(errs) > 0 {
        for i := range errs {
            fmt.Println(errs[i])
        }
    }
}

```