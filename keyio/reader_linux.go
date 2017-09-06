package keyio

import "syscall"

const TCGETS  = syscall.TCGETS
const TCSETS = syscall.TCSETS

func (r Reader) Read(input []byte) {
    singleInput := []byte{ 0 }

    r.reader.Read(singleInput)
    input[0] = singleInput[0]
    if input[0] == 27 {
        r.reader.Read(singleInput)
        input[1] = singleInput[0]
        r.reader.Read(singleInput)
        input[2] = singleInput[0]
    }
}
