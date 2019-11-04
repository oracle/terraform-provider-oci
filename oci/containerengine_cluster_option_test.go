// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	clusterOptionSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_option_id": Representation{repType: Required, create: `all`},
		"compartment_id":    Representation{repType: Optional, create: `${var.compartment_id}`},
	}

	ClusterOptionResourceConfig = ""
)

func TestContainerengineClusterOptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterOptionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_containerengine_cluster_option.test_cluster_option"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", Required, Create, clusterOptionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ClusterOptionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_option_id"),

					resource.TestMatchResourceAttr(singularDatasourceName, "kubernetes_versions.#", regexp.MustCompile("[1-9][0-9]*")),
				),
			},
			// verify singular datasource with compartment_id
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", Optional, Create, clusterOptionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ClusterOptionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_option_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

					resource.TestMatchResourceAttr(singularDatasourceName, "kubernetes_versions.#", regexp.MustCompile("[1-9][0-9]*")),
				),
			},
		},
	})
}
