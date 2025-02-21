package utils

import "iter"


func Map[T, U any](s iter.Seq[T], f func(T) U) iter.Seq[U] {
    return func(yield func(U) bool) {
        for a := range s {
            if !yield(f(a)) {
                return
            }
        }
    }
}
