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
	autonomousDatabaseInsightRequiredRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.autonomous_database_id}`},
		"database_resource_type":       acctest.Representation{RepType: acctest.Required, Create: `autonomousdatabase`},
		"is_advanced_features_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"credential_details":           acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousDatabaseInsightCredentialDetailsRepresentation},
		"connection_details":           acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousDatabaseInsightConnectionDetailsRepresentation},
		"status":                       acctest.Representation{RepType: acctest.Required, Create: `DISABLED`},
		"entity_source":                acctest.Representation{RepType: acctest.Required, Create: `AUTONOMOUS_DATABASE`, Update: `AUTONOMOUS_DATABASE`},
		//"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		//"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangesADIRepresentation},
	}
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiAutonomousResourceDatabaseInsight(t *testing.T) {
	httpreplay.SetScenario("TestOpsiAutonomousResourceDatabaseInsight")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	autonomousDatabaseId := utils.GetEnvSettingWithBlankDefault("autonomous_database_id")
	autonomousDatabaseIdVariableStr := fmt.Sprintf("variable \"autonomous_database_id\" { default = \"%s\" }\n", autonomousDatabaseId)

	adbHostName := utils.GetEnvSettingWithBlankDefault("adb_host")
	adbHostNameVariableStr := fmt.Sprintf("variable \"adb_host\" { default = \"%s\" }\n", adbHostName)

	adbPort := utils.GetEnvSettingWithBlankDefault("adb_port")
	adbPortVariableStr := fmt.Sprintf("variable \"adb_port\" { default = \"%s\" }\n", adbPort)

	serviceName := utils.GetEnvSettingWithBlankDefault("service_name")
	serviceNameVariableStr := fmt.Sprintf("variable \"service_name\" { default = \"%s\" }\n", serviceName)

	/*secretId := utils.GetEnvSettingWithBlankDefault("secret_id")
	secretIdVariableStr := fmt.Sprintf("variable \"secret_id\" { default = \"%s\" }\n", secretId)

	userName := utils.GetEnvSettingWithBlankDefault("user_name")
	userNamedVariableStr := fmt.Sprintf("variable \"user_name\" { default = \"%s\" }\n", userName)*/

	resourceName := "oci_opsi_database_insight.test_database_insight"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+autonomousDatabaseIdVariableStr+adbHostNameVariableStr+adbPortVariableStr+serviceNameVariableStr+AutonomousDatabaseInsightResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, autonomousDatabaseInsightRequiredRepresentation), "opsi", "databaseInsight", t)

	acctest.ResourceTest(t, testAccCheckOpsiDatabaseInsightDestroy, []resource.TestStep{
		// verify Create with Required
		{
			Config: config + compartmentIdVariableStr + autonomousDatabaseIdVariableStr + adbHostNameVariableStr + adbPortVariableStr + serviceNameVariableStr + AutonomousDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, autonomousDatabaseInsightRequiredRepresentation),
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
			Config:            config + AutonomousDatabaseInsightRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"entity_source",
				"is_advanced_features_enable",
			},
			ResourceName: resourceName,
		},
	})
}
