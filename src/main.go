package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <uuid> [flags]")
		return
	}

	uuid := os.Args[1]
	flags := os.Args[2:]

	for _, flag := range flags {
		switch {
		case flag == "-u":
			uuid = strings.ToUpper(uuid)
		case flag == "-l":
			uuid = strings.ToLower(uuid)
		case flag == "--remove-dash":
			uuid = strings.ReplaceAll(uuid, "-", "")
		case flag == "--include-dash":
			if len(uuid) == 32 {
				uuid = uuid[0:8] + "-" + uuid[8:12] + "-" + uuid[12:16] + "-" + uuid[16:20] + "-" + uuid[20:32]
			} else {
				fmt.Println("Invalid UUID length for --include-dash")
				return
			}
		case flag == "--remove-prefix":
			uuid = uuid[len(uuid)-32:]
		case strings.HasPrefix(flag, "--prefix="):
			prefix := strings.TrimPrefix(flag, "--prefix=")
			if prefix == "" {
				fmt.Println("Error: prefix not defined")
				return
			}
			uuid = prefix + ":" + uuid
		default:
			fmt.Printf("Unknown flag: %s\n", flag)
			return
		}
	}

	fmt.Println(uuid)
}
