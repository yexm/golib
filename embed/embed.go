package embed

import (
	_ "embed"
)

//go:embed macbeth.txt
var Macbeth string
