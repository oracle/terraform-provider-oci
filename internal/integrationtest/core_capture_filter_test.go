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
	CoreCaptureFilterRequiredOnlyResource = CoreCaptureFilterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Required, acctest.Create, CoreCaptureFilterRepresentation)

	CoreCaptureFilterResourceConfig = CoreCaptureFilterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Optional, acctest.Update, CoreCaptureFilterRepresentation)

	CoreCoreCaptureFilterSingularDataSourceRepresentation = map[string]interface{}{
		"capture_filter_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_capture_filter.test_capture_filter.id}`},
	}

	CoreCoreCaptureFilterDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `MyCaptureFilter`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreCaptureFilterDataSourceFilterRepresentation}}
	CoreCaptureFilterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_capture_filter.test_capture_filter.id}`}},
	}

	CoreCaptureFilterRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter_type":               acctest.Representation{RepType: acctest.Required, Create: `VTAP`},
		"vtap_capture_filter_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreCaptureFilterVtapCaptureFilterRulesRepresentation},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `MyCaptureFilter`, Update: `displayName2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CoreCaptureFilterVtapCaptureFilterRulesRepresentation = map[string]interface{}{
		"traffic_direction": acctest.Representation{RepType: acctest.Required, Create: `INGRESS`, Update: `EGRESS`},
		"destination_cidr":  acctest.Representation{RepType: acctest.Optional, Create: `10.3.0.0/16`, Update: `10.4.0.0/16`},
		"icmp_options":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreCaptureFilterVtapCaptureFilterRulesIcmpOptionsRepresentation},
		"protocol":          acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"rule_action":       acctest.Representation{RepType: acctest.Optional, Create: `INCLUDE`, Update: `EXCLUDE`},
		"source_cidr":       acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.1/32`, Update: `10.0.0.2/32`},
	}

	CoreCaptureFilterVtapCaptureFilterRulesIcmpOptionsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"code": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	CoreCaptureFilterVtapCaptureFilterRulesTcpOptionsRepresentation = map[string]interface{}{
		"destination_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreCaptureFilterVtapCaptureFilterRulesTcpOptionsDestinationPortRangeRepresentation},
		"source_port_range":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreCaptureFilterVtapCaptureFilterRulesTcpOptionsSourcePortRangeRepresentation},
	}

	CoreCaptureFilterVtapCaptureFilterRulesUdpOptionsRepresentation = map[string]interface{}{
		"destination_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreCaptureFilterVtapCaptureFilterRulesUdpOptionsDestinationPortRangeRepresentation},
		"source_port_range":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreCaptureFilterVtapCaptureFilterRulesUdpOptionsSourcePortRangeRepresentation},
	}

	CoreCaptureFilterVtapCaptureFilterRulesTcpOptionsDestinationPortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}

	CoreCaptureFilterVtapCaptureFilterRulesTcpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}

	CoreCaptureFilterVtapCaptureFilterRulesUdpOptionsDestinationPortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}

	CoreCaptureFilterVtapCaptureFilterRulesUdpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}

	/*
		These are needed only if the protocol is set to UDP. If its set to ICMP/ TCP then these fields will be null

			captureFilterVtapCaptureFilterRulesUdpOptionsRepresentation = map[string]interface{}{
				"destination_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: captureFilterVtapCaptureFilterRulesUdpOptionsDestinationPortRangeRepresentation},
				"source_port_range":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: captureFilterVtapCaptureFilterRulesUdpOptionsSourcePortRangeRepresentation},
			}

			captureFilterVtapCaptureFilterRulesUdpOptionsDestinationPortRangeRepresentation = map[string]interface{}{
				"max": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
				"min": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
			}
			captureFilterVtapCaptureFilterRulesUdpOptionsSourcePortRangeRepresentation = map[string]interface{}{
				"max": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
				"min": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
			}

			captureFilterVtapCaptureFilterRulesTcpOptionsRepresentation = map[string]interface{}{
				"destination_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: captureFilterVtapCaptureFilterRulesTcpOptionsDestinationPortRangeRepresentation},
				"source_port_range":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: captureFilterVtapCaptureFilterRulesTcpOptionsSourcePortRangeRepresentation},
			}

			captureFilterVtapCaptureFilterRulesTcpOptionsDestinationPortRangeRepresentation = map[string]interface{}{
				"max": acctest.Representation{RepType: acctest.Optional, Update: `10`},
				"min": acctest.Representation{RepType: acctest.Optional, Update: `11`},
			}
			captureFilterVtapCaptureFilterRulesTcpOptionsSourcePortRangeRepresentation = map[string]interface{}{
				"max": acctest.Representation{RepType: acctest.Optional, Update: `10`},
				"min": acctest.Representation{RepType: acctest.Optional, Update: `11`},
			}
	*/

	CoreCaptureFilterResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreCaptureFilterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCaptureFilterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_capture_filter.test_capture_filter"
	datasourceName := "data.oci_core_capture_filters.test_capture_filters"
	singularDatasourceName := "data.oci_core_capture_filter.test_capture_filter"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreCaptureFilterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Optional, acctest.Create, CoreCaptureFilterRepresentation), "core", "captureFilter", t)

	acctest.ResourceTest(t, testAccCheckCoreCaptureFilterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreCaptureFilterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Required, acctest.Create, CoreCaptureFilterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "filter_type", "VTAP"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreCaptureFilterResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreCaptureFilterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Optional, acctest.Create, CoreCaptureFilterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyCaptureFilter"),
				resource.TestCheckResourceAttr(resourceName, "filter_type", "VTAP"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.destination_cidr", "10.3.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.icmp_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.icmp_options.0.code", "10"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.icmp_options.0.type", "10"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.protocol", "1"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.rule_action", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.source_cidr", "10.0.0.1/32"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.traffic_direction", "INGRESS"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreCaptureFilterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreCaptureFilterRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyCaptureFilter"),
				resource.TestCheckResourceAttr(resourceName, "filter_type", "VTAP"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.destination_cidr", "10.3.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.icmp_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.icmp_options.0.code", "10"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.icmp_options.0.type", "10"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.protocol", "1"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.rule_action", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.source_cidr", "10.0.0.1/32"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.traffic_direction", "INGRESS"),

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
			Config: config + compartmentIdVariableStr + CoreCaptureFilterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Optional, acctest.Update, CoreCaptureFilterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "filter_type", "VTAP"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.destination_cidr", "10.4.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.icmp_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.icmp_options.0.code", "11"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.icmp_options.0.type", "11"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.protocol", "1"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.rule_action", "EXCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.source_cidr", "10.0.0.2/32"),
				resource.TestCheckResourceAttr(resourceName, "vtap_capture_filter_rules.0.traffic_direction", "EGRESS"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_capture_filters", "test_capture_filters", acctest.Optional, acctest.Update, CoreCoreCaptureFilterDataSourceRepresentation) +
				compartmentIdVariableStr + CoreCaptureFilterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Optional, acctest.Update, CoreCaptureFilterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "capture_filters.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.filter_type", "VTAP"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "capture_filters.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "capture_filters.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "capture_filters.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.vtap_capture_filter_rules.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.vtap_capture_filter_rules.0.destination_cidr", "10.4.0.0/16"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.vtap_capture_filter_rules.0.icmp_options.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.vtap_capture_filter_rules.0.icmp_options.0.code", "11"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.vtap_capture_filter_rules.0.icmp_options.0.type", "11"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.vtap_capture_filter_rules.0.protocol", "1"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.vtap_capture_filter_rules.0.rule_action", "EXCLUDE"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.vtap_capture_filter_rules.0.source_cidr", "10.0.0.2/32"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.vtap_capture_filter_rules.0.traffic_direction", "EGRESS"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Required, acctest.Create, CoreCoreCaptureFilterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreCaptureFilterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "capture_filter_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "filter_type", "VTAP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vtap_capture_filter_rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vtap_capture_filter_rules.0.destination_cidr", "10.4.0.0/16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vtap_capture_filter_rules.0.icmp_options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vtap_capture_filter_rules.0.icmp_options.0.code", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vtap_capture_filter_rules.0.icmp_options.0.type", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vtap_capture_filter_rules.0.protocol", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vtap_capture_filter_rules.0.rule_action", "EXCLUDE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vtap_capture_filter_rules.0.source_cidr", "10.0.0.2/32"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vtap_capture_filter_rules.0.traffic_direction", "EGRESS"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreCaptureFilterRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreCaptureFilterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_capture_filter" {
			noResourceFound = false
			request := oci_core.GetCaptureFilterRequest{}

			tmp := rs.Primary.ID
			request.CaptureFilterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetCaptureFilter(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.CaptureFilterLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreCaptureFilter") {
		resource.AddTestSweepers("CoreCaptureFilter", &resource.Sweeper{
			Name:         "CoreCaptureFilter",
			Dependencies: acctest.DependencyGraph["captureFilter"],
			F:            sweepCoreCaptureFilterResource,
		})
	}
}

func sweepCoreCaptureFilterResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	captureFilterIds, err := getCoreCaptureFilterIds(compartment)
	if err != nil {
		return err
	}
	for _, captureFilterId := range captureFilterIds {
		if ok := acctest.SweeperDefaultResourceId[captureFilterId]; !ok {
			deleteCaptureFilterRequest := oci_core.DeleteCaptureFilterRequest{}

			deleteCaptureFilterRequest.CaptureFilterId = &captureFilterId

			deleteCaptureFilterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteCaptureFilter(context.Background(), deleteCaptureFilterRequest)
			if error != nil {
				fmt.Printf("Error deleting CaptureFilter %s %s, It is possible that the resource is already deleted. Please verify manually \n", captureFilterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &captureFilterId, CoreCaptureFilterSweepWaitCondition, time.Duration(3*time.Minute),
				CoreCaptureFilterSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreCaptureFilterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CaptureFilterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listCaptureFiltersRequest := oci_core.ListCaptureFiltersRequest{}
	listCaptureFiltersRequest.CompartmentId = &compartmentId
	listCaptureFiltersRequest.LifecycleState = oci_core.CaptureFilterLifecycleStateAvailable
	listCaptureFiltersResponse, err := virtualNetworkClient.ListCaptureFilters(context.Background(), listCaptureFiltersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CaptureFilter list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, captureFilter := range listCaptureFiltersResponse.Items {
		id := *captureFilter.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CaptureFilterId", id)
	}
	return resourceIds, nil
}

func CoreCaptureFilterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if captureFilterResponse, ok := response.Response.(oci_core.GetCaptureFilterResponse); ok {
		return captureFilterResponse.LifecycleState != oci_core.CaptureFilterLifecycleStateTerminated
	}
	return false
}

func CoreCaptureFilterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetCaptureFilter(context.Background(), oci_core.GetCaptureFilterRequest{
		CaptureFilterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
