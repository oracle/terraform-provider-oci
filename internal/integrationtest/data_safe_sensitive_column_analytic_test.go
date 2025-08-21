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
	DataSafeSensitiveColumnAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"column_name":               acctest.Representation{RepType: acctest.Optional, Create: []string{`FIRST_NAME`}},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"group_by":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`groupBy`}},
		"object":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`object`}},
		"schema_name":               acctest.Representation{RepType: acctest.Optional, Create: []string{`schemaName`}},
		"sensitive_data_model_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model.id}`},
		"sensitive_type_group_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sensitive_type_group.test_sensitive_type_group.id}`},
		"sensitive_type_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sensitive_type.test_sensitive_type.id}`},
		"target_database_group_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_target_database_group.test_target_database_group.id}`},
		"target_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_target.test_target.id}`},
	}

	DataSafeSensitiveColumnAnalyticResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_column_analytics", "test_sensitive_column_analytics", acctest.Required, acctest.Create, DataSafeSensitiveColumnAnalyticDataSourceRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSensitiveColumnAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSensitiveColumnAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_sensitive_column_analytics.test_sensitive_column_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSensitiveColumnAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_column_analytics_collection.#"),
			),
		},
	})
}
