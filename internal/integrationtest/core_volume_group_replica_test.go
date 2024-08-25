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
	CoreVolumeGroupReplicationRepresentation = map[string]interface{}{
		"availability_domain":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"source_details":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVolumeGroupSourceDetailsRepresentation},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"volume_group_replicas":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreVolumeGroupVolumeGroupReplicasRepresentation},
		"volume_group_replicas_deletion": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"preserve_volume_replica":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreIgnoreSystemTagsChangesRepresentation},
	}

	//hardcode availability_domain here to meet the cross region replicas requirement
	CoreVolumeGroupVolumeGroupReplicasRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `pjBI:US-ASHBURN-AD-1`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"xrr_kms_key_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_ocid_cross_region}`},
	}
)

// issue-routing-tag: core/blockStorage
func TestCoreVolumeGroupReplicaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeGroupReplicaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume_group.test_volume_group"

	createConfig := config + VolumeGroupRequiredOnlyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", acctest.Optional, acctest.Create, CoreVolumeGroupReplicationRepresentation) +
		compartmentIdVariableStr

	updateConfig := config + VolumeGroupRequiredOnlyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", acctest.Optional, acctest.Update,
			acctest.RepresentationCopyWithRemovedProperties(CoreVolumeGroupReplicationRepresentation, []string{"volume_group_replicas"})) +
		compartmentIdVariableStr

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create Volume Group with Replication
		{
			Config: createConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(resourceName, "volume_group_replicas.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "volume_group_replicas.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "volume_group_replicas.0.display_name"),

				func(s *terraform.State) (err error) {
					time.Sleep(35 * time.Minute)
					return
				},
			),
		},
		// Disable replication without preserve replica
		{
			Config: updateConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "state", "UPDATE_PENDING"),
				func(s *terraform.State) (err error) {
					time.Sleep(10 * time.Minute)
					return
				},
			),
		},
		// Delete
		{
			Config:  updateConfig,
			Destroy: true,
		},
	})
}
