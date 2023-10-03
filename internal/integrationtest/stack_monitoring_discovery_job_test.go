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

/*
*

	Dependency variables:
	  management_agent_id = var.stack_mon_management_agent_id_discovery
*/
var (
	StackMonitoringDiscoveryJobRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_discovery_job", "test_discovery_job", acctest.Required, acctest.Create, StackMonitoringDiscoveryJobRepresentation)

	StackMonitoringDiscoveryJobResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_discovery_job", "test_discovery_job", acctest.Optional, acctest.Update, StackMonitoringDiscoveryJobRepresentation)

	StackMonitoringStackMonitoringDiscoveryJobSingularDataSourceRepresentation = map[string]interface{}{
		"discovery_job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_discovery_job.test_discovery_job.id}`},
	}

	StackMonitoringStackMonitoringDiscoveryJobDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `terraformDiscoveryJob`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringDiscoveryJobDataSourceFilterRepresentation}}
	StackMonitoringDiscoveryJobDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_discovery_job.test_discovery_job.id}`}},
	}

	StackMonitoringDiscoveryJobRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"discovery_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringDiscoveryJobDiscoveryDetailsRepresentation},
		"discovery_client":  acctest.Representation{RepType: acctest.Optional, Create: `LA_SERVICE`},
		"discovery_type":    acctest.Representation{RepType: acctest.Optional, Create: `ADD`},
		"should_propagate_tags_to_discovered_resources": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSensitiveDiscoveryJobDataRepresentation},
	}
	StackMonitoringDiscoveryJobDiscoveryDetailsRepresentation = map[string]interface{}{
		"agent_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.stack_mon_management_agent_id_discovery}`},
		"properties":    acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringDiscoveryJobDiscoveryDetailsPropertiesRepresentation},
		"resource_name": acctest.Representation{RepType: acctest.Required, Create: `terraformDiscoveryJob`},
		"resource_type": acctest.Representation{RepType: acctest.Required, Create: `WEBLOGIC_DOMAIN`},
		"credentials":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringDiscoveryJobDiscoveryDetailsCredentialsRepresentation},
		"license":       acctest.Representation{RepType: acctest.Optional, Create: `STANDARD_EDITION`},
		"tags":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringDiscoveryJobDiscoveryDetailsTagsRepresentation},
	}
	StackMonitoringDiscoveryJobDiscoveryDetailsPropertiesRepresentation = map[string]interface{}{
		"properties_map": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"admin_server_host": "somehost.us.oracle.com",
			"admin_server_port":     "7001",
			"admin_server_protocol": "t3"}},
	}
	//Get API does not return sensitive data, it returns null
	ignoreSensitiveDiscoveryJobDataRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`discovery_details`, `system_tags`, `defined_tags`}},
	}
	StackMonitoringDiscoveryJobDiscoveryDetailsCredentialsRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringDiscoveryJobDiscoveryDetailsCredentialsItemsRepresentation},
	}
	StackMonitoringDiscoveryJobDiscoveryDetailsTagsRepresentation = map[string]interface{}{
		"properties_map": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"propertiesMap": "propertiesMap"}},
	}
	StackMonitoringDiscoveryJobDiscoveryDetailsCredentialsItemsRepresentation = map[string]interface{}{
		"credential_name": acctest.Representation{RepType: acctest.Required, Create: `Sk1YQ3JlZHM=`},
		"credential_type": acctest.Representation{RepType: acctest.Required, Create: `Sk1YQ3JlZHM=`},
		"properties":      acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringDiscoveryJobDiscoveryDetailsCredentialsItemsPropertiesRepresentation},
	}
	StackMonitoringDiscoveryJobDiscoveryDetailsCredentialsItemsPropertiesRepresentation = map[string]interface{}{
		"properties_map": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Username": "d2VibG9naWM=",
			"Password": "d2VibG9naWM="}},
	}
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringDiscoveryJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringDiscoveryJobResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentId := utils.GetEnvSettingWithBlankDefault("stack_mon_management_agent_id_discovery")
	if managementAgentId == "" {
		t.Skip("Setting environmental variable stack_mon_management_agent_id_discovery that represents management agent capable of running stack monitoring discovery is pre-requisite for this test")
	}
	managementAgentIdVariableStr := fmt.Sprintf("variable \"stack_mon_management_agent_id_discovery\" { default = \"%s\" }\n", managementAgentId)

	resourceName := "oci_stack_monitoring_discovery_job.test_discovery_job"
	datasourceName := "data.oci_stack_monitoring_discovery_jobs.test_discovery_jobs"
	singularDatasourceName := "data.oci_stack_monitoring_discovery_job.test_discovery_job"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+managementAgentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_discovery_job", "test_discovery_job", acctest.Optional, acctest.Create, StackMonitoringDiscoveryJobRepresentation), "stackmonitoring", "discoveryJob", t)

	acctest.ResourceTest(t, testAccCheckStackMonitoringDiscoveryJobDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_discovery_job", "test_discovery_job", acctest.Optional, acctest.Create, StackMonitoringDiscoveryJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "discovery_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "discovery_details.0.agent_id", managementAgentId),
				resource.TestCheckResourceAttr(resourceName, "discovery_details.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "discovery_details.0.resource_name", "terraformDiscoveryJob"),
				resource.TestCheckResourceAttr(resourceName, "discovery_details.0.resource_type", "WEBLOGIC_DOMAIN"),
				resource.TestCheckResourceAttr(resourceName, "should_propagate_tags_to_discovered_resources", "false"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_discovery_job", "test_discovery_job", acctest.Optional, acctest.Create, StackMonitoringDiscoveryJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "discovery_client", "LA_SERVICE"),
				resource.TestCheckResourceAttr(resourceName, "discovery_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "discovery_details.0.license", "STANDARD_EDITION"),
				resource.TestCheckResourceAttr(resourceName, "discovery_details.0.agent_id", managementAgentId),
				resource.TestCheckResourceAttr(resourceName, "discovery_details.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "discovery_details.0.properties.0.properties_map.%", "3"),
				resource.TestCheckResourceAttr(resourceName, "discovery_details.0.resource_name", "terraformDiscoveryJob"),
				resource.TestCheckResourceAttr(resourceName, "discovery_details.0.resource_type", "WEBLOGIC_DOMAIN"),
				resource.TestCheckResourceAttr(resourceName, "discovery_type", "ADD"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "should_propagate_tags_to_discovered_resources", "false"),

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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_discovery_jobs", "test_discovery_jobs", acctest.Optional, acctest.Update, StackMonitoringStackMonitoringDiscoveryJobDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_discovery_job", "test_discovery_job", acctest.Optional, acctest.Update, StackMonitoringDiscoveryJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "terraformDiscoveryJob"),

				resource.TestCheckResourceAttr(datasourceName, "discovery_job_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "discovery_job_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_discovery_job", "test_discovery_job", acctest.Required, acctest.Create, StackMonitoringStackMonitoringDiscoveryJobSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr + StackMonitoringDiscoveryJobResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "discovery_job_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "discovery_client", "LA_SERVICE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "discovery_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "discovery_details.0.properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "discovery_details.0.properties.0.properties_map.%", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "discovery_details.0.resource_name", "terraformDiscoveryJob"),
				resource.TestCheckResourceAttr(singularDatasourceName, "discovery_details.0.resource_type", "WEBLOGIC_DOMAIN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "discovery_details.0.tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "discovery_details.0.tags.0.properties_map.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "discovery_details.0.license", "STANDARD_EDITION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "discovery_type", "ADD"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status_message"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_id"),
			),
		},

		// remove datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + StackMonitoringDiscoveryJobResourceConfig,
		},
		// verify resource import
		{
			Config:            config + compartmentIdVariableStr + managementAgentIdVariableStr + StackMonitoringDiscoveryJobResourceConfig,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"should_propagate_tags_to_discovered_resources",
				"defined_tags",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckStackMonitoringDiscoveryJobDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).StackMonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_stack_monitoring_discovery_job" {
			noResourceFound = false
			request := oci_stack_monitoring.GetDiscoveryJobRequest{}

			tmp := rs.Primary.ID
			request.DiscoveryJobId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")

			response, err := client.GetDiscoveryJob(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_stack_monitoring.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("StackMonitoringDiscoveryJob") {
		resource.AddTestSweepers("StackMonitoringDiscoveryJob", &resource.Sweeper{
			Name:         "StackMonitoringDiscoveryJob",
			Dependencies: acctest.DependencyGraph["discoveryJob"],
			F:            sweepStackMonitoringDiscoveryJobResource,
		})
	}
}

func sweepStackMonitoringDiscoveryJobResource(compartment string) error {
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()
	discoveryJobIds, err := getStackMonitoringDiscoveryJobIds(compartment)
	if err != nil {
		return err
	}
	for _, discoveryJobId := range discoveryJobIds {
		if ok := acctest.SweeperDefaultResourceId[discoveryJobId]; !ok {
			deleteDiscoveryJobRequest := oci_stack_monitoring.DeleteDiscoveryJobRequest{}

			deleteDiscoveryJobRequest.DiscoveryJobId = &discoveryJobId

			deleteDiscoveryJobRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")
			_, error := stackMonitoringClient.DeleteDiscoveryJob(context.Background(), deleteDiscoveryJobRequest)
			if error != nil {
				fmt.Printf("Error deleting DiscoveryJob %s %s, It is possible that the resource is already deleted. Please verify manually \n", discoveryJobId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &discoveryJobId, StackMonitoringDiscoveryJobSweepWaitCondition, time.Duration(3*time.Minute),
				StackMonitoringDiscoveryJobSweepResponseFetchOperation, "stack_monitoring", true)
		}
	}
	return nil
}

func getStackMonitoringDiscoveryJobIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DiscoveryJobId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()

	listDiscoveryJobsRequest := oci_stack_monitoring.ListDiscoveryJobsRequest{}
	listDiscoveryJobsRequest.CompartmentId = &compartmentId
	listDiscoveryJobsResponse, err := stackMonitoringClient.ListDiscoveryJobs(context.Background(), listDiscoveryJobsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DiscoveryJob list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, discoveryJob := range listDiscoveryJobsResponse.Items {
		id := *discoveryJob.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DiscoveryJobId", id)
	}
	return resourceIds, nil
}

func StackMonitoringDiscoveryJobSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if discoveryJobResponse, ok := response.Response.(oci_stack_monitoring.GetDiscoveryJobResponse); ok {
		return discoveryJobResponse.LifecycleState != oci_stack_monitoring.LifecycleStateDeleted
	}
	return false
}

func StackMonitoringDiscoveryJobSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.StackMonitoringClient().GetDiscoveryJob(context.Background(), oci_stack_monitoring.GetDiscoveryJobRequest{
		DiscoveryJobId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
