package utils

import (
    "os"
    "testing"
    "bufio"
    "log"
)

func fvCompareTest(t *testing.T) {
    const testFile = "test/utilsText.txt"
    const textContent = "This is a test file. Used in `utils/test.go`."

    f, err := os.Open(testFile)
    if err != nil {
	log.Fatal(err)
    }

    fscan := bufio.NewScanner(f)

    if !fvCompare(fscan.Text(), textContent) {
	t.Error("Content is not the same!")
    }

}
