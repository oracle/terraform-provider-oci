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
	VcnRequiredOnlyResource = VcnRequiredOnlyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation)

	VcnResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, vcnRepresentation)

	vcnSingularDataSourceRepresentation = map[string]interface{}{
		"vcn_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	vcnDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: vcnDataSourceFilterRepresentation}}
	vcnDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_vcn.test_vcn.id}`}},
	}

	vcnRepresentation = map[string]interface{}{
		"cidr_block":     acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"dns_label":      acctest.Representation{RepType: acctest.Optional, Create: `dnslabel`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	VcnRequiredOnlyResourceDependencies = ``
	VcnResourceDependencies             = DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreVcnResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVcnResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_vcn.test_vcn"
	datasourceName := "data.oci_core_vcns.test_vcns"
	singularDatasourceName := "data.oci_core_vcn.test_vcn"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+VcnResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, vcnRepresentation), "core", "vcn", t)

	acctest.ResourceTest(t, testAccCheckCoreVcnDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VcnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VcnResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VcnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(vcnRepresentation, []string{"cidr_blocks"}), map[string]interface{}{
						"is_ipv6enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VcnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + VcnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
					"is_ipv6enabled": acctest.Representation{RepType: acctest.Optional, Update: `true`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcns", "test_vcns", acctest.Optional, acctest.Update, vcnDataSourceRepresentation) +
				compartmentIdVariableStr + VcnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
					"is_ipv6enabled": acctest.Representation{RepType: acctest.Optional, Update: `true`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.default_dhcp_options_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.default_route_table_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.default_security_list_id"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.vcn_domain_name"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VcnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
					"is_ipv6enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_dhcp_options_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_route_table_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_security_list_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcn_domain_name"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + VcnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Update, vcnRepresentation),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_ipv6enabled",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckCoreVcnDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_vcn" {
			noResourceFound = false
			request := oci_core.GetVcnRequest{}

			tmp := rs.Primary.ID
			request.VcnId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetVcn(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VcnLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreVcn") {
		resource.AddTestSweepers("CoreVcn", &resource.Sweeper{
			Name:         "CoreVcn",
			Dependencies: acctest.DependencyGraph["vcn"],
			F:            sweepCoreVcnResource,
		})
	}
}

func sweepCoreVcnResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	vcnIds, err := getVcnIds(compartment)
	if err != nil {
		return err
	}
	for _, vcnId := range vcnIds {
		if ok := acctest.SweeperDefaultResourceId[vcnId]; !ok {
			deleteVcnRequest := oci_core.DeleteVcnRequest{}

			deleteVcnRequest.VcnId = &vcnId

			deleteVcnRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteVcn(context.Background(), deleteVcnRequest)
			if error != nil {
				fmt.Printf("Error deleting Vcn %s %s, It is possible that the resource is already deleted. Please verify manually \n", vcnId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &vcnId, vcnSweepWaitCondition, time.Duration(3*time.Minute),
				vcnSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVcnIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VcnId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listVcnsRequest := oci_core.ListVcnsRequest{}
	listVcnsRequest.CompartmentId = &compartmentId
	listVcnsRequest.LifecycleState = oci_core.VcnLifecycleStateAvailable
	listVcnsResponse, err := virtualNetworkClient.ListVcns(context.Background(), listVcnsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Vcn list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, vcn := range listVcnsResponse.Items {
		id := *vcn.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VcnId", id)
		acctest.SweeperDefaultResourceId[*vcn.DefaultDhcpOptionsId] = true
		acctest.SweeperDefaultResourceId[*vcn.DefaultRouteTableId] = true
		acctest.SweeperDefaultResourceId[*vcn.DefaultSecurityListId] = true

	}
	return resourceIds, nil
}

func vcnSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vcnResponse, ok := response.Response.(oci_core.GetVcnResponse); ok {
		return vcnResponse.LifecycleState != oci_core.VcnLifecycleStateTerminated
	}
	return false
}

func vcnSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetVcn(context.Background(), oci_core.GetVcnRequest{
		VcnId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
