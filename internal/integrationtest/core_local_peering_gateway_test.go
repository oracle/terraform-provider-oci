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
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreLocalPeeringGatewayRequiredOnlyResource = CoreLocalPeeringGatewayResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_local_peering_gateway", "test_local_peering_gateway", acctest.Required, acctest.Create, CoreLocalPeeringGatewayRepresentation)

	CoreCoreLocalPeeringGatewayDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreLocalPeeringGatewayDataSourceFilterRepresentation}}
	CoreLocalPeeringGatewayDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_local_peering_gateway.test_local_peering_gateway.id}`}},
	}

	CoreSecondLocalPeeringGatewayDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn2.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreSecondLocalPeeringGatewayDataSourceFilterRepresentation}}
	CoreSecondLocalPeeringGatewayDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_local_peering_gateway.test_local_peering_gateway2.id}`}},
	}

	CoreLocalPeeringGatewayRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`},
	}

	secondLocalPeeringGatewayRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn2.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `requestor`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"peer_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_local_peering_gateway.test_local_peering_gateway.id}`},
	}

	secondLocalPeeringGatewayWithPeerId = `
variable "vcn_cidr_block2" { default = "10.1.0.0/16" }
variable "vcn_display_name2" { default = "displayName2" }
variable "vcn_dns_label2" { default = "dnslabel2" }

resource "oci_core_vcn" "test_vcn2" {
	#Required
	cidr_block = "${var.vcn_cidr_block2}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.vcn_display_name2}"
	dns_label = "${var.vcn_dns_label2}"
}
`
	CoreLocalPeeringGatewayResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, CoreInternetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, CoreRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreLocalPeeringGatewayResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreLocalPeeringGatewayResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_local_peering_gateway.test_local_peering_gateway"
	datasourceName := "data.oci_core_local_peering_gateways.test_local_peering_gateways"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreLocalPeeringGatewayResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_local_peering_gateway", "test_local_peering_gateway", acctest.Optional, acctest.Create, CoreLocalPeeringGatewayRepresentation), "core", "localPeeringGateway", t)

	acctest.ResourceTest(t, testAccCheckCoreLocalPeeringGatewayDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreLocalPeeringGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_local_peering_gateway", "test_local_peering_gateway", acctest.Required, acctest.Create, CoreLocalPeeringGatewayRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreLocalPeeringGatewayResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreLocalPeeringGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_local_peering_gateway", "test_local_peering_gateway", acctest.Optional, acctest.Create, CoreLocalPeeringGatewayRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
				resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreLocalPeeringGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_local_peering_gateway", "test_local_peering_gateway", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreLocalPeeringGatewayRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
				resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + CoreLocalPeeringGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_local_peering_gateway", "test_local_peering_gateway", acctest.Optional, acctest.Update, CoreLocalPeeringGatewayRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
				resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_local_peering_gateways", "test_local_peering_gateways", acctest.Optional, acctest.Update, CoreCoreLocalPeeringGatewayDataSourceRepresentation) +
				compartmentIdVariableStr + CoreLocalPeeringGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_local_peering_gateway", "test_local_peering_gateway", acctest.Optional, acctest.Update, CoreLocalPeeringGatewayRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_local_peering_gateway", "test_local_peering_gateway2", acctest.Optional, acctest.Update, secondLocalPeeringGatewayRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_local_peering_gateways", "test_local_peering_gateways2", acctest.Optional, acctest.Update, CoreSecondLocalPeeringGatewayDataSourceRepresentation) +
				secondLocalPeeringGatewayWithPeerId,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "local_peering_gateways.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "local_peering_gateways.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "local_peering_gateways.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "local_peering_gateways.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.is_cross_tenancy_peering"),
				resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.peering_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.peering_status_details"),
				resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.route_table_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.vcn_id"),
				resource.TestCheckResourceAttrSet(datasourceName+"2", "local_peering_gateways.0.peer_id"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreLocalPeeringGatewayRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
		// verify connect functionality
		{
			Config: config + compartmentIdVariableStr + CoreLocalPeeringGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_local_peering_gateway", "test_local_peering_gateway", acctest.Optional, acctest.Update, CoreLocalPeeringGatewayRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_local_peering_gateway", "test_local_peering_gateway2", acctest.Optional, acctest.Update, secondLocalPeeringGatewayRepresentation) +
				secondLocalPeeringGatewayWithPeerId,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
				resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_id"),
				resource.TestCheckResourceAttr(resourceName+"2", "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName+"2", "display_name", "requestor"),
				resource.TestCheckResourceAttrSet(resourceName+"2", "id"),
				resource.TestCheckResourceAttrSet(resourceName+"2", "is_cross_tenancy_peering"),
				resource.TestCheckResourceAttrSet(resourceName+"2", "peer_id"),
				resource.TestCheckResourceAttr(resourceName+"2", "peering_status", string(oci_core.LocalPeeringGatewayPeeringStatusPeered)),
				resource.TestCheckResourceAttr(resourceName+"2", "state", string(oci_core.LocalPeeringGatewayLifecycleStateAvailable)),
				resource.TestCheckResourceAttrSet(resourceName+"2", "time_created"),
				resource.TestCheckResourceAttrSet(resourceName+"2", "vcn_id"),
			),
		},
	})
}

func testAccCheckCoreLocalPeeringGatewayDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_local_peering_gateway" {
			noResourceFound = false
			request := oci_core.GetLocalPeeringGatewayRequest{}

			tmp := rs.Primary.ID
			request.LocalPeeringGatewayId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetLocalPeeringGateway(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.LocalPeeringGatewayLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreLocalPeeringGateway") {
		resource.AddTestSweepers("CoreLocalPeeringGateway", &resource.Sweeper{
			Name:         "CoreLocalPeeringGateway",
			Dependencies: acctest.DependencyGraph["localPeeringGateway"],
			F:            sweepCoreLocalPeeringGatewayResource,
		})
	}
}

func sweepCoreLocalPeeringGatewayResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	localPeeringGatewayIds, err := getCoreLocalPeeringGatewayIds(compartment)
	if err != nil {
		return err
	}
	for _, localPeeringGatewayId := range localPeeringGatewayIds {
		if ok := acctest.SweeperDefaultResourceId[localPeeringGatewayId]; !ok {
			deleteLocalPeeringGatewayRequest := oci_core.DeleteLocalPeeringGatewayRequest{}

			deleteLocalPeeringGatewayRequest.LocalPeeringGatewayId = &localPeeringGatewayId

			deleteLocalPeeringGatewayRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteLocalPeeringGateway(context.Background(), deleteLocalPeeringGatewayRequest)
			if error != nil {
				fmt.Printf("Error deleting LocalPeeringGateway %s %s, It is possible that the resource is already deleted. Please verify manually \n", localPeeringGatewayId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &localPeeringGatewayId, CoreLocalPeeringGatewaySweepWaitCondition, time.Duration(3*time.Minute),
				CoreLocalPeeringGatewaySweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreLocalPeeringGatewayIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "LocalPeeringGatewayId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listLocalPeeringGatewaysRequest := oci_core.ListLocalPeeringGatewaysRequest{}
	listLocalPeeringGatewaysRequest.CompartmentId = &compartmentId
	listLocalPeeringGatewaysResponse, err := virtualNetworkClient.ListLocalPeeringGateways(context.Background(), listLocalPeeringGatewaysRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting LocalPeeringGateway list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, localPeeringGateway := range listLocalPeeringGatewaysResponse.Items {
		id := *localPeeringGateway.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "LocalPeeringGatewayId", id)
	}
	return resourceIds, nil
}

func CoreLocalPeeringGatewaySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if localPeeringGatewayResponse, ok := response.Response.(oci_core.GetLocalPeeringGatewayResponse); ok {
		return localPeeringGatewayResponse.LifecycleState != oci_core.LocalPeeringGatewayLifecycleStateTerminated
	}
	return false
}

func CoreLocalPeeringGatewaySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetLocalPeeringGateway(context.Background(), oci_core.GetLocalPeeringGatewayRequest{
		LocalPeeringGatewayId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
