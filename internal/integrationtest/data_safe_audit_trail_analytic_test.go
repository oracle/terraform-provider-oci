// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafeauditTrailAnalyticSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"group_by":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`lifecycleState`}},
		"target_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.target_ocid}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditTrailAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeAuditTrailAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetOcid := utils.GetEnvSettingWithBlankDefault("target_ocid")
	targetOcidVariableStr := fmt.Sprintf("variable \"target_ocid\" { default = \"%s\" }\n", targetOcid)

	singularDatasourceName := "data.oci_data_safe_audit_trail_analytic.test_audit_trail_analytic"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_trail_analytic", "test_audit_trail_analytic", acctest.Optional, acctest.Create, DataSafeauditTrailAnalyticSingularDataSourceRepresentation) +
				compartmentIdVariableStr + targetOcidVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.0.count.#", "0"),
			),
		},
	})
}
