package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeAuditProfileManagementResource = DataSafeAuditProfileManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile_management", "test_audit_profile_management", acctest.Required, acctest.Create, auditProfileManagementRepresentation)

	auditProfileManagementRepresentation = map[string]interface{}{
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.compartmentId}`},
		"target_id":                            acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"description":                          acctest.Representation{RepType: acctest.Optional, Create: `updated-description`, Update: `description2`},
		"change_retention_trigger":             acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"offline_months":                       acctest.Representation{RepType: acctest.Optional, Create: `5`, Update: `2`},
		"online_months":                        acctest.Representation{RepType: acctest.Optional, Create: `2`, Update: `1`},
		"is_paid_usage_enabled":                acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"is_override_global_retention_setting": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
	}
	DataSafeAuditProfileManagementResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditProfileManagementResource_basic(t *testing.T) {
	t.Skip("Create operation is not available for Audit Profile Management resource")
	httpreplay.SetScenario("TestDataSafeAuditProfileManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_audit_profile_management.test_audit_profile_management"

	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+DataSafeAuditProfileManagementResourceDependencies+targetIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile_management", "test_audit_profile_management", acctest.Optional, acctest.Update,
			acctest.RepresentationCopyWithNewProperties(auditProfileManagementRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "datasafe", "auditProfileManagement", t)
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Update Retention and Paid Usage Setting
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditProfileManagementResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile_management", "test_audit_profile_management", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(auditProfileManagementRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "offline_months"),
				resource.TestCheckResourceAttrSet(resourceName, "online_months"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
			),
		},
	})
}
