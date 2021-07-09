package arg

import "flag"

var (
	Name string
	Shhh bool
)

// this function will be called automatically
func init() {
	// store name flag into name var
	flag.StringVar(&Name, "name", "Unknown", "your name")

	// check that is in silence mode
	flag.BoolVar(&Shhh, "shhh", false, "disable logging")

	// parse flags (options)
	flag.Parse()
}
