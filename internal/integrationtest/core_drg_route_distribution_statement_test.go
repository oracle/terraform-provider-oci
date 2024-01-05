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
	CoreDrgRouteDistributionStatementsRequiredOnlyResource = CoreDrgRouteDistributionStatementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", acctest.Required, acctest.Create, CoreDrgRouteDistributionStatementRepresentation)

	CoreCoreDrgRouteDistributionStatementDataSourceRepresentation = map[string]interface{}{
		"drg_route_distribution_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
	}

	CoreDrgRouteDistributionStatementRepresentation = map[string]interface{}{
		"drg_route_distribution_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
		"action":                    acctest.Representation{RepType: acctest.Required, Create: `ACCEPT`},
		"match_criteria":            acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDrgRouteDistributionStatementStatementsMatchCriteriaRepresentation},
		"priority":                  acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `15`},
	}

	CoreDrgRouteDistributionStatementRepresentation2 = map[string]interface{}{
		"drg_route_distribution_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
		"action":                    acctest.Representation{RepType: acctest.Required, Create: `ACCEPT`},
		"match_criteria":            acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDrgRouteDistributionStatementStatementsMatchCriteriaRepresentation2},
		"priority":                  acctest.Representation{RepType: acctest.Required, Create: `20`, Update: `25`},
	}

	CoreDrgRouteDistributionStatementRepresentation3 = map[string]interface{}{
		"drg_route_distribution_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
		"action":                    acctest.Representation{RepType: acctest.Required, Create: `ACCEPT`},
		"match_criteria":            acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDrgRouteDistributionStatementStatementsMatchCriteriaRepresentation3},
		"priority":                  acctest.Representation{RepType: acctest.Required, Create: `30`, Update: `35`},
	}

	CoreDrgRouteDistributionStatementRepresentation4 = map[string]interface{}{
		"drg_route_distribution_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
		"action":                    acctest.Representation{RepType: acctest.Required, Create: `ACCEPT`},
		"match_criteria":            acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDrgRouteDistributionStatementStatementsMatchCriteriaRepresentation4},
		"priority":                  acctest.Representation{RepType: acctest.Required, Create: `40`, Update: `45`},
	}

	CoreDrgRouteDistributionStatementStatementsMatchCriteriaRepresentation = map[string]interface{}{
		"match_type":      acctest.Representation{RepType: acctest.Required, Create: `DRG_ATTACHMENT_TYPE`},
		"attachment_type": acctest.Representation{RepType: acctest.Required, Create: `VCN`, Update: `VIRTUAL_CIRCUIT`},
	}

	CoreDrgRouteDistributionStatementStatementsMatchCriteriaRepresentation2 = map[string]interface{}{
		"match_type":      acctest.Representation{RepType: acctest.Required, Create: `DRG_ATTACHMENT_TYPE`},
		"attachment_type": acctest.Representation{RepType: acctest.Required, Create: `REMOTE_PEERING_CONNECTION`},
	}

	CoreDrgRouteDistributionStatementStatementsMatchCriteriaRepresentation3 = map[string]interface{}{
		"match_type":        acctest.Representation{RepType: acctest.Required, Create: `DRG_ATTACHMENT_ID`},
		"drg_attachment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_attachment.test_drg_attachment2.id}`},
	}

	CoreDrgRouteDistributionStatementStatementsMatchCriteriaRepresentation4 = map[string]interface{}{
		"match_type": acctest.Representation{RepType: acctest.Required, Create: `MATCH_ALL`},
	}

	CoreDrgRouteDistributionStatementResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment2", acctest.Required, acctest.Create, CoreDrgAttachmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", acctest.Required, acctest.Create, CoreDrgRouteDistributionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, CoreRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, CoreDrgRepresentation)
)

