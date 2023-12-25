package govault

import (
	"fmt"
	"log"

	"golang.org/x/sys/unix"

	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "List available disk storage on your system.",
	Long:  `Get an overview of how much space is used by your disk and how much is available.`,
	Run: func(cmd *cobra.Command, args []string) {
		listDiskStorage()
	},
}

func init() {
	rootCmd.AddCommand(diskCmd)
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type DiskStatus struct {
	Total     uint64 `json:"Total"`
	Used      uint64 `json:"Used"`
	Available uint64 `json:"Available"`
}

func diskStorage() (disk DiskStatus) {

	fs := unix.Statfs_t{}
	err := unix.Statfs("/", &fs)
	if err != nil {
		log.Fatal(err)
	}

	total := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := total - free
	return DiskStatus{
		Total:     total,
		Used:      used,
		Available: free,
	}

}

func listDiskStorage() {
	diskStatus := diskStorage()

	fmt.Printf("Total Storage: %.2f GB\n", float64(diskStatus.Total)/float64(GB))
	fmt.Printf("Used Storage: %.2f GB\n", float64(diskStatus.Used)/float64(GB))
	fmt.Printf("Available Storage: %.2f GB\n", float64(diskStatus.Available)/float64(GB))
}
