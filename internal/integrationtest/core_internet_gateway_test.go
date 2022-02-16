// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	InternetGatewayRequiredOnlyResource = InternetGatewayResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, internetGatewayRepresentation)

	internetGatewayDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `MyInternetGateway`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"vcn_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: internetGatewayDataSourceFilterRepresentation}}
	internetGatewayDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_internet_gateway.test_internet_gateway.id}`}},
	}

	internetGatewayRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `MyInternetGateway`, Update: `displayName2`},
		"enabled":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	InternetGatewayResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreInternetGatewayResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInternetGatewayResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_internet_gateway.test_internet_gateway"
	datasourceName := "data.oci_core_internet_gateways.test_internet_gateways"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+InternetGatewayResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Optional, acctest.Create, internetGatewayRepresentation), "core", "internetGateway", t)

	acctest.ResourceTest(t, testAccCheckCoreInternetGatewayDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + InternetGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, internetGatewayRepresentation),
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
			Config: config + compartmentIdVariableStr + InternetGatewayResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + InternetGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Optional, acctest.Create, internetGatewayRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyInternetGateway"),
				resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + InternetGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(internetGatewayRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyInternetGateway"),
				resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + InternetGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Optional, acctest.Update, internetGatewayRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_internet_gateways", "test_internet_gateways", acctest.Optional, acctest.Update, internetGatewayDataSourceRepresentation) +
				compartmentIdVariableStr + InternetGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Optional, acctest.Update, internetGatewayRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "gateways.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "gateways.0.compartment_id", compartmentId),
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
	})
}

func testAccCheckCoreInternetGatewayDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_internet_gateway" {
			noResourceFound = false
			request := oci_core.GetInternetGatewayRequest{}

			tmp := rs.Primary.ID
			request.IgId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CoreInternetGateway") {
		resource.AddTestSweepers("CoreInternetGateway", &resource.Sweeper{
			Name:         "CoreInternetGateway",
			Dependencies: acctest.DependencyGraph["internetGateway"],
			F:            sweepCoreInternetGatewayResource,
		})
	}
}

func sweepCoreInternetGatewayResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	internetGatewayIds, err := getInternetGatewayIds(compartment)
	if err != nil {
		return err
	}
	for _, internetGatewayId := range internetGatewayIds {
		if ok := acctest.SweeperDefaultResourceId[internetGatewayId]; !ok {
			deleteInternetGatewayRequest := oci_core.DeleteInternetGatewayRequest{}

			deleteInternetGatewayRequest.IgId = &internetGatewayId

			deleteInternetGatewayRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteInternetGateway(context.Background(), deleteInternetGatewayRequest)
			if error != nil {
				fmt.Printf("Error deleting InternetGateway %s %s, It is possible that the resource is already deleted. Please verify manually \n", internetGatewayId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &internetGatewayId, internetGatewaySweepWaitCondition, time.Duration(3*time.Minute),
				internetGatewaySweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getInternetGatewayIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "InternetGatewayId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "InternetGatewayId", id)
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

func internetGatewaySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetInternetGateway(context.Background(), oci_core.GetInternetGatewayRequest{
		IgId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
