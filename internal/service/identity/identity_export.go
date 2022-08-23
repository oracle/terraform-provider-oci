package identity

import (
	"context"
	"fmt"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportIdentityApiKeyHints.GetIdFn = getIdentityApiKeyId
	exportIdentityAuthTokenHints.GetIdFn = getIdentityAuthTokenId
	exportIdentityCustomerSecretKeyHints.GetIdFn = getIdentityCustomerSecretKeyId
	exportIdentityIdpGroupMappingHints.GetIdFn = getIdentityIdpGroupMappingId
	exportIdentitySmtpCredentialHints.GetIdFn = getIdentitySmtpCredentialId
	exportIdentitySwiftPasswordHints.GetIdFn = getIdentitySwiftPasswordId
	exportIdentityDbCredentialHints.GetIdFn = getIdentityDbCredentialId
	exportIdentityAvailabilityDomainHints.IsDataSource = true
	exportIdentityAvailabilityDomainHints.ResourceAbbreviation = "ad"
	exportIdentityAvailabilityDomainHints.AlwaysExportable = true
	exportIdentityAvailabilityDomainHints.ProcessDiscoveredResourcesFn = processAvailabilityDomains
	exportIdentityAvailabilityDomainHints.GetHCLStringOverrideFn = getAvailabilityDomainHCLDatasource
	exportIdentityAuthenticationPolicyHints.ProcessDiscoveredResourcesFn = processIdentityAuthenticationPolicies
	exportIdentityTagHints.FindResourcesOverrideFn = findIdentityTags
	exportIdentityTagHints.ProcessDiscoveredResourcesFn = processTagDefinitions
	tf_export.RegisterTenancyGraphs("identity", identityResourceGraph)
	tf_export.RegisterCompartmentGraphs("tagging", taggingResourceGraph)
	tf_export.BuildAvailabilityResourceGraph("oci_identity_compartment", customAssociationIdentityCompartment)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func getIdentityDbCredentialId(resource *tf_export.OCIResource) (string, error) {

	dbCredentialId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dbCredentialId for Identity DbCredential")
	}
	userId := resource.Parent.Id
	return GetDbCredentialCompositeId(dbCredentialId, userId), nil
}

func processTagDefinitions(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		if resource.Parent == nil {
			resource.ImportId = fmt.Sprintf("tagNamespaces/%s/tags/%s", resource.SourceAttributes["tag_namespace_id"], resource.SourceAttributes["name"].(string))
			continue
		}

		resource.SourceAttributes["tag_namespace_id"] = resource.Parent.Id
		resource.ImportId = fmt.Sprintf("tagNamespaces/%s/tags/%s", resource.Parent.Id, resource.SourceAttributes["name"].(string))
		resource.Id = resource.ImportId
	}
	return resources, nil
}

