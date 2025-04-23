// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DblmPatchManagementSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_release":                      acctest.Representation{RepType: acctest.Optional, Create: `19.0.0.0.0`},
		"time_started_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2006-01-02T15:04:05Z`},
		"time_started_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2026-01-02T15:04:05Z`},
	}

	DblmPatchManagementResourceConfig = ""
)

// issue-routing-tag: dblm/default
func TestDblmPatchManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDblmPatchManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_dblm_patch_management.test_patch_management"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dblm_patch_management", "test_patch_management", acctest.Optional, acctest.Create, DblmPatchManagementSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DblmPatchManagementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_release", "19.0.0.0.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_started_greater_than_or_equal_to", "2006-01-02T15:04:05Z"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_started_less_than", "2026-01-02T15:04:05Z"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "images_patch_recommendation_summary.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_operations_summary.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resources_patch_compliance_summary.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
	})
}
