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
	DataSafemaskingAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"group_by":                  acctest.Representation{RepType: acctest.Optional, Create: `targetId`},
		"masking_policy_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_masking_policy.test_masking_policy.id}`},
		"sensitive_type_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sensitive_type.test_sensitive_type.id}`},
		"target_database_group_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_target_database_group.test_target_database_group.id}`},
		"target_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_target.test_target.id}`},
	}

	DataSafeMaskingAnalyticResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_analytics", "test_masking_analytics", acctest.Required, acctest.Create, DataSafemaskingAnalyticDataSourceRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeMaskingAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeMaskingAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_masking_analytics.test_masking_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				compartmentIdVariableStr + DataSafeMaskingAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "masking_analytics_collection.#"),
			),
		},
	})
}
