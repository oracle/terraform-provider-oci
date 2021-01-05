// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v31/common"
	oci_common "github.com/oracle/oci-go-sdk/v31/common"
	oci_core "github.com/oracle/oci-go-sdk/v31/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BootVolumeWaitConditionDuration = time.Duration(20 * time.Minute)

	BootVolumeRequiredOnlyResource = BootVolumeResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Required, Create, bootVolumeRepresentation)

	BootVolumeOptionalResource = BootVolumeResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Optional, Create, bootVolumeRepresentation)

	BootVolumeResourceConfig = BootVolumeResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Optional, Update, bootVolumeRepresentation)

	bootVolumeSingularDataSourceRepresentation = map[string]interface{}{
		"boot_volume_id": Representation{repType: Required, create: `${oci_core_boot_volume.test_boot_volume.id}`},
	}

	bootVolumeDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"filter":              RepresentationGroup{Required, bootVolumeDataSourceFilterRepresentation}}
	bootVolumeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_boot_volume.test_boot_volume.id}`}},
	}

	bootVolumeRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"source_details":      RepresentationGroup{Required, bootVolumeSourceDetailsRepresentation},
		"backup_policy_id":    Representation{repType: Optional, create: `${data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id}`},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"kms_key_id":          Representation{repType: Optional, create: `${var.kms_key_id_for_create}`, update: `${var.kms_key_id_for_update}`},
		"size_in_gbs":         Representation{repType: Optional, create: `50`, update: `51`},
		"vpus_per_gb":         Representation{repType: Optional, create: `10`, update: `20`},
	}
	bootVolumeSourceDetailsRepresentation = map[string]interface{}{
		"id":   Representation{repType: Required, create: `${oci_core_instance.test_instance.boot_volume_id}`},
		"type": Representation{repType: Required, create: `bootVolume`},
	}

	BootVolumeResourceDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		VolumeBackupPolicyDependency +
		generateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Required, Create, volumeGroupRepresentation) +
		SourceVolumeListDependency +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig + kmsKeyIdCreateVariableStr + kmsKeyIdUpdateVariableStr
)

func TestCoreBootVolumeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreBootVolumeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_boot_volume.test_boot_volume"
	datasourceName := "data.oci_core_boot_volumes.test_boot_volumes"
	singularDatasourceName := "data.oci_core_boot_volume.test_boot_volume"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreBootVolumeDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + BootVolumeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Required, Create, bootVolumeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckNoResourceAttr(resourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "bootVolume"),
					resource.TestCheckNoResourceAttr(resourceName, "volume_group_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + BootVolumeResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + BootVolumeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Optional, Create, bootVolumeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + BootVolumeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Optional, Create,
						representationCopyWithNewProperties(bootVolumeRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				PreConfig: waitTillCondition(testAccProvider, &resId, bootVolumeWaitCondition, BootVolumeWaitConditionDuration,
					bootVolumeResponseFetchOperation, "core", false),
				Config: config + compartmentIdVariableStr + BootVolumeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Optional, Update, bootVolumeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resource.TestCheckResourceAttr(resourceName, "vpus_per_gb", "20"),
					resource.TestCheckNoResourceAttr(resourceName, "volume_group_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				PreConfig: waitTillCondition(testAccProvider, &resId, bootVolumeWaitCondition, BootVolumeWaitConditionDuration,
					bootVolumeResponseFetchOperation, "core", false),
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_boot_volumes", "test_boot_volumes", Optional, Update, bootVolumeDataSourceRepresentation) +
					compartmentIdVariableStr + BootVolumeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Optional, Update, bootVolumeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckNoResourceAttr(datasourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "boot_volumes.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.defined_tags.%", "1"),
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
					resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.vpus_per_gb", "20"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Required, Create, bootVolumeSingularDataSourceRepresentation) +
					compartmentIdVariableStr + BootVolumeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr(singularDatasourceName, "backup_policy_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "kms_key_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
					resource.TestCheckResourceAttr(singularDatasourceName, "vpus_per_gb", "20"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + BootVolumeResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"backup_policy_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckCoreBootVolumeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).blockstorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_boot_volume" {
			noResourceFound = false
			request := oci_core.GetBootVolumeRequest{}

			tmp := rs.Primary.ID
			request.BootVolumeId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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

func bootVolumeResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *oci_common.RetryPolicy) error {
	_, err := client.blockstorageClient().GetBootVolume(context.Background(), oci_core.GetBootVolumeRequest{
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreBootVolume") {
		resource.AddTestSweepers("CoreBootVolume", &resource.Sweeper{
			Name:         "CoreBootVolume",
			Dependencies: DependencyGraph["bootVolume"],
			F:            sweepCoreBootVolumeResource,
		})
	}
}

func sweepCoreBootVolumeResource(compartment string) error {
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient()
	bootVolumeIds, err := getBootVolumeIds(compartment)
	if err != nil {
		return err
	}
	for _, bootVolumeId := range bootVolumeIds {
		if ok := SweeperDefaultResourceId[bootVolumeId]; !ok {
			deleteBootVolumeRequest := oci_core.DeleteBootVolumeRequest{}

			deleteBootVolumeRequest.BootVolumeId = &bootVolumeId

			deleteBootVolumeRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := blockstorageClient.DeleteBootVolume(context.Background(), deleteBootVolumeRequest)
			if error != nil {
				fmt.Printf("Error deleting BootVolume %s %s, It is possible that the resource is already deleted. Please verify manually \n", bootVolumeId, error)
				continue
			}
			waitTillCondition(testAccProvider, &bootVolumeId, bootVolumeSweepWaitCondition, time.Duration(3*time.Minute),
				bootVolumeSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getBootVolumeIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "BootVolumeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient()

	listBootVolumesRequest := oci_core.ListBootVolumesRequest{}
	listBootVolumesRequest.CompartmentId = &compartmentId

	availabilityDomains, err := getAvalabilityDomains(compartment)
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
			addResourceIdToSweeperResourceIdMap(compartmentId, "BootVolumeId", id)
		}

	}
	return resourceIds, nil
}

func bootVolumeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bootVolumeResponse, ok := response.Response.(oci_core.GetBootVolumeResponse); ok {
		return bootVolumeResponse.LifecycleState != oci_core.BootVolumeLifecycleStateTerminated
	}
	return false
}

func bootVolumeSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.blockstorageClient().GetBootVolume(context.Background(), oci_core.GetBootVolumeRequest{
		BootVolumeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
