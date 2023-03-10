package identity_domains

import (
	"fmt"
	"regexp"

	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportIdentityDomainsGroupHints.GetIdFn = getGetIdFn("groups")
	exportIdentityDomainsGroupHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsUserHints.GetIdFn = getGetIdFn("users")
	exportIdentityDomainsUserHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsDynamicResourceGroupHints.GetIdFn = getGetIdFn("dynamicResourceGroups")
	exportIdentityDomainsDynamicResourceGroupHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsPasswordPolicyHints.GetIdFn = getGetIdFn("passwordPolicies")
	exportIdentityDomainsPasswordPolicyHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsIdentityProviderHints.GetIdFn = getGetIdFn("identityProviders")
	exportIdentityDomainsIdentityProviderHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsKmsiSettingHints.GetIdFn = getGetIdFn("kmsiSettings")
	exportIdentityDomainsKmsiSettingHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources
	exportIdentityDomainsAuthenticationFactorSettingHints.GetIdFn = getGetIdFn("authenticationFactorSettings")
	exportIdentityDomainsAuthenticationFactorSettingHints.ProcessDiscoveredResourcesFn = processIdentityDomainsResources

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

	tf_export.RegisterTenancyGraphs("identity_domains", identityDomainsResourceGraph)

}

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

// Custom overrides for generating composite IDs within the resource discovery framework

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

//region User sub-resources
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

//endregion User sub-resources

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

var exportIdentityDomainsUserHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domains_user",
	DatasourceClass:      "oci_identity_domains_users",
	DatasourceItemsAttr:  "users",
	ResourceAbbreviation: "user",
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

/**
Same as in identity_export.go
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
