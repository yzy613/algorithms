package common

import flag "github.com/spf13/pflag"

var (
	File = flag.StringP("file", "f", "", "file path")
)
