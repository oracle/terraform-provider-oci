// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	nodePoolOptionSingularDataSourceRepresentation = map[string]interface{}{
		"node_pool_option_id": Representation{RepType: Required, Create: `all`},
		"compartment_id":      Representation{RepType: Optional, Create: `${var.compartment_id}`},
	}

	NodePoolOptionResourceConfig = ""
)

// issue-routing-tag: containerengine/default
func TestContainerengineNodePoolOptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineNodePoolOptionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_containerengine_node_pool_option.test_node_pool_option"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool_option", "test_node_pool_option", Required, Create, nodePoolOptionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NodePoolOptionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_pool_option_id"),

				resource.TestMatchResourceAttr(singularDatasourceName, "images.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestMatchResourceAttr(singularDatasourceName, "kubernetes_versions.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestMatchResourceAttr(singularDatasourceName, "shapes.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestMatchResourceAttr(singularDatasourceName, "sources.#", regexp.MustCompile("[1-9][0-9]*")),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool_option", "test_node_pool_option", Optional, Create, nodePoolOptionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NodePoolOptionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_pool_option_id"),

				resource.TestMatchResourceAttr(singularDatasourceName, "images.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestMatchResourceAttr(singularDatasourceName, "kubernetes_versions.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestMatchResourceAttr(singularDatasourceName, "shapes.#", regexp.MustCompile("[1-9][0-9]*")),
			),
		},
	})
}
