// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DistributedDatabaseDistributedDatabasePrivateEndpointRequiredOnlyResource = /*DistributedDatabaseDistributedDatabasePrivateEndpointResourceDependencies +*/
	acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoint", "test_distributed_database_private_endpoint", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabasePrivateEndpointRepresentation)

	DistributedDatabaseDistributedDatabasePrivateEndpointResourceConfig = /*DistributedDatabaseDistributedDatabasePrivateEndpointResourceDependencies +*/
	acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoint", "test_distributed_database_private_endpoint", acctest.Optional, acctest.Update, DistributedDatabaseDistributedDatabasePrivateEndpointRepresentation)

	DistributedDatabaseDistributedDatabasePrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"distributed_database_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_distributed_database_distributed_database_private_endpoint.test_distributed_database_private_endpoint.id}`},
	}

	DistributedDatabaseDistributedDatabasePrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `pe_tf_int_test`, Update: `pe_tf_int_test`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabasePrivateEndpointDataSourceFilterRepresentation}}
	DistributedDatabaseDistributedDatabasePrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_distributed_database_distributed_database_private_endpoint.test_distributed_database_private_endpoint.id}`}},
	}

	DistributedDatabaseDistributedDatabasePrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `pe_tf_int_test`, Update: `pe_tf_int_test`},
		// "subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		// NOTE: We intentionally do not include DefinedTagsDependencies or any defined_tags/freeform_tags
		// assertions in these integration tests. Tag namespace/tag resources require extra IAM
		// permissions (Identity Tag Namespace/Tag management) and can cause test failures with
		// 404 NotAuthorizedOrNotFound in restricted test tenancies, even when the service resource
		// under test (e.g., Distributed DB Private Endpoint) works correctly.
		//"defined_tags":                     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		// NOTE: We intentionally do not include DefinedTagsDependencies or any defined_tags/freeform_tags
		// assertions in these integration tests. Tag namespace/tag resources require extra IAM
		// permissions (Identity Tag Namespace/Tag management) and can cause test failures with
		// 404 NotAuthorizedOrNotFound in restricted test tenancies, even when the service resource
		// under test (e.g., Distributed DB Private Endpoint) works correctly.
		//"freeform_tags":                    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		// Use nsgs defined as env. variables to avoid creating/deleting NSGs in the test

		//"nsg_ids":                          acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}, Update: []string{`nsgIds2`}},
		"nsg_ids": acctest.Representation{
			RepType: acctest.Optional,
			Create:  []string{`${var.nsg_id}`},
			Update:  []string{`${var.nsg_idU}`},
		},

		"reinstate_proxy_instance_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}

	//DistributedDatabaseDistributedDatabasePrivateEndpointResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) /*+
	//acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
	//acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)*/
	// NOTE: We intentionally do not include DefinedTagsDependencies or any defined_tags/freeform_tags
// assertions in these integration tests. Tag namespace/tag resources require extra IAM
// permissions (Identity Tag Namespace/Tag management) and can cause test failures with
// 404 NotAuthorizedOrNotFound in restricted test tenancies, even when the service resource
// under test (e.g., Distributed DB Private Endpoint) works correctly.
/*+
DefinedTagsDependencies*/
)

// issue-routing-tag: distributed_database/default
func TestDistributedDatabaseDistributedDatabasePrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDistributedDatabaseDistributedDatabasePrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	vcnId := utils.GetEnvSettingWithBlankDefault("vcn_ocid")
	vcnIdVariableStr := fmt.Sprintf("variable \"vcn_id\" { default = \"%s\" }\n", vcnId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	nsgId := utils.GetEnvSettingWithBlankDefault("nsg_ocid")
	nsgIdVariableStr := fmt.Sprintf("variable \"nsg_id\" { default = \"%s\" }\n", nsgId)

	nsgIdU := utils.GetEnvSettingWithDefault("nsgU_ocid", nsgId)
	nsgIdUVariableStr := fmt.Sprintf("variable \"nsg_idU\" { default = \"%s\" }\n", nsgIdU)

	// Excluding Update Compartment
	// compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	// compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_distributed_database_distributed_database_private_endpoint.test_distributed_database_private_endpoint"
	datasourceName := "data.oci_distributed_database_distributed_database_private_endpoints.test_distributed_database_private_endpoints"
	singularDatasourceName := "data.oci_distributed_database_distributed_database_private_endpoint.test_distributed_database_private_endpoint"
	privateEndpointRepresentation := DistributedDatabaseDistributedDatabasePrivateEndpointRepresentation
	if nsgIdU == "" {
		privateEndpointRepresentation = acctest.RepresentationCopyWithRemovedProperties(privateEndpointRepresentation, []string{"nsg_ids"})
	}

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+vcnIdVariableStr+subnetIdVariableStr+ /*DistributedDatabaseDistributedDatabasePrivateEndpointResourceDependencies+*/
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoint", "test_distributed_database_private_endpoint", acctest.Optional, acctest.Create, privateEndpointRepresentation), "distributeddatabase", "distributedDatabasePrivateEndpoint", t)

	acctest.ResourceTest(t, testAccCheckDistributedDatabaseDistributedDatabasePrivateEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + nsgIdVariableStr + nsgIdUVariableStr + vcnIdVariableStr + subnetIdVariableStr /*+ DistributedDatabaseDistributedDatabasePrivateEndpointResourceDependencies*/ + acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoint", "test_distributed_database_private_endpoint", acctest.Required, acctest.Create, privateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "pe_tf_int_test"),
				//resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Excluding Create with optionals
		// // delete before next Create
		// {
		// 	Config: config + nsgIdVariableStr + nsgIdUVariableStr + compartmentIdVariableStr + vcnIdVariableStr + subnetIdVariableStr, /*+ DistributedDatabaseDistributedDatabasePrivateEndpointResourceDependencies*/
		// },
		// verify Create with optionals
		// {
		// 	Config: config + nsgIdVariableStr + nsgIdUVariableStr + compartmentIdVariableStr + vcnIdVariableStr + subnetIdVariableStr /*+ DistributedDatabaseDistributedDatabasePrivateEndpointResourceDependencies*/ +
		// 		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoint", "test_distributed_database_private_endpoint", acctest.Optional, acctest.Create, DistributedDatabaseDistributedDatabasePrivateEndpointRepresentation),
		// 	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		// 		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
		// 		resource.TestCheckResourceAttr(resourceName, "description", "description"),
		// 		resource.TestCheckResourceAttr(resourceName, "display_name", "pe0001"),
		// 		// NOTE: We intentionally do not include DefinedTagsDependencies or any defined_tags/freeform_tags
		// 		// assertions in these integration tests. Tag namespace/tag resources require extra IAM
		// 		// permissions (Identity Tag Namespace/Tag management) and can cause test failures with
		// 		// 404 NotAuthorizedOrNotFound in restricted test tenancies, even when the service resource
		// 		// under test (e.g., Distributed DB Private Endpoint) works correctly.
		// 		//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "state"),
		// 		//resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),
		// 		resource.TestCheckResourceAttrSet(resourceName, "time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
		// 		// resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "vcn_id", vcnId),

		// 		func(s *terraform.State) (err error) {
		// 			resId, err = acctest.FromInstanceState(s, resourceName, "id")
		// 			if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
		// 				if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
		// 					return errExport
		// 				}
		// 			}
		// 			return err
		// 		},
		// 	),
		// },

		// verify Update to the compartment (the compartment will be switched back in the next step)
		// {
		// 	Config: config + nsgIdVariableStr + nsgIdUVariableStr + compartmentIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + compartmentIdUVariableStr /*+ DistributedDatabaseDistributedDatabasePrivateEndpointResourceDependencies*/ +
		// 		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoint", "test_distributed_database_private_endpoint", acctest.Optional, acctest.Create,
		// 			acctest.RepresentationCopyWithNewProperties(DistributedDatabaseDistributedDatabasePrivateEndpointRepresentation, map[string]interface{}{
		// 				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
		// 			})),
		// 	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		// 		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
		// 		resource.TestCheckResourceAttr(resourceName, "description", "description"),
		// 		resource.TestCheckResourceAttr(resourceName, "display_name", "pe0001"),
		// 		// NOTE: We intentionally do not include DefinedTagsDependencies or any defined_tags/freeform_tags
		// 		// assertions in these integration tests. Tag namespace/tag resources require extra IAM
		// 		// permissions (Identity Tag Namespace/Tag management) and can cause test failures with
		// 		// 404 NotAuthorizedOrNotFound in restricted test tenancies, even when the service resource
		// 		// under test (e.g., Distributed DB Private Endpoint) works correctly.
		// 		//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "state"),
		// 		//resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),
		// 		resource.TestCheckResourceAttrSet(resourceName, "time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
		// 		//resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "vcn_id", vcnId),

		// 		func(s *terraform.State) (err error) {
		// 			resId2, err = acctest.FromInstanceState(s, resourceName, "id")
		// 			if resId != resId2 {
		// 				return fmt.Errorf("resource recreated when it was supposed to be updated")
		// 			}
		// 			return err
		// 		},
		// 	),
		// },

		// verify updates to updatable parameters
		{
			Config: config + nsgIdVariableStr + nsgIdUVariableStr + compartmentIdVariableStr + vcnIdVariableStr + subnetIdVariableStr /*+ DistributedDatabaseDistributedDatabasePrivateEndpointResourceDependencies*/ +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoint", "test_distributed_database_private_endpoint", acctest.Optional, acctest.Update, privateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "pe_tf_int_test"),
				// NOTE: We intentionally do not include DefinedTagsDependencies or any defined_tags/freeform_tags
				// assertions in these integration tests. Tag namespace/tag resources require extra IAM
				// permissions (Identity Tag Namespace/Tag management) and can cause test failures with
				// 404 NotAuthorizedOrNotFound in restricted test tenancies, even when the service resource
				// under test (e.g., Distributed DB Private Endpoint) works correctly.
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				//resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				//resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoints", "test_distributed_database_private_endpoints", acctest.Optional, acctest.Update, DistributedDatabaseDistributedDatabasePrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + nsgIdVariableStr + nsgIdUVariableStr + vcnIdVariableStr + subnetIdVariableStr /*+ DistributedDatabaseDistributedDatabasePrivateEndpointResourceDependencies*/ +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoint", "test_distributed_database_private_endpoint", acctest.Optional, acctest.Update, privateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "pe_tf_int_test"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "distributed_database_private_endpoint_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "distributed_database_private_endpoint_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoint", "test_distributed_database_private_endpoint", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabasePrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + nsgIdVariableStr + nsgIdUVariableStr + vcnIdVariableStr + subnetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoint", "test_distributed_database_private_endpoint", acctest.Optional, acctest.Update, privateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "distributed_database_private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "pe_tf_int_test"),
				// NOTE: We intentionally do not include DefinedTagsDependencies or any defined_tags/freeform_tags
				// assertions in these integration tests. Tag namespace/tag resources require extra IAM
				// permissions (Identity Tag Namespace/Tag management) and can cause test failures with
				// 404 NotAuthorizedOrNotFound in restricted test tenancies, even when the service resource
				// under test (e.g., Distributed DB Private Endpoint) works correctly.
				//resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "globally_distributed_autonomous_databases.#", "1"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "globally_distributed_databases.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "proxy_compute_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcn_id"),
			),
		},
		// verify resource import
		{
			Config:                  config + DistributedDatabaseDistributedDatabasePrivateEndpointRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"reinstate_proxy_instance_trigger"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDistributedDatabaseDistributedDatabasePrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DistributedDbPrivateEndpointServiceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_distributed_database_distributed_database_private_endpoint" {
			noResourceFound = false
			request := oci_distributed_database.GetDistributedDatabasePrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.DistributedDatabasePrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "distributed_database")

			response, err := client.GetDistributedDatabasePrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_distributed_database.DistributedDatabasePrivateEndpointLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DistributedDatabaseDistributedDatabasePrivateEndpoint") {
		resource.AddTestSweepers("DistributedDatabaseDistributedDatabasePrivateEndpoint", &resource.Sweeper{
			Name:         "DistributedDatabaseDistributedDatabasePrivateEndpoint",
			Dependencies: acctest.DependencyGraph["distributedDatabasePrivateEndpoint"],
			F:            sweepDistributedDatabaseDistributedDatabasePrivateEndpointResource,
		})
	}
}

func sweepDistributedDatabaseDistributedDatabasePrivateEndpointResource(compartment string) error {
	distributedDbPrivateEndpointServiceClient := acctest.GetTestClients(&schema.ResourceData{}).DistributedDbPrivateEndpointServiceClient()
	distributedDatabasePrivateEndpointIds, err := getDistributedDatabaseDistributedDatabasePrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, distributedDatabasePrivateEndpointId := range distributedDatabasePrivateEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[distributedDatabasePrivateEndpointId]; !ok {
			deleteDistributedDatabasePrivateEndpointRequest := oci_distributed_database.DeleteDistributedDatabasePrivateEndpointRequest{}

			deleteDistributedDatabasePrivateEndpointRequest.DistributedDatabasePrivateEndpointId = &distributedDatabasePrivateEndpointId

			deleteDistributedDatabasePrivateEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "distributed_database")
			_, error := distributedDbPrivateEndpointServiceClient.DeleteDistributedDatabasePrivateEndpoint(context.Background(), deleteDistributedDatabasePrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting DistributedDatabasePrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", distributedDatabasePrivateEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &distributedDatabasePrivateEndpointId, DistributedDatabaseDistributedDatabasePrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				DistributedDatabaseDistributedDatabasePrivateEndpointSweepResponseFetchOperation, "distributed_database", true)
		}
	}
	return nil
}

func getDistributedDatabaseDistributedDatabasePrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DistributedDatabasePrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	distributedDbPrivateEndpointServiceClient := acctest.GetTestClients(&schema.ResourceData{}).DistributedDbPrivateEndpointServiceClient()

	listDistributedDatabasePrivateEndpointsRequest := oci_distributed_database.ListDistributedDatabasePrivateEndpointsRequest{}
	listDistributedDatabasePrivateEndpointsRequest.CompartmentId = &compartmentId
	listDistributedDatabasePrivateEndpointsRequest.LifecycleState = oci_distributed_database.DistributedDatabasePrivateEndpointLifecycleStateActive
	listDistributedDatabasePrivateEndpointsResponse, err := distributedDbPrivateEndpointServiceClient.ListDistributedDatabasePrivateEndpoints(context.Background(), listDistributedDatabasePrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DistributedDatabasePrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, distributedDatabasePrivateEndpoint := range listDistributedDatabasePrivateEndpointsResponse.Items {
		id := *distributedDatabasePrivateEndpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DistributedDatabasePrivateEndpointId", id)
	}
	return resourceIds, nil
}

func DistributedDatabaseDistributedDatabasePrivateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if distributedDatabasePrivateEndpointResponse, ok := response.Response.(oci_distributed_database.GetDistributedDatabasePrivateEndpointResponse); ok {
		return distributedDatabasePrivateEndpointResponse.LifecycleState != oci_distributed_database.DistributedDatabasePrivateEndpointLifecycleStateDeleted
	}
	return false
}

func DistributedDatabaseDistributedDatabasePrivateEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DistributedDbPrivateEndpointServiceClient().GetDistributedDatabasePrivateEndpoint(context.Background(), oci_distributed_database.GetDistributedDatabasePrivateEndpointRequest{
		DistributedDatabasePrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
