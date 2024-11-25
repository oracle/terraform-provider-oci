// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetAppsManagementComplianceRecordCountDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	FleetAppsManagementComplianceRecordCountResourceConfig = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementComplianceRecordCountResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementComplianceRecordCountResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_fleet_apps_management_compliance_record_counts.test_compliance_record_counts"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_compliance_record_counts", "test_compliance_record_counts", acctest.Required, acctest.Create, FleetAppsManagementComplianceRecordCountDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementComplianceRecordCountResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "compliance_record_aggregation_collection.#"),
				resource.TestMatchResourceAttr(datasourceName, "compliance_record_aggregation_collection.0.items.#",
					regexp.MustCompile("[1-9][0-9]*")),
			),
		},
	})
}
