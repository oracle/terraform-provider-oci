---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_decryption_profiles"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_decryption_profiles"
description: |-
  Provides the list of Network Firewall Policy Decryption Profiles in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_decryption_profiles
This data source provides the list of Network Firewall Policy Decryption Profiles in Oracle Cloud Infrastructure Network Firewall service.

Returns a list of Decryption Profile for the Network Firewall Policy.


## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_decryption_profiles" "test_network_firewall_policy_decryption_profiles" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

	#Optional
	display_name = var.network_firewall_policy_decryption_profile_display_name
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

The following attributes are exported:

* `decryption_profile_summary_collection` - The list of decryption_profile_summary_collection.

### NetworkFirewallPolicyDecryptionProfile Reference

The following attributes are exported:

* `are_certificate_extensions_restricted` - Whether to block sessions if the server's certificate uses extensions other than key usage and/or extended key usage.
* `is_auto_include_alt_name` - Whether to automatically append SAN to impersonating certificate if server certificate is missing SAN.
* `is_expired_certificate_blocked` - Whether to block sessions if server's certificate is expired.
* `is_out_of_capacity_blocked` - Whether to block sessions if the firewall is temporarily unable to decrypt their traffic.
* `is_revocation_status_timeout_blocked` - Whether to block sessions if the revocation status check for server's certificate does not succeed within the maximum allowed time (defaulting to 5 seconds). 
* `is_unknown_revocation_status_blocked` - Whether to block sessions if the revocation status check for server's certificate results in "unknown".
* `is_unsupported_cipher_blocked` - Whether to block sessions if SSL cipher suite is not supported.
* `is_unsupported_version_blocked` - Whether to block sessions if SSL version is not supported.
* `is_untrusted_issuer_blocked` - Whether to block sessions if server's certificate is issued by an untrusted certificate authority (CA).
* `name` - Unique Name of the decryption profile.
* `parent_resource_id` - OCID of the Network Firewall Policy this decryption profile belongs to.
* `type` - Describes the type of Decryption Profile SslForwardProxy or SslInboundInspection.

