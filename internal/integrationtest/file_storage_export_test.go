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
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/v56/filestorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ExportRequiredOnlyResource = ExportResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_export", "test_export", acctest.Required, acctest.Create, exportRepresentation)

	exportDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"export_set_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_export_set.test_export_set.id}`},
		"file_system_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_file_system.test_file_system.id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_export.test_export.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: exportDataSourceFilterRepresentation}}
	exportDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_file_storage_export.test_export.id}`}},
	}

	exportRepresentation = map[string]interface{}{
		"export_set_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_export_set.test_export_set.id}`},
		"file_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_file_system.test_file_system.id}`},
		"path":           acctest.Representation{RepType: acctest.Required, Create: `/files-5`},
		"export_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exportExportOptionsRepresentation},
	}
	exportExportOptionsRepresentation = map[string]interface{}{
		"source":                         acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"access":                         acctest.Representation{RepType: acctest.Optional, Create: `READ_WRITE`, Update: `READ_ONLY`},
		"anonymous_gid":                  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"anonymous_uid":                  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"identity_squash":                acctest.Representation{RepType: acctest.Optional, Create: `NONE`, Update: `ALL`},
		"require_privileged_source_port": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	ExportResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_file_storage_export_set", "test_export_set", acctest.Required, acctest.Create, exportSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", acctest.Required, acctest.Create, mountTargetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Required, acctest.Create, fileSystemRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: file_storage/default
func TestFileStorageExportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageExportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_export.test_export"
	datasourceName := "data.oci_file_storage_exports.test_exports"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ExportResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_export", "test_export", acctest.Optional, acctest.Create, exportRepresentation), "filestorage", "export", t)

	acctest.ResourceTest(t, testAccCheckFileStorageExportDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ExportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_export", "test_export", acctest.Required, acctest.Create, exportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
				resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
				resource.TestCheckResourceAttr(resourceName, "path", "/files-5"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ExportResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ExportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_export", "test_export", acctest.Optional, acctest.Create, exportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "export_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "export_options.0.access", "READ_WRITE"),
				resource.TestCheckResourceAttr(resourceName, "export_options.0.anonymous_gid", "10"),
				resource.TestCheckResourceAttr(resourceName, "export_options.0.anonymous_uid", "10"),
				resource.TestCheckResourceAttr(resourceName, "export_options.0.identity_squash", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "export_options.0.require_privileged_source_port", "false"),
				resource.TestCheckResourceAttr(resourceName, "export_options.0.source", "0.0.0.0/0"),
				resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
				resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "path", "/files-5"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ExportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_export", "test_export", acctest.Optional, acctest.Update, exportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "export_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "export_options.0.access", "READ_ONLY"),
				resource.TestCheckResourceAttr(resourceName, "export_options.0.anonymous_gid", "11"),
				resource.TestCheckResourceAttr(resourceName, "export_options.0.anonymous_uid", "11"),
				resource.TestCheckResourceAttr(resourceName, "export_options.0.identity_squash", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "export_options.0.require_privileged_source_port", "true"),
				resource.TestCheckResourceAttr(resourceName, "export_options.0.source", "0.0.0.0/0"),
				resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
				resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "path", "/files-5"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_exports", "test_exports", acctest.Optional, acctest.Update, exportDataSourceRepresentation) +
				compartmentIdVariableStr + ExportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_export", "test_export", acctest.Optional, acctest.Update, exportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "exports.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "exports.0.export_set_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "exports.0.file_system_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "exports.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "exports.0.path", "/files-5"),
				resource.TestCheckResourceAttr(datasourceName, "exports.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "exports.0.time_created"),
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
	})
}

func testAccCheckFileStorageExportDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FileStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_file_storage_export" {
			noResourceFound = false
			request := oci_file_storage.GetExportRequest{}

			tmp := rs.Primary.ID
			request.ExportId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")

			response, err := client.GetExport(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_file_storage.ExportLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FileStorageExport") {
		resource.AddTestSweepers("FileStorageExport", &resource.Sweeper{
			Name:         "FileStorageExport",
			Dependencies: acctest.DependencyGraph["export"],
			F:            sweepFileStorageExportResource,
		})
	}
}

func sweepFileStorageExportResource(compartment string) error {
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()
	exportIds, err := getExportIds(compartment)
	if err != nil {
		return err
	}
	for _, exportId := range exportIds {
		if ok := acctest.SweeperDefaultResourceId[exportId]; !ok {
			deleteExportRequest := oci_file_storage.DeleteExportRequest{}

			deleteExportRequest.ExportId = &exportId

			deleteExportRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")
			_, error := fileStorageClient.DeleteExport(context.Background(), deleteExportRequest)
			if error != nil {
				fmt.Printf("Error deleting Export %s %s, It is possible that the resource is already deleted. Please verify manually \n", exportId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &exportId, exportSweepWaitCondition, time.Duration(3*time.Minute),
				exportSweepResponseFetchOperation, "file_storage", true)
		}
	}
	return nil
}

func getExportIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExportId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()

	listExportsRequest := oci_file_storage.ListExportsRequest{}
	listExportsRequest.CompartmentId = &compartmentId
	listExportsRequest.LifecycleState = oci_file_storage.ListExportsLifecycleStateActive
	listExportsResponse, err := fileStorageClient.ListExports(context.Background(), listExportsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Export list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, export := range listExportsResponse.Items {
		id := *export.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExportId", id)
	}
	return resourceIds, nil
}

func exportSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if exportResponse, ok := response.Response.(oci_file_storage.GetExportResponse); ok {
		return exportResponse.LifecycleState != oci_file_storage.ExportLifecycleStateDeleted
	}
	return false
}

func exportSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FileStorageClient().GetExport(context.Background(), oci_file_storage.GetExportRequest{
		ExportId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
