// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	nodePoolOptionSingularDataSourceRepresentation = map[string]interface{}{
		"node_pool_option_id": Representation{repType: Required, create: `all`},
		"compartment_id":      Representation{repType: Optional, create: `${var.compartment_id}`},
	}

	NodePoolOptionResourceConfig = ""
)

func TestContainerengineNodePoolOptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineNodePoolOptionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_containerengine_node_pool_option.test_node_pool_option"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_containerengine_node_pool_option", "test_node_pool_option", Required, Create, nodePoolOptionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + NodePoolOptionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_containerengine_node_pool_option", "test_node_pool_option", Optional, Create, nodePoolOptionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + NodePoolOptionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "node_pool_option_id"),

					resource.TestMatchResourceAttr(singularDatasourceName, "images.#", regexp.MustCompile("[1-9][0-9]*")),
					resource.TestMatchResourceAttr(singularDatasourceName, "kubernetes_versions.#", regexp.MustCompile("[1-9][0-9]*")),
					resource.TestMatchResourceAttr(singularDatasourceName, "shapes.#", regexp.MustCompile("[1-9][0-9]*")),
				),
			},
		},
	})
}
