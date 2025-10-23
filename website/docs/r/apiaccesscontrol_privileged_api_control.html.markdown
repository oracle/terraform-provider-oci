---
subcategory: "Apiaccesscontrol"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apiaccesscontrol_privileged_api_control"
sidebar_current: "docs-oci-resource-apiaccesscontrol-privileged_api_control"
description: |-
  Provides the Privileged Api Control resource in Oracle Cloud Infrastructure Apiaccesscontrol service
---

# oci_apiaccesscontrol_privileged_api_control
This resource provides the Privileged Api Control resource in Oracle Cloud Infrastructure Apiaccesscontrol service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/oracle-api-access-control/latest/PrivilegedApiControl

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/apiaccesscontrol

Creates a PrivilegedApiControl.


## Example Usage

```hcl
resource "oci_apiaccesscontrol_privileged_api_control" "test_privileged_api_control" {
	#Required
	approver_group_id_list = var.privileged_api_control_approver_group_id_list
	compartment_id = var.compartment_id
	notification_topic_id = oci_ons_notification_topic.test_notification_topic.id
	privileged_operation_list {
		#Required
		api_name = oci_apigateway_api.test_api.name

		#Optional
		attribute_names = var.privileged_api_control_privileged_operation_list_attribute_names
		entity_type = var.privileged_api_control_privileged_operation_list_entity_type
	}
	resource_type = var.privileged_api_control_resource_type
	resources = var.privileged_api_control_resources

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.privileged_api_control_description
	display_name = var.privileged_api_control_display_name
	freeform_tags = {"Department"= "Finance"}
	number_of_approvers = var.privileged_api_control_number_of_approvers
}
```

## Argument Reference

The following arguments are supported:

* `approver_group_id_list` - (Required) (Updatable) List of user IAM group ids who can approve an privilegedApi request associated with a resource governed by this operator control.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the PrivilegedApiControl in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description of the privilegedApi control.
* `display_name` - (Optional) (Updatable) Name of the privilegedApi control It has to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `notification_topic_id` - (Required) (Updatable) The OCID of the Oracle Cloud Infrastructure Notification topic to publish messages related to this Delegation Control.
* `number_of_approvers` - (Optional) (Updatable) Number of approvers required to approve an privilegedApi request.
* `privileged_operation_list` - (Required) (Updatable) List of privileged operator operations. If Privileged API Managment is enabled for a resource it will be validated whether the operation done by the operator is a part of privileged operation. 
	* `api_name` - (Required) (Updatable) name of the api which needs to be protected.
	* `attribute_names` - (Optional) (Updatable) list of attributes belonging to the above api which needs to be protected.
	* `entity_type` - (Optional) (Updatable) type of the entity which needs to be protected.
* `resource_type` - (Required) (Updatable) resourceType for which the PrivilegedApiControl is applicable
* `resources` - (Required) (Updatable) contains Resource details


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `approver_group_id_list` - List of IAM user group ids who can approve an privilegedApi request associated with a target resource under the governance of this operator control.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of privilegedApi control.
* `display_name` - Name of the privilegedApi control. The name must be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the PrivilegedApiControl.
* `lifecycle_details` - A message that describes the current state of the PrivilegedApiControl in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `notification_topic_id` - The OCID of the Oracle Cloud Infrastructure Notification topic to publish messages related to this Privileged Api Control.
* `number_of_approvers` - Number of approvers required to approve an privilegedApi request.
* `privileged_operation_list` - List of privileged operations/apis. These operations/apis will be treaated as secured, once enabled by the Privileged API Managment for a resource. Any of these operations, if needs to be executed, needs to be raised as a PrivilegedApi Request which needs to be approved by customers or it can be pre-approved. 
	* `api_name` - name of the api which needs to be protected.
	* `attribute_names` - list of attributes belonging to the above api which needs to be protected.
	* `entity_type` - type of the entity which needs to be protected.
* `resource_type` - resourceType for which the PrivilegedApiControl is applicable
* `resources` - contains Resource details
* `state` - The current state of the PrivilegedApiControl.
* `state_details` - A message that describes the current state of the PrivilegedApiControl in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the PrivilegedApiControl was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_deleted` - The date and time the PrivilegedApiControl was marked for delete, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the PrivilegedApiControl was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Privileged Api Control
	* `update` - (Defaults to 20 minutes), when updating the Privileged Api Control
	* `delete` - (Defaults to 20 minutes), when destroying the Privileged Api Control


## Import

PrivilegedApiControls can be imported using the `id`, e.g.

```
$ terraform import oci_apiaccesscontrol_privileged_api_control.test_privileged_api_control "id"
```

