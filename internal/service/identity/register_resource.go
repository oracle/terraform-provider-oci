// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_identity_api_key", IdentityApiKeyResource())
	tfresource.RegisterResource("oci_identity_auth_token", IdentityAuthTokenResource())
	tfresource.RegisterResource("oci_identity_authentication_policy", IdentityAuthenticationPolicyResource())
	tfresource.RegisterResource("oci_identity_compartment", IdentityCompartmentResource())
	tfresource.RegisterResource("oci_identity_customer_secret_key", IdentityCustomerSecretKeyResource())
	tfresource.RegisterResource("oci_identity_db_credential", IdentityDbCredentialResource())
	tfresource.RegisterResource("oci_identity_domain", IdentityDomainResource())
	tfresource.RegisterResource("oci_identity_domain_replication_to_region", IdentityDomainReplicationToRegionResource())
	tfresource.RegisterResource("oci_identity_dynamic_group", IdentityDynamicGroupResource())
	tfresource.RegisterResource("oci_identity_group", IdentityGroupResource())
	tfresource.RegisterResource("oci_identity_identity_provider", IdentityIdentityProviderResource())
	tfresource.RegisterResource("oci_identity_idp_group_mapping", IdentityIdpGroupMappingResource())
	tfresource.RegisterResource("oci_identity_import_standard_tags_management", IdentityImportStandardTagsManagementResource())
	tfresource.RegisterResource("oci_identity_network_source", IdentityNetworkSourceResource())
	tfresource.RegisterResource("oci_identity_policy", IdentityPolicyResource())
	tfresource.RegisterResource("oci_identity_smtp_credential", IdentitySmtpCredentialResource())
	tfresource.RegisterResource("oci_identity_tag", IdentityTagResource())
	tfresource.RegisterResource("oci_identity_tag_default", IdentityTagDefaultResource())
	tfresource.RegisterResource("oci_identity_tag_namespace", IdentityTagNamespaceResource())
	tfresource.RegisterResource("oci_identity_ui_password", IdentityUiPasswordResource())
	tfresource.RegisterResource("oci_identity_user", IdentityUserResource())
	tfresource.RegisterResource("oci_identity_user_group_membership", IdentityUserGroupMembershipResource())
	tfresource.RegisterResource("oci_identity_user_capabilities_management", IdentityUserCapabilitiesManagementResource())
}
