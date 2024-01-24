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
	CoreNatGatewayRequiredOnlyResource = CoreNatGatewayResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", acctest.Required, acctest.Create, CoreNatGatewayRepresentation)

	CoreNatGatewayResourceConfig = CoreNatGatewayResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", acctest.Optional, acctest.Update, CoreNatGatewayRepresentation)

	CoreCoreNatGatewaySingularDataSourceRepresentation = map[string]interface{}{
		"nat_gateway_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_nat_gateway.test_nat_gateway.id}`},
	}

	CoreCoreNatGatewayDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"vcn_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreNatGatewayDataSourceFilterRepresentation}}
	CoreNatGatewayDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_nat_gateway.test_nat_gateway.id}`}},
	}

	CoreNatGatewayRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"block_traffic":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"route_table_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_route_table.test_route_table.id}`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	CoreNatGatewayResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, CoreRouteTableRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/pnp
func TestCoreNatGatewayResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreNatGatewayResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_nat_gateway.test_nat_gateway"
	datasourceName := "data.oci_core_nat_gateways.test_nat_gateways"
	singularDatasourceName := "data.oci_core_nat_gateway.test_nat_gateway"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreNatGatewayResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", acctest.Optional, acctest.Create, CoreNatGatewayRepresentation), "core", "natGateway", t)

	acctest.ResourceTest(t, testAccCheckCoreNatGatewayDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreNatGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", acctest.Required, acctest.Create, CoreNatGatewayRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreNatGatewayResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreNatGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", acctest.Optional, acctest.Create, CoreNatGatewayRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "block_traffic", "false"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "nat_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "public_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					// use resource discovery to discover the NAT gateway and verify that it is discovered
					// when NAT gateway is in the same compartment as the VCN
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreNatGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreNatGatewayRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "block_traffic", "false"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "nat_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "public_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}

					// use resource discovery to discover the NAT gateway and verify that it is discovered
					// even when NAT gateway is in a different compartment than the VCN
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId2, &compartmentIdU, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreNatGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", acctest.Optional, acctest.Update, CoreNatGatewayRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "block_traffic", "true"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "nat_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "public_ip_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_nat_gateways", "test_nat_gateways", acctest.Optional, acctest.Update, CoreCoreNatGatewayDataSourceRepresentation) +
				compartmentIdVariableStr + CoreNatGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", acctest.Optional, acctest.Update, CoreNatGatewayRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "nat_gateways.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.block_traffic", "true"),
				resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.nat_ip"),
				resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.public_ip_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.route_table_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.vcn_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", acctest.Required, acctest.Create, CoreCoreNatGatewaySingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreNatGatewayResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nat_gateway_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "block_traffic", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nat_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreNatGatewayRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreNatGatewayDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_nat_gateway" {
			noResourceFound = false
			request := oci_core.GetNatGatewayRequest{}

			tmp := rs.Primary.ID
			request.NatGatewayId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetNatGateway(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.NatGatewayLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreNatGateway") {
		resource.AddTestSweepers("CoreNatGateway", &resource.Sweeper{
			Name:         "CoreNatGateway",
			Dependencies: acctest.DependencyGraph["natGateway"],
			F:            sweepCoreNatGatewayResource,
		})
	}
}

func sweepCoreNatGatewayResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	natGatewayIds, err := getCoreNatGatewayIds(compartment)
	if err != nil {
		return err
	}
	for _, natGatewayId := range natGatewayIds {
		if ok := acctest.SweeperDefaultResourceId[natGatewayId]; !ok {
			deleteNatGatewayRequest := oci_core.DeleteNatGatewayRequest{}

			deleteNatGatewayRequest.NatGatewayId = &natGatewayId

			deleteNatGatewayRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteNatGateway(context.Background(), deleteNatGatewayRequest)
			if error != nil {
				fmt.Printf("Error deleting NatGateway %s %s, It is possible that the resource is already deleted. Please verify manually \n", natGatewayId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &natGatewayId, CoreNatGatewaySweepWaitCondition, time.Duration(3*time.Minute),
				CoreNatGatewaySweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreNatGatewayIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NatGatewayId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listNatGatewaysRequest := oci_core.ListNatGatewaysRequest{}
	listNatGatewaysRequest.CompartmentId = &compartmentId
	listNatGatewaysRequest.LifecycleState = oci_core.NatGatewayLifecycleStateAvailable
	listNatGatewaysResponse, err := virtualNetworkClient.ListNatGateways(context.Background(), listNatGatewaysRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NatGateway list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, natGateway := range listNatGatewaysResponse.Items {
		id := *natGateway.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NatGatewayId", id)
	}
	return resourceIds, nil
}

func CoreNatGatewaySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if natGatewayResponse, ok := response.Response.(oci_core.GetNatGatewayResponse); ok {
		return natGatewayResponse.LifecycleState != oci_core.NatGatewayLifecycleStateTerminated
	}
	return false
}

func CoreNatGatewaySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetNatGateway(context.Background(), oci_core.GetNatGatewayRequest{
		NatGatewayId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
