// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreDrgRouteTableRouteRuleRequiredOnlyResource = CoreDrgRouteTableRouteRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule", acctest.Required, acctest.Create, CoreDrgRouteTableRouteRuleRepresentation)

	CoreCoreDrgRouteTableRouteRuleDataSourceRepresentation = map[string]interface{}{
		"drg_route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_table.test_drg_route_table.id}`},
	}

	CoreDrgRouteTableRouteRuleRepresentation = map[string]interface{}{
		"drg_route_table_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_table.test_drg_route_table.id}`},
		"destination_type":           acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
		"destination":                acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`, Update: `192.0.0.0/24`},
		"next_hop_drg_attachment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_attachment.test_drg_attachment.id}`},
	}

	CoreDrgRouteTableRouteRuleRepresentation2 = map[string]interface{}{
		"drg_route_table_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_table.test_drg_route_table.id}`},
		"destination_type":           acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
		"destination":                acctest.Representation{RepType: acctest.Required, Create: `1.1.1.0/24`, Update: `1.1.11.0/24`},
		"next_hop_drg_attachment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_attachment.test_drg_attachment.id}`},
	}

	CoreDrgRouteTableRouteRuleRepresentation3 = map[string]interface{}{
		"drg_route_table_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_table.test_drg_route_table.id}`},
		"destination_type":           acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
		"destination":                acctest.Representation{RepType: acctest.Required, Create: `1.1.2.0/24`, Update: `1.1.12.0/24`},
		"next_hop_drg_attachment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_attachment.test_drg_attachment.id}`},
	}

	CoreDrgRouteTableRouteRuleResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Required, acctest.Create, CoreDrgAttachmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Required, acctest.Create, CoreDrgRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, CoreDrgRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, CoreRouteTableRepresentation)
)

// issue-routing-tag: core/pnp
func TestCoreDrgRouteTableRouteRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgRouteTableRouteRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_route_table_route_rule.test_drg_route_table_route_rule"
	datasourceName := "data.oci_core_drg_route_table_route_rules.test_drg_route_table_route_rules"

	resourceName2 := "oci_core_drg_route_table_route_rule.test_drg_route_table_route_rule2"
	resourceName3 := "oci_core_drg_route_table_route_rule.test_drg_route_table_route_rule3"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreDrgRouteTableRouteRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule", acctest.Optional, acctest.Create, CoreDrgRouteTableRouteRuleRepresentation), "core", "drgRouteTableRouteRule", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteTableRouteRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule", acctest.Required, acctest.Create, CoreDrgRouteTableRouteRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteTableRouteRuleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteTableRouteRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule", acctest.Optional, acctest.Create, CoreDrgRouteTableRouteRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "destination"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_type"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "next_hop_drg_attachment_id"),

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
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteTableRouteRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule", acctest.Optional, acctest.Update,
					CoreDrgRouteTableRouteRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "destination", "192.0.0.0/24"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_drg_route_table_route_rules", "test_drg_route_table_route_rules", acctest.Optional, acctest.Create, CoreCoreDrgRouteTableRouteRuleDataSourceRepresentation) +
				compartmentIdVariableStr + CoreDrgRouteTableRouteRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule", acctest.Optional, acctest.Update, CoreDrgRouteTableRouteRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_table_id"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_rules.0.destination_type", "CIDR_BLOCK"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_rules.0.destination", "192.0.0.0/24"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_rules.0.next_hop_drg_attachment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
			),
		},
		//verify resource import
		{
			Config:                  config + CoreDrgRouteTableRouteRuleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteTableRouteRuleResourceDependencies,
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteTableRouteRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule2", acctest.Optional, acctest.Create, CoreDrgRouteTableRouteRuleRepresentation2) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table_route_rule", "test_drg_route_table_route_rule3", acctest.Required, acctest.Create, CoreDrgRouteTableRouteRuleRepresentation3),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
