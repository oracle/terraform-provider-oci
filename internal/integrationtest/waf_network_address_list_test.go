// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_waf "github.com/oracle/oci-go-sdk/v58/waf"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NetworkAddressListResourceConfig = NetworkAddressListResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", acctest.Optional, acctest.Update, networkAddressListRepresentation)

	NetworkAddressListVcnResourceConfig = NetworkAddressListResourceDependenciesVcn +
		acctest.GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", acctest.Optional, acctest.Update, networkAddressListRepresentationVcn)

	networkAddressListSingularDataSourceRepresentation = map[string]interface{}{
		"network_address_list_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waf_network_address_list.test_network_address_list.id}`},
	}

	networkAddressListVcnSingularDataSourceRepresentation = map[string]interface{}{
		"network_address_list_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waf_network_address_list.test_network_address_list_vcn.id}`},
	}

	networkAddressListVcnDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_waf_network_address_list.test_network_address_list_vcn.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: networkAddressListVcnDataSourceFilterRepresentation}}

	networkAddressListVcnDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_waf_network_address_list.test_network_address_list_vcn.id}`}},
	}

	networkAddressListDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_waf_network_address_list.test_network_address_list.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: networkAddressListDataSourceFilterRepresentation}}

	networkAddressListDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_waf_network_address_list.test_network_address_list.id}`}},
	}

	networkAddressListRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `ADDRESSES`},
		"addresses":      acctest.Representation{RepType: acctest.Required, Create: []string{`10.1.2.3`}, Update: []string{`10.1.2.4`}},
		//"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	networkAddressListRepresentationVcn = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `VCN_ADDRESSES`},
		//"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: "displayName2"},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"addresses":     acctest.Representation{RepType: acctest.Required, Create: []string{`10.1.2.3`}},
		"vcn_addresses": acctest.RepresentationGroup{RepType: acctest.Required, Group: networkAddressListVcnAddressesRepresentation},
	}

	networkAddressListVcnAddressesRepresentation = map[string]interface{}{
		"addresses": acctest.Representation{RepType: acctest.Required, Create: `10.1.2.3`, Update: `10.1.2.0/24`},
		"vcn_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	NetworkAddressListResourceDependencies = ""

	NetworkAddressListResourceDependenciesVcn = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation)
	//+DefinedTagsDependencies
)

// issue-routing-tag: waf/default
func TestWafNetworkAddressListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWafNetworkAddressListResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waf_network_address_list.test_network_address_list"
	datasourceName := "data.oci_waf_network_address_lists.test_network_address_lists"
	singularDatasourceName := "data.oci_waf_network_address_list.test_network_address_list"

	resourceNameVcn := "oci_waf_network_address_list.test_network_address_list_vcn"
	datasourceNameVcn := "data.oci_waf_network_address_lists.test_network_address_lists_vcn"
	singularDatasourceNameVcn := "data.oci_waf_network_address_list.test_network_address_list_vcn"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+NetworkAddressListResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", acctest.Optional, acctest.Create, networkAddressListRepresentation), "waf", "networkAddressList", t)

	acctest.ResourceTest(t, testAccCheckWafNetworkAddressListDestroy, []resource.TestStep{
		// WAF Network Address List VCN_ADDRESSES tests
		// verify Create VCN_Addresses
		{
			Config: config + compartmentIdVariableStr + NetworkAddressListResourceDependenciesVcn +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", acctest.Required, acctest.Create, networkAddressListRepresentationVcn),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNameVcn, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameVcn, "type", "VCN_ADDRESSES"),
				resource.TestCheckResourceAttr(resourceNameVcn, "vcn_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceNameVcn, "vcn_addresses.0.addresses", "10.1.2.3"),
				resource.TestCheckResourceAttrSet(resourceNameVcn, "vcn_addresses.0.vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceNameVcn, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", acctest.Optional, acctest.Create, networkAddressListRepresentationVcn),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = acctest.FromInstanceState(s, resourceNameVcn, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceNameVcn); errExport != nil {
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
				acctest.GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(networkAddressListRepresentationVcn, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceNameVcn, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", acctest.Optional, acctest.Update, networkAddressListRepresentationVcn),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceNameVcn, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waf_network_address_lists", "test_network_address_lists_vcn", acctest.Optional, acctest.Update, networkAddressListVcnDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkAddressListResourceDependenciesVcn +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", acctest.Optional, acctest.Update, networkAddressListRepresentationVcn),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list_vcn", acctest.Required, acctest.Create, networkAddressListVcnSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkAddressListVcnResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", acctest.Required, acctest.Create, networkAddressListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "type", "ADDRESSES"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", acctest.Optional, acctest.Create, networkAddressListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NetworkAddressListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(networkAddressListRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + compartmentIdVariableStr + NetworkAddressListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", acctest.Optional, acctest.Update, networkAddressListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waf_network_address_lists", "test_network_address_lists", acctest.Optional, acctest.Update, networkAddressListDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkAddressListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", acctest.Optional, acctest.Update, networkAddressListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waf_network_address_list", "test_network_address_list", acctest.Required, acctest.Create, networkAddressListSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkAddressListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).WafClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waf_network_address_list" {
			noResourceFound = false
			request := oci_waf.GetNetworkAddressListRequest{}

			tmp := rs.Primary.ID
			request.NetworkAddressListId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waf")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("WafNetworkAddressList") {
		resource.AddTestSweepers("WafNetworkAddressList", &resource.Sweeper{
			Name:         "WafNetworkAddressList",
			Dependencies: acctest.DependencyGraph["networkAddressList"],
			F:            sweepWafNetworkAddressListResource,
		})
	}
}

func sweepWafNetworkAddressListResource(compartment string) error {
	wafClient := acctest.GetTestClients(&schema.ResourceData{}).WafClient()
	networkAddressListIds, err := getNetworkAddressListIds(compartment)
	if err != nil {
		return err
	}
	for _, networkAddressListId := range networkAddressListIds {
		if ok := acctest.SweeperDefaultResourceId[networkAddressListId]; !ok {
			deleteNetworkAddressListRequest := oci_waf.DeleteNetworkAddressListRequest{}

			deleteNetworkAddressListRequest.NetworkAddressListId = &networkAddressListId

			deleteNetworkAddressListRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waf")
			_, error := wafClient.DeleteNetworkAddressList(context.Background(), deleteNetworkAddressListRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkAddressList %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkAddressListId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &networkAddressListId, networkAddressListSweepWaitCondition, time.Duration(3*time.Minute),
				networkAddressListSweepResponseFetchOperation, "waf", true)
		}
	}
	return nil
}

func getNetworkAddressListIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkAddressListId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	wafClient := acctest.GetTestClients(&schema.ResourceData{}).WafClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkAddressListId", id)
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

func networkAddressListSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.WafClient().GetNetworkAddressList(context.Background(), oci_waf.GetNetworkAddressListRequest{
		NetworkAddressListId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
