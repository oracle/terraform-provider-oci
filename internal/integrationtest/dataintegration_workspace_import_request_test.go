// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataintegrationWorkspaceImportRequestRequiredOnlyResource = DataintegrationWorkspaceImportRequestResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_import_request", "test_workspace_import_request", acctest.Required, acctest.Create, DataintegrationWorkspaceImportRequestRepresentation)

	DataintegrationWorkspaceImportRequestResourceConfig = DataintegrationWorkspaceImportRequestResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_import_request", "test_workspace_import_request", acctest.Optional, acctest.Update, DataintegrationWorkspaceImportRequestRepresentation)

	DataintegrationWorkspaceImportRequestSingularDataSourceRepresentation = map[string]interface{}{
		"import_request_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_import_request.test_workspace_import_request.key}`},
		"workspace_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
	}

	DataintegrationWorkspaceImportRequestDataSourceRepresentation = map[string]interface{}{
		"workspace_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
		"import_status": acctest.Representation{RepType: acctest.Optional, Create: `SUCCESSFUL`},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_dataintegration_workspace_import_request.test_workspace_import_request.name}`},
		"projection":    acctest.Representation{RepType: acctest.Optional, Create: `SUMMARY`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceImportRequestDataSourceFilterRepresentation}}
	DataintegrationWorkspaceImportRequestDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataintegration_workspace_import_request.test_workspace_import_request.name}`}},
	}

	DataintegrationWorkspaceImportRequestRepresentation = map[string]interface{}{
		"bucket":                             acctest.Representation{RepType: acctest.Required, Create: `${var.bucket_name}`},
		"file_name":                          acctest.Representation{RepType: acctest.Required, Create: `MyExportObjects.zip`},
		"workspace_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
		"are_data_asset_references_included": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"import_conflict_resolution":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceImportRequestImportConflictResolutionRepresentation},
		"object_key_for_import":              acctest.Representation{RepType: acctest.Optional, Create: nil},
		"object_storage_region":              acctest.Representation{RepType: acctest.Optional, Create: `${var.region_name}`},
		"object_storage_tenancy_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
	}
	DataintegrationWorkspaceImportRequestImportConflictResolutionRepresentation = map[string]interface{}{
		"import_conflict_resolution_type": acctest.Representation{RepType: acctest.Required, Create: `REPLACE`},
		"duplicate_prefix":                acctest.Representation{RepType: acctest.Optional, Create: `duplicatePrefix`},
		"duplicate_suffix":                acctest.Representation{RepType: acctest.Optional, Create: `duplicateSuffix`},
	}

	DataintegrationWorkspaceImportRequestResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_export_request", "test_workspace_export_request", acctest.Required, acctest.Create, acctest.GetMultipleUpdatedRepresenationCopy([]string{"bucket", "workspace_id", "file_name", "is_object_overwrite_enabled", "object_storage_region", "object_storage_tenancy_id"},
		[]interface{}{
			acctest.Representation{RepType: acctest.Required, Create: `${var.bucket_name}`},
			acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
			acctest.Representation{RepType: acctest.Required, Create: `MyExportObjects.zip`},
			acctest.Representation{RepType: acctest.Required, Create: `true`},
			acctest.Representation{RepType: acctest.Required, Create: `${var.region_name}`},
			acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
		}, DataintegrationWorkspaceExportRequestRepresentation))
)

