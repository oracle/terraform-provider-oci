---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitored_resources_list_member"
sidebar_current: "docs-oci-resource-stack_monitoring-monitored_resources_list_member"
description: |-
  Provides the Monitored Resources List Member resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_monitored_resources_list_member
This resource provides the Monitored Resources List Member resource in Oracle Cloud Infrastructure Stack Monitoring service.

List resources which are members of the given monitored resource

## Example Usage

```hcl
resource "oci_stack_monitoring_monitored_resources_list_member" "test_monitored_resources_list_member" {
	#Required
	monitored_resource_id = oci_stack_monitoring_monitored_resource.test_monitored_resource.id

	#Optional
	destination_resource_id = oci_stack_monitoring_destination_resource.test_destination_resource.id
	limit_level = var.monitored_resources_list_member_limit_level
}
```

## Argument Reference

The following arguments are supported:

* `destination_resource_id` - (Optional) Destination Monitored Resource Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `limit_level` - (Optional) The field which determines the depth of hierarchy while searching for members
* `monitored_resource_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of monitored resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - List of Members.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `external_id` - External resource is any Oracle Cloud Infrastructure resource identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) which is not a Stack Monitoring service resource. Currently supports only following resource type identifiers - externalcontainerdatabase, externalnoncontainerdatabase, externalpluggabledatabase and Oracle Cloud Infrastructure compute instance. 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `host_name` - Monitored Resource Host
	* `parent_id` - Parent monitored resource identifier
	* `resource_display_name` - Monitored resource display name.
	* `resource_id` - Monitored resource identifier
	* `resource_name` - Monitored resource name
	* `resource_type` - Monitored resource type
	* `state` - The current state of the Resource.
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitored Resources List Member
	* `update` - (Defaults to 20 minutes), when updating the Monitored Resources List Member
	* `delete` - (Defaults to 20 minutes), when destroying the Monitored Resources List Member


## Import

MonitoredResourcesListMembers can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_monitored_resources_list_member.test_monitored_resources_list_member "id"
```

