// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsSecurityQuestionSettingRequiredOnlyResource = IdentityDomainsSecurityQuestionSettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_security_question_setting", "test_security_question_setting", acctest.Required, acctest.Create, IdentityDomainsSecurityQuestionSettingRepresentation)

	IdentityDomainsSecurityQuestionSettingResourceConfig = IdentityDomainsSecurityQuestionSettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_security_question_setting", "test_security_question_setting", acctest.Optional, acctest.Update, IdentityDomainsSecurityQuestionSettingRepresentation)

	IdentityDomainsIdentityDomainsSecurityQuestionSettingSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"security_question_setting_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_security_question_setting.test_security_question_setting.id}`},
		"attribute_sets":               acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsSecurityQuestionSettingDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsSecurityQuestionSettingRepresentation = map[string]interface{}{
		"idcs_endpoint":                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"max_field_length":             acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"min_answer_length":            acctest.Representation{RepType: acctest.Required, Create: `6`, Update: `8`},
		"num_questions_to_ans":         acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `1`},
		"num_questions_to_setup":       acctest.Representation{RepType: acctest.Required, Create: `5`, Update: `4`},
		"schemas":                      acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:SecurityQuestionSettings`}},
		"security_question_setting_id": acctest.Representation{RepType: acctest.Required, Create: `SecurityQuestionSettings`},
		"attribute_sets":               acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"external_id":                  acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"tags":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSecurityQuestionSettingTagsRepresentation},
	}
	IdentityDomainsSecurityQuestionSettingTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsSecurityQuestionSettingResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsSecurityQuestionSettingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsSecurityQuestionSettingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_security_question_setting.test_security_question_setting"
	datasourceName := "data.oci_identity_domains_security_question_settings.test_security_question_settings"
	singularDatasourceName := "data.oci_identity_domains_security_question_setting.test_security_question_setting"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsSecurityQuestionSettingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_security_question_setting", "test_security_question_setting", acctest.Optional, acctest.Create, IdentityDomainsSecurityQuestionSettingRepresentation), "identitydomains", "securityQuestionSetting", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSecurityQuestionSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_security_question_setting", "test_security_question_setting", acctest.Required, acctest.Create, IdentityDomainsSecurityQuestionSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "max_field_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "min_answer_length", "6"),
				resource.TestCheckResourceAttr(resourceName, "num_questions_to_ans", "2"),
				resource.TestCheckResourceAttr(resourceName, "num_questions_to_setup", "5"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "security_question_setting_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSecurityQuestionSettingResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSecurityQuestionSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_security_question_setting", "test_security_question_setting", acctest.Optional, acctest.Create, IdentityDomainsSecurityQuestionSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "id", "SecurityQuestionSettings"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "max_field_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "min_answer_length", "6"),
				resource.TestCheckResourceAttr(resourceName, "num_questions_to_ans", "2"),
				resource.TestCheckResourceAttr(resourceName, "num_questions_to_setup", "5"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "security_question_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "securityQuestionSettings", resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}

					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSecurityQuestionSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_security_question_setting", "test_security_question_setting", acctest.Optional, acctest.Update, IdentityDomainsSecurityQuestionSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "id", "SecurityQuestionSettings"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "max_field_length", "11"),
				resource.TestCheckResourceAttr(resourceName, "min_answer_length", "8"),
				resource.TestCheckResourceAttr(resourceName, "num_questions_to_ans", "1"),
				resource.TestCheckResourceAttr(resourceName, "num_questions_to_setup", "4"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "security_question_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_security_question_settings", "test_security_question_settings", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsSecurityQuestionSettingDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsSecurityQuestionSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_security_question_setting", "test_security_question_setting", acctest.Optional, acctest.Update, IdentityDomainsSecurityQuestionSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "security_question_settings.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "security_question_settings.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_security_question_setting", "test_security_question_setting", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsSecurityQuestionSettingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsSecurityQuestionSettingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_question_setting_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "max_field_length", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_answer_length", "8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "num_questions_to_ans", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "num_questions_to_setup", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsSecurityQuestionSettingRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_security_question_setting", "securityQuestionSettings"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"idcs_last_upgraded_in_release",
				"idcs_prevented_operations",
				"tags",
				"security_question_setting_id",
			},
			ResourceName: resourceName,
		},
	})
}
