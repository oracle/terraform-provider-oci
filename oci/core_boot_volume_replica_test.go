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
		"boot_volume_replica_id": Representation{RepType: Required, Create: `${data.oci_core_boot_volume_replicas.test_boot_volume_replicas.boot_volume_replicas.0.id}`},
	}

	bootVolumeReplicaDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":        Representation{RepType: Optional, Create: `displayName`},
		"state":               Representation{RepType: Optional, Create: `AVAILABLE`},
	}
	dependenceBootVolumeRepresentation = map[string]interface{}{
		"availability_domain":           Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                Representation{RepType: Required, Create: `${var.compartment_id}`},
		"source_details":                RepresentationGroup{Required, bootVolumeSourceDetailsRepresentation},
		"display_name":                  Representation{RepType: Optional, Create: `boot volume with replica`, Update: `boot volume without replica`},
		"boot_volume_replicas":          RepresentationGroup{Optional, dependenceBootVolumeReplicasRepresentation},
		"boot_volume_replicas_deletion": Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	dependenceBootVolumeReplicasRepresentation = map[string]interface{}{
		"availability_domain": Representation{RepType: Required, Create: `NyKp:US-ASHBURN-AD-1`},
		"display_name":        Representation{RepType: Optional, Create: `displayName`},
	}
	BootVolumeReplicaResourceConfig = BootVolumeResourceDependencies
)

// issue-routing-tag: core/blockStorage
func TestCoreBootVolumeReplicaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreBootVolumeReplicaResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_boot_volume.test_boot_volume"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// Create volume and enable replicas
		{
			Config: config +
				GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Optional, Create, dependenceBootVolumeRepresentation) +
				compartmentIdVariableStr + BootVolumeReplicaResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					time.Sleep(2 * time.Minute)
					return
				},
			),
		},

		{
			Config: config +
				GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Optional, Create, dependenceBootVolumeRepresentation) +
				compartmentIdVariableStr + BootVolumeReplicaResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Optional, Update,
					RepresentationCopyWithRemovedNestedProperties("boot_volume_replicas", dependenceBootVolumeRepresentation)) +
				compartmentIdVariableStr + BootVolumeReplicaResourceConfig,
		},
	})
}
