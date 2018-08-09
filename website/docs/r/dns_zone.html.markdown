---
layout: "oci"
page_title: "OCI: oci_dns_zone"
sidebar_current: "docs-oci-resource-dns-zone"
description: |-
  Creates and manages an OCI DnsZone
---

# oci_dns_zone
The `oci_dns_zone` resource creates and manages an OCI DnsZone

Creates a new zone in the specified compartment.


## Example Usage

```hcl
resource "oci_dns_zone" "test_zone" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${var.zone_name}"
	zone_type = "${var.zone_zone_type}"

	#Optional
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
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment the resource belongs to.
* `external_masters` - (Optional) (Updatable) External master servers for the zone.
	* `address` - (Required) (Updatable) The server's IP address (IPv4 or IPv6).
	* `port` - (Optional) (Updatable) The server's port.
	* `tsig` - (Optional) (Updatable) 
		* `algorithm` - (Required) (Updatable) TSIG Algorithms are encoded as domain names, but most consist of only one non-empty label, which is not required to be explicitly absolute. For a full list of TSIG algorithms, see [Secret Key Transaction Authentication for DNS (TSIG) Algorithm Names](http://www.iana.org/assignments/tsig-algorithm-names/tsig-algorithm-names.xhtml#tsig-algorithm-names-1) 
		* `name` - (Required) (Updatable) A domain name identifying the key for a given pair of hosts.
		* `secret` - (Required) (Updatable) A base64 string encoding the binary shared secret.
* `name` - (Required) The name of the zone.
* `zone_type` - (Required) The type of the zone. Must be either `PRIMARY` or `SECONDARY`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the zone.
* `external_masters` - External master servers for the zone.
	* `address` - The server's IP address (IPv4 or IPv6).
	* `port` - The server's port.
	* `tsig` - A TSIG key
		* `algorithm` - TSIG Algorithms are encoded as domain names, but most consist of only one non-empty label, which is not required to be explicitly absolute. For a full list of TSIG algorithms, see [Secret Key Transaction Authentication for DNS (TSIG) Algorithm Names](http://www.iana.org/assignments/tsig-algorithm-names/tsig-algorithm-names.xhtml#tsig-algorithm-names-1) 
		* `name` - A domain name identifying the key for a given pair of hosts.
		* `secret` - A base64 string encoding the binary shared secret.
* `id` - The OCID of the zone.
* `name` - The name of the zone.
* `self` - The canonical absolute URL of the resource.
* `serial` - The current serial of the zone. As seen in the zone's SOA record. 
* `state` - The current state of the zone resource.
* `time_created` - The date and time the image was created in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.  **Example:** `2016-07-22T17:23:59:60Z` 
* `version` - Version is the never-repeating, totally-orderable, version of the zone, from which the serial field of the zone's SOA record is derived. 
* `zone_type` - The type of the zone. Must be either `PRIMARY` or `SECONDARY`. 

## Import

Zones can be imported using the `id`, e.g.

```
$ terraform import oci_dns_zone.test_zone "id"
```
