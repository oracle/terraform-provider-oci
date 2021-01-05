---
subcategory: "Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_analytics_analytics_instance"
sidebar_current: "docs-oci-resource-analytics-analytics_instance"
description: |-
  Provides the Analytics Instance resource in Oracle Cloud Infrastructure Analytics service
---

# oci_analytics_analytics_instance
This resource provides the Analytics Instance resource in Oracle Cloud Infrastructure Analytics service.

Create a new AnalyticsInstance in the specified compartment. The operation is long-running
and creates a new WorkRequest.


## Example Usage

```hcl
resource "oci_analytics_analytics_instance" "test_analytics_instance" {
	#Required
	capacity {
		#Required
		capacity_type = var.analytics_instance_capacity_capacity_type
		capacity_value = var.analytics_instance_capacity_capacity_value
	}
	compartment_id = var.compartment_id
	feature_set = var.analytics_instance_feature_set
	idcs_access_token = var.analytics_instance_idcs_access_token
	license_type = var.analytics_instance_license_type
	name = var.analytics_instance_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.analytics_instance_description
	email_notification = var.analytics_instance_email_notification
	freeform_tags = {"Department"= "Finance"}
	network_endpoint_details {
		#Required
		network_endpoint_type = var.analytics_instance_network_endpoint_details_network_endpoint_type

		#Optional
		subnet_id = oci_core_subnet.test_subnet.id
		vcn_id = oci_core_vcn.test_vcn.id
		whitelisted_ips = var.analytics_instance_network_endpoint_details_whitelisted_ips
		whitelisted_vcns {

			#Optional
			id = var.analytics_instance_network_endpoint_details_whitelisted_vcns_id
			whitelisted_ips = var.analytics_instance_network_endpoint_details_whitelisted_vcns_whitelisted_ips
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `capacity` - (Required) Service instance capacity metadata (e.g.: OLPU count, number of users, ...etc...). 
	* `capacity_type` - (Required) The capacity model to use. 
	* `capacity_value` - (Required) (Updatable) The capacity value selected (OLPU count, number of users, ...etc...). This parameter affects the number of CPUs, amount of memory or other resources allocated to the instance. 
* `compartment_id` - (Required) (Updatable) The OCID of the compartment. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Optional description. 
* `email_notification` - (Optional) (Updatable) Email address receiving notifications. 
* `feature_set` - (Required) Analytics feature set. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `idcs_access_token` - (Required) IDCS access token identifying a stripe and service administrator user. 
* `license_type` - (Required) (Updatable) The license used for the service. 
* `name` - (Required) The name of the Analytics instance. This name must be unique in the tenancy and cannot be changed. 
* `state` - (Optional) (Updatable) The target state for the instance. Could be set to `ACTIVE` or `INACTIVE`. 
* `network_endpoint_details` - (Optional) Base representation of a network endpoint. 
	* `network_endpoint_type` - (Required) The type of network endpoint. 
	* `subnet_id` - (Required when network_endpoint_type=PRIVATE) The subnet OCID for the private endpoint. 
	* `vcn_id` - (Required when network_endpoint_type=PRIVATE) The VCN OCID for the private endpoint. 
	* `whitelisted_ips` - (Applicable when network_endpoint_type=PUBLIC) Source IP addresses or IP address ranges igress rules. 
	* `whitelisted_vcns` - (Applicable when network_endpoint_type=PUBLIC) Virtual Cloud Networks allowed to access this network endpoint. 
		* `id` - (Required when network_endpoint_type=PUBLIC) The Virtual Cloud Network OCID. 
		* `whitelisted_ips` - (Applicable when network_endpoint_type=PUBLIC) Source IP addresses or IP address ranges igress rules. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `capacity` - Service instance capacity metadata (e.g.: OLPU count, number of users, ...etc...). 
	* `capacity_type` - The capacity model to use. 
	* `capacity_value` - The capacity value selected (OLPU count, number of users, ...etc...). This parameter affects the number of CPUs, amount of memory or other resources allocated to the instance. 
* `compartment_id` - The OCID of the compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Optional description. 
* `email_notification` - Email address receiving notifications. 
* `feature_set` - Analytics feature set. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The resource OCID. 
* `license_type` - The license used for the service. 
* `name` - The name of the Analytics instance. This name must be unique in the tenancy and cannot be changed. 
* `network_endpoint_details` - Base representation of a network endpoint. 
	* `network_endpoint_type` - The type of network endpoint. 
	* `subnet_id` - The subnet OCID for the private endpoint. 
	* `vcn_id` - The VCN OCID for the private endpoint. 
	* `whitelisted_ips` - Source IP addresses or IP address ranges igress rules. 
	* `whitelisted_vcns` - Virtual Cloud Networks allowed to access this network endpoint. 
		* `id` - The Virtual Cloud Network OCID. 
		* `whitelisted_ips` - Source IP addresses or IP address ranges igress rules. 
* `service_url` - URL of the Analytics service. 
* `state` - The current state of an instance. 
* `time_created` - The date and time the instance was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the instance was last updated (in the format defined by RFC3339). This timestamp represents updates made through this API. External events do not influence it. 

## Import

AnalyticsInstances can be imported using the `id`, e.g.

```
$ terraform import oci_analytics_analytics_instance.test_analytics_instance "id"
```

