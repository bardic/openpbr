package cmd

import (
	"github.com/bardic/openpbr/utils"
	cp "github.com/otiai10/copy"
)

type Copy struct {
	Target string
	Dest   string
}

func (cmd *Copy) Perform() error {
	utils.AppendLoadOut("--- Copy custom configs")
	err := cp.Copy(
		utils.LocalPath(cmd.Target),
		utils.LocalPath(cmd.Dest),
	)

	return err
}
