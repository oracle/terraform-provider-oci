// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	databaseInsightRequiredRepresentation = map[string]interface{}{
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"enterprise_manager_bridge_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.enterprise_manager_bridge_id}`},
		"enterprise_manager_entity_identifier": acctest.Representation{RepType: acctest.Required, Create: `${var.enterprise_manager_entity_id}`},
		"enterprise_manager_identifier":        acctest.Representation{RepType: acctest.Required, Create: `${var.enterprise_manager_id}`},
		"status":                               acctest.Representation{RepType: acctest.Required, Create: `DISABLED`},
		"entity_source":                        acctest.Representation{RepType: acctest.Required, Create: `EM_MANAGED_EXTERNAL_DATABASE`, Update: `EM_MANAGED_EXTERNAL_DATABASE`},
		"defined_tags":                         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags":                        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"lifecycle":                            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesDIRepresentation},
	}

	ignoreChangesDIRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiResourceDatabaseInsight(t *testing.T) {
	httpreplay.SetScenario("TestOpsiResourceDatabaseInsight")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	emBridgeId := utils.GetEnvSettingWithBlankDefault("enterprise_manager_bridge_ocid")
	emBridgeIdVariableStr := fmt.Sprintf("variable \"enterprise_manager_bridge_id\" { default = \"%s\" }\n", emBridgeId)

	enterpriseManagerId := utils.GetEnvSettingWithBlankDefault("enterprise_manager_id")
	enterpriseManagerIdVariableStr := fmt.Sprintf("variable \"enterprise_manager_id\" { default = \"%s\" }\n", enterpriseManagerId)

	enterpriseManagerEntityId := utils.GetEnvSettingWithBlankDefault("enterprise_manager_entity_id")
	enterpriseManagerEntityIdVariableStr := fmt.Sprintf("variable \"enterprise_manager_entity_id\" { default = \"%s\" }\n", enterpriseManagerEntityId)

	resourceName := "oci_opsi_database_insight.test_database_insight"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+emBridgeIdVariableStr+enterpriseManagerIdVariableStr+enterpriseManagerEntityIdVariableStr+OpsiDatabaseInsightResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, databaseInsightRequiredRepresentation), "opsi", "databaseInsight", t)

	acctest.ResourceTest(t, testAccCheckOpsiDatabaseInsightDestroy, []resource.TestStep{
		// verify Create with Required
		{
			Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + OpsiDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, databaseInsightRequiredRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify resource import
		{
			Config:                  config + OpsiDatabaseInsightRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
