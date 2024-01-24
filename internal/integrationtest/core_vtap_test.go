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
	CoreVtapRequiredOnlyResource = CoreVtapResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vtap", "test_vtap", acctest.Required, acctest.Create, CoreVtapRepresentation)

	CoreVtapResourceConfig = CoreVtapResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vtap", "test_vtap", acctest.Optional, acctest.Update, CoreVtapRepresentation)

	CoreCoreVtapSingularDataSourceRepresentation = map[string]interface{}{
		"vtap_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vtap.test_vtap.id}`},
	}

	CoreCoreVtapDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `MyVTAP`, Update: `displayName2`},
		"is_vtap_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"source":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet2.id}`, Update: `${oci_core_subnet.test_subnet.id}`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"target_id":       acctest.Representation{RepType: acctest.Optional, Create: `null`},
		"target_ip":       acctest.Representation{RepType: acctest.Optional, Create: `1.1.1.1`, Update: `2.2.2.2`},
		"vcn_id":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVtapDataSourceFilterRepresentation}}

	CoreVtapDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_vtap.test_vtap.id}`}},
	}

	CoreVtapRepresentation = map[string]interface{}{
		"capture_filter_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_core_capture_filter.test_capture_filter.id}`},
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"source_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet2.id}`, Update: `${oci_core_subnet.test_subnet.id}`}, // TODO
		"vcn_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: `MyVTAP`, Update: `displayName2`},
		"encapsulation_protocol":   acctest.Representation{RepType: acctest.Optional, Create: `VXLAN`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_vtap_enabled":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"max_packet_size":          acctest.Representation{RepType: acctest.Optional, Create: `64`, Update: `75`},
		"source_type":              acctest.Representation{RepType: acctest.Required, Create: `SUBNET`},
		"target_ip":                acctest.Representation{RepType: acctest.Required, Create: `1.1.1.1`, Update: `2.2.2.2`},
		"target_type":              acctest.Representation{RepType: acctest.Optional, Create: `IP_ADDRESS`},
		"traffic_mode":             acctest.Representation{RepType: acctest.Optional, Create: `DEFAULT`, Update: `PRIORITY`},
		"vxlan_network_identifier": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		/* Below fields will be set to null as per service implementation conditionally and depend on other resource creations (may take upto 6 hours to create) hence omitting the check */
		//"source_private_endpoint_ip":        acctest.Representation{RepType: acctest.Optional, Create: `sourcePrivateEndpointIp`, Update: `sourcePrivateEndpointIp2`},
		//"source_private_endpoint_subnet_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		//"target_id":                         acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_target.test_target.id}`},
	}

	CoreVtapResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Required, acctest.Create, CoreCaptureFilterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.1.0/24`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.128.0/24`},
		})) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreVtapResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVtapResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_vtap.test_vtap"
	datasourceName := "data.oci_core_vtaps.test_vtaps"
	singularDatasourceName := "data.oci_core_vtap.test_vtap"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreVtapResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_vtap", "test_vtap", acctest.Optional, acctest.Create, CoreVtapRepresentation), "core", "vtap", t)

	acctest.ResourceTest(t, testAccCheckCoreVtapDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreVtapResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vtap", "test_vtap", acctest.Required, acctest.Create, CoreVtapRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "capture_filter_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreVtapResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreVtapResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vtap", "test_vtap", acctest.Optional, acctest.Create, CoreVtapRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "capture_filter_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyVTAP"),
				resource.TestCheckResourceAttr(resourceName, "encapsulation_protocol", "VXLAN"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_vtap_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_packet_size", "64"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttr(resourceName, "source_type", "SUBNET"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_ip", "1.1.1.1"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "IP_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "traffic_mode", "DEFAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "vxlan_network_identifier", "10"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreVtapResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vtap", "test_vtap", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreVtapRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "capture_filter_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyVTAP"),
				resource.TestCheckResourceAttr(resourceName, "encapsulation_protocol", "VXLAN"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_vtap_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_packet_size", "64"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttr(resourceName, "source_type", "SUBNET"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_ip", "1.1.1.1"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "IP_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "traffic_mode", "DEFAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "vxlan_network_identifier", "10"),

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
			Config: config + compartmentIdVariableStr + CoreVtapResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vtap", "test_vtap", acctest.Optional, acctest.Update, CoreVtapRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "capture_filter_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "encapsulation_protocol", "VXLAN"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_vtap_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_packet_size", "75"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttr(resourceName, "source_type", "SUBNET"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_ip", "2.2.2.2"),
				resource.TestCheckResourceAttr(resourceName, "traffic_mode", "PRIORITY"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "vxlan_network_identifier", "11"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vtaps", "test_vtaps", acctest.Required, acctest.Update, CoreCoreVtapDataSourceRepresentation) +
				compartmentIdVariableStr + CoreVtapResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vtap", "test_vtap", acctest.Optional, acctest.Update, CoreVtapRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "vtaps.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "vtaps.0.capture_filter_id"),
				resource.TestCheckResourceAttr(datasourceName, "vtaps.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "vtaps.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "vtaps.0.encapsulation_protocol", "VXLAN"),
				resource.TestCheckResourceAttr(datasourceName, "vtaps.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "vtaps.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "vtaps.0.is_vtap_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "vtaps.0.max_packet_size", "75"),
				resource.TestCheckResourceAttrSet(datasourceName, "vtaps.0.source_id"),
				resource.TestCheckResourceAttr(datasourceName, "vtaps.0.source_type", "SUBNET"),
				resource.TestCheckResourceAttrSet(datasourceName, "vtaps.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "vtaps.0.target_ip", "2.2.2.2"),
				resource.TestCheckResourceAttrSet(datasourceName, "vtaps.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "vtaps.0.traffic_mode", "PRIORITY"),
				resource.TestCheckResourceAttrSet(datasourceName, "vtaps.0.vcn_id"),
				resource.TestCheckResourceAttr(datasourceName, "vtaps.0.vxlan_network_identifier", "11"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vtap", "test_vtap", acctest.Required, acctest.Create, CoreCoreVtapSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreVtapResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vtap_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "encapsulation_protocol", "VXLAN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_vtap_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_packet_size", "75"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_type", "SUBNET"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_ip", "2.2.2.2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_type", "IP_ADDRESS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "traffic_mode", "PRIORITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vxlan_network_identifier", "11"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreVtapRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreVtapDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_vtap" {
			noResourceFound = false
			request := oci_core.GetVtapRequest{}

			tmp := rs.Primary.ID
			request.VtapId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetVtap(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VtapLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreVtap") {
		resource.AddTestSweepers("CoreVtap", &resource.Sweeper{
			Name:         "CoreVtap",
			Dependencies: acctest.DependencyGraph["vtap"],
			F:            sweepCoreVtapResource,
		})
	}
}

func sweepCoreVtapResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	vtapIds, err := getCoreVtapIds(compartment)
	if err != nil {
		return err
	}
	for _, vtapId := range vtapIds {
		if ok := acctest.SweeperDefaultResourceId[vtapId]; !ok {
			deleteVtapRequest := oci_core.DeleteVtapRequest{}

			deleteVtapRequest.VtapId = &vtapId

			deleteVtapRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteVtap(context.Background(), deleteVtapRequest)
			if error != nil {
				fmt.Printf("Error deleting Vtap %s %s, It is possible that the resource is already deleted. Please verify manually \n", vtapId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &vtapId, CoreVtapSweepWaitCondition, time.Duration(3*time.Minute),
				CoreVtapSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreVtapIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VtapId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listVtapsRequest := oci_core.ListVtapsRequest{}
	listVtapsRequest.CompartmentId = &compartmentId
	listVtapsRequest.LifecycleState = oci_core.VtapLifecycleStateAvailable
	listVtapsResponse, err := virtualNetworkClient.ListVtaps(context.Background(), listVtapsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Vtap list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, vtap := range listVtapsResponse.Items {
		id := *vtap.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VtapId", id)
	}
	return resourceIds, nil
}

func CoreVtapSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vtapResponse, ok := response.Response.(oci_core.GetVtapResponse); ok {
		return vtapResponse.LifecycleState != oci_core.VtapLifecycleStateTerminated
	}
	return false
}

func CoreVtapSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetVtap(context.Background(), oci_core.GetVtapRequest{
		VtapId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
