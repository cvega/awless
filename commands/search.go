/*
Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wallix/awless/aws/services"
	"github.com/wallix/awless/config"
	"github.com/wallix/awless/logger"
)

var (
	showIdsOnlyFlag, showIdOnlyFlag, showLatestIdOnly bool
)

func init() {
	RootCmd.AddCommand(searchCmd)

	awsImagesCmd.Flags().BoolVar(&showLatestIdOnly, "latest-id", false, "Returns the id only of the latest AMI matching your query")
	awsImagesCmd.Flags().BoolVar(&showIdOnlyFlag, "id-only", false, "(DEPRECATED, use latest-id) Returns only one (the latest) AMI id matching the query")

	searchCmd.AddCommand(awsImagesCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Perform various searches and resolution",
}

var awsImagesCmd = &cobra.Command{
	Use:               "images",
	PersistentPreRun:  applyHooks(initAwlessEnvHook, initLoggerHook, initCloudServicesHook, firstInstallDoneHook),
	PersistentPostRun: applyHooks(networkMonitorHook),
	Short:             fmt.Sprintf("Find corresponding bare images according to a bare image query, ordering by latest first. Supported owners: %s", strings.Join(awsservices.SupportedAMIOwners, ", ")),
	Long:              fmt.Sprintf("Find corresponding bare images according to a bare image query, ordering by latest first.\n\nQuery string specification is the following column separated format:\n\n\t\t%s\n\nEverything optional expect for the 'owner'. Supported owners: %s", awsservices.ImageQuerySpec, strings.Join(awsservices.SupportedAMIOwners, ", ")),
	Example:           "  awless search images redhat:rhel:7.2\n  awless search images debian::jessie\n  awless search images canonical --latest-id\n  awless search images amazonlinux:::::instance-store",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			exitOn(fmt.Errorf("expecting image query string. Expecting: %s (with everything optional expect for the owner)", awsservices.ImageQuerySpec))
		}

		resolver := &awsservices.ImageResolver{InfraService: awsservices.InfraService.(*awsservices.Infra)}

		query, err := awsservices.ParseImageQuery(args[0])
		exitOn(err)

		logger.Infof("launching search for image in '%s' region. Query: '%s'", config.GetAWSRegion(), query)
		imgs, err := resolver.Resolve(query)
		exitOn(err)

		var ids []string
		for _, img := range imgs {
			ids = append(ids, img.Id)
		}

		if showIdsOnlyFlag {
			for _, id := range ids {
				fmt.Println(id)
			}
			return
		}

		if showLatestIdOnly || showIdOnlyFlag {
			for i, id := range ids {
				fmt.Println(id)
				if i == 0 {
					break
				}
			}
			return
		}

		b, err := json.MarshalIndent(imgs, "", " ")
		exitOn(err)

		fmt.Fprintln(os.Stdout, string(b))
	},
}
