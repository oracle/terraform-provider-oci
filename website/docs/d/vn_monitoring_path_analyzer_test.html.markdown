---
subcategory: "Vn Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_vn_monitoring_path_analyzer_test"
sidebar_current: "docs-oci-datasource-vn_monitoring-path_analyzer_test"
description: |-
  Provides details about a specific Path Analyzer Test in Oracle Cloud Infrastructure Vn Monitoring service
---

# Data Source: oci_vn_monitoring_path_analyzer_test
This data source provides details about a specific Path Analyzer Test resource in Oracle Cloud Infrastructure Vn Monitoring service.

Gets a `PathAnalyzerTest` using its identifier.

## Example Usage

```hcl
data "oci_vn_monitoring_path_analyzer_test" "test_path_analyzer_test" {
	#Required
	path_analyzer_test_id = oci_vn_monitoring_path_analyzer_test.test_path_analyzer_test.id
}
```

## Argument Reference

The following arguments are supported:

* `path_analyzer_test_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `PathAnalyzerTest` resource.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `PathAnalyzerTest` resource's compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `destination_endpoint` - Information describing a source or destination in a `PathAnalyzerTest` resource. 
	* `address` - The IPv4 address of the COMPUTE_INSTANCE-type `Endpoint` object. 
	* `instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute instance. 
	* `listener_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer listener. 
	* `load_balancer_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the listener's load balancer. 
	* `network_load_balancer_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the listener's network load balancer. 
	* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet containing the IP address. This can be used to disambiguate which subnet is intended, in case the IP address is used in more than one subnet (when there are subnets with overlapping IP ranges). 
	* `type` - The type of the `Endpoint`.
	* `vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN containing the IP address. This can be used to disambiguate which VLAN is queried, in case the endpoint IP address belongs to more than one VLAN (when there are VLANs with overlapping IP ranges). 
	* `vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC attached to the compute instance. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - A unique identifier established when the resource is created. The identifier can't be changed later. 
* `protocol` - The IP protocol to use for the `PathAnalyzerTest` resource.
* `protocol_parameters` - Defines the IP protocol parameters for a `PathAnalyzerTest` resource.
	* `destination_port` - The destination port to use in a `PathAnalyzerTest` resource.
	* `icmp_code` - The [ICMP](https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml) code.
	* `icmp_type` - The [ICMP](https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml) type.
	* `source_port` - The source port to use in a `PathAnalyzerTest` resource.
	* `type` - The type of the `ProtocolParameters` object.
* `query_options` - Defines the query options required for a `PathAnalyzerTest` resource.
	* `is_bi_directional_analysis` - If true, a path analysis is done for both the forward and reverse routes.
* `source_endpoint` - Information describing a source or destination in a `PathAnalyzerTest` resource. 
	* `address` - The IPv4 address of the COMPUTE_INSTANCE-type `Endpoint` object. 
	* `instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute instance. 
	* `listener_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer listener. 
	* `load_balancer_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the listener's load balancer. 
	* `network_load_balancer_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the listener's network load balancer. 
	* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet containing the IP address. This can be used to disambiguate which subnet is intended, in case the IP address is used in more than one subnet (when there are subnets with overlapping IP ranges). 
	* `type` - The type of the `Endpoint`.
	* `vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN containing the IP address. This can be used to disambiguate which VLAN is queried, in case the endpoint IP address belongs to more than one VLAN (when there are VLANs with overlapping IP ranges). 
	* `vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC attached to the compute instance. 
* `state` - The current state of the `PathAnalyzerTest` resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the `PathAnalyzerTest` resource was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_updated` - The date and time the `PathAnalyzerTest` resource was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

