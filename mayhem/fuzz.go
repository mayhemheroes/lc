package fuzz

import "strconv"
import "github.com/boyter/lc/parsers"
import "github.com/boyter/lc/processor"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) < 1 {
        num = 0
    } else {
        num, _ = strconv.Atoi(string(bytes[0]))
    }

    switch num {
    
    case 0:
        parsers.GuessLicense(bytes)
        return 0

    case 1:
        content := string(bytes)
        processor.LcCleanText(content)
        return 0

    default:
        var test processor.LicenceGuesser
        test.KeyWordGuessLicence(bytes)
        return 0
    }
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}