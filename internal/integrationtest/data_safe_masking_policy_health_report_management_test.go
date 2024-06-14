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
	DataSafeMaskingPolicyHealthReportManagementResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy_health_report_management", "test_masking_policy_health_report_management", acctest.Optional, acctest.Create, DataSafeMaskingPolicyHealthReportManagementRepresentation)

	ignoreMaskPolicyHealthReportSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DataSafeMaskingPolicyHealthReportManagementRepresentation = map[string]interface{}{
		"target_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"masking_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.masking_policy_id}`},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMaskPolicyHealthReportSystemTagsChangesRep},
	}
)

func TestDataSafeMaskingPolicyHealthReportManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeMaskingPolicyHealthReportManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	maskingPolicyId := utils.GetEnvSettingWithBlankDefault("data_safe_masking_policy_id")
	maskingPolicyIdVariableStr := fmt.Sprintf("variable \"masking_policy_id\" { default = \"%s\" }\n", maskingPolicyId)

	resourceName := "oci_data_safe_masking_policy_health_report_management.test_masking_policy_health_report_management"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + maskingPolicyIdVariableStr + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy_health_report_management", "test_masking_policy_health_report_management", acctest.Optional, acctest.Create, DataSafeMaskingPolicyHealthReportManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "target_id", targetId),
				resource.TestCheckResourceAttr(resourceName, "masking_policy_id", maskingPolicyId),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
			),
		},
	})
}
