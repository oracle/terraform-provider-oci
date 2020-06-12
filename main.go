// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"

	provider "github.com/terraform-providers/terraform-provider-oci/oci"
)

func main() {
	var command = flag.String("command", "", "Command to run. Supported commands include: 'export' and 'list_export_resources'")
	var compartmentId = flag.String("compartment_id", "", "[export] OCID of a compartment to export. If no compartment id nor name is specified, the root compartment will be used.")
	var compartmentName = flag.String("compartment_name", "", "[export] The name of a compartment to export.")
	var outputPath = flag.String("output_path", "", "[export] Path to output generated configurations and state files of the exported compartment")
	var services = flag.String("services", "", "[export] Comma-separated list of service resources to export. By default, all compartment-scope resources are exported.")
	var ids = flag.String("ids", "", "[export] Comma-separated list of resource IDs to export. The ID could either be an OCID or a Terraform import ID. By default, all resources are exported.")
	var generateStateFile = flag.Bool("generate_state", false, "[export][experimental] Set this to import the discovered resources into a state file along with the Terraform configuration")
	var help = flag.Bool("help", false, "Prints usage options")
	var tfVersion = flag.String("tf_version", "0.12", "The version of terraform syntax to generate for configurations. The state file will be written in v0.12 only. The allowed values are :\n * 0.11\n * 0.12")

	flag.Parse()
	provider.PrintVersion()

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

			var terraformVersion provider.TfHclVersion
			if provider.TfVersionEnum(*tfVersion) == provider.TfVersion11 {
				terraformVersion = &provider.TfHclVersion11{Value: provider.TfVersionEnum(*tfVersion)}
			} else if *tfVersion == "" || provider.TfVersionEnum(*tfVersion) == provider.TfVersion12 {
				terraformVersion = &provider.TfHclVersion12{Value: provider.TfVersionEnum(*tfVersion)}
			} else {
				log.Printf("[ERROR]: Invalid tf_version '%s', supported values: 0.11, 0.12\n", *tfVersion)
				os.Exit(1)
			}
			args := &provider.ExportCommandArgs{
				CompartmentId:   compartmentId,
				CompartmentName: compartmentName,
				OutputDir:       outputPath,
				GenerateState:   *generateStateFile,
				TFVersion:       &terraformVersion,
			}

			if services != nil && *services != "" {
				args.Services = strings.Split(*services, ",")
			}

			if ids != nil && *ids != "" {
				args.IDs = strings.Split(*ids, ",")
			}
			err, status := provider.RunExportCommand(args)
			if err != nil {
				color.Red("%v", err)
			}
			os.Exit(int(status))

		case "list_export_resources":
			if err := provider.RunListExportableResourcesCommand(); err != nil {
				log.Printf("%v", err)
				os.Exit(1)
			}
		default:
			log.Printf("[ERROR]: No command '%s' supported\n", *command)
			os.Exit(1)
		}
	}
}
