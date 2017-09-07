package main

import (
    "os"
    "os/signal"
    "syscall"
    "unsafe"
    "bytes"
    "fmt"
    "math/rand"
    "time"
    "flag"
    "strings"

    "github.com/mnito/charmine/keyio"
)

// Text styling
const REVERSE = "7"
const BOLD = "1"

func style(s string, style string) string {
    return "\u001B[" + style + "m" + s + "\u001B[0m"
}

// Terminal size struct
type termsize struct {
    Rows uint16
    Columns uint16
    _ uint16
    _ uint16
}

// Clear all rows
func clear(rows int) {
    for i := 0; i < rows; i += 1 {
        fmt.Printf("\x1B[A\r")
    }
}

// Render objects
func render(objects [][]string) {
    for _, row := range objects {
        for _, object := range row {
            fmt.Printf(object)
        }
    }
}

func main() {
    /* Key codes */
    LEFT := []byte{ 27, 91, 68, 0 }
    UP := []byte{ 27, 91, 65, 0 }
    RIGHT := []byte{ 27, 91, 67, 0 }
    DOWN := []byte{ 27, 91, 66, 0 }
    ESC := []byte{ 27, 0, 0, 0 }
    Q := []byte{ 113, 0, 0, 0 }
    R := []byte{ 114, 0, 0, 0 }
    Y := []byte{ 121, 0, 0, 0 }

    /* Set up flags */
    var DENSITY int
    var ALLOWEDCONTROLS string
    var SEED int64

    flag.IntVar(&DENSITY, "density", 45, "set density of letters (1-100)")
    flag.StringVar(&ALLOWEDCONTROLS, "controls", "ludr", "set controls (ludr)")
    flag.Int64Var(&SEED, "seed", time.Now().UTC().UnixNano(), "set seed")
    flag.Parse()

    /* Set up gameplay conditions */
    PLAYAGAIN := fmt.Sprintf(
        "%s -density %d -controls %s -seed %d",
        os.Args[0],
        DENSITY,
        ALLOWEDCONTROLS,
        SEED)

    LEFTALLOWED := strings.Contains(ALLOWEDCONTROLS, "l")
    UPALLOWED := strings.Contains(ALLOWEDCONTROLS, "u")
    DOWNALLOWED := strings.Contains(ALLOWEDCONTROLS, "d")
    RIGHTALLOWED := strings.Contains(ALLOWEDCONTROLS, "r")

    DENSITYPROB := 99 - DENSITY

    rand.Seed(SEED)

    /* Get terminal size to fill whole terminal */
    termsize := &termsize{}
    syscall.Syscall(
        syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(termsize)))

    rows := int(termsize.Rows)
    columns := int(termsize.Columns)

    /* Game setup */
    // Instantiate multidimensional slice for game objects
    objects := make([][]string, rows)
    for row := range objects {
        objects[row] = make([]string, columns)
        for column := range objects[row] {
            x := rand.Intn(100)
            c := " "
            if x > DENSITYPROB {
                c = string(65 + rand.Intn(6))
            }

            objects[row][column] = c
        }
    }

    GOAL := columns

    char := 65
    startchar := char
    charlen := 6

    initialX := 0
    initialY := rows - 1
    x := initialX
    y := initialY

    // Set object at point and render objects
    objects[y][x] = style(style(string(char), REVERSE), BOLD)

    render(objects)

    /* Set up reader */
    reader := keyio.NewReader()

    // Close reader on ctrl-c
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
    go func(r keyio.Reader) {
        <-signals
        r.Close()
        os.Exit(0)
    }(reader)

    /* Game loop */
    for {
        // Get input
        input := []byte{0, 0, 0, 0}
        reader.Read(input)

        // Process input
        objects[y][x] = " "
        if RIGHTALLOWED && bytes.Compare(input, RIGHT) == 0 {
            x += 1
            char += 1
        } else if UPALLOWED && bytes.Compare(input, UP) == 0 && y > 0 {
            y -= 1
            char += 1
        } else if LEFTALLOWED && bytes.Compare(input, LEFT) == 0 && x > 0{
            x -= 1
            char += 1
        } else if DOWNALLOWED && bytes.Compare(input, DOWN) == 0 && y + 1 < rows {
            y += 1
            char += 1
        } else if bytes.Compare(input, R) == 0 {
            x = initialX
            y = initialY
            char = startchar
        } else if bytes.Compare(input, ESC) == 0 || bytes.Compare(input, Q) == 0 {
            break
        }

        // Loop char back around if necessary
        if char >= startchar + charlen {
            char = startchar
        }

        // Check for win
        if x == GOAL {
            fmt.Printf(style("You won!", BOLD) + "\nTo play again:%s\n", PLAYAGAIN)
            break
        }

        // Check for loss collision
        if !(objects[y][x] == string(char) || objects[y][x] == " ") {
            input := []byte{ 0, 0, 0, 0 }
            fmt.Printf("You lost. Try again?")
            reader.Read(input)
            if bytes.Compare(input, Y) != 0 && bytes.Compare(input, R) != 0 {
                fmt.Printf("\nTo play again: %s\n", PLAYAGAIN)
                break
            }

            // Set back to original values if play again
            char = startchar
            x = initialX
            y = initialY
        }

        // Update player object location
        objects[y][x] = style(style(string(char), REVERSE), BOLD)

        // Clear board and render objects
        clear(rows)
        render(objects)
    }

    // Restore to original terminal
    reader.Close()
}
