package identity_domains

import (
	"fmt"
	"regexp"

	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportIdentityDomainsUserHints.GetIdFn = getGetIdFn("users")
	exportIdentityDomainsUserHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsGroupHints.GetIdFn = getGetIdFn("groups")
	exportIdentityDomainsGroupHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsDynamicResourceGroupHints.GetIdFn = getGetIdFn("dynamicResourceGroups")
	exportIdentityDomainsDynamicResourceGroupHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsPasswordPolicyHints.GetIdFn = getGetIdFn("passwordPolicies")
	exportIdentityDomainsPasswordPolicyHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsIdentityProviderHints.GetIdFn = getGetIdFn("identityProviders")
	exportIdentityDomainsIdentityProviderHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsAuthenticationFactorSettingHints.GetIdFn = getGetIdFn("authenticationFactorSettings")
	exportIdentityDomainsAuthenticationFactorSettingHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsKmsiSettingHints.GetIdFn = getGetIdFn("kmsiSettings")
	exportIdentityDomainsKmsiSettingHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsAccountRecoverySettingHints.GetIdFn = getGetIdFn("accountRecoverySettings")
	exportIdentityDomainsAccountRecoverySettingHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsIdentitySettingHints.GetIdFn = getGetIdFn("identitySettings")
	exportIdentityDomainsIdentitySettingHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsSecurityQuestionSettingHints.GetIdFn = getGetIdFn("securityQuestionSettings")
	exportIdentityDomainsSecurityQuestionSettingHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsGrantHints.GetIdFn = getGetIdFn("grants")
	exportIdentityDomainsGrantHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsAppRoleHints.GetIdFn = getGetIdFn("appRoles")
	exportIdentityDomainsAppRoleHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsAppHints.GetIdFn = getGetIdFn("apps")
	exportIdentityDomainsAppHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsSecurityQuestionHints.GetIdFn = getGetIdFn("securityQuestions")
	exportIdentityDomainsSecurityQuestionHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsCloudGateHints.GetIdFn = getGetIdFn("cloudGates")
	exportIdentityDomainsCloudGateHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsCloudGateServerHints.GetIdFn = getGetIdFn("cloudGateServers")
	exportIdentityDomainsCloudGateServerHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsCloudGateMappingHints.GetIdFn = getGetIdFn("cloudGateMappings")
	exportIdentityDomainsCloudGateMappingHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsSelfRegistrationProfileHints.GetIdFn = getGetIdFn("selfRegistrationProfiles")
	exportIdentityDomainsSelfRegistrationProfileHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsPolicyHints.GetIdFn = getGetIdFn("policies")
	exportIdentityDomainsPolicyHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsConditionHints.GetIdFn = getGetIdFn("conditions")
	exportIdentityDomainsConditionHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsNotificationSettingHints.GetIdFn = getGetIdFn("notificationSettings")
	exportIdentityDomainsNotificationSettingHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsSettingHints.GetIdFn = getGetIdFn("settings")
	exportIdentityDomainsSettingHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsOAuthClientCertificateHints.GetIdFn = getGetIdFn("oauthClientCertificates")
	exportIdentityDomainsOAuthClientCertificateHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsOAuthPartnerCertificateHints.GetIdFn = getGetIdFn("oauthPartnerCertificates")
	exportIdentityDomainsOAuthPartnerCertificateHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsNetworkPerimeterHints.GetIdFn = getGetIdFn("networkPerimeters")
	exportIdentityDomainsNetworkPerimeterHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsApprovalWorkflowAssignmentHints.GetIdFn = getGetIdFn("approvalWorkflowAssignments")
	exportIdentityDomainsApprovalWorkflowAssignmentHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsApprovalWorkflowStepHints.GetIdFn = getGetIdFn("approvalWorkflowSteps")
	exportIdentityDomainsApprovalWorkflowStepHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsApprovalWorkflowHints.GetIdFn = getGetIdFn("approvalWorkflows")
	exportIdentityDomainsApprovalWorkflowHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsRuleHints.GetIdFn = getGetIdFn("rules")
	exportIdentityDomainsRuleHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsIdentityPropagationTrustHints.GetIdFn = getGetIdFn("identityPropagationTrusts")
	exportIdentityDomainsIdentityPropagationTrustHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources

	//// User sub-resources
	exportIdentityDomainsApiKeyHints.GetIdFn = getUserSubResourceGetIdFn("apiKeys")
	exportIdentityDomainsApiKeyHints.ProcessDiscoveredResourcesFn = processUserSubResources
	exportIdentityDomainsApiKeyHints.FindResourcesOverrideFn = findUserSubResources
	exportIdentityDomainsAuthTokenHints.GetIdFn = getUserSubResourceGetIdFn("authTokens")
	exportIdentityDomainsAuthTokenHints.ProcessDiscoveredResourcesFn = processUserSubResources
	exportIdentityDomainsAuthTokenHints.FindResourcesOverrideFn = findUserSubResources
	exportIdentityDomainsCustomerSecretKeyHints.GetIdFn = getUserSubResourceGetIdFn("customerSecretKeys")
	exportIdentityDomainsCustomerSecretKeyHints.ProcessDiscoveredResourcesFn = processUserSubResources
	exportIdentityDomainsCustomerSecretKeyHints.FindResourcesOverrideFn = findUserSubResources
	exportIdentityDomainsOAuth2ClientCredentialHints.GetIdFn = getUserSubResourceGetIdFn("oAuth2ClientCredentials")
	exportIdentityDomainsOAuth2ClientCredentialHints.ProcessDiscoveredResourcesFn = processUserSubResources
	exportIdentityDomainsOAuth2ClientCredentialHints.FindResourcesOverrideFn = findUserSubResources
	exportIdentityDomainsSmtpCredentialHints.GetIdFn = getUserSubResourceGetIdFn("smtpCredentials")
	exportIdentityDomainsSmtpCredentialHints.ProcessDiscoveredResourcesFn = processUserSubResources
	exportIdentityDomainsSmtpCredentialHints.FindResourcesOverrideFn = findUserSubResources
	exportIdentityDomainsUserDbCredentialHints.GetIdFn = getUserSubResourceGetIdFn("userDbCredentials")
	exportIdentityDomainsUserDbCredentialHints.ProcessDiscoveredResourcesFn = processUserSubResources
	exportIdentityDomainsUserDbCredentialHints.FindResourcesOverrideFn = findUserSubResources

	// My/Self resources
	exportIdentityDomainsMyApiKeyHints.GetIdFn = getGetIdFn("myApiKeys")
	exportIdentityDomainsMyApiKeyHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsMyApiKeyHints.FindResourcesOverrideFn = getFindMyResources("myApiKeys")
	exportIdentityDomainsMyAuthTokenHints.GetIdFn = getGetIdFn("myAuthTokens")
	exportIdentityDomainsMyAuthTokenHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsMyAuthTokenHints.FindResourcesOverrideFn = getFindMyResources("myAuthTokens")
	exportIdentityDomainsMyCustomerSecretKeyHints.GetIdFn = getGetIdFn("myCustomerSecretKeys")
	exportIdentityDomainsMyCustomerSecretKeyHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsMyCustomerSecretKeyHints.FindResourcesOverrideFn = getFindMyResources("myCustomerSecretKeys")
	exportIdentityDomainsMyOAuth2ClientCredentialHints.GetIdFn = getGetIdFn("myOAuth2ClientCredentials")
	exportIdentityDomainsMyOAuth2ClientCredentialHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsMyOAuth2ClientCredentialHints.FindResourcesOverrideFn = getFindMyResources("myOAuth2ClientCredentials")
	exportIdentityDomainsMySmtpCredentialHints.GetIdFn = getGetIdFn("mySmtpCredentials")
	exportIdentityDomainsMySmtpCredentialHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsMySmtpCredentialHints.FindResourcesOverrideFn = getFindMyResources("mySmtpCredentials")
	exportIdentityDomainsMyUserDbCredentialHints.GetIdFn = getGetIdFn("myUserDbCredentials")
	exportIdentityDomainsMyUserDbCredentialHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsMyUserDbCredentialHints.FindResourcesOverrideFn = getFindMyResources("myUserDbCredentials")
	exportIdentityDomainsMySupportAccountHints.GetIdFn = getGetIdFn("mySupportAccounts")
	exportIdentityDomainsMySupportAccountHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsMySupportAccountHints.FindResourcesOverrideFn = getFindMyResources("mySupportAccounts")
	exportIdentityDomainsMyRequestHints.GetIdFn = getGetIdFn("myRequests")
	exportIdentityDomainsMyRequestHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsMyRequestHints.FindResourcesOverrideFn = getFindMyResources("myRequests")

	tf_export.RegisterTenancyGraphs("identity_domains", identityDomainsResourceGraph)

}

