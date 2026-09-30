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
	DataSafeSubsettingPolicySubsettingSchemaObjectDataSourceRepresentation = map[string]interface{}{
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"object":               acctest.Representation{RepType: acctest.Required, Create: []string{`EMPLOYEES`}},
		"schema_name":          acctest.Representation{RepType: acctest.Required, Create: []string{`HR_TEST`}},
	}

	DataSafeSubsettingPolicySubsettingSchemaObjectResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Required, acctest.Create, DataSafeSubsettingPolicyRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSubsettingPolicySubsettingSchemaObjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingPolicySubsettingSchemaObjectResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_subsetting_policy_subsetting_schema_objects.test_subsetting_policy_subsetting_schema_objects"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_schema_objects", "test_subsetting_policy_subsetting_schema_objects", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingSchemaObjectDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingSchemaObjectResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "object.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "schema_name.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_policy_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_schema_object_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "subsetting_schema_object_collection.0.items.#", "1"),
			),
		},
	})
}
