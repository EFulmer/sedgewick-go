package main

import (
    "cmp"
    "fmt"
    "os"
)

func Sort[T cmp.Ordered](a []T) {
    n := len(a)
    for i := 0; i < n; i++ {
        minIndex := i
        for j := i + 1; j < n; j++ {
            if cmp.Less(a[j], a[minIndex]) {
                minIndex = j
            }
            exch(a, i, minIndex)
        }
    }
}

func exch[T cmp.Ordered](a []T, i, j int) {
    t := a[i]
    a[i] = a[j]
    a[j] = t
}

func IsSorted[T cmp.Ordered](a []T) bool {
    for i := 0; i < len(a) - 1; i++ {
        if !(cmp.Less(a[i], a[i+1]) || a[i] == a[i+1]) {
            return false
        }
    }
    return true
}

func main() {
    a := os.Args[1:]
    Sort(a)
    if !IsSorted(a) {
        fmt.Println("Eric messed up!")
    }
    for _, elem := range a {
        fmt.Printf("%v ", elem)
    }
    fmt.Println()
}
