
---
subcategory: "Dns"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_rrset"
sidebar_current: "docs-oci-datasource-dns-rrset"
description: |-
  Provides details about a specific Rrset in Oracle Cloud Infrastructure Dns service
---

# Data Source: oci_dns_rrset
This data source provides details about a specific Rrset resource in Oracle Cloud Infrastructure Dns service.

Gets a list of all records in the specified RRSet. The results are
sorted by `recordHash` by default.


## Example Usage

```hcl
data "oci_dns_rrset" "test_rrset" {
	#Required
	domain = var.rrset_domain
	rtype = var.rrset_rtype
	zone_name_or_id = oci_dns_zone.test_zone.id

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment the resource belongs to.
* `domain` - (Required) The target fully-qualified domain name (FQDN) within the target zone.
* `rtype` - (Required) The type of the target RRSet within the target zone.
* `zone_name_or_id` - (Required) The name or OCID of the target zone.
* `zone_version` - (Optional) The version of the zone for which data is requested. 


## Attributes Reference

The following attributes are exported:

* `items` - 
	* `domain` - The fully qualified domain name where the record can be located. 
	* `is_protected` - A Boolean flag indicating whether or not parts of the record are unable to be explicitly managed. 
	* `rdata` - The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see [Supported DNS Resource Record Types](https://docs.cloud.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm) 
	* `record_hash` - A unique identifier for the record within its zone. 
	* `rrset_version` - The latest version of the record's zone in which its RRSet differs from the preceding version. 
	* `rtype` - The canonical name for the record's type, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
	* `ttl` - The Time To Live for the record, in seconds.

