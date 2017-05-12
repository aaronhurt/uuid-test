package main

import (
	"fmt"

	"github.com/leprechau/uuid-test/uuidf"
	fuuid "github.com/rogpeppe/fastuuid"
)

func main() {
	g := fuuid.MustNewGenerator()
	fmt.Println(uuidf.ToString(g.Next()))
}
