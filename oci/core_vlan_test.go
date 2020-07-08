// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VlanRequiredOnlyResource = VlanResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", Required, Create, vlanRepresentation)

	VlanResourceConfig = VlanResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", Optional, Update, vlanRepresentation)

	vlanSingularDataSourceRepresentation = map[string]interface{}{
		"vlan_id": Representation{repType: Required, create: `${oci_core_vlan.test_vlan.id}`},
	}

	vlanDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"vcn_id":         Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, vlanDataSourceFilterRepresentation}}
	vlanDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_vlan.test_vlan.id}`}},
	}

	vlanRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"cidr_block":          Representation{repType: Required, create: `10.0.1.0/24`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"vcn_id":              Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"nsg_ids":             Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, update: []string{}},
		"route_table_id":      Representation{repType: Optional, create: `${oci_core_route_table.test_route_table.id}`},
		"vlan_tag":            Representation{repType: Optional, create: `10`},
	}

	VlanResourceDependencies = generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

func TestCoreVlanResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVlanResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_vlan.test_vlan"
	datasourceName := "data.oci_core_vlans.test_vlans"
	singularDatasourceName := "data.oci_core_vlan.test_vlan"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVlanDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VlanResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", Required, Create, vlanRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.1.0/24"),
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
				Config: config + compartmentIdVariableStr + VlanResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + VlanResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", Optional, Create, vlanRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.1.0/24"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttr(resourceName, "vlan_tag", "10"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VlanResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", Optional, Create,
						representationCopyWithNewProperties(vlanRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.1.0/24"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttr(resourceName, "vlan_tag", "10"),

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
				Config: config + compartmentIdVariableStr + VlanResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", Optional, Update, vlanRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.1.0/24"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttr(resourceName, "vlan_tag", "10"),

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
					generateDataSourceFromRepresentationMap("oci_core_vlans", "test_vlans", Optional, Update, vlanDataSourceRepresentation) +
					compartmentIdVariableStr + VlanResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", Optional, Update, vlanRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "vlans.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "vlans.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "vlans.0.cidr_block", "10.0.1.0/24"),
					resource.TestCheckResourceAttr(datasourceName, "vlans.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "vlans.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "vlans.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "vlans.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "vlans.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vlans.0.route_table_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vlans.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "vlans.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "vlans.0.vcn_id"),
					resource.TestCheckResourceAttr(datasourceName, "vlans.0.vlan_tag", "10"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_vlan", "test_vlan", Required, Create, vlanSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VlanResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vlan_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "cidr_block", "10.0.1.0/24"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "vlan_tag", "10"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + VlanResourceConfig,
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

func testAccCheckCoreVlanDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_vlan" {
			noResourceFound = false
			request := oci_core.GetVlanRequest{}

			tmp := rs.Primary.ID
			request.VlanId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

			response, err := client.GetVlan(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VlanLifecycleStateTerminated): true,
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
	if !inSweeperExcludeList("CoreVlan") {
		resource.AddTestSweepers("CoreVlan", &resource.Sweeper{
			Name:         "CoreVlan",
			Dependencies: DependencyGraph["vlan"],
			F:            sweepCoreVlanResource,
		})
	}
}

func sweepCoreVlanResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	vlanIds, err := getVlanIds(compartment)
	if err != nil {
		return err
	}
	for _, vlanId := range vlanIds {
		if ok := SweeperDefaultResourceId[vlanId]; !ok {
			deleteVlanRequest := oci_core.DeleteVlanRequest{}

			deleteVlanRequest.VlanId = &vlanId

			deleteVlanRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteVlan(context.Background(), deleteVlanRequest)
			if error != nil {
				fmt.Printf("Error deleting Vlan %s %s, It is possible that the resource is already deleted. Please verify manually \n", vlanId, error)
				continue
			}
			waitTillCondition(testAccProvider, &vlanId, vlanSweepWaitCondition, time.Duration(3*time.Minute),
				vlanSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVlanIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "VlanId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listVlansRequest := oci_core.ListVlansRequest{}
	listVlansRequest.CompartmentId = &compartmentId

	vcnIds, error := getVcnIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting vcnId required for Vlan resource requests \n")
	}
	for _, vcnId := range vcnIds {
		listVlansRequest.VcnId = &vcnId

		listVlansRequest.LifecycleState = oci_core.VlanLifecycleStateAvailable
		listVlansResponse, err := virtualNetworkClient.ListVlans(context.Background(), listVlansRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Vlan list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, vlan := range listVlansResponse.Items {
			id := *vlan.Id
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "VlanId", id)
		}

	}
	return resourceIds, nil
}

func vlanSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vlanResponse, ok := response.Response.(oci_core.GetVlanResponse); ok {
		return vlanResponse.LifecycleState != oci_core.VlanLifecycleStateTerminated
	}
	return false
}

func vlanSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetVlan(context.Background(), oci_core.GetVlanRequest{
		VlanId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
