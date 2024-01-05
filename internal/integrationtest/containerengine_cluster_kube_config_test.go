// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ContainerengineContainerengineClusterKubeConfigSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"endpoint":      acctest.Representation{RepType: acctest.Optional, Create: `LEGACY_KUBERNETES`},
		"token_version": acctest.Representation{RepType: acctest.Optional, Create: `2.0.0`},
	}

	ContainerengineClusterKubeConfigResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.20.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster1`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.21.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster2`}})) +
		AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, ContainerengineContainerengineClusterOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		}))
)

// issue-routing-tag: containerengine/default
func TestContainerengineClusterKubeConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterKubeConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_containerengine_cluster_kube_config.test_cluster_kube_config"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_kube_config", "test_cluster_kube_config", acctest.Optional, acctest.Create, ContainerengineContainerengineClusterKubeConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterKubeConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoint", "LEGACY_KUBERNETES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "token_version", "2.0.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
			),
		},
	})
}
