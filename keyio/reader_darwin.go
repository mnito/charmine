package keyio

import "syscall"

const TCGETS  = syscall.TIOCGETA
const TCSETS = syscall.TIOCSETA

func (r Reader) Read(input []byte) {
    r.reader.Read(input)
}
