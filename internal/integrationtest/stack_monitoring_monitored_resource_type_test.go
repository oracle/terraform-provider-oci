// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
	StackMonitoringMonitoredResourceTypeRequiredOnlyResource = StackMonitoringMonitoredResourceTypeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_type", "test_monitored_resource_type", acctest.Required, acctest.Create, StackMonitoringMonitoredResourceTypeRepresentation)

	StackMonitoringMonitoredResourceTypeResourceConfig = StackMonitoringMonitoredResourceTypeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_type", "test_monitored_resource_type", acctest.Optional, acctest.Update, StackMonitoringMonitoredResourceTypeRepresentation)

	StackMonitoringMonitoredResourceTypeSingularDataSourceRepresentation = map[string]interface{}{
		"monitored_resource_type_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitored_resource_type.test_monitored_resource_type.id}`},
	}

	StackMonitoringMonitoredResourceTypeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"exclude_fields":          acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName`}},
		"fields":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`metricNamespace`}},
		"is_exclude_system_types": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"metric_namespace":        acctest.Representation{RepType: acctest.Optional, Create: `terraform_namespace`},
		"name":                    acctest.Representation{RepType: acctest.Optional, Create: `terraform_test_restype`},
		"status":                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMonitoredResourceTypeDataSourceFilterRepresentation}}
	StackMonitoringMonitoredResourceTypeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_monitored_resource_type.test_monitored_resource_type.id}`}},
	}

	//Get API does not return sensitive data, it returns null
	ignoreResourceTypeSensitiveDataRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			`freeform_tags`, `defined_tags`, `system_tags`,
			`compartment_id`, `display_name`, `metadata[0].valid_sub_resource_types`,
			`metadata[0].unique_property_sets`, `metadata[0].unique_property_sets[0].properties`}},
	}

	//Create uses this
	StackMonitoringMonitoredResourceTypeRepresentation = map[string]interface{}{
		"resource_category": acctest.Representation{RepType: acctest.Optional, Create: `APPLICATION`},
		"source_type":       acctest.Representation{RepType: acctest.Optional, Create: `SM_MGMT_AGENT_MONITORED`},
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":              acctest.Representation{RepType: acctest.Required, Create: `terraform_test_restype`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `Created for terraform testing.`, Update: `description2`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `Terraform Resource Type`, Update: `displayName2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"metadata":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringMonitoredResourceTypeMetadataRepresentation},
		"metric_namespace":  acctest.Representation{RepType: acctest.Optional, Create: `terraform_namespace`},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreResourceTypeSensitiveDataRepresentation},
	}

	StackMonitoringMonitoredResourceTypeMetadataRepresentation = map[string]interface{}{
		"format":                      acctest.Representation{RepType: acctest.Required, Create: `SYSTEM_FORMAT`},
		"agent_properties":            acctest.Representation{RepType: acctest.Optional, Create: []string{`agentProperties`}, Update: []string{`agentProperties2`}},
		"required_properties":         acctest.Representation{RepType: acctest.Optional, Create: []string{`requiredProperties`}, Update: []string{`requiredProperties2`}},
		"unique_property_sets":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringMonitoredResourceTypeMetadataUniquePropertySetsRepresentation},
		"valid_properties_for_create": acctest.Representation{RepType: acctest.Optional, Create: []string{`validPropertiesForCreate`}, Update: []string{`validPropertiesForCreate2`}},
		"valid_properties_for_update": acctest.Representation{RepType: acctest.Optional, Create: []string{`validPropertiesForUpdate`}, Update: []string{`validPropertiesForUpdate2`}},
		"valid_sub_resource_types":    acctest.Representation{RepType: acctest.Optional, Create: []string{`validSubResourceTypes`}, Update: []string{`validSubResourceTypes2`}},
		"valid_property_values":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"osType": "Linux,Windows,Solaris", "osVersion": "v6.0,v7.0"}, Update: map[string]string{"prop1": "Linux,Windows,Solaris", "osVersion": "v6.0,v7.0,v8.0"}},
	}
	StackMonitoringMonitoredResourceTypeMetadataUniquePropertySetsRepresentation = map[string]interface{}{
		"properties": acctest.Representation{RepType: acctest.Required, Create: []string{`properties`}, Update: []string{`properties2`}},
	}

	StackMonitoringMonitoredResourceTypeResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMonitoredResourceTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMonitoredResourceTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	tenantId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	tenantIdVariableStr := fmt.Sprintf("variable \"tenant_id\" { default = \"%s\" }\n", tenantId)

	resourceName := "oci_stack_monitoring_monitored_resource_type.test_monitored_resource_type"
	datasourceName := "data.oci_stack_monitoring_monitored_resource_types.test_monitored_resource_types"
	singularDatasourceName := "data.oci_stack_monitoring_monitored_resource_type.test_monitored_resource_type"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+tenantIdVariableStr+StackMonitoringMonitoredResourceTypeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_type", "test_monitored_resource_type", acctest.Optional, acctest.Create, StackMonitoringMonitoredResourceTypeRepresentation), "stackmonitoring", "monitoredResourceType", t)

	acctest.ResourceTest(t, testAccCheckStackMonitoringMonitoredResourceTypeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + tenantIdVariableStr + StackMonitoringMonitoredResourceTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_type", "test_monitored_resource_type", acctest.Required, acctest.Create, StackMonitoringMonitoredResourceTypeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "terraform_test_restype"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + tenantIdVariableStr + StackMonitoringMonitoredResourceTypeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + tenantIdVariableStr + StackMonitoringMonitoredResourceTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_type", "test_monitored_resource_type", acctest.Optional, acctest.Create, StackMonitoringMonitoredResourceTypeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "terraform_test_restype"),
				resource.TestCheckResourceAttr(resourceName, "description", "Created for terraform testing."),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform Resource Type"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.agent_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.format", "SYSTEM_FORMAT"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.required_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.unique_property_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.unique_property_sets.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.valid_properties_for_create.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.valid_properties_for_update.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metric_namespace", "terraform_namespace"),
				resource.TestCheckResourceAttr(resourceName, "resource_category", "APPLICATION"),
				resource.TestCheckResourceAttr(resourceName, "source_type", "SM_MGMT_AGENT_MONITORED"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.valid_property_values.%", "2"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + tenantIdVariableStr + StackMonitoringMonitoredResourceTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_type", "test_monitored_resource_type", acctest.Optional, acctest.Update, StackMonitoringMonitoredResourceTypeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "terraform_test_restype"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform Resource Type"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.agent_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.format", "SYSTEM_FORMAT"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.required_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.unique_property_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.unique_property_sets.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.valid_properties_for_create.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.valid_properties_for_update.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metric_namespace", "terraform_namespace"),
				resource.TestCheckResourceAttr(resourceName, "resource_category", "APPLICATION"),
				resource.TestCheckResourceAttr(resourceName, "source_type", "SM_MGMT_AGENT_MONITORED"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.valid_property_values.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "metric_namespace", "terraform_namespace"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_types", "test_monitored_resource_types", acctest.Optional, acctest.Update, StackMonitoringMonitoredResourceTypeDataSourceRepresentation) +
				compartmentIdVariableStr + tenantIdVariableStr + StackMonitoringMonitoredResourceTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_type", "test_monitored_resource_type", acctest.Optional, acctest.Update, StackMonitoringMonitoredResourceTypeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "exclude_fields.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fields.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "is_exclude_system_types", "false"),
				resource.TestCheckResourceAttr(datasourceName, "metric_namespace", "terraform_namespace"),
				resource.TestCheckResourceAttr(datasourceName, "name", "terraform_test_restype"),
				resource.TestCheckResourceAttr(datasourceName, "status", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "monitored_resource_types_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "monitored_resource_types_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_type", "test_monitored_resource_type", acctest.Required, acctest.Create, StackMonitoringMonitoredResourceTypeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + tenantIdVariableStr + StackMonitoringMonitoredResourceTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitored_resource_type_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Terraform Resource Type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_metrics_config.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "handler_config.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_system_defined"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.0.agent_properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.0.format", "SYSTEM_FORMAT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.0.required_properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.0.unique_property_sets.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.0.unique_property_sets.0.properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.0.valid_properties_for_create.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.0.valid_properties_for_update.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_category", "APPLICATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_type", "SM_MGMT_AGENT_MONITORED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.0.valid_property_values.%", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_namespace", "terraform_namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "terraform_test_restype"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + tenantIdVariableStr + StackMonitoringMonitoredResourceTypeRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckStackMonitoringMonitoredResourceTypeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).StackMonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_stack_monitoring_monitored_resource_type" {
			noResourceFound = false
			request := oci_stack_monitoring.GetMonitoredResourceTypeRequest{}

			tmp := rs.Primary.ID
			request.MonitoredResourceTypeId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")

			response, err := client.GetMonitoredResourceType(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_stack_monitoring.ResourceTypeLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("StackMonitoringMonitoredResourceType") {
		resource.AddTestSweepers("StackMonitoringMonitoredResourceType", &resource.Sweeper{
			Name:         "StackMonitoringMonitoredResourceType",
			Dependencies: acctest.DependencyGraph["monitoredResourceType"],
			F:            sweepStackMonitoringMonitoredResourceTypeResource,
		})
	}
}

func sweepStackMonitoringMonitoredResourceTypeResource(compartment string) error {
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()
	monitoredResourceTypeIds, err := getStackMonitoringMonitoredResourceTypeIds(compartment)
	if err != nil {
		return err
	}
	for _, monitoredResourceTypeId := range monitoredResourceTypeIds {
		if ok := acctest.SweeperDefaultResourceId[monitoredResourceTypeId]; !ok {
			deleteMonitoredResourceTypeRequest := oci_stack_monitoring.DeleteMonitoredResourceTypeRequest{}

			deleteMonitoredResourceTypeRequest.MonitoredResourceTypeId = &monitoredResourceTypeId

			deleteMonitoredResourceTypeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")
			_, error := stackMonitoringClient.DeleteMonitoredResourceType(context.Background(), deleteMonitoredResourceTypeRequest)
			if error != nil {
				fmt.Printf("Error deleting MonitoredResourceType %s %s, It is possible that the resource is already deleted. Please verify manually \n", monitoredResourceTypeId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &monitoredResourceTypeId, StackMonitoringMonitoredResourceTypeSweepWaitCondition, time.Duration(3*time.Minute),
				StackMonitoringMonitoredResourceTypeSweepResponseFetchOperation, "stack_monitoring", true)
		}
	}
	return nil
}

func getStackMonitoringMonitoredResourceTypeIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MonitoredResourceTypeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()

	listMonitoredResourceTypesRequest := oci_stack_monitoring.ListMonitoredResourceTypesRequest{}
	listMonitoredResourceTypesRequest.CompartmentId = &compartmentId
	listMonitoredResourceTypesResponse, err := stackMonitoringClient.ListMonitoredResourceTypes(context.Background(), listMonitoredResourceTypesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MonitoredResourceType list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, monitoredResourceType := range listMonitoredResourceTypesResponse.Items {
		id := *monitoredResourceType.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MonitoredResourceTypeId", id)
	}
	return resourceIds, nil
}

func StackMonitoringMonitoredResourceTypeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if monitoredResourceTypeResponse, ok := response.Response.(oci_stack_monitoring.GetMonitoredResourceTypeResponse); ok {
		return monitoredResourceTypeResponse.LifecycleState != oci_stack_monitoring.ResourceTypeLifecycleStateDeleted
	}
	return false
}

func StackMonitoringMonitoredResourceTypeSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.StackMonitoringClient().GetMonitoredResourceType(context.Background(), oci_stack_monitoring.GetMonitoredResourceTypeRequest{
		MonitoredResourceTypeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
