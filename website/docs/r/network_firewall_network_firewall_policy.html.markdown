---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall_policy"
description: |-
  Provides the Network Firewall Policy resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall_policy
This resource provides the Network Firewall Policy resource in Oracle Cloud Infrastructure Network Firewall service.

Creates a new Network Firewall Policy.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall_policy" "test_network_firewall_policy" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	application_lists = var.network_firewall_policy_application_lists
	decryption_profiles {
		#Required
		is_out_of_capacity_blocked = var.network_firewall_policy_decryption_profiles_is_out_of_capacity_blocked
		is_unsupported_cipher_blocked = var.network_firewall_policy_decryption_profiles_is_unsupported_cipher_blocked
		is_unsupported_version_blocked = var.network_firewall_policy_decryption_profiles_is_unsupported_version_blocked
		type = var.network_firewall_policy_decryption_profiles_type

		#Optional
		are_certificate_extensions_restricted = var.network_firewall_policy_decryption_profiles_are_certificate_extensions_restricted
		is_auto_include_alt_name = var.network_firewall_policy_decryption_profiles_is_auto_include_alt_name
		is_expired_certificate_blocked = var.network_firewall_policy_decryption_profiles_is_expired_certificate_blocked
		is_revocation_status_timeout_blocked = var.network_firewall_policy_decryption_profiles_is_revocation_status_timeout_blocked
		is_unknown_revocation_status_blocked = var.network_firewall_policy_decryption_profiles_is_unknown_revocation_status_blocked
		is_untrusted_issuer_blocked = var.network_firewall_policy_decryption_profiles_is_untrusted_issuer_blocked
	}
	decryption_rules {
		#Required
		action = var.network_firewall_policy_decryption_rules_action
		condition {

			#Optional
			destinations = var.network_firewall_policy_decryption_rules_condition_destinations
			sources = var.network_firewall_policy_decryption_rules_condition_sources
		}
		name = var.network_firewall_policy_decryption_rules_name

		#Optional
		decryption_profile = var.network_firewall_policy_decryption_rules_decryption_profile
		secret = var.network_firewall_policy_decryption_rules_secret
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.network_firewall_policy_display_name
	freeform_tags = {"bar-key"= "value"}
	ip_address_lists = var.network_firewall_policy_ip_address_lists
	mapped_secrets {
		#Required
		source = var.network_firewall_policy_mapped_secrets_source
		type = var.network_firewall_policy_mapped_secrets_type
		vault_secret_id = oci_vault_secret.test_secret.id
		version_number = var.network_firewall_policy_mapped_secrets_version_number
	}
	security_rules {
		#Required
		action = var.network_firewall_policy_security_rules_action
		condition {

			#Optional
			applications = var.network_firewall_policy_security_rules_condition_applications
			destinations = var.network_firewall_policy_security_rules_condition_destinations
			sources = var.network_firewall_policy_security_rules_condition_sources
			urls = var.network_firewall_policy_security_rules_condition_urls
		}
		name = var.network_firewall_policy_security_rules_name

		#Optional
		inspection = var.network_firewall_policy_security_rules_inspection
	}
	url_lists = var.network_firewall_policy_url_lists
}
```

## Argument Reference

The following arguments are supported:

* `application_lists` - (Optional) (Updatable) Map defining application lists of the policy. The value of an entry is a list of "applications", each consisting of a protocol identifier (such as TCP, UDP, or ICMP) and protocol-specific parameters (such as a port range). The associated key is the identifier by which the application list is referenced. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the NetworkFirewall Policy.
* `decryption_profiles` - (Optional) (Updatable) Map defining decryption profiles of the policy. The value of an entry is a decryption profile. The associated key is the identifier by which the decryption profile is referenced. 
	* `are_certificate_extensions_restricted` - (Required when type=SSL_FORWARD_PROXY) (Updatable) Whether to block sessions if the server's certificate uses extensions other than key usage and/or extended key usage.
	* `is_auto_include_alt_name` - (Required when type=SSL_FORWARD_PROXY) (Updatable) Whether to automatically append SAN to impersonating certificate if server certificate is missing SAN.
	* `is_expired_certificate_blocked` - (Required when type=SSL_FORWARD_PROXY) (Updatable) Whether to block sessions if server's certificate is expired.
	* `is_out_of_capacity_blocked` - (Required) (Updatable) Whether to block sessions if the firewall is temporarily unable to decrypt their traffic.
	* `is_revocation_status_timeout_blocked` - (Required when type=SSL_FORWARD_PROXY) (Updatable) Whether to block sessions if the revocation status check for server's certificate does not succeed within the maximum allowed time (defaulting to 5 seconds). 
	* `is_unknown_revocation_status_blocked` - (Required when type=SSL_FORWARD_PROXY) (Updatable) Whether to block sessions if the revocation status check for server's certificate results in "unknown".
	* `is_unsupported_cipher_blocked` - (Required) (Updatable) Whether to block sessions if SSL cipher suite is not supported.
	* `is_unsupported_version_blocked` - (Required) (Updatable) Whether to block sessions if SSL version is not supported.
	* `is_untrusted_issuer_blocked` - (Required when type=SSL_FORWARD_PROXY) (Updatable) Whether to block sessions if server's certificate is issued by an untrusted certificate authority (CA).
	* `type` - (Required) (Updatable) Describes the type of Decryption Profile SslForwardProxy or SslInboundInspection.
* `decryption_rules` - (Optional) (Updatable) List of Decryption Rules defining the behavior of the policy. The first rule with a matching condition determines the action taken upon network traffic. 
	* `action` - (Required) (Updatable) Action:
		* NO_DECRYPT - Matching traffic is not decrypted.
		* DECRYPT - Matching traffic is decrypted with the specified `secret` according to the specified `decryptionProfile`. 
	* `condition` - (Required) (Updatable) Match criteria used in Decryption Rule used on the firewall policy rules.
		* `destinations` - (Optional) (Updatable) An array of IP address list names to be evaluated against the traffic destination address.
		* `sources` - (Optional) (Updatable) An array of IP address list names to be evaluated against the traffic source address.
	* `decryption_profile` - (Optional) (Updatable) The name of the decryption profile to use.
	* `name` - (Required) (Updatable) Name for the decryption rule, must be unique within the policy.
	* `secret` - (Optional) (Updatable) The name of a mapped secret. Its `type` must match that of the specified decryption profile.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) A user-friendly optional name for the firewall policy. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `ip_address_lists` - (Optional) (Updatable) Map defining IP address lists of the policy. The value of an entry is a list of IP addresses or prefixes in CIDR notation. The associated key is the identifier by which the IP address list is referenced. 
* `mapped_secrets` - (Optional) (Updatable) Map defining secrets of the policy. The value of an entry is a "mapped secret" consisting of a purpose and source. The associated key is the identifier by which the mapped secret is referenced. 
	* `source` - (Required) (Updatable) Source of the secrets, where the secrets are stored.
	* `type` - (Required) (Updatable) Type of the secrets mapped based on the policy.
		* `SSL_INBOUND_INSPECTION`: For Inbound inspection of SSL traffic.
		* `SSL_FORWARD_PROXY`: For forward proxy certificates for SSL inspection. 
	* `vault_secret_id` - (Required) (Updatable) OCID for the Vault Secret to be used.
	* `version_number` - (Required) (Updatable) Version number of the secret to be used.
* `security_rules` - (Optional) (Updatable) List of Security Rules defining the behavior of the policy. The first rule with a matching condition determines the action taken upon network traffic. 
	* `action` - (Required) (Updatable) Types of Action on the Traffic flow.
		* ALLOW - Allows the traffic.
		* DROP - Silently drops the traffic, e.g. without sending a TCP reset.
		* REJECT - Rejects the traffic, sending a TCP reset to client and/or server as applicable.
		* INSPECT - Inspects traffic for vulnerability as specified in `inspection`, which may result in rejection. 
	* `condition` - (Required) (Updatable) Criteria to evaluate against network traffic. A match occurs when at least one item in the array associated with each specified property corresponds with the relevant aspect of the traffic. 
		* `applications` - (Optional) (Updatable) An array of application list names to be evaluated against the traffic protocol and protocol-specific parameters.
		* `destinations` - (Optional) (Updatable) An array of IP address list names to be evaluated against the traffic destination address.
		* `sources` - (Optional) (Updatable) An array of IP address list names to be evaluated against the traffic source address.
		* `urls` - (Optional) (Updatable) An array of URL pattern list names to be evaluated against the HTTP(S) request target.
	* `inspection` - (Optional) (Updatable) Type of inspection to affect the Traffic flow. This is only applicable if action is INSPECT.
		* INTRUSION_DETECTION - Intrusion Detection.
		* INTRUSION_PREVENTION - Intrusion Detection and Prevention. Traffic classified as potentially malicious will be rejected as described in `type`. 
	* `name` - (Required) (Updatable) Name for the Security rule, must be unique within the policy.
* `url_lists` - (Optional) (Updatable) Map defining URL pattern lists of the policy. The value of an entry is a list of URL patterns. The associated key is the identifier by which the URL pattern list is referenced. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `application_lists` - Map defining application lists of the policy. The value of an entry is a list of "applications", each consisting of a protocol identifier (such as TCP, UDP, or ICMP) and protocol-specific parameters (such as a port range). The associated key is the identifier by which the application list is referenced. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the NetworkFirewall Policy.
* `decryption_profiles` - Map defining decryption profiles of the policy. The value of an entry is a decryption profile. The associated key is the identifier by which the decryption profile is referenced. 
	* `are_certificate_extensions_restricted` - Whether to block sessions if the server's certificate uses extensions other than key usage and/or extended key usage.
	* `is_auto_include_alt_name` - Whether to automatically append SAN to impersonating certificate if server certificate is missing SAN.
	* `is_expired_certificate_blocked` - Whether to block sessions if server's certificate is expired.
	* `is_out_of_capacity_blocked` - Whether to block sessions if the firewall is temporarily unable to decrypt their traffic.
	* `is_revocation_status_timeout_blocked` - Whether to block sessions if the revocation status check for server's certificate does not succeed within the maximum allowed time (defaulting to 5 seconds). 
	* `is_unknown_revocation_status_blocked` - Whether to block sessions if the revocation status check for server's certificate results in "unknown".
	* `is_unsupported_cipher_blocked` - Whether to block sessions if SSL cipher suite is not supported.
	* `is_unsupported_version_blocked` - Whether to block sessions if SSL version is not supported.
	* `is_untrusted_issuer_blocked` - Whether to block sessions if server's certificate is issued by an untrusted certificate authority (CA).
	* `type` - Describes the type of Decryption Profile SslForwardProxy or SslInboundInspection.
* `decryption_rules` - List of Decryption Rules defining the behavior of the policy. The first rule with a matching condition determines the action taken upon network traffic. 
	* `action` - Action:
		* NO_DECRYPT - Matching traffic is not decrypted.
		* DECRYPT - Matching traffic is decrypted with the specified `secret` according to the specified `decryptionProfile`. 
	* `condition` - Match criteria used in Decryption Rule used on the firewall policy rules.
		* `destinations` - An array of IP address list names to be evaluated against the traffic destination address.
		* `sources` - An array of IP address list names to be evaluated against the traffic source address.
	* `decryption_profile` - The name of the decryption profile to use.
	* `name` - Name for the decryption rule, must be unique within the policy.
	* `secret` - The name of a mapped secret. Its `type` must match that of the specified decryption profile.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly optional name for the firewall policy. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource - Network Firewall Policy.
* `ip_address_lists` - Map defining IP address lists of the policy. The value of an entry is a list of IP addresses or prefixes in CIDR notation. The associated key is the identifier by which the IP address list is referenced. 
* `is_firewall_attached` - To determine if any Network Firewall is associated with this Network Firewall Policy. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `mapped_secrets` - Map defining secrets of the policy. The value of an entry is a "mapped secret" consisting of a purpose and source. The associated key is the identifier by which the mapped secret is referenced. 
	* `source` - Source of the secrets, where the secrets are stored.
	* `type` - Type of the secrets mapped based on the policy.
		* `SSL_INBOUND_INSPECTION`: For Inbound inspection of SSL traffic.
		* `SSL_FORWARD_PROXY`: For forward proxy certificates for SSL inspection. 
	* `vault_secret_id` - OCID for the Vault Secret to be used.
	* `version_number` - Version number of the secret to be used.
* `security_rules` - List of Security Rules defining the behavior of the policy. The first rule with a matching condition determines the action taken upon network traffic. 
	* `action` - Types of Action on the Traffic flow.
		* ALLOW - Allows the traffic.
		* DROP - Silently drops the traffic, e.g. without sending a TCP reset.
		* REJECT - Rejects the traffic, sending a TCP reset to client and/or server as applicable.
		* INSPECT - Inspects traffic for vulnerability as specified in `inspection`, which may result in rejection. 
	* `condition` - Criteria to evaluate against network traffic. A match occurs when at least one item in the array associated with each specified property corresponds with the relevant aspect of the traffic. 
		* `applications` - An array of application list names to be evaluated against the traffic protocol and protocol-specific parameters.
		* `destinations` - An array of IP address list names to be evaluated against the traffic destination address.
		* `sources` - An array of IP address list names to be evaluated against the traffic source address.
		* `urls` - An array of URL pattern list names to be evaluated against the HTTP(S) request target.
	* `inspection` - Type of inspection to affect the Traffic flow. This is only applicable if action is INSPECT.
		* INTRUSION_DETECTION - Intrusion Detection.
		* INTRUSION_PREVENTION - Intrusion Detection and Prevention. Traffic classified as potentially malicious will be rejected as described in `type`. 
	* `name` - Name for the Security rule, must be unique within the policy.
* `state` - The current state of the Network Firewall Policy.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time instant at which the Network Firewall Policy was created in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The time instant at which the Network Firewall Policy was updated in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `url_lists` - Map defining URL pattern lists of the policy. The value of an entry is a list of URL patterns. The associated key is the identifier by which the URL pattern list is referenced. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall Policy
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall Policy


## Import

NetworkFirewallPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_network_firewall_network_firewall_policy.test_network_firewall_policy "id"
```

