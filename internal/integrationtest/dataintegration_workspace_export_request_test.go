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
	DataintegrationWorkspaceExportRequestRequiredOnlyResource = DataintegrationWorkspaceExportRequestResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_export_request", "test_workspace_export_request", acctest.Required, acctest.Create, DataintegrationWorkspaceExportRequestRepresentation)

	DataintegrationWorkspaceExportRequestResourceConfig = DataintegrationWorkspaceExportRequestResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_export_request", "test_workspace_export_request", acctest.Optional, acctest.Update, DataintegrationWorkspaceExportRequestRepresentation)

	DataintegrationWorkspaceExportRequestSingularDataSourceRepresentation = map[string]interface{}{
		"export_request_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_export_request.test_workspace_export_request.key}`},
		"workspace_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
	}

	DataintegrationWorkspaceExportRequestDataSourceRepresentation = map[string]interface{}{
		"workspace_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
		"export_status": acctest.Representation{RepType: acctest.Optional, Create: `SUCCESSFUL`},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_dataintegration_workspace_export_request.test_workspace_export_request.name}`},
		"projection":    acctest.Representation{RepType: acctest.Optional, Create: `SUMMARY`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceExportRequestDataSourceFilterRepresentation}}
	DataintegrationWorkspaceExportRequestDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataintegration_workspace_export_request.test_workspace_export_request.name}`}},
	}

	DataintegrationWorkspaceExportRequestRepresentation = map[string]interface{}{
		"bucket":                      acctest.Representation{RepType: acctest.Required, Create: `${var.bucket_name}`},
		"workspace_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
		"are_references_included":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"file_name":                   acctest.Representation{RepType: acctest.Optional, Create: `MyExportObjects.zip`},
		"filters":                     acctest.Representation{RepType: acctest.Optional, Create: nil},
		"is_object_overwrite_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"object_keys":                 acctest.Representation{RepType: acctest.Optional, Create: nil},
		"object_storage_region":       acctest.Representation{RepType: acctest.Optional, Create: `${var.region_name}`},
		"object_storage_tenancy_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.tenancy_id}`},
	}

	DataintegrationWorkspaceExportRequestResourceDependencies = ""
)

