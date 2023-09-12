---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_record"
sidebar_current: "docs-oci-resource-dns-record"
description: |-
  Provides the Record resource in Oracle Cloud Infrastructure DNS service
---

# oci_dns_record

**Deprecated. Use [oci_dns_rrset](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/dns_rrset) instead.**

This resource provides the Record resource in Oracle Cloud Infrastructure DNS service.

  Updates a collection of records in the specified zone.

You can update one record or all records for the specified zone depending on the changes provided in the
request body. You can also add or remove records using this function. When the zone name is provided as
a path parameter and `PRIVATE` is used for the scope query parameter then the viewId query parameter is
required.


## Example Usage

```hcl
resource "oci_dns_record" "test_record" {
	#Required
	zone_name_or_id = oci_dns_zone_name_or.test_zone_name_or.id
	domain = var.record_items_domain
	rtype = var.record_items_rtype

	#Optional
	rdata = var.record_items_rdata
	ttl = var.record_items_ttl
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment the resource belongs to. If supplied, it must match the Zone's compartment ocid. 
* `domain` - (Required) The fully qualified domain name where the record can be located. Domain value is case insensitive. 
* `rdata` - (Optional) (Updatable) The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see [Supported DNS Resource Record Types](https://docs.cloud.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm) 
* `rtype` - (Required) The canonical name for the record's type, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
* `ttl` - (Optional) (Updatable) The Time To Live for the record, in seconds.
* `zone_name_or_id` - (Required) The name or OCID of the target zone.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment the resource belongs to.
* `domain` - The fully qualified domain name where the record can be located. 
* `is_protected` - A Boolean flag indicating whether or not parts of the record are unable to be explicitly managed. 
* `rdata` - The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see [Supported DNS Resource Record Types](https://docs.cloud.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm) 
* `record_hash` - A unique identifier for the record within its zone. 
* `rrset_version` - The latest version of the record's zone in which its RRSet differs from the preceding version. 
* `rtype` - The canonical name for the record's type, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
* `ttl` - The Time To Live for the record, in seconds. Using a TTL lower than 30 seconds is not recommended.
* `zone_name_or_id` - The name or OCID of the target zone.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Record
	* `update` - (Defaults to 20 minutes), when updating the Record
	* `delete` - (Defaults to 20 minutes), when destroying the Record


## Import

Import is not supported for this resource.

