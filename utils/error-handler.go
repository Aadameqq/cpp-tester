package utils

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println(Colours.Purple, "Error occurred", Colours.Reset)
		fmt.Println(" ", err)

		os.Exit(0)
	}

}
