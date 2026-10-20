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
	DataSafeSubsettingPolicySubsettingSchemaDataSourceRepresentation = map[string]interface{}{
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"is_derived_schema":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"schema_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`HR_TEST`}},
	}

	DataSafeSubsettingPolicySubsettingSchemaResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Required, acctest.Create, DataSafeSubsettingPolicyRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSubsettingPolicySubsettingSchemaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingPolicySubsettingSchemaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_subsetting_policy_subsetting_schemas.test_subsetting_policy_subsetting_schemas"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_schemas", "test_subsetting_policy_subsetting_schemas", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingSchemaDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingSchemaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_policy_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_schema_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "subsetting_schema_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subsetting_schema_collection.0.items.0.is_derived", "false"),
				resource.TestCheckResourceAttr(datasourceName, "subsetting_schema_collection.0.items.0.schema_name", "HR_TEST"),
			),
		},
	})
}
