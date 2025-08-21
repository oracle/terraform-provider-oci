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
	DataSafeAttributeSetAssociatedResourceDataSourceRepresentation = map[string]interface{}{
		"attribute_set_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_attribute_set.test_attribute_set.id}`},
	}

	DataSafeAttributeSetAssociatedResourceResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_attribute_set", "test_attribute_set", acctest.Required, acctest.Create, DataSafeAttributeSetRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeAttributeSetAssociatedResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeAttributeSetAssociatedResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_attribute_set_associated_resources.test_attribute_set_associated_resources"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_attribute_set_associated_resources", "test_attribute_set_associated_resources", acctest.Required, acctest.Create, DataSafeAttributeSetAssociatedResourceDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeAttributeSetAssociatedResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "attribute_set_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "associated_resource_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "associated_resource_collection.0.items.#", "0"),
			),
		},
	})
}