// Custom overrides for generating composite IDs within the resource discovery framework
func processIdentityDomainsResources(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}
		resource.SourceAttributes["idcs_endpoint"] = resource.Parent.SourceAttributes["url"].(string)
		resource.ImportId = resource.Id
	}
	return resources, nil
}

func getResourceUrl(resource *tf_export.OCIResource) (string, error) {
	idcsEndpoint, ok := resource.SourceAttributes["url"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find url for idcsEndpoint")
	}
	return idcsEndpoint, nil
}

func getResourceId(resource *tf_export.OCIResource) (string, error) {
	var resourceMap map[string]interface{} = resource.RawResource.(map[string]interface{})
	id, ok := resourceMap["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find id in resourceMap")
	}

	return id, nil
}

func getDomainResourceInfo(resource *tf_export.OCIResource) (string, string, error) {
	idcsEndpoint, err := getResourceUrl(resource.Parent)
	if err != nil {
		return "", "", err
	}
	id, err := getResourceId(resource)
	if err != nil {
		return idcsEndpoint, "", err
	}

	return idcsEndpoint, id, nil
}

func getGetIdFn(resourceName string) func(*tf_export.OCIResource) (string, error) {
	return func(resource *tf_export.OCIResource) (string, error) {
		idcsEndpoint, id, err := getDomainResourceInfo(resource)
		if err != nil {
			return "", err
		}

		return GetIdentityDomainsCompositeId(idcsEndpoint, resourceName, id), nil
	}
}

