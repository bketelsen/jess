package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createconfigCmd represents the createconfig command
var createconfigCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a config file in your home directory",
	Long:  `Creates a .jess.yaml file in your home directory, which you must edit to provide default configuration values for the app.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := createTheConfig(cmd)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(createconfigCmd)
	createconfigCmd.Flags().BoolP("force", "f", false, "Overwrite existing")
}

func createTheConfig(cmd *cobra.Command) error {
	name := ".jess.yaml"
	path, err := homeDir()
	if err != nil {
		return errors.Wrap(err, "get home directory")
	}
	e, err := exists(filepath.Join(path, name))
	if err != nil {
		return errors.Wrap(err, "check existing config")
	}
	f, err := cmd.Flags().GetBool("force")
	if err != nil {
		return errors.Wrap(err, "getting force flag")
	}
	if e && !f {
		return errors.New("File exists, use --force")
	}
	if e {
		err := os.Remove(filepath.Join(path, name))

		if err != nil {
			return errors.Wrap(err, "removing existing configuration file")
		}
	}
	conf := Conf{
		Author:    viper.GetString("author"),
		Twitter:   viper.GetString("twitter"),
		Email:     viper.GetString("email"),
		ModuleDir: "/Users/me/github.com/me/modules",
		CourseDir: "/Users/me/courses",
	}
	return writeTemplateToFile(path, name, Config, conf)
}
