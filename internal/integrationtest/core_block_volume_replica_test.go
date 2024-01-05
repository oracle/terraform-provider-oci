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
	CoreCoreBlockVolumeReplicaSingularDataSourceRepresentation = map[string]interface{}{
		"block_volume_replica_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_block_volume_replicas.test_block_volume_replicas.block_volume_replicas.0.id}`},
	}
	CoreCoreBlockVolumeReplicaDataSourceRepresentation = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"volume_group_replica_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_volume_group_replica.test_volume_group_replica.id}`},
	}

	CoreBlockDependenceVolumeRepresentation = map[string]interface{}{
		"availability_domain":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"block_volume_replicas":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreBlockDependenceVolumeBlockVolumeReplicasRepresentation},
		"block_volume_replicas_deletion": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	//hardcode availability_domain here to meet the cross region replicas requirement
	CoreBlockDependenceVolumeBlockVolumeReplicasRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `KvuH:US-ASHBURN-AD-1`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	CoreBlockVolumeReplicaResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: core/blockStorage
func TestCoreBlockVolumeReplicaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreBlockVolumeReplicaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume.test_volume"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create volume and enable replicas
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Optional, acctest.Create, CoreBlockDependenceVolumeRepresentation) +
				compartmentIdVariableStr + CoreBlockVolumeReplicaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					time.Sleep(2 * time.Minute)
					return
				},
			),
		},

		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Optional, acctest.Create, CoreBlockDependenceVolumeRepresentation) +
				compartmentIdVariableStr + CoreBlockVolumeReplicaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedNestedProperties("block_volume_replicas", CoreBlockDependenceVolumeRepresentation)) +
				compartmentIdVariableStr + CoreBlockVolumeReplicaResourceConfig,
		},
	})
}
