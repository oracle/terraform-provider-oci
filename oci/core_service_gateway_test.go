// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

var (
	ServiceGatewayRequiredOnlyResource = ServiceGatewayResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", Required, Create, serviceGatewayRepresentation)

	serviceGatewayDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"vcn_id":         Representation{repType: Optional, create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         RepresentationGroup{Required, serviceGatewayDataSourceFilterRepresentation}}
	serviceGatewayDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_service_gateway.test_service_gateway.id}`}},
	}

	serviceGatewayRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"services":       RepresentationGroup{Required, serviceGatewayServicesRepresentation},
		"vcn_id":         Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}
	serviceGatewayServicesRepresentation = map[string]interface{}{
		"service_id": Representation{repType: Required, create: `${lookup(data.oci_core_services.test_services.services[0], "id")}`},
	}

	ServiceGatewayResourceDependencies = VcnRequiredOnlyResource + VcnResourceDependencies + `
data "oci_core_services" "test_services" {
}
`
)

func TestCoreServiceGatewayResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_service_gateway.test_service_gateway"
	datasourceName := "data.oci_core_service_gateways.test_service_gateways"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreServiceGatewayDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ServiceGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", Required, Create, serviceGatewayRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "services", map[string]string{},
						[]string{
							"service_id",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ServiceGatewayResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ServiceGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", Optional, Create, serviceGatewayRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "block_traffic"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "services", map[string]string{},
						[]string{
							"service_id",
							"service_name",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ServiceGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", Optional, Update, serviceGatewayRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "block_traffic"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "services", map[string]string{},
						[]string{
							"service_id",
							"service_name",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
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
					generateDataSourceFromRepresentationMap("oci_core_service_gateways", "test_service_gateways", Optional, Update, serviceGatewayDataSourceRepresentation) +
					compartmentIdVariableStr + ServiceGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", Optional, Update, serviceGatewayRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "service_gateways.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "service_gateways.0.block_traffic"),
					resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "service_gateways.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.services.#", "1"),
					CheckResourceSetContainsElementWithProperties(datasourceName, "service_gateways.0.services", map[string]string{},
						[]string{
							"service_id",
							"service_name",
						}),
					resource.TestCheckResourceAttrSet(datasourceName, "service_gateways.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "service_gateways.0.vcn_id"),
				),
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

func testAccCheckCoreServiceGatewayDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_service_gateway" {
			noResourceFound = false
			request := oci_core.GetServiceGatewayRequest{}

			tmp := rs.Primary.ID
			request.ServiceGatewayId = &tmp

			response, err := client.GetServiceGateway(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.ServiceGatewayLifecycleStateTerminated): true,
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

func initCoreServiceGatewaySweeper() {
	resource.AddTestSweepers("CoreServiceGateway", &resource.Sweeper{
		Name:         "CoreServiceGateway",
		Dependencies: DependencyGraph["serviceGateway"],
		F:            sweepCoreServiceGatewayResource,
	})
}

func sweepCoreServiceGatewayResource(compartment string) error {
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient

	listServiceGatewaysRequest := oci_core.ListServiceGatewaysRequest{}
	listServiceGatewaysRequest.CompartmentId = &compartmentId
	listServiceGatewaysRequest.LifecycleState = oci_core.ServiceGatewayLifecycleStateAvailable
	listServiceGatewaysResponse, err := virtualNetworkClient.ListServiceGateways(context.Background(), listServiceGatewaysRequest)

	if err != nil {
		return fmt.Errorf("Error getting ServiceGateway list for compartment id : %s , %s \n", compartmentId, err)
	}

	for _, serviceGateway := range listServiceGatewaysResponse.Items {
		if serviceGateway.LifecycleState != oci_core.ServiceGatewayLifecycleStateTerminated {
			log.Printf("deleting serviceGateway %s ", *serviceGateway.Id)

			deleteServiceGatewayRequest := oci_core.DeleteServiceGatewayRequest{}

			deleteServiceGatewayRequest.ServiceGatewayId = serviceGateway.Id

			deleteServiceGatewayRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteServiceGateway(context.Background(), deleteServiceGatewayRequest)
			if error != nil {
				fmt.Printf("Error deleting ServiceGateway %s %s, It is possible that the resource is already deleted. Please verify manually \n", *serviceGateway.Id, error)
				continue
			}

			getServiceGatewayRequest := oci_core.GetServiceGatewayRequest{}

			getServiceGatewayRequest.ServiceGatewayId = serviceGateway.Id

			_, error = virtualNetworkClient.GetServiceGateway(context.Background(), getServiceGatewayRequest)
			if error != nil {
				fmt.Printf("Error retrieving ServiceGateway state %s \n", error)
				continue
			}

			waitTillCondition(testAccProvider, serviceGateway.Id, serviceGatewaySweepWaitCondition, time.Duration(3*time.Minute),
				serviceGatewaySweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func serviceGatewaySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if serviceGatewayResponse, ok := response.Response.(oci_core.GetServiceGatewayResponse); ok {
		return serviceGatewayResponse.LifecycleState == oci_core.ServiceGatewayLifecycleStateTerminated
	}
	return false
}

func serviceGatewaySweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient.GetServiceGateway(context.Background(), oci_core.GetServiceGatewayRequest{
		ServiceGatewayId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
