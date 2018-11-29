# bits2str
Golang bits to string converter

# Install

  go get -u github.com/moisespsena-go/bits2str/...
  
# Usage

```go
package main

import github.com/moisespsena-go/bits2str

func main() {
  println(bits2str.Bits(1024).String())
}
```

# CLI

  bits2str 1024
  
  or
  
  echo '1024
  4
  4200' | bits2str
