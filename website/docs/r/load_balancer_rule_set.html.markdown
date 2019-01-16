---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_rule_set"
sidebar_current: "docs-oci-resource-load_balancer-rule_set"
description: |-
  Provides the Rule Set resource in Oracle Cloud Infrastructure Load Balancer service
---

# oci_load_balancer_rule_set
This resource provides the Rule Set resource in Oracle Cloud Infrastructure Load Balancer service.

Creates a new rule set associated with the specified load balancer.


## Example Usage

```hcl
resource "oci_load_balancer_rule_set" "test_rule_set" {
	#Required
	items {
		#Required
		action = "${var.rule_set_items_action}"
		header = "${var.rule_set_items_header}"

		#Optional
		prefix = "${var.rule_set_items_prefix}"
		suffix = "${var.rule_set_items_suffix}"
		value = "${var.rule_set_items_value}"
	}
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.rule_set_name}"
}
```

## Argument Reference

The following arguments are supported:

* `items` - (Required) (Updatable) 
	* `action` - (Required) (Updatable) The action can be one of these values: `ADD_HTTP_REQUEST_HEADER`, `ADD_HTTP_RESPONSE_HEADER`, `EXTEND_HTTP_REQUEST_HEADER_VALUE`, `EXTEND_HTTP_RESPONSE_HEADER_VALUE`, `REMOVE_HTTP_REQUEST_HEADER`, `REMOVE_HTTP_RESPONSE_HEADER`
	* `header` - (Required) (Updatable) A header name that conforms to RFC 7230.  Example: `example_header_name` 
	* `prefix` - (Applicable when action=EXTEND_HTTP_REQUEST_HEADER_VALUE | EXTEND_HTTP_RESPONSE_HEADER_VALUE) (Updatable) A string to prepend to the header value. The resulting header value must still conform to RFC 7230.  Example: `example_prefix_value` 
	* `suffix` - (Applicable when action=EXTEND_HTTP_REQUEST_HEADER_VALUE | EXTEND_HTTP_RESPONSE_HEADER_VALUE) (Updatable) A string to append to the header value. The resulting header value must still conform to RFC 7230.  Example: `example_suffix_value` 
	* `value` - (Required when action=ADD_HTTP_REQUEST_HEADER | ADD_HTTP_RESPONSE_HEADER) (Updatable) A header value that conforms to RFC 7230.  Example: `example_value` 
* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the specified load balancer.
* `name` - (Required) The name for this set of rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_rule_set` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - 
	* `action` - The action can be one of these values: `ADD_HTTP_REQUEST_HEADER`, `ADD_HTTP_RESPONSE_HEADER`, `EXTEND_HTTP_REQUEST_HEADER_VALUE`, `EXTEND_HTTP_RESPONSE_HEADER_VALUE`, `REMOVE_HTTP_REQUEST_HEADER`, `REMOVE_HTTP_RESPONSE_HEADER`
	* `header` - A header name that conforms to RFC 7230.  Example: `example_header_name` 
	* `prefix` - A string to prepend to the header value. The resulting header value must still conform to RFC 7230.  Example: `example_prefix_value` 
	* `suffix` - A string to append to the header value. The resulting header value must still conform to RFC 7230.  Example: `example_suffix_value` 
	* `value` - A header value that conforms to RFC 7230.  Example: `example_value` 
* `name` - The name for this set of rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_rule_set` 

## Import

RuleSets can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_rule_set.test_rule_set "loadBalancers/{loadBalancerId}/ruleSets/{ruleSetName}" 
```

