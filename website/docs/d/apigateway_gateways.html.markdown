---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_gateways"
sidebar_current: "docs-oci-datasource-apigateway-gateways"
description: |-
  Provides the list of Gateways in Oracle Cloud Infrastructure API Gateway service
---

# Data Source: oci_apigateway_gateways
This data source provides the list of Gateways in Oracle Cloud Infrastructure API Gateway service.

Returns a list of gateways.


## Example Usage

```hcl
data "oci_apigateway_gateways" "test_gateways" {
	#Required
	compartment_id = var.compartment_id

	#Optional
    certificate_id = var.oci_apigateway_certificate.test_certificate.id
	display_name = var.gateway_display_name
	state = var.gateway_state
}
```

## Argument Reference

The following arguments are supported:

* `certificate_id` - (Optional) Filter gateways by the certificate ocid.
* `compartment_id` - (Required) The ocid of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state.  Example: `SUCCEEDED` 


## Attributes Reference

The following attributes are exported:

* `gateway_collection` - The list of gateway_collection.

### Gateway Reference

The following attributes are exported:

* `ca_bundles` - An array of CA bundles that should be used on the Gateway for TLS validation.
	* `ca_bundle_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
	* `certificate_authority_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
	* `type` - Type of the CA bundle
* `certificate_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `endpoint_type` - Gateway endpoint type. `PUBLIC` will have a public ip address assigned to it, while `PRIVATE` will only be accessible on a private IP address on the subnet.  Example: `PUBLIC` or `PRIVATE` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - The hostname for APIs deployed on the gateway.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `ip_addresses` - An array of IP addresses associated with the gateway.
	* `ip_address` - An IP address.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state. 
* `network_security_group_ids` - An array of Network Security Groups OCIDs associated with this API Gateway. 
* `response_cache_details` - Base Gateway response cache. 
	* `authentication_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
	* `authentication_secret_version_number` - The version number of the authentication secret to use. 
	* `connect_timeout_in_ms` - Defines the timeout for establishing a connection with the Response Cache. 
	* `is_ssl_enabled` - Defines if the connection should be over SSL. 
	* `is_ssl_verify_disabled` - Defines whether or not to uphold SSL verification. 
	* `read_timeout_in_ms` - Defines the timeout for reading data from the Response Cache. 
	* `send_timeout_in_ms` - Defines the timeout for transmitting data to the Response Cache. 
	* `servers` - The set of cache store members to connect to. At present only a single server is supported. 
		* `host` - Hostname or IP address (IPv4 only) where the cache store is running.
		* `port` - The port the cache store is exposed on.
	* `type` - Type of the Response Cache.
* `state` - The current state of the gateway.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet in which related resources are created. 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

