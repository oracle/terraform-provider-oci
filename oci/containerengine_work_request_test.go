// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	workRequestDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"cluster_id":     Representation{repType: Optional, create: `${oci_containerengine_cluster.test_cluster.id}`},
		"resource_id":    Representation{repType: Optional, create: `${oci_containerengine_cluster.test_cluster.id}`},
		"resource_type":  Representation{repType: Optional, create: `CLUSTER`},
		"status":         Representation{repType: Optional, create: []string{}},
	}

	WorkRequestResourceConfig = generateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Required, Create, clusterRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{repType: Required, create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{repType: Required, create: `10.0.20.0/24`}, "dns_label": Representation{repType: Required, create: `cluster1`}})) +
		generateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{repType: Required, create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{repType: Required, create: `10.0.21.0/24`}, "dns_label": Representation{repType: Required, create: `cluster2`}})) +
		AvailabilityDomainConfig +
		generateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", Required, Create, clusterOptionSingularDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		}))
)

func TestContainerengineWorkRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineWorkRequestResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_containerengine_work_requests.test_work_requests"

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
					generateDataSourceFromRepresentationMap("oci_containerengine_work_requests", "test_work_requests", Optional, Create, workRequestDataSourceRepresentation) +
					compartmentIdVariableStr + WorkRequestResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "cluster_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "work_requests.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "work_requests.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "work_requests.0.operation_type"),
					resource.TestCheckResourceAttr(datasourceName, "work_requests.0.resources.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "work_requests.0.status"),
					resource.TestCheckResourceAttrSet(datasourceName, "work_requests.0.time_accepted"),
					resource.TestCheckResourceAttrSet(datasourceName, "work_requests.0.time_finished"),
					resource.TestCheckResourceAttrSet(datasourceName, "work_requests.0.time_started"),
				),
			},
		},
	})
}
