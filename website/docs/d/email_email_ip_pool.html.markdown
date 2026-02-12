---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_email_ip_pool"
sidebar_current: "docs-oci-datasource-email-email_ip_pool"
description: |-
  Provides details about a specific Email Ip Pool in Oracle Cloud Infrastructure Email service
---

# Data Source: oci_email_email_ip_pool
This data source provides details about a specific Email Ip Pool resource in Oracle Cloud Infrastructure Email service.

Retrieves the specified IpPool by identifier

## Example Usage

```hcl
data "oci_email_email_ip_pool" "test_email_ip_pool" {
	#Required
	email_ip_pool_id = oci_email_email_ip_pool.test_email_ip_pool.id
}
```

## Argument Reference

The following arguments are supported:

* `email_ip_pool_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IpPool.


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

