// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
	DataintegrationWorkspaceApplicationRequiredOnlyResource = DataintegrationWorkspaceApplicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application", "test_workspace_application", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationRepresentation)

	DataintegrationWorkspaceApplicationResourceConfig = DataintegrationWorkspaceApplicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application", "test_workspace_application", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationRepresentation)

	DataintegrationWorkspaceApplicationSingularDataSourceRepresentation = map[string]interface{}{
		"application_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_application.test_workspace_application.key}`},
		"workspace_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
	}

	DataintegrationWorkspaceApplicationDataSourceRepresentation = map[string]interface{}{
		"workspace_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
		"identifier":    acctest.Representation{RepType: acctest.Optional, Create: []string{`APPLICATION_TF_TEST_1`}, Update: []string{`APPLICATION_TF_TEST_2`}},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: `application_tf_test_1`, Update: `application_tf_test_2`},
		"name_contains": acctest.Representation{RepType: acctest.Optional, Create: `application_tf_test_2`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceApplicationDataSourceFilterRepresentation}}
	DataintegrationWorkspaceApplicationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataintegration_workspace_application.test_workspace_application.key}`}},
	}

	DataintegrationWorkspaceApplicationRepresentation = map[string]interface{}{
		"identifier":        acctest.Representation{RepType: acctest.Required, Create: `APPLICATION_TF_TEST_1`, Update: `APPLICATION_TF_TEST_2`},
		"name":              acctest.Representation{RepType: acctest.Required, Create: `application_tf_test_1`, Update: `application_tf_test_2`},
		"workspace_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"model_type":        acctest.Representation{RepType: acctest.Required, Create: `INTEGRATION_APPLICATION`},
		"model_version":     acctest.Representation{RepType: acctest.Optional, Create: `20230426`, Update: `20230426`},
		"object_status":     acctest.Representation{RepType: acctest.Optional, Create: `8`, Update: `8`},
		"registry_metadata": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceApplicationRegistryMetadataRepresentation},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `ACTIVE`},
	}
	DataintegrationWorkspaceApplicationRegistryMetadataRepresentation = map[string]interface{}{
		"is_favorite": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"labels":      acctest.Representation{RepType: acctest.Optional, Create: []string{`labels`}},
	}

	DataintegrationWorkspaceApplicationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("is_private_network_enabled", acctest.Representation{RepType: acctest.Required, Create: `false`}, DataintegrationWorkspaceRepresentation)) +
		DefinedTagsDependencies
)

// issue-routing-tag: dataintegration/default
func TestDataintegrationWorkspaceApplicationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataintegrationWorkspaceApplicationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dataintegration_workspace_application.test_workspace_application"
	datasourceName := "data.oci_dataintegration_workspace_applications.test_workspace_applications"
	singularDatasourceName := "data.oci_dataintegration_workspace_application.test_workspace_application"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataintegrationWorkspaceApplicationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application", "test_workspace_application", acctest.Optional, acctest.Create, DataintegrationWorkspaceApplicationRepresentation), "dataintegration", "workspaceApplication", t)

	acctest.ResourceTest(t, testAccCheckDataintegrationWorkspaceApplicationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application", "test_workspace_application", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "identifier", "APPLICATION_TF_TEST_1"),
				resource.TestCheckResourceAttr(resourceName, "name", "application_tf_test_1"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceApplicationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application", "test_workspace_application", acctest.Optional, acctest.Create, DataintegrationWorkspaceApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "APPLICATION_TF_TEST_1"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "INTEGRATION_APPLICATION"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "20230426"),
				resource.TestCheckResourceAttr(resourceName, "name", "application_tf_test_1"),
				resource.TestCheckResourceAttr(resourceName, "object_status", "8"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.is_favorite", "false"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.labels.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_metadata.0.registry_version"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application", "test_workspace_application", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "APPLICATION_TF_TEST_2"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "INTEGRATION_APPLICATION"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "20230426"),
				resource.TestCheckResourceAttr(resourceName, "name", "application_tf_test_2"),
				resource.TestCheckResourceAttr(resourceName, "object_status", "8"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.is_favorite", "false"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_applications", "test_workspace_applications", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application", "test_workspace_application", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "identifier.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "name", "application_tf_test_2"),
				resource.TestCheckResourceAttr(datasourceName, "name_contains", "application_tf_test_2"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(datasourceName, "application_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "application_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_application", "test_workspace_application", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceApplicationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "workspace_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "identifier", "APPLICATION_TF_TEST_2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_type", "INTEGRATION_APPLICATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_version", "20230426"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "application_tf_test_2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_status", "8"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DataintegrationWorkspaceApplicationRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"registry_metadata",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDataintegrationWorkspaceApplicationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataIntegrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataintegration_workspace_application" {
			noResourceFound = false
			request := oci_dataintegration.GetApplicationRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.ApplicationKey = &value
			}

			if value, ok := rs.Primary.Attributes["workspace_id"]; ok {
				request.WorkspaceId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")

			response, err := client.GetApplication(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dataintegration.ApplicationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataintegrationWorkspaceApplication") {
		resource.AddTestSweepers("DataintegrationWorkspaceApplication", &resource.Sweeper{
			Name:         "DataintegrationWorkspaceApplication",
			Dependencies: acctest.DependencyGraph["workspaceApplication"],
			F:            sweepDataintegrationWorkspaceApplicationResource,
		})
	}
}

func sweepDataintegrationWorkspaceApplicationResource(compartment string) error {
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()
	workspaceApplicationIds, err := getDataintegrationWorkspaceApplicationIds(compartment)
	if err != nil {
		return err
	}
	for _, workspaceApplicationId := range workspaceApplicationIds {
		if ok := acctest.SweeperDefaultResourceId[workspaceApplicationId]; !ok {
			deleteApplicationRequest := oci_dataintegration.DeleteApplicationRequest{}

			deleteApplicationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")
			_, error := dataIntegrationClient.DeleteApplication(context.Background(), deleteApplicationRequest)
			if error != nil {
				fmt.Printf("Error deleting WorkspaceApplication %s %s, It is possible that the resource is already deleted. Please verify manually \n", workspaceApplicationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &workspaceApplicationId, DataintegrationWorkspaceApplicationSweepWaitCondition, time.Duration(3*time.Minute),
				DataintegrationWorkspaceApplicationSweepResponseFetchOperation, "dataintegration", true)
		}
	}
	return nil
}

func getDataintegrationWorkspaceApplicationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WorkspaceApplicationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()

	listApplicationsRequest := oci_dataintegration.ListApplicationsRequest{}

	workspaceIds, error := getDataintegrationWorkspaceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting workspaceId required for WorkspaceApplication resource requests \n")
	}
	for _, workspaceId := range workspaceIds {
		listApplicationsRequest.WorkspaceId = &workspaceId

		listApplicationsResponse, err := dataIntegrationClient.ListApplications(context.Background(), listApplicationsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting WorkspaceApplication list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, workspaceApplication := range listApplicationsResponse.Items {
			id := *workspaceApplication.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WorkspaceApplicationId", id)
		}

	}
	return resourceIds, nil
}

func DataintegrationWorkspaceApplicationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if workspaceApplicationResponse, ok := response.Response.(oci_dataintegration.GetApplicationResponse); ok {
		return workspaceApplicationResponse.LifecycleState != oci_dataintegration.ApplicationLifecycleStateDeleted
	}
	return false
}

func DataintegrationWorkspaceApplicationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataIntegrationClient().GetApplication(context.Background(), oci_dataintegration.GetApplicationRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
