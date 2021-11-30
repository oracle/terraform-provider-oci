// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v53/common"
	oci_devops "github.com/oracle/oci-go-sdk/v53/devops"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DevopsConnectionResourceConfig = GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Optional, Update, devopsConnectionRepresentation) + DevopsConnectionResourceDependencies

	devopsConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"connection_id": Representation{RepType: Required, Create: `${oci_devops_connection.test_connection.id}`},
	}

	devopsConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"connection_type": Representation{RepType: Optional, Create: `GithubAccessToken`},
		"display_name":    Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":              Representation{RepType: Optional, Create: `${oci_devops_connection.test_connection.id}`},
		"project_id":      Representation{RepType: Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":           Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":          RepresentationGroup{Required, devopsConnectionDataSourceFilterRepresentation}}
	devopsConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_devops_connection.test_connection.id}`}},
	}

	devopsConnectionRepresentation = map[string]interface{}{
		"access_token":    Representation{RepType: Required, Create: `${var.github_access_token_vault_id}`},
		"connection_type": Representation{RepType: Required, Create: `GITHUB_ACCESS_TOKEN`},
		"project_id":      Representation{RepType: Required, Create: `${oci_devops_project.test_project.id}`},
		"defined_tags":    Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":     Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"display_name":    Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":   Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DevopsConnectionResourceDependencies = GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsConnectionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	githubAccessTokenVaultId := getEnvSettingWithBlankDefault("github_access_token_vault_id")
	githubAccessTokenVaultIdStr := fmt.Sprintf("variable \"github_access_token_vault_id\" { default = \"%s\" }\n", githubAccessTokenVaultId)

	resourceName := "oci_devops_connection.test_connection"
	datasourceName := "data.oci_devops_connections.test_connections"
	singularDatasourceName := "data.oci_devops_connection.test_connection"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+githubAccessTokenVaultIdStr+DevopsConnectionResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Optional, Create, devopsConnectionRepresentation), "devops", "connection", t)

	ResourceTest(t, testAccCheckDevopsConnectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsConnectionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Required, Create, devopsConnectionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "GITHUB_ACCESS_TOKEN"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Optional, Create, devopsConnectionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "GITHUB_ACCESS_TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Optional, Update, devopsConnectionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "GITHUB_ACCESS_TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_devops_connections", "test_connections", Optional, Update, devopsConnectionDataSourceRepresentation) +
				compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsConnectionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Optional, Update, devopsConnectionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "connection_type", "GithubAccessToken"), //TODO: Needs to accept `GITHUB_ACCESS_TOKEN`
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
				GenerateDataSourceFromRepresentationMap("oci_devops_connection", "test_connection", Required, Create, devopsConnectionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsConnectionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_type", "GITHUB_ACCESS_TOKEN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsConnectionResourceConfig,
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

func testAccCheckDevopsConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).devopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_connection" {
			noResourceFound = false
			request := oci_devops.GetConnectionRequest{}

			tmp := rs.Primary.ID
			request.ConnectionId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("DevopsConnection") {
		resource.AddTestSweepers("DevopsConnection", &resource.Sweeper{
			Name:         "DevopsConnection",
			Dependencies: DependencyGraph["connection"],
			F:            sweepDevopsConnectionResource,
		})
	}
}

func sweepDevopsConnectionResource(compartment string) error {
	devopsClient := GetTestClients(&schema.ResourceData{}).devopsClient()
	connectionIds, err := devopsGetConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, connectionId := range connectionIds {
		if ok := SweeperDefaultResourceId[connectionId]; !ok {
			deleteConnectionRequest := oci_devops.DeleteConnectionRequest{}

			deleteConnectionRequest.ConnectionId = &connectionId

			deleteConnectionRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")
			_, error := devopsClient.DeleteConnection(context.Background(), deleteConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting Connection %s %s, It is possible that the resource is already deleted. Please verify manually \n", connectionId, error)
				continue
			}
		}
	}
	return nil
}

func devopsGetConnectionIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "ConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	devopsClient := GetTestClients(&schema.ResourceData{}).devopsClient()

	listConnectionsRequest := oci_devops.ListConnectionsRequest{}
	listConnectionsRequest.CompartmentId = &compartmentId
	listConnectionsResponse, err := devopsClient.ListConnections(context.Background(), listConnectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Connection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, connection := range listConnectionsResponse.Items {
		id := *connection.GetId()
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "ConnectionId", id)
	}
	return resourceIds, nil
}
