package provider

import (
	tf_identity "github.com/terraform-providers/terraform-provider-oci/oci/service/identity"
)

func init() {
	RegisterDatasource("oci_identity_compartment", tf_identity.IdentityCompartmentDataSource())
	RegisterDatasource("oci_identity_api_keys", tf_identity.IdentityApiKeysDataSource())
	RegisterDatasource("oci_identity_auth_tokens", tf_identity.IdentityAuthTokensDataSource())
	RegisterDatasource("oci_identity_authentication_policy", tf_identity.IdentityAuthenticationPolicyDataSource())
	RegisterDatasource("oci_identity_availability_domain", tf_identity.IdentityAvailabilityDomainDataSource())
	RegisterDatasource("oci_identity_availability_domains", tf_identity.IdentityAvailabilityDomainsDataSource())
	RegisterDatasource("oci_identity_compartments", tf_identity.IdentityCompartmentsDataSource())
	RegisterDatasource("oci_identity_cost_tracking_tags", tf_identity.IdentityCostTrackingTagsDataSource())
	RegisterDatasource("oci_identity_customer_secret_keys", tf_identity.IdentityCustomerSecretKeysDataSource())
	RegisterDatasource("oci_identity_dynamic_groups", tf_identity.IdentityDynamicGroupsDataSource())
	RegisterDatasource("oci_identity_fault_domains", tf_identity.IdentityFaultDomainsDataSource())
	RegisterDatasource("oci_identity_group", tf_identity.IdentityGroupDataSource())
	RegisterDatasource("oci_identity_groups", tf_identity.IdentityGroupsDataSource())
	RegisterDatasource("oci_identity_identity_provider_groups", tf_identity.IdentityIdentityProviderGroupsDataSource())
	RegisterDatasource("oci_identity_identity_providers", tf_identity.IdentityIdentityProvidersDataSource())
	RegisterDatasource("oci_identity_idp_group_mappings", tf_identity.IdentityIdpGroupMappingsDataSource())
	RegisterDatasource("oci_identity_network_source", tf_identity.IdentityNetworkSourceDataSource())
	RegisterDatasource("oci_identity_network_sources", tf_identity.IdentityNetworkSourcesDataSource())
	RegisterDatasource("oci_identity_policies", tf_identity.IdentityPoliciesDataSource())
	RegisterDatasource("oci_identity_region_subscriptions", tf_identity.IdentityRegionSubscriptionsDataSource())
	RegisterDatasource("oci_identity_regions", tf_identity.IdentityRegionsDataSource())
	RegisterDatasource("oci_identity_smtp_credentials", tf_identity.IdentitySmtpCredentialsDataSource())
	RegisterDatasource("oci_identity_swift_passwords", tf_identity.IdentitySwiftPasswordsDataSource())
	RegisterDatasource("oci_identity_tag", tf_identity.IdentityTagDataSource())
	RegisterDatasource("oci_identity_tag_default", tf_identity.IdentityTagDefaultDataSource())
	RegisterDatasource("oci_identity_tag_defaults", tf_identity.IdentityTagDefaultsDataSource())
	RegisterDatasource("oci_identity_tag_namespaces", tf_identity.IdentityTagNamespacesDataSource())
	RegisterDatasource("oci_identity_tags", tf_identity.IdentityTagsDataSource())
	RegisterDatasource("oci_identity_tenancy", tf_identity.IdentityTenancyDataSource())
	RegisterDatasource("oci_identity_ui_password", tf_identity.IdentityUiPasswordDataSource())
	RegisterDatasource("oci_identity_user", tf_identity.IdentityUserDataSource())
	RegisterDatasource("oci_identity_user_group_memberships", tf_identity.IdentityUserGroupMembershipsDataSource())

}
