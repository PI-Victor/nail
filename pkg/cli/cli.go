package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/PI-Victor/nail/pkg/platform"
	"github.com/PI-Victor/nail/pkg/util"
)

var (
	repoName     string
	platformName string

	AddRepo = &cobra.Command{
		Use:   "add",
		Short: "Add a new repository to the platform",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	CreateRelease = &cobra.Command{
		Use:   "release",
		Short: "Create a new platform release",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	ValidateRelease = &cobra.Command{
		Use:   "validate",
		Short: "Validate the current release",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	InitPlatform = &cobra.Command{
		Use:   "init",
		Short: "Initialize a new empty state file",
		Run: func(cmd *cobra.Command, args []string) {
			p := platform.NewPlatform(platformName)
			w := util.NewWorkspace(p)
			if err := w.CreateWorkspace(); err != nil {
				logrus.Fatalf("Failed to create a new workspace: %v", err)
			}
		},
	}

	DeployRelease = &cobra.Command{
		Use:   "deploy",
		Short: "Deploy a specific version of the application",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func init() {
	InitPlatform.PersistentFlags().StringVar(
		&platformName,
		"name",
		"",
		"The name of the platform you want to create a new workspace for",
	)
}
