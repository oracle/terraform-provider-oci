// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	StackMonitoringMonitoredResourceRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Required, acctest.Create, StackMonitoringMonitoredResourceRepresentation)

	StackMonitoringMonitoredResourceResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Optional, acctest.Update, StackMonitoringMonitoredResourceRepresentation)

	StackMonitoringMonitoredResourceSingularDataSourceRepresentation = map[string]interface{}{
		"monitored_resource_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitored_resource.test_monitored_resource.id}`},
	}

	StackMonitoringMonitoredResourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `terraformResource`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMonitoredResourceDataSourceFilterRepresentation}}
	StackMonitoringMonitoredResourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_monitored_resource.test_monitored_resource.id}`}},
	}

	StackMonitoringMonitoredResourceOSCreateProperty1 = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `osName`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `Linux`},
	}

	StackMonitoringMonitoredResourceOSCreateProperty2 = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `osVersion`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `7.0`},
	}

	StackMonitoringMonitoredResourceRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                   acctest.Representation{RepType: acctest.Required, Create: `terraformResource`},
		"type":                   acctest.Representation{RepType: acctest.Required, Create: `host`},
		"additional_aliases":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringMonitoredResourceAdditionalAliasesCredentialRepresentation},
		"additional_credentials": acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringMonitoredResourceAdditionalCredentialsRepresentation},
		"aliases":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringMonitoredResourceAliasesRepresentation},
		"credentials":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringMonitoredResourceCredentialsRepresentation},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayNameTerra`, Update: `displayNameTerra2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"host_name":              acctest.Representation{RepType: acctest.Optional, Create: `${var.stack_mon_hostname_resource1}`},
		"license":                acctest.Representation{RepType: acctest.Optional, Create: `STANDARD_EDITION`, Update: `ENTERPRISE_EDITION`},
		"management_agent_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.stack_mon_management_agent_id_resource1}`},
		"properties":             []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: StackMonitoringMonitoredResourceOSCreateProperty1}, {RepType: acctest.Optional, Group: StackMonitoringMonitoredResourceOSCreateProperty2}},
		"resource_time_zone":     acctest.Representation{RepType: acctest.Optional, Create: `en`},
		"lifecycle":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSensitiveDataRepresentation},
	}

	//Get API does not return sensitive data, it returns null
	ignoreSensitiveDataRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`credentials`, `aliases`, `properties`, `external_id`, `defined_tags`}},
	}

	StackMonitoringMonitoredResourceRepresentation2 = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `terraformSecondaryResource`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `host`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displaySecondaryNameTerra`, Update: `displaySecondaryNameTerra2`},
		"host_name":           acctest.Representation{RepType: acctest.Optional, Create: `${var.stack_mon_hostname_resource2}`},
		"management_agent_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.stack_mon_management_agent_id_resource2}`},
		"properties":          []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: StackMonitoringMonitoredResourceOSCreateProperty1}, {RepType: acctest.Optional, Group: StackMonitoringMonitoredResourceOSCreateProperty2}},
		"resource_time_zone":  acctest.Representation{RepType: acctest.Optional, Create: `en`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSensitiveDataRepresentation},
	}

	StackMonitoringMonitoredResourceAliasesRepresentation = map[string]interface{}{
		"credential": acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMonitoredResourceAliasesCredentialRepresentation},
		"name":       acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"source":     acctest.Representation{RepType: acctest.Required, Create: `host.terraformResource`, Update: `host.terraformResource`},
	}
	StackMonitoringMonitoredResourceCredentialsRepresentation = map[string]interface{}{
		"credential_type": acctest.Representation{RepType: acctest.Optional, Create: `PLAINTEXT`, Update: `PLAINTEXT`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"key_id":          acctest.Representation{RepType: acctest.Optional, Create: `somekeyid`},
		"name":            acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"properties":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringMonitoredResourceCredentialsPropertiesRepresentation},
		"source":          acctest.Representation{RepType: acctest.Optional, Create: `host.terraformResource`},
		"type":            acctest.Representation{RepType: acctest.Optional, Create: `type`, Update: `type2`},
	}

	StackMonitoringMonitoredResourceAdditionalCredentialsRepresentation = map[string]interface{}{
		"credential_type": acctest.Representation{RepType: acctest.Optional, Create: `PLAINTEXT`, Update: `PLAINTEXT`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description3`, Update: `description4`},
		"key_id":          acctest.Representation{RepType: acctest.Optional, Create: `somekeyid`},
		"name":            acctest.Representation{RepType: acctest.Optional, Create: `name3`, Update: `name4`},
		"properties":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringMonitoredResourceAdditionalCredentialsPropertiesRepresentation},
		"source":          acctest.Representation{RepType: acctest.Optional, Create: `host.terraformResource`},
		"type":            acctest.Representation{RepType: acctest.Optional, Create: `type3`, Update: `type4`},
	}

	StackMonitoringMonitoredResourceAdditionalAliasesCredentialRepresentation = map[string]interface{}{
		"credential": acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMonitoredResourceAliasesCredentialRepresentation},
		"name":       acctest.Representation{RepType: acctest.Required, Create: `credAliasName`},
		"source":     acctest.Representation{RepType: acctest.Required, Create: `host.terraformResource`},
	}

	StackMonitoringMonitoredResourceAdditionalCredentialsPropertiesRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `JMXUserName`, Update: `name3`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `WebLogic`, Update: `value3`},
	}
	StackMonitoringMonitoredResourceAliasesCredentialRepresentation = map[string]interface{}{
		"name":    acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"service": acctest.Representation{RepType: acctest.Required, Create: `service`, Update: `service2`},
		"source":  acctest.Representation{RepType: acctest.Required, Create: `host.terraformResource`, Update: `host.terraformResource`},
	}
	StackMonitoringMonitoredResourceCredentialsPropertiesRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMonitoredResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMonitoredResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	managementAgentId1 := utils.GetEnvSettingWithBlankDefault("stack_mon_management_agent_id_resource1")
	if managementAgentId1 == "" {
		t.Skip("Setting environmental variable stack_mon_management_agent_id_resource1 that represents management agent with resource monitoring plugin is pre-requisite for this test")
	}
	managementAgentId1VariableStr := fmt.Sprintf("variable \"stack_mon_management_agent_id_resource1\" { default = \"%s\" }\n", managementAgentId1)

	hostname1 := utils.GetEnvSettingWithBlankDefault("stack_mon_hostname_resource1")
	if hostname1 == "" {
		t.Skip("Setting environmental variable stack_mon_hostname_resource1 that host accessible by agent defined by stack_mon_management_agent_id_resource1 variable is pre-requisite for this test")
	}
	hostname1VariableStr := fmt.Sprintf("variable \"stack_mon_hostname_resource1\" { default = \"%s\" }\n", hostname1)

	resourceName := "oci_stack_monitoring_monitored_resource.test_monitored_resource"
	datasourceName := "data.oci_stack_monitoring_monitored_resources.test_monitored_resources"
	singularDatasourceName := "data.oci_stack_monitoring_monitored_resource.test_monitored_resource"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+managementAgentId1VariableStr+hostname1VariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Optional, acctest.Create, StackMonitoringMonitoredResourceRepresentation), "stackmonitoring", "monitoredResource", t)

	acctest.ResourceTest(t, testAccCheckStackMonitoringMonitoredResourceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Optional, acctest.Create, StackMonitoringMonitoredResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "terraformResource"),
				resource.TestCheckResourceAttr(resourceName, "type", "host"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Optional, acctest.Create, StackMonitoringMonitoredResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameTerra"),
				resource.TestCheckResourceAttrSet(resourceName, "host_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license", "STANDARD_EDITION"),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "terraformResource"),
				resource.TestCheckResourceAttr(resourceName, "properties.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "resource_time_zone", "en"),
				resource.TestCheckResourceAttr(resourceName, "type", "host"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + managementAgentId1VariableStr + hostname1VariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(StackMonitoringMonitoredResourceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameTerra"),
				resource.TestCheckResourceAttrSet(resourceName, "host_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license", "STANDARD_EDITION"),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "properties.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "resource_time_zone", "en"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "host"),

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
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Optional, acctest.Update, StackMonitoringMonitoredResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameTerra2"),
				resource.TestCheckResourceAttrSet(resourceName, "host_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "terraformResource"),
				resource.TestCheckResourceAttr(resourceName, "properties.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "resource_time_zone", "en"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "host"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitored_resources", "test_monitored_resources", acctest.Optional, acctest.Update, StackMonitoringMonitoredResourceDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Optional, acctest.Update, StackMonitoringMonitoredResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "terraformResource"),

				resource.TestCheckResourceAttr(datasourceName, "monitored_resource_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "monitored_resource_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Required, acctest.Create, StackMonitoringMonitoredResourceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + StackMonitoringMonitoredResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitored_resource_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayNameTerra2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "terraformResource"),
				resource.TestCheckResourceAttr(singularDatasourceName, "properties.#", "4"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "properties.0.name", "osName"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "properties.0.value", "Linux"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_time_zone", "en"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "host"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + StackMonitoringMonitoredResourceResourceConfig,
		},
		// verify resource import
		{
			Config:            config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + StackMonitoringMonitoredResourceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"additional_aliases",
				"additional_credentials",
				"external_resource_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckStackMonitoringMonitoredResourceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).StackMonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_stack_monitoring_monitored_resource" {
			noResourceFound = false
			request := oci_stack_monitoring.GetMonitoredResourceRequest{}

			tmp := rs.Primary.ID
			request.MonitoredResourceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")

			response, err := client.GetMonitoredResource(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_stack_monitoring.ResourceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("StackMonitoringMonitoredResource") {
		resource.AddTestSweepers("StackMonitoringMonitoredResource", &resource.Sweeper{
			Name:         "StackMonitoringMonitoredResource",
			Dependencies: acctest.DependencyGraph["monitoredResource"],
			F:            sweepStackMonitoringMonitoredResourceResource,
		})
	}
}

func sweepStackMonitoringMonitoredResourceResource(compartment string) error {
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()
	monitoredResourceIds, err := getStackMonitoringMonitoredResourceIds(compartment)
	if err != nil {
		return err
	}
	for _, monitoredResourceId := range monitoredResourceIds {
		if ok := acctest.SweeperDefaultResourceId[monitoredResourceId]; !ok {
			deleteMonitoredResourceRequest := oci_stack_monitoring.DeleteMonitoredResourceRequest{}

			deleteMonitoredResourceRequest.MonitoredResourceId = &monitoredResourceId

			deleteMonitoredResourceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")
			_, error := stackMonitoringClient.DeleteMonitoredResource(context.Background(), deleteMonitoredResourceRequest)
			if error != nil {
				fmt.Printf("Error deleting MonitoredResource %s %s, It is possible that the resource is already deleted. Please verify manually \n", monitoredResourceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &monitoredResourceId, StackMonitoringMonitoredResourceSweepWaitCondition, time.Duration(3*time.Minute),
				StackMonitoringMonitoredResourceSweepResponseFetchOperation, "stack_monitoring", true)
		}
	}
	return nil
}

func getStackMonitoringMonitoredResourceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MonitoredResourceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()

	listMonitoredResourcesRequest := oci_stack_monitoring.ListMonitoredResourcesRequest{}
	listMonitoredResourcesRequest.CompartmentId = &compartmentId
	listMonitoredResourcesResponse, err := stackMonitoringClient.ListMonitoredResources(context.Background(), listMonitoredResourcesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MonitoredResource list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, monitoredResource := range listMonitoredResourcesResponse.Items {
		id := *monitoredResource.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MonitoredResourceId", id)
	}
	return resourceIds, nil
}

func StackMonitoringMonitoredResourceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if monitoredResourceResponse, ok := response.Response.(oci_stack_monitoring.GetMonitoredResourceResponse); ok {
		return monitoredResourceResponse.LifecycleState != oci_stack_monitoring.ResourceLifecycleStateDeleted
	}
	return false
}

func StackMonitoringMonitoredResourceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.StackMonitoringClient().GetMonitoredResource(context.Background(), oci_stack_monitoring.GetMonitoredResourceRequest{
		MonitoredResourceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
