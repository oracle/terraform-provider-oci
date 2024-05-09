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
	DataSafeMaskingReportManagementResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_report_management", "test_masking_report_management", acctest.Optional, acctest.Create, DataSafeMaskingReportManagementRepresentation)
	DataSafeMaskingReportManagementRepresentation = map[string]interface{}{
		"target_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"masking_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.masking_policy_id}`},
	}
)

func TestDataSafeMaskingReportManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeMaskingReportManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	maskingPolicyId := utils.GetEnvSettingWithBlankDefault("masking_policy_id")
	maskingPolicyIdVariableStr := fmt.Sprintf("variable \"masking_policy_id\" { default = \"%s\" }\n", maskingPolicyId)

	resourceName := "oci_data_safe_masking_report_management.test_masking_report_management"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + maskingPolicyIdVariableStr + targetIdVariableStr +
				DataSafeMaskingReportManagementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "target_id", targetId),
				resource.TestCheckResourceAttr(resourceName, "masking_policy_id", maskingPolicyId),
				resource.TestCheckResourceAttrSet(resourceName, "is_drop_temp_tables_enabled"),
				resource.TestCheckResourceAttrSet(resourceName, "is_redo_logging_enabled"),
				resource.TestCheckResourceAttrSet(resourceName, "is_refresh_stats_enabled"),
				resource.TestCheckResourceAttrSet(resourceName, "masking_work_request_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parallel_degree"),
				resource.TestCheckResourceAttrSet(resourceName, "recompile"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_masking_finished"),
				resource.TestCheckResourceAttrSet(resourceName, "time_masking_started"),
				resource.TestCheckResourceAttrSet(resourceName, "total_masked_columns"),
				resource.TestCheckResourceAttrSet(resourceName, "total_masked_objects"),
				resource.TestCheckResourceAttrSet(resourceName, "total_masked_schemas"),
				resource.TestCheckResourceAttrSet(resourceName, "total_masked_sensitive_types"),
				resource.TestCheckResourceAttrSet(resourceName, "total_masked_values"),
			),
		},
	})
}
