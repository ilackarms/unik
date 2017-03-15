package cmd

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/client"
)

var volumeName string

var rmvCmd = &cobra.Command{
	Use:     "delete-volume",
	Aliases: []string{"rmv"},
	Short:   "Delete a unikernel volume",
	Long: `Deletes a volume.
You may specify the volume by name or id.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := func() error {
			if err := readClientConfig(); err != nil {
				return err
			}
			if volumeName == "" {
				return errors.New("must specify --volume", nil)
			}
			if host == "" {
				host = clientConfig.Host
			}
			logrus.WithFields(logrus.Fields{"host": host, "force": force, "volume": volumeName}).Info("deleting volume")
			if err := client.UnikClient(host).Volumes().Delete(volumeName, force); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			logrus.Errorf("failed deleting volume: %v", err)
			os.Exit(-1)
		}
	},
}

func init() {
	RootCmd.AddCommand(rmvCmd)
	rmvCmd.Flags().StringVar(&volumeName, "volume", "", "<string,required> name or id of volume. unik accepts a prefix of the name or id")
	rmvCmd.Flags().BoolVar(&force, "force", false, "<bool, optional> forces detaching the volume before deletion if it is currently attached")
}
