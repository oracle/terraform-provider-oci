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
	FleetAppsManagementCompliancePolicySingularDataSourceRepresentation = map[string]interface{}{
		"compliance_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compliance_policy_id}`},
	}

	FleetAppsManagementCompliancePolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.compliance_policy_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	FleetAppsManagementCompliancePolicyResourceConfig = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementCompliancePolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementCompliancePolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compliancePolicyId := utils.GetEnvSettingWithBlankDefault("compliance_policy_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compliancePolicyIdVariableStr := fmt.Sprintf("variable \"compliance_policy_id\" { default = \"%s\" }\n", compliancePolicyId)

	datasourceName := "data.oci_fleet_apps_management_compliance_policies.test_compliance_policies"
	singularDatasourceName := "data.oci_fleet_apps_management_compliance_policy.test_compliance_policy"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_compliance_policies", "test_compliance_policies", acctest.Required, acctest.Create, FleetAppsManagementCompliancePolicyDataSourceRepresentation) +
				compartmentIdVariableStr + compliancePolicyIdVariableStr + FleetAppsManagementCompliancePolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				// Below criteria checks collection is returned
				resource.TestCheckResourceAttrSet(datasourceName, "compliance_policy_collection.#"),
				// Considering at least one compliance rule exists in entire tenancy, below criteria checks collection is non-empty
				resource.TestCheckResourceAttr(datasourceName, "compliance_policy_collection.0.%", "1"),
				// Verify each entity has total 11 attributes
				resource.TestCheckResourceAttr(datasourceName, "compliance_policy_collection.0.items.0.%", "11"),

				//Check for some values
				resource.TestCheckResourceAttr(datasourceName, "compliance_policy_collection.0.items.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "compliance_policy_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "compliance_policy_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttr(datasourceName, "compliance_policy_collection.0.items.0.state", "ACTIVE"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_compliance_policy", "test_compliance_policy", acctest.Required, acctest.Create, FleetAppsManagementCompliancePolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + compliancePolicyIdVariableStr + FleetAppsManagementCompliancePolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compliance_policy_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "product_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
