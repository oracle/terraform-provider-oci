// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	FileStorageFileSystemRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Required, acctest.Create, FileStorageFileSystemRepresentation)

	FileStorageFileSystem2RequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system2", acctest.Required, acctest.Create, FileStorageFileSystemRepresentation)

	FileStorageFileStorageFileSystemDataSourceRepresentation = map[string]interface{}{
		"availability_domain":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `media-files-1`, Update: `displayName2`},
		"filesystem_snapshot_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_filesystem_snapshot_policy.test_filesystem_snapshot_policy.id}`},
		"id":                            acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_file_system.test_file_system2.id}`},
		"parent_file_system_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_file_system.test_file_system.id}`},
		"source_snapshot_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_snapshot.test_snapshot.id}`},
		"state":                         acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: FileStorageFileSystemDataSourceFilterRepresentation}}
	FileStorageFileSystemDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_file_storage_file_system.test_file_system2.id}`}},
	}

	FileStorageFileSystemRepresentation = map[string]interface{}{
		"availability_domain":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"clone_attach_status":           acctest.Representation{RepType: acctest.Optional, Create: `DETACH`},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `media-files-1`, Update: `displayName2`},
		"filesystem_snapshot_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_filesystem_snapshot_policy.test_filesystem_snapshot_policy.id}`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"kms_key_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key.kms_key_id_for_create.id}`, Update: `${oci_kms_key.kms_key_id_for_update.id}`},
		"source_snapshot_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_snapshot.test_snapshot.id}`},
		"detach_clone_trigger":          acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `0`},
		"lifecycle":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	FileStorageFileSystemResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Required, acctest.Create, FileStorageFileSystemRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Required, acctest.Create, FileStorageFilesystemSnapshotPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", acctest.Required, acctest.Create, FileStorageSnapshotRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_kms_vault", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(KmsVaultRepresentation, map[string]interface{}{
			"vault_type": acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "kms_key_id_for_create", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(KmsKeyRepresentation, map[string]interface{}{
			"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_vault.test_kms_vault.management_endpoint}`},
			"desired_state":       acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "kms_key_id_for_update", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(KmsKeyRepresentation, map[string]interface{}{
			"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_vault.test_kms_vault.management_endpoint}`},
			"desired_state":       acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		}))
)

// issue-routing-tag: file_storage/default
func TestFileStorageFileSystemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageFileSystemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_file_storage_file_system.test_file_system2"
	datasourceName := "data.oci_file_storage_file_systems.test_file_systems"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FileStorageFileSystemResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Optional, acctest.Create, FileStorageFileSystemRepresentation), "filestorage", "fileSystem", t)

	acctest.ResourceTest(t, testAccCheckFileStorageFileSystemDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FileStorageFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system2", acctest.Required, acctest.Create, FileStorageFileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FileStorageFileSystemResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FileStorageFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system2", acctest.Optional, acctest.Create, FileStorageFileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "clone_attach_status", "ATTACHED"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "media-files-1"),
				resource.TestCheckResourceAttrSet(resourceName, "filesystem_snapshot_policy_id"),
				resource.TestMatchResourceAttr(resourceName, "filesystem_snapshot_policy_id", regexp.MustCompile("ocid*")),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "metered_bytes"),
				resource.TestCheckResourceAttrSet(resourceName, "source_snapshot_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FileStorageFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system2", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FileStorageFileSystemRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "clone_attach_status", "DETACHING"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "media-files-1"),
				resource.TestCheckResourceAttrSet(resourceName, "filesystem_snapshot_policy_id"),
				resource.TestMatchResourceAttr(resourceName, "filesystem_snapshot_policy_id", regexp.MustCompile("ocid*")),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "metered_bytes"),
				resource.TestCheckResourceAttrSet(resourceName, "source_snapshot_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// Detach policy from FS (policy will get attached back to FS in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FileStorageFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system2", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FileStorageFileSystemRepresentation, map[string]interface{}{
						"filesystem_snapshot_policy_id": acctest.Representation{RepType: acctest.Optional, Create: ""},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "media-files-1"),
				resource.TestCheckResourceAttr(resourceName, "filesystem_snapshot_policy_id", ""),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "metered_bytes"),
				resource.TestCheckResourceAttrSet(resourceName, "source_snapshot_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
		// Includes attach policy to FS
		{
			Config: config + compartmentIdVariableStr + FileStorageFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system2", acctest.Optional, acctest.Update, FileStorageFileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "clone_attach_status", "DETACHING"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "filesystem_snapshot_policy_id"),
				resource.TestMatchResourceAttr(resourceName, "filesystem_snapshot_policy_id", regexp.MustCompile("ocid*")),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "metered_bytes"),
				resource.TestCheckResourceAttrSet(resourceName, "source_snapshot_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_file_systems", "test_file_systems", acctest.Optional, acctest.Update, FileStorageFileStorageFileSystemDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system2", acctest.Optional, acctest.Update, FileStorageFileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "filesystem_snapshot_policy_id"),
				resource.TestMatchResourceAttr(datasourceName, "filesystem_snapshot_policy_id", regexp.MustCompile("ocid*")),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "parent_file_system_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "source_snapshot_id"),
				acctest.TestCheckResourceAttributesEqual(datasourceName, "state", "oci_file_storage_file_system.test_file_system2", "state"),

				resource.TestCheckResourceAttr(datasourceName, "file_systems.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "file_systems.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "file_systems.0.clone_attach_status", "DETACHING"),
				resource.TestCheckResourceAttr(datasourceName, "file_systems.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "file_systems.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "file_systems.0.freeform_tags.%", "1"),

				acctest.TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.id", "oci_file_storage_file_system.test_file_system2", "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "file_systems.0.is_clone_parent"),
				resource.TestCheckResourceAttrSet(datasourceName, "file_systems.0.is_hydrated"),
				acctest.TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.kms_key_id", "oci_file_storage_file_system.test_file_system2", "kms_key_id"),
				acctest.TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.metered_bytes", "oci_file_storage_file_system.test_file_system2", "metered_bytes"),
				resource.TestCheckResourceAttr(datasourceName, "file_systems.0.source_details.#", "1"),
				acctest.TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.state", "oci_file_storage_file_system.test_file_system2", "state"),
				acctest.TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.time_created", "oci_file_storage_file_system.test_file_system2", "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + FileStorageFileSystem2RequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"source_snapshot_id", "parent_file_system_id", "detach_clone_trigger"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFileStorageFileSystemDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FileStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_file_storage_file_system" {
			noResourceFound = false
			request := oci_file_storage.GetFileSystemRequest{}

			tmp := rs.Primary.ID
			request.FileSystemId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")

			response, err := client.GetFileSystem(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_file_storage.FileSystemLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FileStorageFileSystem") {
		resource.AddTestSweepers("FileStorageFileSystem", &resource.Sweeper{
			Name:         "FileStorageFileSystem",
			Dependencies: acctest.DependencyGraph["fileSystem"],
			F:            sweepFileStorageFileSystemResource,
		})
	}
}

