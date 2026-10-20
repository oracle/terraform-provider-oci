---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_nat_policy_drg_nat_rules"
sidebar_current: "docs-oci-datasource-core-drg_nat_policy_drg_nat_rules"
description: |-
  Provides the list of DrgNatRules that are associated with a DrgNatPolicy  in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_drg_nat_policy_drg_nat_rules
This data source provides the list of DrgNatRules that are associated with a DrgNatPolicy in Oracle Cloud Infrastructure Core service.

Lists the rules for the specified DrgNatPolicy

## Example Usage

```hcl
data "oci_core_drg_nat_policy_drg_nat_rules" "test_drg_nat_policy_drg_nat_rules" {
	#Required
	drg_nat_policy_id = oci_core_drg_nat_policy.test_drg_nat_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `drg_nat_policy_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DrgNatPolicy.


## Attributes Reference

The following attributes are exported:

* `drg_nat_rules` - The list of drg_nat_rules.

### DrgNatPolicyDrgNatRule Reference

The following attributes are exported:

* `drg_nat_policy_id` - The DrgNatPolicy's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `drg_nat_rule_priority` - The priority associated with each DrgNatRule.
* `id` - The Oracle-assigned ID of the DrgNatRule.
* `original_source` - Represents the range of IP addresses to match against when routing traffic. Original CIDR range for Source NAT.

  Potential values:
	* An IPv4 address range in CIDR notation. For example: `192.168.1.0/24`.
* `translated_source` - Represents the range of IP addresses to match against when routing traffic. Translated CIDR range for Source NAT.

  Potential values:
	* An IPv4 address range in CIDR notation. For example: `192.168.1.0/24`.
* `original_destination` - Represents the range of IP addresses to match against when routing traffic. Original CIDR range for Destination NAT.

  Potential values:
	* An IPv4 address range in CIDR notation. For example: `192.168.1.0/24`.
* `translated_destination` - Represents the range of IP addresses to match against when routing traffic. Translated CIDR range for Destination NAT.

  Potential values:
	* An IPv4 address range in CIDR notation. For example: `192.168.1.0/24`.
