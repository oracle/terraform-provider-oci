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
	"github.com/oracle/oci-go-sdk/v43/common"
	oci_core "github.com/oracle/oci-go-sdk/v43/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DhcpOptionsRequiredOnlyResource = generateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Required, Create, dhcpOptionsRepresentation)

	dhcpOptionsDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `MyDhcpOptions`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"vcn_id":         Representation{repType: Optional, create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         RepresentationGroup{Required, dhcpOptionsDataSourceFilterRepresentation}}
	dhcpOptionsDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_dhcp_options.test_dhcp_options.id}`}},
	}

	dhcpOptionsRepresentation = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
		"options":          []RepresentationGroup{{Required, optionsRepresentation1}, {Required, optionsRepresentation2}},
		"vcn_id":           Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":     Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     Representation{repType: Optional, create: `MyDhcpOptions`, update: `displayName2`},
		"domain_name_type": Representation{repType: Optional, create: `CUSTOM_DOMAIN`, update: `VCN_DOMAIN`},
		"freeform_tags":    Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}
	optionsRepresentation1 = map[string]interface{}{
		"type":        Representation{repType: Required, create: `DomainNameServer`},
		"server_type": Representation{repType: Required, create: `VcnLocalPlusInternet`},
	}

	optionsRepresentation2 = map[string]interface{}{
		"type":                Representation{repType: Required, create: `SearchDomain`},
		"search_domain_names": Representation{repType: Required, create: []string{"test.com"}},
	}

	DhcpOptionsResourceDependencies = generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

func TestCoreDhcpOptionsResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDhcpOptionsResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_dhcp_options.test_dhcp_options"
	datasourceName := "data.oci_core_dhcp_options.test_dhcp_options"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DhcpOptionsResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Optional, Create, dhcpOptionsRepresentation), "core", "dhcpOptions", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreDhcpOptionsDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DhcpOptionsResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Required, Create, dhcpOptionsRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					ComposeAggregateTestCheckFuncWrapper(
						CheckResourceSetContainsElementWithProperties(resourceName, "options", map[string]string{
							"type":        "DomainNameServer",
							"server_type": "VcnLocalPlusInternet",
						}, []string{}),
					),
					CheckResourceSetContainsElementWithProperties(resourceName, "options", map[string]string{
						"type":                  "SearchDomain",
						"search_domain_names.0": "test.com",
					}, []string{}),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DhcpOptionsResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DhcpOptionsResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Optional, Create, dhcpOptionsRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					// resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyDhcpOptions"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_type", "CUSTOM_DOMAIN"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					ComposeAggregateTestCheckFuncWrapper(
						CheckResourceSetContainsElementWithProperties(resourceName, "options", map[string]string{
							"type":        "DomainNameServer",
							"server_type": "VcnLocalPlusInternet",
						}, []string{}),
					),
					CheckResourceSetContainsElementWithProperties(resourceName, "options", map[string]string{
						"type":                  "SearchDomain",
						"search_domain_names.0": "test.com",
					}, []string{}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DhcpOptionsResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Optional, Create,
						representationCopyWithNewProperties(dhcpOptionsRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					// resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyDhcpOptions"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_type", "CUSTOM_DOMAIN"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "options", map[string]string{
						"server_type": "VcnLocalPlusInternet",
						"type":        "DomainNameServer",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "options", map[string]string{
						"type":                  "SearchDomain",
						"search_domain_names.0": "test.com",
					}, []string{}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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
				Config: config + compartmentIdVariableStr + DhcpOptionsResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Optional, Update, dhcpOptionsRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_type", "VCN_DOMAIN"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					ComposeAggregateTestCheckFuncWrapper(
						CheckResourceSetContainsElementWithProperties(resourceName, "options", map[string]string{
							"type":        "DomainNameServer",
							"server_type": "VcnLocalPlusInternet",
						}, []string{}),
					),
					CheckResourceSetContainsElementWithProperties(resourceName, "options", map[string]string{
						"type":                  "SearchDomain",
						"search_domain_names.0": "test.com",
					}, []string{}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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
					generateDataSourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Optional, Update, dhcpOptionsDataSourceRepresentation) +
					compartmentIdVariableStr + DhcpOptionsResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Optional, Update, dhcpOptionsRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "options.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "options.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "options.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "options.0.domain_name_type", "VCN_DOMAIN"),
					resource.TestCheckResourceAttr(datasourceName, "options.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "options.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "options.0.options.#", "2"),
					resource.TestCheckResourceAttrSet(datasourceName, "options.0.options.0.type"),
					resource.TestCheckResourceAttrSet(datasourceName, "options.0.options.1.type"),
					resource.TestCheckResourceAttrSet(datasourceName, "options.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "options.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "options.0.vcn_id"),
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

func testAccCheckCoreDhcpOptionsDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_dhcp_options" {
			noResourceFound = false
			request := oci_core.GetDhcpOptionsRequest{}

			tmp := rs.Primary.ID
			request.DhcpId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

			response, err := client.GetDhcpOptions(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.DhcpOptionsLifecycleStateTerminated): true,
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
	if !inSweeperExcludeList("CoreDhcpOptions") {
		resource.AddTestSweepers("CoreDhcpOptions", &resource.Sweeper{
			Name:         "CoreDhcpOptions",
			Dependencies: DependencyGraph["dhcpOptions"],
			F:            sweepCoreDhcpOptionsResource,
		})
	}
}

func sweepCoreDhcpOptionsResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	dhcpOptionsIds, err := getDhcpOptionsIds(compartment)
	if err != nil {
		return err
	}
	for _, dhcpOptionsId := range dhcpOptionsIds {
		if ok := SweeperDefaultResourceId[dhcpOptionsId]; !ok {
			deleteDhcpOptionsRequest := oci_core.DeleteDhcpOptionsRequest{}

			deleteDhcpOptionsRequest.DhcpId = &dhcpOptionsId

			deleteDhcpOptionsRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteDhcpOptions(context.Background(), deleteDhcpOptionsRequest)
			if error != nil {
				fmt.Printf("Error deleting DhcpOptions %s %s, It is possible that the resource is already deleted. Please verify manually \n", dhcpOptionsId, error)
				continue
			}
			waitTillCondition(testAccProvider, &dhcpOptionsId, dhcpOptionsSweepWaitCondition, time.Duration(3*time.Minute),
				dhcpOptionsSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getDhcpOptionsIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DhcpOptionsId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listDhcpOptionsRequest := oci_core.ListDhcpOptionsRequest{}
	listDhcpOptionsRequest.CompartmentId = &compartmentId
	listDhcpOptionsRequest.LifecycleState = oci_core.DhcpOptionsLifecycleStateAvailable
	listDhcpOptionsResponse, err := virtualNetworkClient.ListDhcpOptions(context.Background(), listDhcpOptionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DhcpOptions list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dhcpOptions := range listDhcpOptionsResponse.Items {
		id := *dhcpOptions.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "DhcpOptionsId", id)
	}
	return resourceIds, nil
}

func dhcpOptionsSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dhcpOptionsResponse, ok := response.Response.(oci_core.GetDhcpOptionsResponse); ok {
		return dhcpOptionsResponse.LifecycleState != oci_core.DhcpOptionsLifecycleStateTerminated
	}
	return false
}

func dhcpOptionsSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetDhcpOptions(context.Background(), oci_core.GetDhcpOptionsRequest{
		DhcpId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
