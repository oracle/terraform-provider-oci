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
	IdentityDomainsAppRequiredOnlyResource = IdentityDomainsAppResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app", "test_app", acctest.Required, acctest.Create, IdentityDomainsAppRepresentation)

	IdentityDomainsAppResourceConfig = IdentityDomainsAppResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app", "test_app", acctest.Optional, acctest.Update, IdentityDomainsAppRepresentation)

	IdentityDomainsIdentityDomainsAppSingularDataSourceRepresentation = map[string]interface{}{
		"app_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_app.test_app.id}`},
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsAppDataSourceRepresentation = map[string]interface{}{
		"app_count":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"start_index":    acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsAppRepresentation = map[string]interface{}{
		"based_on_template":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsAppBasedOnTemplateRepresentation},
		"display_name":                      acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"idcs_endpoint":                     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":                           acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:App`}},
		"access_token_expiry":               acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"active":                            acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"all_url_schemes_allowed":           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"allow_access_control":              acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"allow_offline":                     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"allowed_grants":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`authorization_code`}, Update: []string{`client_credentials`}},
		"allowed_operations":                acctest.Representation{RepType: acctest.Optional, Create: []string{`introspect`}, Update: []string{`onBehalfOfUser`}},
		"allowed_tags":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppAllowedTagsRepresentation},
		"app_icon":                          acctest.Representation{RepType: acctest.Optional, Create: `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=`, Update: `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO+ip1sAAAAASUVORK5CYII=`},
		"app_thumbnail":                     acctest.Representation{RepType: acctest.Optional, Create: `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=`, Update: `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO+ip1sAAAAASUVORK5CYII=`},
		"attr_rendering_metadata":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppAttrRenderingMetadataRepresentation},
		"attribute_sets":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"audience":                          acctest.Representation{RepType: acctest.Optional, Create: `audience`, Update: `audience2`},
		"bypass_consent":                    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"client_ip_checking":                acctest.Representation{RepType: acctest.Optional, Create: `anywhere`, Update: `whitelisted`},
		"client_type":                       acctest.Representation{RepType: acctest.Optional, Create: `confidential`, Update: `trusted`},
		"contact_email_address":             acctest.Representation{RepType: acctest.Optional, Create: `contact@email.com`, Update: `contact2@email.com`},
		"delegated_service_names":           acctest.Representation{RepType: acctest.Optional, Create: []string{`delegatedServiceNames`}, Update: []string{`delegatedServiceNames2`}},
		"description":                       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"disable_kmsi_token_authentication": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"error_page_url":                    acctest.Representation{RepType: acctest.Optional, Create: `https://testurl.com`, Update: `https://testurl2.com`},
		"home_page_url":                     acctest.Representation{RepType: acctest.Optional, Create: `https://testurl.com`, Update: `https://testurl2.com`},
		"icon":                              acctest.Representation{RepType: acctest.Optional, Create: `icon`, Update: `icon2`},
		"id_token_enc_algo":                 acctest.Representation{RepType: acctest.Optional, Create: `A128CBC-HS256`, Update: `A192CBC-HS384`},
		"identity_providers":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppIdentityProvidersRepresentation},
		"is_alias_app":                      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_enterprise_app":                 acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_form_fill":                      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_kerberos_realm":                 acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_login_target":                   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_mobile_target":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_multicloud_service_app":         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_oauth_client":                   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_oauth_resource":                 acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_obligation_capable":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_radius_app":                     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_saml_service_provider":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_unmanaged_app":                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_web_tier_policy":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"landing_page_url":                  acctest.Representation{RepType: acctest.Optional, Create: `https://testurl.com`, Update: `https://testurl2.com`},
		"linking_callback_url":              acctest.Representation{RepType: acctest.Optional, Create: `https://testurl.com`, Update: `https://testurl2.com`},
		"login_mechanism":                   acctest.Representation{RepType: acctest.Optional, Create: `OIDC`, Update: `SAML`},
		"login_page_url":                    acctest.Representation{RepType: acctest.Optional, Create: `https://testurl.com`, Update: `https://testurl2.com`},
		"logout_page_url":                   acctest.Representation{RepType: acctest.Optional, Create: `https://testurl.com`, Update: `https://testurl2.com`},
		"logout_uri":                        acctest.Representation{RepType: acctest.Optional, Create: `logoutUri`, Update: `logoutUri2`},
		"name":                              acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"post_logout_redirect_uris":         acctest.Representation{RepType: acctest.Optional, Create: []string{`postLogoutRedirectUris`}, Update: []string{`postLogoutRedirectUris2`}},
		"privacy_policy_url":                acctest.Representation{RepType: acctest.Optional, Create: `https://testurl.com`, Update: `https://testurl2.com`},
		"product_logo_url":                  acctest.Representation{RepType: acctest.Optional, Create: `https://testurl.com`, Update: `https://testurl2.com`},
		"product_name":                      acctest.Representation{RepType: acctest.Optional, Create: `productName`, Update: `productName2`},
		"protectable_secondary_audiences":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppProtectableSecondaryAudiencesRepresentation},
		"redirect_uris":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`redirectUris`}, Update: []string{`redirectUris2`}},
		"refresh_token_expiry":              acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"saml_service_provider":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppSamlServiceProviderRepresentation},
		"scopes":                            acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppScopesRepresentation},
		"secondary_audiences":               acctest.Representation{RepType: acctest.Optional, Create: []string{`secondaryAudiences`}, Update: []string{`secondaryAudiences2`}},
		"service_params":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppServiceParamsRepresentation},
		"service_type_urn":                  acctest.Representation{RepType: acctest.Optional, Create: `serviceTypeURN`, Update: `serviceTypeURN2`},
		"service_type_version":              acctest.Representation{RepType: acctest.Optional, Create: `serviceTypeVersion`, Update: `serviceTypeVersion2`},
		"show_in_my_apps":                   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"tags":                              acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppTagsRepresentation},
		"terms_of_service_url":              acctest.Representation{RepType: acctest.Optional, Create: `https://testurl.com`, Update: `https://testurl2.com`},
		"trust_scope":                       acctest.Representation{RepType: acctest.Optional, Create: `Explicit`, Update: `Account`},
		"urnietfparamsscimschemasoracleidcsextension_oci_tags":                           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionOCITagsRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionenterprise_app_app":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionenterpriseAppAppRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionform_fill_app_app":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionformFillAppAppRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionformFillAppTemplateAppTemplateRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionkerberosRealmAppRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionmanagedapp_app":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionmanagedappAppRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionmulticloudServiceAppAppRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionopc_service_app":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionopcServiceAppRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionradius_app_app":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionradiusAppAppRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionrequestable_app":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionrequestableAppRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionsamlServiceProviderAppRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionwebTierPolicyAppRepresentation},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangeForIdentityDomainsApp},
	}

	ignoreChangeForIdentityDomainsApp = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			`urnietfparamsscimschemasoracleidcsextension_oci_tags[0].defined_tags`,
			`schemas`,
			`urnietfparamsscimschemasoracleidcsextensionradius_app_app[0].group_membership_to_return`,
		}},
	}
	IdentityDomainsAppBasedOnTemplateRepresentation = map[string]interface{}{
		"value":         acctest.Representation{RepType: acctest.Required, Create: `CustomWebAppTemplateId`},
		"well_known_id": acctest.Representation{RepType: acctest.Optional, Create: `CustomWebAppTemplateId`},
	}
	IdentityDomainsAppAllowedTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsAppAttrRenderingMetadataRepresentation = map[string]interface{}{
		"name":       acctest.Representation{RepType: acctest.Required, Create: `name`},
		"datatype":   acctest.Representation{RepType: acctest.Optional, Create: `datatype`, Update: `datatype2`},
		"helptext":   acctest.Representation{RepType: acctest.Optional, Create: `helptext`, Update: `helptext2`},
		"label":      acctest.Representation{RepType: acctest.Optional, Create: `label`, Update: `label2`},
		"max_length": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_size":   acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"min_length": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"min_size":   acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"order":      acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"read_only":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"regexp":     acctest.Representation{RepType: acctest.Optional, Create: `regexp`, Update: `regexp2`},
		"required":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"section":    acctest.Representation{RepType: acctest.Optional, Create: `saml`, Update: `general`},
		"visible":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"widget":     acctest.Representation{RepType: acctest.Optional, Create: `inputtext`, Update: `checkbox`},
	}
	IdentityDomainsAppIdentityProvidersRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_identity_provider.test_identity_provider.id}`},
	}
	IdentityDomainsAppProtectableSecondaryAudiencesRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `secondaryAudiences`, Update: `secondaryAudiences2`},
	}
	IdentityDomainsAppSamlServiceProviderRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_app.test_saml_app.id}`},
	}
	IdentityDomainsAppScopesRepresentation = map[string]interface{}{
		"value":            acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
		"description":      acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"requires_consent": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsAppServiceParamsRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
	IdentityDomainsAppTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionOCITagsRepresentation = map[string]interface{}{
		"defined_tags":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionOCITagsDefinedTagsRepresentation},
		"freeform_tags": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionOCITagsFreeformTagsRepresentation},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionenterpriseAppAppRepresentation = map[string]interface{}{
		"allow_authz_decision_ttl": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"deny_authz_decision_ttl":  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionformFillAppAppRepresentation = map[string]interface{}{
		"configuration":                    acctest.Representation{RepType: acctest.Optional, Create: `configuration`, Update: `configuration2`},
		"form_cred_method":                 acctest.Representation{RepType: acctest.Optional, Create: `ADMIN_SETS_CREDENTIALS`, Update: `ADMIN_SETS_SHARED_CREDENTIALS`},
		"form_credential_sharing_group_id": acctest.Representation{RepType: acctest.Optional, Create: `formCredentialSharingGroupID`, Update: `formCredentialSharingGroupID2`},
		"form_fill_url_match":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionformFillAppAppFormFillUrlMatchRepresentation},
		"form_type":                        acctest.Representation{RepType: acctest.Optional, Create: `WebApplication`},
		"reveal_password_on_form":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"user_name_form_expression":        acctest.Representation{RepType: acctest.Optional, Create: `concat($user.firstname,\".\",$user.lastname)`},
		"user_name_form_template":          acctest.Representation{RepType: acctest.Optional, Create: `username`, Update: `email address`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionformFillAppTemplateAppTemplateRepresentation = map[string]interface{}{
		"configuration":                    acctest.Representation{RepType: acctest.Optional, Create: `configuration`, Update: `configuration2`},
		"form_cred_method":                 acctest.Representation{RepType: acctest.Optional, Create: `ADMIN_SETS_CREDENTIALS`, Update: `ADMIN_SETS_SHARED_CREDENTIALS`},
		"form_credential_sharing_group_id": acctest.Representation{RepType: acctest.Optional, Create: `formCredentialSharingGroupID`, Update: `formCredentialSharingGroupID2`},
		"form_fill_url_match":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionformFillAppTemplateAppTemplateFormFillUrlMatchRepresentation},
		"form_type":                        acctest.Representation{RepType: acctest.Optional, Create: `WebApplication`},
		"reveal_password_on_form":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"sync_from_template":               acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"user_name_form_expression":        acctest.Representation{RepType: acctest.Optional, Create: `concat($user.firstname,\".\",$user.lastname)`},
		"user_name_form_template":          acctest.Representation{RepType: acctest.Optional, Create: `username`, Update: `email address`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionkerberosRealmAppRepresentation = map[string]interface{}{
		"default_encryption_salt_type":    acctest.Representation{RepType: acctest.Optional, Create: `defaultEncryptionSaltType`, Update: `defaultEncryptionSaltType2`},
		"master_key":                      acctest.Representation{RepType: acctest.Optional, Create: `masterKey`, Update: `masterKey2`},
		"max_renewable_age":               acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_ticket_life":                 acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"realm_name":                      acctest.Representation{RepType: acctest.Optional, Create: `realmName`, Update: `realmName2`},
		"supported_encryption_salt_types": acctest.Representation{RepType: acctest.Optional, Create: []string{`supportedEncryptionSaltTypes`}, Update: []string{`supportedEncryptionSaltTypes2`}},
		"ticket_flags":                    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionmanagedappAppRepresentation = map[string]interface{}{
		"admin_consent_granted":                     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"bundle_configuration_properties":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionmanagedappAppBundleConfigurationPropertiesRepresentation},
		"bundle_pool_configuration":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionmanagedappAppBundlePoolConfigurationRepresentation},
		"connected":                                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"enable_auth_sync_new_user_notification":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"enable_sync":                               acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"enable_sync_summary_report_notification":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"flat_file_bundle_configuration_properties": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionmanagedappAppFlatFileBundleConfigurationPropertiesRepresentation},
		"is_authoritative":                          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"three_legged_oauth_credential":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionmanagedappAppThreeLeggedOAuthCredentialRepresentation},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionmulticloudServiceAppAppRepresentation = map[string]interface{}{
		"multicloud_service_type": acctest.Representation{RepType: acctest.Required, Create: `AWSCognito`},
		"multicloud_platform_url": acctest.Representation{RepType: acctest.Optional, Create: `multicloudPlatformUrl`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionopcServiceAppRepresentation = map[string]interface{}{
		"service_instance_identifier": acctest.Representation{RepType: acctest.Optional, Create: `serviceInstanceIdentifier`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionradiusAppAppRepresentation = map[string]interface{}{
		"client_ip":                          acctest.Representation{RepType: acctest.Required, Create: `clientIP`, Update: `clientIP2`},
		"include_group_in_response":          acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"port":                               acctest.Representation{RepType: acctest.Required, Create: `port`, Update: `port2`},
		"secret_key":                         acctest.Representation{RepType: acctest.Required, Create: `secretKey`, Update: `secretKey2`},
		"capture_client_ip":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"country_code_response_attribute_id": acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"end_user_ip_attribute":              acctest.Representation{RepType: acctest.Optional, Create: `31 Calling-Station-Id`, Update: `26 Vendor-Specific`},
		"group_membership_radius_attribute":  acctest.Representation{RepType: acctest.Optional, Create: `groupMembershipRadiusAttribute`, Update: `groupMembershipRadiusAttribute2`},
		"group_membership_to_return":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionradiusAppAppGroupMembershipToReturnRepresentation},
		"group_name_format":                  acctest.Representation{RepType: acctest.Optional, Create: `groupNameFormat`, Update: `groupNameFormat2`},
		"password_and_otp_together":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"radius_vendor_specific_id":          acctest.Representation{RepType: acctest.Optional, Create: `radiusVendorSpecificId`},
		"response_format":                    acctest.Representation{RepType: acctest.Optional, Create: `responseFormat`, Update: `responseFormat2`},
		"response_format_delimiter":          acctest.Representation{RepType: acctest.Optional, Create: `responseFormatDelimiter`, Update: `responseFormatDelimiter2`},
		"type_of_radius_app":                 acctest.Representation{RepType: acctest.Optional, Create: `Oracle Database`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionrequestableAppRepresentation = map[string]interface{}{
		"requestable": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionsamlServiceProviderAppRepresentation = map[string]interface{}{
		"assertion_consumer_url":            acctest.Representation{RepType: acctest.Optional, Create: `https://testurl.com`, Update: `https://testurl2.com`},
		"encrypt_assertion":                 acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"encryption_algorithm":              acctest.Representation{RepType: acctest.Optional, Create: `3DES`, Update: `AES-128`},
		"encryption_certificate":            acctest.Representation{RepType: acctest.Optional, Create: `MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw==`},
		"federation_protocol":               acctest.Representation{RepType: acctest.Optional, Create: `SAML2.0`, Update: `WS-Fed1.1`},
		"group_assertion_attributes":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionsamlServiceProviderAppGroupAssertionAttributesRepresentation},
		"hok_required":                      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"include_signing_cert_in_signature": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key_encryption_algorithm":          acctest.Representation{RepType: acctest.Optional, Create: `RSA-v1.5`, Update: `RSA-OAEP`},
		"logout_binding":                    acctest.Representation{RepType: acctest.Optional, Create: `Redirect`, Update: `Post`},
		"logout_enabled":                    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"logout_request_url":                acctest.Representation{RepType: acctest.Optional, Create: `https://testurl.com`, Update: `https://testurl2.com`},
		"logout_response_url":               acctest.Representation{RepType: acctest.Optional, Create: `https://testurl.com`, Update: `https://testurl2.com`},
		"metadata":                          acctest.Representation{RepType: acctest.Optional, Create: `<md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" xmlns:dsig=\"http://www.w3.org/2000/09/xmldsig#\" xmlns:enc=\"http://www.w3.org/2001/04/xmlenc#\" xmlns:mdattr=\"urn:oasis:names:tc:SAML:metadata:attribute\" xmlns:query=\"urn:oasis:names:tc:SAML:metadata:ext:query\" xmlns:saml=\"urn:oasis:names:tc:SAML:2.0:assertion\" xmlns:x500=\"urn:oasis:names:tc:SAML:2.0:profiles:attribute:X500\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" ID=\"id-zzU36agM7bKRB32xe6Ronm131S0-\" cacheDuration=\"P3633DT0H0M0S\" entityID=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com:443/fed\" validUntil=\"2031-06-16T06:38:32Z\"><dsig:Signature><dsig:SignedInfo><dsig:CanonicalizationMethod Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/><dsig:SignatureMethod Algorithm=\"http://www.w3.org/2001/04/xmldsig-more#rsa-sha256\"/><dsig:Reference URI=\"#id-zzU36agM7bKRB32xe6Ronm131S0-\"><dsig:Transforms><dsig:Transform Algorithm=\"http://www.w3.org/2000/09/xmldsig#enveloped-signature\"/><dsig:Transform Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/></dsig:Transforms><dsig:DigestMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#sha256\"/><dsig:DigestValue>NZnYsjLx3UbuL43iu3jo0mJUg/Rv9DTPNB5IQPRaD6g=</dsig:DigestValue></dsig:Reference></dsig:SignedInfo><dsig:SignatureValue>KRIgTD7//x/uT73veS0iGcWWw8uprjd+MtREu3vlbFTk0BNgkeSOYItx2LDQhnHP\nO0zsTmtOHlVIsDXQL3KysHwzYndIuMJtETqEC6NpMw3ZF108IK0eT+o/2xC9u13/\nGq10z/MagGvco1mM/RIzX5e2omGyZcKARiDoeNPwg2znmV0WcifntVqn4Y0rnWM7\no0M5HFHZQEgICdTJbC5d6DwLgfnI4ck505fHNRYLsRqj9IGLukKx9kocSG1xzCye\nHlffU4CDyEA7dptEUH59dZmY0Xy35/aepNc7W6IovWsJ2Otr+qDUp207ZCKuISF0\nMEX5hX5VJzVlHDwxkEcYCA==</dsig:SignatureValue><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==</dsig:X509Certificate></dsig:X509Data></dsig:KeyInfo></dsig:Signature><md:IDPSSODescriptor WantAuthnRequestsSigned=\"false\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/></md:IDPSSODescriptor><md:SPSSODescriptor AuthnRequestsSigned=\"true\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/sso\" index=\"1\" isDefault=\"true\"/></md:SPSSODescriptor></md:EntityDescriptor>`},
		"name_id_format":                    acctest.Representation{RepType: acctest.Optional, Create: `nameIdFormat`, Update: `nameIdFormat2`},
		"name_id_userstore_attribute":       acctest.Representation{RepType: acctest.Optional, Create: `emails.primary.value`, Update: `userName`},
		"partner_provider_id":               acctest.Representation{RepType: acctest.Optional, Create: `https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com:443/fed`},
		"partner_provider_pattern":          acctest.Representation{RepType: acctest.Optional, Create: `partnerProviderPattern`, Update: `partnerProviderPattern2`},
		"sign_response_or_assertion":        acctest.Representation{RepType: acctest.Optional, Create: `Assertion`, Update: `Response`},
		"signature_hash_algorithm":          acctest.Representation{RepType: acctest.Optional, Create: `SHA-1`, Update: `SHA-256`},
		"signing_certificate":               acctest.Representation{RepType: acctest.Optional, Create: `MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw==`},
		"succinct_id":                       acctest.Representation{RepType: acctest.Optional, Create: `succinctId`},
		"user_assertion_attributes":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionsamlServiceProviderAppUserAssertionAttributesRepresentation},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionwebTierPolicyAppRepresentation = map[string]interface{}{
		"web_tier_policy_json": acctest.Representation{RepType: acctest.Optional, Create: `{\"cloudgatePolicy\":{\"version\":\"2.6\",\"disableAuthorize\":false,\"webtierPolicy\":[{\"policyName\":\"test\",\"resourceFilters\":[]}]}}`, Update: `{\"cloudgatePolicy\":{\"version\":\"2.6\",\"disableAuthorize\":false,\"webtierPolicy\":[{\"policyName\":\"test2\",\"resourceFilters\":[]}]}}`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionOCITagsDefinedTagsRepresentation = map[string]interface{}{
		"key":       acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag.tag1.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag_namespace.tag-namespace1.name}`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionOCITagsFreeformTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `freeformKey`, Update: `freeformKey2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `freeformValue`, Update: `freeformValue2`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionformFillAppAppFormFillUrlMatchRepresentation = map[string]interface{}{
		"form_url":            acctest.Representation{RepType: acctest.Required, Create: `formUrl`, Update: `formUrl2`},
		"form_url_match_type": acctest.Representation{RepType: acctest.Optional, Create: `exact`, Update: `match`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionformFillAppTemplateAppTemplateFormFillUrlMatchRepresentation = map[string]interface{}{
		"form_url":            acctest.Representation{RepType: acctest.Required, Create: `formUrl`, Update: `formUrl2`},
		"form_url_match_type": acctest.Representation{RepType: acctest.Optional, Create: `exact`, Update: `match`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionmanagedappAppBundleConfigurationPropertiesRepresentation = map[string]interface{}{
		"icf_type":     acctest.Representation{RepType: acctest.Required, Create: `Long`},
		"name":         acctest.Representation{RepType: acctest.Required, Create: `name`},
		"required":     acctest.Representation{RepType: acctest.Required, Create: `false`},
		"confidential": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"help_message": acctest.Representation{RepType: acctest.Optional, Create: `helpMessage`, Update: `helpMessage2`},
		"order":        acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"value":        acctest.Representation{RepType: acctest.Optional, Create: []string{`value`}, Update: []string{`value2`}},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionmanagedappAppBundlePoolConfigurationRepresentation = map[string]interface{}{
		"max_idle":                       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_objects":                    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_wait":                       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"min_evictable_idle_time_millis": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"min_idle":                       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionmanagedappAppFlatFileBundleConfigurationPropertiesRepresentation = map[string]interface{}{
		"icf_type":     acctest.Representation{RepType: acctest.Required, Create: `Long`},
		"name":         acctest.Representation{RepType: acctest.Required, Create: `name`},
		"required":     acctest.Representation{RepType: acctest.Required, Create: `false`},
		"confidential": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"help_message": acctest.Representation{RepType: acctest.Optional, Create: `helpMessage`, Update: `helpMessage2`},
		"order":        acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"value":        acctest.Representation{RepType: acctest.Optional, Create: []string{`value`}, Update: []string{`value2`}},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionmanagedappAppThreeLeggedOAuthCredentialRepresentation = map[string]interface{}{
		"access_token":        acctest.Representation{RepType: acctest.Optional, Create: `accessToken`, Update: `accessToken2`},
		"access_token_expiry": acctest.Representation{RepType: acctest.Optional, Create: `2032-01-01T00:00:00Z`, Update: `2032-01-01T00:00:01Z`},
		"refresh_token":       acctest.Representation{RepType: acctest.Optional, Create: `refreshToken`, Update: `refreshToken2`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionradiusAppAppGroupMembershipToReturnRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_group.test_group.id}`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionsamlServiceProviderAppGroupAssertionAttributesRepresentation = map[string]interface{}{
		"name":       acctest.Representation{RepType: acctest.Required, Create: `groupName`, Update: `groupName2`},
		"condition":  acctest.Representation{RepType: acctest.Optional, Create: `Starts With`, Update: `Equals`},
		"format":     acctest.Representation{RepType: acctest.Optional, Create: `Basic`},
		"group_name": acctest.Representation{RepType: acctest.Optional, Create: `groupName`},
	}
	IdentityDomainsAppUrnietfparamsscimschemasoracleidcsextensionsamlServiceProviderAppUserAssertionAttributesRepresentation = map[string]interface{}{
		"name":                      acctest.Representation{RepType: acctest.Required, Create: `userName`, Update: `userName2`},
		"user_store_attribute_name": acctest.Representation{RepType: acctest.Required, Create: `emails.primary.value`, Update: `userName`},
		"format":                    acctest.Representation{RepType: acctest.Optional, Create: `Basic`},
	}

	IdentityDomainsAppResourceDependencies = DefinedTagsDependencies + TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_provider", "test_identity_provider", acctest.Required, acctest.Update, IdentityDomainsIdentityProviderRepresentation) + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Required, acctest.Create, IdentityDomainsGroupRepresentation) + SamlServiceProviderDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsAppResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsAppResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_app.test_app"
	datasourceName := "data.oci_identity_domains_apps.test_apps"
	singularDatasourceName := "data.oci_identity_domains_app.test_app"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsAppResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app", "test_app", acctest.Optional, acctest.Create, IdentityDomainsAppRepresentation), "identitydomains", "app", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsAppDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAppResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app", "test_app", acctest.Required, acctest.Create, IdentityDomainsAppRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "based_on_template.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "based_on_template.0.value", "CustomWebAppTemplateId"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAppResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAppResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app", "test_app", acctest.Optional, acctest.Create, IdentityDomainsAppRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_token_expiry", "10"),
				resource.TestCheckResourceAttr(resourceName, "active", "false"),
				resource.TestCheckResourceAttr(resourceName, "all_url_schemes_allowed", "false"),
				resource.TestCheckResourceAttr(resourceName, "allow_access_control", "false"),
				resource.TestCheckResourceAttr(resourceName, "allow_offline", "false"),
				resource.TestCheckResourceAttr(resourceName, "allowed_grants.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "allowed_operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "allowed_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "allowed_tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "allowed_tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "app_icon", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII="),
				resource.TestCheckResourceAttr(resourceName, "app_thumbnail", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII="),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.datatype", "datatype"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.helptext", "helptext"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.max_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.max_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.min_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.min_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.order", "10"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.read_only", "false"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.regexp", "regexp"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.required", "false"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.section", "saml"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.visible", "false"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.widget", "inputtext"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "audience", "audience"),
				resource.TestCheckResourceAttr(resourceName, "based_on_template.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "based_on_template.0.value", "CustomWebAppTemplateId"),
				resource.TestCheckResourceAttrSet(resourceName, "based_on_template.0.well_known_id"),
				resource.TestCheckResourceAttr(resourceName, "bypass_consent", "false"),
				resource.TestCheckResourceAttr(resourceName, "client_ip_checking", "anywhere"),
				resource.TestCheckResourceAttr(resourceName, "client_type", "confidential"),
				resource.TestCheckResourceAttr(resourceName, "contact_email_address", "contact@email.com"),
				resource.TestCheckResourceAttr(resourceName, "delegated_service_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "disable_kmsi_token_authentication", "false"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "error_page_url", "https://testurl.com"),
				resource.TestCheckResourceAttr(resourceName, "home_page_url", "https://testurl.com"),
				resource.TestCheckResourceAttr(resourceName, "icon", "icon"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "id_token_enc_algo", "A128CBC-HS256"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "identity_providers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "identity_providers.0.value"),
				resource.TestCheckResourceAttr(resourceName, "is_alias_app", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_enterprise_app", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_form_fill", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_kerberos_realm", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_login_target", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_mobile_target", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_multicloud_service_app", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_oauth_client", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_oauth_resource", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_obligation_capable", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_radius_app", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_saml_service_provider", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_unmanaged_app", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_web_tier_policy", "false"),
				resource.TestCheckResourceAttr(resourceName, "landing_page_url", "https://testurl.com"),
				resource.TestCheckResourceAttr(resourceName, "linking_callback_url", "https://testurl.com"),
				resource.TestCheckResourceAttr(resourceName, "login_mechanism", "OIDC"),
				resource.TestCheckResourceAttr(resourceName, "login_page_url", "https://testurl.com"),
				resource.TestCheckResourceAttr(resourceName, "logout_page_url", "https://testurl.com"),
				resource.TestCheckResourceAttr(resourceName, "logout_uri", "logoutUri"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "post_logout_redirect_uris.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "privacy_policy_url", "https://testurl.com"),
				resource.TestCheckResourceAttr(resourceName, "product_logo_url", "https://testurl.com"),
				resource.TestCheckResourceAttr(resourceName, "product_name", "productName"),
				resource.TestCheckResourceAttr(resourceName, "protectable_secondary_audiences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "protectable_secondary_audiences.0.value", "secondaryAudiences"),
				resource.TestCheckResourceAttr(resourceName, "redirect_uris.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "refresh_token_expiry", "10"),
				resource.TestCheckResourceAttr(resourceName, "saml_service_provider.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "saml_service_provider.0.value"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(resourceName, "scopes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.requires_consent", "false"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "secondary_audiences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_params.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_params.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_params.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "service_type_urn", "serviceTypeURN"),
				resource.TestCheckResourceAttr(resourceName, "service_type_version", "serviceTypeVersion"),
				resource.TestCheckResourceAttr(resourceName, "show_in_my_apps", "false"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "terms_of_service_url", "https://testurl.com"),
				resource.TestCheckResourceAttr(resourceName, "trust_scope", "Explicit"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.key", "freeformKey"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.value", "freeformValue"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionenterprise_app_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionenterprise_app_app.0.allow_authz_decision_ttl", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionenterprise_app_app.0.deny_authz_decision_ttl", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.configuration", "configuration"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_cred_method", "ADMIN_SETS_CREDENTIALS"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_credential_sharing_group_id", "formCredentialSharingGroupID"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_fill_url_match.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_fill_url_match.0.form_url", "formUrl"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_fill_url_match.0.form_url_match_type", "exact"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_type", "WebApplication"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.reveal_password_on_form", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.user_name_form_expression", "concat($user.firstname,\".\",$user.lastname)"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.user_name_form_template", "username"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.configuration", "configuration"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_cred_method", "ADMIN_SETS_CREDENTIALS"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_credential_sharing_group_id", "formCredentialSharingGroupID"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_fill_url_match.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_fill_url_match.0.form_url", "formUrl"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_fill_url_match.0.form_url_match_type", "exact"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_type", "WebApplication"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.reveal_password_on_form", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.sync_from_template", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.user_name_form_expression", "concat($user.firstname,\".\",$user.lastname)"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.user_name_form_template", "username"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.default_encryption_salt_type", "defaultEncryptionSaltType"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.master_key", "masterKey"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.max_renewable_age", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.max_ticket_life", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.realm_name", "realmName"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.supported_encryption_salt_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.ticket_flags", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.admin_consent_granted", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.confidential", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.help_message", "helpMessage"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.icf_type", "Long"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.order", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.required", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.value.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.max_idle", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.max_objects", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.max_wait", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.min_evictable_idle_time_millis", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.min_idle", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.connected", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.enable_auth_sync_new_user_notification", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.enable_sync", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.enable_sync_summary_report_notification", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.confidential", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.help_message", "helpMessage"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.icf_type", "Long"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.order", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.required", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.value.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.is_authoritative", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.three_legged_oauth_credential.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.three_legged_oauth_credential.0.access_token", "accessToken"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.three_legged_oauth_credential.0.access_token_expiry", "2032-01-01T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.three_legged_oauth_credential.0.refresh_token", "refreshToken"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app.0.multicloud_platform_url", "multicloudPlatformUrl"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app.0.multicloud_service_type", "AWSCognito"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionopc_service_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionopc_service_app.0.service_instance_identifier", "serviceInstanceIdentifier"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.capture_client_ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.client_ip", "clientIP"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.country_code_response_attribute_id"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.end_user_ip_attribute", "31 Calling-Station-Id"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.group_membership_radius_attribute", "groupMembershipRadiusAttribute"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.group_name_format", "groupNameFormat"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.include_group_in_response", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.password_and_otp_together", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.port", "port"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.radius_vendor_specific_id"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.response_format", "responseFormat"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.response_format_delimiter", "responseFormatDelimiter"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.secret_key", "secretKey"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.type_of_radius_app", "Oracle Database"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionrequestable_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionrequestable_app.0.requestable", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.assertion_consumer_url", "https://testurl.com"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.encrypt_assertion", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.encryption_algorithm", "3DES"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.encryption_certificate", "MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw=="),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.federation_protocol", "SAML2.0"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.0.condition", "Starts With"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.0.format", "Basic"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.0.group_name"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.0.name", "groupName"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.hok_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.include_signing_cert_in_signature", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.key_encryption_algorithm", "RSA-v1.5"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.logout_binding", "Redirect"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.logout_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.logout_request_url", "https://testurl.com"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.logout_response_url", "https://testurl.com"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.metadata", "<md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" xmlns:dsig=\"http://www.w3.org/2000/09/xmldsig#\" xmlns:enc=\"http://www.w3.org/2001/04/xmlenc#\" xmlns:mdattr=\"urn:oasis:names:tc:SAML:metadata:attribute\" xmlns:query=\"urn:oasis:names:tc:SAML:metadata:ext:query\" xmlns:saml=\"urn:oasis:names:tc:SAML:2.0:assertion\" xmlns:x500=\"urn:oasis:names:tc:SAML:2.0:profiles:attribute:X500\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" ID=\"id-zzU36agM7bKRB32xe6Ronm131S0-\" cacheDuration=\"P3633DT0H0M0S\" entityID=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com:443/fed\" validUntil=\"2031-06-16T06:38:32Z\"><dsig:Signature><dsig:SignedInfo><dsig:CanonicalizationMethod Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/><dsig:SignatureMethod Algorithm=\"http://www.w3.org/2001/04/xmldsig-more#rsa-sha256\"/><dsig:Reference URI=\"#id-zzU36agM7bKRB32xe6Ronm131S0-\"><dsig:Transforms><dsig:Transform Algorithm=\"http://www.w3.org/2000/09/xmldsig#enveloped-signature\"/><dsig:Transform Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/></dsig:Transforms><dsig:DigestMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#sha256\"/><dsig:DigestValue>NZnYsjLx3UbuL43iu3jo0mJUg/Rv9DTPNB5IQPRaD6g=</dsig:DigestValue></dsig:Reference></dsig:SignedInfo><dsig:SignatureValue>KRIgTD7//x/uT73veS0iGcWWw8uprjd+MtREu3vlbFTk0BNgkeSOYItx2LDQhnHP\nO0zsTmtOHlVIsDXQL3KysHwzYndIuMJtETqEC6NpMw3ZF108IK0eT+o/2xC9u13/\nGq10z/MagGvco1mM/RIzX5e2omGyZcKARiDoeNPwg2znmV0WcifntVqn4Y0rnWM7\no0M5HFHZQEgICdTJbC5d6DwLgfnI4ck505fHNRYLsRqj9IGLukKx9kocSG1xzCye\nHlffU4CDyEA7dptEUH59dZmY0Xy35/aepNc7W6IovWsJ2Otr+qDUp207ZCKuISF0\nMEX5hX5VJzVlHDwxkEcYCA==</dsig:SignatureValue><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==</dsig:X509Certificate></dsig:X509Data></dsig:KeyInfo></dsig:Signature><md:IDPSSODescriptor WantAuthnRequestsSigned=\"false\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/></md:IDPSSODescriptor><md:SPSSODescriptor AuthnRequestsSigned=\"true\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/sso\" index=\"1\" isDefault=\"true\"/></md:SPSSODescriptor></md:EntityDescriptor>"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.name_id_format", "nameIdFormat"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.name_id_userstore_attribute", "emails.primary.value"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.partner_provider_id"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.partner_provider_pattern", "partnerProviderPattern"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.sign_response_or_assertion", "Assertion"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.signature_hash_algorithm", "SHA-1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.signing_certificate", "MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw=="),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.succinct_id"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.user_assertion_attributes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.user_assertion_attributes.0.format", "Basic"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.user_assertion_attributes.0.name", "userName"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.user_assertion_attributes.0.user_store_attribute_name", "emails.primary.value"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app.0.web_tier_policy_json", "{\"cloudgatePolicy\":{\"version\":\"2.6\",\"disableAuthorize\":false,\"webtierPolicy\":[{\"policyName\":\"test\",\"resourceFilters\":[]}]}}"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "apps", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsAppResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app", "test_app", acctest.Optional, acctest.Update, IdentityDomainsAppRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_token_expiry", "11"),
				resource.TestCheckResourceAttr(resourceName, "active", "false"),
				resource.TestCheckResourceAttr(resourceName, "all_url_schemes_allowed", "true"),
				resource.TestCheckResourceAttr(resourceName, "allow_access_control", "true"),
				resource.TestCheckResourceAttr(resourceName, "allow_offline", "true"),
				resource.TestCheckResourceAttr(resourceName, "allowed_grants.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "allowed_operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "allowed_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "allowed_tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "allowed_tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "app_icon", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO+ip1sAAAAASUVORK5CYII="),
				resource.TestCheckResourceAttr(resourceName, "app_thumbnail", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO+ip1sAAAAASUVORK5CYII="),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.datatype", "datatype2"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.helptext", "helptext2"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.label", "label2"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.max_length", "11"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.max_size", "11"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.min_length", "11"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.min_size", "11"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.order", "11"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.read_only", "true"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.regexp", "regexp2"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.required", "true"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.section", "general"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.visible", "true"),
				resource.TestCheckResourceAttr(resourceName, "attr_rendering_metadata.0.widget", "checkbox"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "audience", "audience2"),
				resource.TestCheckResourceAttr(resourceName, "based_on_template.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "based_on_template.0.value", "CustomWebAppTemplateId"),
				resource.TestCheckResourceAttrSet(resourceName, "based_on_template.0.well_known_id"),
				resource.TestCheckResourceAttr(resourceName, "bypass_consent", "true"),
				resource.TestCheckResourceAttr(resourceName, "client_ip_checking", "whitelisted"),
				resource.TestCheckResourceAttr(resourceName, "client_type", "trusted"),
				resource.TestCheckResourceAttr(resourceName, "contact_email_address", "contact2@email.com"),
				resource.TestCheckResourceAttr(resourceName, "delegated_service_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "disable_kmsi_token_authentication", "true"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "error_page_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(resourceName, "home_page_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(resourceName, "icon", "icon2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "id_token_enc_algo", "A192CBC-HS384"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "identity_providers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "identity_providers.0.value"),
				resource.TestCheckResourceAttr(resourceName, "is_alias_app", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_enterprise_app", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_form_fill", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_kerberos_realm", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_login_target", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_mobile_target", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_multicloud_service_app", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_oauth_client", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_oauth_resource", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_obligation_capable", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_radius_app", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_saml_service_provider", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_unmanaged_app", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_web_tier_policy", "true"),
				resource.TestCheckResourceAttr(resourceName, "landing_page_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(resourceName, "linking_callback_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(resourceName, "login_mechanism", "SAML"),
				resource.TestCheckResourceAttr(resourceName, "login_page_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(resourceName, "logout_page_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(resourceName, "logout_uri", "logoutUri2"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "post_logout_redirect_uris.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "privacy_policy_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(resourceName, "product_logo_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(resourceName, "product_name", "productName2"),
				resource.TestCheckResourceAttr(resourceName, "protectable_secondary_audiences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "protectable_secondary_audiences.0.value", "secondaryAudiences2"),
				resource.TestCheckResourceAttr(resourceName, "redirect_uris.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "refresh_token_expiry", "11"),
				resource.TestCheckResourceAttr(resourceName, "saml_service_provider.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "saml_service_provider.0.value"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(resourceName, "scopes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.requires_consent", "true"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "secondary_audiences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_params.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_params.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "service_params.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "service_type_urn", "serviceTypeURN2"),
				resource.TestCheckResourceAttr(resourceName, "service_type_version", "serviceTypeVersion2"),
				resource.TestCheckResourceAttr(resourceName, "show_in_my_apps", "true"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "terms_of_service_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(resourceName, "trust_scope", "Account"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.key", "freeformKey2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.value", "freeformValue2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionenterprise_app_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionenterprise_app_app.0.allow_authz_decision_ttl", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionenterprise_app_app.0.deny_authz_decision_ttl", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.configuration", "configuration2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_cred_method", "ADMIN_SETS_SHARED_CREDENTIALS"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_credential_sharing_group_id", "formCredentialSharingGroupID2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_fill_url_match.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_fill_url_match.0.form_url", "formUrl2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_fill_url_match.0.form_url_match_type", "match"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_type", "WebApplication"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.reveal_password_on_form", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.user_name_form_expression", "concat($user.firstname,\".\",$user.lastname)"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.user_name_form_template", "email address"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.configuration", "configuration2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_cred_method", "ADMIN_SETS_SHARED_CREDENTIALS"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_credential_sharing_group_id", "formCredentialSharingGroupID2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_fill_url_match.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_fill_url_match.0.form_url", "formUrl2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_fill_url_match.0.form_url_match_type", "match"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_type", "WebApplication"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.reveal_password_on_form", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.sync_from_template", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.user_name_form_expression", "concat($user.firstname,\".\",$user.lastname)"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.user_name_form_template", "email address"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.default_encryption_salt_type", "defaultEncryptionSaltType2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.master_key", "masterKey2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.max_renewable_age", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.max_ticket_life", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.realm_name", "realmName2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.supported_encryption_salt_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app.0.ticket_flags", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.admin_consent_granted", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.confidential", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.help_message", "helpMessage2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.icf_type", "Long"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.order", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.required", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.value.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.max_idle", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.max_objects", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.max_wait", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.min_evictable_idle_time_millis", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.min_idle", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.connected", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.enable_auth_sync_new_user_notification", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.enable_sync", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.enable_sync_summary_report_notification", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.confidential", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.help_message", "helpMessage2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.icf_type", "Long"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.order", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.required", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.value.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.is_authoritative", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.three_legged_oauth_credential.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.three_legged_oauth_credential.0.access_token", "accessToken2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.three_legged_oauth_credential.0.access_token_expiry", "2032-01-01T00:00:01Z"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.three_legged_oauth_credential.0.refresh_token", "refreshToken2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app.0.multicloud_platform_url", "multicloudPlatformUrl"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app.0.multicloud_service_type", "AWSCognito"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionopc_service_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionopc_service_app.0.service_instance_identifier", "serviceInstanceIdentifier"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.capture_client_ip", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.client_ip", "clientIP2"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.country_code_response_attribute_id"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.end_user_ip_attribute", "26 Vendor-Specific"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.group_membership_radius_attribute", "groupMembershipRadiusAttribute2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.group_name_format", "groupNameFormat2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.include_group_in_response", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.password_and_otp_together", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.port", "port2"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.radius_vendor_specific_id"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.response_format", "responseFormat2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.response_format_delimiter", "responseFormatDelimiter2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.secret_key", "secretKey2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.type_of_radius_app", "Oracle Database"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionrequestable_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionrequestable_app.0.requestable", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.assertion_consumer_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.encrypt_assertion", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.encryption_algorithm", "AES-128"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.encryption_certificate", "MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw=="),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.federation_protocol", "WS-Fed1.1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.0.condition", "Equals"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.0.format", "Basic"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.0.group_name"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.0.name", "groupName2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.hok_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.include_signing_cert_in_signature", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.key_encryption_algorithm", "RSA-OAEP"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.logout_binding", "Post"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.logout_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.logout_request_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.logout_response_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.metadata", "<md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" xmlns:dsig=\"http://www.w3.org/2000/09/xmldsig#\" xmlns:enc=\"http://www.w3.org/2001/04/xmlenc#\" xmlns:mdattr=\"urn:oasis:names:tc:SAML:metadata:attribute\" xmlns:query=\"urn:oasis:names:tc:SAML:metadata:ext:query\" xmlns:saml=\"urn:oasis:names:tc:SAML:2.0:assertion\" xmlns:x500=\"urn:oasis:names:tc:SAML:2.0:profiles:attribute:X500\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" ID=\"id-zzU36agM7bKRB32xe6Ronm131S0-\" cacheDuration=\"P3633DT0H0M0S\" entityID=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com:443/fed\" validUntil=\"2031-06-16T06:38:32Z\"><dsig:Signature><dsig:SignedInfo><dsig:CanonicalizationMethod Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/><dsig:SignatureMethod Algorithm=\"http://www.w3.org/2001/04/xmldsig-more#rsa-sha256\"/><dsig:Reference URI=\"#id-zzU36agM7bKRB32xe6Ronm131S0-\"><dsig:Transforms><dsig:Transform Algorithm=\"http://www.w3.org/2000/09/xmldsig#enveloped-signature\"/><dsig:Transform Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/></dsig:Transforms><dsig:DigestMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#sha256\"/><dsig:DigestValue>NZnYsjLx3UbuL43iu3jo0mJUg/Rv9DTPNB5IQPRaD6g=</dsig:DigestValue></dsig:Reference></dsig:SignedInfo><dsig:SignatureValue>KRIgTD7//x/uT73veS0iGcWWw8uprjd+MtREu3vlbFTk0BNgkeSOYItx2LDQhnHP\nO0zsTmtOHlVIsDXQL3KysHwzYndIuMJtETqEC6NpMw3ZF108IK0eT+o/2xC9u13/\nGq10z/MagGvco1mM/RIzX5e2omGyZcKARiDoeNPwg2znmV0WcifntVqn4Y0rnWM7\no0M5HFHZQEgICdTJbC5d6DwLgfnI4ck505fHNRYLsRqj9IGLukKx9kocSG1xzCye\nHlffU4CDyEA7dptEUH59dZmY0Xy35/aepNc7W6IovWsJ2Otr+qDUp207ZCKuISF0\nMEX5hX5VJzVlHDwxkEcYCA==</dsig:SignatureValue><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==</dsig:X509Certificate></dsig:X509Data></dsig:KeyInfo></dsig:Signature><md:IDPSSODescriptor WantAuthnRequestsSigned=\"false\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/></md:IDPSSODescriptor><md:SPSSODescriptor AuthnRequestsSigned=\"true\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/sso\" index=\"1\" isDefault=\"true\"/></md:SPSSODescriptor></md:EntityDescriptor>"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.name_id_format", "nameIdFormat2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.name_id_userstore_attribute", "userName"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.partner_provider_id"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.partner_provider_pattern", "partnerProviderPattern2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.sign_response_or_assertion", "Response"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.signature_hash_algorithm", "SHA-256"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.signing_certificate", "MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw=="),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.succinct_id"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.user_assertion_attributes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.user_assertion_attributes.0.format", "Basic"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.user_assertion_attributes.0.name", "userName2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.user_assertion_attributes.0.user_store_attribute_name", "userName"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app.0.web_tier_policy_json", "{\"cloudgatePolicy\":{\"version\":\"2.6\",\"disableAuthorize\":false,\"webtierPolicy\":[{\"policyName\":\"test2\",\"resourceFilters\":[]}]}}"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_apps", "test_apps", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsAppDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsAppResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app", "test_app", acctest.Optional, acctest.Update, IdentityDomainsAppRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "app_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "apps.#", regexp.MustCompile("[1-9]+")),
				resource.TestMatchResourceAttr(datasourceName, "apps.0.schemas.#", regexp.MustCompile("[1-9]+")),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_app", "test_app", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsAppSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsAppResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "app_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttr(singularDatasourceName, "access_token_expiry", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "active", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "all_url_schemes_allowed", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allow_access_control", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allow_offline", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_grants.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_operations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_tags.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_tags.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.datatype", "datatype2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.helptext", "helptext2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.label", "label2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.max_length", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.max_size", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.min_length", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.min_size", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.order", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.read_only", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.regexp", "regexp2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.required", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.section", "general"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.visible", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attr_rendering_metadata.0.widget", "checkbox"),
				resource.TestCheckResourceAttr(singularDatasourceName, "audience", "audience2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "based_on_template.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "based_on_template.0.value", "CustomWebAppTemplateId"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bypass_consent", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_ip_checking", "whitelisted"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_type", "trusted"),
				resource.TestCheckResourceAttr(singularDatasourceName, "contact_email_address", "contact2@email.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "delegated_service_names.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "disable_kmsi_token_authentication", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "error_page_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "home_page_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "icon", "icon2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "id_token_enc_algo", "A192CBC-HS384"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_alias_app", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enterprise_app", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_form_fill", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_kerberos_realm", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_login_target", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_mobile_target", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_multicloud_service_app", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_oauth_client", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_oauth_resource", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_obligation_capable", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_radius_app", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_saml_service_provider", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_unmanaged_app", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_web_tier_policy", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "landing_page_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "linking_callback_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "login_mechanism", "SAML"),
				resource.TestCheckResourceAttr(singularDatasourceName, "login_page_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logout_page_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logout_uri", "logoutUri2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "post_logout_redirect_uris.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "privacy_policy_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "product_logo_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "product_name", "productName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protectable_secondary_audiences.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protectable_secondary_audiences.0.value", "secondaryAudiences2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "redirect_uris.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "refresh_token_expiry", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "saml_service_provider.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "saml_service_provider.0.value"),
				resource.TestMatchResourceAttr(singularDatasourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(singularDatasourceName, "scopes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scopes.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scopes.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scopes.0.requires_consent", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scopes.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "secondary_audiences.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_params.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_params.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_params.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_type_urn", "serviceTypeURN2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_type_version", "serviceTypeVersion2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "show_in_my_apps", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "terms_of_service_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trust_scope", "Account"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.key", "freeformKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.value", "freeformValue2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionenterprise_app_app.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionenterprise_app_app.0.allow_authz_decision_ttl", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionenterprise_app_app.0.deny_authz_decision_ttl", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.configuration", "configuration2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_cred_method", "ADMIN_SETS_SHARED_CREDENTIALS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_credential_sharing_group_id", "formCredentialSharingGroupID2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_fill_url_match.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_fill_url_match.0.form_url", "formUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_fill_url_match.0.form_url_match_type", "match"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.form_type", "WebApplication"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.reveal_password_on_form", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.user_name_form_expression", "concat($user.firstname,\".\",$user.lastname)"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app.0.user_name_form_template", "email address"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.configuration", "configuration2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_cred_method", "ADMIN_SETS_SHARED_CREDENTIALS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_credential_sharing_group_id", "formCredentialSharingGroupID2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_fill_url_match.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_fill_url_match.0.form_url", "formUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_fill_url_match.0.form_url_match_type", "match"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.form_type", "WebApplication"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.reveal_password_on_form", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.sync_from_template", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.user_name_form_expression", "concat($user.firstname,\".\",$user.lastname)"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template.0.user_name_form_template", "email address"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.admin_consent_granted", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.confidential", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.help_message", "helpMessage2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.icf_type", "Long"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.order", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.required", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_configuration_properties.0.value.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.max_idle", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.max_objects", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.max_wait", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.min_evictable_idle_time_millis", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.bundle_pool_configuration.0.min_idle", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.connected", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.enable_auth_sync_new_user_notification", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.enable_sync", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.enable_sync_summary_report_notification", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.confidential", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.help_message", "helpMessage2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.icf_type", "Long"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.order", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.required", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.flat_file_bundle_configuration_properties.0.value.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.is_authoritative", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.three_legged_oauth_credential.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.three_legged_oauth_credential.0.access_token", "accessToken2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.three_legged_oauth_credential.0.access_token_expiry", "2032-01-01T00:00:01Z"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app.0.three_legged_oauth_credential.0.refresh_token", "refreshToken2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionopc_service_app.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionopc_service_app.0.service_instance_identifier", "serviceInstanceIdentifier"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.capture_client_ip", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.client_ip", "clientIP2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.country_code_response_attribute_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.end_user_ip_attribute", "26 Vendor-Specific"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.group_membership_radius_attribute", "groupMembershipRadiusAttribute2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.group_name_format", "groupNameFormat2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.include_group_in_response", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.password_and_otp_together", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.port", "port2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.radius_vendor_specific_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.response_format", "responseFormat2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.response_format_delimiter", "responseFormatDelimiter2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.secret_key", "secretKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionradius_app_app.0.type_of_radius_app", "Oracle Database"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.assertion_consumer_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.encrypt_assertion", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.encryption_algorithm", "AES-128"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.encryption_certificate", "MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw=="),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.federation_protocol", "WS-Fed1.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.0.condition", "Equals"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.0.format", "Basic"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.group_assertion_attributes.0.name", "groupName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.hok_required", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.include_signing_cert_in_signature", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.key_encryption_algorithm", "RSA-OAEP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.logout_binding", "Post"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.logout_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.logout_request_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.logout_response_url", "https://testurl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.metadata", "<md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" xmlns:dsig=\"http://www.w3.org/2000/09/xmldsig#\" xmlns:enc=\"http://www.w3.org/2001/04/xmlenc#\" xmlns:mdattr=\"urn:oasis:names:tc:SAML:metadata:attribute\" xmlns:query=\"urn:oasis:names:tc:SAML:metadata:ext:query\" xmlns:saml=\"urn:oasis:names:tc:SAML:2.0:assertion\" xmlns:x500=\"urn:oasis:names:tc:SAML:2.0:profiles:attribute:X500\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" ID=\"id-zzU36agM7bKRB32xe6Ronm131S0-\" cacheDuration=\"P3633DT0H0M0S\" entityID=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com:443/fed\" validUntil=\"2031-06-16T06:38:32Z\"><dsig:Signature><dsig:SignedInfo><dsig:CanonicalizationMethod Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/><dsig:SignatureMethod Algorithm=\"http://www.w3.org/2001/04/xmldsig-more#rsa-sha256\"/><dsig:Reference URI=\"#id-zzU36agM7bKRB32xe6Ronm131S0-\"><dsig:Transforms><dsig:Transform Algorithm=\"http://www.w3.org/2000/09/xmldsig#enveloped-signature\"/><dsig:Transform Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/></dsig:Transforms><dsig:DigestMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#sha256\"/><dsig:DigestValue>NZnYsjLx3UbuL43iu3jo0mJUg/Rv9DTPNB5IQPRaD6g=</dsig:DigestValue></dsig:Reference></dsig:SignedInfo><dsig:SignatureValue>KRIgTD7//x/uT73veS0iGcWWw8uprjd+MtREu3vlbFTk0BNgkeSOYItx2LDQhnHP\nO0zsTmtOHlVIsDXQL3KysHwzYndIuMJtETqEC6NpMw3ZF108IK0eT+o/2xC9u13/\nGq10z/MagGvco1mM/RIzX5e2omGyZcKARiDoeNPwg2znmV0WcifntVqn4Y0rnWM7\no0M5HFHZQEgICdTJbC5d6DwLgfnI4ck505fHNRYLsRqj9IGLukKx9kocSG1xzCye\nHlffU4CDyEA7dptEUH59dZmY0Xy35/aepNc7W6IovWsJ2Otr+qDUp207ZCKuISF0\nMEX5hX5VJzVlHDwxkEcYCA==</dsig:SignatureValue><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==</dsig:X509Certificate></dsig:X509Data></dsig:KeyInfo></dsig:Signature><md:IDPSSODescriptor WantAuthnRequestsSigned=\"false\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/></md:IDPSSODescriptor><md:SPSSODescriptor AuthnRequestsSigned=\"true\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/sso\" index=\"1\" isDefault=\"true\"/></md:SPSSODescriptor></md:EntityDescriptor>"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.name_id_format", "nameIdFormat2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.name_id_userstore_attribute", "userName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.sign_response_or_assertion", "Response"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.signature_hash_algorithm", "SHA-256"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.signing_certificate", "MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw=="),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.user_assertion_attributes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.user_assertion_attributes.0.format", "Basic"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.user_assertion_attributes.0.name", "userName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app.0.user_assertion_attributes.0.user_store_attribute_name", "userName"),
			),
		},
		// reset to IdentityDomainsAppRequiredOnlyResource before verify import
		{
			Config: config + IdentityDomainsAppRequiredOnlyResource,
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsAppRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_app", "apps"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"app_icon",
				"app_thumbnail",
				"hashed_client_secret",
			},
			ResourceName: resourceName,
		},
		// delete the test_app resource before clean up the dependencies
		{
			Config: config + IdentityDomainsAppResourceDependencies,
		},
	})
}

func testAccCheckIdentityDomainsAppDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_app" {
			noResourceFound = false
			request := oci_identity_domains.GetAppRequest{}

			tmp := rs.Primary.ID
			request.AppId = &tmp

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetApp(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsApp") {
		resource.AddTestSweepers("IdentityDomainsApp", &resource.Sweeper{
			Name:         "IdentityDomainsApp",
			Dependencies: acctest.DependencyGraph["app"],
			F:            sweepIdentityDomainsAppResource,
		})
	}
}

func sweepIdentityDomainsAppResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	appIds, err := getIdentityDomainsAppIds(compartment)
	if err != nil {
		return err
	}
	for _, appId := range appIds {
		if ok := acctest.SweeperDefaultResourceId[appId]; !ok {
			deleteAppRequest := oci_identity_domains.DeleteAppRequest{}

			deleteAppRequest.AppId = &appId

			deleteAppRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteApp(context.Background(), deleteAppRequest)
			if error != nil {
				fmt.Printf("Error deleting App %s %s, It is possible that the resource is already deleted. Please verify manually \n", appId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsAppIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AppId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listAppsRequest := oci_identity_domains.ListAppsRequest{}
	listAppsResponse, err := identityDomainsClient.ListApps(context.Background(), listAppsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting App list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, app := range listAppsResponse.Resources {
		id := *app.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AppId", id)
	}
	return resourceIds, nil
}
