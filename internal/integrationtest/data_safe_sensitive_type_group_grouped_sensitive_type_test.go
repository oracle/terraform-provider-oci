// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSensitiveTypeGroupGroupedSensitiveTypeDataSourceRepresentation = map[string]interface{}{
		"sensitive_type_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_type_group.test_sensitive_type_group.id}`},
		"sensitive_type_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.sensitive_type_id}`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSensitiveTypeGroupGroupedSensitiveTypeDataSourceFilterRepresentation}}
	DataSafeSensitiveTypeGroupGroupedSensitiveTypeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_sensitive_type_group_grouped_sensitive_type.test_sensitive_type_group_grouped_sensitive_type.id}`}},
	}

	DataSafeSensitiveTypeGroupGroupedSensitiveTypeRepresentation = map[string]interface{}{
		"sensitive_type_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_type_group.test_sensitive_type_group.id}`},
		"patch_operations":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataSafeSensitiveTypeGroupGroupedSensitiveTypePatchOperationsRepresentation},
	}
	DataSafeSensitiveTypeGroupGroupedSensitiveTypePatchOperationsRepresentation = map[string]interface{}{
		"operation": acctest.Representation{RepType: acctest.Required, Create: `INSERT`},
		"selection": acctest.Representation{RepType: acctest.Required, Create: `items`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: map[string]string{"sensitiveTypeId": `${var.sensitive_type_id}`}},
	}

	DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type_group", "test_sensitive_type_group", acctest.Required, acctest.Create, DataSafeSensitiveTypeGroupRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSensitiveTypeGroupGroupedSensitiveTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSensitiveTypeGroupGroupedSensitiveTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	sensitiveTypeId := utils.GetEnvSettingWithBlankDefault("sensitive_type_id")
	sensitiveTypeIdVariableStr := fmt.Sprintf("variable \"sensitive_type_id\" { default = \"%s\" }\n", sensitiveTypeId)

	resourceName := "oci_data_safe_sensitive_type_group_grouped_sensitive_type.test_sensitive_type_group_grouped_sensitive_type"
	datasourceName := "data.oci_data_safe_sensitive_type_group_grouped_sensitive_types.test_sensitive_type_group_grouped_sensitive_types"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+sensitiveTypeIdVariableStr+DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type_group_grouped_sensitive_type", "test_sensitive_type_group_grouped_sensitive_type", acctest.Required, acctest.Create, DataSafeSensitiveTypeGroupGroupedSensitiveTypeRepresentation), "datasafe", "sensitiveTypeGroupGroupedSensitiveType", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + sensitiveTypeIdVariableStr + DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type_group_grouped_sensitive_type", "test_sensitive_type_group_grouped_sensitive_type", acctest.Optional, acctest.Create, DataSafeSensitiveTypeGroupGroupedSensitiveTypeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_type_group_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_type_group_grouped_sensitive_types", "test_sensitive_type_group_grouped_sensitive_types", acctest.Optional, acctest.Update, DataSafeSensitiveTypeGroupGroupedSensitiveTypeDataSourceRepresentation) +
				compartmentIdVariableStr + sensitiveTypeIdVariableStr + DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type_group_grouped_sensitive_type", "test_sensitive_type_group_grouped_sensitive_type", acctest.Optional, acctest.Update, DataSafeSensitiveTypeGroupGroupedSensitiveTypeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_type_id"),

				resource.TestCheckResourceAttr(datasourceName, "grouped_sensitive_type_collection.#", "1"),
			),
		},
		// verify resource import
		{
			Config: config + DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type_group_grouped_sensitive_type", "test_sensitive_type_group_grouped_sensitive_type", acctest.Required, acctest.Create, DataSafeSensitiveTypeGroupGroupedSensitiveTypeRepresentation),
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"patch_operations",
			},
			ResourceName: resourceName,
		},
	})
}
