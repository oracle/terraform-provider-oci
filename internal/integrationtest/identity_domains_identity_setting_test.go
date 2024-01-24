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
	IdentityDomainsIdentitySettingRequiredOnlyResource = IdentityDomainsIdentitySettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_setting", "test_identity_setting", acctest.Required, acctest.Create, IdentityDomainsIdentitySettingRepresentation)

	IdentityDomainsIdentitySettingResourceConfig = IdentityDomainsIdentitySettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_setting", "test_identity_setting", acctest.Optional, acctest.Update, IdentityDomainsIdentitySettingRepresentation)

	IdentityDomainsIdentityDomainsIdentitySettingSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"identity_setting_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_identity_setting.test_identity_setting.id}`},
		"attribute_sets":      acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsIdentitySettingDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentitySettingRepresentation = map[string]interface{}{
		"idcs_endpoint":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"identity_setting_id": acctest.Representation{RepType: acctest.Required, Create: `IdentitySettings`},
		"schemas":             acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:IdentitySettings`}},
		"posix_gid":           acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsIdentitySettingPosixGidRepresentation},
		"posix_uid":           acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsIdentitySettingPosixUidRepresentation},
		"attribute_sets":      acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"emit_locked_message_when_user_is_locked": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"external_id":                         acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"my_profile":                          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsIdentitySettingMyProfileRepresentation},
		"primary_email_required":              acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"tags":                                acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsIdentitySettingTagsRepresentation},
		"return_inactive_over_locked_message": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"tokens": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: IdentityDomainsIdentitySettingTokensRepresentation},
			{RepType: acctest.Required, Group: IdentityDomainsIdentitySettingTokensRepresentation1},
			{RepType: acctest.Required, Group: IdentityDomainsIdentitySettingTokensRepresentation2},
		},
		"user_allowed_to_set_recovery_email": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangeForIdentityDomainsMyProfile},
	}
	ignoreChangeForIdentityDomainsMyProfile = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{
			`my_profile`,
		}},
	}
	IdentityDomainsIdentitySettingPosixGidRepresentation = map[string]interface{}{
		"manual_assignment_ends_at":     acctest.Representation{RepType: acctest.Required, Create: `1000`, Update: `2000`},
		"manual_assignment_starts_from": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}
	IdentityDomainsIdentitySettingPosixUidRepresentation = map[string]interface{}{
		"manual_assignment_ends_at":     acctest.Representation{RepType: acctest.Required, Create: `1000`, Update: `2000`},
		"manual_assignment_starts_from": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}
	IdentityDomainsIdentitySettingMyProfileRepresentation = map[string]interface{}{
		"allow_end_users_to_change_their_password":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"allow_end_users_to_link_their_support_account":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"allow_end_users_to_manage_their_capabilities":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"allow_end_users_to_update_their_security_settings": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsIdentitySettingTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsIdentitySettingTokensRepresentation = map[string]interface{}{
		"type":          acctest.Representation{RepType: acctest.Required, Create: `emailVerification`},
		"expires_after": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}
	IdentityDomainsIdentitySettingTokensRepresentation1 = map[string]interface{}{
		"type":          acctest.Representation{RepType: acctest.Required, Create: `passwordReset`},
		"expires_after": acctest.Representation{RepType: acctest.Required, Create: `120`, Update: `240`},
	}
	IdentityDomainsIdentitySettingTokensRepresentation2 = map[string]interface{}{
		"type":          acctest.Representation{RepType: acctest.Required, Create: `createUser`},
		"expires_after": acctest.Representation{RepType: acctest.Required, Create: `2500`, Update: `3000`},
	}

	IdentityDomainsIdentitySettingResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsIdentitySettingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsIdentitySettingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_identity_setting.test_identity_setting"
	datasourceName := "data.oci_identity_domains_identity_settings.test_identity_settings"
	singularDatasourceName := "data.oci_identity_domains_identity_setting.test_identity_setting"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsIdentitySettingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_setting", "test_identity_setting", acctest.Optional, acctest.Create, IdentityDomainsIdentitySettingRepresentation), "identitydomains", "identitySetting", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentitySettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_setting", "test_identity_setting", acctest.Required, acctest.Create, IdentityDomainsIdentitySettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "identity_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentitySettingResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentitySettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_setting", "test_identity_setting", acctest.Optional, acctest.Create, IdentityDomainsIdentitySettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "posix_gid.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "posix_gid.0.manual_assignment_ends_at", "1000"),
				resource.TestCheckResourceAttr(resourceName, "posix_gid.0.manual_assignment_starts_from", "10"),
				resource.TestCheckResourceAttr(resourceName, "posix_uid.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "posix_uid.0.manual_assignment_ends_at", "1000"),
				resource.TestCheckResourceAttr(resourceName, "posix_uid.0.manual_assignment_starts_from", "10"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "emit_locked_message_when_user_is_locked", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "id", "IdentitySettings"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "identity_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "primary_email_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "tokens.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "tokens.0.expires_after", "10"),
				resource.TestCheckResourceAttr(resourceName, "tokens.0.type", "emailVerification"),
				resource.TestCheckResourceAttr(resourceName, "user_allowed_to_set_recovery_email", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "identitySettings", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentitySettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_setting", "test_identity_setting", acctest.Optional, acctest.Update, IdentityDomainsIdentitySettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "posix_gid.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "posix_gid.0.manual_assignment_ends_at", "2000"),
				resource.TestCheckResourceAttr(resourceName, "posix_gid.0.manual_assignment_starts_from", "11"),
				resource.TestCheckResourceAttr(resourceName, "posix_uid.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "posix_uid.0.manual_assignment_ends_at", "2000"),
				resource.TestCheckResourceAttr(resourceName, "posix_uid.0.manual_assignment_starts_from", "11"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "emit_locked_message_when_user_is_locked", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "id", "IdentitySettings"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "identity_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "primary_email_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "return_inactive_over_locked_message", "true"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "tokens.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "tokens.0.expires_after", "11"),
				resource.TestCheckResourceAttr(resourceName, "tokens.1.expires_after", "240"),
				resource.TestCheckResourceAttr(resourceName, "tokens.2.expires_after", "3000"),
				resource.TestCheckResourceAttr(resourceName, "user_allowed_to_set_recovery_email", "true"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_identity_settings", "test_identity_settings", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsIdentitySettingDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsIdentitySettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_setting", "test_identity_setting", acctest.Optional, acctest.Update, IdentityDomainsIdentitySettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "identity_settings.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "identity_settings.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_identity_setting", "test_identity_setting", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsIdentitySettingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsIdentitySettingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "identity_setting_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "posix_gid.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "posix_gid.0.manual_assignment_ends_at", "2000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "posix_gid.0.manual_assignment_starts_from", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "posix_uid.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "posix_uid.0.manual_assignment_ends_at", "2000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "posix_uid.0.manual_assignment_starts_from", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "emit_locked_message_when_user_is_locked", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "id", "IdentitySettings"),
				resource.TestCheckResourceAttr(singularDatasourceName, "primary_email_required", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "return_inactive_over_locked_message", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tokens.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tokens.0.expires_after", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tokens.0.type", "emailVerification"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_allowed_to_set_recovery_email", "true"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsIdentitySettingRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_identity_setting", "identitySettings"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"idcs_last_upgraded_in_release",
				"tags",
				"identity_setting_id",
			},
			ResourceName: resourceName,
		},
	})
}
