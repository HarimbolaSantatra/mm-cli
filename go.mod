module mm

go 1.22.0

require github.com/spf13/cobra v1.8.1

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace mm/cmd => ./cmd

replace mm/client => ./client

replace mm/utils => ./utils

replace mm/datatype => ./datatype/
