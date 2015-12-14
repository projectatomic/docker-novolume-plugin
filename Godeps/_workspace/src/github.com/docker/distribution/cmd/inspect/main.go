package main

import (
	"fmt"

	"github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/reference"
	"github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/docker/registry"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "inspect NAME",
	Short: "inspect inspects an image in a remote registry",
	Long:  "inspect inspects an image in a remote registry.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 || args[0] == "" {
			logrus.Fatal("please provide an image name")
		}

		ref, err := reference.ParseNamed(args[0])
		if err != nil {
			logrus.Fatal(err)
		}

		hostname, name := reference.SplitHostname(ref)
		if hostname == "" {
			logrus.Info("using docker.io")
			hostname = "docker.io"
		}

		repoInfo, err := registry.ResolveRepository(ref)
		if err != nil {
			logrus.Fatal(err)
		}

		fmt.Printf("%v\n", ref)
		fmt.Printf("%v\n", hostname)
		fmt.Printf("%v\n", name)
		fmt.Printf("%v\n", repoInfo)
	},
}

//var format string

func init() {
	//Cmd.Flags().StringVarP(&format, "format", "f", "", "format data returned")
}

func main() {
	Cmd.Execute()
}
