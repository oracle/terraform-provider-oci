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
	DrgRouteDistributionStatementsDrgAttachmentRequiredOnlyResource = DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies +
		generateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Required, Create, drgRouteDistributionStatementDrgAttachmentIdRepresentation)

	drgRouteDistributionStatementDrgAttachmentIdDataSourceRepresentation = map[string]interface{}{
		"drg_route_distribution_id": Representation{repType: Required, create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
	}

	drgRouteDistributionStatementDrgAttachmentIdRepresentation = map[string]interface{}{
		"drg_route_distribution_id": Representation{repType: Required, create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
		"action":                    Representation{repType: Required, create: `ACCEPT`},
		"match_criteria":            RepresentationGroup{Required, drgRouteDistributionStatementStatementsMatchCriteriaDrgAttachmentIdRepresentation},
		"priority":                  Representation{repType: Required, create: `25`, update: `30`},
	}

	drgRouteDistributionStatementStatementsMatchCriteriaDrgAttachmentIdRepresentation = map[string]interface{}{
		"match_type":        Representation{repType: Required, create: `DRG_ATTACHMENT_ID`},
		"drg_attachment_id": Representation{repType: Required, create: `${oci_core_drg_attachment.test_drg_attachment.id}`},
	}

	DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies = generateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Required, Create, drgAttachmentRepresentation) +
		generateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", Required, Create, drgRouteDistributionRepresentation) +
		generateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation)
)

// issue-routing-tag: core/pnp
func TestCoreDrgRouteDistributionStatementResource_DrgAttachmentId(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgRouteDistributionStatementResource_DrgAttachmentId")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_route_distribution_statement.test_drg_route_distribution_statement"
	datasourceName := "data.oci_core_drg_route_distribution_statements.test_drg_route_distribution_statements"

	var resId, resId2 string
	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies+
		generateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Required, Create, drgRouteDistributionStatementDrgAttachmentIdRepresentation), "core", "drgRouteDistributionStatement", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			//verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Required, Create,
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
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Required, Update, drgRouteDistributionStatementDrgAttachmentIdRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "drg_route_distribution_id"),
					resource.TestCheckResourceAttr(resourceName, "action", "ACCEPT"),
					resource.TestCheckResourceAttr(resourceName, "match_criteria.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "match_criteria.0.drg_attachment_id"),
					resource.TestCheckResourceAttr(resourceName, "match_criteria.0.match_type", "DRG_ATTACHMENT_ID"),
					resource.TestCheckResourceAttr(resourceName, "priority", "30"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_core_drg_route_distribution_statements", "test_drg_route_distribution_statements", Optional, Create, drgRouteDistributionStatementDrgAttachmentIdDataSourceRepresentation) +
					compartmentIdVariableStr + DrgRouteDistributionStatementResourceDrgAttachmentIdDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Optional, Update, drgRouteDistributionStatementDrgAttachmentIdRepresentation),
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
		},
	})
}
