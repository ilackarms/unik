package cmd

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/client"
)

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "List available unikernel images",
	Long: `Lists all available unikernel images across providers.
Includes important information for running and managing instances,
including bind mounts required at runtime.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := func() error {
			if err := readClientConfig(); err != nil {
				return err
			}
			if host == "" {
				host = clientConfig.Host
			}
			logrus.WithField("host", host).Info("listing images")
			images, err := client.UnikClient(host).Images().All()
			if err != nil {
				return errors.New("listing images failed", err)
			}
			printImages(images...)
			return nil
		}(); err != nil {
			logrus.Errorf("failed listing images: %v", err)
			os.Exit(-1)
		}
	},
}

func init() {
	RootCmd.AddCommand(imagesCmd)
}
