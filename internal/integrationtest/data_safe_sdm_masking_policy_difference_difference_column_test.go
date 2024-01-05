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
	DataSafeSdmMaskingPolicyDifferenceDifferenceColumnSingularDataSourceRepresentation = map[string]interface{}{
		"difference_column_key":            acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"sdm_masking_policy_difference_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sdm_masking_policy_difference.test_sdm_masking_policy_difference.id}`},
	}

	DataSafeSdmMaskingPolicyDifferenceDifferenceColumnDataSourceRepresentation = map[string]interface{}{
		"sdm_masking_policy_difference_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sdm_masking_policy_difference.test_sdm_masking_policy_difference.id}`},
		"column_name":                      acctest.Representation{RepType: acctest.Optional, Create: []string{`columnName`}},
		"difference_type":                  acctest.Representation{RepType: acctest.Optional, Create: `ALL`},
		"object":                           acctest.Representation{RepType: acctest.Optional, Create: []string{`object`}},
		"planned_action":                   acctest.Representation{RepType: acctest.Optional, Create: `SYNC`},
		"schema_name":                      acctest.Representation{RepType: acctest.Optional, Create: []string{`schemaName`}},
		"sync_status":                      acctest.Representation{RepType: acctest.Optional, Create: `SYNCED`},
	}

	DataSafeSdmMaskingPolicyDifferenceDifferenceColumnResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy", "test_masking_policy", acctest.Required, acctest.Create, maskingPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_difference", "test_sdm_masking_policy_difference", acctest.Required, acctest.Create, DataSafeSdmMaskingPolicyDifferenceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model1", acctest.Optional, acctest.Create, sensitiveDataModelRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSdmMaskingPolicyDifferenceDifferenceColumnResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSdmMaskingPolicyDifferenceDifferenceColumnResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	sensitiveTypeId := utils.GetEnvSettingWithBlankDefault("sensitive_type_id")
	sensitiveTypeIdVariableStr := fmt.Sprintf("variable \"sensitive_type_id\" { default = \"%s\" }\n", sensitiveTypeId)

	datasourceName := "data.oci_data_safe_sdm_masking_policy_difference_difference_columns.test_sdm_masking_policy_difference_difference_columns"
	singularDatasourceName := "data.oci_data_safe_sdm_masking_policy_difference_difference_column.test_sdm_masking_policy_difference_difference_column"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + targetIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_difference_difference_columns", "test_sdm_masking_policy_difference_difference_columns", acctest.Required, acctest.Create, DataSafeSdmMaskingPolicyDifferenceDifferenceColumnDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSdmMaskingPolicyDifferenceDifferenceColumnResourceConfig + sensitiveTypeIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "sdm_masking_policy_difference_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sdm_masking_policy_difference_column_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "sdm_masking_policy_difference_column_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sdm_masking_policy_difference_column_collection.0.items.0.schema_name", "HCM"),
				resource.TestCheckResourceAttr(datasourceName, "sdm_masking_policy_difference_column_collection.0.items.0.column_name", "FIRST_NAME"),
				resource.TestCheckResourceAttr(datasourceName, "sdm_masking_policy_difference_column_collection.0.items.0.object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(datasourceName, "sdm_masking_policy_difference_column_collection.0.items.0.difference_type", "NEW"),
				resource.TestCheckResourceAttr(datasourceName, "sdm_masking_policy_difference_column_collection.0.items.0.planned_action", "SYNC"),
			),
		},
		// verify singular datasource
		{
			Config: config + targetIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_difference_difference_column", "test_sdm_masking_policy_difference_difference_column", acctest.Required, acctest.Create, DataSafeSdmMaskingPolicyDifferenceDifferenceColumnSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSdmMaskingPolicyDifferenceDifferenceColumnResourceConfig + sensitiveTypeIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sdm_masking_policy_difference_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "column_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "difference_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "planned_action"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schema_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sensitive_columnkey"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sensitive_type_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sync_status"),
			),
		},
	})
}
