---
subcategory: "Vn Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_vn_monitoring_path_analysi"
sidebar_current: "docs-oci-resource-vn_monitoring-path_analysi"
description: |-
  Provides the Path Analysi resource in Oracle Cloud Infrastructure Vn Monitoring service
---

# oci_vn_monitoring_path_analysi
This resource provides the Path Analysi resource in Oracle Cloud Infrastructure Vn Monitoring service.

Use this method to initiate a [Network Path Analyzer](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) analysis. This method returns
an opc-work-request-id, and you can poll the status of the work request until it either fails or succeeds.

If the work request status is successful, use [ListWorkRequestResults](https://docs.cloud.oracle.com/iaas/api/#/en/VnConfigAdvisor/latest/WorkRequestResult/ListWorkRequestResults)
with the work request ID to ask for the successful analysis results. If the work request status is failed, use
[ListWorkRequestErrors](https://docs.cloud.oracle.com/iaas/api/#/en/VnConfigAdvisor/latest/WorkRequestError/ListWorkRequestErrors)
with the work request ID to ask for the analysis failure information. The information 
returned from either of these methods can be used to build a final report. 


## Example Usage

```hcl
resource "oci_vn_monitoring_path_analysi" "test_path_analysi" {
	#Required
	type = var.path_analysi_type

	#Optional
	cache_control = var.path_analysi_cache_control
	compartment_id = var.compartment_id
	destination_endpoint {
		#Required
		type = var.path_analysi_destination_endpoint_type

		#Optional
		address = var.path_analysi_destination_endpoint_address
		instance_id = oci_core_instance.test_instance.id
		listener_id = oci_load_balancer_listener.test_listener.id
		load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
		network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
		subnet_id = oci_core_subnet.test_subnet.id
		vlan_id = oci_core_vlan.test_vlan.id
		vnic_id = oci_core_vnic_attachment.test_vnic_attachment.id
	}
	path_analyzer_test_id = oci_vn_monitoring_path_analyzer_test.test_path_analyzer_test.id
	protocol = var.path_analysi_protocol
	protocol_parameters {
		#Required
		type = var.path_analysi_protocol_parameters_type

		#Optional
		destination_port = var.path_analysi_protocol_parameters_destination_port
		icmp_code = var.path_analysi_protocol_parameters_icmp_code
		icmp_type = var.path_analysi_protocol_parameters_icmp_type
		source_port = var.path_analysi_protocol_parameters_source_port
	}
	query_options {

		#Optional
		is_bi_directional_analysis = var.path_analysi_query_options_is_bi_directional_analysis
	}
	source_endpoint {
		#Required
		type = var.path_analysi_source_endpoint_type

		#Optional
		address = var.path_analysi_source_endpoint_address
		instance_id = oci_core_instance.test_instance.id
		listener_id = oci_load_balancer_listener.test_listener.id
		load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
		network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
		subnet_id = oci_core_subnet.test_subnet.id
		vlan_id = oci_core_vlan.test_vlan.id
		vnic_id = oci_core_vnic_attachment.test_vnic_attachment.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `cache_control` - (Optional) The Cache-Control HTTP header holds directives (instructions) for caching in both requests and responses. 
* `compartment_id` - (Required when type=ADHOC_QUERY) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the compartment.
* `destination_endpoint` - (Required when type=ADHOC_QUERY) Information describing a source or destination in a `PathAnalyzerTest` resource. 
	* `address` - (Required when type=COMPUTE_INSTANCE | IP_ADDRESS | ON_PREM | SUBNET | VLAN | VNIC) The IPv4 address of the COMPUTE_INSTANCE-type `Endpoint` object. 
	* `instance_id` - (Required when type=COMPUTE_INSTANCE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute instance. 
	* `listener_id` - (Required when type=LOAD_BALANCER_LISTENER | NETWORK_LOAD_BALANCER_LISTENER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer listener. 
	* `load_balancer_id` - (Required when type=LOAD_BALANCER | LOAD_BALANCER_LISTENER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the listener's load balancer. 
	* `network_load_balancer_id` - (Required when type=NETWORK_LOAD_BALANCER | NETWORK_LOAD_BALANCER_LISTENER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the listener's network load balancer. 
	* `subnet_id` - (Required when type=SUBNET) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet containing the IP address. This can be used to disambiguate which subnet is intended, in case the IP address is used in more than one subnet (when there are subnets with overlapping IP ranges). 
	* `type` - (Required) The type of the `Endpoint`.
	* `vlan_id` - (Required when type=VLAN) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN containing the IP address. This can be used to disambiguate which VLAN is queried, in case the endpoint IP address belongs to more than one VLAN (when there are VLANs with overlapping IP ranges). 
	* `vnic_id` - (Required when type=COMPUTE_INSTANCE | VNIC) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC attached to the compute instance. 
* `path_analyzer_test_id` - (Required when type=PERSISTED_QUERY) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `PathAnalyzerTest` resource. 
* `protocol` - (Required when type=ADHOC_QUERY) The IP protocol to used for the path analysis.
* `protocol_parameters` - (Applicable when type=ADHOC_QUERY) Defines the IP protocol parameters for a `PathAnalyzerTest` resource.
	* `destination_port` - (Required when type=TCP | UDP) The destination port to use in a `PathAnalyzerTest` resource.
	* `icmp_code` - (Applicable when type=ICMP) The [ICMP](https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml) code.
	* `icmp_type` - (Required when type=ICMP) The [ICMP](https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml) type.
	* `source_port` - (Applicable when type=TCP | UDP) The source port to use in a `PathAnalyzerTest` resource.
	* `type` - (Required) The type of the `ProtocolParameters` object.
* `query_options` - (Applicable when type=ADHOC_QUERY) Defines the query options required for a `PathAnalyzerTest` resource.
	* `is_bi_directional_analysis` - (Applicable when type=ADHOC_QUERY) If true, a path analysis is done for both the forward and reverse routes.
* `source_endpoint` - (Required when type=ADHOC_QUERY) Information describing a source or destination in a `PathAnalyzerTest` resource. 
	* `address` - (Required when type=COMPUTE_INSTANCE | IP_ADDRESS | ON_PREM | SUBNET | VLAN | VNIC) The IPv4 address of the COMPUTE_INSTANCE-type `Endpoint` object. 
	* `instance_id` - (Required when type=COMPUTE_INSTANCE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute instance. 
	* `listener_id` - (Required when type=LOAD_BALANCER_LISTENER | NETWORK_LOAD_BALANCER_LISTENER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer listener. 
	* `load_balancer_id` - (Required when type=LOAD_BALANCER | LOAD_BALANCER_LISTENER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the listener's load balancer. 
	* `network_load_balancer_id` - (Required when type=NETWORK_LOAD_BALANCER | NETWORK_LOAD_BALANCER_LISTENER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the listener's network load balancer. 
	* `subnet_id` - (Required when type=SUBNET) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet containing the IP address. This can be used to disambiguate which subnet is intended, in case the IP address is used in more than one subnet (when there are subnets with overlapping IP ranges). 
	* `type` - (Required) The type of the `Endpoint`.
	* `vlan_id` - (Required when type=VLAN) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN containing the IP address. This can be used to disambiguate which VLAN is queried, in case the endpoint IP address belongs to more than one VLAN (when there are VLANs with overlapping IP ranges). 
	* `vnic_id` - (Required when type=COMPUTE_INSTANCE | VNIC) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC attached to the compute instance. 
* `type` - (Required) The type of the `PathAnalysis` query.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Path Analysi
	* `update` - (Defaults to 20 minutes), when updating the Path Analysi
	* `delete` - (Defaults to 20 minutes), when destroying the Path Analysi


## Import

PathAnalysis can be imported using the `id`, e.g.

```
$ terraform import oci_vn_monitoring_path_analysi.test_path_analysi "id"
```

