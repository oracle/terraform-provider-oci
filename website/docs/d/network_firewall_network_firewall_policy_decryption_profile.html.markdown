---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_decryption_profile"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_decryption_profile"
description: |-
  Provides details about a specific Network Firewall Policy Decryption Profile in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_decryption_profile
This data source provides details about a specific Network Firewall Policy Decryption Profile resource in Oracle Cloud Infrastructure Network Firewall service.

Get Decryption Profile by the given name in the context of network firewall policy.

## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_decryption_profile" "test_network_firewall_policy_decryption_profile" {
	#Required
	decryption_profile_name = var.oci_network_firewall_network_firewall_policy_decryption_profile_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `decryption_profile_name` - (Required) Unique identifier for Decryption Profiles in the scope of Network Firewall Policy.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

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

