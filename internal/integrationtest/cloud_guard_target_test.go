// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v58/cloudguard"
	"github.com/oracle/oci-go-sdk/v58/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	TargetRequiredOnlyResource = TargetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", acctest.Required, acctest.Create, targetRepresentation)

	TargetResourceConfig = TargetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", acctest.Optional, acctest.Update, targetRepresentation)

	targetSingularDataSourceRepresentation = map[string]interface{}{
		"target_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_target.test_target.id}`},
	}

	targetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `ACTIVE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: targetDataSourceFilterRepresentation}}
	targetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_guard_target.test_target.id}`}},
	}

	targetRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		//For now can be equal to compartmentId only
		"target_resource_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		//For now can be equal to COMPARTMENT only
		"target_resource_type":     acctest.Representation{RepType: acctest.Required, Create: `COMPARTMENT`},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":              acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"state":                    acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `ACTIVE`},
		"target_detector_recipes":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: targetTargetDetectorRecipesRepresentation},
		"target_responder_recipes": acctest.RepresentationGroup{RepType: acctest.Optional, Group: targetTargetResponderRecipesRepresentation},
	}
	//Getting detectorRecipeId and responderRecipeId from a plural datasource call same as in detectorRecipeTest and responderRecipeTest
	targetTargetDetectorRecipesRepresentation = map[string]interface{}{
		"detector_recipe_id": acctest.Representation{RepType: acctest.Required, Create: detectorRecipeId},
		"detector_rules":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: targetTargetDetectorRecipesDetectorRulesRepresentation},
	}
	targetTargetResponderRecipesRepresentation = map[string]interface{}{
		"responder_recipe_id": acctest.Representation{RepType: acctest.Required, Create: responderRecipeId},
		"responder_rules":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: targetTargetResponderRecipesResponderRulesRepresentation},
	}
	//Hardcoding a detectorRuleId and responderRuleId as the condition and configuration are dependent on the rules, so hardcoding one for testing purposes.
	targetTargetDetectorRecipesDetectorRulesRepresentation = map[string]interface{}{
		"details":          acctest.RepresentationGroup{RepType: acctest.Required, Group: targetTargetDetectorRecipesDetectorRulesDetailsRepresentation},
		"detector_rule_id": acctest.Representation{RepType: acctest.Required, Create: `LB_CERTIFICATE_EXPIRING_SOON`},
	}
	targetTargetResponderRecipesResponderRulesRepresentation = map[string]interface{}{
		"details":           acctest.RepresentationGroup{RepType: acctest.Required, Group: targetTargetResponderRecipesResponderRulesDetailsRepresentation},
		"responder_rule_id": acctest.Representation{RepType: acctest.Required, Create: `ENABLE_DB_BACKUP`},
	}
	targetTargetDetectorRecipesDetectorRulesDetailsRepresentation = map[string]interface{}{
		"condition_groups": acctest.RepresentationGroup{RepType: acctest.Optional, Group: targetTargetDetectorRecipesDetectorRulesDetailsConditionGroupsRepresentation},
	}
	//Making correct conditions, mode and configuration Object for responder/detector rule
	targetTargetResponderRecipesResponderRulesDetailsRepresentation = map[string]interface{}{
		"condition":      acctest.Representation{RepType: acctest.Optional, Create: `{\"kind\":\"SIMPLE\",\"parameter\":\"resourceName\",\"operator\":\"EQUALS\",\"value\":\"bucket1\",\"valueType\":\"CUSTOM\"}`, Update: `{\"kind\":\"COMPOSITE\",\"leftOperand\":{\"kind\" :\"COMPOSITE\",\"leftOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"resourceName\",\"operator\":\"EQUALS\",\"value\":\"bucket3\",\"valueType\":\"CUSTOM\"},\"compositeOperator\":\"AND\",\"rightOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"resourceName\",\"operator\":\"NOT_EQUALS\",\"value\":\"bucket12\",\"valueType\":\"CUSTOM\"}},\"compositeOperator\":\"AND\",\"rightOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"resourceName\",\"operator\":\"EQUALS\",\"value\":\"bucket101\",\"valueType\":\"CUSTOM\"}}`},
		"configurations": acctest.RepresentationGroup{RepType: acctest.Optional, Group: targetTargetResponderRecipesResponderRulesDetailsConfigurationsRepresentation},
		"mode":           acctest.Representation{RepType: acctest.Optional, Create: `USERACTION`, Update: `AUTOACTION`},
	}
	targetTargetDetectorRecipesDetectorRulesDetailsConditionGroupsRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"condition":      acctest.Representation{RepType: acctest.Required, Create: `{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"EQUALS\",\"value\":\"10\",\"valueType\":\"CUSTOM\"}`, Update: `{\"kind\":\"COMPOSITE\",\"leftOperand\":{\"kind\" :\"COMPOSITE\",\"leftOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"EQUALS\",\"value\":\"12\",\"valueType\":\"CUSTOM\"},\"compositeOperator\":\"AND\",\"rightOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"NOT_EQUALS\",\"value\":\"12\",\"valueType\":\"CUSTOM\"}},\"compositeOperator\":\"AND\",\"rightOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"EQUALS\",\"value\":\"10\",\"valueType\":\"CUSTOM\"}}`},
	}
	targetTargetResponderRecipesResponderRulesDetailsConfigurationsRepresentation = map[string]interface{}{
		"config_key": acctest.Representation{RepType: acctest.Required, Create: `autoBackupWindowConfig`, Update: `recoveryWindowInDaysConfig`},
		"name":       acctest.Representation{RepType: acctest.Required, Create: `Backup time window (Slot)`, Update: `Backup retention period in days`},
		"value":      acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `20`},
	}

	//Correcting the dependencies
	TargetResourceDependencies = ResponderRecipeResourceDependencies + DetectorRecipeResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_detector_recipe", "test_detector_recipe", acctest.Required, acctest.Create,
		detectorRecipeRepresentation)
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardTargetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardTargetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_cloud_guard_target.test_target"
	datasourceName := "data.oci_cloud_guard_targets.test_targets"
	singularDatasourceName := "data.oci_cloud_guard_target.test_target"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+TargetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", acctest.Optional, acctest.Create, targetRepresentation), "cloudguard", "target", t)

	acctest.ResourceTest(t, testAccCheckCloudGuardTargetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + TargetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", acctest.Required, acctest.Create, targetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_type", "COMPARTMENT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + TargetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + TargetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", acctest.Optional, acctest.Create, targetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "recipe_count"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "target_detector_recipes.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_recipe_id"),
				resource.TestCheckResourceAttr(resourceName, "target_detector_recipes.0.detector_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_detector_recipes.0.detector_rules.0.details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_detector_recipes.0.detector_rules.0.details.0.condition_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_detector_recipes.0.detector_rules.0.details.0.condition_groups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_rules.0.details.0.condition_groups.0.condition"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_rules.0.details.0.is_enabled"),
				//Not in I/P so can't be in O/P as well, but will be in effective_rules
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.effective_detector_rules.0.details.0.risk_level"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_rules.0.detector"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_rules.0.detector_rule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_rules.0.resource_type"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_rules.0.service_type"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.owner"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_type", "COMPARTMENT"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.description"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.owner"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.responder_recipe_id"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.0.details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.condition"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.configurations.0.config_key", "autoBackupWindowConfig"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.configurations.0.name", "Backup time window (Slot)"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.configurations.0.value", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.is_enabled"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.mode", "USERACTION"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.responder_rules.0.responder_rule_id"),

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
			Config: config + compartmentIdVariableStr + TargetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", acctest.Optional, acctest.Update, targetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "recipe_count"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "target_detector_recipes.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_recipe_id"),
				resource.TestCheckResourceAttr(resourceName, "target_detector_recipes.0.detector_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_detector_recipes.0.detector_rules.0.details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_detector_recipes.0.detector_rules.0.details.0.condition_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_detector_recipes.0.detector_rules.0.details.0.condition_groups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_rules.0.details.0.condition_groups.0.condition"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_rules.0.details.0.is_enabled"),
				//Not in I/P so can't be in O/P as well, but will be in effective_rules
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.effective_detector_rules.0.details.0.risk_level"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_rules.0.detector"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_rules.0.detector_rule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_rules.0.resource_type"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.detector_rules.0.service_type"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_detector_recipes.0.owner"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_type", "COMPARTMENT"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.description"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.owner"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.responder_recipe_id"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.0.details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.condition"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.configurations.0.config_key", "recoveryWindowInDaysConfig"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.configurations.0.name", "Backup retention period in days"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.configurations.0.value", "20"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.is_enabled"),
				resource.TestCheckResourceAttr(resourceName, "target_responder_recipes.0.responder_rules.0.details.0.mode", "AUTOACTION"),
				resource.TestCheckResourceAttrSet(resourceName, "target_responder_recipes.0.responder_rules.0.responder_rule_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_targets", "test_targets", acctest.Optional, acctest.Update, targetDataSourceRepresentation) +
				compartmentIdVariableStr + TargetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", acctest.Optional, acctest.Update, targetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "target_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "target_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_target", "test_target", acctest.Required, acctest.Create, targetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + TargetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				//This depends on where target is created, not a fixed value
				resource.TestCheckResourceAttrSet(singularDatasourceName, "inherited_by_compartments.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recipe_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_detector_recipes.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.detector"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_detector_recipes.0.detector_rules.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.details.0.condition_groups.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.details.0.condition_groups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.details.0.condition_groups.0.condition"),
				//These is not input so can't be in output (configuration)
				//This will be in effective detector rules after applying defaults.
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.effective_detector_rules.0.details.0.is_configuration_allowed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.effective_detector_rules.0.details.0.is_enabled"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_detector_recipes.0.effective_detector_rules.0.details.0.labels.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.effective_detector_rules.0.details.0.risk_level"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.detector"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.managed_list_types.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.recommendation"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.service_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.detector_rules.0.time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.display_name"),
				//Count will be more after applying defaults
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.effective_detector_rules.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.owner"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_detector_recipes.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_resource_type", "COMPARTMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_responder_recipes.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.display_name"),
				//Effective count will be more having defaults
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.effective_responder_rules.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.owner"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_responder_recipes.0.responder_rules.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.details.0.condition"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.details.0.configurations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.details.0.configurations.0.config_key", "recoveryWindowInDaysConfig"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.details.0.configurations.0.name", "Backup retention period in days"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.details.0.configurations.0.value", "20"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.details.0.is_enabled"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.details.0.mode", "AUTOACTION"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.display_name"),
				//Changes from backend
				resource.TestCheckResourceAttr(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.policies.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.state"),
				//2 supported modes possible as of now
				resource.TestCheckResourceAttr(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.supported_modes.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.responder_rules.0.type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_responder_recipes.0.time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + TargetResourceConfig,
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

func testAccCheckCloudGuardTargetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CloudGuardClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_guard_target" {
			noResourceFound = false
			request := oci_cloud_guard.GetTargetRequest{}

			tmp := rs.Primary.ID
			request.TargetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")

			response, err := client.GetTarget(context.Background(), request)

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CloudGuardTarget") {
		resource.AddTestSweepers("CloudGuardTarget", &resource.Sweeper{
			Name:         "CloudGuardTarget",
			Dependencies: acctest.DependencyGraph["target"],
			F:            sweepCloudGuardTargetResource,
		})
	}
}

