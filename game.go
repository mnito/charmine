package main

import "fmt"
import "time"
import "strings"

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

func render2(objects [][]string, format string) {
    for _, row := range objects {
        args := []interface{}{}
        for _, object := range row {
            args = append(args, object)
        }
        fmt.Printf(format, args...)
    }
}

func main() {
    // Get terminal size to fill whole terminal
    termsize := Gettermsize()

    termsize.Rows = termsize.Rows

    // Instantiate multidimensional slice for game objects
    objects := make([][]string, termsize.Rows)
    for row := range objects {
        objects[row] = make([]string, termsize.Columns)
        for column := range objects[row] {
            objects[row][column] = " "
        }
    }

    // Set object at point and render objects
    objects[1][5] = red("A")
    render(objects)

    // Delay
    time.Sleep(500 * time.Millisecond)

    // Show movement
    clear(termsize.Rows)
    objects[1][5] = " "
    objects[1][6] = red("A")
    render(objects)

    // Delay
    time.Sleep(500 * time.Millisecond)

    // Alternate rendering
    format := strings.Repeat("%s", termsize.Columns)

    objects[8][9] = "B"
    clear(termsize.Rows)
    render2(objects, format)

    time.Sleep(250 * time.Millisecond)

    objects[8][9] = " "
    objects[8][10] = "B"
    time.Sleep(250 * time.Millisecond)

    clear(termsize.Rows)
    render2(objects, format)

    time.Sleep(250 * time.Millisecond)
}
