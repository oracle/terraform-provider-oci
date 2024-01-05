// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	oci_oda "github.com/oracle/oci-go-sdk/v65/oda"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OdaOdaPrivateEndpointRequiredOnlyResource = OdaOdaPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint", "test_oda_private_endpoint", acctest.Required, acctest.Create, OdaOdaPrivateEndpointRepresentation)

	OdaOdaPrivateEndpointResourceConfig = OdaOdaPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint", "test_oda_private_endpoint", acctest.Optional, acctest.Update, OdaOdaPrivateEndpointRepresentation)

	OdaOdaOdaPrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"oda_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_oda_oda_private_endpoint.test_oda_private_endpoint.id}`},
	}

	OdaOdaOdaPrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `CREATING`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OdaOdaPrivateEndpointDataSourceFilterRepresentation},
	}

	OdaOdaPrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_oda_oda_private_endpoint.test_oda_private_endpoint.id}`}},
	}

	OdaOdaPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
	}

	PESubnetRepresentation = map[string]interface{}{
		"cidr_block":                 acctest.Representation{RepType: acctest.Required, Create: `10.0.0.16/28`},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"dns_label":                  acctest.Representation{RepType: acctest.Required, Create: `testsubnet`},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
		"security_list_ids":          acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_sec_list.id}`}},
		"route_table_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`},
		"prohibit_public_ip_on_vnic": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"prohibit_internet_ingress":  acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	PEVcnRepresentation = map[string]interface{}{
		"cidr_block":     acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/25`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dns_label":      acctest.Representation{RepType: acctest.Required, Create: `testvcn`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	PESecurityListRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"egress_security_rules":  []acctest.RepresentationGroup{{RepType: acctest.Required, Group: PESecurityListEgressSecurityRulesRepresentation}},
		"ingress_security_rules": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: PESecurityListIngressSecurityRulesRepresentation}},
	}

	PESecurityListEgressSecurityRulesRepresentation = map[string]interface{}{
		"destination": acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `all`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
	}

	PESecurityListIngressSecurityRulesRepresentation = map[string]interface{}{
		"source":    acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"protocol":  acctest.Representation{RepType: acctest.Required, Create: `all`},
		"stateless": acctest.Representation{RepType: acctest.Required, Create: `false`},
	}

	PENatGatewayRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	PERouteTableRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"route_rules":    acctest.RepresentationGroup{RepType: acctest.Required, Group: PERouteTableRouteRulesRepresentation},
	}

	PERouteTableRouteRulesRepresentation = map[string]interface{}{
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_nat_gateway.test_ng.id}`},
		"destination":       acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
	}

	OdaOdaPrivateEndpointResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_sec_list", acctest.Required, acctest.Create, PESecurityListRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, PESubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, PEVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_nat_gateway", "test_ng", acctest.Required, acctest.Create, PENatGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, PERouteTableRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: oda/default
func TestOdaOdaPrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOdaOdaPrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_oda_oda_private_endpoint.test_oda_private_endpoint"
	datasourceName := "data.oci_oda_oda_private_endpoints.test_oda_private_endpoints"
	singularDatasourceName := "data.oci_oda_oda_private_endpoint.test_oda_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OdaOdaPrivateEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint", "test_oda_private_endpoint", acctest.Optional, acctest.Create, OdaOdaPrivateEndpointRepresentation), "oda", "odaPrivateEndpoint", t)

	acctest.ResourceTest(t, testAccCheckOdaOdaPrivateEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OdaOdaPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint", "test_oda_private_endpoint", acctest.Required, acctest.Create, OdaOdaPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OdaOdaPrivateEndpointResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OdaOdaPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint", "test_oda_private_endpoint", acctest.Optional, acctest.Create, OdaOdaPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OdaOdaPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint", "test_oda_private_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OdaOdaPrivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
			Config: config + compartmentIdVariableStr + OdaOdaPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint", "test_oda_private_endpoint", acctest.Optional, acctest.Update, OdaOdaPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_oda_oda_private_endpoints", "test_oda_private_endpoints", acctest.Optional, acctest.Update, OdaOdaOdaPrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + OdaOdaPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint", "test_oda_private_endpoint", acctest.Optional, acctest.Update, OdaOdaPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "CREATING"),

				resource.TestCheckResourceAttr(datasourceName, "oda_private_endpoint_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oda_private_endpoint_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_oda_oda_private_endpoint", "test_oda_private_endpoint", acctest.Required, acctest.Create, OdaOdaOdaPrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OdaOdaPrivateEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oda_private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + OdaOdaPrivateEndpointRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOdaOdaPrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_oda_oda_private_endpoint" {
			noResourceFound = false
			request := oci_oda.GetOdaPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.OdaPrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "oda")

			response, err := client.GetOdaPrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_oda.OdaPrivateEndpointLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OdaOdaPrivateEndpoint") {
		resource.AddTestSweepers("OdaOdaPrivateEndpoint", &resource.Sweeper{
			Name:         "OdaOdaPrivateEndpoint",
			Dependencies: acctest.DependencyGraph["odaPrivateEndpoint"],
			F:            sweepOdaOdaPrivateEndpointResource,
		})
	}
}

func sweepOdaOdaPrivateEndpointResource(compartment string) error {
	managementClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementClient()
	odaPrivateEndpointIds, err := getOdaOdaPrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, odaPrivateEndpointId := range odaPrivateEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[odaPrivateEndpointId]; !ok {
			deleteOdaPrivateEndpointRequest := oci_oda.DeleteOdaPrivateEndpointRequest{}

			deleteOdaPrivateEndpointRequest.OdaPrivateEndpointId = &odaPrivateEndpointId

			deleteOdaPrivateEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "oda")
			_, error := managementClient.DeleteOdaPrivateEndpoint(context.Background(), deleteOdaPrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting OdaPrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", odaPrivateEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &odaPrivateEndpointId, OdaOdaPrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				OdaOdaPrivateEndpointSweepResponseFetchOperation, "oda", true)
		}
	}
	return nil
}

func getOdaOdaPrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OdaPrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	managementClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementClient()

	listOdaPrivateEndpointsRequest := oci_oda.ListOdaPrivateEndpointsRequest{}
	listOdaPrivateEndpointsRequest.CompartmentId = &compartmentId
	listOdaPrivateEndpointsRequest.LifecycleState = oci_oda.OdaPrivateEndpointLifecycleStateActive
	listOdaPrivateEndpointsResponse, err := managementClient.ListOdaPrivateEndpoints(context.Background(), listOdaPrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OdaPrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, odaPrivateEndpoint := range listOdaPrivateEndpointsResponse.Items {
		id := *odaPrivateEndpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OdaPrivateEndpointId", id)
	}
	return resourceIds, nil
}

func OdaOdaPrivateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if odaPrivateEndpointResponse, ok := response.Response.(oci_oda.GetOdaPrivateEndpointResponse); ok {
		return odaPrivateEndpointResponse.LifecycleState != oci_oda.OdaPrivateEndpointLifecycleStateDeleted
	}
	return false
}

func OdaOdaPrivateEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ManagementClient().GetOdaPrivateEndpoint(context.Background(), oci_oda.GetOdaPrivateEndpointRequest{
		OdaPrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
