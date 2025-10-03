// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_lustre_file_storage "github.com/oracle/oci-go-sdk/v65/lustrefilestorage"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LustreFileStorageObjectStorageLinkRequiredOnlyResource = LustreFileStorageObjectStorageLinkResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_object_storage_link", "test_object_storage_link", acctest.Required, acctest.Create, LustreFileStorageObjectStorageLinkRepresentation)

	LustreFileStorageObjectStorageLinkResourceConfig = LustreFileStorageObjectStorageLinkResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_object_storage_link", "test_object_storage_link", acctest.Optional, acctest.Update, LustreFileStorageObjectStorageLinkRepresentation)

	LustreFileStorageObjectStorageLinkSingularDataSourceRepresentation = map[string]interface{}{
		"object_storage_link_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_lustre_file_storage_object_storage_link.test_object_storage_link.id}`},
	}

	LustreFileStorageObjectStorageLinkDataSourceRepresentation = map[string]interface{}{
		"availability_domain":   acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_lustre_file_storage_object_storage_link.test_object_storage_link.id}`},
		"lustre_file_system_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id}`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: LustreFileStorageObjectStorageLinkDataSourceFilterRepresentation}}
	LustreFileStorageObjectStorageLinkDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_lustre_file_storage_object_storage_link.test_object_storage_link.id}`}},
	}

	LustreFileStorageObjectStorageLinkRepresentation = map[string]interface{}{
		"availability_domain":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"file_system_path":      acctest.Representation{RepType: acctest.Required, Create: `/example/file`},
		"lustre_file_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id}`},
		"object_storage_prefix": acctest.Representation{RepType: acctest.Required, Create: `aaa:/aaa/b`},
		"is_overwrite":          acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		// "start_export_to_object_trigger":   acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "start_import_from_object_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "stop_export_to_object_trigger":    acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "stop_import_from_object_trigger":  acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreOslinkSystemTagsChangesRep},
	}
	IgnoreOslinkSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	LustreFileStorageObjectStorageLinkResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_lustre_file_system", "test_lustre_file_system", acctest.Required, acctest.Create, LustreFileStorageLustreFileSystemRepresentation)
)

