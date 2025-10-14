---
subcategory: "Apiaccesscontrol"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apiaccesscontrol_privileged_api_request"
sidebar_current: "docs-oci-resource-apiaccesscontrol-privileged_api_request"
description: |-
  Provides the Privileged Api Request resource in Oracle Cloud Infrastructure Apiaccesscontrol service
---

# oci_apiaccesscontrol_privileged_api_request
This resource provides the Privileged Api Request resource in Oracle Cloud Infrastructure Apiaccesscontrol service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/oracle-api-access-control/latest/PrivilegedApiRequest

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/apiaccesscontrol

Creates a PrivilegedApiRequest.


## Example Usage

```hcl
resource "oci_apiaccesscontrol_privileged_api_request" "test_privileged_api_request" {
	#Required
	privileged_operation_list {
		#Required
		api_name = oci_apigateway_api.test_api.name

		#Optional
		attribute_names = var.privileged_api_request_privileged_operation_list_attribute_names
	}
	reason_summary = var.privileged_api_request_reason_summary
	resource_id = oci_cloud_guard_resource.test_resource.id

	#Optional
	compartment_id = var.compartment_id
	defined_tags = {"Operations.CostCenter"= "42"}
	duration_in_hrs = var.privileged_api_request_duration_in_hrs
	freeform_tags = {"Department"= "Finance"}
	notification_topic_id = oci_ons_notification_topic.test_notification_topic.id
	reason_detail = var.privileged_api_request_reason_detail
	severity = var.privileged_api_request_severity
	sub_resource_name_list = var.privileged_api_request_sub_resource_name_list
	ticket_numbers = var.privileged_api_request_ticket_numbers
	time_requested_for_future_access = var.privileged_api_request_time_requested_for_future_access
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `duration_in_hrs` - (Optional) Duration in hours for which access is sought on the target resource.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `notification_topic_id` - (Optional) The OCID of the Oracle Cloud Infrastructure Notification topic to publish messages related to this Privileged Api Request.
* `privileged_operation_list` - (Required) List of api names, attributes for which approval is sought by the user. 
	* `api_name` - (Required) name of the api which needs to be protected.
	* `attribute_names` - (Optional) list of attributes belonging to the above api which needs to be protected.
* `reason_detail` - (Optional) Reason in detail for which the operator is requesting access on the target resource.
* `reason_summary` - (Required) Summary comment by the operator creating the access request.
* `resource_id` - (Required) The OCID of the target resource associated with the access request. The operator raises an access request to get approval to access the target resource. 
* `severity` - (Optional) Priority assigned to the access request by the operator
* `sub_resource_name_list` - (Optional) The subresource names requested for approval.
* `ticket_numbers` - (Optional) A list of ticket numbers related to this Privileged Api Access Request, e.g. Service Request (SR) number and JIRA ticket number. 
* `time_requested_for_future_access` - (Optional) Time in future when the user for the privilegedApi request needs to be created in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z' 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `approver_details` - Contains the approver details who have approved the privilegedApi Request during the initial request.
	* `approval_action` - The action done by the approver.
	* `approval_comment` - Comment specified by the approver of the request.
	* `approver_id` - The userId of the approver.
	* `time_approved_for_access` - Time for when the privilegedApi request should start that is authorized by the customer in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z' 
	* `time_of_authorization` - Time when the privilegedApi request was authorized by the customer in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z' 
* `closure_comment` - The comment entered by the operator while closing the request.
* `compartment_id` - The OCID of the compartment that contains the access request.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Name of the privilegedApi control. The name must be unique.
* `duration_in_hrs` - Duration in hours for which access is sought on the target resource.
* `entity_type` - entityType of resource for which the AccessRequest is applicable
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the privilegedApi request.
* `lifecycle_details` - more in detail about the lifeCycleState.
* `notification_topic_id` - The OCID of the Oracle Cloud Infrastructure Notification topic to publish messages related to this privileged api request.
* `number_of_approvers_required` - Number of approvers required to approve an privilegedApi request.
* `privileged_api_control_id` - The OCID of the privilegedApi control governing the target resource.
* `privileged_api_control_name` - Name of the privilegedApi control governing the target resource.
* `privileged_operation_list` - List of api names, attributes for which approval is sought by the user. 
	* `api_name` - name of the api which needs to be protected.
	* `attribute_names` - list of attributes belonging to the above api which needs to be protected.
* `reason_detail` - Reason in Detail for which the operator is requesting access on the target resource.
* `reason_summary` - Summary comment by the operator creating the access request.
* `request_id` - This is an automatic identifier generated by the system which is easier for human comprehension.
* `requested_by` - List of Users who has created this privilegedApiRequest. 
* `resource_id` - The OCID of the target resource associated with the access request. The operator raises an access request to get approval to access the target resource. 
* `resource_name` - resourceName for which the PrivilegedApiRequest is applicable
* `resource_type` - resourceType for which the AccessRequest is applicable
* `severity` - Priority assigned to the access request by the operator
* `state` - The current state of the PrivilegedApiRequest.
* `state_details` - A message that describes the current state of the PrivilegedApiControl in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `sub_resource_name_list` - The subresource names requested for approval.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `ticket_numbers` - A list of ticket numbers related to this Privileged Api Access Request, e.g. Service Request (SR) number and JIRA ticket number. 
* `time_created` - Time when the privilegedApi request was created in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z' 
* `time_requested_for_future_access` - Time in future when the user for the privilegedApi request needs to be created in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the privilegedApi request was last modified in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z' 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Privileged Api Request
	* `update` - (Defaults to 20 minutes), when updating the Privileged Api Request
	* `delete` - (Defaults to 20 minutes), when destroying the Privileged Api Request


## Import

PrivilegedApiRequests can be imported using the `id`, e.g.

```
$ terraform import oci_apiaccesscontrol_privileged_api_request.test_privileged_api_request "id"
```

