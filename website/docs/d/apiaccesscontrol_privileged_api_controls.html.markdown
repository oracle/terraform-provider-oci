---
subcategory: "Apiaccesscontrol"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apiaccesscontrol_privileged_api_controls"
sidebar_current: "docs-oci-datasource-apiaccesscontrol-privileged_api_controls"
description: |-
  Provides the list of Privileged Api Controls in Oracle Cloud Infrastructure Apiaccesscontrol service
---

# Data Source: oci_apiaccesscontrol_privileged_api_controls
This data source provides the list of Privileged Api Controls in Oracle Cloud Infrastructure Apiaccesscontrol service.

Gets a list of PrivilegedApiControls.


## Example Usage

```hcl
data "oci_apiaccesscontrol_privileged_api_controls" "test_privileged_api_controls" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.privileged_api_control_display_name
	id = var.privileged_api_control_id
	resource_type = var.privileged_api_control_resource_type
	state = var.privileged_api_control_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the PrivilegedApiControl.
* `resource_type` - (Optional) A filter to return only lists of resources that match the entire given service type.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `privileged_api_control_collection` - The list of privileged_api_control_collection.

### PrivilegedApiControl Reference

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

