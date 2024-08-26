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
	macsCloudDatabaseInsightRequiredRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connection_credential_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiMacsCloudExadataInsightCredentialDetailsRepresentation},
		"connection_details":            acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiMacsCloudExadataInsightConnectionDetailsRepresentation},
		"database_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.macs_database_id}`},
		"database_resource_type":        acctest.Representation{RepType: acctest.Required, Create: `database`},
		"deployment_type":               acctest.Representation{RepType: acctest.Required, Create: `EXACC`},
		"management_agent_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.management_agent_id}`},
		"status":                        acctest.Representation{RepType: acctest.Required, Create: `DISABLED`},
		"entity_source":                 acctest.Representation{RepType: acctest.Required, Create: `MACS_MANAGED_CLOUD_DATABASE`, Update: `MACS_MANAGED_CLOUD_DATABASE`},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"lifecycle":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesOpsiMacsCloudDatabaseInsightRepresentation},
	}
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiMacsCloudResourceDatabaseInsight(t *testing.T) {
	httpreplay.SetScenario("TestOpsiMacsCloudResourceDatabaseInsight")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentId := utils.GetEnvSettingWithBlankDefault("management_agent_id")
	managementAgentIdVariableStr := fmt.Sprintf("variable \"management_agent_id\" { default = \"%s\" }\n", managementAgentId)

	macsDatabaseId := utils.GetEnvSettingWithBlankDefault("macs_database_id")
	macsDatabaseIdVariableStr := fmt.Sprintf("variable \"macs_database_id\" { default = \"%s\" }\n", macsDatabaseId)

	serviceName := utils.GetEnvSettingWithBlankDefault("service_name")
	serviceNameVariableStr := fmt.Sprintf("variable \"service_name\" { default = \"%s\" }\n", serviceName)

	secretId := utils.GetEnvSettingWithBlankDefault("password_secret_id")
	secretIdVariableStr := fmt.Sprintf("variable \"password_secret_id\" { default = \"%s\" }\n", secretId)

	userName := utils.GetEnvSettingWithBlankDefault("user_name")
	userNamedVariableStr := fmt.Sprintf("variable \"user_name\" { default = \"%s\" }\n", userName)

	dbPort := utils.GetEnvSettingWithBlankDefault("db_port")
	dbPortVariableStr := fmt.Sprintf("variable \"db_port\" { default = \"%s\" }\n", dbPort)

	resourceName := "oci_opsi_database_insight.test_database_insight"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+macsDatabaseIdVariableStr+managementAgentIdVariableStr+serviceNameVariableStr+secretIdVariableStr+userNamedVariableStr+dbPortVariableStr+OpsiMacsCloudDatabaseInsightResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, macsCloudDatabaseInsightRequiredRepresentation), "opsi", "databaseInsight", t)

	acctest.ResourceTest(t, testAccCheckOpsiDatabaseInsightDestroy, []resource.TestStep{
		// verify Create with Required
		{
			Config: config + compartmentIdVariableStr + macsDatabaseIdVariableStr + managementAgentIdVariableStr + serviceNameVariableStr + secretIdVariableStr + userNamedVariableStr + dbPortVariableStr + OpsiMacsCloudDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, macsCloudDatabaseInsightRequiredRepresentation),
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
			Config:            config + OpsiMacsCloudDatabaseInsightRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"deployment_type",
				"service_name",
			},
			ResourceName: resourceName,
		},
	})
}
