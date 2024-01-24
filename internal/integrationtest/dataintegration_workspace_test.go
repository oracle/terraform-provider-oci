// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataintegrationWorkspaceRequiredOnlyResource = DataintegrationWorkspaceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Required, acctest.Create, DataintegrationWorkspaceRepresentation)

	DataintegrationWorkspaceResourceConfig = DataintegrationWorkspaceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Optional, acctest.Update, DataintegrationWorkspaceRepresentation)

	DataintegrationDataintegrationWorkspaceSingularDataSourceRepresentation = map[string]interface{}{
		"workspace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
	}

	DataintegrationDataintegrationWorkspaceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceDataSourceFilterRepresentation}}
	DataintegrationWorkspaceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataintegration_workspace.test_workspace.id}`}},
	}

	DataintegrationWorkspaceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"endpoint_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.private_endpoint_ocid}`},
		"endpoint_name":              acctest.Representation{RepType: acctest.Optional, Create: `PE2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_private_network_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"registry_compartment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.dcms_registry_comp_id}`},
		"registry_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.dcms_registry_id}`},
		"dns_server_ip":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataintegrationWorkspaceDefinedTagsChangesRepresentation},
	}

	ignoreDataintegrationWorkspaceDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DataintegrationWorkspaceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`}})) +
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

	dcmsRegistryId := utils.GetEnvSettingWithBlankDefault("dcms_registry_ocid")
	dcmsRegistryIdVariableStr := fmt.Sprintf("variable \"dcms_registry_id\" { default = \"%s\" }\n", dcmsRegistryId)
	dcmsRegistryCompId := utils.GetEnvSettingWithBlankDefault("dcms_registry_compartment_ocid")
	dcmsRegistryCompIdVariableStr := fmt.Sprintf("variable \"dcms_registry_comp_id\" { default = \"%s\" }\n", dcmsRegistryCompId)
	dcmsRegistryName := utils.GetEnvSettingWithBlankDefault("dcms_registry_name")
	dcmsRegistryNameVariableStr := fmt.Sprintf("variable \"dcms_registry_name\" { default = \"%s\" }\n", dcmsRegistryName)
	endpointId := utils.GetEnvSettingWithBlankDefault("private_endpoint_ocid")
	endpointIdVariableStr := fmt.Sprintf("variable \"private_endpoint_ocid\" { default = \"%s\" }\n", endpointId)

	resourceName := "oci_dataintegration_workspace.test_workspace"
	datasourceName := "data.oci_dataintegration_workspaces.test_workspaces"
	singularDatasourceName := "data.oci_dataintegration_workspace.test_workspace"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dcmsRegistryIdVariableStr+dcmsRegistryNameVariableStr+dcmsRegistryCompIdVariableStr+endpointIdVariableStr+DataintegrationWorkspaceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Optional, acctest.Create, DataintegrationWorkspaceRepresentation), "dataintegration", "workspace", t)

	acctest.ResourceTest(t, testAccCheckDataintegrationWorkspaceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("is_private_network_enabled", acctest.Representation{RepType: acctest.Required, Create: `false`}, DataintegrationWorkspaceRepresentation)),
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
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + dcmsRegistryIdVariableStr + dcmsRegistryNameVariableStr + dcmsRegistryCompIdVariableStr + endpointIdVariableStr + DataintegrationWorkspaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Optional, acctest.Create, DataintegrationWorkspaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_name", "PE2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_private_network_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_id"),

				//TEMPORARILY DISABLING RESOURCE DISCOVERY TEST WHICH ALWAYS FAILS DUE TO AD IN ASHBURN NOT LONDON ISSUES
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					/*if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}*/
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + dcmsRegistryIdVariableStr + dcmsRegistryNameVariableStr + dcmsRegistryCompIdVariableStr + endpointIdVariableStr + DataintegrationWorkspaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataintegrationWorkspaceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_name", "PE2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_private_network_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be moved")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + dcmsRegistryIdVariableStr + dcmsRegistryNameVariableStr + dcmsRegistryCompIdVariableStr + endpointIdVariableStr + DataintegrationWorkspaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Optional, acctest.Update, DataintegrationWorkspaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_name", "PE2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_private_network_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspaces", "test_workspaces", acctest.Optional, acctest.Update, DataintegrationDataintegrationWorkspaceDataSourceRepresentation) +
				compartmentIdVariableStr + dcmsRegistryIdVariableStr + dcmsRegistryNameVariableStr + dcmsRegistryCompIdVariableStr + endpointIdVariableStr + DataintegrationWorkspaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Optional, acctest.Update, DataintegrationWorkspaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "workspaces.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "workspaces.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "workspaces.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "workspaces.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspaces.0.endpoint_id"),
				resource.TestCheckResourceAttr(datasourceName, "workspaces.0.endpoint_name", "PE2"),
				resource.TestCheckResourceAttr(datasourceName, "workspaces.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspaces.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspaces.0.registry_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspaces.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspaces.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspaces.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Required, acctest.Create, DataintegrationDataintegrationWorkspaceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dcmsRegistryIdVariableStr + dcmsRegistryNameVariableStr + dcmsRegistryCompIdVariableStr + endpointIdVariableStr + DataintegrationWorkspaceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "workspace_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_private_network_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DataintegrationWorkspaceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{"is_private_network_enabled", "state_message",
				"registry_compartment_id",
				"registry_name",
			},
			ResourceName: resourceName,
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
	workspaceIds, err := getDataintegrationWorkspaceIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &workspaceId, DataintegrationWorkspaceSweepWaitCondition, time.Duration(3*time.Minute),
				DataintegrationWorkspaceSweepResponseFetchOperation, "dataintegration", true)
		}
	}
	return nil
}

func getDataintegrationWorkspaceIds(compartment string) ([]string, error) {
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

func DataintegrationWorkspaceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if workspaceResponse, ok := response.Response.(oci_dataintegration.GetWorkspaceResponse); ok {
		return workspaceResponse.LifecycleState != oci_dataintegration.WorkspaceLifecycleStateDeleted
	}
	return false
}

func DataintegrationWorkspaceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataIntegrationClient().GetWorkspace(context.Background(), oci_dataintegration.GetWorkspaceRequest{
		WorkspaceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