func sweepFileStorageFileSystemResource(compartment string) error {
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()
	fileSystemIds, err := getFileStorageFileSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, fileSystemId := range fileSystemIds {
		if ok := acctest.SweeperDefaultResourceId[fileSystemId]; !ok {
			deleteFileSystemRequest := oci_file_storage.DeleteFileSystemRequest{}

			deleteFileSystemRequest.FileSystemId = &fileSystemId

			deleteFileSystemRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")
			_, error := fileStorageClient.DeleteFileSystem(context.Background(), deleteFileSystemRequest)
			if error != nil {
				fmt.Printf("Error deleting FileSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", fileSystemId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &fileSystemId, FileStorageFileSystemSweepWaitCondition, time.Duration(3*time.Minute),
				FileStorageFileSystemSweepResponseFetchOperation, "file_storage", true)
		}
	}
	return nil
}

func getFileStorageFileSystemIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FileSystemId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()

	listFileSystemsRequest := oci_file_storage.ListFileSystemsRequest{}
	listFileSystemsRequest.CompartmentId = &compartmentId

	availabilityDomains, err := acctest.GetAvalabilityDomains(compartment)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting availabilityDomains required for FileSystem list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, availabilityDomainName := range availabilityDomains {
		listFileSystemsRequest.AvailabilityDomain = &availabilityDomainName

		listFileSystemsRequest.LifecycleState = oci_file_storage.ListFileSystemsLifecycleStateActive
		listFileSystemsResponse, err := fileStorageClient.ListFileSystems(context.Background(), listFileSystemsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting FileSystem list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, fileSystem := range listFileSystemsResponse.Items {
			id := *fileSystem.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FileSystemId", id)
		}

	}
	return resourceIds, nil
}

func FileStorageFileSystemSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fileSystemResponse, ok := response.Response.(oci_file_storage.GetFileSystemResponse); ok {
		return fileSystemResponse.LifecycleState != oci_file_storage.FileSystemLifecycleStateDeleted
	}
	return false
}

func FileStorageFileSystemSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FileStorageClient().GetFileSystem(context.Background(), oci_file_storage.GetFileSystemRequest{
		FileSystemId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
