package main
import (
	"syscall"
	"fmt"
)

func main() {
	printCurrentCoreRlimit()

	err := syscall.Setrlimit(syscall.RLIMIT_CORE, &syscall.Rlimit{syscall.RLIM_INFINITY, syscall.RLIM_INFINITY})
	if err != nil {
		panic(err)
	}

	printCurrentCoreRlimit()
}

func printCurrentCoreRlimit() {
	rlimit := syscall.Rlimit{}
	err := syscall.Getrlimit(syscall.RLIMIT_CORE, &rlimit)
	if err != nil {
		panic(err)
	}
	fmt.Printf("core-setting: soft=%d, hard=%d\n", rlimit.Cur, rlimit.Max)
}
