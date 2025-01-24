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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	StackMonitoringMonitoringTemplateRequiredOnlyResource = StackMonitoringMonitoringTemplateResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template", "test_monitoring_template", acctest.Required, acctest.Create, StackMonitoringMonitoringTemplateRepresentation)

	StackMonitoringMonitoringTemplateResourceConfig = StackMonitoringMonitoringTemplateResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template", "test_monitoring_template", acctest.Optional, acctest.Update, StackMonitoringMonitoringTemplateRepresentation)

	StackMonitoringMonitoringTemplateSingularDataSourceRepresentation = map[string]interface{}{
		"monitoring_template_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitoring_template.test_monitoring_template.id}`},
	}

	StackMonitoringMonitoringTemplateDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `MT_MonitoringTemplateTerraformExample`, Update: `MT_MonitoringTemplateTerraformExample_Updated`},
		"metric_name":            acctest.Representation{RepType: acctest.Optional, Create: []string{`CpuUtilization`}},
		"monitoring_template_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_stack_monitoring_monitoring_template.test_monitoring_template.id}`},
		"namespace":              acctest.Representation{RepType: acctest.Optional, Create: []string{`namespace`}},
		"resource_types":         acctest.Representation{RepType: acctest.Optional, Create: []string{`resourceTypes`}},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"status":                 acctest.Representation{RepType: acctest.Optional, Create: `NOT_APPLIED`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMonitoringTemplateDataSourceFilterRepresentation}}

	StackMonitoringMonitoringTemplateDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_monitoring_template.test_monitoring_template.id}`}},
	}

	destinationId = utils.GetEnvSettingWithBlankDefault("destination_ocid")

	StackMonitoringMonitoringTemplateRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `Test MT for resource type Apache Tomcat`, Update: `Test MT updated for resource type Apache Tomcat`},
		"destinations":                  acctest.Representation{RepType: acctest.Required, Create: []string{destinationId}},
		"display_name":                  acctest.Representation{RepType: acctest.Required, Create: `MT_MonitoringTemplateTerraformExample`, Update: `MT_MonitoringTemplateTerraformExample_Updated`},
		"members":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMonitoringTemplateMembersRepresentation},
		"is_alarms_enabled":             acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"is_split_notification_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"message_format":                acctest.Representation{RepType: acctest.Required, Create: `ONS_OPTIMIZED`, Update: `PRETTY_JSON`},
		"repeat_notification_duration":  acctest.Representation{RepType: acctest.Optional, Create: `PT2H`, Update: `PT3H`},
		"lifecycle":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMonitoringTemplateSensitiveDataRepresentation},
	}

	StackMonitoringMonitoringTemplateMembersRepresentation = map[string]interface{}{
		"id":   acctest.Representation{RepType: acctest.Required, Create: `ocid1.stackmonitoringresourcetype.apache_tomcat`, Update: `ocid1.stackmonitoringresourcetype.apache_tomcat`},
		"type": acctest.Representation{RepType: acctest.Required, Create: `RESOURCE_TYPE`, Update: `RESOURCE_TYPE`},
	}

	ignoreMonitoringTemplateSensitiveDataRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			`system_tags`}},
	}

	StackMonitoringMonitoringTemplateResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMonitoringTemplateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMonitoringTemplateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_stack_monitoring_monitoring_template.test_monitoring_template"
	datasourceName := "data.oci_stack_monitoring_monitoring_templates.test_monitoring_templates"
	singularDatasourceName := "data.oci_stack_monitoring_monitoring_template.test_monitoring_template"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+StackMonitoringMonitoringTemplateResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template", "test_monitoring_template", acctest.Optional, acctest.Create, StackMonitoringMonitoringTemplateRepresentation), "stackmonitoring", "monitoringTemplate", t)

	acctest.ResourceTest(t, testAccCheckStackMonitoringMonitoringTemplateDestroy, []resource.TestStep{
		// Verify Create
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMonitoringTemplateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template", "test_monitoring_template", acctest.Required, acctest.Create, StackMonitoringMonitoringTemplateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MT_MonitoringTemplateTerraformExample"),
				resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "members.0.id", "ocid1.stackmonitoringresourcetype.apache_tomcat"),
				resource.TestCheckResourceAttr(resourceName, "members.0.type", "RESOURCE_TYPE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					fmt.Println(err)
					return err
				},
			),
		},

		// Delete before next Create
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMonitoringTemplateResourceDependencies,
		},

		// Verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMonitoringTemplateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template", "test_monitoring_template", acctest.Optional, acctest.Create, StackMonitoringMonitoringTemplateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Test MT for resource type Apache Tomcat"),
				resource.TestCheckResourceAttr(resourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MT_MonitoringTemplateTerraformExample"),
				resource.TestCheckResourceAttr(resourceName, "is_alarms_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_split_notification_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "members.0.id", "ocid1.stackmonitoringresourcetype.apache_tomcat"),
				resource.TestCheckResourceAttr(resourceName, "members.0.type", "RESOURCE_TYPE"),
				resource.TestCheckResourceAttr(resourceName, "message_format", "ONS_OPTIMIZED"),
				resource.TestCheckResourceAttr(resourceName, "repeat_notification_duration", "PT2H"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "total_alarm_conditions"),
				resource.TestCheckResourceAttrSet(resourceName, "total_applied_alarm_conditions"),

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

		// Verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMonitoringTemplateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template", "test_monitoring_template", acctest.Optional, acctest.Update, StackMonitoringMonitoringTemplateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Test MT updated for resource type Apache Tomcat"),
				resource.TestCheckResourceAttr(resourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MT_MonitoringTemplateTerraformExample_Updated"),
				resource.TestCheckResourceAttr(resourceName, "is_alarms_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_split_notification_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "members.0.id", "ocid1.stackmonitoringresourcetype.apache_tomcat"),
				resource.TestCheckResourceAttr(resourceName, "members.0.type", "RESOURCE_TYPE"),
				resource.TestCheckResourceAttr(resourceName, "message_format", "PRETTY_JSON"),
				resource.TestCheckResourceAttr(resourceName, "repeat_notification_duration", "PT3H"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "total_alarm_conditions"),
				resource.TestCheckResourceAttrSet(resourceName, "total_applied_alarm_conditions"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// Verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitoring_templates", "test_monitoring_templates", acctest.Optional, acctest.Update, StackMonitoringMonitoringTemplateDataSourceRepresentation) +
				compartmentIdVariableStr + StackMonitoringMonitoringTemplateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template", "test_monitoring_template", acctest.Optional, acctest.Create, StackMonitoringMonitoringTemplateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "MT_MonitoringTemplateTerraformExample_Updated"),
				resource.TestCheckResourceAttr(datasourceName, "metric_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "namespace.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "resource_types.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "status", "NOT_APPLIED"),
				resource.TestCheckResourceAttrSet(datasourceName, "monitoring_template_id"),
			),
		},

		// Verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitoring_template", "test_monitoring_template", acctest.Required, acctest.Create, StackMonitoringMonitoringTemplateSingularDataSourceRepresentation) +
				compartmentIdVariableStr + StackMonitoringMonitoringTemplateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "Test MT updated for resource type Apache Tomcat"),
				resource.TestCheckResourceAttr(singularDatasourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "MT_MonitoringTemplateTerraformExample_Updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_alarms_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_split_notification_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "members.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "members.0.id", "ocid1.stackmonitoringresourcetype.apache_tomcat"),
				resource.TestCheckResourceAttr(singularDatasourceName, "members.0.type", "RESOURCE_TYPE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "message_format", "PRETTY_JSON"),
				resource.TestCheckResourceAttr(singularDatasourceName, "repeat_notification_duration", "PT3H"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitoring_template_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_alarm_conditions"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_applied_alarm_conditions"),
			),
		},

		// Verify resource import
		{
			Config:                  config + StackMonitoringMonitoringTemplateRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckStackMonitoringMonitoringTemplateDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).StackMonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_stack_monitoring_monitoring_template" {
			noResourceFound = false
			request := oci_stack_monitoring.GetMonitoringTemplateRequest{}

			tmp := rs.Primary.ID
			request.MonitoringTemplateId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")

			response, err := client.GetMonitoringTemplate(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_stack_monitoring.MonitoringTemplateLifeCycleStatesDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					// Resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				// Resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			// Verify that exception is for '404 not found'.
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
	if !acctest.InSweeperExcludeList("StackMonitoringMonitoringTemplate") {
		resource.AddTestSweepers("StackMonitoringMonitoringTemplate", &resource.Sweeper{
			Name:         "StackMonitoringMonitoringTemplate",
			Dependencies: acctest.DependencyGraph["monitoringTemplate"],
			F:            sweepStackMonitoringMonitoringTemplateResource,
		})
	}
}

func sweepStackMonitoringMonitoringTemplateResource(compartment string) error {
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()
	monitoringTemplateIds, err := getStackMonitoringMonitoringTemplateIds(compartment)
	if err != nil {
		return err
	}
	for _, monitoringTemplateId := range monitoringTemplateIds {
		if ok := acctest.SweeperDefaultResourceId[monitoringTemplateId]; !ok {
			deleteMonitoringTemplateRequest := oci_stack_monitoring.DeleteMonitoringTemplateRequest{}

			deleteMonitoringTemplateRequest.MonitoringTemplateId = &monitoringTemplateId

			deleteMonitoringTemplateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")
			_, error := stackMonitoringClient.DeleteMonitoringTemplate(context.Background(), deleteMonitoringTemplateRequest)
			if error != nil {
				fmt.Printf("Error deleting MonitoringTemplate %s %s, It is possible that the resource is already deleted. Please verify manually \n", monitoringTemplateId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &monitoringTemplateId, StackMonitoringMonitoringTemplateSweepWaitCondition, time.Duration(3*time.Minute),
				StackMonitoringMonitoringTemplateSweepResponseFetchOperation, "stack_monitoring", true)
		}
	}
	return nil
}

func getStackMonitoringMonitoringTemplateIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MonitoringTemplateId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()

	listMonitoringTemplatesRequest := oci_stack_monitoring.ListMonitoringTemplatesRequest{}
	listMonitoringTemplatesRequest.CompartmentId = &compartmentId
	listMonitoringTemplatesRequest.LifecycleState = oci_stack_monitoring.ListMonitoringTemplatesLifecycleStateActive
	listMonitoringTemplatesResponse, err := stackMonitoringClient.ListMonitoringTemplates(context.Background(), listMonitoringTemplatesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MonitoringTemplate list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, monitoringTemplate := range listMonitoringTemplatesResponse.Items {
		id := *monitoringTemplate.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MonitoringTemplateId", id)
	}
	return resourceIds, nil
}

func StackMonitoringMonitoringTemplateSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if monitoringTemplateResponse, ok := response.Response.(oci_stack_monitoring.GetMonitoringTemplateResponse); ok {
		return monitoringTemplateResponse.LifecycleState != oci_stack_monitoring.MonitoringTemplateLifeCycleStatesActive
	}
	return false
}

func StackMonitoringMonitoringTemplateSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.StackMonitoringClient().GetMonitoringTemplate(context.Background(), oci_stack_monitoring.GetMonitoringTemplateRequest{
		MonitoringTemplateId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
