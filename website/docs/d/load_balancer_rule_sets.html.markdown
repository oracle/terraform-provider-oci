---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_rule_sets"
sidebar_current: "docs-oci-datasource-load_balancer-rule_sets"
description: |-
  Provides the list of Rule Sets in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_rule_sets
This data source provides the list of Rule Sets in Oracle Cloud Infrastructure Load Balancer service.

Lists all rule sets associated with the specified load balancer.

## Example Usage

```hcl
data "oci_load_balancer_rule_sets" "test_rule_sets" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the specified load balancer.


## Attributes Reference

The following attributes are exported:

* `rule_sets` - The list of rule_sets.

### RuleSet Reference

The following attributes are exported:

* `items` - An array of rules that compose the rule set.
	* `action` - The action can be one of these values: `ADD_HTTP_REQUEST_HEADER`, `ADD_HTTP_RESPONSE_HEADER`, `ALLOW`, `CONTROL_ACCESS_USING_HTTP_METHODS`, `EXTEND_HTTP_REQUEST_HEADER_VALUE`, `EXTEND_HTTP_RESPONSE_HEADER_VALUE`, `REMOVE_HTTP_REQUEST_HEADER`, `REMOVE_HTTP_RESPONSE_HEADER`
	* `allowed_methods` - The list of HTTP methods allowed for this listener.

		By default, you can specify only the standard HTTP methods defined in the [HTTP Method Registry](http://www.iana.org/assignments/http-methods/http-methods.xhtml). You can also see a list of supported standard HTTP methods in the Load Balancing service documentation at [Managing Rule Sets](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrulesets.htm).

		Your backend application must be able to handle the methods specified in this list.

		The list of HTTP methods is extensible. If you need to configure custom HTTP methods, contact [My Oracle Support](http://support.oracle.com/) to remove the restriction for your tenancy.

		Example: ["GET", "PUT", "POST", "PROPFIND"] 
	* `conditions` - 
        * `attribute_name` - The attribute_name can be one of these values: `SOURCE_IP_ADDRESS`, `SOURCE_VCN_ID`, `SOURCE_VCN_IP_ADDRESS`
        * `attribute_value` - Depends on `attribute_name`:
            - when `attribute_name` = `SOURCE_IP_ADDRESS` | IPv4 or IPv6 address range to which the source IP address of incoming packet would be matched against
            - when `attribute_name` = `SOURCE_VCN_IP_ADDRESS` | IPv4 address range to which the original client IP address (in customer VCN) of incoming packet would be matched against
            - when `attribute_name` = `SOURCE_VCN_ID` | OCID of the customer VCN to which the service gateway embedded VCN ID of incoming packet would be matched against
	* `description` - Brief description of the access control rule. 
	* `header` - A header name that conforms to RFC 7230.  Example: `example_header_name` 
	* `prefix` - A string to prepend to the header value. The resulting header value must still conform to RFC 7230.  Example: `example_prefix_value` 
	* `status_code` - The HTTP status code to return when the requested HTTP method is not in the list of allowed methods. The associated status line returned with the code is mapped from the standard HTTP specification. The default value is `405 (Method Not Allowed)`.  Example: 403 
	* `suffix` - A string to append to the header value. The resulting header value must still conform to RFC 7230.  Example: `example_suffix_value` 
	* `value` - A header value that conforms to RFC 7230.  Example: `example_value` 
* `name` - The name for this set of rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_rule_set` 

