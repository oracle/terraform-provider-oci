---
subcategory: "Vn Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_vn_monitoring_path_analyzer_tests"
sidebar_current: "docs-oci-datasource-vn_monitoring-path_analyzer_tests"
description: |-
  Provides the list of Path Analyzer Tests in Oracle Cloud Infrastructure Vn Monitoring service
---

# Data Source: oci_vn_monitoring_path_analyzer_tests
This data source provides the list of Path Analyzer Tests in Oracle Cloud Infrastructure Vn Monitoring service.

Returns a list of all `PathAnalyzerTests` in a compartment.


## Example Usage

```hcl
data "oci_vn_monitoring_path_analyzer_tests" "test_path_analyzer_tests" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.path_analyzer_test_display_name
	state = var.path_analyzer_test_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter that returns only resources that match the entire display name given.
* `state` - (Optional) A filter that returns only resources whose `lifecycleState` matches the given `lifecycleState`.


## Attributes Reference

The following attributes are exported:

* `path_analyzer_test_collection` - The list of path_analyzer_test_collection.

### PathAnalyzerTest Reference

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

