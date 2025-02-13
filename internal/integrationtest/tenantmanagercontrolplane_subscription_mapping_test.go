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
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	TenantmanagercontrolplaneSubscriptionMappingSingularDataSourceRepresentation = map[string]interface{}{
		"subscription_mapping_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_tenantmanagercontrolplane_subscription_mapping.test_subscription_mapping.id}`},
	}

	TenantmanagercontrolplaneSubscriptionMappingDataSourceRepresentation = map[string]interface{}{
		"subscription_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_tenantmanagercontrolplane_subscription_mapping.test_subscription_mapping.subscription_id}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_tenantmanagercontrolplane_subscription_mapping.test_subscription_mapping.compartment_id}`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_tenantmanagercontrolplane_subscription_mapping.test_subscription_mapping.state}`},
		"subscription_mapping_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_tenantmanagercontrolplane_subscription_mapping.test_subscription_mapping.id}`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: TenantmanagercontrolplaneSubscriptionMappingDataSourceFilterRepresentation}}
	TenantmanagercontrolplaneSubscriptionMappingDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_tenantmanagercontrolplane_subscription_mapping.test_subscription_mapping.id}`}},
	}

	TenantmanagercontrolplaneSubscriptionMappingRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_mapping_compartment_id}`},
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_mapping_subscription_id}`},
	}
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneSubscriptionMappingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneSubscriptionMappingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("subscription_mapping_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"subscription_mapping_compartment_id\" { default = \"%s\" }\n", compartmentId)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_mapping_subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_mapping_subscription_id\" { default = \"%s\" }\n", subscriptionId)

	resourceName := "oci_tenantmanagercontrolplane_subscription_mapping.test_subscription_mapping"
	datasourceName := "data.oci_tenantmanagercontrolplane_subscription_mappings.test_subscription_mappings"
	singularDatasourceName := "data.oci_tenantmanagercontrolplane_subscription_mapping.test_subscription_mapping"

	var resId string
	acctest.SaveConfigContent("", "", "", t)

	createConfig := config + compartmentIdVariableStr + subscriptionIdVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_tenantmanagercontrolplane_subscription_mapping", "test_subscription_mapping", acctest.Required, acctest.Create, TenantmanagercontrolplaneSubscriptionMappingRepresentation)
	dataSourceConfig := config + compartmentIdVariableStr + subscriptionIdVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_tenantmanagercontrolplane_subscription_mapping", "test_subscription_mapping", acctest.Required, acctest.Create, TenantmanagercontrolplaneSubscriptionMappingRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_subscription_mappings", "test_subscription_mappings", acctest.Optional, acctest.Create, TenantmanagercontrolplaneSubscriptionMappingDataSourceRepresentation)
	singularDataSourceConfig := config + compartmentIdVariableStr + subscriptionIdVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_tenantmanagercontrolplane_subscription_mapping", "test_subscription_mapping", acctest.Required, acctest.Create, TenantmanagercontrolplaneSubscriptionMappingRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_subscription_mapping", "test_subscription_mapping", acctest.Required, acctest.Create, TenantmanagercontrolplaneSubscriptionMappingSingularDataSourceRepresentation)
	importConfig := config + compartmentIdVariableStr + subscriptionIdVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_tenantmanagercontrolplane_subscription_mapping", "test_subscription_mapping", acctest.Required, acctest.Create, TenantmanagercontrolplaneSubscriptionMappingRepresentation)

	fmt.Printf("Create Config: %s\n", createConfig)
	fmt.Printf("Data Source Config: %s\n", dataSourceConfig)
	fmt.Printf("Singular Data Source Config: %s\n", singularDataSourceConfig)
	fmt.Printf("Import Config: %s\n", importConfig)

	acctest.ResourceTest(t, testAccCheckTenantmanagercontrolplaneSubscriptionMappingDestroy, []resource.TestStep{
		// verify Create
		{
			Config: createConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subscription_id"),

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
		// verify datasource
		{
			Config: dataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "state"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_mapping_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "subscription_mapping_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_mapping_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_mapping_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_explicitly_assigned"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  importConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckTenantmanagercontrolplaneSubscriptionMappingDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OrganizationsSubscriptionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_tenantmanagercontrolplane_subscription_mapping" {
			noResourceFound = false
			request := oci_tenantmanagercontrolplane.GetSubscriptionMappingRequest{}

			tmp := rs.Primary.ID
			request.SubscriptionMappingId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "tenantmanagercontrolplane")

			response, err := client.GetSubscriptionMapping(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_tenantmanagercontrolplane.SubscriptionMappingLifecycleStateDeleted):  true,
					string(oci_tenantmanagercontrolplane.SubscriptionMappingLifecycleStateInactive): true,
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
	if !acctest.InSweeperExcludeList("TenantmanagercontrolplaneSubscriptionMapping") {
		resource.AddTestSweepers("TenantmanagercontrolplaneSubscriptionMapping", &resource.Sweeper{
			Name:         "TenantmanagercontrolplaneSubscriptionMapping",
			Dependencies: acctest.DependencyGraph["subscriptionMapping"],
			F:            sweepTenantmanagercontrolplaneSubscriptionMappingResource,
		})
	}
}

func sweepTenantmanagercontrolplaneSubscriptionMappingResource(compartment string) error {
	subscriptionClient := acctest.GetTestClients(&schema.ResourceData{}).OrganizationsSubscriptionClient()
	subscriptionMappingIds, err := getTenantmanagercontrolplaneSubscriptionMappingIds(compartment)
	if err != nil {
		return err
	}
	for _, subscriptionMappingId := range subscriptionMappingIds {
		if ok := acctest.SweeperDefaultResourceId[subscriptionMappingId]; !ok {
			deleteSubscriptionMappingRequest := oci_tenantmanagercontrolplane.DeleteSubscriptionMappingRequest{}

			deleteSubscriptionMappingRequest.SubscriptionMappingId = &subscriptionMappingId

			deleteSubscriptionMappingRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "tenantmanagercontrolplane")
			_, error := subscriptionClient.DeleteSubscriptionMapping(context.Background(), deleteSubscriptionMappingRequest)
			if error != nil {
				fmt.Printf("Error deleting SubscriptionMapping %s %s, It is possible that the resource is already deleted. Please verify manually \n", subscriptionMappingId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &subscriptionMappingId, TenantmanagercontrolplaneSubscriptionMappingSweepWaitCondition, time.Duration(3*time.Minute),
				TenantmanagercontrolplaneSubscriptionMappingSweepResponseFetchOperation, "tenantmanagercontrolplane", true)
		}
	}
	return nil
}

func getTenantmanagercontrolplaneSubscriptionMappingIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SubscriptionMappingId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	subscriptionClient := acctest.GetTestClients(&schema.ResourceData{}).OrganizationsSubscriptionClient()

	listSubscriptionMappingsRequest := oci_tenantmanagercontrolplane.ListSubscriptionMappingsRequest{}
	listSubscriptionMappingsRequest.CompartmentId = &compartmentId

	subscriptionIds, error := getTenantmanagercontrolplaneSubscriptionIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting subscriptionId required for SubscriptionMapping resource requests \n")
	}
	for _, subscriptionId := range subscriptionIds {
		listSubscriptionMappingsRequest.SubscriptionId = &subscriptionId

		listSubscriptionMappingsRequest.LifecycleState = oci_tenantmanagercontrolplane.SubscriptionMappingLifecycleStateActive
		listSubscriptionMappingsResponse, err := subscriptionClient.ListSubscriptionMappings(context.Background(), listSubscriptionMappingsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting SubscriptionMapping list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, subscriptionMapping := range listSubscriptionMappingsResponse.Items {
			id := *subscriptionMapping.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SubscriptionMappingId", id)
		}

	}
	return resourceIds, nil
}

func TenantmanagercontrolplaneSubscriptionMappingSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if subscriptionMappingResponse, ok := response.Response.(oci_tenantmanagercontrolplane.GetSubscriptionMappingResponse); ok {
		return subscriptionMappingResponse.LifecycleState != oci_tenantmanagercontrolplane.SubscriptionMappingLifecycleStateInactive
	}
	return false
}

func TenantmanagercontrolplaneSubscriptionMappingSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OrganizationsSubscriptionClient().GetSubscriptionMapping(context.Background(), oci_tenantmanagercontrolplane.GetSubscriptionMappingRequest{
		SubscriptionMappingId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
