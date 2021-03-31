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
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v38/cloudguard"
	"github.com/oracle/oci-go-sdk/v38/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	TargetRequiredOnlyResource = TargetResourceDependencies +
		generateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", Required, Create, targetRepresentation)

	TargetResourceConfig = TargetResourceDependencies +
		generateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", Optional, Update, targetRepresentation)

	targetSingularDataSourceRepresentation = map[string]interface{}{
		"target_id": Representation{repType: Required, create: `${oci_cloud_guard_target.test_target.id}`},
	}

	targetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"access_level":              Representation{repType: Optional, create: `ACCESSIBLE`},
		"compartment_id_in_subtree": Representation{repType: Optional, create: `true`},
		"display_name":              Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":                     Representation{repType: Optional, create: `ACTIVE`, update: `ACTIVE`},
		"filter":                    RepresentationGroup{Required, targetDataSourceFilterRepresentation}}
	targetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_cloud_guard_target.test_target.id}`}},
	}

	targetRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `displayName`, update: `displayName2`},
		//For now can be equal to compartmentId only
		"target_resource_id": Representation{repType: Required, create: `${var.compartment_id}`},
		//For now can be equal to COMPARTMENT only
		"target_resource_type":     Representation{repType: Required, create: `COMPARTMENT`},
		"defined_tags":             Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":              Representation{repType: Optional, create: `description`},
		"freeform_tags":            Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"state":                    Representation{repType: Optional, create: `ACTIVE`, update: `ACTIVE`},
		"target_detector_recipes":  RepresentationGroup{Optional, targetTargetDetectorRecipesRepresentation},
		"target_responder_recipes": RepresentationGroup{Optional, targetTargetResponderRecipesRepresentation},
	}
	//Getting detectorRecipeId and responderRecipeId from a plural datasource call same as in detectorRecipeTest and responderRecipeTest
	targetTargetDetectorRecipesRepresentation = map[string]interface{}{
		"detector_recipe_id": Representation{repType: Required, create: detectorRecipeId},
		"detector_rules":     RepresentationGroup{Optional, targetTargetDetectorRecipesDetectorRulesRepresentation},
	}
	targetTargetResponderRecipesRepresentation = map[string]interface{}{
		"responder_recipe_id": Representation{repType: Required, create: responderRecipeId},
		"responder_rules":     RepresentationGroup{Optional, targetTargetResponderRecipesResponderRulesRepresentation},
	}
	//Hardcoding a detectorRuleId and responderRuleId as the condition and configuration are dependent on the rules, so hardcoding one for testing purposes.
	targetTargetDetectorRecipesDetectorRulesRepresentation = map[string]interface{}{
		"details":          RepresentationGroup{Required, targetTargetDetectorRecipesDetectorRulesDetailsRepresentation},
		"detector_rule_id": Representation{repType: Required, create: `LB_CERTIFICATE_EXPIRING_SOON`},
	}
	targetTargetResponderRecipesResponderRulesRepresentation = map[string]interface{}{
		"details":           RepresentationGroup{Required, targetTargetResponderRecipesResponderRulesDetailsRepresentation},
		"responder_rule_id": Representation{repType: Required, create: `ENABLE_DB_BACKUP`},
	}
	targetTargetDetectorRecipesDetectorRulesDetailsRepresentation = map[string]interface{}{
		"condition_groups": RepresentationGroup{Optional, targetTargetDetectorRecipesDetectorRulesDetailsConditionGroupsRepresentation},
	}
	//Making correct conditions, mode and configuration Object for responder/detector rule
	targetTargetResponderRecipesResponderRulesDetailsRepresentation = map[string]interface{}{
		"condition":      Representation{repType: Optional, create: `{\"kind\":\"SIMPLE\",\"parameter\":\"resourceName\",\"operator\":\"EQUALS\",\"value\":\"bucket1\",\"valueType\":\"CUSTOM\"}`, update: `{\"kind\":\"COMPOSITE\",\"leftOperand\":{\"kind\" :\"COMPOSITE\",\"leftOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"resourceName\",\"operator\":\"EQUALS\",\"value\":\"bucket3\",\"valueType\":\"CUSTOM\"},\"compositeOperator\":\"AND\",\"rightOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"resourceName\",\"operator\":\"NOT_EQUALS\",\"value\":\"bucket12\",\"valueType\":\"CUSTOM\"}},\"compositeOperator\":\"AND\",\"rightOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"resourceName\",\"operator\":\"EQUALS\",\"value\":\"bucket101\",\"valueType\":\"CUSTOM\"}}`},
		"configurations": RepresentationGroup{Optional, targetTargetResponderRecipesResponderRulesDetailsConfigurationsRepresentation},
		"mode":           Representation{repType: Optional, create: `USERACTION`, update: `AUTOACTION`},
	}
	targetTargetDetectorRecipesDetectorRulesDetailsConditionGroupsRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"condition":      Representation{repType: Required, create: `{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"EQUALS\",\"value\":\"10\",\"valueType\":\"CUSTOM\"}`, update: `{\"kind\":\"COMPOSITE\",\"leftOperand\":{\"kind\" :\"COMPOSITE\",\"leftOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"EQUALS\",\"value\":\"12\",\"valueType\":\"CUSTOM\"},\"compositeOperator\":\"AND\",\"rightOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"NOT_EQUALS\",\"value\":\"12\",\"valueType\":\"CUSTOM\"}},\"compositeOperator\":\"AND\",\"rightOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"EQUALS\",\"value\":\"10\",\"valueType\":\"CUSTOM\"}}`},
	}
	targetTargetResponderRecipesResponderRulesDetailsConfigurationsRepresentation = map[string]interface{}{
		"config_key": Representation{repType: Required, create: `autoBackupWindowConfig`, update: `recoveryWindowInDaysConfig`},
		"name":       Representation{repType: Required, create: `Backup time window (Slot)`, update: `Backup retention period in days`},
		"value":      Representation{repType: Required, create: `10`, update: `20`},
	}

	//Correcting the dependencies
	TargetResourceDependencies = ResponderRecipeResourceDependencies + DetectorRecipeResourceDependencies + generateResourceFromRepresentationMap("oci_cloud_guard_detector_recipe", "test_detector_recipe", Required, Create,
		detectorRecipeRepresentation)
)

func TestCloudGuardTargetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardTargetResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_cloud_guard_target.test_target"
	datasourceName := "data.oci_cloud_guard_targets.test_targets"
	singularDatasourceName := "data.oci_cloud_guard_target.test_target"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+TargetResourceDependencies+
		generateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", Optional, Create, targetRepresentation), "cloudguard", "target", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCloudGuardTargetDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + TargetResourceDependencies +
					generateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", Required, Create, targetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "target_resource_id"),
					resource.TestCheckResourceAttr(resourceName, "target_resource_type", "COMPARTMENT"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + TargetResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + TargetResourceDependencies +
					generateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", Optional, Create, targetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", Optional, Update, targetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_cloud_guard_targets", "test_targets", Optional, Update, targetDataSourceRepresentation) +
					compartmentIdVariableStr + TargetResourceDependencies +
					generateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", Optional, Update, targetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_cloud_guard_target", "test_target", Required, Create, targetSingularDataSourceRepresentation) +
					compartmentIdVariableStr + TargetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "target_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
		},
	})
}

func testAccCheckCloudGuardTargetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).cloudGuardClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_guard_target" {
			noResourceFound = false
			request := oci_cloud_guard.GetTargetRequest{}

			tmp := rs.Primary.ID
			request.TargetId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "cloud_guard")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CloudGuardTarget") {
		resource.AddTestSweepers("CloudGuardTarget", &resource.Sweeper{
			Name:         "CloudGuardTarget",
			Dependencies: DependencyGraph["target"],
			F:            sweepCloudGuardTargetResource,
		})
	}
}

func sweepCloudGuardTargetResource(compartment string) error {
	cloudGuardClient := GetTestClients(&schema.ResourceData{}).cloudGuardClient()
	targetIds, err := getTargetIds(compartment)
	if err != nil {
		return err
	}
	for _, targetId := range targetIds {
		if ok := SweeperDefaultResourceId[targetId]; !ok {
			deleteTargetRequest := oci_cloud_guard.DeleteTargetRequest{}

			deleteTargetRequest.TargetId = &targetId

			deleteTargetRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "cloud_guard")
			_, error := cloudGuardClient.DeleteTarget(context.Background(), deleteTargetRequest)
			if error != nil {
				fmt.Printf("Error deleting Target %s %s, It is possible that the resource is already deleted. Please verify manually \n", targetId, error)
				continue
			}
			waitTillCondition(testAccProvider, &targetId, targetSweepWaitCondition, time.Duration(3*time.Minute),
				targetSweepResponseFetchOperation, "cloud_guard", true)
		}
	}
	return nil
}

func getTargetIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "TargetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	cloudGuardClient := GetTestClients(&schema.ResourceData{}).cloudGuardClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "TargetId", id)
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

func targetSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.cloudGuardClient().GetTarget(context.Background(), oci_cloud_guard.GetTargetRequest{
		TargetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
