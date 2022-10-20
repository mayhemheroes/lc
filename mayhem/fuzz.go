package fuzz

import "github.com/boyter/lc/parsers"

func mayhemit(bytes []byte) int {

    parsers.GuessLicense(bytes)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}