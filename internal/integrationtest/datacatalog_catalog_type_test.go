// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	catalogTypeSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"type_key":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_datacatalog_catalog_types.test_catalog_types.type_collection.0.items.0.key}`},
		"fields":     acctest.Representation{RepType: acctest.Optional, Create: []string{}},
	}

	catalogTypeDataSourceRepresentation = map[string]interface{}{
		"catalog_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"state":         acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type_category": acctest.Representation{RepType: acctest.Optional, Create: `dataAsset`, Update: `connection`},
	}

	CatalogTypeResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", acctest.Required, acctest.Create, catalogRepresentation)
)

// issue-routing-tag: datacatalog/default
func TestDatacatalogCatalogTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogCatalogTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datacatalog_catalog_types.test_catalog_types"
	singularDatasourceName := "data.oci_datacatalog_catalog_type.test_catalog_type"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalog_types", "test_catalog_types", acctest.Optional, acctest.Create, catalogTypeDataSourceRepresentation) +
				compartmentIdVariableStr + CatalogTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalog_types", "test_catalog_types", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(catalogTypeDataSourceRepresentation, map[string]interface{}{
						"name": acctest.Representation{RepType: acctest.Optional, Create: `Oracle Database`}})) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalog_type", "test_catalog_type", acctest.Optional, acctest.Create, catalogTypeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CatalogTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fields.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type_key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "properties.%", "16"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
			),
		},
	})
}
