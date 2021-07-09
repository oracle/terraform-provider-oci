---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_zone"
sidebar_current: "docs-oci-resource-dns-zone"
description: |-
  Provides the Zone resource in Oracle Cloud Infrastructure DNS service
---

# oci_dns_zone
This resource provides the Zone resource in Oracle Cloud Infrastructure DNS service.

Creates a new zone in the specified compartment. Additionally, for Private DNS, 
the `scope` and `viewId` query parameters are required when creating private zones.

## Example Usage

```hcl
resource "oci_dns_zone" "test_zone" {
	#Required
	compartment_id = var.compartment_id
	name = var.zone_name
	zone_type = var.zone_zone_type

	#Optional
	defined_tags = var.zone_defined_tags
	external_masters {
		#Required
		address = var.zone_external_masters_address

		#Optional
		port = var.zone_external_masters_port
		tsig_key_id = oci_dns_tsig_key.test_tsig_key.id
	}
	freeform_tags = var.zone_freeform_tags
	scope = var.zone_scope
	view_id = oci_dns_view.test_view.id
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
	* `tsig_key_id` - (Optional) (Updatable) The OCID of the TSIG key.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `name` - (Required) The name of the zone.
* `scope` - (Optional) Specifies to operate only on resources that have a matching DNS scope. 
This value will be null for zones in the global DNS and `PRIVATE` when creating a private zone.
* `view_id` - (Optional) The OCID of the private view containing the zone. This value will be null for zones in the global DNS, which are publicly resolvable and not part of a private view. 
* `zone_type` - (Required) The type of the zone. Must be either `PRIMARY` or `SECONDARY`. `SECONDARY` is only supported for GLOBAL zones. 


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
* `zone_type` - The type of the zone. Must be either `PRIMARY` or `SECONDARY`. `SECONDARY` is only supported for GLOBAL zones. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Zone
	* `update` - (Defaults to 20 minutes), when updating the Zone
	* `delete` - (Defaults to 20 minutes), when destroying the Zone


## Import

For legacy Zones that were created without using `scope`, these Zones can be imported using the `id`, e.g.

```
$ terraform import oci_dns_zone.test_zone "id"
```
For Zones created using `scope` and `view_id`, these Zones can be imported using the `id`, e.g.

```
$ terraform import oci_dns_zone.test_zone "zoneNameOrId/{zoneNameOrId}/scope/{scope}/viewId/{viewId}"
```

skip adding `{view_id}` at the end if Zone was created without `view_id`.
