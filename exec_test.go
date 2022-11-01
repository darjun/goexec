package goexec_test

import (
	"fmt"
	"testing"

	"github.com/darjun/goexec"
)

func TestCombinedOutputString(t *testing.T) {
	fmt.Println(goexec.CombinedOutputString("cal", nil, goexec.WithEnv("LANG", "en_US.UTF-8")))
}
