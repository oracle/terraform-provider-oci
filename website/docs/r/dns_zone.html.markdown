---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_zone"
sidebar_current: "docs-oci-resource-dns-zone"
description: |-
  Provides the Zone resource in Oracle Cloud Infrastructure Dns service
---

# oci_dns_zone
This resource provides the Zone resource in Oracle Cloud Infrastructure Dns service.

Creates a new zone in the specified compartment.


## Example Usage

```hcl
resource "oci_dns_zone" "test_zone" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${var.zone_name}"
	zone_type = "${var.zone_zone_type}"

	#Optional
	defined_tags = "${var.zone_defined_tags}"
	external_masters {
		#Required
		address = "${var.zone_external_masters_address}"

		#Optional
		port = "${var.zone_external_masters_port}"
		tsig {
			#Required
			algorithm = "${var.zone_external_masters_tsig_algorithm}"
			name = "${var.zone_external_masters_tsig_name}"
			secret = "${var.zone_external_masters_tsig_secret}"
		}
	}
	freeform_tags = "${var.zone_freeform_tags}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment the resource belongs to.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Operations.CostCenter": "42"}` 
* `external_masters` - (Optional) (Updatable) External master servers for the zone. `externalMasters` becomes a required parameter when the `zoneType` value is `SECONDARY`. 
	* `address` - (Required) (Updatable) The server's IP address (IPv4 or IPv6).
	* `port` - (Optional) (Updatable) The server's port. Port value must be a value of 53, otherwise omit the port value. 
	* `tsig` - (Optional) (Updatable) 
		* `algorithm` - (Required) (Updatable) TSIG Algorithms are encoded as domain names, but most consist of only one non-empty label, which is not required to be explicitly absolute. Applicable algorithms include: hmac-sha1, hmac-sha224, hmac-sha256, hmac-sha512. For more information on these algorithms, see [RFC 4635](https://tools.ietf.org/html/rfc4635#section-2). 
		* `name` - (Required) (Updatable) A domain name identifying the key for a given pair of hosts.
		* `secret` - (Required) (Updatable) A base64 string encoding the binary shared secret.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `name` - (Required) The name of the zone.
* `zone_type` - (Required) The type of the zone. Must be either `PRIMARY` or `SECONDARY`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Import

Zones can be imported using the `id`, e.g.

```
$ terraform import oci_dns_zone.test_zone "id"
```

