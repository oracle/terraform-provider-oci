// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/oracle/terraform-provider-oci/internal/provider"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
)

var filterFlag tf_export.Filter

func init() {
	// Tie the command-line flag to the intervalFlag variable and
	// set a usage message.
	flag.Var(&filterFlag, "filter", "pass a filter to filter resources discovered. Use the flag multiple times to pass multiple filters")
}

func main() {
	// TODO: input for resource discovery from a config file
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
	var varsResourceLevel = flag.String("variables_resource_level", "", "[export] List of top-level attributes to be export as variable following format resourceType.attribute, if attribute is present in variables_global_level, it will be excluded for this resourceType")
	var varsGlobalLevel = flag.String("variables_global_level", "", "[export] List of top-level attributes to be export as variable following format attribute1,attribute2, if attribute present in variables_resource_level, it will be excluded for this resourceType")

	flag.Parse()
	globalvar.PrintVersion()

	if help != nil && *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if command == nil || *command == "" {
		log.Println("Executable runs in Terraform plugin mode by default. For additional usage options, please run with the '-help' flag.")
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: func() *schema.Provider {
				return provider.Provider()
			},
		})
	} else {
		switch *command {
		case "export":

			var terraformVersion tf_export.TfHclVersion
			if tf_export.TfVersionEnum(*tfVersion) == tf_export.TfVersion11 {
				terraformVersion = &tf_export.TfHclVersion11{Value: tf_export.TfVersionEnum(*tfVersion)}
			} else if *tfVersion == "" || tf_export.TfVersionEnum(*tfVersion) == tf_export.TfVersion12 {
				terraformVersion = &tf_export.TfHclVersion12{Value: tf_export.TfVersionEnum(*tfVersion)}
			} else {
				color.Red("[ERROR]: Invalid tf_version '%s', supported values: 0.11, 0.12\n", *tfVersion)
				os.Exit(1)
			}

			if *parallelism < 1 {
				color.Red("[ERROR] parallelism cannot be less than 1, specify at least 1")
				os.Exit(1)
			}

			args := &tf_export.ExportCommandArgs{
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

			if varsResourceLevel != nil && *varsResourceLevel != "" {
				args.VarsExportResourceLevel = strings.Split(*varsResourceLevel, ",")
			}

			if varsGlobalLevel != nil && *varsGlobalLevel != "" {
				args.VarExportGlobalLevel = strings.Split(*varsGlobalLevel, ",")
			}

			if excludeServices != nil && *excludeServices != "" {
				args.ExcludeServices = strings.Split(*excludeServices, ",")
			}

			if ids != nil && *ids != "" {
				args.IDs = strings.Split(*ids, ",")
			}

			if filterFlag != nil {
				args.Filters = filterFlag
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
