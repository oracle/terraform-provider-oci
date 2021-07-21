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
	oci_apigateway "github.com/oracle/oci-go-sdk/v45/apigateway"
	"github.com/oracle/oci-go-sdk/v45/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	GatewayRequiredOnlyResource = GatewayResourceDependencies +
		generateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", Required, Create, gatewayRepresentation)

	GatewayResourceConfig = GatewayResourceDependencies +
		generateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", Optional, Update, gatewayRepresentation)

	gatewaySingularDataSourceRepresentation = map[string]interface{}{
		"gateway_id": Representation{repType: Required, create: `${oci_apigateway_gateway.test_gateway.id}`},
	}

	gatewayDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"certificate_id": Representation{repType: Optional, create: `oci_apigateway_certificate.test_certificate.id`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, gatewayDataSourceFilterRepresentation}}
	gatewayDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_apigateway_gateway.test_gateway.id}`}},
	}

	gatewayRepresentation = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"endpoint_type":          Representation{repType: Required, create: `PUBLIC`},
		"subnet_id":              Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"certificate_id":         Representation{repType: Optional, create: `${oci_apigateway_certificate.test_certificate.id}`},
		"defined_tags":           Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":          Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"response_cache_details": RepresentationGroup{Optional, gatewayResponseCacheDetailsRepresentation},
	}
	gatewayResponseCacheDetailsRepresentation = map[string]interface{}{
		"type":                                 Representation{repType: Required, create: `EXTERNAL_RESP_CACHE`},
		"authentication_secret_id":             Representation{repType: Optional, create: `${var.oci_vault_secret_id}`},
		"authentication_secret_version_number": Representation{repType: Optional, create: `1`, update: `2`},
		"connect_timeout_in_ms":                Representation{repType: Optional, create: `10`, update: `11`},
		"is_ssl_enabled":                       Representation{repType: Optional, create: `false`, update: `true`},
		"is_ssl_verify_disabled":               Representation{repType: Optional, create: `false`, update: `true`},
		"read_timeout_in_ms":                   Representation{repType: Optional, create: `10`, update: `11`},
		"send_timeout_in_ms":                   Representation{repType: Optional, create: `10`, update: `11`},
		"servers":                              RepresentationGroup{Optional, gatewayResponseCacheDetailsServersRepresentation},
	}
	gatewayResponseCacheDetailsServersRepresentation = map[string]interface{}{
		"host": Representation{repType: Optional, create: `host`, update: `host2`},
		"port": Representation{repType: Optional, create: `10`, update: `11`},
	}

	GatewayResourceDependencies = generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Required, Create, apiGatewaycertificateRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies +
		apiCertificateVariableStr + apiPrivateKeyVariableStr
)

func TestApigatewayGatewayResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayGatewayResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	vaultSecretId := getEnvSettingWithBlankDefault("oci_vault_secret_id")
	vaultSecretIdStr := fmt.Sprintf("variable \"oci_vault_secret_id\" { default = \"%s\" }\n", vaultSecretId)

	resourceName := "oci_apigateway_gateway.test_gateway"
	datasourceName := "data.oci_apigateway_gateways.test_gateways"
	singularDatasourceName := "data.oci_apigateway_gateway.test_gateway"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+vaultSecretIdStr+GatewayResourceDependencies+
		generateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", Optional, Create, gatewayRepresentation), "apigateway", "gateway", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApigatewayGatewayDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + vaultSecretIdStr + GatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", Required, Create, gatewayRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "endpoint_type", "PUBLIC"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + GatewayResourceDependencies,
				Check: ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						time.Sleep(3 * time.Minute)
						return err
					},
				),
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + vaultSecretIdStr + GatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", Optional, Create, gatewayRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "certificate_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "endpoint_type", "PUBLIC"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "response_cache_details.0.authentication_secret_id"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.authentication_secret_version_number", "1"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.connect_timeout_in_ms", "10"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.is_ssl_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.is_ssl_verify_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.read_timeout_in_ms", "10"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.send_timeout_in_ms", "10"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.servers.0.host", "host"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.servers.0.port", "10"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.type", "EXTERNAL_RESP_CACHE"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + vaultSecretIdStr + GatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", Optional, Create,
						representationCopyWithNewProperties(gatewayRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "certificate_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "endpoint_type", "PUBLIC"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "response_cache_details.0.authentication_secret_id"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.authentication_secret_version_number", "1"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.connect_timeout_in_ms", "10"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.is_ssl_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.is_ssl_verify_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.read_timeout_in_ms", "10"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.send_timeout_in_ms", "10"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.servers.0.host", "host"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.servers.0.port", "10"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.type", "EXTERNAL_RESP_CACHE"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
				Config: config + compartmentIdVariableStr + vaultSecretIdStr + GatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", Optional, Update, gatewayRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "certificate_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "endpoint_type", "PUBLIC"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "response_cache_details.0.authentication_secret_id"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.authentication_secret_version_number", "2"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.connect_timeout_in_ms", "11"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.is_ssl_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.is_ssl_verify_disabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.read_timeout_in_ms", "11"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.send_timeout_in_ms", "11"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.servers.0.host", "host2"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.servers.0.port", "11"),
					resource.TestCheckResourceAttr(resourceName, "response_cache_details.0.type", "EXTERNAL_RESP_CACHE"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
					generateDataSourceFromRepresentationMap("oci_apigateway_gateways", "test_gateways", Optional, Update, gatewayDataSourceRepresentation) +
					compartmentIdVariableStr + vaultSecretIdStr + GatewayResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", Optional, Update, gatewayRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "certificate_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "gateway_collection.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "gateway_collection.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "gateway_collection.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "gateway_collection.0.freeform_tags.%", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", Required, Create, gatewaySingularDataSourceRepresentation) +
					compartmentIdVariableStr + vaultSecretIdStr + GatewayResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "gateway_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "endpoint_type", "PUBLIC"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "hostname"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ip_addresses.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "response_cache_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "response_cache_details.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "response_cache_details.0.authentication_secret_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "response_cache_details.0.authentication_secret_version_number", "2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "response_cache_details.0.connect_timeout_in_ms", "11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "response_cache_details.0.is_ssl_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "response_cache_details.0.is_ssl_verify_disabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "response_cache_details.0.read_timeout_in_ms", "11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "response_cache_details.0.send_timeout_in_ms", "11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "response_cache_details.0.servers.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "response_cache_details.0.servers.0.host", "host2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "response_cache_details.0.servers.0.port", "11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "response_cache_details.0.type", "EXTERNAL_RESP_CACHE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + vaultSecretIdStr + GatewayResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lifecycle_details"},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckApigatewayGatewayDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).gatewayClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apigateway_gateway" {
			noResourceFound = false
			request := oci_apigateway.GetGatewayRequest{}

			tmp := rs.Primary.ID
			request.GatewayId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apigateway")

			response, err := client.GetGateway(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_apigateway.GatewayLifecycleStateDeleted): true,
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
	if !inSweeperExcludeList("ApigatewayGateway") {
		resource.AddTestSweepers("ApigatewayGateway", &resource.Sweeper{
			Name:         "ApigatewayGateway",
			Dependencies: DependencyGraph["gateway"],
			F:            sweepApigatewayGatewayResource,
		})
	}
}

func sweepApigatewayGatewayResource(compartment string) error {
	gatewayClient := GetTestClients(&schema.ResourceData{}).gatewayClient()
	gatewayIds, err := getGatewayIds(compartment)
	if err != nil {
		return err
	}
	for _, gatewayId := range gatewayIds {
		if ok := SweeperDefaultResourceId[gatewayId]; !ok {
			deleteGatewayRequest := oci_apigateway.DeleteGatewayRequest{}

			deleteGatewayRequest.GatewayId = &gatewayId

			deleteGatewayRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apigateway")
			_, error := gatewayClient.DeleteGateway(context.Background(), deleteGatewayRequest)
			if error != nil {
				fmt.Printf("Error deleting Gateway %s %s, It is possible that the resource is already deleted. Please verify manually \n", gatewayId, error)
				continue
			}
			waitTillCondition(testAccProvider, &gatewayId, gatewaySweepWaitCondition, time.Duration(3*time.Minute),
				gatewaySweepResponseFetchOperation, "apigateway", true)
		}
	}
	return nil
}

func getGatewayIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "GatewayId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	gatewayClient := GetTestClients(&schema.ResourceData{}).gatewayClient()

	listGatewaysRequest := oci_apigateway.ListGatewaysRequest{}
	listGatewaysRequest.CompartmentId = &compartmentId
	listGatewaysRequest.LifecycleState = oci_apigateway.GatewayLifecycleStateActive
	listGatewaysResponse, err := gatewayClient.ListGateways(context.Background(), listGatewaysRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Gateway list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, gateway := range listGatewaysResponse.Items {
		id := *gateway.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "GatewayId", id)
	}
	return resourceIds, nil
}

func gatewaySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if gatewayResponse, ok := response.Response.(oci_apigateway.GetGatewayResponse); ok {
		return gatewayResponse.LifecycleState != oci_apigateway.GatewayLifecycleStateDeleted
	}
	return false
}

func gatewaySweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.gatewayClient().GetGateway(context.Background(), oci_apigateway.GetGatewayRequest{
		GatewayId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
