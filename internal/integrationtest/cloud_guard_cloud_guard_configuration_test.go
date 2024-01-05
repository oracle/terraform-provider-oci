// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	CloudGuardCloudGuardConfigurationRequiredOnlyResource = CloudGuardCloudGuardConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", acctest.Required, acctest.Create, CloudGuardCloudGuardConfigurationRepresentation)

	CloudGuardCloudGuardConfigurationResourceConfig = CloudGuardCloudGuardConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", acctest.Optional, acctest.Update, CloudGuardCloudGuardConfigurationRepresentation)

	CloudGuardCloudGuardCloudGuardConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	//Has to be a valid reporting region where the tenant is subscribed to
	reportingRegion                                 = utils.GetEnvSettingWithDefault("region", "us-phoenix-1")
	CloudGuardCloudGuardConfigurationRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"reporting_region": acctest.Representation{RepType: acctest.Required, Create: reportingRegion, Update: reportingRegion},
		//Only "ENABLED" and "DISABLED" status fields are allowed, the latter will off-board the customer; soft deleting CP components and disallowing ops which we dont want.
		"status":                acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `ENABLED`},
		"self_manage_resources": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: nil},
	}

	CloudGuardCloudGuardConfigurationResourceDependencies = ""
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardCloudGuardConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardCloudGuardConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	//Enable/Disable CG is a tenant-level operation
	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_cloud_guard_cloud_guard_configuration.test_cloud_guard_configuration"

	singularDatasourceName := "data.oci_cloud_guard_cloud_guard_configuration.test_cloud_guard_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudGuardCloudGuardConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", acctest.Optional, acctest.Create, CloudGuardCloudGuardConfigurationRepresentation), "cloudguard", "cloudGuardConfiguration", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardCloudGuardConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", acctest.Required, acctest.Create, CloudGuardCloudGuardConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "reporting_region", reportingRegion),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardCloudGuardConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudGuardCloudGuardConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", acctest.Optional, acctest.Create, CloudGuardCloudGuardConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "reporting_region", reportingRegion),
				resource.TestCheckResourceAttr(resourceName, "self_manage_resources", "false"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CloudGuardCloudGuardConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", acctest.Optional, acctest.Update, CloudGuardCloudGuardConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "reporting_region", reportingRegion),
				resource.TestCheckResourceAttr(resourceName, "self_manage_resources", "false"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_cloud_guard_configuration", "test_cloud_guard_configuration", acctest.Required, acctest.Create, CloudGuardCloudGuardCloudGuardConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardCloudGuardConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(singularDatasourceName, "reporting_region", reportingRegion),
				resource.TestCheckResourceAttr(singularDatasourceName, "self_manage_resources", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "ENABLED"),
			),
		},
	})
}
