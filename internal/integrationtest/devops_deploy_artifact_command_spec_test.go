package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DevopsDeployCommandSpecArtifactRequiredOnlyResource = DevopsDeployCommandSpecArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, DevopsDeployCommandSpecArtifactRepresentation)

	DevopsDeployCommandSpecArtifactResourceConfig = DevopsDeployCommandSpecArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Update, DevopsDeployCommandSpecArtifactRepresentation)

	DevopsDeployCommandSpecArtifactSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_artifact_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_artifact.test_deploy_artifact.id}`},
	}

	DevopsDeployCommandSpecArtifactDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_deploy_artifact.test_deploy_artifact.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsDeployCommandSpecArtifactDataSourceFilterRepresentation}}
	DevopsDeployCommandSpecArtifactDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_deploy_artifact.test_deploy_artifact.id}`}},
	}

	DevopsDeployCommandSpecArtifactRepresentation = map[string]interface{}{
		"argument_substitution_mode": acctest.Representation{RepType: acctest.Required, Create: `NONE`, Update: `SUBSTITUTE_PLACEHOLDERS`},
		"deploy_artifact_source":     acctest.RepresentationGroup{RepType: acctest.Required, Group: deployCommandSpecArtifactDeployArtifactSourceRepresentation},
		"deploy_artifact_type":       acctest.Representation{RepType: acctest.Required, Create: `COMMAND_SPEC`},
		"project_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}
	deployCommandSpecArtifactDeployArtifactSourceRepresentation = map[string]interface{}{
		"deploy_artifact_source_type": acctest.Representation{RepType: acctest.Required, Create: `INLINE`},
		"base64encoded_content":       acctest.Representation{RepType: acctest.Required, Create: " base64_encode", Update: "base64_encode_update"},
	}

	DevopsDeployCommandSpecArtifactResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_devops_log_group", acctest.Required, acctest.Create, DevopsLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsDeployCommandSpecArtifactResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployArtifactResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_artifact.test_deploy_artifact"
	datasourceName := "data.oci_devops_deploy_artifacts.test_deploy_artifacts"
	singularDatasourceName := "data.oci_devops_deploy_artifact.test_deploy_artifact"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsDeployCommandSpecArtifactResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Create, DevopsDeployCommandSpecArtifactRepresentation), "devops", "deployArtifact", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployArtifactDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DevopsDeployCommandSpecArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, DevopsDeployCommandSpecArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.base64encoded_content", "base64_encode"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "COMMAND_SPEC"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DevopsDeployCommandSpecArtifactResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DevopsDeployCommandSpecArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Create, DevopsDeployCommandSpecArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.base64encoded_content", "base64_encode"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "COMMAND_SPEC"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DevopsDeployCommandSpecArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Update, DevopsDeployCommandSpecArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "SUBSTITUTE_PLACEHOLDERS"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.base64encoded_content", "base64_encode_update"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "COMMAND_SPEC"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_artifacts", "test_deploy_artifacts", acctest.Optional, acctest.Update, DevopsDeployCommandSpecArtifactDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsDeployCommandSpecArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Update, DevopsDeployCommandSpecArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "deploy_artifact_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "deploy_artifact_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, DevopsDeployCommandSpecArtifactSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsDeployCommandSpecArtifactResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_artifact_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "argument_substitution_mode", "SUBSTITUTE_PLACEHOLDERS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.base64encoded_content", "base64_encode_update"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "INLINE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_type", "COMMAND_SPEC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DevopsDeployCommandSpecArtifactRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
