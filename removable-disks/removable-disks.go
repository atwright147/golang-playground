package main

import (
	"fmt"
	"strings"

	"github.com/jaypipes/ghw"
)

func main() {
	block, err := ghw.Block()
	if err != nil {
		fmt.Printf("Error getting block storage info: %v", err)
		return
	}

	for _, disk := range block.Disks {
		if disk.IsRemovable {
			fmt.Printf("Disk Name: %s\n", disk.Name)
			fmt.Printf("  Vendor: %s\n", disk.Vendor)
			fmt.Printf("  Model: %s\n", disk.Model)
			fmt.Printf("  Serial Number: %s\n", disk.SerialNumber)
			fmt.Printf("  WWN: %s\n", disk.WWN)

			for _, partition := range disk.Partitions {
				var label string
				if partition.FilesystemLabel != "" {
					label = strings.TrimSpace(partition.FilesystemLabel)
				}
				if label == "" && partition.Label != "" {
					label = strings.TrimSpace(partition.Label)
				}
				if label == "" && disk.Model != "" {
					label = strings.TrimSpace(disk.Model)
				}
				if label == "" {
					label = "Untitled"
				}

				fmt.Printf("  Partition Name: %s\n", partition.Name)
				fmt.Printf("    Label: %s\n", partition.Label)
				fmt.Printf("    Filesystem Label: %s\n", partition.FilesystemLabel)
				fmt.Printf("    My Label: %s\n", label)
			}
			fmt.Println()
		}
	}
}
