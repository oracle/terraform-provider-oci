---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_zone_stage_dnssec_key_version"
sidebar_current: "docs-oci-resource-dns-zone_stage_dnssec_key_version"
description: |-
  Provides the Zone Stage Dnssec Key Version resource in Oracle Cloud Infrastructure DNS service
---

# oci_dns_zone_stage_dnssec_key_version
This resource provides the Zone Stage Dnssec Key Version resource in Oracle Cloud Infrastructure DNS service.

Stages a new `DnssecKeyVersion` on the zone. Staging is a process that generates a new "successor" key version
that replaces an existing "predecessor" key version.
**Note:** A new key-signing key (KSK) version is inert until you update the parent zone DS records.

For more information, see the [DNSSEC](https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnssec.htm) documentation.


## Example Usage

```hcl
resource "oci_dns_zone_stage_dnssec_key_version" "test_zone_stage_dnssec_key_version" {
	#Required
	predecessor_dnssec_key_version_uuid = var.zone_stage_dnssec_key_version_predecessor_dnssec_key_version_uuid
	zone_id = oci_dns_zone.test_zone.id

	#Optional
	scope = var.zone_stage_dnssec_key_version_scope
}
```

## Argument Reference

The following arguments are supported:

* `predecessor_dnssec_key_version_uuid` - (Required) The UUID of the `DnssecKeyVersion` for which a new successor should be generated. 
* `scope` - (Optional) Specifies to operate only on resources that have a matching DNS scope. 
* `zone_id` - (Required) The OCID of the target zone.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Zone Stage Dnssec Key Version
	* `update` - (Defaults to 20 minutes), when updating the Zone Stage Dnssec Key Version
	* `delete` - (Defaults to 20 minutes), when destroying the Zone Stage Dnssec Key Version


## Import

Import is not supported for this resource.

