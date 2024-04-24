// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ContainerengineContainerengineNodePoolOptionSingularDataSourceRepresentation = map[string]interface{}{
		"node_pool_option_id": acctest.Representation{RepType: acctest.Required, Create: `all`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	ContainerengineNodePoolOptionResourceConfig = ""
)

// issue-routing-tag: containerengine/default
func TestContainerengineNodePoolOptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineNodePoolOptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_containerengine_node_pool_option.test_node_pool_option"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool_option", "test_node_pool_option", acctest.Required, acctest.Create, ContainerengineContainerengineNodePoolOptionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineNodePoolOptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool_option", "test_node_pool_option", acctest.Optional, acctest.Create, ContainerengineContainerengineNodePoolOptionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineNodePoolOptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_pool_option_id"),

				resource.TestMatchResourceAttr(singularDatasourceName, "images.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestMatchResourceAttr(singularDatasourceName, "kubernetes_versions.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestMatchResourceAttr(singularDatasourceName, "shapes.#", regexp.MustCompile("[1-9][0-9]*")),
			),
		},
	})
}
