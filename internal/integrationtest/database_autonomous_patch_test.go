// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseDatabaseAutonomousPatchSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_patch_id": acctest.Representation{RepType: acctest.Required, Create: "LATEST"},
	}

	DatabaseAutonomousPatchResourceConfig = ""
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousPatchResource_basic(t *testing.T) {
	t.Skip("Skip this test till the MR resource test is skipped since the patch id is reliably fetched from MR resource")
	httpreplay.SetScenario("TestDatabaseAutonomousPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_autonomous_patch.test_autonomous_patch"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_patch", "test_autonomous_patch", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousPatchSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_patch_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_patch_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_model"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "quarter"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_released"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "year"),
			),
		},
	})
}
