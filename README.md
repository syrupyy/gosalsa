# gosalsa
gbsalsa but horribly transpiled to work as a go package
## but what does gbsalsa do?
it decompiles MIDI files from Cookie Run: OvenBreak
## and where can i find gbsalsa?
https://github.com/simontime/gbsalsa
## also how do i use this?
like this
```go
package main

import (
    "fmt"
    "github.com/syrupyy/gosalsa"
)

func main() {
    gosalsa.CryptFile("sound/magma/bgm_whipped_cream.mid")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("file decrypted/encrypted successfully")
    }
}
```
