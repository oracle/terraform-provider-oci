// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	workRequestDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cluster_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"resource_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"resource_type":  acctest.Representation{RepType: acctest.Optional, Create: `CLUSTER`},
		"status":         acctest.Representation{RepType: acctest.Optional, Create: []string{}},
	}

	WorkRequestResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, clusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.20.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster1`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.21.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster2`}})) +
		AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, clusterOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		}))
)

// issue-routing-tag: containerengine/default
func TestContainerengineWorkRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineWorkRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_containerengine_work_requests.test_work_requests"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_work_requests", "test_work_requests", acctest.Optional, acctest.Create, workRequestDataSourceRepresentation) +
				compartmentIdVariableStr + WorkRequestResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
