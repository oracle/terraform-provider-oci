---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_capture_filters"
sidebar_current: "docs-oci-datasource-core-capture_filters"
description: |-
  Provides the list of Capture Filters in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_capture_filters
This data source provides the list of Capture Filters in Oracle Cloud Infrastructure Core service.

Lists the capture filters in the specified compartment.


## Example Usage

```hcl
data "oci_core_capture_filters" "test_capture_filters" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.capture_filter_display_name
	filter_type = var.capture_filter_filter_type
	state = var.capture_filter_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `filter_type` - (Optional) A filter to only return resources that match the given capture filterType. The filterType value is the string representation of enum - VTAP, FLOWLOG.
* `state` - (Optional) A filter to return only resources that match the given capture filter lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `capture_filters` - The list of capture_filters.

### CaptureFilter Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the capture filter. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `filter_type` - Indicates which service will use this capture filter
* `flow_log_capture_filter_rules` - The set of rules governing what traffic the Flow Log collects when creating a flow log capture filter. 
	* `destination_cidr` - Traffic to this CIDR will be captured in the flow log. 
	* `flow_log_type` - Type or types of flow logs to store. `ALL` includes records for both accepted traffic and rejected traffic. 
	* `icmp_options` - Optional and valid only for ICMP and ICMPv6. Use to specify a particular ICMP type and code as defined in:
		* [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml)
		* [ICMPv6 Parameters](https://www.iana.org/assignments/icmpv6-parameters/icmpv6-parameters.xhtml)

		If you specify ICMP or ICMPv6 as the protocol but omit this object, then all ICMP types and codes are allowed. If you do provide this object, the type is required and the code is optional. To enable MTU negotiation for ingress internet traffic via IPv4, make sure to allow type 3 ("Destination Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set"). If you need to specify multiple codes for a single type, create a separate security list rule for each. 
		* `code` - The ICMP code (optional).
		* `type` - The ICMP type.
	* `is_enabled` - Indicates whether a flow log capture filter rule is enabled. 
	* `priority` - A lower number indicates a higher priority, range 0-9. Each rule must have a distinct priority. 
	* `protocol` - The transport protocol the filter uses. 
	* `rule_action` - Include or exclude a ruleAction object. 
	* `sampling_rate` - Sampling interval as 1 of X, where X is an integer not greater than 100000. 
	* `source_cidr` - Traffic from this CIDR will be captured in the flow log. 
	* `tcp_options` - Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed. 
		* `destination_port_range` - 
			* `max` - The maximum port number, which must not be less than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number, which must not be greater than the maximum port number. 
		* `source_port_range` - 
			* `max` - The maximum port number, which must not be less than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number, which must not be greater than the maximum port number. 
	* `udp_options` - Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed. 
		* `destination_port_range` - 
			* `max` - The maximum port number, which must not be less than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number, which must not be greater than the maximum port number. 
		* `source_port_range` - 
			* `max` - The maximum port number, which must not be less than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number, which must not be greater than the maximum port number. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The capture filter's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)). 
* `state` - The capture filter's current administrative state.
* `time_created` - The date and time the capture filter was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2021-08-25T21:10:29.600Z` 
* `vtap_capture_filter_rules` - The set of rules governing what traffic a VTAP mirrors. 
	* `destination_cidr` - Traffic sent to this CIDR block through the VTAP source will be mirrored to the VTAP target. 
	* `icmp_options` - Optional and valid only for ICMP and ICMPv6. Use to specify a particular ICMP type and code as defined in:
		* [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml)
		* [ICMPv6 Parameters](https://www.iana.org/assignments/icmpv6-parameters/icmpv6-parameters.xhtml)

		If you specify ICMP or ICMPv6 as the protocol but omit this object, then all ICMP types and codes are allowed. If you do provide this object, the type is required and the code is optional. To enable MTU negotiation for ingress internet traffic via IPv4, make sure to allow type 3 ("Destination Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set"). If you need to specify multiple codes for a single type, create a separate security list rule for each. 
		* `code` - The ICMP code (optional).
		* `type` - The ICMP type.
	* `protocol` - The transport protocol used in the filter. If do not choose a protocol, all protocols will be used in the filter. Supported options are:
		* 1 = ICMP
		* 6 = TCP
		* 17 = UDP 
	* `rule_action` - Include or exclude packets meeting this definition from mirrored traffic. 
	* `source_cidr` - Traffic from this CIDR block to the VTAP source will be mirrored to the VTAP target. 
	* `tcp_options` - Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed. 
		* `destination_port_range` - 
			* `max` - The maximum port number, which must not be less than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number, which must not be greater than the maximum port number. 
		* `source_port_range` - 
			* `max` - The maximum port number, which must not be less than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number, which must not be greater than the maximum port number. 
	* `traffic_direction` - The traffic direction the VTAP is configured to mirror. 
	* `udp_options` - Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed. 
		* `destination_port_range` - 
			* `max` - The maximum port number, which must not be less than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number, which must not be greater than the maximum port number. 
		* `source_port_range` - 
			* `max` - The maximum port number, which must not be less than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number, which must not be greater than the maximum port number. 

