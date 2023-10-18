---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_mapped_secret"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall_policy_mapped_secret"
description: |-
  Provides the Network Firewall Policy Mapped Secret resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall_policy_mapped_secret
This resource provides the Network Firewall Policy Mapped Secret resource in Oracle Cloud Infrastructure Network Firewall service.

Creates a new Mapped Secret for the Network Firewall Policy.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall_policy_mapped_secret" "test_network_firewall_policy_mapped_secret" {
	#Required
	name = var.network_firewall_policy_mapped_secret_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	source = var.network_firewall_policy_mapped_secret_source
	type = var.network_firewall_policy_mapped_secret_type
	vault_secret_id = oci_vault_secret.test_secret.id
	version_number = var.network_firewall_policy_mapped_secret_version_number
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Unique name to identify the group of urls to be used in the policy rules.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `source` - (Required) Source of the secrets, where the secrets are stored. The only accepted value is `OCI_VAULT`
* `type` - (Required) Type of the secrets mapped based on the policy.
	* `SSL_INBOUND_INSPECTION`: For Inbound inspection of SSL traffic.
	* `SSL_FORWARD_PROXY`: For forward proxy certificates for SSL inspection. 
* `vault_secret_id` - (Required) (Updatable) OCID for the Vault Secret to be used.
* `version_number` - (Required) (Updatable) Version number of the secret to be used.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `name` - Name of the secret.
* `parent_resource_id` - OCID of the Network Firewall Policy this Mapped Secret belongs to.
* `source` - Source of the secrets, where the secrets are stored.
* `type` - Type of the secrets mapped based on the policy.
	* `SSL_INBOUND_INSPECTION`: For Inbound inspection of SSL traffic.
	* `SSL_FORWARD_PROXY`: For forward proxy certificates for SSL inspection. 
* `vault_secret_id` - OCID for the Vault Secret to be used.
* `version_number` - Version number of the secret to be used.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall Policy Mapped Secret
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall Policy Mapped Secret
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall Policy Mapped Secret


## Import

NetworkFirewallPolicyMappedSecrets can be imported using the `name`, e.g.

```
$ terraform import oci_network_firewall_network_firewall_policy_mapped_secret.test_network_firewall_policy_mapped_secret "networkFirewallPolicies/{networkFirewallPolicyId}/mappedSecrets/{mappedSecretName}" 
```

