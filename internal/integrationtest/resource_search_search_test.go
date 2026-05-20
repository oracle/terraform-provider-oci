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
	ResourceSearchSearchSingularDataSourceRepresentation = map[string]interface{}{
		"query": acctest.Representation{RepType: acctest.Required, Create: `query all resources`},
	}

	ResourceSearchSearchDataSourceRepresentation = map[string]interface{}{
		"query": acctest.Representation{RepType: acctest.Required, Create: `query all resources`},
	}

	ResourceSearchSearchResourceConfig = ""
)

// issue-routing-tag: resource_search/default
func TestResourceSearchSearchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceSearchSearchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_resource_search.test_search"
	singularDatasourceName := "data.oci_resource_search.test_search"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resource_search", "test_search", acctest.Required, acctest.Create, ResourceSearchSearchDataSourceRepresentation) +
				compartmentIdVariableStr + ResourceSearchSearchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "query", "query all resources"),
				resource.TestCheckResourceAttrSet(datasourceName, "results.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "results.0.identifier"),
				resource.TestCheckResourceAttrSet(datasourceName, "results.0.resource_type"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resource_search", "test_search", acctest.Required, acctest.Create, ResourceSearchSearchSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ResourceSearchSearchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "query", "query all resources"),
				resource.TestCheckResourceAttrSet(datasourceName, "results.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "results.0.identifier"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "results.0.resource_type"),
			),
		},
	})
}
