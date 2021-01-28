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
		"block_volume_replica_id": Representation{repType: Required, create: `${data.oci_core_block_volume_replicas.test_block_volume_replicas.block_volume_replicas.0.id}`},
	}
	blockVolumeReplicaDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `NyKp:US-ASHBURN-AD-1`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":        Representation{repType: Optional, create: `displayName`},
		"state":               Representation{repType: Optional, create: `AVAILABLE`},
	}

	dependenceVolumeRepresentation = map[string]interface{}{
		"availability_domain":            Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                 Representation{repType: Required, create: `${var.compartment_id}`},
		"block_volume_replicas":          RepresentationGroup{Optional, dependenceVolumeBlockVolumeReplicasRepresentation},
		"block_volume_replicas_deletion": Representation{repType: Optional, create: `false`, update: `true`},
	}

	//hardcode availability_domain here to meet the cross region replicas requirement
	dependenceVolumeBlockVolumeReplicasRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `NyKp:US-ASHBURN-AD-1`},
		"display_name":        Representation{repType: Optional, create: `displayName`},
	}

	BlockVolumeReplicaResourceConfig = AvailabilityDomainConfig
)

func TestCoreBlockVolumeReplicaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreBlockVolumeReplicaResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume.test_volume"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},

		Steps: []resource.TestStep{
			// create volume and enable replicas
			{
				Config: config +
					generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Create, dependenceVolumeRepresentation) +
					compartmentIdVariableStr + BlockVolumeReplicaResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					func(s *terraform.State) (err error) {
						time.Sleep(2 * time.Minute)
						return
					},
				),
			},

			{
				Config: config +
					generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Create, dependenceVolumeRepresentation) +
					compartmentIdVariableStr + BlockVolumeReplicaResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Update,
						representationCopyWithRemovedNestedProperties("block_volume_replicas", dependenceVolumeRepresentation)) +
					compartmentIdVariableStr + BlockVolumeReplicaResourceConfig,
			},
		},
	})
}
