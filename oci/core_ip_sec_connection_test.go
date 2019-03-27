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
	IpSecConnectionRequiredOnlyResource = IpSecConnectionResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", Required, Create, ipSecConnectionRepresentation)

	ipSecConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"cpe_id":         Representation{repType: Optional, create: `${oci_core_cpe.test_cpe.id}`},
		"drg_id":         Representation{repType: Optional, create: `${oci_core_drg.test_drg.id}`},
		"filter":         RepresentationGroup{Required, ipSecConnectionDataSourceFilterRepresentation}}
	ipSecConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_ipsec.test_ip_sec_connection.id}`}},
	}

	ipSecConnectionRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"cpe_id":         Representation{repType: Required, create: `${oci_core_cpe.test_cpe.id}`},
		"drg_id":         Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
		"static_routes":  Representation{repType: Required, create: []string{`10.0.0.0/16`}},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `MyIPSecConnection`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	IpSecConnectionResourceDependencies = CpeRequiredOnlyResource + DrgRequiredOnlyResource
)

func TestCoreIpSecConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpSecConnectionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_ipsec.test_ip_sec_connection"
	datasourceName := "data.oci_core_ipsec_connections.test_ip_sec_connections"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreIpSecConnectionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + IpSecConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", Required, Create, ipSecConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cpe_id"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttr(resourceName, "static_routes.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + IpSecConnectionResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + IpSecConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", Optional, Create, ipSecConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cpe_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyIPSecConnection"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "static_routes.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + IpSecConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", Optional, Update, ipSecConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cpe_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "static_routes.#", "1"),

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
					generateDataSourceFromRepresentationMap("oci_core_ipsec_connections", "test_ip_sec_connections", Optional, Update, ipSecConnectionDataSourceRepresentation) +
					compartmentIdVariableStr + IpSecConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", Optional, Update, ipSecConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "cpe_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_id"),

					resource.TestCheckResourceAttr(datasourceName, "connections.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "connections.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "connections.0.cpe_id"),
					resource.TestCheckResourceAttr(datasourceName, "connections.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "connections.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "connections.0.drg_id"),
					resource.TestCheckResourceAttr(datasourceName, "connections.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "connections.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "connections.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "connections.0.static_routes.#", "1"),
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

func testAccCheckCoreIpSecConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_ipsec" {
			noResourceFound = false
			request := oci_core.GetIPSecConnectionRequest{}

			tmp := rs.Primary.ID
			request.IpscId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

			response, err := client.GetIPSecConnection(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.IpSecConnectionLifecycleStateTerminated): true,
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
	resource.AddTestSweepers("CoreIpSecConnection", &resource.Sweeper{
		Name:         "CoreIpSecConnection",
		Dependencies: DependencyGraph["ipSecConnection"],
		F:            sweepCoreIpSecConnectionResource,
	})
}

func sweepCoreIpSecConnectionResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient
	ipSecConnectionIds, err := getIpSecConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, ipSecConnectionId := range ipSecConnectionIds {
		if ok := SweeperDefaultResourceId[ipSecConnectionId]; !ok {
			deleteIPSecConnectionRequest := oci_core.DeleteIPSecConnectionRequest{}

			deleteIPSecConnectionRequest.IpscId = &ipSecConnectionId

			deleteIPSecConnectionRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteIPSecConnection(context.Background(), deleteIPSecConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting IpSecConnection %s %s, It is possible that the resource is already deleted. Please verify manually \n", ipSecConnectionId, error)
				continue
			}
			waitTillCondition(testAccProvider, &ipSecConnectionId, ipSecConnectionSweepWaitCondition, time.Duration(3*time.Minute),
				ipSecConnectionSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getIpSecConnectionIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "IpSecConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient

	listIPSecConnectionsRequest := oci_core.ListIPSecConnectionsRequest{}
	listIPSecConnectionsRequest.CompartmentId = &compartmentId
	listIPSecConnectionsResponse, err := virtualNetworkClient.ListIPSecConnections(context.Background(), listIPSecConnectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting IpSecConnection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, ipSecConnection := range listIPSecConnectionsResponse.Items {
		id := *ipSecConnection.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "IpSecConnectionId", id)
	}
	return resourceIds, nil
}

func ipSecConnectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if ipSecConnectionResponse, ok := response.Response.(oci_core.GetIPSecConnectionResponse); ok {
		return ipSecConnectionResponse.LifecycleState == oci_core.IpSecConnectionLifecycleStateTerminated
	}
	return false
}

func ipSecConnectionSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient.GetIPSecConnection(context.Background(), oci_core.GetIPSecConnectionRequest{
		IpscId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
