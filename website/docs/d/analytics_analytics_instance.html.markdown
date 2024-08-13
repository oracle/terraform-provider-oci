---
subcategory: "Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_analytics_analytics_instance"
sidebar_current: "docs-oci-datasource-analytics-analytics_instance"
description: |-
  Provides details about a specific Analytics Instance in Oracle Cloud Infrastructure Analytics service
---

# Data Source: oci_analytics_analytics_instance
This data source provides details about a specific Analytics Instance resource in Oracle Cloud Infrastructure Analytics service.

Info for a specific Analytics instance.


## Example Usage

```hcl
data "oci_analytics_analytics_instance" "test_analytics_instance" {
	#Required
	analytics_instance_id = oci_analytics_analytics_instance.test_analytics_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `analytics_instance_id` - (Required) The OCID of the AnalyticsInstance. 


## Attributes Reference

The following attributes are exported:

* `capacity` - Service instance capacity metadata (e.g.: OLPU count, number of users, ...etc...). 
	* `capacity_type` - The capacity model to use. Accepted values are: OLPU_COUNT, USER_COUNT 
	* `capacity_value` - The capacity value selected, either the number of OCPUs (OLPU_COUNT) or the number of users (USER_COUNT). This parameter affects the number of OCPUs, amount of memory, and other resources allocated to the instance. 
* `compartment_id` - The OCID of the compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Optional description. 
* `domain_id` - Identity domain OCID. 
* `email_notification` - Email address receiving notifications. 
* `feature_bundle` - The feature set of an Analytics instance. 
* `feature_set` - Analytics feature set. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The resource OCID. 
* `kms_key_id` - OCID of the Oracle Cloud Infrastructure Vault Key encrypting the customer data stored in this Analytics instance. A null value indicates Oracle managed default encryption. 
* `license_type` - The license used for the service. 
* `name` - The name of the Analytics instance. This name must be unique in the tenancy and cannot be changed. 
* `network_endpoint_details` - Base representation of a network endpoint. 
	* `network_endpoint_type` - The type of network endpoint. 
	* `network_security_group_ids` - Network Security Group OCIDs for an Analytics instance. 
	* `subnet_id` - The subnet OCID for the private endpoint. 
	* `vcn_id` - The VCN OCID for the private endpoint. 
	* `whitelisted_ips` - Source IP addresses or IP address ranges in ingress rules. 
	* `whitelisted_services` - Oracle Cloud Services that are allowed to access this Analytics instance. 
	* `whitelisted_vcns` - Virtual Cloud Networks allowed to access this network endpoint. 
		* `id` - The Virtual Cloud Network OCID. 
		* `whitelisted_ips` - Source IP addresses or IP address ranges in ingress rules. 
* `private_access_channels` - Map of PrivateAccessChannel unique identifier key as KEY and PrivateAccessChannel Object as VALUE. 
	* `display_name` - Display Name of the Private Access Channel. 
	* `egress_source_ip_addresses` - The list of IP addresses from the customer subnet connected to private access channel, used as a source Ip by Private Access Channel for network traffic from the AnalyticsInstance to Private Sources. 
	* `ip_address` - IP Address of the Private Access channel. 
	* `key` - Private Access Channel unique identifier key. 
	* `network_security_group_ids` - Network Security Group OCIDs for an Analytics instance. 
	* `private_source_dns_zones` - List of Private Source DNS zones registered with Private Access Channel, where datasource hostnames from these dns zones / domains will be resolved in the peered VCN for access from Analytics Instance. Min of 1 is required and Max of 30 Private Source DNS zones can be registered. 
		* `description` - Description of private source dns zone. 
		* `dns_zone` - Private Source DNS Zone. Ex: example-vcn.oraclevcn.com, corp.example.com. 
	* `private_source_scan_hosts` - List of Private Source DB SCAN hosts registered with Private Access Channel for access from Analytics Instance. 
		* `description` - Description of private source scan host zone. 
		* `scan_hostname` - Private Source Scan hostname. Ex: db01-scan.corp.example.com, prd-db01-scan.mycompany.com. 
		* `scan_port` - Private Source Scan host port. This is the source port where SCAN protocol will get connected (e.g. 1521). 
	* `subnet_id` - OCID of the customer subnet connected to private access channel. 
	* `vcn_id` - OCID of the customer VCN peered with private access channel. 
* `service_url` - URL of the Analytics service. 
* `state` - The current state of an instance. 
* `system_tags` - System tags for this resource. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.key": "value"}` 
* `time_created` - The date and time the instance was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the instance was last updated (in the format defined by RFC3339). This timestamp represents updates made through this API. External events do not influence it. 
* `update_channel` - Analytics instance update channel. 
* `vanity_url_details` - Map of VanityUrl unique identifier key as KEY and VanityUrl Object as VALUE. 
	* `description` - Description of the vanity url. 
	* `hosts` - List of fully qualified hostnames supported by this vanity URL definition (max of 3). 
	* `key` - The vanity url unique identifier key. 
	* `public_certificate` - PEM certificate for HTTPS connections. 
	* `urls` - List of urls supported by this vanity URL definition (max of 3). 

