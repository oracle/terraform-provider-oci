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
	bootVolumeReplicaSingularDataSourceRepresentation = map[string]interface{}{
		"boot_volume_replica_id": Representation{repType: Required, create: `${data.oci_core_boot_volume_replicas.test_boot_volume_replicas.boot_volume_replicas.0.id}`},
	}

	bootVolumeReplicaDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":        Representation{repType: Optional, create: `displayName`},
		"state":               Representation{repType: Optional, create: `AVAILABLE`},
	}
	dependenceBootVolumeRepresentation = map[string]interface{}{
		"availability_domain":           Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                Representation{repType: Required, create: `${var.compartment_id}`},
		"source_details":                RepresentationGroup{Required, bootVolumeSourceDetailsRepresentation},
		"display_name":                  Representation{repType: Optional, create: `boot volume with replica`, update: `boot volume without replica`},
		"boot_volume_replicas":          RepresentationGroup{Optional, dependenceBootVolumeReplicasRepresentation},
		"boot_volume_replicas_deletion": Representation{repType: Optional, create: `false`, update: `true`},
	}
	dependenceBootVolumeReplicasRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `NyKp:US-ASHBURN-AD-1`},
		"display_name":        Representation{repType: Optional, create: `displayName`},
	}
	BootVolumeReplicaResourceConfig = BootVolumeResourceDependencies
)

func TestCoreBootVolumeReplicaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreBootVolumeReplicaResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_boot_volume.test_boot_volume"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// create volume and enable replicas
			{
				Config: config +
					generateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Optional, Create, dependenceBootVolumeRepresentation) +
					compartmentIdVariableStr + BootVolumeReplicaResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					func(s *terraform.State) (err error) {
						time.Sleep(2 * time.Minute)
						return
					},
				),
			},

			{
				Config: config +
					generateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Optional, Create, dependenceBootVolumeRepresentation) +
					compartmentIdVariableStr + BootVolumeReplicaResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "boot_volume_replicas.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "boot_volume_replicas.0.availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "boot_volume_replicas.0.boot_volume_replica_id"),
					resource.TestCheckResourceAttr(resourceName, "boot_volume_replicas.0.display_name", "displayName"),
				),
			},
			// disabled replicas
			{
				Config: config +
					generateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Optional, Update,
						representationCopyWithRemovedNestedProperties("boot_volume_replicas", dependenceBootVolumeRepresentation)) +
					compartmentIdVariableStr + BootVolumeReplicaResourceConfig,
			},
		},
	})
}
