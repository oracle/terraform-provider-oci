// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	CoreCoreBootVolumeReplicaSingularDataSourceRepresentation = map[string]interface{}{
		"boot_volume_replica_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_boot_volume_replicas.test_boot_volume_replicas.boot_volume_replicas.0.id}`},
	}

	CoreCoreBootVolumeReplicaDataSourceRepresentation = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"volume_group_replica_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_volume_group_replica.test_volume_group_replica.id}`},
	}

	CoreCoreBootVolumeDependenceBootVolumeRepresentation = map[string]interface{}{
		"availability_domain":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"source_details":                acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreBootVolumeSourceDetailsRepresentation},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `boot volume with replica`, Update: `boot volume without replica`},
		"boot_volume_replicas":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreCoreBootVolumeDependenceBootVolumeReplicasRepresentation},
		"boot_volume_replicas_deletion": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	CoreCoreBootVolumeDependenceBootVolumeReplicasRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `KvuH:US-ASHBURN-AD-1`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}
	CoreBootVolumeReplicaResourceConfig = CoreBootVolumeResourceDependencies
)

// issue-routing-tag: core/blockStorage
func TestCoreBootVolumeReplicaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreBootVolumeReplicaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_boot_volume.test_boot_volume"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create volume and enable replicas
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Optional, acctest.Create, CoreCoreBootVolumeDependenceBootVolumeRepresentation) +
				compartmentIdVariableStr + CoreBootVolumeReplicaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					time.Sleep(2 * time.Minute)
					return
				},
			),
		},

		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Optional, acctest.Create, CoreCoreBootVolumeDependenceBootVolumeRepresentation) +
				compartmentIdVariableStr + CoreBootVolumeReplicaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedNestedProperties("boot_volume_replicas", CoreCoreBootVolumeDependenceBootVolumeRepresentation)) +
				compartmentIdVariableStr + CoreBootVolumeReplicaResourceConfig,
		},
	})
}
