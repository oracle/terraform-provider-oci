// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v49/cloudguard"
	"github.com/oracle/oci-go-sdk/v49/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DetectorRecipeRequiredOnlyResource = DetectorRecipeResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_cloud_guard_detector_recipe", "test_detector_recipe", Required, Create, detectorRecipeRepresentation)

	DetectorRecipeResourceConfig = DetectorRecipeResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_cloud_guard_detector_recipe", "test_detector_recipe", Optional, Update, detectorRecipeRepresentation)

	detectorRecipeSingularDataSourceRepresentation = map[string]interface{}{
		"detector_recipe_id": Representation{RepType: Required, Create: `${oci_cloud_guard_detector_recipe.test_detector_recipe.id}`},
	}

	detectorRecipeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"access_level":              Representation{RepType: Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": Representation{RepType: Optional, Create: `true`},
		"display_name":              Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"resource_metadata_only":    Representation{RepType: Optional, Create: `false`},
		"state":                     Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":                    RepresentationGroup{Required, detectorRecipeDataSourceFilterRepresentation}}
	detectorRecipeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_cloud_guard_detector_recipe.test_detector_recipe.id}`}},
	}

	//Making a list call and getting a source detectorRecipeId
	detectorRecipeId             = `${data.oci_cloud_guard_detector_recipes.oracle_detector_recipe.detector_recipe_collection.0.items.0.id}`
	detectorRecipeRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":              Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"source_detector_recipe_id": Representation{RepType: Required, Create: detectorRecipeId},
		"defined_tags":              Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":               Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"detector_rules":            RepresentationGroup{Optional, detectorRecipeDetectorRulesRepresentation},
		"freeform_tags":             Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}
	//Configurations and Conditions are dependent on the detectorRuleId selected, hence hardcoding one for testing purposes
	detectorRecipeDetectorRulesRepresentation = map[string]interface{}{
		"details":          RepresentationGroup{Required, detectorRecipeDetectorRulesDetailsRepresentation},
		"detector_rule_id": Representation{RepType: Required, Create: `LB_CERTIFICATE_EXPIRING_SOON`},
	}
	detectorRecipeDetectorRulesDetailsRepresentation = map[string]interface{}{
		"is_enabled": Representation{RepType: Required, Create: `false`, Update: `true`},
		//Only valid riskLevels allowed
		"risk_level": Representation{RepType: Required, Create: `CRITICAL`, Update: `LOW`},
		//Making a valid condition Object
		"condition":      Representation{RepType: Optional, Create: `{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"EQUALS\",\"value\":\"10\",\"valueType\":\"CUSTOM\"}`, Update: `{\"kind\":\"COMPOSITE\",\"leftOperand\":{\"kind\" :\"COMPOSITE\",\"leftOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"EQUALS\",\"value\":\"12\",\"valueType\":\"CUSTOM\"},\"compositeOperator\":\"AND\",\"rightOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"NOT_EQUALS\",\"value\":\"12\",\"valueType\":\"CUSTOM\"}},\"compositeOperator\":\"AND\",\"rightOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"EQUALS\",\"value\":\"10\",\"valueType\":\"CUSTOM\"}}`},
		"configurations": RepresentationGroup{Optional, detectorRecipeDetectorRulesDetailsConfigurationsRepresentation},
		"labels":         Representation{RepType: Optional, Create: []string{`labels`}, Update: []string{`labels2`}},
	}
	//Making a valid configuration Object
	detectorRecipeDetectorRulesDetailsConfigurationsRepresentation = map[string]interface{}{
		"config_key": Representation{RepType: Required, Create: `lbCertificateExpiringSoonConfig`, Update: `lbCertificateExpiringSoonConfig`},
		"name":       Representation{RepType: Required, Create: `Days before expiring - 1`, Update: `Days before expiring - 2`},
		"data_type":  Representation{RepType: Optional, Create: `multiList`, Update: `multiList`},
		"value":      Representation{RepType: Optional, Create: `30`, Update: `20`},
		"values":     RepresentationGroup{Optional, detectorRecipeDetectorRulesDetailsConfigurationsValuesRepresentation},
	}
	//Making a valid configuration values object
	detectorRecipeDetectorRulesDetailsConfigurationsValuesRepresentation = map[string]interface{}{
		"list_type":         Representation{RepType: Required, Create: `CUSTOM`, Update: `MANAGED`},
		"managed_list_type": Representation{RepType: Required, Create: `RESOURCE_OCID`, Update: `RESOURCE_OCID`},
		"value":             Representation{RepType: Required, Create: `resourceOcid1`, Update: `resourceOcid2`},
	}
	//Make a representation for plural datasource
	detectorRecipeDataSourceRepresentationPluralDataSource = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: GetEnvSettingWithBlankDefault("tenancy_ocid")},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
	}

	//Corrected the dependencies.
	DetectorRecipeResourceDependencies = GenerateDataSourceFromRepresentationMap("oci_cloud_guard_detector_recipes", "oracle_detector_recipe", Required, Create, detectorRecipeDataSourceRepresentationPluralDataSource) +
		DefinedTagsDependencies
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardDetectorRecipeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardDetectorRecipeResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_guard_detector_recipe.test_detector_recipe"
	datasourceName := "data.oci_cloud_guard_detector_recipes.test_detector_recipes"
	singularDatasourceName := "data.oci_cloud_guard_detector_recipe.test_detector_recipe"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DetectorRecipeResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_cloud_guard_detector_recipe", "test_detector_recipe", Optional, Create, detectorRecipeRepresentation), "cloudguard", "detectorRecipe", t)

	ResourceTest(t, testAccCheckCloudGuardDetectorRecipeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DetectorRecipeResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_cloud_guard_detector_recipe", "test_detector_recipe", Required, Create, detectorRecipeRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "source_detector_recipe_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DetectorRecipeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DetectorRecipeResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_cloud_guard_detector_recipe", "test_detector_recipe", Optional, Create, detectorRecipeRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "detector"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.#", "1"),
				//Just checking it being set, it being a JSON
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.details.0.condition"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.config_key", "lbCertificateExpiringSoonConfig"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.data_type", "multiList"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.name", "Days before expiring - 1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.value", "30"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.values.0.list_type", "CUSTOM"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.values.0.managed_list_type", "RESOURCE_OCID"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.values.0.value", "resourceOcid1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.risk_level", "CRITICAL"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.detector"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.detector_rule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.resource_type"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.service_type"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "owner"),
				resource.TestCheckResourceAttrSet(resourceName, "source_detector_recipe_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DetectorRecipeResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_cloud_guard_detector_recipe", "test_detector_recipe", Optional, Create,
					RepresentationCopyWithNewProperties(detectorRecipeRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "detector"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.details.0.condition"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.config_key", "lbCertificateExpiringSoonConfig"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.data_type", "multiList"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.name", "Days before expiring - 1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.value", "30"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.values.0.list_type", "CUSTOM"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.values.0.managed_list_type", "RESOURCE_OCID"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.values.0.value", "resourceOcid1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.risk_level", "CRITICAL"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.detector"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.detector_rule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.resource_type"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.service_type"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "owner"),
				resource.TestCheckResourceAttrSet(resourceName, "source_detector_recipe_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DetectorRecipeResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_cloud_guard_detector_recipe", "test_detector_recipe", Optional, Update, detectorRecipeRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "detector"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.details.0.condition"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.config_key", "lbCertificateExpiringSoonConfig"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.data_type", "multiList"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.name", "Days before expiring - 2"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.value", "20"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.values.0.list_type", "MANAGED"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.values.0.managed_list_type", "RESOURCE_OCID"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.configurations.0.values.0.value", "resourceOcid2"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detector_rules.0.details.0.risk_level", "LOW"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.detector"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.detector_rule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.resource_type"),
				resource.TestCheckResourceAttrSet(resourceName, "detector_rules.0.service_type"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "owner"),
				resource.TestCheckResourceAttrSet(resourceName, "source_detector_recipe_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_cloud_guard_detector_recipes", "test_detector_recipes", Optional, Update, detectorRecipeDataSourceRepresentation) +
				compartmentIdVariableStr + DetectorRecipeResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_cloud_guard_detector_recipe", "test_detector_recipe", Optional, Update, detectorRecipeRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "resource_metadata_only", "false"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "detector_recipe_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "detector_recipe_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_cloud_guard_detector_recipe", "test_detector_recipe", Required, Create, detectorRecipeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DetectorRecipeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detector_recipe_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detector"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.#", "1"),
				//This field may or may not be present - depends on the rule.
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detector_rules.0.description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detector_rules.0.details.0.condition"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.details.0.configurations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.details.0.configurations.0.config_key", "lbCertificateExpiringSoonConfig"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.details.0.configurations.0.data_type", "multiList"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.details.0.configurations.0.name", "Days before expiring - 2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.details.0.configurations.0.value", "20"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.details.0.configurations.0.values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.details.0.configurations.0.values.0.list_type", "MANAGED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.details.0.configurations.0.values.0.managed_list_type", "RESOURCE_OCID"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.details.0.configurations.0.values.0.value", "resourceOcid2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detector_rules.0.details.0.is_configuration_allowed"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.details.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.details.0.labels.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.details.0.risk_level", "LOW"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detector_rules.0.detector"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detector_rules.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detector_rules.0.managed_list_types.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detector_rules.0.recommendation"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detector_rules.0.resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detector_rules.0.service_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detector_rules.0.state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detector_rules.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detector_rules.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				//These are effective rules, after applying defaults over user input so here the count is more, but count can change on adding more rules,
				//so we will check for existence only
				resource.TestCheckResourceAttrSet(singularDatasourceName, "effective_detector_rules.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "owner"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DetectorRecipeResourceConfig,
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

func testAccCheckCloudGuardDetectorRecipeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := TestAccProvider.Meta().(*OracleClients).cloudGuardClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_guard_detector_recipe" {
			noResourceFound = false
			request := oci_cloud_guard.GetDetectorRecipeRequest{}

			tmp := rs.Primary.ID
			request.DetectorRecipeId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "cloud_guard")

			response, err := client.GetDetectorRecipe(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_guard.LifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		InitDependencyGraph()
	}
	if !InSweeperExcludeList("CloudGuardDetectorRecipe") {
		resource.AddTestSweepers("CloudGuardDetectorRecipe", &resource.Sweeper{
			Name:         "CloudGuardDetectorRecipe",
			Dependencies: DependencyGraph["detectorRecipe"],
			F:            sweepCloudGuardDetectorRecipeResource,
		})
	}
}

