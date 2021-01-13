// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	CloudGuardConfigurationRequiredOnlyResource = CloudGuardConfigurationResourceDependencies +
		generateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", Required, Create, cloudGuardConfigurationRepresentation)

	CloudGuardConfigurationResourceConfig = CloudGuardConfigurationResourceDependencies +
		generateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", Optional, Update, cloudGuardConfigurationRepresentation)

	cloudGuardConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}

	//Has to be a valid reporting region where the tenant is subscribed to
	reportingRegion                       = getEnvSettingWithDefault("region", "us-phoenix-1")
	cloudGuardConfigurationRepresentation = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
		"reporting_region": Representation{repType: Required, create: reportingRegion, update: reportingRegion},
		//Only "ENABLED" and "DISABLED" status fields are allowed, the latter will off-board the customer; soft deleting CP components and disallowing ops which we dont want.
		"status":                Representation{repType: Required, create: `ENABLED`, update: `ENABLED`},
		"self_manage_resources": Representation{repType: Optional, create: `false`, update: nil},
	}

	CloudGuardConfigurationResourceDependencies = ""
)

func TestCloudGuardCloudGuardConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardCloudGuardConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	//Enable/Disable CG is a tenant-level operation
	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_cloud_guard_cloud_guard_configuration.test_cloud_guard_configuration"

	singularDatasourceName := "data.oci_cloud_guard_cloud_guard_configuration.test_cloud_guard_configuration"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + CloudGuardConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", Required, Create, cloudGuardConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "reporting_region", reportingRegion),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CloudGuardConfigurationResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CloudGuardConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", Optional, Create, cloudGuardConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "reporting_region", reportingRegion),
					resource.TestCheckResourceAttr(resourceName, "self_manage_resources", "false"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + CloudGuardConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", Optional, Update, cloudGuardConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "reporting_region", reportingRegion),
					resource.TestCheckResourceAttr(resourceName, "self_manage_resources", "false"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", Required, Create, cloudGuardConfigurationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CloudGuardConfigurationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(singularDatasourceName, "reporting_region", reportingRegion),
					resource.TestCheckResourceAttr(singularDatasourceName, "self_manage_resources", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "status", "ENABLED"),
				),
			},
		},
	})
}
