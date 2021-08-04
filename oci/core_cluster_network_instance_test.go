// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	clusterNetworkInstanceDataSourceRepresentation = map[string]interface{}{
		"cluster_network_id": Representation{repType: Required, create: `${oci_core_cluster_network.test_cluster_network.id}`},
		"compartment_id":     Representation{repType: Required, create: `${var.compartment_id}`},
	}

	ClusterNetworkInstanceResourceConfig = ClusterNetworkResourceConfig
)

// issue-routing-tag: core/computeManagement
func TestCoreClusterNetworkInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreClusterNetworkInstanceResource_basic")
	defer httpreplay.SaveScenario()
	if !strings.Contains(getEnvSettingWithBlankDefault("enabled_tests"), "ClusterNetwork") {
		t.Skip("ClusterNetwork test not supported due to limited host capacity")
	}
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_cluster_network_instances.test_cluster_network_instances"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_cluster_network_instances", "test_cluster_network_instances", Required, Create, clusterNetworkInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + ClusterNetworkInstanceResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "cluster_network_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "instances.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.fault_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.instance_configuration_id"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.load_balancer_backends.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.region"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.time_created"),
				),
			},
		},
	})
}
