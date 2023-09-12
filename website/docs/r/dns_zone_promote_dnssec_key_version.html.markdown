---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_zone_promote_dnssec_key_version"
sidebar_current: "docs-oci-resource-dns-zone_promote_dnssec_key_version"
description: |-
  Provides the Zone Promote Dnssec Key Version resource in Oracle Cloud Infrastructure DNS service
---

# oci_dns_zone_promote_dnssec_key_version
This resource provides the Zone Promote Dnssec Key Version resource in Oracle Cloud Infrastructure DNS service.

Promotes a specified `DnssecKeyVersion` on the zone.

If the `DnssecKeyVersion` identified in the request body is a key signing key (KSK) that is replacing
another `DnssecKeyVersion`, then the old `DnssecKeyVersion` is scheduled for removal from the zone.

For key signing keys (KSKs), you must create the DS record with the new key information **before** promoting
the new key to establish a chain of trust. To avoid a service disruption, remove the old DS record as soon
as its TTL (time to live) expires.

For more information, see [DNSSEC](https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnssec.htm).


## Example Usage

```hcl
resource "oci_dns_zone_promote_dnssec_key_version" "test_zone_promote_dnssec_key_version" {
	#Required
	dnssec_key_version_uuid = var.zone_promote_dnssec_key_version_dnssec_key_version_uuid
	zone_id = oci_dns_zone.test_zone.id

	#Optional
	scope = var.zone_promote_dnssec_key_version_scope
}
```

## Argument Reference

The following arguments are supported:

* `dnssec_key_version_uuid` - (Required) The UUID of the `DnssecKeyVersion` that is being promoted. 
* `scope` - (Optional) Specifies to operate only on resources that have a matching DNS scope. 
* `zone_id` - (Required) The OCID of the target zone.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Zone Promote Dnssec Key Version
	* `update` - (Defaults to 20 minutes), when updating the Zone Promote Dnssec Key Version
	* `delete` - (Defaults to 20 minutes), when destroying the Zone Promote Dnssec Key Version


## Import

Import is not supported for this resource.

