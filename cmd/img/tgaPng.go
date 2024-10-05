package img

import (
	"os/exec"
	"syscall"

	"github.com/bardic/openpbr/cmd/utils"
	"github.com/spf13/cobra"
)

var TgaPngCmd = &cobra.Command{
	Use:   "tgapng",
	Short: "convert tgas to pngs",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		in := args[0]
		out := args[1]
		c := exec.Command(utils.IM_CMD, in, "png32:"+out)
		c.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
		go c.Run()
		return nil
	},
}
