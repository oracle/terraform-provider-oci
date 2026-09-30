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

var dataSafeSubsetDataRepresentation = map[string]interface{}{
	"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
	"target_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
	"target_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
		"user_name": acctest.Representation{RepType: acctest.Required, Create: `MASKADMIN28`},
		"password":  acctest.Representation{RepType: acctest.Required, Create: `Maskadminmaskadmin$1`},
	}},
}

func TestDataSafeSubsetDataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsetDataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	compartmentID := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	targetID := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")

	variables := fmt.Sprintf(`
variable "compartment_id" {
	default = "%s"
}

variable "target_id" {
	default = "%s"
}
`, compartmentID, targetID)
	dependencies := acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Required, acctest.Create, DataSafeSubsettingPolicyRepresentation)
	dependencies += acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_subsetting_policy_subsetting_rule", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingRuleRepresentation)
	subsetResource := acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subset_data", "test_subset_data", acctest.Required, acctest.Create, dataSafeSubsetDataRepresentation)
	subsetResource = strings.TrimSuffix(subsetResource, "}\n") + "depends_on = [\"oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule\"]\n}\n"
	resourceName := "oci_data_safe_subset_data.test_subset_data"

	acctest.ResourceTest(t, nil, []resource.TestStep{{
		Config: config + variables + dependencies + subsetResource,
		Check:  resource.TestCheckResourceAttrSet(resourceName, "id"),
	}})
}
