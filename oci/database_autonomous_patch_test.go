// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	autonomousPatchSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_patch_id": Representation{repType: Required, create: "LATEST"},
	}

	AutonomousPatchResourceConfig = ""
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousPatchResource_basic(t *testing.T) {
	t.Skip("Skip this test till the MR resource test is skipped since the patch id is reliably fetched from MR resource")
	httpreplay.SetScenario("TestDatabaseAutonomousPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_autonomous_patch.test_autonomous_patch"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_database_autonomous_patch", "test_autonomous_patch", Required, Create, autonomousPatchSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_patch_id"),

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
