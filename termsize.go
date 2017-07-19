package main

import "syscall"
import "unsafe"

type _winsize struct {
    ws_row uint16
    ws_col uint16
    ws_xpixel uint16
    ws_ypixel uint16
}

type winsize struct {
    Rows int
    Columns int
    Pixelsacross int
    Pixelsdown int
}

func Gettermsize() winsize {
    _size := &_winsize{}

    status, _, error := syscall.Syscall(
        syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(_size)))

    if int(status) == -1 {
        panic("Could not get size information: " + string(error))
    }

    size := winsize{
        int(_size.ws_row),
        int(_size.ws_col),
        int(_size.ws_xpixel),
        int(_size.ws_ypixel)}

    return size
}
