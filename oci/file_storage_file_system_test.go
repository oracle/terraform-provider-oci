// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	FileSystemRequiredOnlyResource = generateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", Required, Create, fileSystemRepresentation)

	fileSystemDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":        Representation{repType: Optional, create: `media-files-1`, update: `displayName2`},
		"id":                  Representation{repType: Optional, create: `${oci_file_storage_file_system.test_file_system.id}`},
		"state":               Representation{repType: Optional, create: `ACTIVE`},
		"filter":              RepresentationGroup{Required, fileSystemDataSourceFilterRepresentation}}
	fileSystemDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_file_storage_file_system.test_file_system.id}`}},
	}

	fileSystemRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        Representation{repType: Optional, create: `media-files-1`, update: `displayName2`},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	FileSystemResourceDependencies = AvailabilityDomainConfig + DefinedTagsDependencies
)

func TestFileStorageFileSystemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageFileSystemResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_file_storage_file_system.test_file_system"
	datasourceName := "data.oci_file_storage_file_systems.test_file_systems"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckFileStorageFileSystemDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + FileSystemResourceDependencies +
					generateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", Required, Create, fileSystemRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + FileSystemResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + FileSystemResourceDependencies +
					generateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", Optional, Create, fileSystemRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "media-files-1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "metered_bytes"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FileSystemResourceDependencies +
					generateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", Optional, Create,
						representationCopyWithNewProperties(fileSystemRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "media-files-1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "metered_bytes"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				Config: config + compartmentIdVariableStr + FileSystemResourceDependencies +
					generateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", Optional, Update, fileSystemRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "metered_bytes"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				Config: config +
					generateDataSourceFromRepresentationMap("oci_file_storage_file_systems", "test_file_systems", Optional, Update, fileSystemDataSourceRepresentation) +
					compartmentIdVariableStr + FileSystemResourceDependencies +
					generateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", Optional, Update, fileSystemRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					TestCheckResourceAttributesEqual(datasourceName, "state", "oci_file_storage_file_system.test_file_system", "state"),

					resource.TestCheckResourceAttr(datasourceName, "file_systems.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "file_systems.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "file_systems.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "file_systems.0.display_name", "displayName2"),
					TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.id", "oci_file_storage_file_system.test_file_system", "id"),
					TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.metered_bytes", "oci_file_storage_file_system.test_file_system", "metered_bytes"),
					TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.state", "oci_file_storage_file_system.test_file_system", "state"),
					TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.time_created", "oci_file_storage_file_system.test_file_system", "time_created"),
				),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckFileStorageFileSystemDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).fileStorageClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_file_storage_file_system" {
			noResourceFound = false
			request := oci_file_storage.GetFileSystemRequest{}

			tmp := rs.Primary.ID
			request.FileSystemId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "file_storage")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("FileStorageFileSystem") {
		resource.AddTestSweepers("FileStorageFileSystem", &resource.Sweeper{
			Name:         "FileStorageFileSystem",
			Dependencies: DependencyGraph["fileSystem"],
			F:            sweepFileStorageFileSystemResource,
		})
	}
}

func sweepFileStorageFileSystemResource(compartment string) error {
	fileStorageClient := GetTestClients(&schema.ResourceData{}).fileStorageClient
	fileSystemIds, err := getFileSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, fileSystemId := range fileSystemIds {
		if ok := SweeperDefaultResourceId[fileSystemId]; !ok {
			deleteFileSystemRequest := oci_file_storage.DeleteFileSystemRequest{}

			deleteFileSystemRequest.FileSystemId = &fileSystemId

			deleteFileSystemRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "file_storage")
			_, error := fileStorageClient.DeleteFileSystem(context.Background(), deleteFileSystemRequest)
			if error != nil {
				fmt.Printf("Error deleting FileSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", fileSystemId, error)
				continue
			}
			waitTillCondition(testAccProvider, &fileSystemId, fileSystemSweepWaitCondition, time.Duration(3*time.Minute),
				fileSystemSweepResponseFetchOperation, "file_storage", true)
		}
	}
	return nil
}

func getFileSystemIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "FileSystemId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fileStorageClient := GetTestClients(&schema.ResourceData{}).fileStorageClient

	listFileSystemsRequest := oci_file_storage.ListFileSystemsRequest{}
	listFileSystemsRequest.CompartmentId = &compartmentId

	availabilityDomains, err := getAvalabilityDomains(compartment)
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
			addResourceIdToSweeperResourceIdMap(compartmentId, "FileSystemId", id)
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

func fileSystemSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.fileStorageClient.GetFileSystem(context.Background(), oci_file_storage.GetFileSystemRequest{
		FileSystemId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
