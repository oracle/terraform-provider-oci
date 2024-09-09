// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetAppsManagementOnboardingRequiredOnlyResource = FleetAppsManagementOnboardingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_onboarding", "test_onboarding", acctest.Required, acctest.Create, FleetAppsManagementOnboardingRepresentation)

	FleetAppsManagementOnboardingDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.tenancy_ocid}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementOnboardingDataSourceFilterRepresentation}}
	FleetAppsManagementOnboardingDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_onboarding.test_onboarding.id}`}},
	}

	FleetAppsManagementOnboardingRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"is_cost_tracking_tag_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_fams_tag_enabled":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	FleetAppsManagementOnboardingResourceDependencies = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementOnboardingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementOnboardingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_fleet_apps_management_onboarding.test_onboarding"
	datasourceName := "data.oci_fleet_apps_management_onboardings.test_onboardings"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementOnboardingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_onboarding", "test_onboarding", acctest.Optional, acctest.Create, FleetAppsManagementOnboardingRepresentation), "fleetappsmanagement", "onboarding", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementOnboardingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_onboarding", "test_onboarding", acctest.Required, acctest.Create, FleetAppsManagementOnboardingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementOnboardingResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementOnboardingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_onboarding", "test_onboarding", acctest.Optional, acctest.Create, FleetAppsManagementOnboardingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_cost_tracking_tag_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_fams_tag_enabled", "false"),

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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_onboardings", "test_onboardings", acctest.Optional, acctest.Update, FleetAppsManagementOnboardingDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementOnboardingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_onboarding", "test_onboarding", acctest.Optional, acctest.Update, FleetAppsManagementOnboardingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "onboarding_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "onboarding_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "onboarding_collection.0.items.#", "1"),
			),
		},
	})
}
