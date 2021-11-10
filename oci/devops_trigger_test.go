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
	"github.com/oracle/oci-go-sdk/v51/common"
	oci_devops "github.com/oracle/oci-go-sdk/v51/devops"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	TriggerRequiredOnlyResource = TriggerResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", Required, Create, triggerRepresentation)

	TriggerResourceConfig = TriggerResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", Optional, Update, triggerRepresentation)

	triggerSingularDataSourceRepresentation = map[string]interface{}{
		"trigger_id": Representation{RepType: Required, Create: `${oci_devops_trigger.test_trigger.id}`},
	}

	triggerDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, triggerDataSourceFilterRepresentation}}
	triggerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_devops_trigger.test_trigger.id}`}},
	}

	triggerRepresentation = map[string]interface{}{
		"actions":        RepresentationGroup{Required, triggerActionsRepresentation},
		"project_id":     Representation{RepType: Required, Create: `${oci_devops_project.test_project.id}`},
		"trigger_source": Representation{RepType: Required, Create: `GITHUB`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}
	triggerActionsRepresentation = map[string]interface{}{
		"build_pipeline_id": Representation{RepType: Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"type":              Representation{RepType: Required, Create: `TRIGGER_BUILD_PIPELINE`},
		"filter":            RepresentationGroup{Optional, triggerActionsFilterRepresentation},
	}
	triggerActionsFilterRepresentation = map[string]interface{}{
		"trigger_source": Representation{RepType: Required, Create: `GITHUB`, Update: `GITHUB`},
		"events":         Representation{RepType: Optional, Create: []string{`PUSH`}, Update: []string{`PULL_REQUEST_CREATED`}},
		"include":        RepresentationGroup{Optional, triggerActionsFilterIncludeRepresentation},
	}
	triggerActionsFilterIncludeRepresentation = map[string]interface{}{
		"base_ref": Representation{RepType: Optional, Create: `baseRef`, Update: `baseRef2`},
		"head_ref": Representation{RepType: Optional, Create: `headRef`, Update: `headRef2`},
	}

	TriggerResourceDependencies = GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", Required, Create, buildPipelineRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsTriggerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsTriggerResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_trigger.test_trigger"
	datasourceName := "data.oci_devops_triggers.test_triggers"
	singularDatasourceName := "data.oci_devops_trigger.test_trigger"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+TriggerResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", Optional, Create, triggerRepresentation), "devops", "trigger", t)

	ResourceTest(t, testAccCheckDevopsTriggerDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + TriggerResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", Required, Create, triggerRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "actions.0.build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "TRIGGER_BUILD_PIPELINE"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "trigger_source", "GITHUB"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + TriggerResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + TriggerResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", Optional, Create, triggerRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "actions.0.build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.events.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.base_ref", "baseRef"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.head_ref", "headRef"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.trigger_source", "GITHUB"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "TRIGGER_BUILD_PIPELINE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "trigger_source", "GITHUB"),
				resource.TestCheckResourceAttrSet(resourceName, "trigger_url"),

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
			Config: config + compartmentIdVariableStr + TriggerResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", Optional, Update, triggerRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "actions.0.build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.events.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.base_ref", "baseRef2"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.head_ref", "headRef2"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.trigger_source", "GITHUB"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "TRIGGER_BUILD_PIPELINE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "trigger_source", "GITHUB"),
				resource.TestCheckResourceAttrSet(resourceName, "trigger_url"),

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
				GenerateDataSourceFromRepresentationMap("oci_devops_triggers", "test_triggers", Optional, Update, triggerDataSourceRepresentation) +
				compartmentIdVariableStr + TriggerResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", Optional, Update, triggerRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				resource.TestCheckResourceAttr(datasourceName, "trigger_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_trigger", "test_trigger", Required, Create, triggerSingularDataSourceRepresentation) +
				compartmentIdVariableStr + TriggerResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trigger_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.events.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.include.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.include.0.base_ref", "baseRef2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.include.0.head_ref", "headRef2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.trigger_source", "GITHUB"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.type", "TRIGGER_BUILD_PIPELINE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + TriggerResourceConfig,
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

func testAccCheckDevopsTriggerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).devopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_trigger" {
			noResourceFound = false
			request := oci_devops.GetTriggerRequest{}

			tmp := rs.Primary.ID
			request.TriggerId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("DevopsTrigger") {
		resource.AddTestSweepers("DevopsTrigger", &resource.Sweeper{
			Name:         "DevopsTrigger",
			Dependencies: DependencyGraph["trigger"],
			F:            sweepDevopsTriggerResource,
		})
	}
}

func sweepDevopsTriggerResource(compartment string) error {
	devopsClient := GetTestClients(&schema.ResourceData{}).devopsClient()
	triggerIds, err := getTriggerIds(compartment)
	if err != nil {
		return err
	}
	for _, triggerId := range triggerIds {
		if ok := SweeperDefaultResourceId[triggerId]; !ok {
			deleteTriggerRequest := oci_devops.DeleteTriggerRequest{}

			deleteTriggerRequest.TriggerId = &triggerId

			deleteTriggerRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")
			_, error := devopsClient.DeleteTrigger(context.Background(), deleteTriggerRequest)
			if error != nil {
				fmt.Printf("Error deleting Trigger %s %s, It is possible that the resource is already deleted. Please verify manually \n", triggerId, error)
				continue
			}
		}
	}
	return nil
}

func getTriggerIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "TriggerId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	devopsClient := GetTestClients(&schema.ResourceData{}).devopsClient()

	listTriggersRequest := oci_devops.ListTriggersRequest{}
	listTriggersRequest.CompartmentId = &compartmentId
	listTriggersResponse, err := devopsClient.ListTriggers(context.Background(), listTriggersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Trigger list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, trigger := range listTriggersResponse.Items {
		id := *trigger.GetId()
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "TriggerId", id)
	}
	return resourceIds, nil
}
