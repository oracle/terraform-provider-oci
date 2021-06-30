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
	migrateToNativeVCNSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_id": Representation{repType: Required, create: `${oci_containerengine_cluster.test_cluster.id}`},
	}
)

func TestContainerengineMigrateToNativeVcnStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineMigrateToNativeVcnStatusResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_cluster.test_cluster"
	singularDatasourceName := "data.oci_containerengine_migrate_to_native_vcn_status.test_migrate_to_native_vcn_status"

	var resId, resId2 string

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckContainerengineClusterDestroy,

		Steps: []resource.TestStep{
			// create V1 Cluster
			{
				Config: config + compartmentIdVariableStr + ClusterResourceDependencies + generateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Required, Create, representationCopyWithRemovedProperties(clusterRepresentation, []string{"kms_key_id", "options", "image_policy_config"})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify V1 Cluster migrates to V2
			{
				Config: config + compartmentIdVariableStr + ClusterResourceDependencies + generateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Optional, Update, representationCopyWithRemovedProperties(clusterRepresentation, []string{"kms_key_id", "options", "image_policy_config"})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "endpoint_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "endpoint_config.0.is_public_ip_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "endpoint_config.0.subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "endpoint_config.0.nsg_ids.#"),
					resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify singular datasource
			{
				Config: config + compartmentIdVariableStr + ClusterResourceDependencies + generateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Optional, Update, representationCopyWithRemovedProperties(clusterRepresentation, []string{"kms_key_id", "options", "image_policy_config"})) + generateDataSourceFromRepresentationMap(
					"oci_containerengine_migrate_to_native_vcn_status", "test_migrate_to_native_vcn_status",
					Optional, Create, migrateToNativeVCNSingularDataSourceRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_decommission_scheduled"),
				),
			},
		},
	})
}
