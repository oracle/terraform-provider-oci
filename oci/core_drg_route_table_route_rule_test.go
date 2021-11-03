// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DrgRouteTableRouteRuleRequiredOnlyResource = DrgRouteTableRouteRuleResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule", Required, Create, drgRouteTableRouteRuleRepresentation)

	drgRouteTableRouteRuleDataSourceRepresentation = map[string]interface{}{
		"drg_route_table_id": Representation{RepType: Required, Create: `${oci_core_drg_route_table.test_drg_route_table.id}`},
	}

	drgRouteTableRouteRuleRepresentation = map[string]interface{}{
		"drg_route_table_id":         Representation{RepType: Required, Create: `${oci_core_drg_route_table.test_drg_route_table.id}`},
		"destination_type":           Representation{RepType: Required, Create: `CIDR_BLOCK`},
		"destination":                Representation{RepType: Required, Create: `0.0.0.0/0`, Update: `192.0.0.0/24`},
		"next_hop_drg_attachment_id": Representation{RepType: Required, Create: `${oci_core_drg_attachment.test_drg_attachment.id}`},
	}

	drgRouteTableRouteRuleRepresentation2 = map[string]interface{}{
		"drg_route_table_id":         Representation{RepType: Required, Create: `${oci_core_drg_route_table.test_drg_route_table.id}`},
		"destination_type":           Representation{RepType: Required, Create: `CIDR_BLOCK`},
		"destination":                Representation{RepType: Required, Create: `1.1.1.0/24`, Update: `1.1.11.0/24`},
		"next_hop_drg_attachment_id": Representation{RepType: Required, Create: `${oci_core_drg_attachment.test_drg_attachment.id}`},
	}

	drgRouteTableRouteRuleRepresentation3 = map[string]interface{}{
		"drg_route_table_id":         Representation{RepType: Required, Create: `${oci_core_drg_route_table.test_drg_route_table.id}`},
		"destination_type":           Representation{RepType: Required, Create: `CIDR_BLOCK`},
		"destination":                Representation{RepType: Required, Create: `1.1.2.0/24`, Update: `1.1.12.0/24`},
		"next_hop_drg_attachment_id": Representation{RepType: Required, Create: `${oci_core_drg_attachment.test_drg_attachment.id}`},
	}

	DrgRouteTableRouteRuleResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Required, Create, drgAttachmentRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Required, Create, drgRouteTableRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation)
)

// issue-routing-tag: core/pnp
func TestCoreDrgRouteTableRouteRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgRouteTableRouteRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_route_table_route_rule.test_drg_route_table_route_rule"
	datasourceName := "data.oci_core_drg_route_table_route_rules.test_drg_route_table_route_rules"

	resourceName2 := "oci_core_drg_route_table_route_rule.test_drg_route_table_route_rule2"
	resourceName3 := "oci_core_drg_route_table_route_rule.test_drg_route_table_route_rule3"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DrgRouteTableRouteRuleResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule", Optional, Create, drgRouteTableRouteRuleRepresentation), "core", "drgRouteTableRouteRule", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DrgRouteTableRouteRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule", Required, Create, drgRouteTableRouteRuleRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DrgRouteTableRouteRuleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DrgRouteTableRouteRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule", Optional, Create, drgRouteTableRouteRuleRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "destination"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_type"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "next_hop_drg_attachment_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DrgRouteTableRouteRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule", Optional, Update,
					drgRouteTableRouteRuleRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "destination", "192.0.0.0/24"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updatedr")
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_drg_route_table_route_rules", "test_drg_route_table_route_rules", Optional, Create, drgRouteTableRouteRuleDataSourceRepresentation) +
				compartmentIdVariableStr + DrgRouteTableRouteRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule", Optional, Update, drgRouteTableRouteRuleRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_table_id"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_rules.0.destination_type", "CIDR_BLOCK"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_rules.0.destination", "192.0.0.0/24"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_rules.0.next_hop_drg_attachment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
			),
		},
		//verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DrgRouteTableRouteRuleResourceDependencies,
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DrgRouteTableRouteRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule2", Optional, Create, drgRouteTableRouteRuleRepresentation2) +
				GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule3", Required, Create, drgRouteTableRouteRuleRepresentation3),
			Check: ComposeAggregateTestCheckFuncWrapper(
				//check first resource
				resource.TestCheckResourceAttrSet(resourceName2, "destination"),
				resource.TestCheckResourceAttrSet(resourceName2, "destination_type"),
				resource.TestCheckResourceAttrSet(resourceName2, "drg_route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName2, "id"),
				resource.TestCheckResourceAttrSet(resourceName2, "next_hop_drg_attachment_id"),
				//check second resource
				resource.TestCheckResourceAttrSet(resourceName3, "destination"),
				resource.TestCheckResourceAttrSet(resourceName3, "destination_type"),
				resource.TestCheckResourceAttrSet(resourceName3, "drg_route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName3, "id"),
				resource.TestCheckResourceAttrSet(resourceName3, "next_hop_drg_attachment_id"),
			),
		},
	})
}
