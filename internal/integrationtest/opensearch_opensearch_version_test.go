// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OpensearchOpensearchVersionSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	OpensearchOpensearchVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	OpensearchOpensearchVersionResourceConfig = ""
)

// issue-routing-tag: opensearch/default
func TestOpensearchOpensearchVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpensearchOpensearchVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_opensearch_opensearch_versions.test_opensearch_versions"
	singularDatasourceName := "data.oci_opensearch_opensearch_version.test_opensearch_version"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opensearch_opensearch_versions", "test_opensearch_versions", acctest.Required, acctest.Create, OpensearchOpensearchVersionDataSourceRepresentation) +
				compartmentIdVariableStr + OpensearchOpensearchVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "opensearch_versions_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "opensearch_versions_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opensearch_opensearch_version", "test_opensearch_version", acctest.Required, acctest.Create, OpensearchOpensearchVersionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OpensearchOpensearchVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
			),
		},
	})
}
