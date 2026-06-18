---
subcategory: "Events"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_events_rule"
sidebar_current: "docs-oci-resource-events-rule"
description: |-
  Provides the Rule resource in Oracle Cloud Infrastructure Events service
---

# oci_events_rule
This resource provides the Rule resource in Oracle Cloud Infrastructure Events service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/events/latest/Rule

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/events

Creates a new rule.


## Example Usage

```hcl
resource "oci_events_rule" "test_rule" {
	#Required
	actions {
		#Required
		action {
			#Required
			action_type = var.rule_actions_action_action_type
			is_enabled = var.rule_actions_action_is_enabled

			#Optional
			description = var.rule_actions_action_description
			function_id = oci_functions_function.test_function.id
			stream_id = oci_streaming_stream.test_stream.id
			topic_id = oci_ons_notification_topic.test_topic.id
		}
	}
	compartment_id = var.compartment_id
	# Recommended for new configurations. Use either condition_details or condition.
	condition_details {
		event_types = [
			"com.oraclecloud.objectstorage.createbucket",
			"com.oraclecloud.objectstorage.deletebucket",
		]
		data = jsonencode({})
	}
	display_name = var.rule_display_name
	is_enabled = var.rule_is_enabled

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.rule_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `actions` - (Required) (Updatable) A list of ActionDetails objects to create for a rule.
	* `action` - (Optional) (Updatable) A list of one or more ActionDetails objects.
		* `action_type` - (Required) (Updatable) The action to perform if the condition in the rule matches an event.
			* **ONS:** Send to an Oracle Notification Service topic.
			* **OSS:** Send to a stream from Oracle Streaming Service.
			* **FAAS:** Send to an Oracle Functions Service endpoint. 
		* `description` - (Optional) (Updatable) A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering confidential information. 
		* `function_id` - (Applicable when action_type=FAAS) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Function hosted by Oracle Functions Service. 
		* `is_enabled` - (Required) (Updatable) Whether or not this action is currently enabled.  Example: `true` 
		* `stream_id` - (Required when action_type=OSS) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream to which messages are delivered. 
		* `topic_id` - (Applicable when action_type=ONS) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic to which messages are delivered. 
	* `actions` - (Optional) (Updatable) Deprecated. Use `action` instead. This nested block is retained for backward compatibility.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this rule belongs. 
* `condition` - (Optional) (Updatable) A JSON string filter that specifies the event that will trigger actions associated with this rule. Use either `condition` or `condition_details`. This argument is retained for backward compatibility. For new configurations, `condition_details` is recommended because it avoids manually escaping JSON and is easier to maintain when matching multiple event types. A few  important things to remember about filters:
	* Fields not mentioned in the condition are ignored. You can create a valid filter that matches all events with two curly brackets: `{}` 

	For more examples, see  [Matching Events with Filters](https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/filterevents.htm).       
	* For a condition with fields to match an event, the event must contain all the field names  listed in the condition. Field names must appear in the condition with the same nesting  structure used in the event. 

	For a list of reference events, see  [Services that Produce Events](https://docs.cloud.oracle.com/iaas/Content/Events/Reference/eventsproducers.htm).       
	* Rules apply to events in the compartment in which you create them and any child compartments.  This means that a condition specified by a rule only matches events emitted from resources in  the compartment or any of its child compartments. 
	* Wildcard matching is supported with the asterisk (*) character. 

	For examples of wildcard matching, see  [Matching Events with Filters](https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/filterevents.htm)

	Example:
	```hcl
	condition = "{\"eventType\":[\"com.oraclecloud.objectstorage.createbucket\",\"com.oraclecloud.objectstorage.deletebucket\"],\"data\":{}}"
	```
* `condition_details` - (Optional) (Updatable) A structured helper for building the rule condition JSON. Use either `condition` or `condition_details`. This is the recommended form for new configurations.
	* `event_types` - (Optional) (Updatable) A list of event types to match.
	* `data` - (Optional) (Updatable) A JSON string containing additional event data filters.

	The following `condition_details` example is equivalent to the `condition` JSON string example above.

	Example:
	```hcl
	condition_details {
		event_types = [
			"com.oraclecloud.objectstorage.createbucket",
			"com.oraclecloud.objectstorage.deletebucket",
		]
		data = jsonencode({})
	}
	```
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A string that describes the details of the rule. It does not have to be unique, and you can change it. Avoid entering confidential information. 
* `display_name` - (Required) (Updatable) A string that describes the rule. It does not have to be unique, and you can change it. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_enabled` - (Required) (Updatable) Whether or not this rule is currently enabled.  Example: `true` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `actions` - A list of Action objects associated with a rule. 
	* `actions` - A list of one or more Action objects. 
		* `action_type` - The action to perform if the condition in the rule matches an event.
			* **ONS:** Send to an Oracle Notification Service topic.
			* **OSS:** Send to a stream from Oracle Streaming Service.
			* **FAAS:** Send to an Oracle Functions Service endpoint. 
		* `description` - A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering confidential information. 
		* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Function hosted by Oracle Functions Service. 
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the action. 
		* `is_enabled` - Whether or not this action is currently enabled.  Example: `true` 
		* `lifecycle_message` - A message generated by the Events service about the current state of this action. 
		* `state` - The current state of the rule. 
		* `stream_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream to which messages are delivered. 
		* `topic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic to which messages are delivered. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this rule belongs. 
* `condition` - A filter that specifies the event that will trigger actions associated with this rule. A few  important things to remember about filters:
	* Fields not mentioned in the condition are ignored. You can create a valid filter that matches all events with two curly brackets: `{}` 

	For more examples, see  [Matching Events with Filters](https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/filterevents.htm).       
	* For a condition with fields to match an event, the event must contain all the field names  listed in the condition. Field names must appear in the condition with the same nesting  structure used in the event. 

	For a list of reference events, see  [Services that Produce Events](https://docs.cloud.oracle.com/iaas/Content/Events/Reference/eventsproducers.htm).       
	* Rules apply to events in the compartment in which you create them and any child compartments.  This means that a condition specified by a rule only matches events emitted from resources in  the compartment or any of its child compartments. 
	* Wildcard matching is supported with the asterisk (*) character. 

	For examples of wildcard matching, see  [Matching Events with Filters](https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/filterevents.htm)

	Example: `\"eventType\": \"com.oraclecloud.databaseservice.autonomous.database.backup.end\"` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A string that describes the details of the rule. It does not have to be unique, and you can change it. Avoid entering confidential information. 
* `display_name` - A string that describes the rule. It does not have to be unique, and you can change it. Avoid entering confidential information.  Example: `"This rule sends a notification upon completion of DbaaS backup."` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this rule. 
* `is_enabled` - Whether or not this rule is currently enabled.  Example: `true` 
* `lifecycle_message` - A message generated by the Events service about the current state of this rule. 
* `state` - The current state of the rule. 
* `time_created` - The time this rule was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-09-12T22:47:12.613Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Rule
	* `update` - (Defaults to 20 minutes), when updating the Rule
	* `delete` - (Defaults to 20 minutes), when destroying the Rule


## Import

Rules can be imported using the `id`, e.g.

```
$ terraform import oci_events_rule.test_rule "id"
```
