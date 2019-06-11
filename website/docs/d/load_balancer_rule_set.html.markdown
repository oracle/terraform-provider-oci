---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_rule_set"
sidebar_current: "docs-oci-datasource-load_balancer-rule_set"
description: |-
  Provides details about a specific Rule Set in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_rule_set
This data source provides details about a specific Rule Set resource in Oracle Cloud Infrastructure Load Balancer service.

Gets the specified set of rules.

## Example Usage

```hcl
data "oci_load_balancer_rule_set" "test_rule_set" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.rule_set_name}"
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the specified load balancer.
* `name` - (Required) The name of the rule set to retrieve.  Example: `example_rule_set` 


## Attributes Reference

The following attributes are exported:

* `items` - An array of rules that compose the rule set.
	* `action` - The action can be one of these values: `ADD_HTTP_REQUEST_HEADER`, `ADD_HTTP_RESPONSE_HEADER`, `CONTROL_ACCESS_USING_HTTP_METHODS`, `EXTEND_HTTP_REQUEST_HEADER_VALUE`, `EXTEND_HTTP_RESPONSE_HEADER_VALUE`, `REMOVE_HTTP_REQUEST_HEADER`, `REMOVE_HTTP_RESPONSE_HEADER`
	* `allowed_methods` - The list of HTTP methods allowed for this listener.

		By default, you can specify only the standard HTTP methods defined in the [HTTP Method Registry](http://www.iana.org/assignments/http-methods/http-methods.xhtml). You can also see a list of supported standard HTTP methods in the Load Balancing service documentation at [Managing Rule Sets](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrulesets.htm).

		Your backend application must be able to handle the methods specified in this list.

		The list of HTTP methods is extensible. If you need to configure custom HTTP methods, contact [My Oracle Support](http://support.oracle.com/) to remove the restriction for your tenancy.

		Example: ["GET", "PUT", "POST", "PROPFIND"] 
	* `header` - A header name that conforms to RFC 7230.  Example: `example_header_name` 
	* `prefix` - A string to prepend to the header value. The resulting header value must still conform to RFC 7230.  Example: `example_prefix_value` 
	* `status_code` - The HTTP status code to return when the requested HTTP method is not in the list of allowed methods. The associated status line returned with the code is mapped from the standard HTTP specification. The default value is `405 (Method Not Allowed)`.  Example: 403 
	* `suffix` - A string to append to the header value. The resulting header value must still conform to RFC 7230.  Example: `example_suffix_value` 
	* `value` - A header value that conforms to RFC 7230.  Example: `example_value` 
* `name` - The name for this set of rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_rule_set` 