func findIdentityTags(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
	// List on Tags does not return validator, and resource Read requires tagNamespaceId
	// which is also not returned in Summary response. Tags also do not have composite id in state.
	// Getting tags using ListTagsRequest and the calling tagResource.Read
	tagNamespaceId := parent.Id
	request := oci_identity.ListTagsRequest{}

	request.TagNamespaceId = &tagNamespaceId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")
	results := []*tf_export.OCIResource{}

	response, err := ctx.Clients.IdentityClient().ListTags(context.Background(), request)
	if err != nil {
		return results, err
	}

	request.Page = response.OpcNextPage

	for request.Page != nil {
		listResponse, err := ctx.Clients.IdentityClient().ListTags(context.Background(), request)
		if err != nil {
			return results, err
		}

		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	for _, tag := range response.Items {
		tagResource := tf_export.ResourcesMap[tfMeta.ResourceClass]

		d := tagResource.TestResourceData()
		d.SetId(GetIdentityTagCompositeId(*tag.Name, parent.Id))

		if err := tagResource.Read(d, ctx.Clients); err != nil {
			rdError := &tf_export.ResourceDiscoveryError{ResourceType: tfMeta.ResourceClass, ParentResource: parent.TerraformName, Error: err, ResourceGraph: resourceGraph}
			ctx.AddErrorToList(rdError)
			continue
		}

		resource := &tf_export.OCIResource{
			CompartmentId:    parent.CompartmentId,
			SourceAttributes: tf_export.ConvertResourceDataToMap(tagResource.Schema, d),
			RawResource:      tag,
			TerraformResource: tf_export.TerraformResource{
				Id:             d.Id(),
				TerraformClass: tfMeta.ResourceClass,
			},
			GetHclStringFn: tf_export.GetHclStringFromGenericMap,
			Parent:         parent,
		}

		if resource.TerraformName, err = tf_export.GenerateTerraformNameFromResource(resource.SourceAttributes, tagResource.Schema); err != nil {
			resource.TerraformName = fmt.Sprintf("%s_%s", parent.Parent.TerraformName, *tag.Name)
		}

		results = append(results, resource)
	}

	return results, nil

}

func processIdentityAuthenticationPolicies(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	// Add composite id as the resource's import ID
	for _, resource := range resources {
		resource.ImportId = GetAuthenticationPolicyCompositeId(resource.CompartmentId)
		resource.Id = resource.ImportId
	}
	return resources, nil
}

func processAvailabilityDomains(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for idx, ad := range resources {
		ad.SourceAttributes["index"] = idx + 1

		adName, ok := ad.SourceAttributes["name"].(string)
		if !ok || adName == "" {
			return resources, fmt.Errorf("[ERROR] availability domain at index '%v' has no name\n", idx)
		}
		tf_export.RefMapLock.Lock()
		tf_export.ReferenceMap[adName] = tf_export.TfHclVersionvar.GetDataSourceHclString(ad.GetTerraformReference(), "name")
		tf_export.RefMapLock.Unlock()
	}

	return resources, nil
}

func getAvailabilityDomainHCLDatasource(builder *strings.Builder, ociRes *tf_export.OCIResource, varMap map[string]string) error {
	builder.WriteString(fmt.Sprintf("data %s %s {\n", ociRes.TerraformClass, ociRes.TerraformName))

	builder.WriteString(fmt.Sprintf("compartment_id = %v\n", varMap[ociRes.CompartmentId]))

	adIndex, ok := ociRes.SourceAttributes["index"]
	if !ok {
		return fmt.Errorf("[ERROR] no index found for availability domain '%s'", ociRes.GetTerraformReference())
	}
	builder.WriteString(fmt.Sprintf("ad_number = \"%v\"\n", adIndex.(int)))
	builder.WriteString("}\n")

	return nil
}

func getIdentityApiKeyId(resource *tf_export.OCIResource) (string, error) {

	fingerprint, ok := resource.SourceAttributes["fingerprint"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find fingerprint for Identity ApiKey")
	}
	userId := resource.Parent.Id
	return GetApiKeyCompositeId(fingerprint, userId), nil
}

func getIdentityAuthenticationPolicyId(resource *tf_export.OCIResource) (string, error) {

	compartmentId, ok := resource.SourceAttributes["compartment_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find compartmentId for Identity AuthenticationPolicy")
	}
	return GetAuthenticationPolicyCompositeId(compartmentId), nil
}

func getIdentityAuthTokenId(resource *tf_export.OCIResource) (string, error) {

	authTokenId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find authTokenId for Identity AuthToken")
	}
	userId := resource.Parent.Id
	return GetAuthTokenCompositeId(authTokenId, userId), nil
}

func getIdentityCustomerSecretKeyId(resource *tf_export.OCIResource) (string, error) {

	customerSecretKeyId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find customerSecretKeyId for Identity CustomerSecretKey")
	}
	userId := resource.Parent.Id
	return GetCustomerSecretKeyCompositeId(customerSecretKeyId, userId), nil
}

func getIdentityIdpGroupMappingId(resource *tf_export.OCIResource) (string, error) {

	identityProviderId := resource.Parent.Id
	mappingId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find mappingId for Identity IdpGroupMapping")
	}
	return GetIdpGroupMappingCompositeId(identityProviderId, mappingId), nil
}

func getIdentitySmtpCredentialId(resource *tf_export.OCIResource) (string, error) {

	smtpCredentialId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find smtpCredentialId for Identity SmtpCredential")
	}
	userId := resource.Parent.Id
	return GetSmtpCredentialCompositeId(smtpCredentialId, userId), nil
}

