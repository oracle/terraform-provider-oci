// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ClusterKubeConfigResourceConfig = ClusterKubeConfigResourceDependencies + `

`
	ClusterKubeConfigPropertyVariables = `
variable "cluster_kube_config_expiration" { default = 2592000 }
variable "cluster_kube_config_token_version" { default = "1.0.0" }

`
	ClusterKubeConfigResourceDependencies = ClusterPropertyVariables + ClusterRequiredOnlyResource
)

func TestContainerengineClusterKubeConfigResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_containerengine_cluster_kube_config.test_cluster_kube_config"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `
variable "cluster_kube_config_expiration" { default = 2592000 }
variable "cluster_kube_config_token_version" { default = "1.0.0" }

data "oci_containerengine_cluster_kube_config" "test_cluster_kube_config" {
	#Required
	cluster_id = "${oci_containerengine_cluster.test_cluster.id}"

	#Optional
	expiration = "${var.cluster_kube_config_expiration}"
	token_version = "${var.cluster_kube_config_token_version}"
}
                ` + compartmentIdVariableStr + ClusterKubeConfigResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "expiration", "2592000"),
					resource.TestCheckResourceAttr(singularDatasourceName, "token_version", "1.0.0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				),
			},
		},
	})
}
