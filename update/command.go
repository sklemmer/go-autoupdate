package update

import "github.com/spf13/cobra"

var UpdateCommand = &cobra.Command{
	Use:   "update",
	Short: "Update updates the application with latest from git",
	Long:  "",
	RunE:  runCmd,
}

func runCmd(cmd *cobra.Command, args []string) (error) {
	updater := NewUpdater()
	ok, err := updater.Check()

	if err != nil {
		return err
	}

	if ok {
		if err := updater.Update(); err != nil {
			return err
		}
	}
	return nil
}
