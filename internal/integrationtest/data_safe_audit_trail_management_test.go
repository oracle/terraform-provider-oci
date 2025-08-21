package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	startAuditTrailManagementRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_ocid}`},
		"target_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"trail_location":              acctest.Representation{RepType: acctest.Optional, Create: `UNIFIED_AUDIT_TRAIL`},
		"audit_collection_start_time": acctest.Representation{RepType: acctest.Required, Create: `2025-01-01T00:00:00Z`, Update: `2025-01-01T00:00:00Z`},
		"start_trigger":               acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreManagementChangesRep},
	}
	stopAuditTrailManagementRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_ocid}`},
		"target_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"trail_location": acctest.Representation{RepType: acctest.Optional, Create: `UNIFIED_AUDIT_TRAIL`},
		"stop_trigger":   acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreManagementChangesRep},
	}
	resumeAuditTrailManagementRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_ocid}`},
		"target_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"trail_location": acctest.Representation{RepType: acctest.Optional, Create: `UNIFIED_AUDIT_TRAIL`},
		"resume_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreManagementChangesRep},
	}
	ignoreManagementChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`, `system_tags`, `audit_collection_start_time`}},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditTrailManagementResource_basic(t *testing.T) {
	t.Skip("Create operation is not available for Audit Trail Management resource")
	httpreplay.SetScenario("TestDataSafeAuditTrailManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_ocid\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_audit_trail_management.test_audit_trail_management"
	acctest.SaveConfigContent("", "", "", t)
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Start trail
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail_management", "test_audit_trail_management", acctest.Optional, acctest.Update, startAuditTrailManagementRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
			),
		},
		// verify Stop trail
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail_management", "test_audit_trail_management", acctest.Optional, acctest.Update, stopAuditTrailManagementRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
			),
		},
		// verify Resume trail
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail_management", "test_audit_trail_management", acctest.Optional, acctest.Update, resumeAuditTrailManagementRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "can_update_last_archive_time_on_target", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
			),
		},
		// verify Stop trail
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail_management", "test_audit_trail_management", acctest.Optional, acctest.Update, stopAuditTrailManagementRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "can_update_last_archive_time_on_target", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
			),
		},
	})
}
