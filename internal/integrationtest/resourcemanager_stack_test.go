// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v65/resourcemanager"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/service/resourcemanager"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ResourcemanagerResourcemanagerStackSingularDataSourceRepresentation = map[string]interface{}{
		"stack_id": acctest.Representation{RepType: acctest.Required, Create: `${var.resource_manager_stack_id}`},
	}

	ResourcemanagerResourcemanagerStackDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `TestResourcemanagerStackResource_basic`, Update: `TestResourcemanagerStackResource_basic`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_resourcemanager_stack.test_stack.id}`},
		"state":          acctest.Representation{RepType: acctest.Required, Create: `ACTIVE`}, // make `required` here so it can be asserted against in step 0
	}

	ResourcemanagerStackResourceConfig = DefinedTagsDependencies
)

// issue-routing-tag: resourcemanager/default
func TestResourcemanagerStackResource_basic(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestResourcemanagerStackResource_basic") {
		t.Skip("Skipping suppressed TestResourcemanagerStackResource_basic")
	}

	httpreplay.SetScenario("TestResourcemanagerStackResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	client := acctest.GetTestClients(&schema.ResourceData{}).ResourceManagerClient()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceManagerStackId, err := resourcemanager.CreateResourceManagerStack(*client, "TestResourcemanagerStackResource_basic", compartmentId)
	if err != nil {
		t.Errorf("cannot Create resource manager stack for the test run: %v", err)
	}

	datasourceName := "data.oci_resourcemanager_stacks.test_stacks"
	singularDatasourceName := "data.oci_resourcemanager_stack.test_stack"

	acctest.SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		CheckDestroy: func(s *terraform.State) error {
			return resourcemanager.DestroyResourceManagerStack(*client, resourceManagerStackId)
		},
		PreventPostDestroyRefresh: true,
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `
					variable "resource_manager_stack_id" { default = "` + resourceManagerStackId + `" }
					` +
					acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stacks", "test_stacks", acctest.Required, acctest.Create, ResourcemanagerResourcemanagerStackDataSourceRepresentation) +
					compartmentIdVariableStr + ResourcemanagerStackResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttrSet(datasourceName, "stacks.#"),
					resource.TestCheckResourceAttr(datasourceName, "stacks.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "stacks.0.description"),
					resource.TestCheckResourceAttr(datasourceName, "stacks.0.display_name", "TestResourcemanagerStackResource_basic"),
					resource.TestCheckResourceAttr(datasourceName, "stacks.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "stacks.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "stacks.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "stacks.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config + `
					variable "resource_manager_stack_id" { default = "` + resourceManagerStackId + `" }
					` +
					acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stacks", "test_stacks", acctest.Required, acctest.Create, ResourcemanagerResourcemanagerStackDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stack", "test_stack", acctest.Required, acctest.Create, ResourcemanagerResourcemanagerStackSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ResourcemanagerStackResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "stack_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "config_source.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "config_source.0.config_source_type", "ZIP_UPLOAD"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TestResourcemanagerStackResource_basic"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "variables.%", "3"),
				),
			},
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ResourcemanagerStack") {
		resource.AddTestSweepers("ResourcemanagerStack", &resource.Sweeper{
			Name:         "ResourcemanagerStack",
			Dependencies: acctest.DependencyGraph["stack"],
			F:            sweepResourcemanagerStackResource,
		})
	}
}

func sweepResourcemanagerStackResource(compartment string) error {
	resourceManagerClient := acctest.GetTestClients(&schema.ResourceData{}).ResourceManagerClient()
	stackIds, err := getResourcemanagerStackIds(compartment)
	if err != nil {
		return err
	}
	for _, stackId := range stackIds {
		if ok := acctest.SweeperDefaultResourceId[stackId]; !ok {
			deleteStackRequest := oci_resourcemanager.DeleteStackRequest{}

			deleteStackRequest.StackId = &stackId

			deleteStackRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "resourcemanager")
			_, error := resourceManagerClient.DeleteStack(context.Background(), deleteStackRequest)
			if error != nil {
				fmt.Printf("Error deleting Stack %s %s, It is possible that the resource is already deleted. Please verify manually \n", stackId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &stackId, ResourcemanagerStackSweepWaitCondition, time.Duration(3*time.Minute),
				ResourcemanagerStackSweepResponseFetchOperation, "resourcemanager", true)
		}
	}
	return nil
}

func getResourcemanagerStackIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "StackId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	resourceManagerClient := acctest.GetTestClients(&schema.ResourceData{}).ResourceManagerClient()

	listStacksRequest := oci_resourcemanager.ListStacksRequest{}
	listStacksRequest.CompartmentId = &compartmentId
	listStacksRequest.LifecycleState = oci_resourcemanager.StackLifecycleStateActive
	listStacksResponse, err := resourceManagerClient.ListStacks(context.Background(), listStacksRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Stack list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, stack := range listStacksResponse.Items {
		id := *stack.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "StackId", id)
	}
	return resourceIds, nil
}

func ResourcemanagerStackSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if stackResponse, ok := response.Response.(oci_resourcemanager.GetStackResponse); ok {
		return stackResponse.LifecycleState != oci_resourcemanager.StackLifecycleStateDeleted
	}
	return false
}

func ResourcemanagerStackSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ResourceManagerClient().GetStack(context.Background(), oci_resourcemanager.GetStackRequest{
		StackId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
