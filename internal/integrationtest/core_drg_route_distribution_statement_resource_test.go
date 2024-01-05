// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DrgRouteDistributionStatementsDrgAttachmentRequiredOnlyResource = DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", acctest.Required, acctest.Create, drgRouteDistributionStatementDrgAttachmentIdRepresentation)

	drgRouteDistributionStatementDrgAttachmentIdDataSourceRepresentation = map[string]interface{}{
		"drg_route_distribution_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
	}

	drgRouteDistributionStatementDrgAttachmentIdRepresentation = map[string]interface{}{
		"drg_route_distribution_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
		"action":                    acctest.Representation{RepType: acctest.Required, Create: `ACCEPT`},
		"match_criteria":            acctest.RepresentationGroup{RepType: acctest.Required, Group: drgRouteDistributionStatementStatementsMatchCriteriaDrgAttachmentIdRepresentation},
		"priority":                  acctest.Representation{RepType: acctest.Required, Create: `25`, Update: `30`},
	}

	drgRouteDistributionStatementStatementsMatchCriteriaDrgAttachmentIdRepresentation = map[string]interface{}{
		"match_type":        acctest.Representation{RepType: acctest.Required, Create: `DRG_ATTACHMENT_ID`},
		"drg_attachment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_attachment.test_drg_attachment.id}`},
	}

	DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Required, acctest.Create, CoreDrgAttachmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", acctest.Required, acctest.Create, CoreDrgRouteDistributionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, CoreDrgRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, CoreRouteTableRepresentation)
)

// issue-routing-tag: core/pnp
func TestCoreDrgRouteDistributionStatementResource_DrgAttachmentId(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgRouteDistributionStatementResource_DrgAttachmentId")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_route_distribution_statement.test_drg_route_distribution_statement"
	datasourceName := "data.oci_core_drg_route_distribution_statements.test_drg_route_distribution_statements"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", acctest.Required, acctest.Create, drgRouteDistributionStatementDrgAttachmentIdRepresentation), "core", "drgRouteDistributionStatement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", acctest.Required, acctest.Create,
					drgRouteDistributionStatementDrgAttachmentIdRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_distribution_id"),
				resource.TestCheckResourceAttr(resourceName, "action", "ACCEPT"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "match_criteria.0.drg_attachment_id"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.0.match_type", "DRG_ATTACHMENT_ID"),
				resource.TestCheckResourceAttr(resourceName, "priority", "25"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", acctest.Required, acctest.Update, drgRouteDistributionStatementDrgAttachmentIdRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_distribution_id"),
				resource.TestCheckResourceAttr(resourceName, "action", "ACCEPT"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "match_criteria.0.drg_attachment_id"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.0.match_type", "DRG_ATTACHMENT_ID"),
				resource.TestCheckResourceAttr(resourceName, "priority", "30"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_drg_route_distribution_statements", "test_drg_route_distribution_statements", acctest.Optional, acctest.Create, drgRouteDistributionStatementDrgAttachmentIdDataSourceRepresentation) +
				compartmentIdVariableStr + DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", acctest.Optional, acctest.Update, drgRouteDistributionStatementDrgAttachmentIdRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			Config:                  config + CoreDrgRouteDistributionStatementsRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
