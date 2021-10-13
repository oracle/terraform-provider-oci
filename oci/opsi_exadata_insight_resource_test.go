package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	exadataInsightRequiredRepresentation = map[string]interface{}{
		"compartment_id":                       Representation{RepType: Required, Create: `${var.compartment_id}`},
		"enterprise_manager_bridge_id":         Representation{RepType: Required, Create: `${var.enterprise_manager_bridge_id}`},
		"enterprise_manager_entity_identifier": Representation{RepType: Required, Create: `${var.enterprise_manager_entity_id}`},
		"enterprise_manager_identifier":        Representation{RepType: Required, Create: `${var.enterprise_manager_id}`},
		"status":                               Representation{RepType: Required, Create: `DISABLED`},
		"entity_source":                        Representation{RepType: Required, Create: `EM_MANAGED_EXTERNAL_EXADATA`, Update: `EM_MANAGED_EXTERNAL_EXADATA`},
		"defined_tags":                         Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags":                        Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}},
		"lifecycle":                            RepresentationGroup{Required, ignoreChangesEIRepresentation},
	}

	ignoreChangesEIRepresentation = map[string]interface{}{
		"ignore_changes": Representation{RepType: Required, Create: []string{`defined_tags`}},
	}
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiResourceExadataInsight(t *testing.T) {
	httpreplay.SetScenario("TestOpsiResourceExadataInsight")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	emBridgeId := getEnvSettingWithBlankDefault("enterprise_manager_bridge_ocid")
	emBridgeIdVariableStr := fmt.Sprintf("variable \"enterprise_manager_bridge_id\" { default = \"%s\" }\n", emBridgeId)

	enterpriseManagerId := getEnvSettingWithBlankDefault("enterprise_manager_id")
	enterpriseManagerIdVariableStr := fmt.Sprintf("variable \"enterprise_manager_id\" { default = \"%s\" }\n", enterpriseManagerId)

	enterpriseManagerEntityId := getEnvSettingWithBlankDefault("enterprise_manager_entity_id")
	enterpriseManagerEntityIdVariableStr := fmt.Sprintf("variable \"enterprise_manager_entity_id\" { default = \"%s\" }\n", enterpriseManagerEntityId)

	resourceName := "oci_opsi_exadata_insight.test_exadata_insight"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+emBridgeIdVariableStr+enterpriseManagerIdVariableStr+enterpriseManagerEntityIdVariableStr+ExadataInsightResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", Required, Create, exadataInsightRequiredRepresentation), "opsi", "exadataInsight", t)

	ResourceTest(t, testAccCheckOpsiExadataInsightDestroy, []resource.TestStep{
		// verify Create with Required
		{
			Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + ExadataInsightResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", Required, Create, exadataInsightRequiredRepresentation),

			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),

				func(s *terraform.State) (err error) {
					_, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