// issue-routing-tag: lustre_file_storage/default
func TestLustreFileStorageObjectStorageLinkResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLustreFileStorageObjectStorageLinkResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_lustre_file_storage_object_storage_link.test_object_storage_link"
	datasourceName := "data.oci_lustre_file_storage_object_storage_links.test_object_storage_links"
	singularDatasourceName := "data.oci_lustre_file_storage_object_storage_link.test_object_storage_link"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LustreFileStorageObjectStorageLinkResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_object_storage_link", "test_object_storage_link", acctest.Optional, acctest.Create, LustreFileStorageObjectStorageLinkRepresentation), "lustrefilestorage", "objectStorageLink", t)

	acctest.ResourceTest(t, testAccCheckLustreFileStorageObjectStorageLinkDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LustreFileStorageObjectStorageLinkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_object_storage_link", "test_object_storage_link", acctest.Required, acctest.Create, LustreFileStorageObjectStorageLinkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "file_system_path", "/example/file"),
				resource.TestCheckResourceAttrSet(resourceName, "lustre_file_system_id"),
				resource.TestCheckResourceAttr(resourceName, "object_storage_prefix", "aaa:/aaa/b"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LustreFileStorageObjectStorageLinkResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LustreFileStorageObjectStorageLinkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_object_storage_link", "test_object_storage_link", acctest.Optional, acctest.Create, LustreFileStorageObjectStorageLinkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "file_system_path", "/example/file"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_overwrite", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "lustre_file_system_id"),
				resource.TestCheckResourceAttr(resourceName, "object_storage_prefix", "aaa:/aaa/b"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LustreFileStorageObjectStorageLinkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_object_storage_link", "test_object_storage_link", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LustreFileStorageObjectStorageLinkRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "file_system_path", "/example/file"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_overwrite", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "lustre_file_system_id"),
				resource.TestCheckResourceAttr(resourceName, "object_storage_prefix", "aaa:/aaa/b"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
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
			Config: config + compartmentIdVariableStr + LustreFileStorageObjectStorageLinkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_object_storage_link", "test_object_storage_link", acctest.Optional, acctest.Update, LustreFileStorageObjectStorageLinkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "file_system_path", "/example/file"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_overwrite", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "lustre_file_system_id"),
				resource.TestCheckResourceAttr(resourceName, "object_storage_prefix", "aaa:/aaa/b"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_lustre_file_storage_object_storage_links", "test_object_storage_links", acctest.Optional, acctest.Update, LustreFileStorageObjectStorageLinkDataSourceRepresentation) +
				compartmentIdVariableStr + LustreFileStorageObjectStorageLinkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_object_storage_link", "test_object_storage_link", acctest.Optional, acctest.Update, LustreFileStorageObjectStorageLinkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "lustre_file_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "object_storage_link_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "object_storage_link_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_lustre_file_storage_object_storage_link", "test_object_storage_link", acctest.Required, acctest.Create, LustreFileStorageObjectStorageLinkSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LustreFileStorageObjectStorageLinkResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_storage_link_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "file_system_path", "/example/file"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_storage_prefix", "aaa:/aaa/b"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + LustreFileStorageObjectStorageLinkRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLustreFileStorageObjectStorageLinkDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LustreFileStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_lustre_file_storage_object_storage_link" {
			noResourceFound = false
			request := oci_lustre_file_storage.GetObjectStorageLinkRequest{}

			tmp := rs.Primary.ID
			request.ObjectStorageLinkId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "lustre_file_storage")

			response, err := client.GetObjectStorageLink(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_lustre_file_storage.ObjectStorageLinkLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("LustreFileStorageObjectStorageLink") {
		resource.AddTestSweepers("LustreFileStorageObjectStorageLink", &resource.Sweeper{
			Name:         "LustreFileStorageObjectStorageLink",
			Dependencies: acctest.DependencyGraph["objectStorageLink"],
			F:            sweepLustreFileStorageObjectStorageLinkResource,
		})
	}
}

func sweepLustreFileStorageObjectStorageLinkResource(compartment string) error {
	lustreFileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).LustreFileStorageClient()
	objectStorageLinkIds, err := getLustreFileStorageObjectStorageLinkIds(compartment)
	if err != nil {
		return err
	}
	for _, objectStorageLinkId := range objectStorageLinkIds {
		if ok := acctest.SweeperDefaultResourceId[objectStorageLinkId]; !ok {
			deleteObjectStorageLinkRequest := oci_lustre_file_storage.DeleteObjectStorageLinkRequest{}

			deleteObjectStorageLinkRequest.ObjectStorageLinkId = &objectStorageLinkId

			deleteObjectStorageLinkRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "lustre_file_storage")
			_, error := lustreFileStorageClient.DeleteObjectStorageLink(context.Background(), deleteObjectStorageLinkRequest)
			if error != nil {
				fmt.Printf("Error deleting ObjectStorageLink %s %s, It is possible that the resource is already deleted. Please verify manually \n", objectStorageLinkId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &objectStorageLinkId, LustreFileStorageObjectStorageLinkSweepWaitCondition, time.Duration(3*time.Minute),
				LustreFileStorageObjectStorageLinkSweepResponseFetchOperation, "lustre_file_storage", true)
		}
	}
	return nil
}

func getLustreFileStorageObjectStorageLinkIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ObjectStorageLinkId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	lustreFileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).LustreFileStorageClient()

	listObjectStorageLinksRequest := oci_lustre_file_storage.ListObjectStorageLinksRequest{}
	listObjectStorageLinksRequest.CompartmentId = &compartmentId
	listObjectStorageLinksRequest.LifecycleState = oci_lustre_file_storage.ObjectStorageLinkLifecycleStateActive
	listObjectStorageLinksResponse, err := lustreFileStorageClient.ListObjectStorageLinks(context.Background(), listObjectStorageLinksRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ObjectStorageLink list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, objectStorageLink := range listObjectStorageLinksResponse.Items {
		id := *objectStorageLink.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ObjectStorageLinkId", id)
	}
	return resourceIds, nil
}

func LustreFileStorageObjectStorageLinkSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if objectStorageLinkResponse, ok := response.Response.(oci_lustre_file_storage.GetObjectStorageLinkResponse); ok {
		return objectStorageLinkResponse.LifecycleState != oci_lustre_file_storage.ObjectStorageLinkLifecycleStateDeleted
	}
	return false
}

func LustreFileStorageObjectStorageLinkSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.LustreFileStorageClient().GetObjectStorageLink(context.Background(), oci_lustre_file_storage.GetObjectStorageLinkRequest{
		ObjectStorageLinkId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
