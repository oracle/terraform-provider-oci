// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsPasswordPolicyRequiredOnlyResource = IdentityDomainsPasswordPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_password_policy", "test_password_policy", acctest.Required, acctest.Create, IdentityDomainsPasswordPolicyRepresentation)

	IdentityDomainsPasswordPolicyResourceConfig = IdentityDomainsPasswordPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_password_policy", "test_password_policy", acctest.Optional, acctest.Update, IdentityDomainsPasswordPolicyRepresentation)

	IdentityDomainsIdentityDomainsPasswordPolicySingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"password_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_password_policy.test_password_policy.id}`},
		"attribute_sets":     acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsPasswordPolicyDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"password_policy_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"password_policy_filter": acctest.Representation{RepType: acctest.Optional, Create: `name eq \"name\"`},
		"attribute_sets":         acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":            acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsPasswordPolicyRepresentation = map[string]interface{}{
		"idcs_endpoint":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `name`},
		"schemas":                    acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:PasswordPolicy`}},
		"allowed_chars":              acctest.Representation{RepType: acctest.Optional, Create: `allowedChars`, Update: `allowedChars2`},
		"attribute_sets":             acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"dictionary_delimiter":       acctest.Representation{RepType: acctest.Optional, Create: `dictionaryDelimiter`, Update: `dictionaryDelimiter2`},
		"dictionary_location":        acctest.Representation{RepType: acctest.Optional, Create: `dictionaryLocation`, Update: `dictionaryLocation2`},
		"dictionary_word_disallowed": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"disallowed_chars":           acctest.Representation{RepType: acctest.Optional, Create: `a,b,c`, Update: `x,y,z`},
		"disallowed_substrings":      acctest.Representation{RepType: acctest.Optional, Create: []string{`disallowedSubstrings`}, Update: []string{`disallowedSubstrings2`}},
		"distinct_characters":        acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"external_id":                acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"first_name_disallowed":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"force_password_reset":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"groups":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsPasswordPolicyGroupsRepresentation},
		"last_name_disallowed":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"lockout_duration":           acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_incorrect_attempts":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_length":                 acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_repeated_chars":         acctest.Representation{RepType: acctest.Optional, Create: `100`, Update: `101`},
		"max_special_chars":          acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"min_alpha_numerals":         acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"min_alphas":                 acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"min_length":                 acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"min_lower_case":             acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"min_numerals":               acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"min_password_age":           acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"min_special_chars":          acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"min_unique_chars":           acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"min_upper_case":             acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"num_passwords_in_history":   acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"password_expire_warning":    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"password_expires_after":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"password_strength":          acctest.Representation{RepType: acctest.Optional, Create: `Custom`},
		"priority":                   acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `9`},
		"required_chars":             acctest.Representation{RepType: acctest.Optional, Create: `x,y,z`, Update: `a,b,c`},
		"starts_with_alphabet":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"tags":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsPasswordPolicyTagsRepresentation},
		"user_name_disallowed":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsPasswordPolicyGroupsRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_group.test_group.id}`},
	}
	IdentityDomainsPasswordPolicyTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsPasswordPolicyResourceDependencies = TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Required, acctest.Create, IdentityDomainsGroupRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsPasswordPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsPasswordPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_password_policy.test_password_policy"
	datasourceName := "data.oci_identity_domains_password_policies.test_password_policies"
	singularDatasourceName := "data.oci_identity_domains_password_policy.test_password_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsPasswordPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_password_policy", "test_password_policy", acctest.Optional, acctest.Create, IdentityDomainsPasswordPolicyRepresentation), "identitydomains", "passwordPolicy", t)

	print(config + compartmentIdVariableStr + IdentityDomainsPasswordPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_password_policy", "test_password_policy", acctest.Optional, acctest.Create, IdentityDomainsPasswordPolicyRepresentation))
	acctest.ResourceTest(t, testAccCheckIdentityDomainsPasswordPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsPasswordPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_password_policy", "test_password_policy", acctest.Required, acctest.Create, IdentityDomainsPasswordPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsPasswordPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsPasswordPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_password_policy", "test_password_policy", acctest.Optional, acctest.Create, IdentityDomainsPasswordPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "allowed_chars", "allowedChars"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "dictionary_delimiter", "dictionaryDelimiter"),
				resource.TestCheckResourceAttr(resourceName, "dictionary_location", "dictionaryLocation"),
				resource.TestCheckResourceAttr(resourceName, "dictionary_word_disallowed", "false"),
				resource.TestCheckResourceAttr(resourceName, "disallowed_chars", "a,b,c"),
				resource.TestCheckResourceAttr(resourceName, "disallowed_substrings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "distinct_characters", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "first_name_disallowed", "false"),
				resource.TestCheckResourceAttr(resourceName, "force_password_reset", "false"),
				resource.TestCheckResourceAttr(resourceName, "groups.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "groups.0.value"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "last_name_disallowed", "false"),
				resource.TestCheckResourceAttr(resourceName, "lockout_duration", "10"),
				resource.TestCheckResourceAttr(resourceName, "max_incorrect_attempts", "10"),
				resource.TestCheckResourceAttr(resourceName, "max_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "max_repeated_chars", "100"),
				resource.TestCheckResourceAttr(resourceName, "max_special_chars", "10"),
				resource.TestCheckResourceAttr(resourceName, "min_alpha_numerals", "1"),
				resource.TestCheckResourceAttr(resourceName, "min_alphas", "1"),
				resource.TestCheckResourceAttr(resourceName, "min_length", "1"),
				resource.TestCheckResourceAttr(resourceName, "min_lower_case", "1"),
				resource.TestCheckResourceAttr(resourceName, "min_numerals", "1"),
				resource.TestCheckResourceAttr(resourceName, "min_password_age", "10"),
				resource.TestCheckResourceAttr(resourceName, "min_special_chars", "1"),
				resource.TestCheckResourceAttr(resourceName, "min_unique_chars", "1"),
				resource.TestCheckResourceAttr(resourceName, "min_upper_case", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "num_passwords_in_history", "10"),
				resource.TestCheckResourceAttr(resourceName, "password_expire_warning", "10"),
				resource.TestCheckResourceAttr(resourceName, "password_expires_after", "10"),
				resource.TestCheckResourceAttr(resourceName, "password_strength", "Custom"),
				resource.TestCheckResourceAttr(resourceName, "priority", "10"),
				resource.TestCheckResourceAttr(resourceName, "required_chars", "x,y,z"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "starts_with_alphabet", "false"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "user_name_disallowed", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "passwordPolicies", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsPasswordPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_password_policy", "test_password_policy", acctest.Optional, acctest.Update, IdentityDomainsPasswordPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "allowed_chars", "allowedChars2"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "dictionary_delimiter", "dictionaryDelimiter2"),
				resource.TestCheckResourceAttr(resourceName, "dictionary_location", "dictionaryLocation2"),
				resource.TestCheckResourceAttr(resourceName, "dictionary_word_disallowed", "true"),
				resource.TestCheckResourceAttr(resourceName, "disallowed_chars", "x,y,z"),
				resource.TestCheckResourceAttr(resourceName, "disallowed_substrings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "distinct_characters", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "first_name_disallowed", "true"),
				resource.TestCheckResourceAttr(resourceName, "force_password_reset", "true"),
				resource.TestCheckResourceAttr(resourceName, "groups.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "groups.0.value"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "last_name_disallowed", "true"),
				resource.TestCheckResourceAttr(resourceName, "lockout_duration", "11"),
				resource.TestCheckResourceAttr(resourceName, "max_incorrect_attempts", "11"),
				resource.TestCheckResourceAttr(resourceName, "max_length", "11"),
				resource.TestCheckResourceAttr(resourceName, "max_repeated_chars", "101"),
				resource.TestCheckResourceAttr(resourceName, "max_special_chars", "11"),
				resource.TestCheckResourceAttr(resourceName, "min_alpha_numerals", "2"),
				resource.TestCheckResourceAttr(resourceName, "min_alphas", "2"),
				resource.TestCheckResourceAttr(resourceName, "min_length", "2"),
				resource.TestCheckResourceAttr(resourceName, "min_lower_case", "2"),
				resource.TestCheckResourceAttr(resourceName, "min_numerals", "2"),
				resource.TestCheckResourceAttr(resourceName, "min_password_age", "11"),
				resource.TestCheckResourceAttr(resourceName, "min_special_chars", "2"),
				resource.TestCheckResourceAttr(resourceName, "min_unique_chars", "2"),
				resource.TestCheckResourceAttr(resourceName, "min_upper_case", "2"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "num_passwords_in_history", "11"),
				resource.TestCheckResourceAttr(resourceName, "password_expire_warning", "11"),
				resource.TestCheckResourceAttr(resourceName, "password_expires_after", "11"),
				resource.TestCheckResourceAttr(resourceName, "password_strength", "Custom"),
				resource.TestCheckResourceAttr(resourceName, "priority", "9"),
				resource.TestCheckResourceAttr(resourceName, "required_chars", "a,b,c"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "starts_with_alphabet", "true"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "user_name_disallowed", "true"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_password_policies", "test_password_policies", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsPasswordPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsPasswordPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_password_policy", "test_password_policy", acctest.Optional, acctest.Update, IdentityDomainsPasswordPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "password_policy_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "password_policies.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "password_policies.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_password_policy", "test_password_policy", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsPasswordPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsPasswordPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "password_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "disallowed_chars", "x,y,z"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dictionary_delimiter", "dictionaryDelimiter2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dictionary_location", "dictionaryLocation2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dictionary_word_disallowed", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_chars", "allowedChars2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "disallowed_substrings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "first_name_disallowed", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "groups.0.value"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "last_name_disallowed", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lockout_duration", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_incorrect_attempts", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_length", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_repeated_chars", "101"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_special_chars", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_alpha_numerals", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_alphas", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_length", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_lower_case", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_numerals", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_password_age", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_special_chars", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_unique_chars", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_upper_case", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "num_passwords_in_history", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password_expire_warning", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password_expires_after", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password_strength", "Custom"),
				resource.TestCheckResourceAttr(singularDatasourceName, "priority", "9"),
				resource.TestCheckResourceAttr(singularDatasourceName, "required_chars", "a,b,c"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "starts_with_alphabet", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_name_disallowed", "true"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsPasswordPolicyRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_password_policy", "passwordPolicies"),
			ImportStateVerifyIgnore: []string{
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"attribute_sets",
				"attributes",
				"force_password_reset",
				"distinct_characters",
				"tags",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsPasswordPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_password_policy" {
			noResourceFound = false
			request := oci_identity_domains.GetPasswordPolicyRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.PasswordPolicyId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetPasswordPolicy(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("IdentityDomainsPasswordPolicy") {
		resource.AddTestSweepers("IdentityDomainsPasswordPolicy", &resource.Sweeper{
			Name:         "IdentityDomainsPasswordPolicy",
			Dependencies: acctest.DependencyGraph["passwordPolicy"],
			F:            sweepIdentityDomainsPasswordPolicyResource,
		})
	}
}

func sweepIdentityDomainsPasswordPolicyResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	passwordPolicyIds, err := getIdentityDomainsPasswordPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, passwordPolicyId := range passwordPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[passwordPolicyId]; !ok {
			deletePasswordPolicyRequest := oci_identity_domains.DeletePasswordPolicyRequest{}

			deletePasswordPolicyRequest.PasswordPolicyId = &passwordPolicyId

			deletePasswordPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeletePasswordPolicy(context.Background(), deletePasswordPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting PasswordPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", passwordPolicyId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsPasswordPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PasswordPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listPasswordPoliciesRequest := oci_identity_domains.ListPasswordPoliciesRequest{}
	listPasswordPoliciesResponse, err := identityDomainsClient.ListPasswordPolicies(context.Background(), listPasswordPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PasswordPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, passwordPolicy := range listPasswordPoliciesResponse.Resources {
		id := *passwordPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PasswordPolicyId", id)
	}
	return resourceIds, nil
}
