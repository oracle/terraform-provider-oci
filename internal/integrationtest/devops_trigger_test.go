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
	DevopsTriggerRequiredOnlyResource = DevopsTriggerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Required, acctest.Create, DevopsTriggerRepresentation)

	DevopsTriggerResourceConfig = DevopsTriggerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Optional, acctest.Update, DevopsTriggerRepresentation)

	DevopsDevopsTriggerSingularDataSourceRepresentation = map[string]interface{}{
		"trigger_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_trigger.test_trigger.id}`},
	}

	DevopsDevopsTriggerDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsTriggerDataSourceFilterRepresentation}}
	DevopsTriggerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_trigger.test_trigger.id}`}},
	}

	DevopsTriggerRepresentation = map[string]interface{}{
		"actions":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsTriggerActionsRepresentation},
		"project_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"trigger_source": acctest.Representation{RepType: acctest.Required, Create: `GITHUB`},
		"connection_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_connection.test_connection.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}
	DevopsTriggerActionsRepresentation = map[string]interface{}{
		"build_pipeline_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"type":              acctest.Representation{RepType: acctest.Required, Create: `TRIGGER_BUILD_PIPELINE`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: DevopsTriggerActionsFilterRepresentation},
	}
	DevopsTriggerActionsFilterRepresentation = map[string]interface{}{
		"trigger_source": acctest.Representation{RepType: acctest.Required, Create: `GITHUB`, Update: `GITHUB`},
		"events":         acctest.Representation{RepType: acctest.Optional, Create: []string{`PUSH`}, Update: []string{`PUSH`}},
		"include":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DevopsTriggerActionsFilterIncludeRepresentation},
		"exclude":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: triggerActionsFilterExcludeRepresentation},
	}
	triggerActionsFilterExcludeRepresentation = map[string]interface{}{
		"file_filter": acctest.RepresentationGroup{RepType: acctest.Optional, Group: triggerActionsFilterExcludeFileFilterRepresentation},
	}
	DevopsTriggerActionsFilterIncludeRepresentation = map[string]interface{}{
		"base_ref":    acctest.Representation{RepType: acctest.Optional, Create: `baseRef`, Update: `baseRef2`},
		"file_filter": acctest.RepresentationGroup{RepType: acctest.Optional, Group: triggerActionsFilterIncludeFileFilterRepresentation},
		"head_ref":    acctest.Representation{RepType: acctest.Optional, Create: `headRef`, Update: `headRef2`},
	}
	triggerActionsFilterExcludeFileFilterRepresentation = map[string]interface{}{
		"file_paths": acctest.Representation{RepType: acctest.Optional, Create: []string{`filePaths1`}, Update: []string{`filePaths1`}},
	}
	triggerActionsFilterIncludeFileFilterRepresentation = map[string]interface{}{
		"file_paths": acctest.Representation{RepType: acctest.Optional, Create: []string{`filePaths2`}, Update: []string{`filePaths2`}},
	}

	DevopsTriggerResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", acctest.Required, acctest.Create, DevopsBuildPipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Required, acctest.Create, DevopsConnectionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Create, DevopsRepositoryRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsTriggerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsTriggerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	githubAccessTokenVaultId := utils.GetEnvSettingWithBlankDefault("github_access_token_vault_id")
	githubAccessTokenVaultIdStr := fmt.Sprintf("variable \"github_access_token_vault_id\" { default = \"%s\" }\n", githubAccessTokenVaultId)

	resourceName := "oci_devops_trigger.test_trigger"
	datasourceName := "data.oci_devops_triggers.test_triggers"
	singularDatasourceName := "data.oci_devops_trigger.test_trigger"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+githubAccessTokenVaultIdStr+DevopsTriggerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Optional, acctest.Create, DevopsTriggerRepresentation), "devops", "trigger", t)

	acctest.ResourceTest(t, testAccCheckDevopsTriggerDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsTriggerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Required, acctest.Create, DevopsTriggerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "actions.0.build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "TRIGGER_BUILD_PIPELINE"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "trigger_source", "GITHUB"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsTriggerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Optional, acctest.Create, DevopsTriggerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "actions.0.build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.events.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.exclude.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.exclude.0.file_filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.exclude.0.file_filter.0.file_paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.base_ref", "baseRef"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.file_filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.file_filter.0.file_paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.head_ref", "headRef"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.trigger_source", "GITHUB"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "TRIGGER_BUILD_PIPELINE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "trigger_source", "GITHUB"),
				resource.TestCheckResourceAttrSet(resourceName, "trigger_url"),

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
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsTriggerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Optional, acctest.Update, DevopsTriggerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "actions.0.build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.events.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.exclude.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.exclude.0.file_filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.exclude.0.file_filter.0.file_paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.base_ref", "baseRef2"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.file_filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.file_filter.0.file_paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.head_ref", "headRef2"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.trigger_source", "GITHUB"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "TRIGGER_BUILD_PIPELINE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "trigger_source", "GITHUB"),
				resource.TestCheckResourceAttrSet(resourceName, "trigger_url"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_triggers", "test_triggers", acctest.Optional, acctest.Update, DevopsDevopsTriggerDataSourceRepresentation) +
				compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsTriggerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Optional, acctest.Update, DevopsTriggerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				resource.TestCheckResourceAttr(datasourceName, "trigger_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Required, acctest.Create, DevopsDevopsTriggerSingularDataSourceRepresentation) +
				compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsTriggerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trigger_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.events.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.exclude.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.exclude.0.file_filter.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.exclude.0.file_filter.0.file_paths.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.include.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.include.0.base_ref", "baseRef2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.include.0.file_filter.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.include.0.file_filter.0.file_paths.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.include.0.head_ref", "headRef2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.trigger_source", "GITHUB"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.type", "TRIGGER_BUILD_PIPELINE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trigger_source", "GITHUB"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trigger_url"),
			),
		},
		// verify resource import
		{
			Config:                  config + DevopsTriggerRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDevopsTriggerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DevopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_trigger" {
			noResourceFound = false
			request := oci_devops.GetTriggerRequest{}

			tmp := rs.Primary.ID
			request.TriggerId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")

			_, err := client.GetTrigger(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DevopsTrigger") {
		resource.AddTestSweepers("DevopsTrigger", &resource.Sweeper{
			Name:         "DevopsTrigger",
			Dependencies: acctest.DependencyGraph["trigger"],
			F:            sweepDevopsTriggerResource,
		})
	}
}

func sweepDevopsTriggerResource(compartment string) error {
	devopsClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()
	triggerIds, err := getDevopsTriggerIds(compartment)
	if err != nil {
		return err
	}
	for _, triggerId := range triggerIds {
		if ok := acctest.SweeperDefaultResourceId[triggerId]; !ok {
			deleteTriggerRequest := oci_devops.DeleteTriggerRequest{}

			deleteTriggerRequest.TriggerId = &triggerId

			deleteTriggerRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")
			_, error := devopsClient.DeleteTrigger(context.Background(), deleteTriggerRequest)
			if error != nil {
				fmt.Printf("Error deleting Trigger %s %s, It is possible that the resource is already deleted. Please verify manually \n", triggerId, error)
				continue
			}
		}
	}
	return nil
}

func getDevopsTriggerIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "TriggerId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	devopsClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()

	listTriggersRequest := oci_devops.ListTriggersRequest{}
	listTriggersRequest.CompartmentId = &compartmentId
	listTriggersResponse, err := devopsClient.ListTriggers(context.Background(), listTriggersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Trigger list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, trigger := range listTriggersResponse.Items {
		id := *trigger.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "TriggerId", id)
	}
	return resourceIds, nil
}
