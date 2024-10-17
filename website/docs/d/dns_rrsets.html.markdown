---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_rrsets"
sidebar_current: "docs-oci-datasource-dns-rrsets"
description: |-
  Provides the list of Rrsets in Oracle Cloud Infrastructure DNS service
---

# Data Source: oci_dns_rrsets
This data source provides the list of RRsets in Oracle Cloud Infrastructure DNS service.

Gets a list of all rrsets in the specified zone.

You can optionally filter the results using the listed parameters. When the zone name
is provided as a path parameter and `PRIVATE` is used for the scope query parameter then
the viewId parameter is required.


## Example Usage

```hcl
data "oci_dns_rrsets" "test_rrsets" {
	#Required
	zone_name_or_id = oci_dns_zone.test_zone.id

	#Optional
	domain = var.rrset_domain
	domain_contains = var.rrset_domain
	rtype = var.rrset_rtype
	scope = var.rrset_scope
	view_id = oci_dns_view.test_view.id
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Optional) The target fully-qualified domain name (FQDN) within the target zone.
* `domain_contains` - (Optional) Matches any rrset whose fully-qualified domain name (FQDN) contains the provided value.
* `rtype` - (Optional) Search by record type. Will match any record whose [type](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4) (case-insensitive) equals the provided value. 
* `scope` - (Optional) Specifies to operate only on resources that have a matching DNS scope. 
* `view_id` - (Optional) The OCID of the view the zone is associated with. Required when accessing a private zone by name.
* `zone_name_or_id` - (Required) The name or OCID of the target zone.


## Attributes Reference

The following attributes are exported:

* `rrsets` - The list of rrsets.

### Rrset Reference

The following attributes are exported:

* `domain` - The fully qualified domain name where the record can be located. 
* `rtype` - The type of DNS record, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
* `items` - 
	* `domain` - The fully qualified domain name where the record can be located. 
	* `is_protected` - A Boolean flag indicating whether or not parts of the record are unable to be explicitly managed. 
	* `rdata` - The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see [Supported DNS Resource Record Types](https://docs.cloud.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm) 
	* `record_hash` - A unique identifier for the record within its zone. 
	* `rrset_version` - The latest version of the record's zone in which its RRSet differs from the preceding version. 
	* `rtype` - The type of DNS record, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
	* `ttl` - The Time To Live for the record, in seconds. Using a TTL lower than 30 seconds is not recommended. 

