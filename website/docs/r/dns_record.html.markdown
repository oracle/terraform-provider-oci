---
layout: "oci"
page_title: "OCI: oci_dns_record"
sidebar_current: "docs-oci-resource-dns-record"
description: |-
  Creates and manages an OCI DnsRecord
---

# oci_dns_record
The `oci_dns_record` resource creates and manages an OCI DnsRecord

Replaces records in the specified zone with the records specified in the
request body. If a specified record does not exist, it will be created.
If the record exists, then it will be updated to represent the record in
the body of the request. If a record in the zone does not exist in the
request body, the record will be removed from the zone.


## Example Usage

```hcl
resource "oci_dns_record" "test_record" {
	#Required
	zone_name_or_id = "${oci_dns_zone_name_or.test_zone_name_or.id}"

	#Optional
	compartment_id = "${var.compartment_id}"
	domain = "${var.record_items_domain}"
	rdata = "${var.record_items_rdata}"
	rtype = "${var.record_items_rtype}"
	ttl = "${var.record_items_ttl}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment the resource belongs to. If supplied, it must match the Zone's compartment ocid. 
* `domain` - (Optional) (Updatable) The fully qualified domain name where the record can be located.  
* `rdata` - (Optional) (Updatable) The record's data, as whitespace-delimited tokens in type-specific presentation format.  
* `rtype` - (Optional) (Updatable) The canonical name for the record's type, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
* `ttl` - (Optional) (Updatable) The Time To Live for the record, in seconds.
* `zone_name_or_id` - (Required) The name or OCID of the target zone.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment the resource belongs to.
* `domain` - The fully qualified domain name where the record can be located. 
* `is_protected` - A Boolean flag indicating whether or not parts of the record are unable to be explicitly managed. 
* `rdata` - The record's data, as whitespace-delimited tokens in type-specific presentation format. 
* `record_hash` - A unique identifier for the record within its zone. 
* `rrset_version` - The latest version of the record's zone in which its RRSet differs from the preceding version. 
* `rtype` - The canonical name for the record's type, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
* `ttl` - The Time To Live for the record, in seconds.
* `zone_name_or_id` - The name or OCID of the target zone.

## Import

PublicIps can be imported using the `id`, e.g.

```
$ terraform import oci_dns_records.test_record "id"
```
