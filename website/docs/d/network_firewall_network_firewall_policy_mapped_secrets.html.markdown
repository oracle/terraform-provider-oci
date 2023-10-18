---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_mapped_secrets"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_mapped_secrets"
description: |-
  Provides the list of Network Firewall Policy Mapped Secrets in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_mapped_secrets
This data source provides the list of Network Firewall Policy Mapped Secrets in Oracle Cloud Infrastructure Network Firewall service.

Returns a list of Mapped Secret for the Network Firewall Policy.


## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_mapped_secrets" "test_network_firewall_policy_mapped_secrets" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

	#Optional
	display_name = var.network_firewall_policy_mapped_secret_display_name
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

The following attributes are exported:

* `mapped_secret_summary_collection` - The list of mapped_secret_summary_collection.

### NetworkFirewallPolicyMappedSecret Reference

The following attributes are exported:

* `name` - Name of the secret.
* `parent_resource_id` - OCID of the Network Firewall Policy this Mapped Secret belongs to.
* `source` - Source of the secrets, where the secrets are stored.
* `type` - Type of the secrets mapped based on the policy.
	* `SSL_INBOUND_INSPECTION`: For Inbound inspection of SSL traffic.
	* `SSL_FORWARD_PROXY`: For forward proxy certificates for SSL inspection. 
* `vault_secret_id` - OCID for the Vault Secret to be used.
* `version_number` - Version number of the secret to be used.

