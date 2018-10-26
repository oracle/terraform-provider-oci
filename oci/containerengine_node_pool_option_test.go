// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	nodePoolOptionSingularDataSourceRepresentation = map[string]interface{}{
		"node_pool_option_id": Representation{repType: Required, create: `all`},
	}

	NodePoolOptionResourceConfig = ""
)

func TestContainerengineNodePoolOptionResource_basic(t *testing.T) {
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
				),
			},
		},
	})
}
