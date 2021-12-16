package provider

import (
	"github.com/terraform-providers/terraform-provider-oci/internal/service/audit"
	tf_identity "github.com/terraform-providers/terraform-provider-oci/internal/service/identity"
	tf_kms "github.com/terraform-providers/terraform-provider-oci/internal/service/kms"
)

func init() {
	RegisterResource("oci_identity_dynamic_group", tf_identity.IdentityDynamicGroupResource())
	RegisterResource("oci_identity_smtp_credential", tf_identity.IdentitySmtpCredentialResource())
	RegisterResource("oci_identity_authentication_policy", tf_identity.IdentityAuthenticationPolicyResource())
	RegisterResource("oci_identity_idp_group_mapping", tf_identity.IdentityIdpGroupMappingResource())
	RegisterResource("oci_identity_ui_password", tf_identity.IdentityUiPasswordResource())
	RegisterResource("oci_identity_user_capabilities_management", tf_identity.IdentityUserCapabilitiesManagementResource())
	RegisterResource("oci_identity_tag_default", tf_identity.IdentityTagDefaultResource())
	RegisterResource("oci_identity_network_source", tf_identity.IdentityNetworkSourceResource())
	RegisterResource("oci_identity_identity_provider", tf_identity.IdentityIdentityProviderResource())
	RegisterResource("oci_identity_data_plane_generate_scoped_access_token", tf_identity.IdentityDataPlaneGenerateScopedAccessTokenResource())
	RegisterResource("oci_identity_swift_password", tf_identity.IdentitySwiftPasswordResource())
	RegisterResource("oci_identity_user_group_membership", tf_identity.IdentityUserGroupMembershipResource())
	RegisterResource("oci_identity_db_credential", tf_identity.IdentityDbCredentialResource())
	RegisterResource("oci_identity_domain_replication_to_region", tf_identity.IdentityDomainReplicationToRegionResource())
	RegisterResource("oci_identity_domain", tf_identity.IdentityDomainResource())
	RegisterResource("oci_identity_auth_token", tf_identity.IdentityAuthTokenResource())
	RegisterResource("oci_identity_tag", tf_identity.IdentityTagResource())
	RegisterResource("oci_identity_api_key", tf_identity.IdentityApiKeyResource())
	RegisterResource("oci_identity_group", tf_identity.IdentityGroupResource())
	RegisterResource("oci_identity_compartment", tf_identity.IdentityCompartmentResource())
	RegisterResource("oci_identity_tag_namespace", tf_identity.IdentityTagNamespaceResource())
	RegisterResource("oci_identity_policy", tf_identity.IdentityPolicyResource())
	RegisterResource("oci_identity_user", tf_identity.IdentityUserResource())
	RegisterResource("oci_identity_customer_secret_key", tf_identity.IdentityCustomerSecretKeyResource())
	RegisterResource("oci_audit_configuration", audit.AuditConfigurationResource())

	// kms service
	RegisterResource("oci_kms_verify", tf_kms.KmsVerifyResource())
	RegisterResource("oci_kms_vault_replication", tf_kms.KmsVaultReplicationResource())
	RegisterResource("oci_kms_generated_key", tf_kms.KmsGeneratedKeyResource())
	RegisterResource("oci_kms_key", tf_kms.KmsKeyResource())
	RegisterResource("oci_kms_key_version", tf_kms.KmsKeyVersionResource())
	RegisterResource("oci_kms_encrypted_data", tf_kms.KmsEncryptedDataResource())
	RegisterResource("oci_kms_sign", tf_kms.KmsSignResource())
	RegisterResource("oci_kms_vault", tf_kms.KmsVaultResource())
}
