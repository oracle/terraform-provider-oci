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
	oci_api_platform "github.com/oracle/oci-go-sdk/v65/apiplatform"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ApiPlatformApiPlatformInstanceRequiredOnlyResource = ApiPlatformApiPlatformInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_api_platform_api_platform_instance", "test_api_platform_instance", acctest.Required, acctest.Create, ApiPlatformApiPlatformInstanceRepresentation)

	ApiPlatformApiPlatformInstanceResourceConfig = ApiPlatformApiPlatformInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_api_platform_api_platform_instance", "test_api_platform_instance", acctest.Optional, acctest.Update, ApiPlatformApiPlatformInstanceRepresentation)

	ApiPlatformApiPlatformInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"api_platform_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_api_platform_api_platform_instance.test_api_platform_instance.id}`},
	}

	ApiPlatformApiPlatformInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_api_platform_api_platform_instance.test_api_platform_instance.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ApiPlatformApiPlatformInstanceDataSourceFilterRepresentation}}
	ApiPlatformApiPlatformInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_api_platform_api_platform_instance.test_api_platform_instance.id}`}},
	}

	ApiPlatformApiPlatformInstanceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `name`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ApiPlatformApiPlatformInstanceIgnoreChangesDeploymentRepresentation},
	}

	ApiPlatformApiPlatformInstanceIgnoreChangesDeploymentRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`}},
	}

	ApiPlatformApiPlatformInstanceResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: api_platform/default
func TestApiPlatformApiPlatformInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApiPlatformApiPlatformInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_api_platform_api_platform_instance.test_api_platform_instance"
	datasourceName := "data.oci_api_platform_api_platform_instances.test_api_platform_instances"
	singularDatasourceName := "data.oci_api_platform_api_platform_instance.test_api_platform_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApiPlatformApiPlatformInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_api_platform_api_platform_instance", "test_api_platform_instance", acctest.Optional, acctest.Create, ApiPlatformApiPlatformInstanceRepresentation), "apiplatform", "apiPlatformInstance", t)

	acctest.ResourceTest(t, testAccCheckApiPlatformApiPlatformInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ApiPlatformApiPlatformInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_api_platform_api_platform_instance", "test_api_platform_instance", acctest.Required, acctest.Create, ApiPlatformApiPlatformInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ApiPlatformApiPlatformInstanceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApiPlatformApiPlatformInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_api_platform_api_platform_instance", "test_api_platform_instance", acctest.Optional, acctest.Create, ApiPlatformApiPlatformInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ApiPlatformApiPlatformInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_api_platform_api_platform_instance", "test_api_platform_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ApiPlatformApiPlatformInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ApiPlatformApiPlatformInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_api_platform_api_platform_instance", "test_api_platform_instance", acctest.Optional, acctest.Update, ApiPlatformApiPlatformInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_api_platform_api_platform_instances", "test_api_platform_instances", acctest.Optional, acctest.Update, ApiPlatformApiPlatformInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + ApiPlatformApiPlatformInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_api_platform_api_platform_instance", "test_api_platform_instance", acctest.Optional, acctest.Update, ApiPlatformApiPlatformInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "api_platform_instance_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "api_platform_instance_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_api_platform_api_platform_instance", "test_api_platform_instance", acctest.Required, acctest.Create, ApiPlatformApiPlatformInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApiPlatformApiPlatformInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "api_platform_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_app.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "uris.#", "1"),
			),
		},
		// verify resource import
		{
			Config:                  config + ApiPlatformApiPlatformInstanceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckApiPlatformApiPlatformInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ApiPlatformClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_api_platform_api_platform_instance" {
			noResourceFound = false
			request := oci_api_platform.GetApiPlatformInstanceRequest{}

			tmp := rs.Primary.ID
			request.ApiPlatformInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "api_platform")

			response, err := client.GetApiPlatformInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_api_platform.ApiPlatformInstanceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ApiPlatformApiPlatformInstance") {
		resource.AddTestSweepers("ApiPlatformApiPlatformInstance", &resource.Sweeper{
			Name:         "ApiPlatformApiPlatformInstance",
			Dependencies: acctest.DependencyGraph["apiPlatformInstance"],
			F:            sweepApiPlatformApiPlatformInstanceResource,
		})
	}
}

func sweepApiPlatformApiPlatformInstanceResource(compartment string) error {
	apiPlatformClient := acctest.GetTestClients(&schema.ResourceData{}).ApiPlatformClient()
	apiPlatformInstanceIds, err := getApiPlatformApiPlatformInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, apiPlatformInstanceId := range apiPlatformInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[apiPlatformInstanceId]; !ok {
			deleteApiPlatformInstanceRequest := oci_api_platform.DeleteApiPlatformInstanceRequest{}

			deleteApiPlatformInstanceRequest.ApiPlatformInstanceId = &apiPlatformInstanceId

			deleteApiPlatformInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "api_platform")
			_, error := apiPlatformClient.DeleteApiPlatformInstance(context.Background(), deleteApiPlatformInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting ApiPlatformInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", apiPlatformInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &apiPlatformInstanceId, ApiPlatformApiPlatformInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				ApiPlatformApiPlatformInstanceSweepResponseFetchOperation, "api_platform", true)
		}
	}
	return nil
}

func getApiPlatformApiPlatformInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ApiPlatformInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	apiPlatformClient := acctest.GetTestClients(&schema.ResourceData{}).ApiPlatformClient()

	listApiPlatformInstancesRequest := oci_api_platform.ListApiPlatformInstancesRequest{}
	listApiPlatformInstancesRequest.CompartmentId = &compartmentId
	listApiPlatformInstancesRequest.LifecycleState = oci_api_platform.ApiPlatformInstanceLifecycleStateActive
	listApiPlatformInstancesResponse, err := apiPlatformClient.ListApiPlatformInstances(context.Background(), listApiPlatformInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ApiPlatformInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, apiPlatformInstance := range listApiPlatformInstancesResponse.Items {
		id := *apiPlatformInstance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ApiPlatformInstanceId", id)
	}
	return resourceIds, nil
}

func ApiPlatformApiPlatformInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if apiPlatformInstanceResponse, ok := response.Response.(oci_api_platform.GetApiPlatformInstanceResponse); ok {
		return apiPlatformInstanceResponse.LifecycleState != oci_api_platform.ApiPlatformInstanceLifecycleStateDeleted
	}
	return false
}

func ApiPlatformApiPlatformInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ApiPlatformClient().GetApiPlatformInstance(context.Background(), oci_api_platform.GetApiPlatformInstanceRequest{
		ApiPlatformInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
