package exts

import(
    "strings"
)

func KeysAsString(m map[string]func(a []int)) string {
    names := make([]string, len(m))

    i := 0

    for name := range m {
        names[i] = name
        i++
    }

    return strings.Join(names[:], ", ")
}
