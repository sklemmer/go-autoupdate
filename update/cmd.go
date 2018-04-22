package update

import (
	"github.com/spf13/cobra"
	"fmt"
)

func (u *Updater) GetUpdateCommand() (*cobra.Command) {
	return &cobra.Command{
		Use:   "update",
		Short: "Update updates the application with latest from git",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) (error) {
			ok, err := u.Check()
			if err != nil {
				return err
			}

			if ok {
				if err := u.Update(); err != nil {
					return err
				}
				fmt.Println("Version successfully updated")
			}
			return nil
		},
	}
}