// region User sub-resources
func findUserSubResources(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
	licenseType, ok := parent.Parent.SourceAttributes["license_type"].(string)
	if !ok {
		return nil, fmt.Errorf("[ERROR] Failed to get domain licenseType from parent resource attribute `license_type`")
	}
	// User subresources e.g. api keys, auth tokens, etc are not available in external-user domain
	if licenseType == "external-user" {
		return nil, nil
	}

	return tf_export.FindResourcesGeneric(ctx, tfMeta, parent, resourceGraph)
}

func processUserSubResources(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}
		resource.SourceAttributes["idcs_endpoint"] = resource.Parent.SourceAttributes["idcs_endpoint"].(string)
		resource.ImportId = resource.Id
	}
	return resources, nil
}

func getUserSubResourceGetIdFn(resourceName string) func(*tf_export.OCIResource) (string, error) {
	return func(resource *tf_export.OCIResource) (string, error) {
		idcsEndpoint, err := getResourceUrl(resource.Parent.Parent)
		if err != nil {
			return "", err
		}
		id, err := getResourceId(resource)
		if err != nil {
			return "", err
		}
		return GetIdentityDomainsCompositeId(idcsEndpoint, resourceName, id), nil
	}
}

// endregion User sub-resources

func getFindMyResources(resourceName string) func(*tf_export.ResourceDiscoveryContext, *tf_export.TerraformResourceAssociation, *tf_export.OCIResource, *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
	return func(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
		idcsEndpoint, ok := parent.SourceAttributes["url"].(string)
		if !ok {
			return nil, fmt.Errorf("[ERROR] Failed to get domain url from parent resource attribute `url`")
		}

		// My* resources only allows GET with current domain's endpoint url,
		// so only call Find when idcs_endpoint matches
		for resourceId, _ := range ctx.ExpectedResourceIds {
			regex, _ := regexp.Compile("^idcsEndpoint/(.*)/" + resourceName + "/(.*)$")
			tokens := regex.FindStringSubmatch(resourceId)
			if len(tokens) == 3 && tokens[1] == idcsEndpoint {
				return tf_export.FindResourcesGeneric(ctx, tfMeta, parent, resourceGraph)
			}
		}

		return nil, nil
	}
}

// Hints for discovering and exporting this resource to configuration and state files
var exportIdentityDomainsUserHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_user",
	DatasourceClass:      "oci_identity_domains_users",
	DatasourceItemsAttr:  "users",
	ResourceAbbreviation: "user",
}

var exportIdentityDomainsPasswordPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_password_policy",
	DatasourceClass:      "oci_identity_domains_password_policies",
	DatasourceItemsAttr:  "password_policies",
	ResourceAbbreviation: "password_policy",
}

var exportIdentityDomainsSmtpCredentialHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_smtp_credential",
	DatasourceClass:      "oci_identity_domains_smtp_credentials",
	DatasourceItemsAttr:  "smtp_credentials",
	ResourceAbbreviation: "smtp_credential",
}

var exportIdentityDomainsApiKeyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_api_key",
	DatasourceClass:      "oci_identity_domains_api_keys",
	DatasourceItemsAttr:  "api_keys",
	ResourceAbbreviation: "api_key",
}

var exportIdentityDomainsMySupportAccountHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_my_support_account",
	DatasourceClass:      "oci_identity_domains_my_support_accounts",
	DatasourceItemsAttr:  "my_support_accounts",
	ResourceAbbreviation: "my_support_account",
}

var exportIdentityDomainsCustomerSecretKeyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_customer_secret_key",
	DatasourceClass:      "oci_identity_domains_customer_secret_keys",
	DatasourceItemsAttr:  "customer_secret_keys",
	ResourceAbbreviation: "customer_secret_key",
}

var exportIdentityDomainsKmsiSettingHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_kmsi_setting",
	DatasourceClass:      "oci_identity_domains_kmsi_settings",
	DatasourceItemsAttr:  "kmsi_settings",
	ResourceAbbreviation: "kmsi_setting",
}

var exportIdentityDomainsMyAuthTokenHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_my_auth_token",
	DatasourceClass:      "oci_identity_domains_my_auth_tokens",
	DatasourceItemsAttr:  "my_auth_tokens",
	ResourceAbbreviation: "my_auth_token",
}

var exportIdentityDomainsUserDbCredentialHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_user_db_credential",
	DatasourceClass:      "oci_identity_domains_user_db_credentials",
	DatasourceItemsAttr:  "user_db_credentials",
	ResourceAbbreviation: "user_db_credential",
}

var exportIdentityDomainsIdentityProviderHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_identity_provider",
	DatasourceClass:      "oci_identity_domains_identity_providers",
	DatasourceItemsAttr:  "identity_providers",
	ResourceAbbreviation: "identity_provider",
}

var exportIdentityDomainsAuthTokenHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_auth_token",
	DatasourceClass:      "oci_identity_domains_auth_tokens",
	DatasourceItemsAttr:  "auth_tokens",
	ResourceAbbreviation: "auth_token",
}

var exportIdentityDomainsDynamicResourceGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_dynamic_resource_group",
	DatasourceClass:      "oci_identity_domains_dynamic_resource_groups",
	DatasourceItemsAttr:  "dynamic_resource_groups",
	ResourceAbbreviation: "dynamic_resource_group",
}

var exportIdentityDomainsMyCustomerSecretKeyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_my_customer_secret_key",
	DatasourceClass:      "oci_identity_domains_my_customer_secret_keys",
	DatasourceItemsAttr:  "my_customer_secret_keys",
	ResourceAbbreviation: "my_customer_secret_key",
}

var exportIdentityDomainsMyUserDbCredentialHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_my_user_db_credential",
	DatasourceClass:      "oci_identity_domains_my_user_db_credentials",
	DatasourceItemsAttr:  "my_user_db_credentials",
	ResourceAbbreviation: "my_user_db_credential",
}

var exportIdentityDomainsMySmtpCredentialHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_my_smtp_credential",
	DatasourceClass:      "oci_identity_domains_my_smtp_credentials",
	DatasourceItemsAttr:  "my_smtp_credentials",
	ResourceAbbreviation: "my_smtp_credential",
}

var exportIdentityDomainsMyApiKeyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_my_api_key",
	DatasourceClass:      "oci_identity_domains_my_api_keys",
	DatasourceItemsAttr:  "my_api_keys",
	ResourceAbbreviation: "my_api_key",
}

var exportIdentityDomainsGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_group",
	DatasourceClass:      "oci_identity_domains_groups",
	DatasourceItemsAttr:  "groups",
	ResourceAbbreviation: "group",
}

var exportIdentityDomainsOAuth2ClientCredentialHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_oauth2client_credential",
	DatasourceClass:      "oci_identity_domains_oauth2client_credentials",
	DatasourceItemsAttr:  "oauth2client_credentials",
	ResourceAbbreviation: "oauth2client_credential",
}

var exportIdentityDomainsMyOAuth2ClientCredentialHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_my_oauth2client_credential",
	DatasourceClass:      "oci_identity_domains_my_oauth2client_credentials",
	DatasourceItemsAttr:  "my_oauth2client_credentials",
	ResourceAbbreviation: "my_oauth2client_credential",
}

