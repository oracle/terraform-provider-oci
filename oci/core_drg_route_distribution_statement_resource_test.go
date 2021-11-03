// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DrgRouteDistributionStatementsDrgAttachmentRequiredOnlyResource = DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies +
		GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Required, Create, drgRouteDistributionStatementDrgAttachmentIdRepresentation)

	drgRouteDistributionStatementDrgAttachmentIdDataSourceRepresentation = map[string]interface{}{
		"drg_route_distribution_id": Representation{RepType: Required, Create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
	}

	drgRouteDistributionStatementDrgAttachmentIdRepresentation = map[string]interface{}{
		"drg_route_distribution_id": Representation{RepType: Required, Create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
		"action":                    Representation{RepType: Required, Create: `ACCEPT`},
		"match_criteria":            RepresentationGroup{Required, drgRouteDistributionStatementStatementsMatchCriteriaDrgAttachmentIdRepresentation},
		"priority":                  Representation{RepType: Required, Create: `25`, Update: `30`},
	}

	drgRouteDistributionStatementStatementsMatchCriteriaDrgAttachmentIdRepresentation = map[string]interface{}{
		"match_type":        Representation{RepType: Required, Create: `DRG_ATTACHMENT_ID`},
		"drg_attachment_id": Representation{RepType: Required, Create: `${oci_core_drg_attachment.test_drg_attachment.id}`},
	}

	DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies = GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Required, Create, drgAttachmentRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", Required, Create, drgRouteDistributionRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation)
)

// issue-routing-tag: core/pnp
func TestCoreDrgRouteDistributionStatementResource_DrgAttachmentId(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgRouteDistributionStatementResource_DrgAttachmentId")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_route_distribution_statement.test_drg_route_distribution_statement"
	datasourceName := "data.oci_core_drg_route_distribution_statements.test_drg_route_distribution_statements"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies+
		GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Required, Create, drgRouteDistributionStatementDrgAttachmentIdRepresentation), "core", "drgRouteDistributionStatement", t)

	ResourceTest(t, nil, []resource.TestStep{
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies +
				GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Required, Create,
					drgRouteDistributionStatementDrgAttachmentIdRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_distribution_id"),
				resource.TestCheckResourceAttr(resourceName, "action", "ACCEPT"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "match_criteria.0.drg_attachment_id"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.0.match_type", "DRG_ATTACHMENT_ID"),
				resource.TestCheckResourceAttr(resourceName, "priority", "25"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies +
				GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Required, Update, drgRouteDistributionStatementDrgAttachmentIdRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_distribution_id"),
				resource.TestCheckResourceAttr(resourceName, "action", "ACCEPT"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "match_criteria.0.drg_attachment_id"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.0.match_type", "DRG_ATTACHMENT_ID"),
				resource.TestCheckResourceAttr(resourceName, "priority", "30"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updatedr")
					}
					return err
				},
			),
		},
		//	//verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_drg_route_distribution_statements", "test_drg_route_distribution_statements", Optional, Create, drgRouteDistributionStatementDrgAttachmentIdDataSourceRepresentation) +
				compartmentIdVariableStr + DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies +
				GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Optional, Update, drgRouteDistributionStatementDrgAttachmentIdRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_distribution_id"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_distribution_statements.0.action", "ACCEPT"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_distribution_statements.0.match_criteria.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_distribution_statements.0.match_criteria.0.drg_attachment_id"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_distribution_statements.0.match_criteria.0.match_type", "DRG_ATTACHMENT_ID"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_distribution_statements.0.priority", "30"),
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
	})
}
