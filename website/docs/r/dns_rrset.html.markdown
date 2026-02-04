---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_rrset"
sidebar_current: "docs-oci-resource-dns-rrset"
description: |-
  Provides the Rrset resource in Oracle Cloud Infrastructure DNS service
---

# oci_dns_rrset
This resource provides the Rrset resource in Oracle Cloud Infrastructure DNS service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/dns/latest/Rrset

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/dns

  Updates records in the specified RRSet.

When accessing a private zone by name, the view_id query parameter is required.

## Example Usage

```hcl
resource "oci_dns_rrset" "test_rrset" {
	#Required
	domain = var.rrset_domain
	rtype = var.rrset_rtype
	zone_name_or_id = oci_dns_zone.test_zone.id

	#Optional
	items {
		#Required
		domain = var.rrset_items_domain
		rdata = var.rrset_items_rdata
		rtype = var.rrset_items_rtype
		ttl = var.rrset_items_ttl
	}
	view_id = oci_dns_view.test_view.id
}
```

## Behavior

- Create returns HTTP 409 (Conflict) if the target RRSet already contains records. Use Update to modify an existing RRSet, or Delete to remove records.

## Argument Reference

The following arguments are supported:


* `domain` - (Required) The target fully-qualified domain name (FQDN) within the target zone.
* `items` - (Optional) (Updatable) 
	* `domain` - (Required) The fully qualified domain name where the record can be located. 
	* `rdata` - (Required) (Updatable) The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see [Supported DNS Resource Record Types](https://docs.cloud.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm)  
	* `rtype` - (Required) The type of DNS record, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
	* `ttl` - (Required) (Updatable) The Time To Live for the record, in seconds. Using a TTL lower than 30 seconds is not recommended. 
* `rtype` - (Required) The type of the target RRSet within the target zone.
* `view_id` - (Optional) The OCID of the view the zone is associated with. Required when accessing a private zone by name.
* `zone_name_or_id` - (Required) The name or OCID of the target zone.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - 
	* `domain` - The fully qualified domain name where the record can be located. 
	* `is_protected` - A Boolean flag indicating whether or not parts of the record are unable to be explicitly managed. 
	* `rdata` - The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see [Supported DNS Resource Record Types](https://docs.cloud.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm) 
	* `record_hash` - A unique identifier for the record within its zone. 
	* `rrset_version` - The latest version of the record's zone in which its RRSet differs from the preceding version. 
	* `rtype` - The type of DNS record, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
	* `ttl` - The Time To Live for the record, in seconds. Using a TTL lower than 30 seconds is not recommended. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Rrset
	* `update` - (Defaults to 20 minutes), when updating the Rrset
	* `delete` - (Defaults to 20 minutes), when destroying the Rrset


## Import

For legacy Rrsets that were created without using `scope`, these Rrsets can be imported using the `id`, e.g.

```
$ terraform import oci_dns_rrset.test_rrset "zoneNameOrId/{zoneNameOrId}/domain/{domain}/rtype/{rtype}" 
```

Note: Legacy RRSet IDs that include scope/viewId remain accepted for import for backward compatibility; however, scope is no longer a supported argument on this resource.
