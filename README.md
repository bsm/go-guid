# GUID [![Build Status](https://travis-ci.org/bsm/guid.png?branch=master)](https://travis-ci.org/bsm/guid)

Simple, thread-safe MongoDB style GUID generator.

### Example

```go
package main

import (
  "encoding/hex"
  "log"

  "github.com/bsm/guid"
)

func main() {
  // Create a new 12-byte guid
  g1 := guid.New96()
  fmt.Println(hex.EncodeToString(g1))

  // Create a new 16-byte guid
  g2 := guid.New128()
  fmt.Println(hex.EncodeToString(g2))
}
```

### Licence

```
Copyright (c) 2015 Black Square Media

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```