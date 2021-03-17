// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	flexComponentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"name":           Representation{repType: Optional, create: `Exadata.X8M.StorageServer`},
	}

	FlexComponentResourceConfig = ""
)

func TestDatabaseFlexComponentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseFlexComponentResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_flex_components.test_flex_components"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_flex_components", "test_flex_components", Optional, Create, flexComponentDataSourceRepresentation) +
					compartmentIdVariableStr + FlexComponentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "name", "Exadata.X8M.StorageServer"),

					resource.TestCheckResourceAttrSet(datasourceName, "flex_component_collection.#"),
					resource.TestCheckResourceAttr(datasourceName, "flex_component_collection.0.items.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "flex_component_collection.0.items.0.available_core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "flex_component_collection.0.items.0.available_db_storage_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "flex_component_collection.0.items.0.minimum_core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "flex_component_collection.0.items.0.name"),
				),
			},
		},
	})
}
