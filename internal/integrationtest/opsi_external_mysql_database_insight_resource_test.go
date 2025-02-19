package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	externalMySqlDatabaseInsightRequiredRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.dbmgmt_external_mysql_database_id}`},
		"database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_connector_id}`},
		"status":                acctest.Representation{RepType: acctest.Required, Create: `DISABLED`},
		"entity_source":         acctest.Representation{RepType: acctest.Required, Create: `EXTERNAL_MYSQL_DATABASE_SYSTEM`, Update: `EXTERNAL_MYSQL_DATABASE_SYSTEM`},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangesExternalMySqlDatabaseInsightRepresentation},
	}
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiResourceExternalMySqlDatabaseInsight(t *testing.T) {
	httpreplay.SetScenario("TestOpsiResourceExternalMySqlDatabaseInsight")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbmgmtExternalMySqlDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_mysql_database_id")
	dbmgmtExternalMySqlDatabaseIdVariableStr := fmt.Sprintf("variable \"dbmgmt_external_mysql_database_id\" { default = \"%s\" }\n", dbmgmtExternalMySqlDatabaseId)

	databaseConnectorId := utils.GetEnvSettingWithBlankDefault("database_connector_id")
	databaseConnectorIdVariableStr := fmt.Sprintf("variable \"database_connector_id\" { default = \"%s\" }\n", databaseConnectorId)

	resourceName := "oci_opsi_database_insight.test_database_insight"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dbmgmtExternalMySqlDatabaseIdVariableStr+databaseConnectorIdVariableStr+ExternalMySqlDatabaseInsightResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, externalMySqlDatabaseInsightRequiredRepresentation), "opsi", "databaseInsight", t)

	acctest.ResourceTest(t, testAccCheckOpsiDatabaseInsightDestroy, []resource.TestStep{
		// verify Create with Required
		{
			Config: config + compartmentIdVariableStr + dbmgmtExternalMySqlDatabaseIdVariableStr + databaseConnectorIdVariableStr + ExternalMySqlDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, externalMySqlDatabaseInsightRequiredRepresentation),
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
			Config:            config + ExternalMySqlDatabaseInsightRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"database_connector_id",
			},
			ResourceName: resourceName,
		},
	})
}