var exportIdentityDomainsAuthenticationFactorSettingHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_authentication_factor_setting",
	DatasourceClass:      "oci_identity_domains_authentication_factor_settings",
	DatasourceItemsAttr:  "authentication_factor_settings",
	ResourceAbbreviation: "authentication_factor_setting",
}

var exportIdentityDomainsAccountRecoverySettingHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_account_recovery_setting",
	DatasourceClass:      "oci_identity_domains_account_recovery_settings",
	DatasourceItemsAttr:  "account_recovery_settings",
	ResourceAbbreviation: "account_recovery_setting",
}

var exportIdentityDomainsIdentitySettingHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_identity_setting",
	DatasourceClass:      "oci_identity_domains_identity_settings",
	DatasourceItemsAttr:  "identity_settings",
	ResourceAbbreviation: "identity_setting",
}

var exportIdentityDomainsMyRequestHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_my_request",
	DatasourceClass:      "oci_identity_domains_my_requests",
	DatasourceItemsAttr:  "my_requests",
	ResourceAbbreviation: "my_request",
}

var exportIdentityDomainsSecurityQuestionSettingHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_security_question_setting",
	DatasourceClass:      "oci_identity_domains_security_question_settings",
	DatasourceItemsAttr:  "security_question_settings",
	ResourceAbbreviation: "security_question_setting",
}

var exportIdentityDomainsGrantHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_grant",
	DatasourceClass:      "oci_identity_domains_grants",
	DatasourceItemsAttr:  "grants",
	ResourceAbbreviation: "grant",
}

var exportIdentityDomainsAppRoleHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_app_role",
	DatasourceClass:      "oci_identity_domains_app_roles",
	DatasourceItemsAttr:  "app_roles",
	ResourceAbbreviation: "app_role",
}

var exportIdentityDomainsAppHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_app",
	DatasourceClass:      "oci_identity_domains_apps",
	DatasourceItemsAttr:  "apps",
	ResourceAbbreviation: "app",
}

var exportIdentityDomainsSecurityQuestionHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_security_question",
	DatasourceClass:      "oci_identity_domains_security_questions",
	DatasourceItemsAttr:  "security_questions",
	ResourceAbbreviation: "security_question",
}

/*
IDCP oci_identity_domain resource is the dependency for all identity_domains resource.
Thus add the hint here, which is the same as the hint in identity_export.
*/
var exportIdentityDomainHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domain",
	DatasourceClass:      "oci_identity_domains",
	DatasourceItemsAttr:  "domains",
	ResourceAbbreviation: "domain",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.DomainLifecycleStateActive),
	},
}

var exportIdentityDomainsApprovalWorkflowAssignmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_approval_workflow_assignment",
	DatasourceClass:      "oci_identity_domains_approval_workflow_assignments",
	DatasourceItemsAttr:  "approval_workflow_assignments",
	ResourceAbbreviation: "approval_workflow_assignment",
}

var exportIdentityDomainsApprovalWorkflowStepHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_approval_workflow_step",
	DatasourceClass:      "oci_identity_domains_approval_workflow_steps",
	DatasourceItemsAttr:  "approval_workflow_steps",
	ResourceAbbreviation: "approval_workflow_step",
}

var exportIdentityDomainsApprovalWorkflowHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_approval_workflow",
	DatasourceClass:      "oci_identity_domains_approval_workflows",
	DatasourceItemsAttr:  "approval_workflows",
	ResourceAbbreviation: "approval_workflow",
}

var exportIdentityDomainsCloudGateMappingHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_cloud_gate_mapping",
	DatasourceClass:      "oci_identity_domains_cloud_gate_mappings",
	DatasourceItemsAttr:  "cloud_gate_mappings",
	ResourceAbbreviation: "cloud_gate_mapping",
}

var exportIdentityDomainsCloudGateServerHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_cloud_gate_server",
	DatasourceClass:      "oci_identity_domains_cloud_gate_servers",
	DatasourceItemsAttr:  "cloud_gate_servers",
	ResourceAbbreviation: "cloud_gate_server",
}

var exportIdentityDomainsCloudGateHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_cloud_gate",
	DatasourceClass:      "oci_identity_domains_cloud_gates",
	DatasourceItemsAttr:  "cloud_gates",
	ResourceAbbreviation: "cloud_gate",
}

var exportIdentityDomainsConditionHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_condition",
	DatasourceClass:      "oci_identity_domains_conditions",
	DatasourceItemsAttr:  "conditions",
	ResourceAbbreviation: "condition",
}

var exportIdentityDomainsNetworkPerimeterHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_network_perimeter",
	DatasourceClass:      "oci_identity_domains_network_perimeters",
	DatasourceItemsAttr:  "network_perimeters",
	ResourceAbbreviation: "network_perimeter",
}

var exportIdentityDomainsPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_policy",
	DatasourceClass:      "oci_identity_domains_policies",
	DatasourceItemsAttr:  "policies",
	ResourceAbbreviation: "policy",
}

var exportIdentityDomainsSelfRegistrationProfileHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_self_registration_profile",
	DatasourceClass:      "oci_identity_domains_self_registration_profiles",
	DatasourceItemsAttr:  "self_registration_profiles",
	ResourceAbbreviation: "self_registration_profile",
}

var exportIdentityDomainsNotificationSettingHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_notification_setting",
	DatasourceClass:      "oci_identity_domains_notification_settings",
	DatasourceItemsAttr:  "notification_settings",
	ResourceAbbreviation: "notification_setting",
}

var exportIdentityDomainsSettingHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_setting",
	DatasourceClass:      "oci_identity_domains_settings",
	DatasourceItemsAttr:  "settings",
	ResourceAbbreviation: "setting",
}

var exportIdentityDomainsOAuthClientCertificateHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_oauth_client_certificate",
	DatasourceClass:      "oci_identity_domains_oauth_client_certificates",
	DatasourceItemsAttr:  "oauth_client_certificates",
	ResourceAbbreviation: "oauth_client_certificate",
}

var exportIdentityDomainsOAuthPartnerCertificateHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_oauth_partner_certificate",
	DatasourceClass:      "oci_identity_domains_oauth_partner_certificates",
	DatasourceItemsAttr:  "oauth_partner_certificates",
	ResourceAbbreviation: "oauth_partner_certificate",
}

var exportIdentityDomainsRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_rule",
	DatasourceClass:      "oci_identity_domains_rules",
	DatasourceItemsAttr:  "rules",
	ResourceAbbreviation: "rule",
}

var exportIdentityDomainsIdentityPropagationTrustHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_identity_propagation_trust",
	DatasourceClass:      "oci_identity_domains_identity_propagation_trusts",
	DatasourceItemsAttr:  "identity_propagation_trusts",
	ResourceAbbreviation: "identity_propagation_trust",
}

var identityDomainsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportIdentityDomainHints},
	},
	"oci_identity_domain": {
		{
			TerraformResourceHints: exportIdentityDomainsUserHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsGroupHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsPasswordPolicyHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsIdentityProviderHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsDynamicResourceGroupHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsKmsiSettingHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsMyApiKeyHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsMyAuthTokenHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsMyCustomerSecretKeyHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsMyOAuth2ClientCredentialHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsMySmtpCredentialHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsMyUserDbCredentialHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsMySupportAccountHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsAuthenticationFactorSettingHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsIdentitySettingHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsAccountRecoverySettingHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsGrantHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsSecurityQuestionSettingHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsMyRequestHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsAppHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsAppRoleHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsCloudGateHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsCloudGateServerHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsCloudGateMappingHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsSelfRegistrationProfileHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsPolicyHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsConditionHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsNotificationSettingHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsSettingHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsNetworkPerimeterHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsOAuthClientCertificateHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsOAuthPartnerCertificateHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsApprovalWorkflowAssignmentHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsApprovalWorkflowStepHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsApprovalWorkflowHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsRuleHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsIdentityPropagationTrustHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint": "url",
			},
		},
	},
	"oci_identity_domains_user": {
		{
			TerraformResourceHints: exportIdentityDomainsApiKeyHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint":  "idcs_endpoint",
				"api_key_filter": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsAuthTokenHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint":     "idcs_endpoint",
				"auth_token_filter": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsCustomerSecretKeyHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint":              "idcs_endpoint",
				"customer_secret_key_filter": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsOAuth2ClientCredentialHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint":                  "idcs_endpoint",
				"oauth2client_credential_filter": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsSmtpCredentialHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint":          "idcs_endpoint",
				"smtp_credential_filter": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityDomainsUserDbCredentialHints,
			DatasourceQueryParams: map[string]string{
				"idcs_endpoint":             "idcs_endpoint",
				"user_db_credential_filter": "id",
			},
		},
	},
}
