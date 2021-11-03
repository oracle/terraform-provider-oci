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
		GenerateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", Required, Create, cloudGuardConfigurationRepresentation)

	CloudGuardConfigurationResourceConfig = CloudGuardConfigurationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", Optional, Update, cloudGuardConfigurationRepresentation)

	cloudGuardConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
	}

	//Has to be a valid reporting region where the tenant is subscribed to
	reportingRegion                       = GetEnvSettingWithDefault("region", "us-phoenix-1")
	cloudGuardConfigurationRepresentation = map[string]interface{}{
		"compartment_id":   Representation{RepType: Required, Create: `${var.compartment_id}`},
		"reporting_region": Representation{RepType: Required, Create: reportingRegion, Update: reportingRegion},
		//Only "ENABLED" and "DISABLED" status fields are allowed, the latter will off-board the customer; soft deleting CP components and disallowing ops which we dont want.
		"status":                Representation{RepType: Required, Create: `ENABLED`, Update: `ENABLED`},
		"self_manage_resources": Representation{RepType: Optional, Create: `false`, Update: nil},
	}

	CloudGuardConfigurationResourceDependencies = ""
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardCloudGuardConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardCloudGuardConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	//Enable/Disable CG is a tenant-level operation
	compartmentId := GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_cloud_guard_cloud_guard_configuration.test_cloud_guard_configuration"

	singularDatasourceName := "data.oci_cloud_guard_cloud_guard_configuration.test_cloud_guard_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+CloudGuardConfigurationResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", Optional, Create, cloudGuardConfigurationRepresentation), "cloudguard", "cloudGuardConfiguration", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardConfigurationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", Required, Create, cloudGuardConfigurationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "reporting_region", reportingRegion),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudGuardConfigurationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", Optional, Create, cloudGuardConfigurationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "reporting_region", reportingRegion),
				resource.TestCheckResourceAttr(resourceName, "self_manage_resources", "false"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				GenerateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", Optional, Update, cloudGuardConfigurationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "reporting_region", reportingRegion),
				resource.TestCheckResourceAttr(resourceName, "self_manage_resources", "false"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", Required, Create, cloudGuardConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardConfigurationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(singularDatasourceName, "reporting_region", reportingRegion),
				resource.TestCheckResourceAttr(singularDatasourceName, "self_manage_resources", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "ENABLED"),
			),
		},
	})
}