// issue-routing-tag: core/pnp
func TestCoreDrgRouteDistributionStatementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgRouteDistributionStatementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_route_distribution_statement.test_drg_route_distribution_statement"
	datasourceName := "data.oci_core_drg_route_distribution_statements.test_drg_route_distribution_statements"
	resourceName1 := "oci_core_drg_route_distribution_statement.test_drg_route_distribution_statement2"
	resourceName2 := "oci_core_drg_route_distribution_statement.test_drg_route_distribution_statement3"
	resourceName3 := "oci_core_drg_route_distribution_statement.test_drg_route_distribution_statement4"
	resourceName4 := "oci_core_drg_route_distribution_statement.test_drg_route_distribution_statement5"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreDrgRouteDistributionStatementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", acctest.Optional, acctest.Create, CoreDrgRouteDistributionStatementRepresentation), "core", "drgRouteDistributionStatement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteDistributionStatementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", acctest.Optional, acctest.Create,
					CoreDrgRouteDistributionStatementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_distribution_id"),
				resource.TestCheckResourceAttr(resourceName, "action", "ACCEPT"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.0.attachment_type", "VCN"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.0.match_type", "DRG_ATTACHMENT_TYPE"),
				resource.TestCheckResourceAttr(resourceName, "priority", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteDistributionStatementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", acctest.Optional, acctest.Update, CoreDrgRouteDistributionStatementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_distribution_id"),
				resource.TestCheckResourceAttr(resourceName, "action", "ACCEPT"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.0.attachment_type", "VIRTUAL_CIRCUIT"),
				resource.TestCheckResourceAttr(resourceName, "match_criteria.0.match_type", "DRG_ATTACHMENT_TYPE"),
				resource.TestCheckResourceAttr(resourceName, "priority", "15"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_drg_route_distribution_statements", "test_drg_route_distribution_statements", acctest.Optional, acctest.Create, CoreCoreDrgRouteDistributionStatementDataSourceRepresentation) +
				compartmentIdVariableStr + CoreDrgRouteDistributionStatementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", acctest.Optional, acctest.Update, CoreDrgRouteDistributionStatementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_distribution_id"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_distribution_statements.0.action", "ACCEPT"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_distribution_statements.0.match_criteria.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_distribution_statements.0.match_criteria.0.attachment_type", "VIRTUAL_CIRCUIT"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_distribution_statements.0.match_criteria.0.match_type", "DRG_ATTACHMENT_TYPE"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_distribution_statements.0.priority", "15"),
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
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteDistributionStatementResourceDependencies,
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteDistributionStatementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement2", acctest.Optional, acctest.Create, CoreDrgRouteDistributionStatementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement3", acctest.Required, acctest.Create, CoreDrgRouteDistributionStatementRepresentation2) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement4", acctest.Required, acctest.Create, CoreDrgRouteDistributionStatementRepresentation3) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement5", acctest.Required, acctest.Create, CoreDrgRouteDistributionStatementRepresentation4),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//check first resource
				resource.TestCheckResourceAttrSet(resourceName1, "drg_route_distribution_id"),
				resource.TestCheckResourceAttr(resourceName1, "action", "ACCEPT"),
				resource.TestCheckResourceAttr(resourceName1, "match_criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName1, "match_criteria.0.attachment_type", "VCN"),
				resource.TestCheckResourceAttr(resourceName1, "match_criteria.0.match_type", "DRG_ATTACHMENT_TYPE"),
				resource.TestCheckResourceAttr(resourceName1, "priority", "10"),
				resource.TestCheckResourceAttrSet(resourceName1, "id"),
				//check second resource
				resource.TestCheckResourceAttrSet(resourceName2, "drg_route_distribution_id"),
				resource.TestCheckResourceAttr(resourceName2, "action", "ACCEPT"),
				resource.TestCheckResourceAttr(resourceName2, "match_criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName2, "match_criteria.0.attachment_type", "REMOTE_PEERING_CONNECTION"),
				resource.TestCheckResourceAttr(resourceName2, "match_criteria.0.match_type", "DRG_ATTACHMENT_TYPE"),
				resource.TestCheckResourceAttr(resourceName2, "priority", "20"),
				resource.TestCheckResourceAttrSet(resourceName2, "id"),
				// check third resource
				resource.TestCheckResourceAttrSet(resourceName3, "drg_route_distribution_id"),
				resource.TestCheckResourceAttr(resourceName3, "action", "ACCEPT"),
				resource.TestCheckResourceAttr(resourceName3, "match_criteria.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName3, "match_criteria.0.drg_attachment_id"),
				resource.TestCheckResourceAttr(resourceName3, "match_criteria.0.match_type", "DRG_ATTACHMENT_ID"),
				resource.TestCheckResourceAttr(resourceName3, "priority", "30"),
				resource.TestCheckResourceAttrSet(resourceName3, "id"),
				// check fourth resource
				resource.TestCheckResourceAttrSet(resourceName4, "drg_route_distribution_id"),
				resource.TestCheckResourceAttr(resourceName4, "action", "ACCEPT"),
				resource.TestCheckResourceAttr(resourceName4, "match_criteria.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName4, "match_criteria.0.drg_attachment_id"),
				resource.TestCheckResourceAttr(resourceName4, "match_criteria.0.match_type", "MATCH_ALL"),
				resource.TestCheckResourceAttr(resourceName4, "priority", "40"),
				resource.TestCheckResourceAttrSet(resourceName4, "id"),
			),
		},
		// delete
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteDistributionStatementResourceDependencies,
		},
	})
}
