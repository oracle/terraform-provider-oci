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
	DataSafeMaskingPolicyReferentialRelationDataSourceRepresentation = map[string]interface{}{
		"masking_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.masking_policy_id}`},
		"relation_type":     acctest.Representation{RepType: acctest.Optional, Create: []string{`DB_DEFINED`}},
		"schema_name":       acctest.Representation{RepType: acctest.Optional, Create: []string{`HR`}},
		"column_name":       acctest.Representation{RepType: acctest.Optional, Create: []string{`EMPLOYEE_ID`}},
		"object":            acctest.Representation{RepType: acctest.Optional, Create: []string{`EMPLOYEES`}},
	}

	DataSafeMaskingPolicyReferentialRelationResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeMaskingPolicyReferentialRelationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeMaskingPolicyReferentialRelationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	maskingPolicyId := utils.GetEnvSettingWithBlankDefault("data_safe_masking_policy_id")
	maskingPolicyIdVariableStr := fmt.Sprintf("variable \"masking_policy_id\" { default = \"%s\" }\n", maskingPolicyId)

	datasourceName := "data.oci_data_safe_masking_policy_referential_relations.test_masking_policy_referential_relations"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_policy_referential_relations", "test_masking_policy_referential_relations", acctest.Optional, acctest.Create, DataSafeMaskingPolicyReferentialRelationDataSourceRepresentation) +
				compartmentIdVariableStr + maskingPolicyIdVariableStr + DataSafeMaskingPolicyReferentialRelationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "masking_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "relation_type.0", "DB_DEFINED"),
				resource.TestCheckResourceAttr(datasourceName, "schema_name.0", "HR"),
				resource.TestCheckResourceAttr(datasourceName, "column_name.0", "EMPLOYEE_ID"),
				resource.TestCheckResourceAttr(datasourceName, "object.0", "EMPLOYEES"),

				resource.TestCheckResourceAttrSet(datasourceName, "masking_policy_referential_relation_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "masking_policy_referential_relation_collection.0.items.#", "5"),
			),
		},
	})
}
