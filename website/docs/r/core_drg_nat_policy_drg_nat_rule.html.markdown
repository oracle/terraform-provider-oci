---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_nat_policy_drg_nat_rule"
sidebar_current: "docs-oci-resource-core-drg_nat_policy_drg_nat_rule"
description: |-
  Provides the DrgNatRule resource as part of a DrgNatPolicy in Oracle Cloud Infrastructure Core service.
---

# oci_core_drg_nat_policy_drg_nat_rule

This resource creates and manages an individual DrgNatRule for a DrgNatPolicy in the Oracle Cloud Infrastructure Core service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iaas/latest/DrgNatRule

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Creates a new DrgNatRule for a DrgNatPolicy. Assign the DrgNatPolicy to a DRG attachment
using the `UpdateDrgAttachment` or `CreateDrgAttachment` operations.

## Example Usage

```hcl

resource "oci_core_drg_nat_policy_drg_nat_rule" "test_drg_nat_rule" {
  drg_nat_policy_id      = oci_core_drg_nat_policy.test_drg_nat_policy.id
  drg_nat_rule_priority  = 100
  original_source        = "10.0.0.0/24"
  translated_source      = "192.168.0.0/24"
  original_destination   = "172.16.1.0/24"
  translated_destination = "10.1.1.0/24"
}
```

## Argument Reference

The following arguments are supported:

* `drg_nat_policy_id` - (Required) The DrgNatPolicy's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `drg_nat_rule_priority` – (Required) (Updatable) The priority associated with each DrgNatRule.
* `original_source` – (Optional) (Updatable) Represents the range of IP addresses to match against when routing traffic. Original CIDR range for Source NAT.
* `translated_source` – (Optional) (Updatable) Represents the range of IP addresses to match against when routing traffic. Translated CIDR range for Source NAT.
* `original_destination` – (Optional) (Updatable) Represents the range of IP addresses to match against when routing traffic. Original CIDR range for Destination NAT.
* `translated_destination` – (Optional) (Updatable) Represents the range of IP addresses to match against when routing traffic. Translated CIDR range for Destination NAT.

**IMPORTANT**  
Changing certain properties that do not support update will force destruction and recreation of the resource with new property values.

## Attributes Reference

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


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` – (Defaults to 20 minutes), when creating the DrgNatRule
* `update` – (Defaults to 20 minutes), when updating the DrgNatRule
* `delete` – (Defaults to 20 minutes), when destroying the DrgNatRule