// issue-routing-tag: dataintegration/default
func TestDataintegrationWorkspaceExportRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataintegrationWorkspaceExportRequestResource_basic")
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

	resourceName := "oci_dataintegration_workspace_export_request.test_workspace_export_request"
	datasourceName := "data.oci_dataintegration_workspace_export_requests.test_workspace_export_requests"
	singularDatasourceName := "data.oci_dataintegration_workspace_export_request.test_workspace_export_request"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+workspaceIdVariableStr+tenancyIdVariableStr+regionVariableStr+bucketVariableStr+DataintegrationWorkspaceExportRequestResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_export_request", "test_workspace_export_request", acctest.Optional, acctest.Create, DataintegrationWorkspaceExportRequestRepresentation), "dataintegration", "workspaceExportRequest", t)

	acctest.ResourceTest(t, testAccCheckDataintegrationWorkspaceExportRequestDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + workspaceIdVariableStr + bucketVariableStr + DataintegrationWorkspaceExportRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_export_request", "test_workspace_export_request", acctest.Required, acctest.Create, DataintegrationWorkspaceExportRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bucket"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceExportRequestResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + workspaceIdVariableStr + tenancyIdVariableStr + regionVariableStr + bucketVariableStr + DataintegrationWorkspaceExportRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_export_request", "test_workspace_export_request", acctest.Optional, acctest.Create, DataintegrationWorkspaceExportRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_references_included", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket"),
				resource.TestCheckResourceAttr(resourceName, "file_name", "MyExportObjects.zip"),
				resource.TestCheckResourceAttr(resourceName, "filters.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "is_object_overwrite_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "object_keys.#", "3"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_export_requests", "test_workspace_export_requests", acctest.Optional, acctest.Update, DataintegrationWorkspaceExportRequestDataSourceRepresentation) +
				compartmentIdVariableStr + workspaceIdVariableStr + tenancyIdVariableStr + regionVariableStr + bucketVariableStr + DataintegrationWorkspaceExportRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_export_request", "test_workspace_export_request", acctest.Optional, acctest.Update, DataintegrationWorkspaceExportRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "export_status", "SUCCESSFUL"),
				resource.TestCheckResourceAttrSet(datasourceName, "name"),
				resource.TestCheckResourceAttr(datasourceName, "projection", "SUMMARY"),
				//resource.TestCheckResourceAttr(datasourceName, "time_ended_in_millis", "10"),
				//resource.TestCheckResourceAttr(datasourceName, "time_started_in_millis", "10"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(datasourceName, "export_request_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "export_request_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_export_request", "test_workspace_export_request", acctest.Required, acctest.Create, DataintegrationWorkspaceExportRequestSingularDataSourceRepresentation) +
				compartmentIdVariableStr + workspaceIdVariableStr + tenancyIdVariableStr + regionVariableStr + bucketVariableStr + DataintegrationWorkspaceExportRequestResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "export_request_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "are_references_included", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bucket"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "exported_items.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "file_name", "MyExportObjects.zip"),
				resource.TestCheckResourceAttr(singularDatasourceName, "filters.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_object_overwrite_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_keys.#", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_storage_region"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "referenced_items"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended_in_millis"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started_in_millis"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_exported_object_count"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataintegrationWorkspaceExportRequestRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataintegrationWorkspaceExportRequestDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataIntegrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataintegration_workspace_export_request" {
			noResourceFound = false
			request := oci_dataintegration.GetExportRequestRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.ExportRequestKey = &value
			}

			if value, ok := rs.Primary.Attributes["workspace_id"]; ok {
				request.WorkspaceId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")

			_, err := client.GetExportRequest(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataintegrationWorkspaceExportRequest") {
		resource.AddTestSweepers("DataintegrationWorkspaceExportRequest", &resource.Sweeper{
			Name:         "DataintegrationWorkspaceExportRequest",
			Dependencies: acctest.DependencyGraph["workspaceExportRequest"],
			F:            sweepDataintegrationWorkspaceExportRequestResource,
		})
	}
}

func sweepDataintegrationWorkspaceExportRequestResource(compartment string) error {
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()
	workspaceExportRequestIds, err := getDataintegrationWorkspaceExportRequestIds(compartment)
	if err != nil {
		return err
	}
	for _, workspaceExportRequestId := range workspaceExportRequestIds {
		if ok := acctest.SweeperDefaultResourceId[workspaceExportRequestId]; !ok {
			deleteExportRequestRequest := oci_dataintegration.DeleteExportRequestRequest{}

			deleteExportRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")
			_, error := dataIntegrationClient.DeleteExportRequest(context.Background(), deleteExportRequestRequest)
			if error != nil {
				fmt.Printf("Error deleting WorkspaceExportRequest %s %s, It is possible that the resource is already deleted. Please verify manually \n", workspaceExportRequestId, error)
				continue
			}
		}
	}
	return nil
}

func getDataintegrationWorkspaceExportRequestIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WorkspaceExportRequestId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()

	listExportRequestsRequest := oci_dataintegration.ListExportRequestsRequest{}
	//listExportRequestsRequest.CompartmentId = &compartmentId

	workspaceIds, error := getDataintegrationWorkspaceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting workspaceId required for WorkspaceExportRequest resource requests \n")
	}
	for _, workspaceId := range workspaceIds {
		listExportRequestsRequest.WorkspaceId = &workspaceId

		listExportRequestsResponse, err := dataIntegrationClient.ListExportRequests(context.Background(), listExportRequestsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting WorkspaceExportRequest list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, workspaceExportRequest := range listExportRequestsResponse.Items {
			id := *workspaceExportRequest.Key
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WorkspaceExportRequestId", id)
		}

	}
	return resourceIds, nil
}
