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
	DataSafeAuditTrailManagementResource = DataSafeAuditTrailManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail_management", "test_audit_trail_management", acctest.Required, acctest.Create, auditTrailManagementRepresentation)

	startAuditTrailManagementRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartmentId}`},
		"target_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"trail_location":              acctest.Representation{RepType: acctest.Optional, Create: `UNIFIED_AUDIT_TRAIL`},
		"audit_collection_start_time": acctest.Representation{RepType: acctest.Optional, Create: `2023-08-01T00:00:00Z`, Update: `2023-08-01T00:00:00Z`},
		"start_trigger":               acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreManagementChangesRep},
	}
	stopAuditTrailManagementRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartmentId}`},
		"target_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"trail_location": acctest.Representation{RepType: acctest.Optional, Create: `UNIFIED_AUDIT_TRAIL`},
		"stop_trigger":   acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreManagementChangesRep},
	}
	resumeAuditTrailManagementRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartmentId}`},
		"target_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"trail_location": acctest.Representation{RepType: acctest.Optional, Create: `UNIFIED_AUDIT_TRAIL`},
		"resume_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreManagementChangesRep},
	}

	auditTrailManagementRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartmentId}`},
		"target_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"trail_location": acctest.Representation{RepType: acctest.Optional, Create: `UNIFIED_AUDIT_TRAIL`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `updated-description`, Update: `description2`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreManagementChangesRep},
	}
	ignoreManagementChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`, `audit_collection_start_time`}},
	}

	DataSafeAuditTrailManagementResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditTrailManagementResource_basic(t *testing.T) {
	t.Skip("Create operation is not available for Audit Trail Management resource")
	httpreplay.SetScenario("TestDataSafeAuditTrailManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_audit_trail_management.test_audit_trail_management"

	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+DataSafeAuditTrailManagementResourceDependencies+targetIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail_management", "test_audit_trail_management", acctest.Optional, acctest.Update,
			acctest.RepresentationCopyWithNewProperties(auditTrailManagementRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "datasafe", "auditTrailManagement", t)
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Start trail
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditTrailManagementResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail_management", "test_audit_trail_management", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(startAuditTrailManagementRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
			),
		},
		// verify Stop trail
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditTrailManagementResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail_management", "test_audit_trail_management", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(stopAuditTrailManagementRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
			),
		},
		// verify Resume trail
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditTrailManagementResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail_management", "test_audit_trail_management", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(resumeAuditTrailManagementRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
			),
		},
		// verify Stop trail
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditTrailManagementResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail_management", "test_audit_trail_management", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(stopAuditTrailManagementRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
			),
		},
	})
}
