
package keyio

import (
    "os"
    "bufio"
    "syscall"
    "unsafe"
)

const ICANON = syscall.ICANON
const ECHO = syscall.ECHO
const VMIN = syscall.VMIN
const VTIME = syscall.VTIME

type Reader struct {
    reader *bufio.Reader
    origterm *syscall.Termios
}

func NewReader() Reader {
    /* Set up terminal to accept input without pressing enter */
    term := &syscall.Termios{ }
    origterm := &syscall.Termios{ }

    syscall.Syscall(
        syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(TCGETS),
        uintptr(unsafe.Pointer(term)))

    *origterm = *term

    term.Lflag &^= (ICANON | ECHO)
    term.Cc[VMIN] = 1
    term.Cc[VTIME] = 0

    syscall.Syscall(
        syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(TCSETS),
        uintptr(unsafe.Pointer(term)))

    // Hide cursor
    print("\x1B[?25l")

    return Reader{bufio.NewReader(os.Stdin), origterm}
}

func (r Reader) Close() {
    print("\x1B[?25h")
    syscall.Syscall(
        syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(TCSETS),
        uintptr(unsafe.Pointer(r.origterm)))
}
