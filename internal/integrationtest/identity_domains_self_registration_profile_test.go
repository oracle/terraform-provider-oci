// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsSelfRegistrationProfileRequiredOnlyResource = IdentityDomainsSelfRegistrationProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_self_registration_profile", "test_self_registration_profile", acctest.Required, acctest.Create, IdentityDomainsSelfRegistrationProfileRepresentation)

	IdentityDomainsSelfRegistrationProfileResourceConfig = IdentityDomainsSelfRegistrationProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_self_registration_profile", "test_self_registration_profile", acctest.Optional, acctest.Update, IdentityDomainsSelfRegistrationProfileRepresentation)

	IdentityDomainsSelfRegistrationProfileSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"self_registration_profile_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_self_registration_profile.test_self_registration_profile.id}`},
		"attribute_sets":               acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsSelfRegistrationProfileDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"self_registration_profile_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"self_registration_profile_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":                      acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsSelfRegistrationProfileRepresentation = map[string]interface{}{
		"activation_email_required":            acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"consent_text_present":                 acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"display_name":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsSelfRegistrationProfileDisplayNameRepresentation},
		"email_template":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsSelfRegistrationProfileEmailTemplateRepresentation},
		"idcs_endpoint":                        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"name":                                 acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"number_of_days_redirect_url_is_valid": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"redirect_url":                         acctest.Representation{RepType: acctest.Required, Create: `https://www.oracle.com`, Update: `https://www.oraclecloud.com`},
		"schemas":                              acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:SelfRegistrationProfile`}},
		"show_on_login_page":                   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"active":                               acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"after_submit_text":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSelfRegistrationProfileAfterSubmitTextRepresentation},
		"allowed_email_domains":                acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"attribute_sets":                       acctest.Representation{RepType: acctest.Required, Create: []string{`all`}},
		"consent_text":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSelfRegistrationProfileConsentTextRepresentation},
		"default_groups":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSelfRegistrationProfileDefaultGroupsRepresentation},
		"disallowed_email_domains":             acctest.Representation{RepType: acctest.Optional, Create: []string{`test.com`}},
		"external_id":                          acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"footer_logo":                          acctest.Representation{RepType: acctest.Optional, Create: `footerLogo`, Update: `footerLogo2`},
		"footer_text":                          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSelfRegistrationProfileFooterTextRepresentation},
		"header_logo":                          acctest.Representation{RepType: acctest.Optional, Create: `headerLogo`, Update: `headerLogo2`},
		"header_text":                          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSelfRegistrationProfileHeaderTextRepresentation},
		"tags":                                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSelfRegistrationProfileTagsRepresentation},
		"user_attributes":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSelfRegistrationProfileUserAttributesRepresentation},
		"lifecycle":                            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesForIdentityDomainsSelfRegistrationProfile},
	}
	IdentityDomainsSelfRegistrationProfileDisplayNameRepresentation = map[string]interface{}{
		"locale":  acctest.Representation{RepType: acctest.Required, Create: `en-US`, Update: `fr`},
		"value":   acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
		"default": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}
	IdentityDomainsSelfRegistrationProfileEmailTemplateRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `MeRegisterVerifyEmail`, Update: `MeRegisterActivationEmail`},
	}
	IdentityDomainsSelfRegistrationProfileAfterSubmitTextRepresentation = map[string]interface{}{
		"locale":  acctest.Representation{RepType: acctest.Required, Create: `en-US`, Update: `fr`},
		"value":   acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
		"default": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	IdentityDomainsSelfRegistrationProfileConsentTextRepresentation = map[string]interface{}{
		"locale":  acctest.Representation{RepType: acctest.Required, Create: `en-US`, Update: `fr`},
		"value":   acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
		"default": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	IdentityDomainsSelfRegistrationProfileDefaultGroupsRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `AllUsersId`},
	}
	IdentityDomainsSelfRegistrationProfileFooterTextRepresentation = map[string]interface{}{
		"locale":  acctest.Representation{RepType: acctest.Required, Create: `en-US`, Update: `fr`},
		"value":   acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
		"default": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	IdentityDomainsSelfRegistrationProfileHeaderTextRepresentation = map[string]interface{}{
		"locale":  acctest.Representation{RepType: acctest.Required, Create: `en-US`, Update: `fr`},
		"value":   acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
		"default": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	IdentityDomainsSelfRegistrationProfileTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsSelfRegistrationProfileUserAttributesRepresentation = map[string]interface{}{
		"seq_number":                     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"value":                          acctest.Representation{RepType: acctest.Required, Create: `employeeNumber`},
		"fully_qualified_attribute_name": acctest.Representation{RepType: acctest.Optional, Create: `urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:employeeNumber`},
	}

	ignoreChangesForIdentityDomainsSelfRegistrationProfile = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			`user_attributes`,
		}},
	}

	IdentityDomainsSelfRegistrationProfileResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsSelfRegistrationProfileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsSelfRegistrationProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_self_registration_profile.test_self_registration_profile"
	datasourceName := "data.oci_identity_domains_self_registration_profiles.test_self_registration_profiles"
	singularDatasourceName := "data.oci_identity_domains_self_registration_profile.test_self_registration_profile"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsSelfRegistrationProfileResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_self_registration_profile", "test_self_registration_profile", acctest.Optional, acctest.Create, IdentityDomainsSelfRegistrationProfileRepresentation), "identitydomains", "selfRegistrationProfile", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsSelfRegistrationProfileDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSelfRegistrationProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_self_registration_profile", "test_self_registration_profile", acctest.Required, acctest.Create, IdentityDomainsSelfRegistrationProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "activation_email_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "consent_text_present", "false"),
				resource.TestCheckResourceAttr(resourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name.0.locale", "en-US"),
				resource.TestCheckResourceAttr(resourceName, "display_name.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "number_of_days_redirect_url_is_valid", "10"),
				resource.TestCheckResourceAttr(resourceName, "redirect_url", "https://www.oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "show_on_login_page", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSelfRegistrationProfileResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSelfRegistrationProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_self_registration_profile", "test_self_registration_profile", acctest.Optional, acctest.Create, IdentityDomainsSelfRegistrationProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "activation_email_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "active", "false"),
				resource.TestCheckResourceAttr(resourceName, "after_submit_text.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "after_submit_text.0.default", "true"),
				resource.TestCheckResourceAttr(resourceName, "after_submit_text.0.locale", "en-US"),
				resource.TestCheckResourceAttr(resourceName, "after_submit_text.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "allowed_email_domains.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "consent_text.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "consent_text.0.default", "true"),
				resource.TestCheckResourceAttr(resourceName, "consent_text.0.locale", "en-US"),
				resource.TestCheckResourceAttr(resourceName, "consent_text.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "consent_text_present", "false"),
				resource.TestCheckResourceAttr(resourceName, "default_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "default_groups.0.display", "All Domain Users"),
				resource.TestCheckResourceAttr(resourceName, "default_groups.0.value", "AllUsersId"),
				resource.TestCheckResourceAttr(resourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name.0.default", "true"),
				resource.TestCheckResourceAttr(resourceName, "display_name.0.locale", "en-US"),
				resource.TestCheckResourceAttr(resourceName, "display_name.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "email_template.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "email_template.0.display", "MeRegisterVerifyEmail"),
				resource.TestCheckResourceAttr(resourceName, "email_template.0.value", "MeRegisterVerifyEmail"),
				resource.TestCheckResourceAttr(resourceName, "footer_logo", "footerLogo"),
				resource.TestCheckResourceAttr(resourceName, "footer_text.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "footer_text.0.default", "true"),
				resource.TestCheckResourceAttr(resourceName, "footer_text.0.locale", "en-US"),
				resource.TestCheckResourceAttr(resourceName, "footer_text.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "header_logo", "headerLogo"),
				resource.TestCheckResourceAttr(resourceName, "header_text.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "header_text.0.default", "true"),
				resource.TestCheckResourceAttr(resourceName, "header_text.0.locale", "en-US"),
				resource.TestCheckResourceAttr(resourceName, "header_text.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "idcs_last_modified_by.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "number_of_days_redirect_url_is_valid", "10"),
				resource.TestCheckResourceAttr(resourceName, "redirect_url", "https://www.oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "show_on_login_page", "false"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestMatchResourceAttr(resourceName, "user_attributes.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(resourceName, "user_attributes.7.deletable", "true"),
				resource.TestCheckResourceAttr(resourceName, "user_attributes.7.fully_qualified_attribute_name", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:employeeNumber"),
				resource.TestCheckResourceAttr(resourceName, "user_attributes.7.seq_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_attributes.7.value", "employeeNumber"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "selfRegistrationProfiles", resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSelfRegistrationProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_self_registration_profile", "test_self_registration_profile", acctest.Optional, acctest.Update, IdentityDomainsSelfRegistrationProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "activation_email_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "active", "true"),
				resource.TestCheckResourceAttr(resourceName, "after_submit_text.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "after_submit_text.0.default", "true"),
				resource.TestCheckResourceAttr(resourceName, "after_submit_text.0.locale", "fr"),
				resource.TestCheckResourceAttr(resourceName, "after_submit_text.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "allowed_email_domains.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "consent_text.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "consent_text.0.default", "true"),
				resource.TestCheckResourceAttr(resourceName, "consent_text.0.locale", "fr"),
				resource.TestCheckResourceAttr(resourceName, "consent_text.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "consent_text_present", "true"),
				resource.TestCheckResourceAttr(resourceName, "default_groups.0.display", "All Domain Users"),
				resource.TestCheckResourceAttr(resourceName, "default_groups.0.value", "AllUsersId"),
				resource.TestCheckResourceAttr(resourceName, "disallowed_email_domains.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name.0.default", "true"),
				resource.TestCheckResourceAttr(resourceName, "display_name.0.locale", "fr"),
				resource.TestCheckResourceAttr(resourceName, "display_name.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "email_template.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "email_template.0.display", "MeRegisterActivationEmail"),
				resource.TestCheckResourceAttr(resourceName, "email_template.0.value", "MeRegisterActivationEmail"),
				resource.TestCheckResourceAttr(resourceName, "footer_logo", "footerLogo2"),
				resource.TestCheckResourceAttr(resourceName, "footer_text.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "footer_text.0.default", "true"),
				resource.TestCheckResourceAttr(resourceName, "footer_text.0.locale", "fr"),
				resource.TestCheckResourceAttr(resourceName, "footer_text.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "header_logo", "headerLogo2"),
				resource.TestCheckResourceAttr(resourceName, "header_text.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "header_text.0.default", "true"),
				resource.TestCheckResourceAttr(resourceName, "header_text.0.locale", "fr"),
				resource.TestCheckResourceAttr(resourceName, "header_text.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "idcs_last_modified_by.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "number_of_days_redirect_url_is_valid", "11"),
				resource.TestCheckResourceAttr(resourceName, "redirect_url", "https://www.oraclecloud.com"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "show_on_login_page", "true"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestMatchResourceAttr(resourceName, "user_attributes.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(resourceName, "user_attributes.7.deletable", "true"),
				resource.TestCheckResourceAttr(resourceName, "user_attributes.7.fully_qualified_attribute_name", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:employeeNumber"),
				resource.TestCheckResourceAttr(resourceName, "user_attributes.7.seq_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_attributes.7.value", "employeeNumber"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_self_registration_profiles", "test_self_registration_profiles", acctest.Optional, acctest.Update, IdentityDomainsSelfRegistrationProfileDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsSelfRegistrationProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_self_registration_profile", "test_self_registration_profile", acctest.Optional, acctest.Update, IdentityDomainsSelfRegistrationProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "self_registration_profile_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_self_registration_profile", "test_self_registration_profile", acctest.Required, acctest.Create, IdentityDomainsSelfRegistrationProfileSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsSelfRegistrationProfileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "self_registration_profile_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "activation_email_required", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "active", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "after_submit_text.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "after_submit_text.0.default", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "after_submit_text.0.locale", "fr"),
				resource.TestCheckResourceAttr(singularDatasourceName, "after_submit_text.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_email_domains.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "consent_text.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "consent_text.0.default", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "consent_text.0.locale", "fr"),
				resource.TestCheckResourceAttr(singularDatasourceName, "consent_text.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "consent_text_present", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name.0.default", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name.0.locale", "fr"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "footer_logo", "footerLogo2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "footer_text.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "footer_text.0.default", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "footer_text.0.locale", "fr"),
				resource.TestCheckResourceAttr(singularDatasourceName, "footer_text.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "header_logo", "headerLogo2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "header_text.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "header_text.0.default", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "header_text.0.locale", "fr"),
				resource.TestCheckResourceAttr(singularDatasourceName, "header_text.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_last_modified_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "meta.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "number_of_days_redirect_url_is_valid", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "redirect_url", "https://www.oraclecloud.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "show_on_login_page", "true"),
				resource.TestMatchResourceAttr(singularDatasourceName, "user_attributes.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_attributes.7.deletable", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_attributes.7.fully_qualified_attribute_name", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:employeeNumber"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_attributes.7.seq_number", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_attributes.7.value", "employeeNumber"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsSelfRegistrationProfileRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_self_registration_profile", "selfRegistrationProfiles"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"email_template",
				"default_groups",
				"tags",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsSelfRegistrationProfileDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_self_registration_profile" {
			noResourceFound = false
			request := oci_identity_domains.GetSelfRegistrationProfileRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			tmp := rs.Primary.ID
			request.SelfRegistrationProfileId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetSelfRegistrationProfile(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsSelfRegistrationProfile") {
		resource.AddTestSweepers("IdentityDomainsSelfRegistrationProfile", &resource.Sweeper{
			Name:         "IdentityDomainsSelfRegistrationProfile",
			Dependencies: acctest.DependencyGraph["selfRegistrationProfile"],
			F:            sweepIdentityDomainsSelfRegistrationProfileResource,
		})
	}
}

func sweepIdentityDomainsSelfRegistrationProfileResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	selfRegistrationProfileIds, err := getIdentityDomainsSelfRegistrationProfileIds(compartment)
	if err != nil {
		return err
	}
	for _, selfRegistrationProfileId := range selfRegistrationProfileIds {
		if ok := acctest.SweeperDefaultResourceId[selfRegistrationProfileId]; !ok {
			deleteSelfRegistrationProfileRequest := oci_identity_domains.DeleteSelfRegistrationProfileRequest{}

			deleteSelfRegistrationProfileRequest.SelfRegistrationProfileId = &selfRegistrationProfileId

			deleteSelfRegistrationProfileRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteSelfRegistrationProfile(context.Background(), deleteSelfRegistrationProfileRequest)
			if error != nil {
				fmt.Printf("Error deleting SelfRegistrationProfile %s %s, It is possible that the resource is already deleted. Please verify manually \n", selfRegistrationProfileId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsSelfRegistrationProfileIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SelfRegistrationProfileId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listSelfRegistrationProfilesRequest := oci_identity_domains.ListSelfRegistrationProfilesRequest{}
	listSelfRegistrationProfilesResponse, err := identityDomainsClient.ListSelfRegistrationProfiles(context.Background(), listSelfRegistrationProfilesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SelfRegistrationProfile list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, selfRegistrationProfile := range listSelfRegistrationProfilesResponse.Resources {
		id := *selfRegistrationProfile.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SelfRegistrationProfileId", id)
		////TODO: remove if not needed
		//acctest.SweeperDefaultResourceId[*selfRegistrationProfile.DefaultGroups] = true

	}
	return resourceIds, nil
}
