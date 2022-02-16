// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/v58/filestorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	FileSystemRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Required, acctest.Create, fileSystemRepresentation)

	fileSystemDataSourceRepresentation = map[string]interface{}{
		"availability_domain":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `media-files-1`, Update: `displayName2`},
		"id":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_file_system.test_file_system2.id}`},
		"parent_file_system_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_file_system.test_file_system.id}`},
		"source_snapshot_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_snapshot.test_snapshot.id}`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: fileSystemDataSourceFilterRepresentation}}
	fileSystemDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_file_storage_file_system.test_file_system2.id}`}},
	}

	fileSystemRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `media-files-1`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"source_snapshot_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_snapshot.test_snapshot.id}`},
		"kms_key_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id_for_create}`, Update: `${var.kms_key_id_for_update}`},
	}

	FileSystemResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Required, acctest.Create, fileSystemRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", acctest.Required, acctest.Create, snapshotRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig + kmsKeyIdCreateVariableStr + kmsKeyIdUpdateVariableStr
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FileSystemResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Optional, acctest.Create, fileSystemRepresentation), "filestorage", "fileSystem", t)

	acctest.ResourceTest(t, testAccCheckFileStorageFileSystemDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system2", acctest.Required, acctest.Create, fileSystemRepresentation),
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
			Config: config + compartmentIdVariableStr + FileSystemResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system2", acctest.Optional, acctest.Create, fileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "media-files-1"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system2", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(fileSystemRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "media-files-1"),
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
		{
			Config: config + compartmentIdVariableStr + FileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system2", acctest.Optional, acctest.Update, fileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_file_systems", "test_file_systems", acctest.Optional, acctest.Update, fileSystemDataSourceRepresentation) +
				compartmentIdVariableStr + FileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system2", acctest.Optional, acctest.Update, fileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "parent_file_system_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "source_snapshot_id"),
				acctest.TestCheckResourceAttributesEqual(datasourceName, "state", "oci_file_storage_file_system.test_file_system2", "state"),

				resource.TestCheckResourceAttr(datasourceName, "file_systems.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "file_systems.0.availability_domain"),
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
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"source_snapshot_id", "parent_file_system_id"},
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
	fileSystemIds, err := getFileSystemIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &fileSystemId, fileSystemSweepWaitCondition, time.Duration(3*time.Minute),
				fileSystemSweepResponseFetchOperation, "file_storage", true)
		}
	}
	return nil
}

func getFileSystemIds(compartment string) ([]string, error) {
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

func fileSystemSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fileSystemResponse, ok := response.Response.(oci_file_storage.GetFileSystemResponse); ok {
		return fileSystemResponse.LifecycleState != oci_file_storage.FileSystemLifecycleStateDeleted
	}
	return false
}

func fileSystemSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FileStorageClient().GetFileSystem(context.Background(), oci_file_storage.GetFileSystemRequest{
		FileSystemId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
