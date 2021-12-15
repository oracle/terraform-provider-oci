---
subcategory: "Service Manager Proxy"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_manager_proxy_service_environment"
sidebar_current: "docs-oci-datasource-service_manager_proxy-service_environment"
description: |-
  Provides details about a specific Service Environment in Oracle Cloud Infrastructure Service Manager Proxy service
---

# Data Source: oci_service_manager_proxy_service_environment
This data source provides details about a specific Service Environment resource in Oracle Cloud Infrastructure Service Manager Proxy service.

Get the detailed information for a specific service environment.


## Example Usage

```hcl
data "oci_service_manager_proxy_service_environment" "test_service_environment" {
	#Required
	compartment_id = var.compartment_id
	service_environment_id = oci_service_manager_proxy_service_environment.test_service_environment.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the compartment.
* `service_environment_id` - (Required) The unique identifier associated with the service environment. 

	**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the compartment.
* `console_url` - The URL for the console.
* `id` - Unqiue identifier for the entitlement related to the environment. 

	**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `service_definition` - Details for a service definition. 
	* `display_name` - Display name of the service. For example, "Oracle Retail Order Management Cloud Service". 
	* `short_display_name` - Short display name of the service. For example, "Retail Order Management". 
	* `type` - The service definition type. For example, a service definition type "RGBUOROMS"  would be for the service "Oracle Retail Order Management Cloud Service". 
* `service_environment_endpoints` - Array of service environment end points.
	* `description` - Description of the environment link
	* `environment_type` - Service environment endpoint type.
	* `url` - Service environment instance URL.
* `status` - Status of the entitlement registration for the service.
* `subscription_id` - The unique subscription ID associated with the service environment ID.

	**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 

