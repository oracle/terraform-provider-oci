---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_service_gateways"
sidebar_current: "docs-oci-datasource-core-service_gateways"
description: |-
  Provides the list of Service Gateways in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_service_gateways
This data source provides the list of Service Gateways in Oracle Cloud Infrastructure Core service.

Lists the service gateways in the specified compartment. You may optionally specify a VCN OCID
to filter the results by VCN.


## Example Usage

```hcl
data "oci_core_service_gateways" "test_service_gateways" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	state = "${var.service_gateway_state}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state.  The state value is case-insensitive. 
* `vcn_id` - (Optional) The OCID of the VCN.


## Attributes Reference

The following attributes are exported:

* `service_gateways` - The list of service_gateways.

### ServiceGateway Reference

The following attributes are exported:

* `block_traffic` - Whether the service gateway blocks all traffic through it. The default is `false`. When this is `true`, traffic is not routed to any services, regardless of route rules.  Example: `true` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the service gateway. 
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "foo-value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the service gateway.
* `services` - List of the services enabled on this service gateway. The list can be empty. You can enable a particular service by using [AttachServiceId](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/ServiceGateway/AttachServiceId). 
	* `service_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the service. 
	* `service_name` - The name of the service.
* `state` - The service gateway's current state.
* `time_created` - The date and time the service gateway was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the service gateway belongs to. 

