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
	catalogTypeSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_id": Representation{repType: Required, create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"type_key":   Representation{repType: Required, create: `${data.oci_datacatalog_catalog_types.test_catalog_types.type_collection.0.items.0.key}`},
		"fields":     Representation{repType: Optional, create: []string{}},
	}

	catalogTypeDataSourceRepresentation = map[string]interface{}{
		"catalog_id":    Representation{repType: Required, create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"state":         Representation{repType: Optional, create: `ACTIVE`},
		"type_category": Representation{repType: Optional, create: `dataAsset`, update: `connection`},
	}

	CatalogTypeResourceConfig = generateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Required, Create, catalogRepresentation)
)

// issue-routing-tag: datacatalog/default
func TestDatacatalogCatalogTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogCatalogTypeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datacatalog_catalog_types.test_catalog_types"
	singularDatasourceName := "data.oci_datacatalog_catalog_type.test_catalog_type"

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
					generateDataSourceFromRepresentationMap("oci_datacatalog_catalog_types", "test_catalog_types", Optional, Create, catalogTypeDataSourceRepresentation) +
					compartmentIdVariableStr + CatalogTypeResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttrSet(datasourceName, "type_collection.0.items.0.name"),
					resource.TestCheckResourceAttrSet(datasourceName, "type_collection.0.items.0.description"),
					resource.TestCheckResourceAttrSet(datasourceName, "type_collection.0.items.0.uri"),
					resource.TestCheckResourceAttrSet(datasourceName, "type_collection.0.items.0.state"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_datacatalog_catalog_types", "test_catalog_types", Optional, Create,
						representationCopyWithNewProperties(catalogTypeDataSourceRepresentation, map[string]interface{}{
							"name": Representation{repType: Optional, create: `Oracle Database`}})) +
					generateDataSourceFromRepresentationMap("oci_datacatalog_catalog_type", "test_catalog_type", Optional, Create, catalogTypeSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CatalogTypeResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "fields.#", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "type_key"),
					resource.TestCheckResourceAttr(singularDatasourceName, "properties.%", "16"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				),
			},
		},
	})
}