// issue-routing-tag: dataintegration/default
func TestDataintegrationWorkspaceImportRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataintegrationWorkspaceImportRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	workspaceId := utils.GetEnvSettingWithBlankDefault("workspace_id")
	workspaceIdVariableStr := fmt.Sprintf("variable \"workspace_id\" { default = \"%s\" }\n", workspaceId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)

	region := utils.GetEnvSettingWithBlankDefault("region")
	regionVariableStr := fmt.Sprintf("variable \"region_name\" { default = \"%s\" }\n", region)

	bucket := utils.GetEnvSettingWithBlankDefault("bucket_name")
	bucketVariableStr := fmt.Sprintf("variable \"bucket_name\" { default = \"%s\" }\n", bucket)

	resourceName := "oci_dataintegration_workspace_import_request.test_workspace_import_request"
	datasourceName := "data.oci_dataintegration_workspace_import_requests.test_workspace_import_requests"
	singularDatasourceName := "data.oci_dataintegration_workspace_import_request.test_workspace_import_request"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+workspaceIdVariableStr+tenancyIdVariableStr+regionVariableStr+bucketVariableStr+DataintegrationWorkspaceImportRequestResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_import_request", "test_workspace_import_request", acctest.Optional, acctest.Create, DataintegrationWorkspaceImportRequestRepresentation), "dataintegration", "workspaceImportRequest", t)

	acctest.ResourceTest(t, testAccCheckDataintegrationWorkspaceImportRequestDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + workspaceIdVariableStr + tenancyIdVariableStr + regionVariableStr + bucketVariableStr + DataintegrationWorkspaceImportRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_import_request", "test_workspace_import_request", acctest.Required, acctest.Create, DataintegrationWorkspaceImportRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bucket"),
				resource.TestCheckResourceAttr(resourceName, "file_name", "MyExportObjects.zip"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + workspaceIdVariableStr + tenancyIdVariableStr + regionVariableStr + bucketVariableStr + DataintegrationWorkspaceImportRequestResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + workspaceIdVariableStr + tenancyIdVariableStr + regionVariableStr + bucketVariableStr + DataintegrationWorkspaceImportRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_import_request", "test_workspace_import_request", acctest.Optional, acctest.Create, DataintegrationWorkspaceImportRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_data_asset_references_included", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket"),
				resource.TestCheckResourceAttr(resourceName, "file_name", "MyExportObjects.zip"),
				resource.TestCheckResourceAttr(resourceName, "import_conflict_resolution.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "import_conflict_resolution.0.duplicate_prefix", "duplicatePrefix"),
				resource.TestCheckResourceAttr(resourceName, "import_conflict_resolution.0.duplicate_suffix", "duplicateSuffix"),
				resource.TestCheckResourceAttr(resourceName, "import_conflict_resolution.0.import_conflict_resolution_type", "REPLACE"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_region"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_tenancy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_import_requests", "test_workspace_import_requests", acctest.Optional, acctest.Update, DataintegrationWorkspaceImportRequestDataSourceRepresentation) +
				compartmentIdVariableStr + workspaceIdVariableStr + tenancyIdVariableStr + regionVariableStr + bucketVariableStr + DataintegrationWorkspaceImportRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_import_request", "test_workspace_import_request", acctest.Optional, acctest.Update, DataintegrationWorkspaceImportRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "import_status", "SUCCESSFUL"),
				resource.TestCheckResourceAttrSet(datasourceName, "name"),
				resource.TestCheckResourceAttr(datasourceName, "projection", "SUMMARY"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(datasourceName, "import_request_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "import_request_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_import_request", "test_workspace_import_request", acctest.Required, acctest.Create, DataintegrationWorkspaceImportRequestSingularDataSourceRepresentation) +
				compartmentIdVariableStr + workspaceIdVariableStr + tenancyIdVariableStr + regionVariableStr + bucketVariableStr + DataintegrationWorkspaceImportRequestResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "import_request_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "are_data_asset_references_included", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bucket"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "file_name", "MyExportObjects.zip"),
				resource.TestCheckResourceAttr(singularDatasourceName, "import_conflict_resolution.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "import_conflict_resolution.0.duplicate_prefix", "duplicatePrefix"),
				resource.TestCheckResourceAttr(singularDatasourceName, "import_conflict_resolution.0.duplicate_suffix", "duplicateSuffix"),
				resource.TestCheckResourceAttr(singularDatasourceName, "import_conflict_resolution.0.import_conflict_resolution_type", "REPLACE"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "imported_objects.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_storage_region"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended_in_millis"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started_in_millis"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_imported_object_count"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataintegrationWorkspaceImportRequestRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataintegrationWorkspaceImportRequestDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataIntegrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataintegration_workspace_import_request" {
			noResourceFound = false
			request := oci_dataintegration.GetImportRequestRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.ImportRequestKey = &value
			}

			if value, ok := rs.Primary.Attributes["workspace_id"]; ok {
				request.WorkspaceId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")

			_, err := client.GetImportRequest(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("DataintegrationWorkspaceImportRequest") {
		resource.AddTestSweepers("DataintegrationWorkspaceImportRequest", &resource.Sweeper{
			Name:         "DataintegrationWorkspaceImportRequest",
			Dependencies: acctest.DependencyGraph["workspaceImportRequest"],
			F:            sweepDataintegrationWorkspaceImportRequestResource,
		})
	}
}

func sweepDataintegrationWorkspaceImportRequestResource(compartment string) error {
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()
	workspaceImportRequestIds, err := getDataintegrationWorkspaceImportRequestIds(compartment)
	if err != nil {
		return err
	}
	for _, workspaceImportRequestId := range workspaceImportRequestIds {
		if ok := acctest.SweeperDefaultResourceId[workspaceImportRequestId]; !ok {
			deleteImportRequestRequest := oci_dataintegration.DeleteImportRequestRequest{}

			deleteImportRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")
			_, error := dataIntegrationClient.DeleteImportRequest(context.Background(), deleteImportRequestRequest)
			if error != nil {
				fmt.Printf("Error deleting WorkspaceImportRequest %s %s, It is possible that the resource is already deleted. Please verify manually \n", workspaceImportRequestId, error)
				continue
			}
		}
	}
	return nil
}

func getDataintegrationWorkspaceImportRequestIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WorkspaceImportRequestId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()

	listImportRequestsRequest := oci_dataintegration.ListImportRequestsRequest{}
	//listImportRequestsRequest.CompartmentId = &compartmentId

	workspaceIds, error := getDataintegrationWorkspaceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting workspaceId required for WorkspaceImportRequest resource requests \n")
	}
	for _, workspaceId := range workspaceIds {
		listImportRequestsRequest.WorkspaceId = &workspaceId

		listImportRequestsResponse, err := dataIntegrationClient.ListImportRequests(context.Background(), listImportRequestsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting WorkspaceImportRequest list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, workspaceImportRequest := range listImportRequestsResponse.Items {
			id := *workspaceImportRequest.Key
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WorkspaceImportRequestId", id)
		}

	}
	return resourceIds, nil
}
