// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

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
	IdentityDomainsUserRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Required, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(IdentityDomainsUserRepresentation, []string{"password"}), map[string]interface{}{
			"password": acctest.Representation{RepType: acctest.Required, Create: `toBeIgnored_#15`},
		}),
	)

	IdentityDomainsUserResourceConfig = IdentityDomainsUserResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Optional, acctest.Update, IdentityDomainsUserRepresentation)

	IdentityDomainsIdentityDomainsUserSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"user_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_user.test_user.id}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsUserDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"user_count":     acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"user_filter":    acctest.Representation{RepType: acctest.Optional, Create: `userName sw \"userName\"`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":    acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsUserRepresentation = map[string]interface{}{
		"idcs_endpoint":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":            acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:core:2.0:User`}},
		"user_name":          acctest.Representation{RepType: acctest.Required, Create: `userName`},
		"active":             acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"addresses":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserAddressesRepresentation},
		"attribute_sets":     acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"emails":             acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsUserEmailsRepresentation},
		"entitlements":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserEntitlementsRepresentation},
		"external_id":        acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"ims":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserImsRepresentation},
		"locale":             acctest.Representation{RepType: acctest.Optional, Create: `en`, Update: `es`},
		"name":               acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsUserNameRepresentation},
		"nick_name":          acctest.Representation{RepType: acctest.Optional, Create: `nickName`, Update: `nickName2`},
		"password":           acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"phone_numbers":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserPhoneNumbersRepresentation},
		"photos":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserPhotosRepresentation},
		"preferred_language": acctest.Representation{RepType: acctest.Optional, Create: `en`, Update: `es`},
		"profile_url":        acctest.Representation{RepType: acctest.Optional, Create: `https://profileUrl.com`, Update: `https://profileUrl2.com`},
		"roles":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserRolesRepresentation},
		"tags":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserTagsRepresentation},
		"timezone":           acctest.Representation{RepType: acctest.Optional, Create: `America/Los_Angeles`, Update: `America/Vancouver`},
		"title":              acctest.Representation{RepType: acctest.Optional, Create: `title`, Update: `title2`},
		"urnietfparamsscimschemasextensionenterprise20user":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasextensionenterprise20UserRepresentation},
		"urnietfparamsscimschemasoracleidcsextension_oci_tags":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionOCITagsRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionadaptive_user":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionadaptiveUserRepresentation},
		"urnietfparamsscimschemasoracleidcsextensioncapabilities_user":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensioncapabilitiesUserRepresentation},
		"urnietfparamsscimschemasoracleidcsextensiondb_credentials_user": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensiondbCredentialsUserRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionpasswordless_user":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionpasswordlessUserRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionposix_user":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionposixUserRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionself_change_user":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionselfChangeUserRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionsff_user":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionsffUserRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionuser_state_user":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionuserStateUserRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionuser_user":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionuserUserRepresentation},
		"user_type":        acctest.Representation{RepType: acctest.Optional, Create: `Contractor`, Update: `Employee`},
		"x509certificates": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserX509CertificatesRepresentation},
		"lifecycle":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangeForIdentityDomainsUser},
	}
	ignoreChangeForIdentityDomainsUser = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			`urnietfparamsscimschemasoracleidcsextension_oci_tags[0].defined_tags`,
			`emails`,
			`schemas`,
			`urnietfparamsscimschemasoracleidcsextensionself_change_user`,
			`urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user[0].sec_questions`,
		}},
	}

	IdentityDomainsUserAddressesRepresentation = map[string]interface{}{
		"type":           acctest.Representation{RepType: acctest.Required, Create: `work`, Update: `home`},
		"country":        acctest.Representation{RepType: acctest.Optional, Create: `us`, Update: `gb`},
		"formatted":      acctest.Representation{RepType: acctest.Optional, Create: `formatted`, Update: `formatted2`},
		"locality":       acctest.Representation{RepType: acctest.Optional, Create: `locality`, Update: `locality2`},
		"postal_code":    acctest.Representation{RepType: acctest.Optional, Create: `postalCode`, Update: `postalCode2`},
		"primary":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"region":         acctest.Representation{RepType: acctest.Optional, Create: `region`, Update: `region2`},
		"street_address": acctest.Representation{RepType: acctest.Optional, Create: `streetAddress`, Update: `streetAddress2`},
	}
	IdentityDomainsUserEmailsRepresentation = map[string]interface{}{
		"type":      acctest.Representation{RepType: acctest.Required, Create: `work`, Update: `home`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `value@email.com`, Update: `value2@email.com`},
		"primary":   acctest.Representation{RepType: acctest.Required, Create: `true`},
		"secondary": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"verified":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsUserEntitlementsRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `type`, Update: `type2`},
		"value":   acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
		"display": acctest.Representation{RepType: acctest.Optional, Create: `display`, Update: `display2`},
		"primary": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsUserImsRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `aim`, Update: `gtalk`},
		"value":   acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
		"display": acctest.Representation{RepType: acctest.Optional, Create: `display`, Update: `display2`},
		"primary": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsUserNameRepresentation = map[string]interface{}{
		"family_name":      acctest.Representation{RepType: acctest.Required, Create: `familyName`, Update: `familyName2`},
		"formatted":        acctest.Representation{RepType: acctest.Optional, Create: `formatted`, Update: `formatted2`},
		"given_name":       acctest.Representation{RepType: acctest.Optional, Create: `givenName`, Update: `givenName2`},
		"honorific_prefix": acctest.Representation{RepType: acctest.Optional, Create: `honorificPrefix`, Update: `honorificPrefix2`},
		"honorific_suffix": acctest.Representation{RepType: acctest.Optional, Create: `honorificSuffix`, Update: `honorificSuffix2`},
		"middle_name":      acctest.Representation{RepType: acctest.Optional, Create: `middleName`, Update: `middleName2`},
	}
	IdentityDomainsUserPhoneNumbersRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `work`, Update: `home`},
		"value":   acctest.Representation{RepType: acctest.Required, Create: `1112223333`, Update: `1112223334`},
		"primary": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsUserPhotosRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `photo`, Update: `thumbnail`},
		"value":   acctest.Representation{RepType: acctest.Required, Create: `https://value.com`, Update: `https://value2.com`},
		"display": acctest.Representation{RepType: acctest.Optional, Create: `display`, Update: `display2`},
		"primary": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsUserRolesRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `type`, Update: `type2`},
		"value":   acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
		"display": acctest.Representation{RepType: acctest.Optional, Create: `display`, Update: `display2`},
		"primary": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsUserTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsUserUrnietfparamsscimschemasextensionenterprise20UserRepresentation = map[string]interface{}{
		"cost_center":     acctest.Representation{RepType: acctest.Optional, Create: `costCenter`, Update: `costCenter2`},
		"department":      acctest.Representation{RepType: acctest.Optional, Create: `department`, Update: `department2`},
		"division":        acctest.Representation{RepType: acctest.Optional, Create: `division`, Update: `division2`},
		"employee_number": acctest.Representation{RepType: acctest.Optional, Create: `employeeNumber`, Update: `employeeNumber2`},
		"manager":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasextensionenterprise20UserManagerRepresentation},
		"organization":    acctest.Representation{RepType: acctest.Optional, Create: `organization`, Update: `organization2`},
	}
	IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionOCITagsRepresentation = map[string]interface{}{
		"defined_tags":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionOCITagsDefinedTagsRepresentation},
		"freeform_tags": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionOCITagsFreeformTagsRepresentation},
	}
	IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionadaptiveUserRepresentation = map[string]interface{}{
		"risk_level": acctest.Representation{RepType: acctest.Optional, Create: `LOW`, Update: `MEDIUM`},
	}
	IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensioncapabilitiesUserRepresentation = map[string]interface{}{
		"can_use_api_keys":                 acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"can_use_auth_tokens":              acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"can_use_console_password":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"can_use_customer_secret_keys":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"can_use_db_credentials":           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"can_use_oauth2client_credentials": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"can_use_smtp_credentials":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensiondbCredentialsUserRepresentation = map[string]interface{}{
		"db_user_name": acctest.Representation{RepType: acctest.Optional, Create: `dbUserName`},
	}
	IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionpasswordlessUserRepresentation = map[string]interface{}{
		"factor_method": acctest.Representation{RepType: acctest.Optional, Create: `factorMethod`, Update: `factorMethod2`},
		"factor_type":   acctest.Representation{RepType: acctest.Optional, Create: `EMAIL`, Update: `SMS`},
	}
	IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionposixUserRepresentation = map[string]interface{}{
		"gecos":          acctest.Representation{RepType: acctest.Optional, Create: `gecos`, Update: `gecos2`},
		"home_directory": acctest.Representation{RepType: acctest.Optional, Create: `homeDirectory`, Update: `homeDirectory2`},
		"login_shell":    acctest.Representation{RepType: acctest.Optional, Create: `loginShell`, Update: `loginShell2`},
		"uid_number":     acctest.Representation{RepType: acctest.Optional, Create: `500`, Update: `501`},
	}
	IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionselfChangeUserRepresentation = map[string]interface{}{
		"allow_self_change": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionsffUserRepresentation = map[string]interface{}{
		"sff_auth_keys": acctest.Representation{RepType: acctest.Optional, Create: `sffAuthKeys`, Update: `sffAuthKeys2`},
	}
	IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionuserStateUserRepresentation = map[string]interface{}{
		"max_concurrent_sessions": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionuserUserRepresentation = map[string]interface{}{
		"user_provider":                              acctest.Representation{RepType: acctest.Optional, Create: `facebook`},
		"account_recovery_required":                  acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"bypass_notification":                        acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"creation_mechanism":                         acctest.Representation{RepType: acctest.Optional, Create: `api`},
		"do_not_show_getting_started":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_authentication_delegated":                acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_federated_user":                          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_group_membership_normalized":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_group_membership_synced_to_users_groups": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"service_user":                               acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"user_flow_controlled_by_external_client":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	IdentityDomainsUserX509CertificatesRepresentation = map[string]interface{}{
		"value":   acctest.Representation{RepType: acctest.Required, Create: `MIIBPDCB56ADAgECAhAucicVzKJi9WsJknMUyFhRMA0GCSqGSIb3DQEBBAUAMA4xDDAKBgNVBAMTA29pZDAeFw0wNDA1MjUwMDEwMjNaFw0wNDExMjEwMDEwMjNaMDExCzAJBgNVBAYTAnVzMQwwCgYDVQQKEwNJTUMxFDASBgNVBAMTC0thbCBCYWlsZXlzMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAMqsqm5/ZE2Fy57YmUNC6Lc6j70z+DNhmCADhsxpA4DXwrOkDDaIpAXG45y4NvsImjjpaFGxSE0upxBIQAHj9CMCAwEAATANBgkqhkiG9w0BAQQFAANBAGktTIBlB3VyN+7a9mRzdeYgS8ZwVsee1iGVRHCTfF1quxtVyWVwMX0dxffwz6pK0Pm3bV7uiEVu5qf3rO1hYSE=`},
		"display": acctest.Representation{RepType: acctest.Optional, Create: `display`, Update: `display2`},
		"primary": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"type":    acctest.Representation{RepType: acctest.Optional, Create: `type`, Update: `type2`},
	}
	IdentityDomainsUserUrnietfparamsscimschemasextensionenterprise20UserManagerRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_domains_user.test_user_manager.id}`},
	}
	IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionOCITagsDefinedTagsRepresentation = map[string]interface{}{
		"key":       acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag.tag1.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag_namespace.tag-namespace1.name}`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsUserUrnietfparamsscimschemasoracleidcsextensionOCITagsFreeformTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `freeformKey`, Update: `freeformKey2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `freeformValue`, Update: `freeformValue2`},
	}

	IdentityDomainsUserResourceDependencies = DefinedTagsDependencies + TestDomainDependencies + IdentityDomainsUserManager
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_user.test_user"
	datasourceName := "data.oci_identity_domains_users.test_users"
	singularDatasourceName := "data.oci_identity_domains_user.test_user"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsUserResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Optional, acctest.Create, IdentityDomainsUserRepresentation), "identitydomains", "user", t)

	print(config + compartmentIdVariableStr + IdentityDomainsUserResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Optional, acctest.Create, IdentityDomainsUserRepresentation))

	acctest.ResourceTest(t, testAccCheckIdentityDomainsUserDestroy, []resource.TestStep{
		// verify Create with required fields
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Required, acctest.Create, IdentityDomainsUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name.0.family_name", "familyName"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update with required fields
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Required, acctest.Update, IdentityDomainsUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name.0.family_name", "familyName2"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsUserResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Optional, acctest.Create, IdentityDomainsUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active", "true"),
				resource.TestCheckResourceAttr(resourceName, "addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.country", "us"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.formatted", "formatted"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.locality", "locality"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.postal_code", "postalCode"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.primary", "false"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.region", "region"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.street_address", "streetAddress"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.type", "work"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "emails.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "emails.0.value", "value@email.com"),
				resource.TestCheckResourceAttr(resourceName, "emails.1.value", "value@email.com"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.display", "display"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.primary", "false"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.type", "type"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "ims.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ims.0.display", "display"),
				resource.TestCheckResourceAttr(resourceName, "ims.0.primary", "false"),
				resource.TestCheckResourceAttr(resourceName, "ims.0.type", "aim"),
				resource.TestCheckResourceAttr(resourceName, "ims.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "locale", "en"),
				resource.TestCheckResourceAttr(resourceName, "name.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name.0.family_name", "familyName"),
				resource.TestCheckResourceAttr(resourceName, "name.0.formatted", "formatted"),
				resource.TestCheckResourceAttr(resourceName, "name.0.given_name", "givenName"),
				resource.TestCheckResourceAttr(resourceName, "name.0.honorific_prefix", "honorificPrefix"),
				resource.TestCheckResourceAttr(resourceName, "name.0.honorific_suffix", "honorificSuffix"),
				resource.TestCheckResourceAttr(resourceName, "name.0.middle_name", "middleName"),
				resource.TestCheckResourceAttr(resourceName, "nick_name", "nickName"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "phone_numbers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "phone_numbers.0.primary", "false"),
				resource.TestCheckResourceAttr(resourceName, "phone_numbers.0.type", "work"),
				resource.TestCheckResourceAttr(resourceName, "phone_numbers.0.value", "1112223333"),
				resource.TestCheckResourceAttr(resourceName, "photos.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "photos.0.display", "display"),
				resource.TestCheckResourceAttr(resourceName, "photos.0.primary", "false"),
				resource.TestCheckResourceAttr(resourceName, "photos.0.type", "photo"),
				resource.TestCheckResourceAttr(resourceName, "photos.0.value", "https://value.com"),
				resource.TestCheckResourceAttr(resourceName, "preferred_language", "en"),
				resource.TestCheckResourceAttr(resourceName, "profile_url", "https://profileUrl.com"),
				resource.TestCheckResourceAttr(resourceName, "roles.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "roles.0.display", "display"),
				resource.TestCheckResourceAttr(resourceName, "roles.0.primary", "false"),
				resource.TestCheckResourceAttr(resourceName, "roles.0.type", "type"),
				resource.TestCheckResourceAttr(resourceName, "roles.0.value", "value"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "timezone", "America/Los_Angeles"),
				resource.TestCheckResourceAttr(resourceName, "title", "title"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.cost_center", "costCenter"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.department", "department"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.division", "division"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.employee_number", "employeeNumber"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.manager.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.manager.0.value"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.organization", "organization"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.key", "freeformKey"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.value", "freeformValue"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionadaptive_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionadaptive_user.0.risk_level", "LOW"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_api_keys", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_auth_tokens", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_console_password", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_customer_secret_keys", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_db_credentials", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_oauth2client_credentials", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_smtp_credentials", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensiondb_credentials_user.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensiondb_credentials_user.0.db_user_name"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionpasswordless_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionpasswordless_user.0.factor_method", "factorMethod"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionpasswordless_user.0.factor_type", "EMAIL"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_user.0.gecos", "gecos"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_user.0.home_directory", "homeDirectory"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_user.0.login_shell", "loginShell"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_user.0.uid_number", "500"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsff_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsff_user.0.sff_auth_keys", "sffAuthKeys"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_state_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_state_user.0.max_concurrent_sessions", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.user_provider", "facebook"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.account_recovery_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.bypass_notification", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.creation_mechanism", "api"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.do_not_show_getting_started", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.is_authentication_delegated", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.is_federated_user", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.is_group_membership_normalized", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.is_group_membership_synced_to_users_groups", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.user_flow_controlled_by_external_client", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				resource.TestCheckResourceAttr(resourceName, "user_type", "Contractor"),
				resource.TestCheckResourceAttr(resourceName, "x509certificates.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "x509certificates.0.display", "display"),
				resource.TestCheckResourceAttr(resourceName, "x509certificates.0.primary", "false"),
				resource.TestCheckResourceAttr(resourceName, "x509certificates.0.type", "type"),
				resource.TestCheckResourceAttr(resourceName, "x509certificates.0.value", "MIIBPDCB56ADAgECAhAucicVzKJi9WsJknMUyFhRMA0GCSqGSIb3DQEBBAUAMA4xDDAKBgNVBAMTA29pZDAeFw0wNDA1MjUwMDEwMjNaFw0wNDExMjEwMDEwMjNaMDExCzAJBgNVBAYTAnVzMQwwCgYDVQQKEwNJTUMxFDASBgNVBAMTC0thbCBCYWlsZXlzMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAMqsqm5/ZE2Fy57YmUNC6Lc6j70z+DNhmCADhsxpA4DXwrOkDDaIpAXG45y4NvsImjjpaFGxSE0upxBIQAHj9CMCAwEAATANBgkqhkiG9w0BAQQFAANBAGktTIBlB3VyN+7a9mRzdeYgS8ZwVsee1iGVRHCTfF1quxtVyWVwMX0dxffwz6pK0Pm3bV7uiEVu5qf3rO1hYSE="),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "users", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Optional, acctest.Update, IdentityDomainsUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active", "true"),
				resource.TestCheckResourceAttr(resourceName, "addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.country", "gb"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.formatted", "formatted2"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.locality", "locality2"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.postal_code", "postalCode2"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.primary", "true"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.region", "region2"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.street_address", "streetAddress2"),
				resource.TestCheckResourceAttr(resourceName, "addresses.0.type", "home"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "emails.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.display", "display2"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.primary", "true"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.type", "type2"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.value", "value2"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "ims.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ims.0.display", "display2"),
				resource.TestCheckResourceAttr(resourceName, "ims.0.primary", "true"),
				resource.TestCheckResourceAttr(resourceName, "ims.0.type", "gtalk"),
				resource.TestCheckResourceAttr(resourceName, "ims.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "locale", "es"),
				resource.TestCheckResourceAttr(resourceName, "name.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name.0.family_name", "familyName2"),
				resource.TestCheckResourceAttr(resourceName, "name.0.formatted", "formatted2"),
				resource.TestCheckResourceAttr(resourceName, "name.0.given_name", "givenName2"),
				resource.TestCheckResourceAttr(resourceName, "name.0.honorific_prefix", "honorificPrefix2"),
				resource.TestCheckResourceAttr(resourceName, "name.0.honorific_suffix", "honorificSuffix2"),
				resource.TestCheckResourceAttr(resourceName, "name.0.middle_name", "middleName2"),
				resource.TestCheckResourceAttr(resourceName, "nick_name", "nickName2"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "phone_numbers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "phone_numbers.0.primary", "true"),
				resource.TestCheckResourceAttr(resourceName, "phone_numbers.0.type", "home"),
				resource.TestCheckResourceAttr(resourceName, "phone_numbers.0.value", "1112223334"),
				resource.TestCheckResourceAttr(resourceName, "photos.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "photos.0.display", "display2"),
				resource.TestCheckResourceAttr(resourceName, "photos.0.primary", "true"),
				resource.TestCheckResourceAttr(resourceName, "photos.0.type", "thumbnail"),
				resource.TestCheckResourceAttr(resourceName, "photos.0.value", "https://value2.com"),
				resource.TestCheckResourceAttr(resourceName, "preferred_language", "es"),
				resource.TestCheckResourceAttr(resourceName, "profile_url", "https://profileUrl2.com"),
				resource.TestCheckResourceAttr(resourceName, "roles.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "roles.0.display", "display2"),
				resource.TestCheckResourceAttr(resourceName, "roles.0.primary", "true"),
				resource.TestCheckResourceAttr(resourceName, "roles.0.type", "type2"),
				resource.TestCheckResourceAttr(resourceName, "roles.0.value", "value2"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "timezone", "America/Vancouver"),
				resource.TestCheckResourceAttr(resourceName, "title", "title2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.cost_center", "costCenter2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.department", "department2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.division", "division2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.employee_number", "employeeNumber2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.manager.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.manager.0.value"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasextensionenterprise20user.0.organization", "organization2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.key", "freeformKey2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.value", "freeformValue2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionadaptive_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionadaptive_user.0.risk_level", "MEDIUM"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_api_keys", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_auth_tokens", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_console_password", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_customer_secret_keys", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_db_credentials", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_oauth2client_credentials", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_smtp_credentials", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensiondb_credentials_user.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensiondb_credentials_user.0.db_user_name"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionpasswordless_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionpasswordless_user.0.factor_method", "factorMethod2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionpasswordless_user.0.factor_type", "SMS"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_user.0.gecos", "gecos2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_user.0.home_directory", "homeDirectory2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_user.0.login_shell", "loginShell2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_user.0.uid_number", "501"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsff_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsff_user.0.sff_auth_keys", "sffAuthKeys2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_state_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_state_user.0.max_concurrent_sessions", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.account_recovery_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.applicable_authentication_target_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.creation_mechanism", "api"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.do_not_show_getting_started", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.is_authentication_delegated", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.is_federated_user", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.is_group_membership_normalized", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.is_group_membership_synced_to_users_groups", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.user_flow_controlled_by_external_client", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				resource.TestCheckResourceAttr(resourceName, "user_type", "Employee"),
				resource.TestCheckResourceAttr(resourceName, "x509certificates.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "x509certificates.0.display", "display2"),
				resource.TestCheckResourceAttr(resourceName, "x509certificates.0.primary", "true"),
				resource.TestCheckResourceAttr(resourceName, "x509certificates.0.type", "type2"),
				resource.TestCheckResourceAttr(resourceName, "x509certificates.0.value", "MIIBPDCB56ADAgECAhAucicVzKJi9WsJknMUyFhRMA0GCSqGSIb3DQEBBAUAMA4xDDAKBgNVBAMTA29pZDAeFw0wNDA1MjUwMDEwMjNaFw0wNDExMjEwMDEwMjNaMDExCzAJBgNVBAYTAnVzMQwwCgYDVQQKEwNJTUMxFDASBgNVBAMTC0thbCBCYWlsZXlzMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAMqsqm5/ZE2Fy57YmUNC6Lc6j70z+DNhmCADhsxpA4DXwrOkDDaIpAXG45y4NvsImjjpaFGxSE0upxBIQAHj9CMCAwEAATANBgkqhkiG9w0BAQQFAANBAGktTIBlB3VyN+7a9mRzdeYgS8ZwVsee1iGVRHCTfF1quxtVyWVwMX0dxffwz6pK0Pm3bV7uiEVu5qf3rO1hYSE="),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_users", "test_users", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsUserDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Optional, acctest.Update, IdentityDomainsUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "user_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),
				resource.TestCheckResourceAttr(datasourceName, "total_results", "1"),

				resource.TestCheckResourceAttr(datasourceName, "users.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsUserSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "active", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "addresses.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "addresses.0.country", "gb"),
				resource.TestCheckResourceAttr(singularDatasourceName, "addresses.0.formatted", "formatted2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "addresses.0.locality", "locality2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "addresses.0.postal_code", "postalCode2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "addresses.0.primary", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "addresses.0.region", "region2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "addresses.0.street_address", "streetAddress2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "addresses.0.type", "home"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "emails.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.display", "display2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.primary", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.type", "type2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ims.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ims.0.display", "display2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ims.0.primary", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ims.0.type", "gtalk"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ims.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locale", "es"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name.0.family_name", "familyName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name.0.formatted", "formatted2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name.0.given_name", "givenName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name.0.honorific_prefix", "honorificPrefix2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name.0.honorific_suffix", "honorificSuffix2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name.0.middle_name", "middleName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nick_name", "nickName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phone_numbers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phone_numbers.0.primary", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phone_numbers.0.type", "home"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phone_numbers.0.value", "1112223334"),
				resource.TestCheckResourceAttr(singularDatasourceName, "photos.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "photos.0.display", "display2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "photos.0.primary", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "photos.0.type", "thumbnail"),
				resource.TestCheckResourceAttr(singularDatasourceName, "photos.0.value", "https://value2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "preferred_language", "es"),
				resource.TestCheckResourceAttr(singularDatasourceName, "profile_url", "https://profileUrl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "roles.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "roles.0.display", "display2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "roles.0.primary", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "roles.0.type", "type2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "roles.0.value", "value2"),
				resource.TestMatchResourceAttr(singularDatasourceName, "schemas.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestCheckResourceAttr(singularDatasourceName, "timezone", "America/Vancouver"),
				resource.TestCheckResourceAttr(singularDatasourceName, "title", "title2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasextensionenterprise20user.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasextensionenterprise20user.0.cost_center", "costCenter2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasextensionenterprise20user.0.department", "department2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasextensionenterprise20user.0.division", "division2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasextensionenterprise20user.0.employee_number", "employeeNumber2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasextensionenterprise20user.0.manager.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "urnietfparamsscimschemasextensionenterprise20user.0.manager.0.value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasextensionenterprise20user.0.organization", "organization2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.key", "freeformKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.value", "freeformValue2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_api_keys", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_auth_tokens", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_console_password", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_customer_secret_keys", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_db_credentials", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_oauth2client_credentials", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensioncapabilities_user.0.can_use_smtp_credentials", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionpasswordless_user.0.factor_method", "factorMethod2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionpasswordless_user.0.factor_type", "SMS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionuser_state_user.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionuser_state_user.0.max_concurrent_sessions", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.user_provider", "facebook"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.account_recovery_required", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.bypass_notification", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.do_not_show_getting_started", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.is_authentication_delegated", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.is_federated_user", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.is_group_membership_normalized", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.is_group_membership_synced_to_users_groups", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionuser_user.0.user_flow_controlled_by_external_client", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_type", "Employee"),
				resource.TestCheckResourceAttr(singularDatasourceName, "x509certificates.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "x509certificates.0.display", "display2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "x509certificates.0.primary", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "x509certificates.0.type", "type2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "x509certificates.0.value", "MIIBPDCB56ADAgECAhAucicVzKJi9WsJknMUyFhRMA0GCSqGSIb3DQEBBAUAMA4xDDAKBgNVBAMTA29pZDAeFw0wNDA1MjUwMDEwMjNaFw0wNDExMjEwMDEwMjNaMDExCzAJBgNVBAYTAnVzMQwwCgYDVQQKEwNJTUMxFDASBgNVBAMTC0thbCBCYWlsZXlzMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAMqsqm5/ZE2Fy57YmUNC6Lc6j70z+DNhmCADhsxpA4DXwrOkDDaIpAXG45y4NvsImjjpaFGxSE0upxBIQAHj9CMCAwEAATANBgkqhkiG9w0BAQQFAANBAGktTIBlB3VyN+7a9mRzdeYgS8ZwVsee1iGVRHCTfF1quxtVyWVwMX0dxffwz6pK0Pm3bV7uiEVu5qf3rO1hYSE="),
			),
		},

		// reset to required only resource for import
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsUserResourceDependencies + IdentityDomainsUserRequiredOnlyResource,
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsUserRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_user", "users"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"password",
				"schemas",
			},
			ResourceName: resourceName,
		},
		// dependency manager user needs to be removed after test user
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsUserResourceDependencies,
		},
	})
}

func testAccCheckIdentityDomainsUserDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_user" {
			noResourceFound = false
			request := oci_identity_domains.GetUserRequest{}

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
			request.UserId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetUser(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsUser") {
		resource.AddTestSweepers("IdentityDomainsUser", &resource.Sweeper{
			Name:         "IdentityDomainsUser",
			Dependencies: acctest.DependencyGraph["user"],
			F:            sweepIdentityDomainsUserResource,
		})
	}
}

func sweepIdentityDomainsUserResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	userIds, err := getIdentityDomainsUserIds(compartment)
	if err != nil {
		return err
	}
	for _, userId := range userIds {
		if ok := acctest.SweeperDefaultResourceId[userId]; !ok {
			deleteUserRequest := oci_identity_domains.DeleteUserRequest{}

			deleteUserRequest.UserId = &userId

			deleteUserRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteUser(context.Background(), deleteUserRequest)
			if error != nil {
				fmt.Printf("Error deleting User %s %s, It is possible that the resource is already deleted. Please verify manually \n", userId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsUserIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "UserId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listUsersRequest := oci_identity_domains.ListUsersRequest{}
	listUsersResponse, err := identityDomainsClient.ListUsers(context.Background(), listUsersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting User list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, user := range listUsersResponse.Resources {
		id := *user.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "UserId", id)
	}
	return resourceIds, nil
}
