---
subcategory: "Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_analytics_analytics_instances"
sidebar_current: "docs-oci-datasource-analytics-analytics_instances"
description: |-
  Provides the list of Analytics Instances in Oracle Cloud Infrastructure Analytics service
---

# Data Source: oci_analytics_analytics_instances
This data source provides the list of Analytics Instances in Oracle Cloud Infrastructure Analytics service.

List Analytics instances.


## Example Usage

```hcl
data "oci_analytics_analytics_instances" "test_analytics_instances" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	capacity_type = var.analytics_instance_capacity_type
	feature_set = var.analytics_instance_feature_set
	name = var.analytics_instance_name
	state = var.analytics_instance_state
}
```

## Argument Reference

The following arguments are supported:

* `capacity_type` - (Optional) A filter to only return resources matching the capacity type enum. Values are case-insensitive. 
* `compartment_id` - (Required) The OCID of the compartment. 
* `feature_set` - (Optional) A filter to only return resources matching the feature set. Values are case-insensitive. 
* `name` - (Optional) A filter to return only resources that match the given name exactly. 
* `state` - (Optional) A filter to only return resources matching the lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `analytics_instances` - The list of analytics_instances.

### AnalyticsInstance Reference

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

