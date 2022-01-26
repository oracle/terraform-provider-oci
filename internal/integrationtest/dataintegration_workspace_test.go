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
	oci_dataintegration "github.com/oracle/oci-go-sdk/v56/dataintegration"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	WorkspaceRequiredOnlyResource = WorkspaceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Required, acctest.Create, workspaceRepresentation)

	WorkspaceResourceConfig = WorkspaceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Optional, acctest.Update, workspaceRepresentation)

	workspaceSingularDataSourceRepresentation = map[string]interface{}{
		"workspace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
	}

	workspaceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: workspaceDataSourceFilterRepresentation}}
	workspaceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataintegration_workspace.test_workspace.id}`}},
	}

	workspaceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_private_network_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"subnet_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	WorkspaceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`}})) +
		DefinedTagsDependencies
)

// issue-routing-tag: dataintegration/default
func TestDataintegrationWorkspaceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataintegrationWorkspaceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dataintegration_workspace.test_workspace"
	datasourceName := "data.oci_dataintegration_workspaces.test_workspaces"
	singularDatasourceName := "data.oci_dataintegration_workspace.test_workspace"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+WorkspaceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Optional, acctest.Create, workspaceRepresentation), "dataintegration", "workspace", t)

	acctest.ResourceTest(t, testAccCheckDataintegrationWorkspaceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + WorkspaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("is_private_network_enabled", acctest.Representation{RepType: acctest.Required, Create: `false`}, workspaceRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "is_private_network_enabled", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + WorkspaceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + WorkspaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Optional, acctest.Create, workspaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_private_network_enabled", "true"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + WorkspaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(workspaceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_private_network_enabled", "true"),

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
			Config: config + compartmentIdVariableStr + WorkspaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Optional, acctest.Update, workspaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_private_network_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspaces", "test_workspaces", acctest.Optional, acctest.Update, workspaceDataSourceRepresentation) +
				compartmentIdVariableStr + WorkspaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Optional, acctest.Update, workspaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "workspaces.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "workspaces.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "workspaces.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "workspaces.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "workspaces.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspaces.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspaces.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspaces.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspaces.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Required, acctest.Create, workspaceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + WorkspaceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_private_network_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "state_message"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + WorkspaceResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"is_private_network_enabled", "state_message"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataintegrationWorkspaceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataIntegrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataintegration_workspace" {
			noResourceFound = false
			request := oci_dataintegration.GetWorkspaceRequest{}

			tmp := rs.Primary.ID
			request.WorkspaceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")

			response, err := client.GetWorkspace(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dataintegration.WorkspaceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataintegrationWorkspace") {
		resource.AddTestSweepers("DataintegrationWorkspace", &resource.Sweeper{
			Name:         "DataintegrationWorkspace",
			Dependencies: acctest.DependencyGraph["workspace"],
			F:            sweepDataintegrationWorkspaceResource,
		})
	}
}

func sweepDataintegrationWorkspaceResource(compartment string) error {
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()
	workspaceIds, err := getWorkspaceIds(compartment)
	if err != nil {
		return err
	}
	for _, workspaceId := range workspaceIds {
		if ok := acctest.SweeperDefaultResourceId[workspaceId]; !ok {
			deleteWorkspaceRequest := oci_dataintegration.DeleteWorkspaceRequest{}

			deleteWorkspaceRequest.WorkspaceId = &workspaceId

			deleteWorkspaceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")
			_, error := dataIntegrationClient.DeleteWorkspace(context.Background(), deleteWorkspaceRequest)
			if error != nil {
				fmt.Printf("Error deleting Workspace %s %s, It is possible that the resource is already deleted. Please verify manually \n", workspaceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &workspaceId, workspaceSweepWaitCondition, time.Duration(3*time.Minute),
				workspaceSweepResponseFetchOperation, "dataintegration", true)
		}
	}
	return nil
}

func getWorkspaceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WorkspaceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()

	listWorkspacesRequest := oci_dataintegration.ListWorkspacesRequest{}
	listWorkspacesRequest.CompartmentId = &compartmentId
	listWorkspacesRequest.LifecycleState = oci_dataintegration.WorkspaceLifecycleStateActive
	listWorkspacesResponse, err := dataIntegrationClient.ListWorkspaces(context.Background(), listWorkspacesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Workspace list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, workspace := range listWorkspacesResponse.Items {
		id := *workspace.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WorkspaceId", id)
	}
	return resourceIds, nil
}

func workspaceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if workspaceResponse, ok := response.Response.(oci_dataintegration.GetWorkspaceResponse); ok {
		return workspaceResponse.LifecycleState != oci_dataintegration.WorkspaceLifecycleStateDeleted
	}
	return false
}

func workspaceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataIntegrationClient().GetWorkspace(context.Background(), oci_dataintegration.GetWorkspaceRequest{
		WorkspaceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
