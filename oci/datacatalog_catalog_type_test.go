// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	catalogTypeSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_id": Representation{RepType: Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"type_key":   Representation{RepType: Required, Create: `${data.oci_datacatalog_catalog_types.test_catalog_types.type_collection.0.items.0.key}`},
		"fields":     Representation{RepType: Optional, Create: []string{}},
	}

	catalogTypeDataSourceRepresentation = map[string]interface{}{
		"catalog_id":    Representation{RepType: Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"state":         Representation{RepType: Optional, Create: `ACTIVE`},
		"type_category": Representation{RepType: Optional, Create: `dataAsset`, Update: `connection`},
	}

	CatalogTypeResourceConfig = GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Required, Create, catalogRepresentation)
)

// issue-routing-tag: datacatalog/default
func TestDatacatalogCatalogTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogCatalogTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datacatalog_catalog_types.test_catalog_types"
	singularDatasourceName := "data.oci_datacatalog_catalog_type.test_catalog_type"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalog_types", "test_catalog_types", Optional, Create, catalogTypeDataSourceRepresentation) +
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
				GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalog_types", "test_catalog_types", Optional, Create,
					RepresentationCopyWithNewProperties(catalogTypeDataSourceRepresentation, map[string]interface{}{
						"name": Representation{RepType: Optional, Create: `Oracle Database`}})) +
				GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalog_type", "test_catalog_type", Optional, Create, catalogTypeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CatalogTypeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fields.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type_key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "properties.%", "16"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
			),
		},
	})
}
