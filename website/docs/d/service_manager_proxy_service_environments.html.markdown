---
subcategory: "Service Manager Proxy"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_manager_proxy_service_environments"
sidebar_current: "docs-oci-datasource-service_manager_proxy-service_environments"
description: |-
  Provides the list of Service Environments in Oracle Cloud Infrastructure Service Manager Proxy service
---

# Data Source: oci_service_manager_proxy_service_environments
This data source provides the list of Service Environments in Oracle Cloud Infrastructure Service Manager Proxy service.

List the details of Software as a Service (SaaS) environments provisioned by Service Manager.
Information includes the service instance endpoints and service definition details.


## Example Usage

```hcl
data "oci_service_manager_proxy_service_environments" "test_service_environments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.service_environment_display_name
	service_environment_id = oci_service_manager_proxy_service_environment.test_service_environment.id
	service_environment_type = var.service_environment_service_environment_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the compartment.
* `display_name` - (Optional) The display name of the resource.
* `service_environment_id` - (Optional) The unique identifier associated with the service environment. 

	**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `service_environment_type` - (Optional) The environment's service definition type.  For example, "RGBUOROMS" is the service definition type for "Oracle Retail Order Management Cloud Service". 


## Attributes Reference

The following attributes are exported:

* `service_environment_collection` - The list of service_environment_collection.

### ServiceEnvironment Reference

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

