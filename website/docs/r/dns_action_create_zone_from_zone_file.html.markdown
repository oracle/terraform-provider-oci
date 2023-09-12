---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_action_create_zone_from_zone_file"
sidebar_current: "docs-oci-resource-dns-action_create_zone_from_zone_file"
description: |-
  Provides the Action Create Zone From Zone File resource in Oracle Cloud Infrastructure DNS service
---

# oci_dns_action_create_zone_from_zone_file
This resource provides the Action Create Zone From Zone File resource in Oracle Cloud Infrastructure DNS service.

Creates a new zone from a zone file in the specified compartment. Not supported for private zones.

After the zone has been created, it should be further managed by importing it to an `oci_dns_zone` resource.


## Example Usage

```hcl
resource "oci_dns_action_create_zone_from_zone_file" "test_action_create_zone_from_zone_file" {
	#Required
	create_zone_from_zone_file_details = var.action_create_zone_from_zone_file_create_zone_from_zone_file_details
	compartment_id = var.compartment_id

	#Optional
	scope = var.action_create_zone_from_zone_file_scope
	view_id = oci_dns_view.test_view.id
}
```

## Argument Reference

The following arguments are supported:

* `create_zone_from_zone_file_details` - (Required) The zone file contents.
* `compartment_id` - (Required) The OCID of the compartment the resource belongs to.
* `scope` - (Optional) Specifies to operate only on resources that have a matching DNS scope. 
* `view_id` - (Optional) The OCID of the view the resource is associated with.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the zone.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Operations": {"CostCenter": "42"}}` 
* `external_downstreams` - External secondary servers for the zone. This field is currently not supported when `zoneType` is `SECONDARY` or `scope` is `PRIVATE`. 
	* `address` - The server's IP address (IPv4 or IPv6).
	* `port` - The server's port. Port value must be a value of 53, otherwise omit the port value. 
	* `tsig_key_id` - The OCID of the TSIG key. A TSIG key is used to secure DNS messages (in this case, zone transfers) between two systems that both have the (shared) secret. 
* `external_masters` - External master servers for the zone. `externalMasters` becomes a required parameter when the `zoneType` value is `SECONDARY`. 
	* `address` - The server's IP address (IPv4 or IPv6).
	* `port` - The server's port. Port value must be a value of 53, otherwise omit the port value. 
	* `tsig_key_id` - The OCID of the TSIG key.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `id` - The OCID of the zone.
* `is_protected` - A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed. 
* `name` - The name of the zone.
* `nameservers` - The authoritative nameservers for the zone.
	* `hostname` - The hostname of the nameserver.
* `scope` - The scope of the zone.
* `self` - The canonical absolute URL of the resource.
* `serial` - The current serial of the zone. As seen in the zone's SOA record. 
* `state` - The current state of the zone resource.
* `time_created` - The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `version` - Version is the never-repeating, totally-orderable, version of the zone, from which the serial field of the zone's SOA record is derived. 
* `view_id` - The OCID of the private view containing the zone. This value will be null for zones in the global DNS, which are publicly resolvable and not part of a private view. 
* `zone_transfer_servers` - The Oracle Cloud Infrastructure nameservers that transfer the zone data with external nameservers. 
	* `address` - The server's IP address (IPv4 or IPv6).
	* `is_transfer_destination` - A Boolean flag indicating whether or not the server is a zone data transfer destination. 
	* `is_transfer_source` - A Boolean flag indicating whether or not the server is a zone data transfer source. 
	* `port` - The server's port. 
* `zone_type` - The type of the zone. Must be either `PRIMARY` or `SECONDARY`. `SECONDARY` is only supported for GLOBAL zones. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Action Create Zone From Zone File
	* `update` - (Defaults to 20 minutes), when updating the Action Create Zone From Zone File
	* `delete` - (Defaults to 20 minutes), when destroying the Action Create Zone From Zone File


## Import

ActionCreateZoneFromZoneFile can be imported using the `id`, e.g.

```
$ terraform import oci_dns_action_create_zone_from_zone_file.test_action_create_zone_from_zone_file "id"
```

