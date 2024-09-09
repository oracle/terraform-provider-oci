// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetAppsManagementOnboardingPolicyDataSourceRepresentation = map[string]interface{}{}

	FleetAppsManagementOnboardingPolicyResourceConfig = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementOnboardingPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementOnboardingPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_fleet_apps_management_onboarding_policies.test_onboarding_policies"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_onboarding_policies", "test_onboarding_policies", acctest.Required, acctest.Create, FleetAppsManagementOnboardingPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementOnboardingPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "onboarding_policy_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "onboarding_policy_collection.0.items.#", "1"),
			),
		},
	})
}
