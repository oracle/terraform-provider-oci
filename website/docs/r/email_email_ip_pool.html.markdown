---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_email_ip_pool"
sidebar_current: "docs-oci-resource-email-email_ip_pool"
description: |-
  Provides the Email Ip Pool resource in Oracle Cloud Infrastructure Email service
---

# oci_email_email_ip_pool
This resource provides the Email Ip Pool resource in Oracle Cloud Infrastructure Email service.

Creates a new EmailIpPool.

## Example Usage

```hcl
resource "oci_email_email_ip_pool" "test_email_ip_pool" {
	#Required
	compartment_id = var.compartment_id
	name = var.email_ip_pool_name
	outbound_ips = []

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.email_ip_pool_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for this IpPool.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the IpPool. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - (Required) The name of the IpPool. The name must be unique within a region.  The name is case sensitive and supported characters include alphanumeric, hyphens ("-") and underscore ("_") characters.  Example: green_pool-1 
* `outbound_ips` - (Required) A list of outbound public IPs for assignment to the IpPool. These IPs must be in the AVAILABLE state to be eligible for assignment.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the IpPool.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the IpPool. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The unique [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IpPool resource that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'INACTIVE' state. 
* `locks` - Locks associated with this resource.
	* `compartment_id` - The lock compartment ID.
	* `message` - A message added by the lock creator. The message typically gives an indication of why the resource is locked. 
	* `related_resource_id` - The resource ID that is locking this resource. Indicates that deleting this resource removes the lock. 
	* `time_created` - Indicates when the lock was created, in the format defined by RFC 3339.
	* `type` - Lock type.
* `name` - The name of the IpPool. The name must be unique within a region. The name is case sensitive and supported characters include alphanumeric, hyphens ("-") and underscore ("_") characters.  Example: green_pool-1 
* `outbound_ips` - Summary of outbound IPs assigned to the IpPool.
	* `assignment_state` - The assignment state of the public IP address.
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'DRAINING' state. 
	* `outbound_ip` - The public IP address assigned to the tenancy.
	* `state` - The current state of the Email Outbound Public IP.
* `state` - The current state of the IpPool.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the IpPool was created. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".  Example: `2021-02-12T22:47:12.613Z` 
* `time_updated` - The time of the last change to the IpPool, due to a state change or an update operation. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ". 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Email Ip Pool
	* `update` - (Defaults to 20 minutes), when updating the Email Ip Pool
	* `delete` - (Defaults to 20 minutes), when destroying the Email Ip Pool


## Import

EmailIpPools can be imported using the `id`, e.g.

```
$ terraform import oci_email_email_ip_pool.test_email_ip_pool "id"
```

