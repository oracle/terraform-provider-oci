---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_zones"
sidebar_current: "docs-oci-datasource-dns-zones"
description: |-
  Provides the list of Zones in Oracle Cloud Infrastructure Dns service
---

# Data Source: oci_dns_zones
This data source provides the list of Zones in Oracle Cloud Infrastructure Dns service.

Gets a list of all zones in the specified compartment. The collection
can be filtered by name, time created, and zone type.


## Example Usage

```hcl
data "oci_dns_zones" "test_zones" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	name = "${var.zone_name}"
	name_contains = "${var.zone_name_contains}"
	state = "${var.zone_state}"
	time_created_greater_than_or_equal_to = "${var.zone_time_created_greater_than_or_equal_to}"
	time_created_less_than = "${var.zone_time_created_less_than}"
	zone_type = "${var.zone_zone_type}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment the resource belongs to.
* `name` - (Optional) A case-sensitive filter for zone names. Will match any zone with a name that equals the provided value. 
* `name_contains` - (Optional) Search by zone name. Will match any zone whose name (case-insensitive) contains the provided value. 
* `sort_by` - (Optional) The field by which to sort zones. Allowed values are: name|zoneType|timeCreated
* `sort_order` - The order to sort the resources. Allowed values are: ASC|DESC  
* `state` - (Optional) The state of a resource.
* `time_created_greater_than_or_equal_to` - (Optional) An [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp that states all returned resources were created on or after the indicated time. 
* `time_created_less_than` - (Optional) An [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp that states all returned resources were created before the indicated time. 
* `zone_type` - (Optional) Search by zone type, `PRIMARY` or `SECONDARY`. Will match any zone whose type equals the provided value. 


## Attributes Reference

The following attributes are exported:

* `zones` - A list of DNS zones.

### DnsZone Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the zone.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Operations.CostCenter": "42"}` 
* `external_masters` - External master servers for the zone. `externalMasters` becomes a required parameter when the `zoneType` value is `SECONDARY`. 
	* `address` - The server's IP address (IPv4 or IPv6).
	* `port` - The server's port. Port value must be a value of 53, otherwise omit the port value. 
	* `tsig` - A TSIG key
		* `algorithm` - TSIG Algorithms are encoded as domain names, but most consist of only one non-empty label, which is not required to be explicitly absolute. Applicable algorithms include: hmac-sha1, hmac-sha224, hmac-sha256, hmac-sha512. For more information on these algorithms, see [RFC 4635](https://tools.ietf.org/html/rfc4635#section-2). 
		* `name` - A domain name identifying the key for a given pair of hosts.
		* `secret` - A base64 string encoding the binary shared secret.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `id` - The OCID of the zone.
* `name` - The name of the zone.
* `nameservers` - The authoritative nameservers for the zone.
	* `hostname` - The hostname of the nameserver.
* `self` - The canonical absolute URL of the resource.
* `serial` - The current serial of the zone. As seen in the zone's SOA record. 
* `state` - The current state of the zone resource.
* `time_created` - The date and time the resource was created in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `version` - Version is the never-repeating, totally-orderable, version of the zone, from which the serial field of the zone's SOA record is derived. 
* `zone_type` - The type of the zone. Must be either `PRIMARY` or `SECONDARY`. 

