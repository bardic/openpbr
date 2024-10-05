package gen

import (
	"os/exec"
	"syscall"

	"github.com/bardic/openpbr/cmd/utils"
	"github.com/spf13/cobra"
)

var HeightCmd = &cobra.Command{
	Use:   "height",
	Short: "create heightmaps based on colour image",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		in := args[0]
		out := args[1]
		c := exec.Command(utils.IM_CMD, in, "-channel", "RGB", "-negate", "-set", "colorspace", "Gray", "png32:"+out)
		c.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
		go c.Run()
		return nil
	},
}
