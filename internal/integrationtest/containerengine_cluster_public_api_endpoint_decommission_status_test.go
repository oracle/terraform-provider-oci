// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ContainerengineClusterPublicApiEndpointDecommissionStatusSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
	}
	ContainerengineClusterDecommissionManagerRepresentation = map[string]interface{}{
		"is_public_api_endpoint_decommissioned": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"cluster_id":                            acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"rollback_deadline_delay":               acctest.Representation{RepType: acctest.Optional, Create: `P1D`},
	}

	ContainerengineClusterDecommissionManagerUpdateRepresentation = map[string]interface{}{
		"is_public_api_endpoint_decommissioned": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"cluster_id":                            acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
	}

	ContainerengineV1hClusterRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-2]}`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `name`},
		"vcn_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	ContainerengineV2ClusterRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-2]}`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `name`},
		"vcn_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"endpoint_config":    acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineClusterEndpointConfigRepresentation},
	}

	ContainerengineClusterDecommissionSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
	}
	ContainerengineClusterPublicApiEndpointDecommissionStatusResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineClusterRepresentation)
	ContainerengineClusterDecommissionResourceDependencies                  = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.20.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster1`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.21.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster2`}})) +
		AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, ContainerengineContainerengineClusterOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.22.0/24`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		KeyResourceDependencyConfig
)

// issue-routing-tag: containerengine/default
func TestContainerengineClusterPublicApiEndpointDecommissionStatusResource_fullFlow(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterPublicApiEndpointDecommissionStatusResource_fullFlow")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_containerengine_cluster.test_cluster"
	managerResourceName := "oci_containerengine_cluster_public_api_endpoint_decommission_manager.test_manager"
	statusDatasourceName := "data.oci_containerengine_cluster_public_api_endpoint_decommission_status.test_status"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// 1. Create v1h cluster
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterDecommissionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineV1hClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
			),
		},
		// 2. Migrate cluster from v1h to v2
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterDecommissionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineV2ClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_config.0.subnet_id"),
			),
		},
		// 3. Call decommission manager to decommission cluster (set is_public_api_endpoint_decommissioned=true)
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterDecommissionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineV2ClusterRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_public_api_endpoint_decommission_manager", "test_manager", acctest.Required, acctest.Create, ContainerengineClusterDecommissionManagerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(managerResourceName, "is_public_api_endpoint_decommissioned", "true"),
			),
		},
		// 4. Verify decommission status (expect status DECOMMISSIONED or IN_PROGRESS)
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterDecommissionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineV2ClusterRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_public_api_endpoint_decommission_manager", "test_manager", acctest.Required, acctest.Create, ContainerengineClusterDecommissionManagerRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_public_api_endpoint_decommission_status", "test_status", acctest.Required, acctest.Create, ContainerengineClusterDecommissionSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(statusDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttr(statusDatasourceName, "status", "DECOMMISSIONED"),
				resource.TestCheckResourceAttrSet(statusDatasourceName, "time_decommission_rollback_deadline"),
			),
		},
		// 5. Extend deadline
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterDecommissionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineV2ClusterRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_public_api_endpoint_decommission_manager", "test_manager", acctest.Optional, acctest.Create, ContainerengineClusterDecommissionManagerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(managerResourceName, "is_public_api_endpoint_decommissioned", "true"),
				resource.TestCheckResourceAttrSet(managerResourceName, "rollback_deadline_delay"),
			),
		},
		// 6. Rollback
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterDecommissionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineV2ClusterRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_public_api_endpoint_decommission_manager", "test_manager", acctest.Required, acctest.Create, ContainerengineClusterDecommissionManagerUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(managerResourceName, "is_public_api_endpoint_decommissioned", "false"),
			),
		},
		// 7. back to pending state
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterDecommissionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineV2ClusterRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_public_api_endpoint_decommission_manager", "test_manager", acctest.Required, acctest.Create, ContainerengineClusterDecommissionManagerUpdateRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_public_api_endpoint_decommission_status", "test_status", acctest.Required, acctest.Create, ContainerengineClusterDecommissionSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(statusDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttr(statusDatasourceName, "status", "PENDING"),
			),
		},
	})
}
