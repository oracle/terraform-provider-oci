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
	DrgRouteDistributionStatementsRequiredOnlyResource = DrgRouteDistributionStatementResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Required, Create, drgRouteDistributionStatementRepresentation)

	drgRouteDistributionStatementDataSourceRepresentation = map[string]interface{}{
		"drg_route_distribution_id": Representation{repType: Required, create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
	}

	drgRouteDistributionStatementRepresentation = map[string]interface{}{
		"drg_route_distribution_id": Representation{repType: Required, create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
		"action":                    Representation{repType: Required, create: `ACCEPT`},
		"match_criteria":            RepresentationGroup{Required, drgRouteDistributionStatementStatementsMatchCriteriaRepresentation},
		"priority":                  Representation{repType: Required, create: `10`, update: `15`},
	}

	drgRouteDistributionStatementStatementsMatchCriteriaRepresentation = map[string]interface{}{
		"match_type":      Representation{repType: Required, create: `DRG_ATTACHMENT_TYPE`},
		"attachment_type": Representation{repType: Required, create: `VCN`, update: `VIRTUAL_CIRCUIT`},
	}

	DrgRouteDistributionStatementResourceDependencies = DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", Required, Create, drgRouteDistributionRepresentation) +
		generateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentation)
)

func TestCoreDrgRouteDistributionStatementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgRouteDistributionStatementResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_route_distribution_statement.test_drg_route_distribution_statement"
	datasourceName := "data.oci_core_drg_route_distribution_statements.test_drg_route_distribution_statements"

	var resId, resId2 string
	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DrgRouteDistributionStatementResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Optional, Create, drgRouteDistributionStatementRepresentation), "core", "drgRouteDistributionStatement", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			//verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DrgRouteDistributionStatementResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Optional, Create,
						drgRouteDistributionStatementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "drg_route_distribution_id"),
					resource.TestCheckResourceAttr(resourceName, "action", "ACCEPT"),
					resource.TestCheckResourceAttr(resourceName, "match_criteria.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "match_criteria.0.attachment_type", "VCN"),
					resource.TestCheckResourceAttr(resourceName, "match_criteria.0.match_type", "DRG_ATTACHMENT_TYPE"),
					resource.TestCheckResourceAttr(resourceName, "priority", "10"),
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
				Config: config + compartmentIdVariableStr + DrgRouteDistributionStatementResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Optional, Update, drgRouteDistributionStatementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "drg_route_distribution_id"),
					resource.TestCheckResourceAttr(resourceName, "action", "ACCEPT"),
					resource.TestCheckResourceAttr(resourceName, "match_criteria.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "match_criteria.0.attachment_type", "VIRTUAL_CIRCUIT"),
					resource.TestCheckResourceAttr(resourceName, "match_criteria.0.match_type", "DRG_ATTACHMENT_TYPE"),
					resource.TestCheckResourceAttr(resourceName, "priority", "15"),
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
			//verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_drg_route_distribution_statements", "test_drg_route_distribution_statements", Optional, Create, drgRouteDistributionStatementDataSourceRepresentation) +
					compartmentIdVariableStr + DrgRouteDistributionStatementResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_route_distribution_statement", "test_drg_route_distribution_statement", Optional, Update, drgRouteDistributionStatementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
