---
subcategory: "Vn Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_vn_monitoring_path_analyzer_test"
sidebar_current: "docs-oci-resource-vn_monitoring-path_analyzer_test"
description: |-
  Provides the Path Analyzer Test resource in Oracle Cloud Infrastructure Vn Monitoring service
---

# oci_vn_monitoring_path_analyzer_test
This resource provides the Path Analyzer Test resource in Oracle Cloud Infrastructure Vn Monitoring service.

Creates a new `PathAnalyzerTest` resource.

## Example Usage

```hcl
resource "oci_vn_monitoring_path_analyzer_test" "test_path_analyzer_test" {
	#Required
	compartment_id = var.compartment_id
	destination_endpoint {
		#Required
		type = var.path_analyzer_test_destination_endpoint_type

		#Optional
		address = var.path_analyzer_test_destination_endpoint_address
		instance_id = oci_core_instance.test_instance.id
		listener_id = oci_load_balancer_listener.test_listener.id
		load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
		network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
		subnet_id = oci_core_subnet.test_subnet.id
		vlan_id = oci_core_vlan.test_vlan.id
		vnic_id = oci_core_vnic_attachment.test_vnic_attachment.id
	}
	protocol = var.path_analyzer_test_protocol
	source_endpoint {
		#Required
		type = var.path_analyzer_test_source_endpoint_type

		#Optional
		address = var.path_analyzer_test_source_endpoint_address
		instance_id = oci_core_instance.test_instance.id
		listener_id = oci_load_balancer_listener.test_listener.id
		load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
		network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
		subnet_id = oci_core_subnet.test_subnet.id
		vlan_id = oci_core_vlan.test_vlan.id
		vnic_id = oci_core_vnic_attachment.test_vnic_attachment.id
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.path_analyzer_test_display_name
	freeform_tags = {"bar-key"= "value"}
	protocol_parameters {
		#Required
		type = var.path_analyzer_test_protocol_parameters_type

		#Optional
		destination_port = var.path_analyzer_test_protocol_parameters_destination_port
		icmp_code = var.path_analyzer_test_protocol_parameters_icmp_code
		icmp_type = var.path_analyzer_test_protocol_parameters_icmp_type
		source_port = var.path_analyzer_test_protocol_parameters_source_port
	}
	query_options {

		#Optional
		is_bi_directional_analysis = var.path_analyzer_test_query_options_is_bi_directional_analysis
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the `PathAnalyzerTest` resource's compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `destination_endpoint` - (Required) (Updatable) Information describing a source or destination in a `PathAnalyzerTest` resource. 
	* `address` - (Required when type=COMPUTE_INSTANCE | IP_ADDRESS | ON_PREM | SUBNET | VLAN | VNIC) (Updatable) The IPv4 address of the COMPUTE_INSTANCE-type `Endpoint` object. 
	* `instance_id` - (Required when type=COMPUTE_INSTANCE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute instance. 
	* `listener_id` - (Required when type=LOAD_BALANCER_LISTENER | NETWORK_LOAD_BALANCER_LISTENER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer listener. 
	* `load_balancer_id` - (Required when type=LOAD_BALANCER | LOAD_BALANCER_LISTENER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the listener's load balancer. 
	* `network_load_balancer_id` - (Required when type=NETWORK_LOAD_BALANCER | NETWORK_LOAD_BALANCER_LISTENER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the listener's network load balancer. 
	* `subnet_id` - (Required when type=SUBNET) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet containing the IP address. This can be used to disambiguate which subnet is intended, in case the IP address is used in more than one subnet (when there are subnets with overlapping IP ranges). 
	* `type` - (Required) (Updatable) The type of the `Endpoint`.
	* `vlan_id` - (Required when type=VLAN) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN containing the IP address. This can be used to disambiguate which VLAN is queried, in case the endpoint IP address belongs to more than one VLAN (when there are VLANs with overlapping IP ranges). 
	* `vnic_id` - (Required when type=COMPUTE_INSTANCE | VNIC) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC attached to the compute instance. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `protocol` - (Required) (Updatable) The IP protocol to use in the `PathAnalyzerTest` resource.
* `protocol_parameters` - (Optional) (Updatable) Defines the IP protocol parameters for a `PathAnalyzerTest` resource.
	* `destination_port` - (Required when type=TCP | UDP) (Updatable) The destination port to use in a `PathAnalyzerTest` resource.
	* `icmp_code` - (Applicable when type=ICMP) (Updatable) The [ICMP](https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml) code.
	* `icmp_type` - (Required when type=ICMP) (Updatable) The [ICMP](https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml) type.
	* `source_port` - (Applicable when type=TCP | UDP) (Updatable) The source port to use in a `PathAnalyzerTest` resource.
	* `type` - (Required) (Updatable) The type of the `ProtocolParameters` object.
* `query_options` - (Optional) (Updatable) Defines the query options required for a `PathAnalyzerTest` resource.
	* `is_bi_directional_analysis` - (Optional) (Updatable) If true, a path analysis is done for both the forward and reverse routes.
* `source_endpoint` - (Required) (Updatable) Information describing a source or destination in a `PathAnalyzerTest` resource. 
	* `address` - (Required when type=COMPUTE_INSTANCE | IP_ADDRESS | ON_PREM | SUBNET | VLAN | VNIC) (Updatable) The IPv4 address of the COMPUTE_INSTANCE-type `Endpoint` object. 
	* `instance_id` - (Required when type=COMPUTE_INSTANCE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute instance. 
	* `listener_id` - (Required when type=LOAD_BALANCER_LISTENER | NETWORK_LOAD_BALANCER_LISTENER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer listener. 
	* `load_balancer_id` - (Required when type=LOAD_BALANCER | LOAD_BALANCER_LISTENER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the listener's load balancer. 
	* `network_load_balancer_id` - (Required when type=NETWORK_LOAD_BALANCER | NETWORK_LOAD_BALANCER_LISTENER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the listener's network load balancer. 
	* `subnet_id` - (Required when type=SUBNET) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet containing the IP address. This can be used to disambiguate which subnet is intended, in case the IP address is used in more than one subnet (when there are subnets with overlapping IP ranges). 
	* `type` - (Required) (Updatable) The type of the `Endpoint`.
	* `vlan_id` - (Required when type=VLAN) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN containing the IP address. This can be used to disambiguate which VLAN is queried, in case the endpoint IP address belongs to more than one VLAN (when there are VLANs with overlapping IP ranges). 
	* `vnic_id` - (Required when type=COMPUTE_INSTANCE | VNIC) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC attached to the compute instance. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Path Analyzer Test
	* `update` - (Defaults to 20 minutes), when updating the Path Analyzer Test
	* `delete` - (Defaults to 20 minutes), when destroying the Path Analyzer Test


## Import

PathAnalyzerTests can be imported using the `id`, e.g.

```
$ terraform import oci_vn_monitoring_path_analyzer_test.test_path_analyzer_test "id"
```

