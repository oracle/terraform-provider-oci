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
	CoreIpSecConnectionRequiredOnlyResource = CoreIpSecConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Required, acctest.Create, CoreIpSecConnectionRepresentation)

	IpSecConnectionOptionalResource = CoreIpSecConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Optional, acctest.Create, CoreIpSecConnectionRepresentation)

	CoreCoreIpSecConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpe_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_cpe.test_cpe.id}`},
		"drg_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg.test_drg.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreIpSecConnectionDataSourceFilterRepresentation}}
	CoreIpSecConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_ipsec.test_ip_sec_connection.id}`}},
	}

	CoreIpSecConnectionRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpe_id":                    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_cpe.test_cpe.id}`},
		"drg_id":                    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"static_routes":             acctest.Representation{RepType: acctest.Required, Create: []string{`10.0.0.0/16`}, Update: []string{`10.1.0.0/16`}},
		"cpe_local_identifier":      acctest.Representation{RepType: acctest.Optional, Create: `189.44.2.135`, Update: `fakehostname`},
		"cpe_local_identifier_type": acctest.Representation{RepType: acctest.Optional, Create: `IP_ADDRESS`, Update: `HOSTNAME`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `MyIPSecConnection`, Update: `displayName2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTags},
	}

	ignoreDefinedTags = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	CoreIpSecConnectionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", acctest.Required, acctest.Create, CoreCpeRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, CoreDrgRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/default
func TestCoreIpSecConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpSecConnectionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_ipsec.test_ip_sec_connection"
	datasourceName := "data.oci_core_ipsec_connections.test_ip_sec_connections"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreIpSecConnectionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Optional, acctest.Create, CoreIpSecConnectionRepresentation), "core", "ipSecConnection", t)

	acctest.ResourceTest(t, testAccCheckCoreIpSecConnectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreIpSecConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Required, acctest.Create, CoreIpSecConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "cpe_id"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttr(resourceName, "static_routes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tunnel_configuration.#", "0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreIpSecConnectionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreIpSecConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Optional, acctest.Create, CoreIpSecConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "cpe_id"),
				resource.TestCheckResourceAttr(resourceName, "cpe_local_identifier", "189.44.2.135"),
				resource.TestCheckResourceAttr(resourceName, "cpe_local_identifier_type", "IP_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyIPSecConnection"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "static_routes.#", "1"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreIpSecConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreIpSecConnectionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "cpe_id"),
				resource.TestCheckResourceAttr(resourceName, "cpe_local_identifier", "189.44.2.135"),
				resource.TestCheckResourceAttr(resourceName, "cpe_local_identifier_type", "IP_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyIPSecConnection"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "static_routes.#", "1"),

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
			Config: config + compartmentIdVariableStr + CoreIpSecConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Optional, acctest.Update, CoreIpSecConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "cpe_id"),
				resource.TestCheckResourceAttr(resourceName, "cpe_local_identifier", "fakehostname"),
				resource.TestCheckResourceAttr(resourceName, "cpe_local_identifier_type", "HOSTNAME"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "static_routes.#", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connections", "test_ip_sec_connections", acctest.Optional, acctest.Update, CoreCoreIpSecConnectionDataSourceRepresentation) +
				compartmentIdVariableStr + CoreIpSecConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Optional, acctest.Update, CoreIpSecConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "cpe_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_id"),

				resource.TestCheckResourceAttr(datasourceName, "connections.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "connections.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.cpe_id"),
				resource.TestCheckResourceAttr(datasourceName, "connections.0.cpe_local_identifier", "fakehostname"),
				resource.TestCheckResourceAttr(datasourceName, "connections.0.cpe_local_identifier_type", "HOSTNAME"),
				resource.TestCheckResourceAttr(datasourceName, "connections.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.drg_id"),
				resource.TestCheckResourceAttr(datasourceName, "connections.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "connections.0.static_routes.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.transport_type"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreIpSecConnectionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreIpSecConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_ipsec" {
			noResourceFound = false
			request := oci_core.GetIPSecConnectionRequest{}

			tmp := rs.Primary.ID
			request.IpscId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CoreIpSecConnection") {
		resource.AddTestSweepers("CoreIpSecConnection", &resource.Sweeper{
			Name:         "CoreIpSecConnection",
			Dependencies: acctest.DependencyGraph["ipSecConnection"],
			F:            sweepCoreIpSecConnectionResource,
		})
	}
}

func sweepCoreIpSecConnectionResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	ipSecConnectionIds, err := getCoreIpSecConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, ipSecConnectionId := range ipSecConnectionIds {
		if ok := acctest.SweeperDefaultResourceId[ipSecConnectionId]; !ok {
			deleteIPSecConnectionRequest := oci_core.DeleteIPSecConnectionRequest{}

			deleteIPSecConnectionRequest.IpscId = &ipSecConnectionId

			deleteIPSecConnectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteIPSecConnection(context.Background(), deleteIPSecConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting IpSecConnection %s %s, It is possible that the resource is already deleted. Please verify manually \n", ipSecConnectionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &ipSecConnectionId, CoreIpSecConnectionSweepWaitCondition, time.Duration(3*time.Minute),
				CoreIpSecConnectionSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreIpSecConnectionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "IpSecConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listIPSecConnectionsRequest := oci_core.ListIPSecConnectionsRequest{}
	listIPSecConnectionsRequest.CompartmentId = &compartmentId
	listIPSecConnectionsResponse, err := virtualNetworkClient.ListIPSecConnections(context.Background(), listIPSecConnectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting IpSecConnection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, ipSecConnection := range listIPSecConnectionsResponse.Items {
		id := *ipSecConnection.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "IpSecConnectionId", id)
	}
	return resourceIds, nil
}

func CoreIpSecConnectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if ipSecConnectionResponse, ok := response.Response.(oci_core.GetIPSecConnectionResponse); ok {
		return ipSecConnectionResponse.LifecycleState != oci_core.IpSecConnectionLifecycleStateTerminated
	}
	return false
}

func CoreIpSecConnectionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetIPSecConnection(context.Background(), oci_core.GetIPSecConnectionRequest{
		IpscId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
