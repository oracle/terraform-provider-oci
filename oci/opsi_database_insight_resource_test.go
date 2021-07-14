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
	databaseInsightRequiredRepresentation = map[string]interface{}{
		"compartment_id":                       Representation{repType: Required, create: `${var.compartment_id}`},
		"enterprise_manager_bridge_id":         Representation{repType: Required, create: `${var.enterprise_manager_bridge_id}`},
		"enterprise_manager_entity_identifier": Representation{repType: Required, create: `${var.enterprise_manager_entity_id}`},
		"enterprise_manager_identifier":        Representation{repType: Required, create: `${var.enterprise_manager_id}`},
		"status":                               Representation{repType: Required, create: `DISABLED`},
		"entity_source":                        Representation{repType: Required, create: `EM_MANAGED_EXTERNAL_DATABASE`, update: `EM_MANAGED_EXTERNAL_DATABASE`},
		"defined_tags":                         Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags":                        Representation{repType: Optional, create: map[string]string{"bar-key": "value"}},
		"lifecycle":                            RepresentationGroup{Required, ignoreChangesDIRepresentation},
	}

	ignoreChangesDIRepresentation = map[string]interface{}{
		"ignore_changes": Representation{repType: Required, create: []string{`defined_tags`}},
	}
)

func TestOpsiResourceDatabaseInsight(t *testing.T) {
	httpreplay.SetScenario("TestOpsiResourceDatabaseInsight")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	emBridgeId := getEnvSettingWithBlankDefault("enterprise_manager_bridge_ocid")
	emBridgeIdVariableStr := fmt.Sprintf("variable \"enterprise_manager_bridge_id\" { default = \"%s\" }\n", emBridgeId)

	enterpriseManagerId := getEnvSettingWithBlankDefault("enterprise_manager_id")
	enterpriseManagerIdVariableStr := fmt.Sprintf("variable \"enterprise_manager_id\" { default = \"%s\" }\n", enterpriseManagerId)

	enterpriseManagerEntityId := getEnvSettingWithBlankDefault("enterprise_manager_entity_id")
	enterpriseManagerEntityIdVariableStr := fmt.Sprintf("variable \"enterprise_manager_entity_id\" { default = \"%s\" }\n", enterpriseManagerEntityId)

	resourceName := "oci_opsi_database_insight.test_database_insight"

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+emBridgeIdVariableStr+enterpriseManagerIdVariableStr+enterpriseManagerEntityIdVariableStr+DatabaseInsightResourceDependencies+
		generateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", Required, Create, databaseInsightRequiredRepresentation), "opsi", "databaseInsight", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckOpsiDatabaseInsightDestroy,
		Steps: []resource.TestStep{
			// verify create with Required
			{
				Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + DatabaseInsightResourceDependencies +
					generateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", Required, Create, databaseInsightRequiredRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
						return err
					},
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
