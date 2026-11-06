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

When the zone name is provided as a path parameter and `PRIVATE` is used for the scope query
parameter then the viewId query parameter is required.

## RRSet identity and records

An RRSet is the collection of records with the same domain name and record type
in one zone. The top-level `zone_name_or_id`, `domain` and `rtype` select that
RRSet. Each `items` block describes one record in it: repeat the selected domain
and type, and provide that record's `rdata` and `ttl`. Multiple A record values
therefore belong in multiple `items` blocks of one resource, rather than separate
resources managing the same RRSet. Use a consistent TTL for all records in a set.

## Example Usage

These examples use a zone resource named `example` and documentation IP addresses.
The zone can be created with `oci_dns_zone` using `compartment_id`, `name` and
`zone_type = "PRIMARY"`. See the [public RRSet example](https://github.com/oracle/terraform-provider-oci/tree/master/examples/dns/global/rrset.tf)
and [private RRSet example](https://github.com/oracle/terraform-provider-oci/tree/master/examples/dns/private/rrset)
for complete configurations. The public directory also includes other DNS examples.

### One A record

```hcl
resource "oci_dns_rrset" "single_record" {
  zone_name_or_id = oci_dns_zone.example.id
  domain          = "www.${oci_dns_zone.example.name}"
  rtype           = "A"

  items {
    domain = "www.${oci_dns_zone.example.name}"
    rtype  = "A"
    rdata  = "192.0.2.10"
    ttl    = 300
  }
}
```

### Multiple A records in one RRSet

```hcl
resource "oci_dns_rrset" "multiple_records" {
  zone_name_or_id = oci_dns_zone.example.id
  domain          = "api.${oci_dns_zone.example.name}"
  rtype           = "A"

  items {
    domain = "api.${oci_dns_zone.example.name}"
    rtype  = "A"
    rdata  = "192.0.2.20"
    ttl    = 300
  }

  items {
    domain = "api.${oci_dns_zone.example.name}"
    rtype  = "A"
    rdata  = "192.0.2.21"
    ttl    = 300
  }
}
```

For private zones, see the [private RRSet example](https://github.com/oracle/terraform-provider-oci/tree/master/examples/dns/private/rrset).
When identifying a private zone by name, provide its `view_id` as well.

## Behavior

- Creating a resource for an RRSet that already contains records fails with
  HTTP 409 (Conflict), including when `items` is omitted. Import an existing
  RRSet before managing it.
- Updating an existing resource replaces its record set with the configured
  `items`; removing an item removes that record. Keep the desired records explicit
  in the configuration, rather than using omission of `items` as a deletion instruction.

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment the zone belongs to.

	This parameter is deprecated and should be omitted. 
* `domain` - (Required) The target fully-qualified domain name (FQDN) within the target zone.
* `items` - (Optional) (Updatable) 
    The records to manage in this RRSet. Each item must use the top-level domain and record type. See Behavior above.
	* `domain` - (Required) The fully qualified domain name where the record can be located. 
	* `rdata` - (Required) (Updatable) The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see [Supported DNS Resource Record Types](https://docs.cloud.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm)  
	* `rtype` - (Required) The type of DNS record, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
	* `ttl` - (Required) (Updatable) The Time To Live for the record, in seconds. Using a TTL lower than 30 seconds is not recommended.
* `rtype` - (Required) The type of the target RRSet within the target zone.
* `scope` - (Optional) Specifies to operate only on resources that have a matching DNS scope. 
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

For Rrsets created using `scope` and `view_id`, these Rrsets can be imported using the `id`, e.g.

```
$ terraform import oci_dns_rrset.test_rrset "zoneNameOrId/{zoneNameOrId}/domain/{domain}/rtype/{rtype}/scope/{scope}/viewId/{viewId}"
```

skip adding `{view_id}` at the end if Rrset was created without `view_id`.
