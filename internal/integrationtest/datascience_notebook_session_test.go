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
	oci_datascience "github.com/oracle/oci-go-sdk/v56/datascience"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NotebookSessionRequiredOnlyResource = NotebookSessionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, notebookSessionRepresentation)

	NotebookSessionResourceConfig = NotebookSessionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, notebookSessionRepresentation)

	notebookSessionSingularDataSourceRepresentation = map[string]interface{}{
		"notebook_session_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_notebook_session.test_notebook_session.id}`},
	}

	notebookSessionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_notebook_session.test_notebook_session.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: notebookSessionDataSourceFilterRepresentation}}
	notebookSessionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_notebook_session.test_notebook_session.id}`}},
	}

	notebookSessionRepresentation = map[string]interface{}{
		"compartment_id":                         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"notebook_session_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: notebookSessionNotebookSessionConfigurationDetailsRepresentation},
		"project_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"defined_tags":                           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	notebookSessionNotebookSessionConfigurationDetailsRepresentation = map[string]interface{}{
		"shape":                     acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E2.2`, Update: `VM.Standard.E2.4`},
		"subnet_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"block_storage_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `50`, Update: `51`},
	}

	NotebookSessionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, projectRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceNotebookSessionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceNotebookSessionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_notebook_session.test_notebook_session"
	datasourceName := "data.oci_datascience_notebook_sessions.test_notebook_sessions"
	singularDatasourceName := "data.oci_datascience_notebook_session.test_notebook_session"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+NotebookSessionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create, notebookSessionRepresentation), "datascience", "notebookSession", t)

	acctest.ResourceTest(t, testAccCheckDatascienceNotebookSessionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, notebookSessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E2.2"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_configuration_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create, notebookSessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.block_storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E2.2"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_configuration_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(notebookSessionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.block_storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E2.2"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_configuration_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
			Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, notebookSessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.block_storage_size_in_gbs", "51"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E2.4"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_configuration_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_notebook_sessions", "test_notebook_sessions", acctest.Optional, acctest.Update, notebookSessionDataSourceRepresentation) +
				compartmentIdVariableStr + NotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, notebookSessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.created_by"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.0.block_storage_size_in_gbs", "51"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.0.shape", "VM.Standard.E2.4"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.notebook_session_url"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.project_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, notebookSessionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NotebookSessionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_configuration_details.0.block_storage_size_in_gbs", "51"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E2.4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + NotebookSessionResourceConfig,
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

func testAccCheckDatascienceNotebookSessionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_notebook_session" {
			noResourceFound = false
			request := oci_datascience.GetNotebookSessionRequest{}

			tmp := rs.Primary.ID
			request.NotebookSessionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetNotebookSession(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.NotebookSessionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatascienceNotebookSession") {
		resource.AddTestSweepers("DatascienceNotebookSession", &resource.Sweeper{
			Name:         "DatascienceNotebookSession",
			Dependencies: acctest.DependencyGraph["notebookSession"],
			F:            sweepDatascienceNotebookSessionResource,
		})
	}
}

func sweepDatascienceNotebookSessionResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	notebookSessionIds, err := getNotebookSessionIds(compartment)
	if err != nil {
		return err
	}
	for _, notebookSessionId := range notebookSessionIds {
		if ok := acctest.SweeperDefaultResourceId[notebookSessionId]; !ok {
			deleteNotebookSessionRequest := oci_datascience.DeleteNotebookSessionRequest{}

			deleteNotebookSessionRequest.NotebookSessionId = &notebookSessionId

			deleteNotebookSessionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteNotebookSession(context.Background(), deleteNotebookSessionRequest)
			if error != nil {
				fmt.Printf("Error deleting NotebookSession %s %s, It is possible that the resource is already deleted. Please verify manually \n", notebookSessionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &notebookSessionId, notebookSessionSweepWaitCondition, time.Duration(3*time.Minute),
				notebookSessionSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getNotebookSessionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NotebookSessionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listNotebookSessionsRequest := oci_datascience.ListNotebookSessionsRequest{}
	listNotebookSessionsRequest.CompartmentId = &compartmentId
	listNotebookSessionsRequest.LifecycleState = oci_datascience.ListNotebookSessionsLifecycleStateActive
	listNotebookSessionsResponse, err := dataScienceClient.ListNotebookSessions(context.Background(), listNotebookSessionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NotebookSession list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, notebookSession := range listNotebookSessionsResponse.Items {
		id := *notebookSession.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NotebookSessionId", id)
	}
	return resourceIds, nil
}

func notebookSessionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if notebookSessionResponse, ok := response.Response.(oci_datascience.GetNotebookSessionResponse); ok {
		return notebookSessionResponse.LifecycleState != oci_datascience.NotebookSessionLifecycleStateDeleted
	}
	return false
}

func notebookSessionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetNotebookSession(context.Background(), oci_datascience.GetNotebookSessionRequest{
		NotebookSessionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
