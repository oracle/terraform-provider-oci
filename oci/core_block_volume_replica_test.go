// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	blockVolumeReplicaSingularDataSourceRepresentation = map[string]interface{}{
		"block_volume_replica_id": Representation{RepType: Required, Create: `${data.oci_core_block_volume_replicas.test_block_volume_replicas.block_volume_replicas.0.id}`},
	}
	blockVolumeReplicaDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{RepType: Required, Create: `NyKp:US-ASHBURN-AD-1`},
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":        Representation{RepType: Optional, Create: `displayName`},
		"state":               Representation{RepType: Optional, Create: `AVAILABLE`},
	}

	dependenceVolumeRepresentation = map[string]interface{}{
		"availability_domain":            Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                 Representation{RepType: Required, Create: `${var.compartment_id}`},
		"block_volume_replicas":          RepresentationGroup{Optional, dependenceVolumeBlockVolumeReplicasRepresentation},
		"block_volume_replicas_deletion": Representation{RepType: Optional, Create: `false`, Update: `true`},
	}

	//hardcode availability_domain here to meet the cross region replicas requirement
	dependenceVolumeBlockVolumeReplicasRepresentation = map[string]interface{}{
		"availability_domain": Representation{RepType: Required, Create: `NyKp:US-ASHBURN-AD-1`},
		"display_name":        Representation{RepType: Optional, Create: `displayName`},
	}

	BlockVolumeReplicaResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: core/blockStorage
func TestCoreBlockVolumeReplicaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreBlockVolumeReplicaResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume.test_volume"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// Create volume and enable replicas
		{
			Config: config +
				GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Create, dependenceVolumeRepresentation) +
				compartmentIdVariableStr + BlockVolumeReplicaResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					time.Sleep(2 * time.Minute)
					return
				},
			),
		},

		{
			Config: config +
				GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Create, dependenceVolumeRepresentation) +
				compartmentIdVariableStr + BlockVolumeReplicaResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "block_volume_replicas.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "block_volume_replicas.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "block_volume_replicas.0.block_volume_replica_id"),
				resource.TestCheckResourceAttr(resourceName, "block_volume_replicas.0.display_name", "displayName"),
			),
		},
		// disabled replicas
		{
			Config: config +
				GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Update,
					RepresentationCopyWithRemovedNestedProperties("block_volume_replicas", dependenceVolumeRepresentation)) +
				compartmentIdVariableStr + BlockVolumeReplicaResourceConfig,
		},
	})
}
