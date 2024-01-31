// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataintegrationWorkspaceApplicationPatchRequiredOnlyResource = DataintegrationWorkspaceApplicationPatchResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_patch", "test_workspace_application_patch", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationPatchRepresentation)

	DataintegrationWorkspaceApplicationPatchResourceConfig = DataintegrationWorkspaceApplicationPatchResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_patch", "test_workspace_application_patch", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationPatchRepresentation)

	DataintegrationWorkspaceApplicationPatchSingularDataSourceRepresentation = map[string]interface{}{
		"application_key": acctest.Representation{RepType: acctest.Required, Create: `${var.application_key}`},
		"patch_key":       acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_application_patch.test_workspace_application_patch.key}`},
		"workspace_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
	}

	DataintegrationWorkspaceApplicationPatchDataSourceRepresentation = map[string]interface{}{
		"application_key": acctest.Representation{RepType: acctest.Required, Create: `${var.application_key}`},
		"workspace_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
		"name":            acctest.Representation{RepType: acctest.Optional, Create: `name63`, Update: `name63_2`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceApplicationPatchDataSourceFilterRepresentation}}
	DataintegrationWorkspaceApplicationPatchDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataintegration_workspace_application_patch.test_workspace_application_patch.name}`}},
	}

	// Patch representation for publish patch.
	DataintegrationWorkspaceApplicationPatchRepresentation = map[string]interface{}{
		"application_key": acctest.Representation{RepType: acctest.Required, Create: `${var.application_key}`},
		"name":            acctest.Representation{RepType: acctest.Required, Create: `name63`, Update: `name63_2`},
		"object_keys":     acctest.Representation{RepType: acctest.Required, Create: []string{`${var.task_key}`}},
		"patch_type":      acctest.Representation{RepType: acctest.Required, Create: `PUBLISH`},
		"workspace_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
	}

	// Patch representation for un-publish patch.
	DataintegrationWorkspaceApplicationPatchUnpublishRepresentation = map[string]interface{}{
		"application_key": acctest.Representation{RepType: acctest.Required, Create: `${var.application_key}`},
		"name":            acctest.Representation{RepType: acctest.Required, Create: `unpublish_name63`},
		"object_keys":     acctest.Representation{RepType: acctest.Required, Create: []string{`${var.task_key}`}},
		"patch_type":      acctest.Representation{RepType: acctest.Required, Create: `UNPUBLISH`},
		"workspace_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
	}

	// Patch representation for publish patch and optional parameters.
	DataintegrationWorkspaceApplicationPatchUnpublishOptionalRepresentation = map[string]interface{}{
		"application_key":   acctest.Representation{RepType: acctest.Required, Create: `${var.application_key}`},
		"name":              acctest.Representation{RepType: acctest.Required, Create: `publish_optional_name63`},
		"object_keys":       acctest.Representation{RepType: acctest.Required, Create: []string{`${var.task_key}`}},
		"patch_type":        acctest.Representation{RepType: acctest.Required, Create: `PUBLISH`},
		"workspace_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
		"identifier":        acctest.Representation{RepType: acctest.Optional, Create: `PUBLISH_OPTIONAL_NAME63`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"object_status":     acctest.Representation{RepType: acctest.Optional, Create: `8`},
		"registry_metadata": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceApplicationPatchUnpublishOptionalRegistryMetadataRepresentation},
	}
	DataintegrationWorkspaceApplicationPatchUnpublishOptionalRegistryMetadataRepresentation = map[string]interface{}{
		"aggregator_key": acctest.Representation{RepType: acctest.Optional, Create: `aggregatorKey`},
		"is_favorite":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"key":            acctest.Representation{RepType: acctest.Optional, Create: `key`},
		"labels":         acctest.Representation{RepType: acctest.Optional, Create: []string{`labels`}},
	}

	DataintegrationWorkspaceApplicationPatchResourceDependencies = ""
)

