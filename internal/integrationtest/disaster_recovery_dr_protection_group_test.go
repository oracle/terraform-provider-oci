// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DisasterRecoveryDrProtectionGroupRequiredOnlyResource = DisasterRecoveryDrProtectionGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Required, acctest.Create, DisasterRecoveryDrProtectionGroupRepresentation)

	DisasterRecoveryDrProtectionGroupResourceConfig = DisasterRecoveryDrProtectionGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Optional, acctest.Update, DisasterRecoveryDrProtectionGroupRepresentation)

	DisasterRecoveryDisasterRecoveryDrProtectionGroupSingularDataSourceRepresentation = map[string]interface{}{
		"dr_protection_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_disaster_recovery_dr_protection_group.test_dr_protection_group.id}`},
	}

	DisasterRecoveryDisasterRecoveryDrProtectionGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `My DR Protection Group`, Update: `displayName2`},
		"dr_protection_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_disaster_recovery_dr_protection_group.test_dr_protection_group.id}`},
		"role":                   acctest.Representation{RepType: acctest.Optional, Create: `PRIMARY`},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DisasterRecoveryDrProtectionGroupDataSourceFilterRepresentation}}
	DisasterRecoveryDrProtectionGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_disaster_recovery_dr_protection_group.test_dr_protection_group.id}`}},
	}

	DisasterRecoveryDrProtectionGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `My DR Protection Group`, Update: `displayName2`},
		"log_location":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DisasterRecoveryDrProtectionGroupLogLocationRepresentation},
		"association":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DisasterRecoveryDrProtectionGroupAssociationRepresentation},
		"members":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DisasterRecoveryDrProtectionGroupMembersRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: DefinedTagsIgnoreRepresentation},
	}

	DisasterRecoveryPeerDrProtectionGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `peerDisplayName`},
		"log_location":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DisasterRecoveryDrProtectionGroupLogLocationRepresentation},
	}

	DefinedTagsIgnoreRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DisasterRecoveryDrProtectionGroupLogLocationRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_bucket.test_bucket.name}`, Update: `${data.oci_objectstorage_bucket.test_bucket2.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`, Update: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}
	DisasterRecoveryDrProtectionGroupAssociationRepresentation = map[string]interface{}{
		"role":        acctest.Representation{RepType: acctest.Required, Create: `PRIMARY`},
		"peer_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_disaster_recovery_dr_protection_group.test_peer.id}`},
		"peer_region": acctest.Representation{RepType: acctest.Optional, Create: `${var.region}`},
	}

	DisasterRecoveryDrProtectionGroupMembersRepresentation = map[string]interface{}{
		"member_id":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_volume_groups.test_volume_groups.volume_groups.0.id}`},
		"member_type": acctest.Representation{RepType: acctest.Required, Create: `VOLUME_GROUP`},
	}

	DisasterRecoveryDrProtectionGroupWithComputeMemberConfig = `
	resource "oci_disaster_recovery_dr_protection_group" "test_dr_protection_group" {
		  #Required
		  compartment_id = var.compartment_id
		  display_name   = "My DR Protection Group"
		  log_location {
		   		#Required
		   		bucket    = data.oci_objectstorage_bucket.test_bucket.name
		   		namespace = data.oci_objectstorage_namespace.test_namespace.namespace
		  }
          members {
		   		member_id = data.oci_core_instances.dr_instances.instances[0].id
		   		member_type = "COMPUTE_INSTANCE_MOVABLE"
		  }
    }
	`

	ObjectStorageBucketDependencyConfig = `
	data "oci_objectstorage_namespace" "test_namespace" {
  		#Optional
  		compartment_id = var.compartment_id
	}
	data "oci_objectstorage_bucket" "test_bucket" {
  		namespace = data.oci_objectstorage_namespace.test_namespace.namespace
  		name      = "testBucketName"
	}
	data "oci_objectstorage_bucket" "test_bucket2" {
  		namespace = data.oci_objectstorage_namespace.test_namespace.namespace
  		name      = "testBucketName_1"
	}
	`

	VolumeGroupDependencyConfig = `
	data "oci_core_volume_groups" "test_volume_groups" {
  		#Required
  		compartment_id = var.compartment_id

  		#Optional
  		availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  		display_name        = "example-volume-group"
  		state               = "AVAILABLE"
	}
	`
	ComputeInstanceDependencyConfig = `
	data "oci_core_instances" "dr_instances" {
	  	#Required
	  	compartment_id = var.compartment_id
	
	  	#Optional
	  	display_name = "example-instance"
	}
	`

	DisasterRecoveryDrProtectionGroupResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_peer", acctest.Optional, acctest.Create, DisasterRecoveryPeerDrProtectionGroupRepresentation) +
		ObjectStorageBucketDependencyConfig +
		VolumeGroupDependencyConfig +
		ComputeInstanceDependencyConfig +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: disaster_recovery/default
func TestDisasterRecoveryDrProtectionGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDisasterRecoveryDrProtectionGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	var region = utils.GetEnvSettingWithBlankDefault("region")

	resourceName := "oci_disaster_recovery_dr_protection_group.test_dr_protection_group"
	datasourceName := "data.oci_disaster_recovery_dr_protection_groups.test_dr_protection_groups"
	singularDatasourceName := "data.oci_disaster_recovery_dr_protection_group.test_dr_protection_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DisasterRecoveryDrProtectionGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Optional, acctest.Create, DisasterRecoveryDrProtectionGroupRepresentation), "disasterrecovery", "drProtectionGroup", t)

	acctest.ResourceTest(t, testAccCheckDisasterRecoveryDrProtectionGroupDestroy, []resource.TestStep{
		// Create Dependencies
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrProtectionGroupResourceDependencies,
		},
		// verify Create DR Protection Group without members
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrProtectionGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Required, acctest.Create, acctest.RepresentationCopyWithRemovedProperties(DisasterRecoveryDrProtectionGroupRepresentation, []string{"members"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My DR Protection Group"),
				resource.TestCheckResourceAttr(resourceName, "log_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "log_location.0.bucket", "testBucketName"),
				resource.TestCheckResourceAttrSet(resourceName, "log_location.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "members.#", "0"),
				resource.TestCheckNoResourceAttr(resourceName, "members.0.member_id"),
				resource.TestCheckNoResourceAttr(resourceName, "members.0.member_type"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Create with Volume Group as member of DR Protection Group
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrProtectionGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Required, acctest.Create, DisasterRecoveryDrProtectionGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My DR Protection Group"),
				resource.TestCheckResourceAttr(resourceName, "log_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "log_location.0.bucket", "testBucketName"),
				resource.TestCheckResourceAttrSet(resourceName, "log_location.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "members.0.member_id"),
				resource.TestCheckResourceAttr(resourceName, "members.0.member_type", "VOLUME_GROUP"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// Update member of DR Protection Group: Remove Volume Group, Add Compute Instance
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrProtectionGroupResourceDependencies +
				DisasterRecoveryDrProtectionGroupWithComputeMemberConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My DR Protection Group"),
				resource.TestCheckResourceAttr(resourceName, "log_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "log_location.0.bucket", "testBucketName"),
				resource.TestCheckResourceAttrSet(resourceName, "log_location.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "members.0.member_id"),
				resource.TestCheckResourceAttr(resourceName, "members.0.member_type", "COMPUTE_INSTANCE_MOVABLE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Remove members from DR Protection Group
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrProtectionGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Required, acctest.Create, acctest.RepresentationCopyWithRemovedProperties(DisasterRecoveryDrProtectionGroupRepresentation, []string{"members"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My DR Protection Group"),
				resource.TestCheckResourceAttr(resourceName, "log_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "log_location.0.bucket", "testBucketName"),
				resource.TestCheckResourceAttrSet(resourceName, "log_location.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "members.#", "0"),
				resource.TestCheckNoResourceAttr(resourceName, "members.0.member_id"),
				resource.TestCheckNoResourceAttr(resourceName, "members.0.member_type"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrProtectionGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrProtectionGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Optional, acctest.Create, DisasterRecoveryDrProtectionGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "association.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "association.0.peer_id"),
				resource.TestCheckResourceAttr(resourceName, "association.0.peer_region", region),
				resource.TestCheckResourceAttr(resourceName, "association.0.role", "PRIMARY"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My DR Protection Group"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "log_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "log_location.0.bucket", "testBucketName"),
				resource.TestCheckResourceAttrSet(resourceName, "log_location.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "members.0.member_id"),
				resource.TestCheckResourceAttr(resourceName, "members.0.member_type", "VOLUME_GROUP"),
				resource.TestCheckResourceAttrSet(resourceName, "role"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DisasterRecoveryDrProtectionGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DisasterRecoveryDrProtectionGroupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "association.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "association.0.peer_id"),
				resource.TestCheckResourceAttr(resourceName, "association.0.peer_region", region),
				resource.TestCheckResourceAttr(resourceName, "association.0.role", "PRIMARY"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My DR Protection Group"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "log_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "log_location.0.bucket", "testBucketName"),
				resource.TestCheckResourceAttrSet(resourceName, "log_location.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "members.0.member_id"),
				resource.TestCheckResourceAttr(resourceName, "members.0.member_type", "VOLUME_GROUP"),
				resource.TestCheckResourceAttrSet(resourceName, "role"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrProtectionGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Optional, acctest.Update, DisasterRecoveryDrProtectionGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "association.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "association.0.peer_id"),
				resource.TestCheckResourceAttr(resourceName, "association.0.peer_region", region),
				resource.TestCheckResourceAttr(resourceName, "association.0.role", "PRIMARY"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "log_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "log_location.0.bucket", "testBucketName_1"),
				resource.TestCheckResourceAttrSet(resourceName, "log_location.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "members.0.member_id"),
				resource.TestCheckResourceAttr(resourceName, "members.0.member_type", "VOLUME_GROUP"),
				resource.TestCheckResourceAttrSet(resourceName, "role"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_disaster_recovery_dr_protection_groups", "test_dr_protection_groups", acctest.Optional, acctest.Update, DisasterRecoveryDisasterRecoveryDrProtectionGroupDataSourceRepresentation) +
				compartmentIdVariableStr + DisasterRecoveryDrProtectionGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Optional, acctest.Update, DisasterRecoveryDrProtectionGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "dr_protection_group_id"),
				resource.TestCheckResourceAttr(datasourceName, "role", "PRIMARY"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "dr_protection_group_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "dr_protection_group_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Required, acctest.Create, DisasterRecoveryDisasterRecoveryDrProtectionGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DisasterRecoveryDrProtectionGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dr_protection_group_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "life_cycle_details"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_location.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_location.0.bucket", "testBucketName_1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_location.0.namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "members.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "members.0.member_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_region"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "role"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DisasterRecoveryDrProtectionGroupRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"association",
			},
			ResourceName: resourceName,
		},
		// Disassociate
		{
			Config: config +
				compartmentIdVariableStr + DisasterRecoveryDrProtectionGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(DisasterRecoveryDrProtectionGroupRepresentation, map[string]interface{}{
					"disassociate_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					time.Sleep(1 * time.Minute)
					return
				},
			),
		},
	})
}

func testAccCheckDisasterRecoveryDrProtectionGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DisasterRecoveryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_disaster_recovery_dr_protection_group" {
			noResourceFound = false
			request := oci_disaster_recovery.GetDrProtectionGroupRequest{}

			tmp := rs.Primary.ID
			request.DrProtectionGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "disaster_recovery")

			response, err := client.GetDrProtectionGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_disaster_recovery.DrProtectionGroupLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DisasterRecoveryDrProtectionGroup") {
		resource.AddTestSweepers("DisasterRecoveryDrProtectionGroup", &resource.Sweeper{
			Name:         "DisasterRecoveryDrProtectionGroup",
			Dependencies: acctest.DependencyGraph["drProtectionGroup"],
			F:            sweepDisasterRecoveryDrProtectionGroupResource,
		})
	}
}

func sweepDisasterRecoveryDrProtectionGroupResource(compartment string) error {
	disasterRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DisasterRecoveryClient()
	drProtectionGroupIds, err := getDisasterRecoveryDrProtectionGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, drProtectionGroupId := range drProtectionGroupIds {
		if ok := acctest.SweeperDefaultResourceId[drProtectionGroupId]; !ok {
			deleteDrProtectionGroupRequest := oci_disaster_recovery.DeleteDrProtectionGroupRequest{}

			deleteDrProtectionGroupRequest.DrProtectionGroupId = &drProtectionGroupId

			deleteDrProtectionGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "disaster_recovery")
			_, error := disasterRecoveryClient.DeleteDrProtectionGroup(context.Background(), deleteDrProtectionGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting DrProtectionGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", drProtectionGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &drProtectionGroupId, DisasterRecoveryDrProtectionGroupSweepWaitCondition, time.Duration(3*time.Minute),
				DisasterRecoveryDrProtectionGroupSweepResponseFetchOperation, "disaster_recovery", true)
		}
	}
	return nil
}

func getDisasterRecoveryDrProtectionGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DrProtectionGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	disasterRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DisasterRecoveryClient()

	listDrProtectionGroupsRequest := oci_disaster_recovery.ListDrProtectionGroupsRequest{}
	listDrProtectionGroupsRequest.CompartmentId = &compartmentId
	listDrProtectionGroupsRequest.LifecycleState = oci_disaster_recovery.ListDrProtectionGroupsLifecycleStateActive
	listDrProtectionGroupsResponse, err := disasterRecoveryClient.ListDrProtectionGroups(context.Background(), listDrProtectionGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DrProtectionGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, drProtectionGroup := range listDrProtectionGroupsResponse.Items {
		id := *drProtectionGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DrProtectionGroupId", id)
	}
	return resourceIds, nil
}

func DisasterRecoveryDrProtectionGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if drProtectionGroupResponse, ok := response.Response.(oci_disaster_recovery.GetDrProtectionGroupResponse); ok {
		return drProtectionGroupResponse.LifecycleState != oci_disaster_recovery.DrProtectionGroupLifecycleStateDeleted
	}
	return false
}

func DisasterRecoveryDrProtectionGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DisasterRecoveryClient().GetDrProtectionGroup(context.Background(), oci_disaster_recovery.GetDrProtectionGroupRequest{
		DrProtectionGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
