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
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BootVolumeWaitConditionDuration = time.Duration(20 * time.Minute)

	BootVolumeRequiredOnlyResource = CoreBootVolumeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Required, acctest.Create, CoreBootVolumeRepresentation)

	BootVolumeOptionalResource = CoreBootVolumeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Optional, acctest.Create, CoreBootVolumeRepresentation)

	BootVolumeResourceConfig = CoreBootVolumeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Optional, acctest.Update, CoreBootVolumeRepresentation)

	CoreCoreBootVolumeSingularDataSourceRepresentation = map[string]interface{}{
		"boot_volume_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_boot_volume.test_boot_volume.id}`},
	}

	CoreCoreBootVolumeDataSourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreBootVolumeDataSourceFilterRepresentation}}
	CoreBootVolumeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_boot_volume.test_boot_volume.id}`}},
	}

	CoreBootVolumeRepresentation = map[string]interface{}{
		"availability_domain":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"source_details":       acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreBootVolumeSourceDetailsRepresentation},
		"backup_policy_id":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id}`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"kms_key_id":           acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"size_in_gbs":          acctest.Representation{RepType: acctest.Optional, Create: `50`, Update: `51`},
		"vpus_per_gb":          acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `10`},
		"autotune_policies":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: bootVolumeAutotunePoliciesRepresentation},
		"is_auto_tune_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
	}
	CoreBootVolumeSourceDetailsRepresentation = map[string]interface{}{
		"id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.boot_volume_id}`},
		"type": acctest.Representation{RepType: acctest.Required, Create: `bootVolume`},
	}
	bootVolumeAutotunePoliciesRepresentation = map[string]interface{}{
		"autotune_type":   acctest.Representation{RepType: acctest.Required, Create: `PERFORMANCE_BASED`, Update: `PERFORMANCE_BASED`},
		"max_vpus_per_gb": acctest.Representation{RepType: acctest.Optional, Create: `20`, Update: `30`},
	}
	CoreBootVolumeBootVolumeReplicasRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`, Update: `availabilityDomain2`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}

	CoreBootVolumeResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		utils.VolumeBackupPolicyDependency +
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", acctest.Required, acctest.Create, CoreVolumeGroupRepresentation) +
		SourceVolumeListDependency +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig + kmsKeyIdCreateVariableStr + kmsKeyIdUpdateVariableStr
)

// issue-routing-tag: core/blockStorage
func TestCoreBootVolumeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreBootVolumeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_boot_volume.test_boot_volume"
	datasourceName := "data.oci_core_boot_volumes.test_boot_volumes"
	singularDatasourceName := "data.oci_core_boot_volume.test_boot_volume"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreBootVolumeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Optional, acctest.Create, CoreBootVolumeRepresentation), "core", "bootVolume", t)

	acctest.ResourceTest(t, testAccCheckCoreBootVolumeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreBootVolumeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Required, acctest.Create, CoreBootVolumeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckNoResourceAttr(resourceName, "backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "bootVolume"),
				resource.TestCheckNoResourceAttr(resourceName, "volume_group_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreBootVolumeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreBootVolumeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Optional, acctest.Create, CoreBootVolumeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.0.autotune_type", "PERFORMANCE_BASED"),
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.0.max_vpus_per_gb", "20"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "bootVolume"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vpus_per_gb", "10"),
				resource.TestCheckNoResourceAttr(resourceName, "volume_group_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreBootVolumeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreBootVolumeRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.0.autotune_type", "PERFORMANCE_BASED"),
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.0.max_vpus_per_gb", "20"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "bootVolume"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vpus_per_gb", "10"),
				resource.TestCheckNoResourceAttr(resourceName, "volume_group_id"),

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
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId, bootVolumeWaitCondition, BootVolumeWaitConditionDuration,
				bootVolumeResponseFetchOperation, "core", false),
			Config: config + compartmentIdVariableStr + CoreBootVolumeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Optional, acctest.Update, CoreBootVolumeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.0.autotune_type", "PERFORMANCE_BASED"),
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.0.max_vpus_per_gb", "30"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "size_in_gbs", "51"),
				resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "bootVolume"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vpus_per_gb", "10"),
				resource.TestCheckNoResourceAttr(resourceName, "volume_group_id"),

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
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId, bootVolumeWaitCondition, BootVolumeWaitConditionDuration,
				bootVolumeResponseFetchOperation, "core", false),
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_boot_volumes", "test_boot_volumes", acctest.Optional, acctest.Update, CoreCoreBootVolumeDataSourceRepresentation) +
				compartmentIdVariableStr + CoreBootVolumeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Optional, acctest.Update, CoreBootVolumeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckNoResourceAttr(datasourceName, "backup_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "boot_volumes.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.autotune_policies.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.autotune_policies.0.autotune_type", "PERFORMANCE_BASED"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.autotune_policies.0.max_vpus_per_gb", "30"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.image_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.is_hydrated"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.size_in_gbs", "51"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.size_in_mbs"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.source_details.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.source_details.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.source_details.0.type", "bootVolume"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.vpus_per_gb", "10"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Required, acctest.Create, CoreCoreBootVolumeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BootVolumeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckNoResourceAttr(singularDatasourceName, "backup_policy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "kms_key_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "autotune_policies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "autotune_policies.0.autotune_type", "PERFORMANCE_BASED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "autotune_policies.0.max_vpus_per_gb", "30"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "image_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_hydrated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "size_in_gbs", "51"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_mbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_details.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.type", "bootVolume"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vpus_per_gb", "10"),
			),
		},
		// verify resource import
		{
			Config:            config + BootVolumeRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"backup_policy_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckCoreBootVolumeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BlockstorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_boot_volume" {
			noResourceFound = false
			request := oci_core.GetBootVolumeRequest{}

			tmp := rs.Primary.ID
			request.BootVolumeId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetBootVolume(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.BootVolumeLifecycleStateTerminated): true,
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

func bootVolumeResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *oci_common.RetryPolicy) error {
	_, err := client.BlockstorageClient().GetBootVolume(context.Background(), oci_core.GetBootVolumeRequest{
		BootVolumeId: resourceId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func bootVolumeWaitCondition(response oci_common.OCIOperationResponse) bool {
	// Only stop if the volume is hydrated
	if bootVolumeResponse, ok := response.Response.(oci_core.GetBootVolumeResponse); ok {
		return *bootVolumeResponse.IsHydrated == false
	}
	return false
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CoreBootVolume") {
		resource.AddTestSweepers("CoreBootVolume", &resource.Sweeper{
			Name:         "CoreBootVolume",
			Dependencies: acctest.DependencyGraph["bootVolume"],
			F:            sweepCoreBootVolumeResource,
		})
	}
}

func sweepCoreBootVolumeResource(compartment string) error {
	blockstorageClient := acctest.GetTestClients(&schema.ResourceData{}).BlockstorageClient()
	bootVolumeIds, err := getCoreBootVolumeIds(compartment)
	if err != nil {
		return err
	}
	for _, bootVolumeId := range bootVolumeIds {
		if ok := acctest.SweeperDefaultResourceId[bootVolumeId]; !ok {
			deleteBootVolumeRequest := oci_core.DeleteBootVolumeRequest{}

			deleteBootVolumeRequest.BootVolumeId = &bootVolumeId

			deleteBootVolumeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := blockstorageClient.DeleteBootVolume(context.Background(), deleteBootVolumeRequest)
			if error != nil {
				fmt.Printf("Error deleting BootVolume %s %s, It is possible that the resource is already deleted. Please verify manually \n", bootVolumeId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &bootVolumeId, CoreBootVolumeSweepWaitCondition, time.Duration(3*time.Minute),
				CoreBootVolumeSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreBootVolumeIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BootVolumeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockstorageClient := acctest.GetTestClients(&schema.ResourceData{}).BlockstorageClient()

	listBootVolumesRequest := oci_core.ListBootVolumesRequest{}
	listBootVolumesRequest.CompartmentId = &compartmentId
	availabilityDomains, err := acctest.GetAvalabilityDomains(compartment)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting availabilityDomains required for BootVolume list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, availabilityDomainName := range availabilityDomains {
		listBootVolumesRequest.AvailabilityDomain = &availabilityDomainName

		listBootVolumesResponse, err := blockstorageClient.ListBootVolumes(context.Background(), listBootVolumesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting BootVolume list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, bootVolume := range listBootVolumesResponse.Items {
			id := *bootVolume.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BootVolumeId", id)
		}
	}
	return resourceIds, nil
}

func CoreBootVolumeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bootVolumeResponse, ok := response.Response.(oci_core.GetBootVolumeResponse); ok {
		return bootVolumeResponse.LifecycleState != oci_core.BootVolumeLifecycleStateTerminated
	}
	return false
}

func CoreBootVolumeSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BlockstorageClient().GetBootVolume(context.Background(), oci_core.GetBootVolumeRequest{
		BootVolumeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