func sweepCloudGuardDetectorRecipeResource(compartment string) error {
	cloudGuardClient := GetTestClients(&schema.ResourceData{}).cloudGuardClient()
	detectorRecipeIds, err := getDetectorRecipeIds(compartment)
	if err != nil {
		return err
	}
	for _, detectorRecipeId := range detectorRecipeIds {
		if ok := SweeperDefaultResourceId[detectorRecipeId]; !ok {
			deleteDetectorRecipeRequest := oci_cloud_guard.DeleteDetectorRecipeRequest{}

			deleteDetectorRecipeRequest.DetectorRecipeId = &detectorRecipeId

			deleteDetectorRecipeRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "cloud_guard")
			_, error := cloudGuardClient.DeleteDetectorRecipe(context.Background(), deleteDetectorRecipeRequest)
			if error != nil {
				fmt.Printf("Error deleting DetectorRecipe %s %s, It is possible that the resource is already deleted. Please verify manually \n", detectorRecipeId, error)
				continue
			}
			WaitTillCondition(TestAccProvider, &detectorRecipeId, detectorRecipeSweepWaitCondition, time.Duration(3*time.Minute),
				detectorRecipeSweepResponseFetchOperation, "cloud_guard", true)
		}
	}
	return nil
}

func getDetectorRecipeIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "DetectorRecipeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	cloudGuardClient := GetTestClients(&schema.ResourceData{}).cloudGuardClient()

	listDetectorRecipesRequest := oci_cloud_guard.ListDetectorRecipesRequest{}
	listDetectorRecipesRequest.CompartmentId = &compartmentId
	listDetectorRecipesRequest.LifecycleState = oci_cloud_guard.ListDetectorRecipesLifecycleStateActive
	listDetectorRecipesResponse, err := cloudGuardClient.ListDetectorRecipes(context.Background(), listDetectorRecipesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DetectorRecipe list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, detectorRecipe := range listDetectorRecipesResponse.Items {
		id := *detectorRecipe.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "DetectorRecipeId", id)
	}
	return resourceIds, nil
}

func detectorRecipeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if detectorRecipeResponse, ok := response.Response.(oci_cloud_guard.GetDetectorRecipeResponse); ok {
		return detectorRecipeResponse.LifecycleState != oci_cloud_guard.LifecycleStateDeleted
	}
	return false
}

func detectorRecipeSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.cloudGuardClient().GetDetectorRecipe(context.Background(), oci_cloud_guard.GetDetectorRecipeRequest{
		DetectorRecipeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