func sweepCloudGuardTargetResource(compartment string) error {
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()
	targetIds, err := getTargetIds(compartment)
	if err != nil {
		return err
	}
	for _, targetId := range targetIds {
		if ok := acctest.SweeperDefaultResourceId[targetId]; !ok {
			deleteTargetRequest := oci_cloud_guard.DeleteTargetRequest{}

			deleteTargetRequest.TargetId = &targetId

			deleteTargetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")
			_, error := cloudGuardClient.DeleteTarget(context.Background(), deleteTargetRequest)
			if error != nil {
				fmt.Printf("Error deleting Target %s %s, It is possible that the resource is already deleted. Please verify manually \n", targetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &targetId, targetSweepWaitCondition, time.Duration(3*time.Minute),
				targetSweepResponseFetchOperation, "cloud_guard", true)
		}
	}
	return nil
}

func getTargetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "TargetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()

	listTargetsRequest := oci_cloud_guard.ListTargetsRequest{}
	listTargetsRequest.CompartmentId = &compartmentId
	listTargetsRequest.LifecycleState = oci_cloud_guard.ListTargetsLifecycleStateActive
	listTargetsResponse, err := cloudGuardClient.ListTargets(context.Background(), listTargetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Target list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, target := range listTargetsResponse.Items {
		id := *target.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "TargetId", id)
	}
	return resourceIds, nil
}

func targetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if targetResponse, ok := response.Response.(oci_cloud_guard.GetTargetResponse); ok {
		return targetResponse.LifecycleState != oci_cloud_guard.LifecycleStateDeleted
	}
	return false
}

func targetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.CloudGuardClient().GetTarget(context.Background(), oci_cloud_guard.GetTargetRequest{
		TargetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
