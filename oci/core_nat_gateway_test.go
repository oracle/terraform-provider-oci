// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NatGatewayRequiredOnlyResource = NatGatewayResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", Required, Create, natGatewayRepresentation)

	NatGatewayResourceConfig = NatGatewayResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", Optional, Update, natGatewayRepresentation)

	natGatewaySingularDataSourceRepresentation = map[string]interface{}{
		"nat_gateway_id": Representation{repType: Required, create: `${oci_core_nat_gateway.test_nat_gateway.id}`},
	}

	natGatewayDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"vcn_id":         Representation{repType: Optional, create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         RepresentationGroup{Required, natGatewayDataSourceFilterRepresentation}}
	natGatewayDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_nat_gateway.test_nat_gateway.id}`}},
	}

	natGatewayRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"vcn_id":         Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
		"block_traffic":  Representation{repType: Optional, create: `false`, update: `true`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	NatGatewayResourceDependencies = generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

func TestCoreNatGatewayResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreNatGatewayResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_nat_gateway.test_nat_gateway"
	datasourceName := "data.oci_core_nat_gateways.test_nat_gateways"
	singularDatasourceName := "data.oci_core_nat_gateway.test_nat_gateway"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreNatGatewayDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + NatGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", Required, Create, natGatewayRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + NatGatewayResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + NatGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", Optional, Create, natGatewayRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "block_traffic", "false"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "nat_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NatGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", Optional, Create,
						representationCopyWithNewProperties(natGatewayRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "block_traffic", "false"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "nat_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + NatGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", Optional, Update, natGatewayRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "block_traffic", "true"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "nat_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_core_nat_gateways", "test_nat_gateways", Optional, Update, natGatewayDataSourceRepresentation) +
					compartmentIdVariableStr + NatGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", Optional, Update, natGatewayRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "nat_gateways.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.block_traffic", "true"),
					resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.nat_ip"),
					resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.vcn_id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_nat_gateway", "test_nat_gateway", Required, Create, natGatewaySingularDataSourceRepresentation) +
					compartmentIdVariableStr + NatGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "nat_gateway_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "block_traffic", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "nat_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + NatGatewayResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreNatGatewayDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_nat_gateway" {
			noResourceFound = false
			request := oci_core.GetNatGatewayRequest{}

			tmp := rs.Primary.ID
			request.NatGatewayId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreNatGateway") {
		resource.AddTestSweepers("CoreNatGateway", &resource.Sweeper{
			Name:         "CoreNatGateway",
			Dependencies: DependencyGraph["natGateway"],
			F:            sweepCoreNatGatewayResource,
		})
	}
}

func sweepCoreNatGatewayResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient
	natGatewayIds, err := getNatGatewayIds(compartment)
	if err != nil {
		return err
	}
	for _, natGatewayId := range natGatewayIds {
		if ok := SweeperDefaultResourceId[natGatewayId]; !ok {
			deleteNatGatewayRequest := oci_core.DeleteNatGatewayRequest{}

			deleteNatGatewayRequest.NatGatewayId = &natGatewayId

			deleteNatGatewayRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteNatGateway(context.Background(), deleteNatGatewayRequest)
			if error != nil {
				fmt.Printf("Error deleting NatGateway %s %s, It is possible that the resource is already deleted. Please verify manually \n", natGatewayId, error)
				continue
			}
			waitTillCondition(testAccProvider, &natGatewayId, natGatewaySweepWaitCondition, time.Duration(3*time.Minute),
				natGatewaySweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getNatGatewayIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "NatGatewayId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "NatGatewayId", id)
	}
	return resourceIds, nil
}

func natGatewaySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if natGatewayResponse, ok := response.Response.(oci_core.GetNatGatewayResponse); ok {
		return natGatewayResponse.LifecycleState != oci_core.NatGatewayLifecycleStateTerminated
	}
	return false
}

func natGatewaySweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient.GetNatGateway(context.Background(), oci_core.GetNatGatewayRequest{
		NatGatewayId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
