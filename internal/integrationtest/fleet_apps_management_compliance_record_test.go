// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetAppsManagementComplianceRecordDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compliance_state": acctest.Representation{RepType: acctest.Optional, Create: `NON_COMPLIANT`},
		"entity_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.test_active_fleet}`},
		"product_name":     acctest.Representation{RepType: acctest.Optional, Create: `Oracle Linux`},
		"product_stack":    acctest.Representation{RepType: acctest.Optional, Create: `Oracle Linux`},
	}

	FleetAppsManagementComplianceRecordResourceConfig = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementComplianceRecordResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementComplianceRecordResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Fleet in ACTIVE state. Fleets require a confirmation action call not supported by Terraform to go active.
	// Thus, this needs to be created and confirmed manually.
	activeFleetId := utils.GetEnvSettingWithBlankDefault("test_active_fleet")
	activeFleetStr := fmt.Sprintf("variable \"test_active_fleet\" { default = \"%s\" }\n", activeFleetId)

	datasourceName := "data.oci_fleet_apps_management_compliance_records.test_compliance_records"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_compliance_records", "test_compliance_records", acctest.Optional, acctest.Create, FleetAppsManagementComplianceRecordDataSourceRepresentation) +
				activeFleetStr + compartmentIdVariableStr + FleetAppsManagementComplianceRecordResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compliance_state", "NON_COMPLIANT"),
				resource.TestCheckResourceAttrSet(datasourceName, "entity_id"),
				resource.TestCheckResourceAttr(datasourceName, "product_name", "Oracle Linux"),
				resource.TestCheckResourceAttr(datasourceName, "product_stack", "Oracle Linux"),

				resource.TestCheckResourceAttrSet(datasourceName, "compliance_record_collection.#"),
				resource.TestMatchResourceAttr(datasourceName, "compliance_record_collection.0.items.#",
					regexp.MustCompile("[1-9][0-9]*")),
			),
		},
	})
}
