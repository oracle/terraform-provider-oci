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
	clusterKubeConfigSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_id":    Representation{RepType: Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"endpoint":      Representation{RepType: Optional, Create: `LEGACY_KUBERNETES`},
		"token_version": Representation{RepType: Optional, Create: `2.0.0`},
	}

	ClusterKubeConfigResourceConfig = GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Required, Create, clusterRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", Required, Create, RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{RepType: Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{RepType: Required, Create: `10.0.20.0/24`}, "dns_label": Representation{RepType: Required, Create: `cluster1`}})) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", Required, Create, RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{RepType: Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{RepType: Required, Create: `10.0.21.0/24`}, "dns_label": Representation{RepType: Required, Create: `cluster2`}})) +
		AvailabilityDomainConfig +
		GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", Required, Create, clusterOptionSingularDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{RepType: Required, Create: `dnslabel`},
		}))
)

// issue-routing-tag: containerengine/default
func TestContainerengineClusterKubeConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterKubeConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_containerengine_cluster_kube_config.test_cluster_kube_config"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_kube_config", "test_cluster_kube_config", Optional, Create, clusterKubeConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ClusterKubeConfigResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoint", "LEGACY_KUBERNETES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "token_version", "2.0.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
			),
		},
	})
}
