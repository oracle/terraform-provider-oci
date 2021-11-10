// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"

	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"

	//	"strings"

	//"github.com/terraform-providers/terraform-provider-oci/oci/resourcediscovery"

	//	"github.com/terraform-providers/terraform-provider-oci/oci/tfresource"

	//	"github.com/fatih/color"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/internal/provider"
)

func main() {
	var command = flag.String("command", "", "Command to run. Supported commands include: 'export', 'list_export_resources' and 'list_export_services'. 'list_export_services' supports json format.")
	var listExportServicesPath = flag.String("list_export_services_path", "", "[export] Path to output list of supported services in json format")
	var compartmentId = flag.String("compartment_id", "", "[export] OCID of a compartment to export. If no compartment id nor name is specified, the root compartment will be used.")
	var compartmentName = flag.String("compartment_name", "", "[export] The name of a compartment to export.")
	var includeRelatedResources = flag.Bool("include_related_resources", false, "[export] Set this flag to discover related resources for the resource OCIDs specified in `ids` argument.")
	var outputPath = flag.String("output_path", "", "[export] Path to output generated configurations and state files of the exported compartment")
	var services = flag.String("services", "", "[export] Comma-separated list of service resources to export. By default, all compartment-scope resources are exported.")
	var excludeServices = flag.String("exclude_services", "", "[export] [experimental] Comma-separated list of service resources to exclude from export. If a service is present in both 'services' and 'exclude_services' argument, it will be excluded.")
	var ids = flag.String("ids", "", "[export] Comma-separated list of tuples <resource Type:resource ID> for resources to export. The ID could either be an OCID or a Terraform import ID. By default, all resources are exported.")
	var generateStateFile = flag.Bool("generate_state", false, "[export][experimental] Set this to import the discovered resources into a state file along with the Terraform configuration")
	var help = flag.Bool("help", false, "Prints usage options")
	var tfVersion = flag.String("tf_version", "0.12", "The version of terraform syntax to generate for configurations. The state file will be written in v0.12 only. The allowed values are :\n * 0.11\n * 0.12")
	var retryTimeout = flag.String("retry_timeout", "15s", "[export] The time duration for which API calls will wait and retry operation in case of API errors. By default, the retry timeout duration is 15s")
	var parallelism = flag.Int("parallelism", 1, "The number of threads to use for resource discovery. By default the value is 1")

	flag.Parse()
	globalvar.PrintVersion()

	if help != nil && *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if command == nil || *command == "" {
		log.Println("Executable runs in Terraform plugin mode by default. For additional usage options, please run with the '-help' flag.")
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: func() terraform.ResourceProvider {
				return provider.Provider()
			},
		})
	} else {
		switch *command {
		case "export":

			var terraformVersion resourcediscovery.TfHclVersion
			if resourcediscovery.TfVersionEnum(*tfVersion) == resourcediscovery.TfVersion11 {
				terraformVersion = &resourcediscovery.TfHclVersion11{Value: resourcediscovery.TfVersionEnum(*tfVersion)}
			} else if *tfVersion == "" || resourcediscovery.TfVersionEnum(*tfVersion) == resourcediscovery.TfVersion12 {
				terraformVersion = &resourcediscovery.TfHclVersion12{Value: resourcediscovery.TfVersionEnum(*tfVersion)}
			} else {
				color.Red("[ERROR]: Invalid tf_version '%s', supported values: 0.11, 0.12\n", *tfVersion)
				os.Exit(1)
			}

			if *parallelism < 1 {
				color.Red("[ERROR] parallelism cannot be less than 1, specify at least 1")
				os.Exit(1)
			}

			args := &resourcediscovery.ExportCommandArgs{
				CompartmentId:                compartmentId,
				CompartmentName:              compartmentName,
				OutputDir:                    outputPath,
				GenerateState:                *generateStateFile,
				TFVersion:                    &terraformVersion,
				RetryTimeout:                 retryTimeout,
				IsExportWithRelatedResources: *includeRelatedResources,
				Parallelism:                  *parallelism,
			}

			if services != nil && *services != "" {
				args.Services = strings.Split(*services, ",")
			}

			if excludeServices != nil && *excludeServices != "" {
				args.ExcludeServices = strings.Split(*excludeServices, ",")
			}

			if ids != nil && *ids != "" {
				args.IDs = strings.Split(*ids, ",")
			}
			err, status := resourcediscovery.RunExportCommand(args)
			if err != nil {
				color.Red("%v", err)
			}
			os.Exit(int(status))

		case "list_export_resources":
			if err := resourcediscovery.RunListExportableResourcesCommand(); err != nil {
				color.Red("%v", err)
				os.Exit(1)
			}
		case "list_export_services":
			if err := resourcediscovery.RunListExportableServicesCommand(*listExportServicesPath); err != nil {
				color.Red("%v", err)
				os.Exit(1)
			}
		default:
			log.Printf("[ERROR]: No command '%s' supported\n", *command)
			os.Exit(1)
		}
	}
}
