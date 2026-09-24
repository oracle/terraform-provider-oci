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
	cryptoAssessmentManagementIgnoreChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`, `system_tags`}},
	}

	cryptoAssessmentManagementRepresentation = map[string]interface{}{
		"crypto_assessment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.crypto_assessment_id}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"is_assessment_scheduled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"schedule":                acctest.Representation{RepType: acctest.Optional, Create: `v1; 00 30 15 * *`, Update: `v1; 00 30 20 * *`},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: cryptoAssessmentManagementIgnoreChangesRep},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeCryptoAssessmentManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCryptoAssessmentManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	cryptoAssessmentId := utils.GetEnvSettingWithBlankDefault("crypto_assessment_id")
	cryptoAssessmentIdVariableStr := fmt.Sprintf("variable \"crypto_assessment_id\" { default = \"%s\" }\n", cryptoAssessmentId)

	resourceName := "oci_data_safe_crypto_assessment_management.test_crypto_assessment_management"

	acctest.SaveConfigContent("", "", "", t)
	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + cryptoAssessmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_crypto_assessment_management", "test_crypto_assessment_management", acctest.Optional, acctest.Update, cryptoAssessmentManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
			),
		},
	})
}
