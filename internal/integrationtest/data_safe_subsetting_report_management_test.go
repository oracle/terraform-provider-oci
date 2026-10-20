package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var dataSafeSubsettingReportManagementRepresentation = map[string]interface{}{
	"target_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
	"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
	"target_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
		"user_name": acctest.Representation{RepType: acctest.Required, Create: `MASKADMIN28`},
		"password":  acctest.Representation{RepType: acctest.Required, Create: `Maskadminmaskadmin$1`},
	}},
}

func TestDataSafeSubsettingReportManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingReportManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	compartmentID := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	targetID := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	variables := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\nvariable \"target_id\" { default = \"%s\" }\n", compartmentID, targetID)
	resourceName := "oci_data_safe_subsetting_report_management.test_subsetting_report_management"
	policyDependencies := acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Required, acctest.Create, DataSafeSubsettingPolicyRepresentation)
	policyDependencies += acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_subsetting_policy_subsetting_rule", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingRuleRepresentation)
	subsetData := acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subset_data", "test_subset_data", acctest.Required, acctest.Create, dataSafeSubsetDataRepresentation)
	subsetData = strings.TrimSuffix(subsetData, "}\n") + "depends_on = [\"oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule\"]\n}\n"
	reportResource := acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_report_management", "test_subsetting_report_management", acctest.Optional, acctest.Create, dataSafeSubsettingReportManagementRepresentation)
	reportResource = strings.TrimSuffix(reportResource, "}\n") + "depends_on = [\"oci_data_safe_subset_data.test_subset_data\"]\n}\n"

	acctest.ResourceTest(t, nil, []resource.TestStep{{
		Config: config + variables + policyDependencies + subsetData + reportResource,
		Check: acctest.ComposeAggregateTestCheckFuncWrapper(
			resource.TestCheckResourceAttrSet(resourceName, "id"),
			resource.TestCheckResourceAttr(resourceName, "target_id", targetID),
			resource.TestCheckResourceAttrSet(resourceName, "subsetting_policy_id"),
			resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
			resource.TestCheckResourceAttrSet(resourceName, "state"),
			resource.TestCheckResourceAttrSet(resourceName, "subsetting_status"),
			resource.TestCheckResourceAttrSet(resourceName, "total_subsetted_objects"),
		),
	}})
}
