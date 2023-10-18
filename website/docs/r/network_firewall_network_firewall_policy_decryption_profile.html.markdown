---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_decryption_profile"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall_policy_decryption_profile"
description: |-
  Provides the Network Firewall Policy Decryption Profile resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall_policy_decryption_profile
This resource provides the Network Firewall Policy Decryption Profile resource in Oracle Cloud Infrastructure Network Firewall service.

Creates a new Decryption Profile for the Network Firewall Policy.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall_policy_decryption_profile" "test_network_firewall_policy_decryption_profile" {
	#Required
	name = var.network_firewall_policy_decryption_profile_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	type = var.network_firewall_policy_decryption_profile_type

	#Optional
	are_certificate_extensions_restricted = var.network_firewall_policy_decryption_profile_are_certificate_extensions_restricted
	is_auto_include_alt_name = var.network_firewall_policy_decryption_profile_is_auto_include_alt_name
	is_expired_certificate_blocked = var.network_firewall_policy_decryption_profile_is_expired_certificate_blocked
	is_out_of_capacity_blocked = var.network_firewall_policy_decryption_profile_is_out_of_capacity_blocked
	is_revocation_status_timeout_blocked = var.network_firewall_policy_decryption_profile_is_revocation_status_timeout_blocked
	is_unknown_revocation_status_blocked = var.network_firewall_policy_decryption_profile_is_unknown_revocation_status_blocked
	is_unsupported_cipher_blocked = var.network_firewall_policy_decryption_profile_is_unsupported_cipher_blocked
	is_unsupported_version_blocked = var.network_firewall_policy_decryption_profile_is_unsupported_version_blocked
	is_untrusted_issuer_blocked = var.network_firewall_policy_decryption_profile_is_untrusted_issuer_blocked
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the decryption profile.
* `type` - (Required) Describes the type of decryption profile. The accepted values are - * SSL_FORWARD_PROXY * SSL_INBOUND_INSPECTION
* `is_out_of_capacity_blocked` - (Optional) (Updatable) Whether to block sessions if the firewall is temporarily unable to decrypt their traffic.
* `is_unsupported_cipher_blocked` - (Optional) (Updatable) Whether to block sessions if SSL cipher suite is not supported.
* `is_unsupported_version_blocked` - (Optional) (Updatable) Whether to block sessions if SSL version is not supported.
* `are_certificate_extensions_restricted` - (Applicable only when type=SSL_FORWARD_PROXY) (Updatable) Whether to block sessions if the server's certificate uses extensions other than key usage and/or extended key usage.
* `is_auto_include_alt_name` - (Applicable only when type=SSL_FORWARD_PROXY) (Updatable) Whether to automatically append SAN to impersonating certificate if server certificate is missing SAN.
* `is_expired_certificate_blocked` - (Applicable only when type=SSL_FORWARD_PROXY) (Updatable) Whether to block sessions if server's certificate is expired.
* `is_revocation_status_timeout_blocked` - (Applicable only when type=SSL_FORWARD_PROXY) (Updatable) Whether to block sessions if the revocation status check for server's certificate does not succeed within the maximum allowed time (defaulting to 5 seconds). 
* `is_unknown_revocation_status_blocked` - (Applicable only when type=SSL_FORWARD_PROXY) (Updatable) Whether to block sessions if the revocation status check for server's certificate results in "unknown".
* `is_untrusted_issuer_blocked` - (Applicable only when type=SSL_FORWARD_PROXY) (Updatable) Whether to block sessions if server's certificate is issued by an untrusted certificate authority (CA).
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `name` - Unique Name of the decryption profile.
* `type` - Describes the type of decryption profile.
* `parent_resource_id` - OCID of the Network Firewall Policy this decryption profile belongs to.
* `are_certificate_extensions_restricted` - Whether to block sessions if the server's certificate uses extensions other than key usage and/or extended key usage.
* `is_auto_include_alt_name` - Whether to automatically append SAN to impersonating certificate if server certificate is missing SAN.
* `is_expired_certificate_blocked` - Whether to block sessions if server's certificate is expired.
* `is_out_of_capacity_blocked` - Whether to block sessions if the firewall is temporarily unable to decrypt their traffic.
* `is_revocation_status_timeout_blocked` - Whether to block sessions if the revocation status check for server's certificate does not succeed within the maximum allowed time (defaulting to 5 seconds). 
* `is_unknown_revocation_status_blocked` - Whether to block sessions if the revocation status check for server's certificate results in "unknown".
* `is_unsupported_cipher_blocked` - Whether to block sessions if SSL cipher suite is not supported.
* `is_unsupported_version_blocked` - Whether to block sessions if SSL version is not supported.
* `is_untrusted_issuer_blocked` - Whether to block sessions if server's certificate is issued by an untrusted certificate authority (CA).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall Policy Decryption Profile
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall Policy Decryption Profile
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall Policy Decryption Profile


## Import

NetworkFirewallPolicyDecryptionProfiles can be imported using the `name`, e.g.

```
$ terraform import oci_network_firewall_network_firewall_policy_decryption_profile.test_network_firewall_policy_decryption_profile "networkFirewallPolicies/{networkFirewallPolicyId}/decryptionProfiles/{decryptionProfileName}" 
```

