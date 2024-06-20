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
	oci_globally_distributed_database "github.com/oracle/oci-go-sdk/v65/globallydistributeddatabase"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GloballyDistributedDatabasePrivateEndpointRequiredOnlyResource = GloballyDistributedDatabasePrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, GloballyDistributedDatabasePrivateEndpointRepresentation)

	GloballyDistributedDatabasePrivateEndpointResourceConfig = GloballyDistributedDatabasePrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Update, GloballyDistributedDatabasePrivateEndpointRepresentation)

	GloballyDistributedDatabasePrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_globally_distributed_database_private_endpoint.test_private_endpoint.id}`},
	}

	GloballyDistributedDatabasePrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `pe0001`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GloballyDistributedDatabasePrivateEndpointDataSourceFilterRepresentation}}
	GloballyDistributedDatabasePrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_globally_distributed_database_private_endpoint.test_private_endpoint.id}`}},
	}

	GloballyDistributedDatabasePrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `pe0001`, Update: `displayName2`},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		//"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.vcn_id}`},
		//"nsg1_ids": acctest.Representation{RepType: acctest.Required, Create: `${var.nsg1_ids}`},
		//"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		//"nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}, Update: []string{`nsgIds2`}},
	}

	GloballyDistributedDatabasePrivateEndpointResourceDependencies = "" /*acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +*/
	/*DefinedTagsDependencies*/
)

// issue-routing-tag: globally_distributed_database/default
func TestGloballyDistributedDatabasePrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGloballyDistributedDatabasePrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	vcnId := utils.GetEnvSettingWithBlankDefault("vcn_ocid")
	vcnIdVariableStr := fmt.Sprintf("variable \"vcn_id\" { default = \"%s\" }\n", vcnId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	nsgIds := "null"
	nsgIdsVariableStr := fmt.Sprintf("variable \"nsg1_ids\" { default = %s }\n", nsgIds)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_globally_distributed_database_private_endpoint.test_private_endpoint"
	datasourceName := "data.oci_globally_distributed_database_private_endpoints.test_private_endpoints"
	singularDatasourceName := "data.oci_globally_distributed_database_private_endpoint.test_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+subnetIdVariableStr+nsgIdsVariableStr+vcnIdVariableStr+GloballyDistributedDatabasePrivateEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Create, GloballyDistributedDatabasePrivateEndpointRepresentation), "globallydistributeddatabase", "privateEndpoint", t)

	acctest.ResourceTest(t, testAccCheckGloballyDistributedDatabasePrivateEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + nsgIdsVariableStr + GloballyDistributedDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, GloballyDistributedDatabasePrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "pe0001"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),
				//resource.TestCheckResourceAttr(resourceName, "nsg_ids", nsgIds),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + nsgIdsVariableStr + vcnIdVariableStr + GloballyDistributedDatabasePrivateEndpointResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + nsgIdsVariableStr + vcnIdVariableStr + GloballyDistributedDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Create, GloballyDistributedDatabasePrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "pe0001"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				//resource.TestCheckResourceAttr(resourceName, "nsg_ids", nsgIds),
				resource.TestCheckResourceAttr(resourceName, "vcn_id", vcnId),
				//resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + nsgIdsVariableStr + vcnIdVariableStr + compartmentIdUVariableStr + GloballyDistributedDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GloballyDistributedDatabasePrivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "pe0001"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				////resource.TestCheckResourceAttr(resourceName, "nsg_ids", nsgIds),
				resource.TestCheckResourceAttr(resourceName, "vcn_id", vcnId),
				//resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + nsgIdsVariableStr + vcnIdVariableStr + GloballyDistributedDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Update, GloballyDistributedDatabasePrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				//resource.TestCheckResourceAttr(resourceName, "nsg_ids", nsgIds),
				resource.TestCheckResourceAttr(resourceName, "vcn_id", vcnId),
				//resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_globally_distributed_database_private_endpoints", "test_private_endpoints", acctest.Optional, acctest.Update, GloballyDistributedDatabasePrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + nsgIdsVariableStr + vcnIdVariableStr + GloballyDistributedDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Update, GloballyDistributedDatabasePrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "private_endpoint_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "private_endpoint_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_globally_distributed_database_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, GloballyDistributedDatabasePrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + nsgIdsVariableStr + vcnIdVariableStr + GloballyDistributedDatabasePrivateEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sharded_databases.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcn_id"),
			),
		},
		// verify resource import
		{
			Config:                  config + GloballyDistributedDatabasePrivateEndpointRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGloballyDistributedDatabasePrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ShardedDatabaseServiceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_globally_distributed_database_private_endpoint" {
			noResourceFound = false
			request := oci_globally_distributed_database.GetPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.PrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "globally_distributed_database")

			response, err := client.GetPrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_globally_distributed_database.PrivateEndpointLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GloballyDistributedDatabasePrivateEndpoint") {
		resource.AddTestSweepers("GloballyDistributedDatabasePrivateEndpoint", &resource.Sweeper{
			Name:         "GloballyDistributedDatabasePrivateEndpoint",
			Dependencies: acctest.DependencyGraph["privateEndpoint"],
			F:            sweepGloballyDistributedDatabasePrivateEndpointResource,
		})
	}
}

func sweepGloballyDistributedDatabasePrivateEndpointResource(compartment string) error {
	shardedDatabaseServiceClient := acctest.GetTestClients(&schema.ResourceData{}).ShardedDatabaseServiceClient()
	privateEndpointIds, err := getGloballyDistributedDatabasePrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, privateEndpointId := range privateEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[privateEndpointId]; !ok {
			deletePrivateEndpointRequest := oci_globally_distributed_database.DeletePrivateEndpointRequest{}

			deletePrivateEndpointRequest.PrivateEndpointId = &privateEndpointId

			deletePrivateEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "globally_distributed_database")
			_, error := shardedDatabaseServiceClient.DeletePrivateEndpoint(context.Background(), deletePrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting PrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", privateEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &privateEndpointId, GloballyDistributedDatabasePrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				GloballyDistributedDatabasePrivateEndpointSweepResponseFetchOperation, "globally_distributed_database", true)
		}
	}
	return nil
}

func getGloballyDistributedDatabasePrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	shardedDatabaseServiceClient := acctest.GetTestClients(&schema.ResourceData{}).ShardedDatabaseServiceClient()

	listPrivateEndpointsRequest := oci_globally_distributed_database.ListPrivateEndpointsRequest{}
	listPrivateEndpointsRequest.CompartmentId = &compartmentId
	listPrivateEndpointsRequest.LifecycleState = oci_globally_distributed_database.PrivateEndpointLifecycleStateActive
	listPrivateEndpointsResponse, err := shardedDatabaseServiceClient.ListPrivateEndpoints(context.Background(), listPrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, privateEndpoint := range listPrivateEndpointsResponse.Items {
		id := *privateEndpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PrivateEndpointId", id)
	}
	return resourceIds, nil
}

func GloballyDistributedDatabasePrivateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if privateEndpointResponse, ok := response.Response.(oci_globally_distributed_database.GetPrivateEndpointResponse); ok {
		return privateEndpointResponse.LifecycleState != oci_globally_distributed_database.PrivateEndpointLifecycleStateDeleted
	}
	return false
}

func GloballyDistributedDatabasePrivateEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ShardedDatabaseServiceClient().GetPrivateEndpoint(context.Background(), oci_globally_distributed_database.GetPrivateEndpointRequest{
		PrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
