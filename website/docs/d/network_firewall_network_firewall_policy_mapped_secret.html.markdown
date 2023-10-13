---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_mapped_secret"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_mapped_secret"
description: |-
  Provides details about a specific Network Firewall Policy Mapped Secret in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_mapped_secret
This data source provides details about a specific Network Firewall Policy Mapped Secret resource in Oracle Cloud Infrastructure Network Firewall service.

Get Mapped Secret by the given name in the context of network firewall policy.

## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_mapped_secret" "test_network_firewall_policy_mapped_secret" {
	#Required
	mapped_secret_name = var.oci_network_firewall_network_firewall_policy_mapped_secret_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `mapped_secret_name` - (Required) Unique identifier for Mapped Secrets in the scope of Network Firewall Policy.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


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