func getIdentitySwiftPasswordId(resource *tf_export.OCIResource) (string, error) {

	swiftPasswordId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find swiftPasswordId for Identity SwiftPassword")
	}
	userId := resource.Parent.Id
	return GetSwiftPasswordCompositeId(swiftPasswordId, userId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportIdentityApiKeyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_api_key",
	DatasourceClass:      "oci_identity_api_keys",
	DatasourceItemsAttr:  "api_keys",
	ResourceAbbreviation: "api_key",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.ApiKeyLifecycleStateActive),
	},
}

var exportIdentityAvailabilityDomainHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_availability_domain",
	DatasourceClass:      "oci_identity_availability_domains",
	DatasourceItemsAttr:  "availability_domains",
	ResourceAbbreviation: "availability_domain",
}

var exportIdentityAuthenticationPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_authentication_policy",
	DatasourceClass:      "oci_identity_authentication_policy",
	ResourceAbbreviation: "authentication_policy",
}

var exportIdentityAuthTokenHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_auth_token",
	DatasourceClass:      "oci_identity_auth_tokens",
	DatasourceItemsAttr:  "tokens",
	ResourceAbbreviation: "auth_token",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.AuthTokenLifecycleStateActive),
	},
}

var exportIdentityCompartmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_compartment",
	DatasourceClass:      "oci_identity_compartments",
	DatasourceItemsAttr:  "compartments",
	ResourceAbbreviation: "compartment",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.CompartmentLifecycleStateActive),
	},
}

var exportIdentityCustomerSecretKeyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_customer_secret_key",
	DatasourceClass:      "oci_identity_customer_secret_keys",
	DatasourceItemsAttr:  "customer_secret_keys",
	ResourceAbbreviation: "customer_secret_key",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.CustomerSecretKeyLifecycleStateActive),
	},
}

var exportIdentityDynamicGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_dynamic_group",
	DatasourceClass:      "oci_identity_dynamic_groups",
	DatasourceItemsAttr:  "dynamic_groups",
	ResourceAbbreviation: "dynamic_group",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.DynamicGroupLifecycleStateActive),
	},
}

var exportIdentityGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_group",
	DatasourceClass:      "oci_identity_groups",
	DatasourceItemsAttr:  "groups",
	ResourceAbbreviation: "Group",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.GroupLifecycleStateActive),
	},
}

var exportIdentityIdentityProviderHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_identity_provider",
	DatasourceClass:      "oci_identity_identity_providers",
	DatasourceItemsAttr:  "identity_providers",
	ResourceAbbreviation: "identity_provider",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.IdentityProviderLifecycleStateActive),
	},
}

var exportIdentityIdpGroupMappingHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_idp_group_mapping",
	DatasourceClass:      "oci_identity_idp_group_mappings",
	DatasourceItemsAttr:  "idp_group_mappings",
	ResourceAbbreviation: "idp_group_mapping",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.IdpGroupMappingLifecycleStateActive),
	},
}

var exportIdentityPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_policy",
	DatasourceClass:      "oci_identity_policies",
	DatasourceItemsAttr:  "policies",
	ResourceAbbreviation: "policy",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.PolicyLifecycleStateActive),
	},
}

var exportIdentitySmtpCredentialHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_smtp_credential",
	DatasourceClass:      "oci_identity_smtp_credentials",
	DatasourceItemsAttr:  "smtp_credentials",
	ResourceAbbreviation: "smtp_credential",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.SmtpCredentialLifecycleStateActive),
	},
}

var exportIdentitySwiftPasswordHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_swift_password",
	DatasourceClass:      "oci_identity_swift_passwords",
	DatasourceItemsAttr:  "passwords",
	ResourceAbbreviation: "swift_password",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.SwiftPasswordLifecycleStateActive),
	},
}

var exportIdentityUiPasswordHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_ui_password",
	DatasourceClass:      "oci_identity_ui_password",
	ResourceAbbreviation: "ui_password",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.UiPasswordLifecycleStateActive),
	},
}

var exportIdentityUserGroupMembershipHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_user_group_membership",
	DatasourceClass:      "oci_identity_user_group_memberships",
	DatasourceItemsAttr:  "memberships",
	ResourceAbbreviation: "user_group_membership",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.UserGroupMembershipLifecycleStateActive),
	},
}

var exportIdentityUserHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_user",
	DatasourceClass:      "oci_identity_users",
	DatasourceItemsAttr:  "users",
	ResourceAbbreviation: "user",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.UserLifecycleStateActive),
	},
}

var exportIdentityTagDefaultHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_tag_default",
	DatasourceClass:      "oci_identity_tag_defaults",
	DatasourceItemsAttr:  "tag_defaults",
	ResourceAbbreviation: "tag_default",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.TagDefaultLifecycleStateActive),
	},
}

var exportIdentityTagNamespaceHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_tag_namespace",
	DatasourceClass:      "oci_identity_tag_namespaces",
	DatasourceItemsAttr:  "tag_namespaces",
	ResourceAbbreviation: "tag_namespace",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.TagNamespaceLifecycleStateActive),
	},
}

var exportIdentityTagHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_identity_tag",
	DatasourceClass:        "oci_identity_tags",
	DatasourceItemsAttr:    "tags",
	ResourceAbbreviation:   "tag",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_identity.TagLifecycleStateActive),
	},
}

var exportIdentityNetworkSourceHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_network_source",
	DatasourceClass:      "oci_identity_network_sources",
	DatasourceItemsAttr:  "network_sources",
	ResourceAbbreviation: "network_source",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.NetworkSourcesLifecycleStateActive),
	},
}

var exportIdentityDomainHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_domain",
	DatasourceClass:      "oci_identity_domains",
	DatasourceItemsAttr:  "domains",
	ResourceAbbreviation: "domain",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.DomainLifecycleStateActive),
	},
}

var exportIdentityDbCredentialHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_db_credential",
	DatasourceClass:      "oci_identity_db_credentials",
	DatasourceItemsAttr:  "db_credentials",
	ResourceAbbreviation: "db_credential",
	DiscoverableLifecycleStates: []string{
		string(oci_identity.DbCredentialLifecycleStateActive),
	},
}

var exportIdentityImportStandardTagsManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_import_standard_tags_management",
	ResourceAbbreviation: "import_standard_tags_management",
}

var identityResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportIdentityAuthenticationPolicyHints},
		{TerraformResourceHints: exportIdentityCompartmentHints},
		{TerraformResourceHints: exportIdentityDynamicGroupHints},
		{TerraformResourceHints: exportIdentityGroupHints},
		{
			TerraformResourceHints: exportIdentityIdentityProviderHints,
			DatasourceQueryParams:  map[string]string{"protocol": "'SAML2'"},
		},
		{TerraformResourceHints: exportIdentityPolicyHints},
		{TerraformResourceHints: exportIdentityUserHints},
		{TerraformResourceHints: exportIdentityNetworkSourceHints},
		{TerraformResourceHints: exportIdentityDomainHints},
	},
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportIdentityCompartmentHints,
			DatasourceQueryParams:  map[string]string{"compartment_id": "id"},
		},
		{
			TerraformResourceHints: exportIdentityPolicyHints,
			DatasourceQueryParams:  map[string]string{"compartment_id": "id"},
		},
	},
	"oci_identity_identity_provider": {
		{
			TerraformResourceHints: exportIdentityIdpGroupMappingHints,
			DatasourceQueryParams: map[string]string{
				"identity_provider_id": "id",
			},
		},
	},
	"oci_identity_user": {
		{
			TerraformResourceHints: exportIdentityApiKeyHints,
			DatasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityAuthTokenHints,
			DatasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityCustomerSecretKeyHints,
			DatasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityDbCredentialHints,
			DatasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentitySmtpCredentialHints,
			DatasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentitySwiftPasswordHints,
			DatasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityUiPasswordHints,
			DatasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityUserGroupMembershipHints,
			DatasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
	},
}

var taggingResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportIdentityTagDefaultHints},
		{TerraformResourceHints: exportIdentityTagNamespaceHints},
	},
	"oci_identity_tag_namespace": {
		{
			TerraformResourceHints: exportIdentityTagHints,
			DatasourceQueryParams: map[string]string{
				"tag_namespace_id": "id",
			},
		},
	},
}

var customAssociationIdentityCompartment = []tf_export.TerraformResourceAssociation{
	{
		TerraformResourceHints: exportIdentityAvailabilityDomainHints,
	},
}
