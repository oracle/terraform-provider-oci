// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v36/common"
	oci_core "github.com/oracle/oci-go-sdk/v36/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InternetGatewayRequiredOnlyResource = InternetGatewayResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation)

	internetGatewayDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `MyInternetGateway`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"vcn_id":         Representation{repType: Optional, create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         RepresentationGroup{Required, internetGatewayDataSourceFilterRepresentation}}
	internetGatewayDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_internet_gateway.test_internet_gateway.id}`}},
	}

	internetGatewayRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"vcn_id":         Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `MyInternetGateway`, update: `displayName2`},
		"enabled":        Representation{repType: Optional, create: `false`, update: `true`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	InternetGatewayResourceDependencies = generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

func TestCoreInternetGatewayResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInternetGatewayResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_internet_gateway.test_internet_gateway"
	datasourceName := "data.oci_core_internet_gateways.test_internet_gateways"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInternetGatewayDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + InternetGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation),
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
				Config: config + compartmentIdVariableStr + InternetGatewayResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + InternetGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Optional, Create, internetGatewayRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyInternetGateway"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + InternetGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Optional, Create,
						representationCopyWithNewProperties(internetGatewayRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyInternetGateway"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				Config: config + compartmentIdVariableStr + InternetGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Optional, Update, internetGatewayRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
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
					generateDataSourceFromRepresentationMap("oci_core_internet_gateways", "test_internet_gateways", Optional, Update, internetGatewayDataSourceRepresentation) +
					compartmentIdVariableStr + InternetGatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Optional, Update, internetGatewayRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "gateways.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "gateways.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "gateways.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "gateways.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "gateways.0.enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "gateways.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "gateways.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "gateways.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "gateways.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "gateways.0.vcn_id"),
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

func testAccCheckCoreInternetGatewayDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_internet_gateway" {
			noResourceFound = false
			request := oci_core.GetInternetGatewayRequest{}

			tmp := rs.Primary.ID
			request.IgId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

			response, err := client.GetInternetGateway(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.InternetGatewayLifecycleStateTerminated): true,
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
	if !inSweeperExcludeList("CoreInternetGateway") {
		resource.AddTestSweepers("CoreInternetGateway", &resource.Sweeper{
			Name:         "CoreInternetGateway",
			Dependencies: DependencyGraph["internetGateway"],
			F:            sweepCoreInternetGatewayResource,
		})
	}
}

func sweepCoreInternetGatewayResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	internetGatewayIds, err := getInternetGatewayIds(compartment)
	if err != nil {
		return err
	}
	for _, internetGatewayId := range internetGatewayIds {
		if ok := SweeperDefaultResourceId[internetGatewayId]; !ok {
			deleteInternetGatewayRequest := oci_core.DeleteInternetGatewayRequest{}

			deleteInternetGatewayRequest.IgId = &internetGatewayId

			deleteInternetGatewayRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteInternetGateway(context.Background(), deleteInternetGatewayRequest)
			if error != nil {
				fmt.Printf("Error deleting InternetGateway %s %s, It is possible that the resource is already deleted. Please verify manually \n", internetGatewayId, error)
				continue
			}
			waitTillCondition(testAccProvider, &internetGatewayId, internetGatewaySweepWaitCondition, time.Duration(3*time.Minute),
				internetGatewaySweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getInternetGatewayIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "InternetGatewayId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listInternetGatewaysRequest := oci_core.ListInternetGatewaysRequest{}
	listInternetGatewaysRequest.CompartmentId = &compartmentId
	listInternetGatewaysRequest.LifecycleState = oci_core.InternetGatewayLifecycleStateAvailable
	listInternetGatewaysResponse, err := virtualNetworkClient.ListInternetGateways(context.Background(), listInternetGatewaysRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting InternetGateway list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, internetGateway := range listInternetGatewaysResponse.Items {
		id := *internetGateway.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "InternetGatewayId", id)
	}
	return resourceIds, nil
}

func internetGatewaySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if internetGatewayResponse, ok := response.Response.(oci_core.GetInternetGatewayResponse); ok {
		return internetGatewayResponse.LifecycleState != oci_core.InternetGatewayLifecycleStateTerminated
	}
	return false
}

func internetGatewaySweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetInternetGateway(context.Background(), oci_core.GetInternetGatewayRequest{
		IgId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
