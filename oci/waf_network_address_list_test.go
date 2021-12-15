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
	"github.com/oracle/oci-go-sdk/v54/common"
	oci_waf "github.com/oracle/oci-go-sdk/v54/waf"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NetworkAddressListResourceConfig = NetworkAddressListResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", Optional, Update, networkAddressListRepresentation)

	NetworkAddressListVcnResourceConfig = NetworkAddressListResourceDependenciesVcn +
		GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", Optional, Update, networkAddressListRepresentationVcn)

	networkAddressListSingularDataSourceRepresentation = map[string]interface{}{
		"network_address_list_id": Representation{RepType: Required, Create: `${oci_waf_network_address_list.test_network_address_list.id}`},
	}

	networkAddressListVcnSingularDataSourceRepresentation = map[string]interface{}{
		"network_address_list_id": Representation{RepType: Required, Create: `${oci_waf_network_address_list.test_network_address_list_vcn.id}`},
	}

	networkAddressListVcnDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":             Representation{RepType: Optional, Create: `${oci_waf_network_address_list.test_network_address_list_vcn.id}`},
		"state":          Representation{RepType: Optional, Create: []string{`ACTIVE`}},
		"filter":         RepresentationGroup{Required, networkAddressListVcnDataSourceFilterRepresentation}}

	networkAddressListVcnDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_waf_network_address_list.test_network_address_list_vcn.id}`}},
	}

	networkAddressListDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":             Representation{RepType: Optional, Create: `${oci_waf_network_address_list.test_network_address_list.id}`},
		"state":          Representation{RepType: Optional, Create: []string{`ACTIVE`}},
		"filter":         RepresentationGroup{Required, networkAddressListDataSourceFilterRepresentation}}

	networkAddressListDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_waf_network_address_list.test_network_address_list.id}`}},
	}

	networkAddressListRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"type":           Representation{RepType: Required, Create: `ADDRESSES`},
		"addresses":      Representation{RepType: Required, Create: []string{`10.1.2.3`}, Update: []string{`10.1.2.4`}},
		//"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	networkAddressListRepresentationVcn = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"type":           Representation{RepType: Required, Create: `VCN_ADDRESSES`},
		//"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  Representation{RepType: Optional, Create: `displayName`, Update: "displayName2"},
		"freeform_tags": Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"addresses":     Representation{RepType: Required, Create: []string{`10.1.2.3`}},
		"vcn_addresses": RepresentationGroup{Required, networkAddressListVcnAddressesRepresentation},
	}

	networkAddressListVcnAddressesRepresentation = map[string]interface{}{
		"addresses": Representation{RepType: Required, Create: `10.1.2.3`, Update: `10.1.2.0/24`},
		"vcn_id":    Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	NetworkAddressListResourceDependencies = ""

	NetworkAddressListResourceDependenciesVcn = GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation)
	//+DefinedTagsDependencies
)

// issue-routing-tag: waf/default
func TestWafNetworkAddressListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWafNetworkAddressListResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waf_network_address_list.test_network_address_list"
	datasourceName := "data.oci_waf_network_address_lists.test_network_address_lists"
	singularDatasourceName := "data.oci_waf_network_address_list.test_network_address_list"

	resourceNameVcn := "oci_waf_network_address_list.test_network_address_list_vcn"
	datasourceNameVcn := "data.oci_waf_network_address_lists.test_network_address_lists_vcn"
	singularDatasourceNameVcn := "data.oci_waf_network_address_list.test_network_address_list_vcn"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+NetworkAddressListResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", Optional, Create, networkAddressListRepresentation), "waf", "networkAddressList", t)

	ResourceTest(t, testAccCheckWafNetworkAddressListDestroy, []resource.TestStep{
		// WAF Network Address List VCN_ADDRESSES tests
		// verify Create VCN_Addresses
		{
			Config: config + compartmentIdVariableStr + NetworkAddressListResourceDependenciesVcn +
				GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", Required, Create, networkAddressListRepresentationVcn),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNameVcn, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameVcn, "type", "VCN_ADDRESSES"),
				resource.TestCheckResourceAttr(resourceNameVcn, "vcn_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceNameVcn, "vcn_addresses.0.addresses", "10.1.2.3"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "vcn_addresses.0.vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceNameVcn, "id")
					return err
				},
			),
		},

		// delete VCN_Addresses
		{
			Config: config + compartmentIdVariableStr + NetworkAddressListResourceDependenciesVcn,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkAddressListResourceDependenciesVcn +
				GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", Optional, Create, networkAddressListRepresentationVcn),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNameVcn, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameVcn, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceNameVcn, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "id"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "state"),
				resource.TestCheckResourceAttr(resourceNameVcn, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "time_created"),
				resource.TestCheckResourceAttr(resourceNameVcn, "type", "VCN_ADDRESSES"),
				resource.TestCheckResourceAttr(resourceNameVcn, "vcn_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceNameVcn, "vcn_addresses.0.addresses", "10.1.2.3"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "vcn_addresses.0.vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceNameVcn, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceNameVcn); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NetworkAddressListResourceDependenciesVcn +
				GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", Optional, Create,
					RepresentationCopyWithNewProperties(networkAddressListRepresentationVcn, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNameVcn, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceNameVcn, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceNameVcn, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "id"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "state"),
				resource.TestCheckResourceAttr(resourceNameVcn, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "time_created"),
				resource.TestCheckResourceAttr(resourceNameVcn, "type", "VCN_ADDRESSES"),
				resource.TestCheckResourceAttr(resourceNameVcn, "vcn_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceNameVcn, "vcn_addresses.0.addresses", "10.1.2.3"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "vcn_addresses.0.vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceNameVcn, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NetworkAddressListResourceDependenciesVcn +
				GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", Optional, Update, networkAddressListRepresentationVcn),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNameVcn, "addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceNameVcn, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameVcn, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceNameVcn, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "id"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "state"),
				resource.TestCheckResourceAttr(resourceNameVcn, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "time_created"),
				resource.TestCheckResourceAttr(resourceNameVcn, "type", "VCN_ADDRESSES"),
				resource.TestCheckResourceAttr(resourceNameVcn, "vcn_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceNameVcn, "vcn_addresses.0.addresses", "10.1.2.0/24"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "vcn_addresses.0.vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceNameVcn, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_waf_network_address_lists", "test_network_address_lists_vcn", Optional, Update, networkAddressListVcnDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkAddressListResourceDependenciesVcn +
				GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", Optional, Update, networkAddressListRepresentationVcn),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceNameVcn, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceNameVcn, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceNameVcn, "id"),
				resource.TestCheckResourceAttr(datasourceNameVcn, "state.#", "1"),

				resource.TestCheckResourceAttr(datasourceNameVcn, "network_address_list_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceNameVcn, "network_address_list_collection.0.items.#", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", Required, Create, networkAddressListVcnSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkAddressListVcnResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceNameVcn, "network_address_list_id"),

				resource.TestCheckResourceAttr(singularDatasourceNameVcn, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceNameVcn, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceNameVcn, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameVcn, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameVcn, "state"),
				resource.TestCheckResourceAttr(singularDatasourceNameVcn, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameVcn, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameVcn, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceNameVcn, "type", "VCN_ADDRESSES"),
				resource.TestCheckResourceAttr(singularDatasourceNameVcn, "vcn_addresses.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameVcn, "vcn_addresses.0.addresses", "10.1.2.0/24"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameVcn, "vcn_addresses.0.vcn_id"),
			),
		},

		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + NetworkAddressListVcnResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"addresses"},
			ResourceName:            resourceNameVcn,
		},

		// WAF Network Address List ADDRESSES tests
		// verify Create Addresses
		{
			Config: config + compartmentIdVariableStr + NetworkAddressListResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", Required, Create, networkAddressListRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "type", "ADDRESSES"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete Addresses before next Create
		{
			Config: config + compartmentIdVariableStr + NetworkAddressListResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkAddressListResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", Optional, Create, networkAddressListRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ADDRESSES"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NetworkAddressListResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", Optional, Create,
					RepresentationCopyWithNewProperties(networkAddressListRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ADDRESSES"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NetworkAddressListResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", Optional, Update, networkAddressListRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ADDRESSES"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_waf_network_address_lists", "test_network_address_lists", Optional, Update, networkAddressListDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkAddressListResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", Optional, Update, networkAddressListRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "network_address_list_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "network_address_list_collection.0.items.#", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", Required, Create, networkAddressListSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkAddressListResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_address_list_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "addresses.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "ADDRESSES"),
			),
		},

		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + NetworkAddressListResourceConfig,
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

func testAccCheckWafNetworkAddressListDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).wafClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waf_network_address_list" {
			noResourceFound = false
			request := oci_waf.GetNetworkAddressListRequest{}

			tmp := rs.Primary.ID
			request.NetworkAddressListId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "waf")

			response, err := client.GetNetworkAddressList(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_waf.NetworkAddressListLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !InSweeperExcludeList("WafNetworkAddressList") {
		resource.AddTestSweepers("WafNetworkAddressList", &resource.Sweeper{
			Name:         "WafNetworkAddressList",
			Dependencies: DependencyGraph["networkAddressList"],
			F:            sweepWafNetworkAddressListResource,
		})
	}
}

func sweepWafNetworkAddressListResource(compartment string) error {
	wafClient := GetTestClients(&schema.ResourceData{}).wafClient()
	networkAddressListIds, err := getNetworkAddressListIds(compartment)
	if err != nil {
		return err
	}
	for _, networkAddressListId := range networkAddressListIds {
		if ok := SweeperDefaultResourceId[networkAddressListId]; !ok {
			deleteNetworkAddressListRequest := oci_waf.DeleteNetworkAddressListRequest{}

			deleteNetworkAddressListRequest.NetworkAddressListId = &networkAddressListId

			deleteNetworkAddressListRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "waf")
			_, error := wafClient.DeleteNetworkAddressList(context.Background(), deleteNetworkAddressListRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkAddressList %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkAddressListId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &networkAddressListId, networkAddressListSweepWaitCondition, time.Duration(3*time.Minute),
				networkAddressListSweepResponseFetchOperation, "waf", true)
		}
	}
	return nil
}

func getNetworkAddressListIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "NetworkAddressListId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	wafClient := GetTestClients(&schema.ResourceData{}).wafClient()

	listNetworkAddressListsRequest := oci_waf.ListNetworkAddressListsRequest{}
	listNetworkAddressListsRequest.CompartmentId = &compartmentId
	listNetworkAddressListsRequest.LifecycleState = []oci_waf.NetworkAddressListLifecycleStateEnum{oci_waf.NetworkAddressListLifecycleStateActive}
	listNetworkAddressListsResponse, err := wafClient.ListNetworkAddressLists(context.Background(), listNetworkAddressListsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NetworkAddressList list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, networkAddressList := range listNetworkAddressListsResponse.Items {
		id := *networkAddressList.GetId()
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkAddressListId", id)
	}
	return resourceIds, nil
}

func networkAddressListSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if networkAddressListResponse, ok := response.Response.(oci_waf.GetNetworkAddressListResponse); ok {
		return networkAddressListResponse.GetLifecycleState() != oci_waf.NetworkAddressListLifecycleStateDeleted
	}
	return false
}

func networkAddressListSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.wafClient().GetNetworkAddressList(context.Background(), oci_waf.GetNetworkAddressListRequest{
		NetworkAddressListId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
