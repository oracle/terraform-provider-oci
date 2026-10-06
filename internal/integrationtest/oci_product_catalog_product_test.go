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
	OciProductCatalogProductSingularDataSourceRepresentation = map[string]interface{}{
		// Use ID from the plural data source instead of referencing a managed resource
		"product_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_oci_product_catalog_products.test_products.products.0.id}`},
	}

	// Create the plural products data source without filters; it will list products
	OciProductCatalogProductResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_oci_product_catalog_products", "test_products", acctest.Required, acctest.Create, map[string]interface{}{})
)

// issue-routing-tag: oci_product_catalog/default
func TestOciProductCatalogProductResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOciProductCatalogProductResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_oci_product_catalog_product.test_product"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_oci_product_catalog_product", "test_product", acctest.Required, acctest.Create, OciProductCatalogProductSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OciProductCatalogProductResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "product_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_excluded"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_name"),
				// Only verify that at least one SKU is returned; exact count varies by product
				resource.TestCheckResourceAttrSet(singularDatasourceName, "skus.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				// time_launched/time_ready may be absent when the product is not in an ACTIVE/READY state.
			),
		},
	})
}
