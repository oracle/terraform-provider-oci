// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	workRequestDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"cluster_id":     Representation{RepType: Optional, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"resource_id":    Representation{RepType: Optional, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"resource_type":  Representation{RepType: Optional, Create: `CLUSTER`},
		"status":         Representation{RepType: Optional, Create: []string{}},
	}

	WorkRequestResourceConfig = GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Required, Create, clusterRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", Required, Create, RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{RepType: Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{RepType: Required, Create: `10.0.20.0/24`}, "dns_label": Representation{RepType: Required, Create: `cluster1`}})) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", Required, Create, RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{RepType: Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{RepType: Required, Create: `10.0.21.0/24`}, "dns_label": Representation{RepType: Required, Create: `cluster2`}})) +
		AvailabilityDomainConfig +
		GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", Required, Create, clusterOptionSingularDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{RepType: Required, Create: `dnslabel`},
		}))
)

// issue-routing-tag: containerengine/default
func TestContainerengineWorkRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineWorkRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_containerengine_work_requests.test_work_requests"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_containerengine_work_requests", "test_work_requests", Optional, Create, workRequestDataSourceRepresentation) +
				compartmentIdVariableStr + WorkRequestResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
	})
}
