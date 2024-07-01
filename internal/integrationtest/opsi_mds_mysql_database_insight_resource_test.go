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
	mdsMySqlDatabaseInsightRequiredRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.mds_mysql_database_id}`},
		"database_resource_type": acctest.Representation{RepType: acctest.Required, Create: `mysqldbsystem`},
		"status":                 acctest.Representation{RepType: acctest.Required, Create: `DISABLED`},
		"entity_source":          acctest.Representation{RepType: acctest.Required, Create: `MDS_MYSQL_DATABASE_SYSTEM`, Update: `MDS_MYSQL_DATABASE_SYSTEM`},
		//"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		//"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangesADIRepresentation},
	}

	ignoreChangesADIRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiResourceMdsMySqlDatabaseInsight(t *testing.T) {
	httpreplay.SetScenario("TestOpsiResourceMdsMySqlDatabaseInsight")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	mdsMySqlDatabaseId := utils.GetEnvSettingWithBlankDefault("mds_mysql_database_id")
	mdsMySqlDatabaseIdDatabaseIdVariableStr := fmt.Sprintf("variable \"mds_mysql_database_id\" { default = \"%s\" }\n", mdsMySqlDatabaseId)

	resourceName := "oci_opsi_database_insight.test_database_insight"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+mdsMySqlDatabaseIdDatabaseIdVariableStr+MdsMySqlDatabaseInsightResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, mdsMySqlDatabaseInsightRequiredRepresentation), "opsi", "databaseInsight", t)

	acctest.ResourceTest(t, testAccCheckOpsiDatabaseInsightDestroy, []resource.TestStep{
		// verify Create with Required
		{
			Config: config + compartmentIdVariableStr + mdsMySqlDatabaseIdDatabaseIdVariableStr + MdsMySqlDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, mdsMySqlDatabaseInsightRequiredRepresentation),
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
			Config:                  config + MdsMySqlDatabaseInsightRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
