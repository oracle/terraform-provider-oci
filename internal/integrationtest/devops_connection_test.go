// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DevopsConnectionRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Required, acctest.Create, DevopsConnectionRepresentation)

	DevopsConnectionResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Optional, acctest.Update, DevopsConnectionRepresentation) + DevopsConnectionResourceDependencies

	DevopsDevopsConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_connection.test_connection.id}`},
	}

	DevopsDevopsConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_connection.test_connection.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsConnectionDataSourceFilterRepresentation}}
	DevopsConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_connection.test_connection.id}`}},
	}

	DevopsConnectionRepresentation = map[string]interface{}{
		"access_token":    acctest.Representation{RepType: acctest.Required, Create: `${var.github_access_token_vault_id}`},
		"connection_type": acctest.Representation{RepType: acctest.Required, Create: `GITHUB_ACCESS_TOKEN`},
		"project_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	DevopsConnectionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsConnectionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	githubAccessTokenVaultId := utils.GetEnvSettingWithBlankDefault("github_access_token_vault_id")
	githubAccessTokenVaultIdStr := fmt.Sprintf("variable \"github_access_token_vault_id\" { default = \"%s\" }\n", githubAccessTokenVaultId)

	resourceName := "oci_devops_connection.test_connection"
	datasourceName := "data.oci_devops_connections.test_connections"
	singularDatasourceName := "data.oci_devops_connection.test_connection"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+githubAccessTokenVaultIdStr+DevopsConnectionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Optional, acctest.Create, DevopsConnectionRepresentation), "devops", "connection", t)

	acctest.ResourceTest(t, testAccCheckDevopsConnectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Required, acctest.Create, DevopsConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "GITHUB_ACCESS_TOKEN"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DevopsConnectionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Optional, acctest.Create, DevopsConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "GITHUB_ACCESS_TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Optional, acctest.Update, DevopsConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "GITHUB_ACCESS_TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_connections", "test_connections", acctest.Optional, acctest.Update, DevopsDevopsConnectionDataSourceRepresentation) +
				compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Optional, acctest.Update, DevopsConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "connection_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "connection_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Required, acctest.Create, DevopsDevopsConnectionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsConnectionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_type", "GITHUB_ACCESS_TOKEN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "last_connection_validation_result.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DevopsConnectionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDevopsConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DevopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_connection" {
			noResourceFound = false
			request := oci_devops.GetConnectionRequest{}

			tmp := rs.Primary.ID
			request.ConnectionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")

			_, err := client.GetConnection(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DevopsConnection") {
		resource.AddTestSweepers("DevopsConnection", &resource.Sweeper{
			Name:         "DevopsConnection",
			Dependencies: acctest.DependencyGraph["connection"],
			F:            sweepDevopsConnectionResource,
		})
	}
}

func sweepDevopsConnectionResource(compartment string) error {
	devopsClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()
	connectionIds, err := getDevopsConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, connectionId := range connectionIds {
		if ok := acctest.SweeperDefaultResourceId[connectionId]; !ok {
			deleteConnectionRequest := oci_devops.DeleteConnectionRequest{}

			deleteConnectionRequest.ConnectionId = &connectionId

			deleteConnectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")
			_, error := devopsClient.DeleteConnection(context.Background(), deleteConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting Connection %s %s, It is possible that the resource is already deleted. Please verify manually \n", connectionId, error)
				continue
			}
		}
	}
	return nil
}

func getDevopsConnectionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	devopsClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()

	listConnectionsRequest := oci_devops.ListConnectionsRequest{}
	listConnectionsRequest.CompartmentId = &compartmentId
	listConnectionsResponse, err := devopsClient.ListConnections(context.Background(), listConnectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Connection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, connection := range listConnectionsResponse.Items {
		id := *connection.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ConnectionId", id)
	}
	return resourceIds, nil
}
