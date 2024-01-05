// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsAppDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["app_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["attribute_sets"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	fieldMap["attributes"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["authorization"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["idcs_endpoint"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["resource_type_schema_version"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsAppResource(), fieldMap, readSingularIdentityDomainsApp)
}

func readSingularIdentityDomainsApp(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsAppDataSourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

type IdentityDomainsAppDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetAppResponse
}

func (s *IdentityDomainsAppDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsAppDataSourceCrud) Get() error {
	request := oci_identity_domains.GetAppRequest{}

	if appId, ok := s.D.GetOkExists("app_id"); ok {
		tmp := appId.(string)
		request.AppId = &tmp
	}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetApp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsAppDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AccessTokenExpiry != nil {
		s.D.Set("access_token_expiry", *s.Res.AccessTokenExpiry)
	}

	accounts := []interface{}{}
	for _, item := range s.Res.Accounts {
		accounts = append(accounts, AppAccountsToMap(item))
	}
	s.D.Set("accounts", accounts)

	if s.Res.Active != nil {
		s.D.Set("active", *s.Res.Active)
	}

	adminRoles := []interface{}{}
	for _, item := range s.Res.AdminRoles {
		adminRoles = append(adminRoles, AppAdminRolesToMap(item))
	}
	s.D.Set("admin_roles", adminRoles)

	aliasApps := []interface{}{}
	for _, item := range s.Res.AliasApps {
		aliasApps = append(aliasApps, AppAliasAppsToMap(item))
	}
	s.D.Set("alias_apps", aliasApps)

	if s.Res.AllUrlSchemesAllowed != nil {
		s.D.Set("all_url_schemes_allowed", *s.Res.AllUrlSchemesAllowed)
	}

	if s.Res.AllowAccessControl != nil {
		s.D.Set("allow_access_control", *s.Res.AllowAccessControl)
	}

	if s.Res.AllowOffline != nil {
		s.D.Set("allow_offline", *s.Res.AllowOffline)
	}

	s.D.Set("allowed_grants", s.Res.AllowedGrants)

	s.D.Set("allowed_operations", s.Res.AllowedOperations)

	allowedScopes := []interface{}{}
	for _, item := range s.Res.AllowedScopes {
		allowedScopes = append(allowedScopes, AppAllowedScopesToMap(item))
	}
	s.D.Set("allowed_scopes", allowedScopes)

	allowedTags := []interface{}{}
	for _, item := range s.Res.AllowedTags {
		allowedTags = append(allowedTags, AppAllowedTagsToMap(item))
	}
	s.D.Set("allowed_tags", allowedTags)

	if s.Res.AppIcon != nil {
		s.D.Set("app_icon", *s.Res.AppIcon)
	}

	if s.Res.AppSignonPolicy != nil {
		s.D.Set("app_signon_policy", []interface{}{AppAppSignonPolicyToMap(s.Res.AppSignonPolicy)})
	} else {
		s.D.Set("app_signon_policy", nil)
	}

	if s.Res.AppThumbnail != nil {
		s.D.Set("app_thumbnail", *s.Res.AppThumbnail)
	}

	appsNetworkPerimeters := []interface{}{}
	for _, item := range s.Res.AppsNetworkPerimeters {
		appsNetworkPerimeters = append(appsNetworkPerimeters, AppAppsNetworkPerimetersToMap(item))
	}
	s.D.Set("apps_network_perimeters", appsNetworkPerimeters)

	if s.Res.AsOPCService != nil {
		s.D.Set("as_opc_service", []interface{}{AppAsOPCServiceToMap(s.Res.AsOPCService)})
	} else {
		s.D.Set("as_opc_service", nil)
	}

	attrRenderingMetadata := []interface{}{}
	for _, item := range s.Res.AttrRenderingMetadata {
		attrRenderingMetadata = append(attrRenderingMetadata, AppAttrRenderingMetadataToMap(item))
	}
	s.D.Set("attr_rendering_metadata", attrRenderingMetadata)

	if s.Res.Audience != nil {
		s.D.Set("audience", *s.Res.Audience)
	}

	if s.Res.BasedOnTemplate != nil {
		s.D.Set("based_on_template", []interface{}{AppBasedOnTemplateToMap(s.Res.BasedOnTemplate)})
	} else {
		s.D.Set("based_on_template", nil)
	}

	if s.Res.BypassConsent != nil {
		s.D.Set("bypass_consent", *s.Res.BypassConsent)
	}

	if s.Res.CallbackServiceUrl != nil {
		s.D.Set("callback_service_url", *s.Res.CallbackServiceUrl)
	}

	certificates := []interface{}{}
	for _, item := range s.Res.Certificates {
		certificates = append(certificates, AppCertificatesToMap(item))
	}
	s.D.Set("certificates", certificates)

	s.D.Set("client_ip_checking", s.Res.ClientIPChecking)

	if s.Res.ClientSecret != nil {
		s.D.Set("client_secret", *s.Res.ClientSecret)
	}

	s.D.Set("client_type", s.Res.ClientType)

	cloudControlProperties := []interface{}{}
	for _, item := range s.Res.CloudControlProperties {
		cloudControlProperties = append(cloudControlProperties, AppCloudControlPropertiesToMap(item))
	}
	s.D.Set("cloud_control_properties", cloudControlProperties)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.ContactEmailAddress != nil {
		s.D.Set("contact_email_address", *s.Res.ContactEmailAddress)
	}

	s.D.Set("delegated_service_names", s.Res.DelegatedServiceNames)

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisableKmsiTokenAuthentication != nil {
		s.D.Set("disable_kmsi_token_authentication", *s.Res.DisableKmsiTokenAuthentication)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	editableAttributes := []interface{}{}
	for _, item := range s.Res.EditableAttributes {
		editableAttributes = append(editableAttributes, AppEditableAttributesToMap(item))
	}
	s.D.Set("editable_attributes", editableAttributes)

	if s.Res.ErrorPageUrl != nil {
		s.D.Set("error_page_url", *s.Res.ErrorPageUrl)
	}

	grantedAppRoles := []interface{}{}
	for _, item := range s.Res.GrantedAppRoles {
		grantedAppRoles = append(grantedAppRoles, AppGrantedAppRolesToMap(item))
	}
	s.D.Set("granted_app_roles", grantedAppRoles)

	grants := []interface{}{}
	for _, item := range s.Res.Grants {
		grants = append(grants, AppGrantsToMap(item))
	}
	s.D.Set("grants", grants)

	if s.Res.HashedClientSecret != nil {
		s.D.Set("hashed_client_secret", *s.Res.HashedClientSecret)
	}

	if s.Res.HomePageUrl != nil {
		s.D.Set("home_page_url", *s.Res.HomePageUrl)
	}

	if s.Res.Icon != nil {
		s.D.Set("icon", *s.Res.Icon)
	}

	if s.Res.IdTokenEncAlgo != nil {
		s.D.Set("id_token_enc_algo", *s.Res.IdTokenEncAlgo)
	}

	if s.Res.IdcsCreatedBy != nil {
		s.D.Set("idcs_created_by", []interface{}{idcsCreatedByToMap(s.Res.IdcsCreatedBy)})
	} else {
		s.D.Set("idcs_created_by", nil)
	}

	if s.Res.IdcsLastModifiedBy != nil {
		s.D.Set("idcs_last_modified_by", []interface{}{idcsLastModifiedByToMap(s.Res.IdcsLastModifiedBy)})
	} else {
		s.D.Set("idcs_last_modified_by", nil)
	}

	if s.Res.IdcsLastUpgradedInRelease != nil {
		s.D.Set("idcs_last_upgraded_in_release", *s.Res.IdcsLastUpgradedInRelease)
	}

	s.D.Set("idcs_prevented_operations", s.Res.IdcsPreventedOperations)

	identityProviders := []interface{}{}
	for _, item := range s.Res.IdentityProviders {
		identityProviders = append(identityProviders, AppIdentityProvidersToMap(item))
	}
	s.D.Set("identity_providers", identityProviders)

	if s.Res.IdpPolicy != nil {
		s.D.Set("idp_policy", []interface{}{AppIdpPolicyToMap(s.Res.IdpPolicy)})
	} else {
		s.D.Set("idp_policy", nil)
	}

	if s.Res.Infrastructure != nil {
		s.D.Set("infrastructure", *s.Res.Infrastructure)
	}

	if s.Res.IsAliasApp != nil {
		s.D.Set("is_alias_app", *s.Res.IsAliasApp)
	}

	if s.Res.IsDatabaseService != nil {
		s.D.Set("is_database_service", *s.Res.IsDatabaseService)
	}

	if s.Res.IsEnterpriseApp != nil {
		s.D.Set("is_enterprise_app", *s.Res.IsEnterpriseApp)
	}

	if s.Res.IsFormFill != nil {
		s.D.Set("is_form_fill", *s.Res.IsFormFill)
	}

	if s.Res.IsKerberosRealm != nil {
		s.D.Set("is_kerberos_realm", *s.Res.IsKerberosRealm)
	}

	if s.Res.IsLoginTarget != nil {
		s.D.Set("is_login_target", *s.Res.IsLoginTarget)
	}

	if s.Res.IsManagedApp != nil {
		s.D.Set("is_managed_app", *s.Res.IsManagedApp)
	}

	if s.Res.IsMobileTarget != nil {
		s.D.Set("is_mobile_target", *s.Res.IsMobileTarget)
	}

	if s.Res.IsMulticloudServiceApp != nil {
		s.D.Set("is_multicloud_service_app", *s.Res.IsMulticloudServiceApp)
	}

	if s.Res.IsOAuthClient != nil {
		s.D.Set("is_oauth_client", *s.Res.IsOAuthClient)
	}

	if s.Res.IsOAuthResource != nil {
		s.D.Set("is_oauth_resource", *s.Res.IsOAuthResource)
	}

	if s.Res.IsOPCService != nil {
		s.D.Set("is_opc_service", *s.Res.IsOPCService)
	}

	if s.Res.IsObligationCapable != nil {
		s.D.Set("is_obligation_capable", *s.Res.IsObligationCapable)
	}

	if s.Res.IsRadiusApp != nil {
		s.D.Set("is_radius_app", *s.Res.IsRadiusApp)
	}

	if s.Res.IsSamlServiceProvider != nil {
		s.D.Set("is_saml_service_provider", *s.Res.IsSamlServiceProvider)
	}

	if s.Res.IsUnmanagedApp != nil {
		s.D.Set("is_unmanaged_app", *s.Res.IsUnmanagedApp)
	}

	if s.Res.IsWebTierPolicy != nil {
		s.D.Set("is_web_tier_policy", *s.Res.IsWebTierPolicy)
	}

	if s.Res.LandingPageUrl != nil {
		s.D.Set("landing_page_url", *s.Res.LandingPageUrl)
	}

	if s.Res.LinkingCallbackUrl != nil {
		s.D.Set("linking_callback_url", *s.Res.LinkingCallbackUrl)
	}

	s.D.Set("login_mechanism", s.Res.LoginMechanism)

	if s.Res.LoginPageUrl != nil {
		s.D.Set("login_page_url", *s.Res.LoginPageUrl)
	}

	if s.Res.LogoutPageUrl != nil {
		s.D.Set("logout_page_url", *s.Res.LogoutPageUrl)
	}

	if s.Res.LogoutUri != nil {
		s.D.Set("logout_uri", *s.Res.LogoutUri)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.MeterAsOPCService != nil {
		s.D.Set("meter_as_opc_service", *s.Res.MeterAsOPCService)
	}

	if s.Res.Migrated != nil {
		s.D.Set("migrated", *s.Res.Migrated)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	s.D.Set("post_logout_redirect_uris", s.Res.PostLogoutRedirectUris)

	if s.Res.PrivacyPolicyUrl != nil {
		s.D.Set("privacy_policy_url", *s.Res.PrivacyPolicyUrl)
	}

	if s.Res.ProductLogoUrl != nil {
		s.D.Set("product_logo_url", *s.Res.ProductLogoUrl)
	}

	if s.Res.ProductName != nil {
		s.D.Set("product_name", *s.Res.ProductName)
	}

	protectableSecondaryAudiences := []interface{}{}
	for _, item := range s.Res.ProtectableSecondaryAudiences {
		protectableSecondaryAudiences = append(protectableSecondaryAudiences, AppProtectableSecondaryAudiencesToMap(item))
	}
	s.D.Set("protectable_secondary_audiences", protectableSecondaryAudiences)

	if s.Res.RadiusPolicy != nil {
		s.D.Set("radius_policy", []interface{}{AppRadiusPolicyToMap(s.Res.RadiusPolicy)})
	} else {
		s.D.Set("radius_policy", nil)
	}

	if s.Res.ReadyToUpgrade != nil {
		s.D.Set("ready_to_upgrade", *s.Res.ReadyToUpgrade)
	}

	s.D.Set("redirect_uris", s.Res.RedirectUris)

	if s.Res.RefreshTokenExpiry != nil {
		s.D.Set("refresh_token_expiry", *s.Res.RefreshTokenExpiry)
	}

	if s.Res.SamlServiceProvider != nil {
		s.D.Set("saml_service_provider", []interface{}{AppSamlServiceProviderToMap(s.Res.SamlServiceProvider)})
	} else {
		s.D.Set("saml_service_provider", nil)
	}

	s.D.Set("schemas", s.Res.Schemas)

	scopes := []interface{}{}
	for _, item := range s.Res.Scopes {
		scopes = append(scopes, AppScopesToMap(item))
	}
	s.D.Set("scopes", scopes)

	s.D.Set("secondary_audiences", s.Res.SecondaryAudiences)

	serviceParams := []interface{}{}
	for _, item := range s.Res.ServiceParams {
		serviceParams = append(serviceParams, AppServiceParamsToMap(item))
	}
	s.D.Set("service_params", serviceParams)

	if s.Res.ServiceTypeURN != nil {
		s.D.Set("service_type_urn", *s.Res.ServiceTypeURN)
	}

	if s.Res.ServiceTypeVersion != nil {
		s.D.Set("service_type_version", *s.Res.ServiceTypeVersion)
	}

	if s.Res.ShowInMyApps != nil {
		s.D.Set("show_in_my_apps", *s.Res.ShowInMyApps)
	}

	if s.Res.SignonPolicy != nil {
		s.D.Set("signon_policy", []interface{}{AppSignonPolicyToMap(s.Res.SignonPolicy)})
	} else {
		s.D.Set("signon_policy", nil)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.TermsOfServiceUrl != nil {
		s.D.Set("terms_of_service_url", *s.Res.TermsOfServiceUrl)
	}

	if s.Res.TermsOfUse != nil {
		s.D.Set("terms_of_use", []interface{}{AppTermsOfUseToMap(s.Res.TermsOfUse)})
	} else {
		s.D.Set("terms_of_use", nil)
	}

	trustPolicies := []interface{}{}
	for _, item := range s.Res.TrustPolicies {
		trustPolicies = append(trustPolicies, AppTrustPoliciesToMap(item))
	}
	s.D.Set("trust_policies", trustPolicies)

	s.D.Set("trust_scope", s.Res.TrustScope)

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextension_oci_tags", []interface{}{ExtensionOCITagsToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextension_oci_tags", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondbcs_app", []interface{}{AppExtensionDbcsAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondbcs_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionEnterpriseAppApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionenterprise_app_app", []interface{}{AppExtensionEnterpriseAppAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionEnterpriseAppApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionenterprise_app_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionform_fill_app_app", []interface{}{AppExtensionFormFillAppAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionform_fill_app_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppTemplateAppTemplate != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template", []interface{}{AppExtensionFormFillAppTemplateAppTemplateToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppTemplateAppTemplate)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosRealmApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app", []interface{}{AppExtensionKerberosRealmAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosRealmApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionManagedappApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionmanagedapp_app", []interface{}{AppExtensionManagedappAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionManagedappApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionmanagedapp_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionMulticloudServiceAppApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app", []interface{}{AppExtensionMulticloudServiceAppAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionMulticloudServiceAppApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOpcServiceApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionopc_service_app", []interface{}{AppExtensionOpcServiceAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOpcServiceApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionopc_service_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionRadiusAppApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionradius_app_app", []interface{}{AppExtensionRadiusAppAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionRadiusAppApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionradius_app_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionrequestable_app", []interface{}{AppExtensionRequestableAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionrequestable_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSamlServiceProviderApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app", []interface{}{AppExtensionSamlServiceProviderAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSamlServiceProviderApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionWebTierPolicyApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app", []interface{}{AppExtensionWebTierPolicyAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionWebTierPolicyApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app", nil)
	}

	userRoles := []interface{}{}
	for _, item := range s.Res.UserRoles {
		userRoles = append(userRoles, AppUserRolesToMap(item))
	}
	s.D.Set("user_roles", userRoles)

	return nil
}