// issue-routing-tag: dataintegration/default
func TestDataintegrationWorkspaceApplicationPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataintegrationWorkspaceApplicationPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	workspaceId := utils.GetEnvSettingWithBlankDefault("workspace_ocid")
	workspaceIdVariableStr := fmt.Sprintf("variable \"workspace_id\" { default = \"%s\" }\n", workspaceId)

	applicationkey := "80a5b482-05da-456c-9438-92decae2de57"
	applicationKeyVariableStr := fmt.Sprintf("variable \"application_key\" { default = \"%s\" }\n", applicationkey)

	taskKey := "5a7e5877-3a51-49d3-9adb-1175ce712359"
	taskKeyVariableStr := fmt.Sprintf("variable \"task_key\" { default = \"%s\" }\n", taskKey)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId) +
		workspaceIdVariableStr + applicationKeyVariableStr + taskKeyVariableStr

	resourceName := "oci_dataintegration_workspace_application_patch.test_workspace_application_patch"
	datasourceName := "data.oci_dataintegration_workspace_application_patches.test_workspace_application_patches"
	singularDatasourceName := "data.oci_dataintegration_workspace_application_patch.test_workspace_application_patch"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_patch", "test_workspace_application_patch", acctest.Optional, acctest.Create, DataintegrationWorkspaceApplicationPatchRepresentation), "dataintegration", "workspaceApplicationPatch", t)

	acctest.ResourceTest(t, testAccCheckDataintegrationWorkspaceApplicationPatchDestroy, []resource.TestStep{
		// verify Create a patch of PUBLISH type.
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_patch", "test_workspace_application_patch", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationPatchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "application_key", applicationkey),
				resource.TestCheckResourceAttrSet(resourceName, "identifier"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "object_keys.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_type", "PUBLISH"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),
			),
		},

		// verify Create a patch of UN-PUBLISH type.
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_patch", "test_workspace_application_patch", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationPatchUnpublishRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "application_key", applicationkey),
				resource.TestCheckResourceAttrSet(resourceName, "identifier"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "object_keys.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_type", "UNPUBLISH"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),
			),
		},

		// Data integration service doesn't support delete patch that is successful.

		// verify Create a patch of PUBLISH type with optionals
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceApplicationPatchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_patch", "test_workspace_application_patch", acctest.Optional, acctest.Create, DataintegrationWorkspaceApplicationPatchUnpublishOptionalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "application_key", applicationkey),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "identifier"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "20230119"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "object_keys.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "object_status", "8"),
				resource.TestCheckResourceAttr(resourceName, "patch_type", "PUBLISH"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_metadata.0.aggregator_key"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.is_favorite", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_metadata.0.key"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.labels.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_metadata.0.registry_version"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_application_patches", "test_workspace_application_patches", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationPatchDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceApplicationPatchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_patch", "test_workspace_application_patch", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationPatchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "application_key", applicationkey),
				resource.TestCheckResourceAttrSet(datasourceName, "name"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(datasourceName, "patch_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "patch_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_application_patch", "test_workspace_application_patch", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationPatchSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceApplicationPatchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "application_key", applicationkey),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "workspace_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dependent_object_metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "identifier"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_version", "20230119"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_status", "8"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_object_metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_type", "PUBLISH"),
			),
		},
		// verify resource import
		{
			Config:            config + DataintegrationWorkspaceApplicationPatchRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"object_keys",
				"registry_metadata",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDataintegrationWorkspaceApplicationPatchDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataIntegrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataintegration_workspace_application_patch" {
			noResourceFound = false
			request := oci_dataintegration.GetPatchRequest{}

			if value, ok := rs.Primary.Attributes["application_key"]; ok {
				request.ApplicationKey = &value
			}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.PatchKey = &value
			}

			if value, ok := rs.Primary.Attributes["workspace_id"]; ok {
				request.WorkspaceId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")

			_, err := client.GetPatch(context.Background(), request)

			if err == nil {
				fmt.Printf("[ERROR] resource still exists")
				//return fmt.Errorf("resource still exists")
				return nil
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
	if !acctest.InSweeperExcludeList("DataintegrationWorkspaceApplicationPatch") {
		resource.AddTestSweepers("DataintegrationWorkspaceApplicationPatch", &resource.Sweeper{
			Name:         "DataintegrationWorkspaceApplicationPatch",
			Dependencies: acctest.DependencyGraph["workspaceApplicationPatch"],
			F:            sweepDataintegrationWorkspaceApplicationPatchResource,
		})
	}
}

func sweepDataintegrationWorkspaceApplicationPatchResource(compartment string) error {
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()
	workspaceApplicationPatchIds, err := getDataintegrationWorkspaceApplicationPatchIds(compartment)
	if err != nil {
		return err
	}
	for _, workspaceApplicationPatchId := range workspaceApplicationPatchIds {
		if ok := acctest.SweeperDefaultResourceId[workspaceApplicationPatchId]; !ok {
			deletePatchRequest := oci_dataintegration.DeletePatchRequest{}

			deletePatchRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")
			_, error := dataIntegrationClient.DeletePatch(context.Background(), deletePatchRequest)
			if error != nil {
				fmt.Printf("Error deleting WorkspaceApplicationPatch %s %s, It is possible that the resource is already deleted. Please verify manually \n", workspaceApplicationPatchId, error)
				continue
			}
		}
	}
	return nil
}

func getDataintegrationWorkspaceApplicationPatchIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WorkspaceApplicationPatchId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()

	listPatchesRequest := oci_dataintegration.ListPatchesRequest{}

	applicationKeys := []string{utils.GetEnvSettingWithBlankDefault("application_key")}

	for _, applicationKey := range applicationKeys {
		listPatchesRequest.ApplicationKey = &applicationKey

		workspaceIds, error := getDataintegrationWorkspaceIds(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting workspaceId required for WorkspaceApplicationPatch resource requests \n")
		}
		for _, workspaceId := range workspaceIds {
			listPatchesRequest.WorkspaceId = &workspaceId

			listPatchesResponse, err := dataIntegrationClient.ListPatches(context.Background(), listPatchesRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting WorkspaceApplicationPatch list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, workspaceApplicationPatch := range listPatchesResponse.Items {
				id := *workspaceApplicationPatch.Key
				resourceIds = append(resourceIds, id)
				acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WorkspaceApplicationPatchId", id)
			}

		}
	}
	return resourceIds, nil
}
