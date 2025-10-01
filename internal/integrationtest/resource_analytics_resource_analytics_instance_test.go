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
	oci_resource_analytics "github.com/oracle/oci-go-sdk/v65/resourceanalytics"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ResourceAnalyticsResourceAnalyticsInstanceRequiredOnlyResource = ResourceAnalyticsResourceAnalyticsInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance", "test_resource_analytics_instance", acctest.Required, acctest.Create, ResourceAnalyticsResourceAnalyticsInstanceRepresentation)

	ResourceAnalyticsResourceAnalyticsInstanceResourceConfig = ResourceAnalyticsResourceAnalyticsInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance", "test_resource_analytics_instance", acctest.Optional, acctest.Update, ResourceAnalyticsResourceAnalyticsInstanceRepresentation)

	ResourceAnalyticsResourceAnalyticsInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"resource_analytics_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id}`},
	}

	ResourceAnalyticsResourceAnalyticsInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `asim-ra-test`, Update: `asim-ra-test2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ResourceAnalyticsResourceAnalyticsInstanceDataSourceFilterRepresentation}}
	ResourceAnalyticsResourceAnalyticsInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id}`}},
	}

	ResourceAnalyticsResourceAnalyticsInstanceRepresentation = map[string]interface{}{
		"adw_admin_password":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ResourceAnalyticsResourceAnalyticsInstanceAdwAdminPasswordRepresentation},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `asim-ra-test`, Update: `asim-ra-test2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_mutual_tls_required": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":          acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
	}
	ResourceAnalyticsResourceAnalyticsInstanceAdwAdminPasswordRepresentation = map[string]interface{}{
		"password_type": acctest.Representation{RepType: acctest.Required, Create: `PLAIN_TEXT`},
		"password":      acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
	}

	ResourceAnalyticsResourceAnalyticsInstanceResourceDependencies = ""
)

// issue-routing-tag: resource_analytics/default
func TestResourceAnalyticsResourceAnalyticsInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceAnalyticsResourceAnalyticsInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	defaultVariablesStr := subnetIdVariableStr + fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance"
	datasourceName := "data.oci_resource_analytics_resource_analytics_instances.test_resource_analytics_instances"
	singularDatasourceName := "data.oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+defaultVariablesStr+ResourceAnalyticsResourceAnalyticsInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance", "test_resource_analytics_instance", acctest.Optional, acctest.Create, ResourceAnalyticsResourceAnalyticsInstanceRepresentation), "resourceanalytics", "resourceAnalyticsInstance", t)

	acctest.ResourceTest(t, testAccCheckResourceAnalyticsResourceAnalyticsInstanceDestroy, []resource.TestStep{
		// STEP 0 (1/8)-- verify Create
		{
			Config: config + defaultVariablesStr + ResourceAnalyticsResourceAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance", "test_resource_analytics_instance", acctest.Required, acctest.Create, ResourceAnalyticsResourceAnalyticsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "adw_admin_password.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "adw_admin_password.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "adw_admin_password.0.password_type", "PLAIN_TEXT"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// STEP 1 (2/8)- delete before next Create
		{
			Config: config + defaultVariablesStr + ResourceAnalyticsResourceAnalyticsInstanceResourceDependencies,
		},
		// STEP 2 (3/8)- verify Create with optionals
		{
			Config: config + defaultVariablesStr + ResourceAnalyticsResourceAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance", "test_resource_analytics_instance", acctest.Optional, acctest.Create, ResourceAnalyticsResourceAnalyticsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "adw_admin_password.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "adw_admin_password.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "adw_admin_password.0.password_type", "PLAIN_TEXT"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "asim-ra-test"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_mutual_tls_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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

		// STEP 3 (4/8)- verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + defaultVariablesStr + compartmentIdUVariableStr + ResourceAnalyticsResourceAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance", "test_resource_analytics_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ResourceAnalyticsResourceAnalyticsInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "adw_admin_password.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "adw_admin_password.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "adw_admin_password.0.password_type", "PLAIN_TEXT"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "asim-ra-test"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_mutual_tls_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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

		// STEP 4 (5/8)- verify updates to updatable parameters
		{
			Config: config + defaultVariablesStr + ResourceAnalyticsResourceAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance", "test_resource_analytics_instance", acctest.Optional, acctest.Update, ResourceAnalyticsResourceAnalyticsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "adw_admin_password.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "adw_admin_password.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "adw_admin_password.0.password_type", "PLAIN_TEXT"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "asim-ra-test2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_mutual_tls_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
		// STEP 5 (6/8)- verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instances", "test_resource_analytics_instances", acctest.Optional, acctest.Update, ResourceAnalyticsResourceAnalyticsInstanceDataSourceRepresentation) +
				defaultVariablesStr + ResourceAnalyticsResourceAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance", "test_resource_analytics_instance", acctest.Optional, acctest.Update, ResourceAnalyticsResourceAnalyticsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "asim-ra-test2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "resource_analytics_instance_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "resource_analytics_instance_collection.0.items.#", "1"),
			),
		},
		// STEP 6 (7/8)- verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance", "test_resource_analytics_instance", acctest.Required, acctest.Create, ResourceAnalyticsResourceAnalyticsInstanceSingularDataSourceRepresentation) +
				defaultVariablesStr + ResourceAnalyticsResourceAnalyticsInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_analytics_instance_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "adw_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "asim-ra-test2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// STEP 7 (8/8)- verify resource import
		{
			Config:            config + ResourceAnalyticsResourceAnalyticsInstanceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"adw_admin_password",
				"is_mutual_tls_required",
				"license_model",
				"nsg_ids",
				"subnet_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckResourceAnalyticsResourceAnalyticsInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ResourceAnalyticsInstanceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_resource_analytics_resource_analytics_instance" {
			noResourceFound = false
			request := oci_resource_analytics.GetResourceAnalyticsInstanceRequest{}

			tmp := rs.Primary.ID
			request.ResourceAnalyticsInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "resource_analytics")

			response, err := client.GetResourceAnalyticsInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_resource_analytics.ResourceAnalyticsInstanceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ResourceAnalyticsResourceAnalyticsInstance") {
		resource.AddTestSweepers("ResourceAnalyticsResourceAnalyticsInstance", &resource.Sweeper{
			Name:         "ResourceAnalyticsResourceAnalyticsInstance",
			Dependencies: acctest.DependencyGraph["resourceAnalyticsInstance"],
			F:            sweepResourceAnalyticsResourceAnalyticsInstanceResource,
		})
	}
}

func sweepResourceAnalyticsResourceAnalyticsInstanceResource(compartment string) error {
	resourceAnalyticsInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).ResourceAnalyticsInstanceClient()
	resourceAnalyticsInstanceIds, err := getResourceAnalyticsResourceAnalyticsInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, resourceAnalyticsInstanceId := range resourceAnalyticsInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[resourceAnalyticsInstanceId]; !ok {
			deleteResourceAnalyticsInstanceRequest := oci_resource_analytics.DeleteResourceAnalyticsInstanceRequest{}

			deleteResourceAnalyticsInstanceRequest.ResourceAnalyticsInstanceId = &resourceAnalyticsInstanceId

			deleteResourceAnalyticsInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "resource_analytics")
			_, error := resourceAnalyticsInstanceClient.DeleteResourceAnalyticsInstance(context.Background(), deleteResourceAnalyticsInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting ResourceAnalyticsInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", resourceAnalyticsInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &resourceAnalyticsInstanceId, ResourceAnalyticsResourceAnalyticsInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				ResourceAnalyticsResourceAnalyticsInstanceSweepResponseFetchOperation, "resource_analytics", true)
		}
	}
	return nil
}

func getResourceAnalyticsResourceAnalyticsInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ResourceAnalyticsInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	resourceAnalyticsInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).ResourceAnalyticsInstanceClient()

	listResourceAnalyticsInstancesRequest := oci_resource_analytics.ListResourceAnalyticsInstancesRequest{}
	listResourceAnalyticsInstancesRequest.CompartmentId = &compartmentId
	listResourceAnalyticsInstancesRequest.LifecycleState = oci_resource_analytics.ResourceAnalyticsInstanceLifecycleStateNeedsAttention
	listResourceAnalyticsInstancesResponse, err := resourceAnalyticsInstanceClient.ListResourceAnalyticsInstances(context.Background(), listResourceAnalyticsInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ResourceAnalyticsInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, resourceAnalyticsInstance := range listResourceAnalyticsInstancesResponse.Items {
		id := *resourceAnalyticsInstance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ResourceAnalyticsInstanceId", id)
	}
	return resourceIds, nil
}

func ResourceAnalyticsResourceAnalyticsInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if resourceAnalyticsInstanceResponse, ok := response.Response.(oci_resource_analytics.GetResourceAnalyticsInstanceResponse); ok {
		return resourceAnalyticsInstanceResponse.LifecycleState != oci_resource_analytics.ResourceAnalyticsInstanceLifecycleStateDeleted
	}
	return false
}

func ResourceAnalyticsResourceAnalyticsInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ResourceAnalyticsInstanceClient().GetResourceAnalyticsInstance(context.Background(), oci_resource_analytics.GetResourceAnalyticsInstanceRequest{
		ResourceAnalyticsInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
