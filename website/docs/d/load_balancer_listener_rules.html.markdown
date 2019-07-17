---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_listener_rules"
sidebar_current: "docs-oci-datasource-load_balancer-listener_rules"
description: |-
  Provides the list of Listener Rules in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_listener_rules
This data source provides the list of Listener Rules in Oracle Cloud Infrastructure Load Balancer service.

Lists all of the rules from all of the rule sets associated with the specified listener. The response organizes
the rules in the following order:

*  Access control rules
*  Allow method rules
*  Request header rules
*  Response header rules


## Example Usage

```hcl
data "oci_load_balancer_listener_rules" "test_listener_rules" {
	#Required
	listener_name = "${oci_load_balancer_listener.test_listener.name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```

## Argument Reference

The following arguments are supported:

* `listener_name` - (Required) The name of the listener the rules are associated with. 
* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer associated with the listener.


## Attributes Reference

The following attributes are exported:

* `listener_rules` - The list of listener_rules.

### ListenerRule Reference

The following attributes are exported:

* `name` - The name of the rule set that the rule belongs to. 
* `rule` - A rule object that applies to the listener.
	* `action` - The action can be one of these values: `ADD_HTTP_REQUEST_HEADER`, `ADD_HTTP_RESPONSE_HEADER`, `ALLOW`, `CONTROL_ACCESS_USING_HTTP_METHODS`, `EXTEND_HTTP_REQUEST_HEADER_VALUE`, `EXTEND_HTTP_RESPONSE_HEADER_VALUE`, `REMOVE_HTTP_REQUEST_HEADER`, `REMOVE_HTTP_RESPONSE_HEADER`
	* `allowed_methods` - The list of HTTP methods allowed for this listener.

		By default, you can specify only the standard HTTP methods defined in the [HTTP Method Registry](http://www.iana.org/assignments/http-methods/http-methods.xhtml). You can also see a list of supported standard HTTP methods in the Load Balancing service documentation at [Managing Rule Sets](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrulesets.htm).

		Your backend application must be able to handle the methods specified in this list.

		The list of HTTP methods is extensible. If you need to configure custom HTTP methods, contact [My Oracle Support](http://support.oracle.com/) to remove the restriction for your tenancy.

		Example: ["GET", "PUT", "POST", "PROPFIND"] 
	* `conditions` - 
		* `attribute_name` - The attribute_name can be one of these values: `SOURCE_IP_ADDRESS`, `SOURCE_VCN_ID`, `SOURCE_VCN_IP_ADDRESS`
		* `attribute_value` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the originating VCN that an incoming packet must match.

			You can use this condition in conjunction with `SourceVcnIpAddressCondition`.

			**NOTE:** If you define this condition for a rule without a `SourceVcnIpAddressCondition`, this condition matches all incoming traffic in the specified VCN. 
	* `description` - A brief description of the access control rule. Avoid entering confidential information.

		example: `192.168.0.0/16 and 2001:db8::/32 are trusted clients. Whitelist them.` 
	* `header` - A header name that conforms to RFC 7230.  Example: `example_header_name` 
	* `prefix` - A string to prepend to the header value. The resulting header value must still conform to RFC 7230.  Example: `example_prefix_value` 
	* `status_code` - The HTTP status code to return when the requested HTTP method is not in the list of allowed methods. The associated status line returned with the code is mapped from the standard HTTP specification. The default value is `405 (Method Not Allowed)`.  Example: 403 
	* `suffix` - A string to append to the header value. The resulting header value must still conform to RFC 7230.  Example: `example_suffix_value` 
	* `value` - A header value that conforms to RFC 7230.  Example: `example_value` 

