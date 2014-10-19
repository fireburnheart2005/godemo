package main
import (
	"fmt"

	"errors"
)

func main() {
	//error
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
	    fmt.Print(err)
	}
}