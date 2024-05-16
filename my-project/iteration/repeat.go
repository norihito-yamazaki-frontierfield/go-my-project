package iteration

import "strings"

// Repeat returns a new string consisting of the specified character repeated the specified number of times.
func Repeat(character string, count int) string {
    var builder strings.Builder
    builder.Grow(len(character) * count)
    for i := 0; i < count; i++ {
        builder.WriteString(character)
    }
    return builder.String()
}
