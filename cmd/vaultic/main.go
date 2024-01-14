package main

import (
	"fmt"
	"os"

	"github.com/smxlong/vaultic"
)

func main() {
	if err := vaultic.New().Command().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
