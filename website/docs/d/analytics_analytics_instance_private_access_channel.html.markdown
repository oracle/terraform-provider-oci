---
subcategory: "Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_analytics_analytics_instance_private_access_channel"
sidebar_current: "docs-oci-datasource-analytics-analytics_instance_private_access_channel"
description: |-
  Provides details about a specific Analytics Instance Private Access Channel in Oracle Cloud Infrastructure Analytics service
---

# Data Source: oci_analytics_analytics_instance_private_access_channel
This data source provides details about a specific Analytics Instance Private Access Channel resource in Oracle Cloud Infrastructure Analytics service.

Retrieve private access channel in the specified Analytics Instance.


## Example Usage

```hcl
data "oci_analytics_analytics_instance_private_access_channel" "test_analytics_instance_private_access_channel" {
	#Required
	analytics_instance_id = oci_analytics_analytics_instance.test_analytics_instance.id
	private_access_channel_key = var.analytics_instance_private_access_channel_private_access_channel_key
}
```

## Argument Reference

The following arguments are supported:

* `analytics_instance_id` - (Required) The OCID of the AnalyticsInstance. 
* `private_access_channel_key` - (Required) The unique identifier key of the Private Access Channel. 


## Attributes Reference

The following attributes are exported:

* `display_name` - Display Name of the Private Access Channel. 
* `egress_source_ip_addresses` - The list of IP addresses from the customer subnet connected to private access channel, used as a source Ip by Private Access Channel for network traffic from the AnalyticsInstance to Private Sources. 
* `ip_address` - IP Address of the Private Access channel. 
* `key` - Private Access Channel unique identifier key. 
* `private_source_dns_zones` - List of Private Source DNS zones registered with Private Access Channel, where datasource hostnames from these dns zones / domains will be resolved in the peered VCN for access from Analytics Instance. Min of 1 is required and Max of 30 Private Source DNS zones can be registered. 
	* `description` - Description of private source dns zone. 
	* `dns_zone` - Private Source DNS Zone. Ex: example-vcn.oraclevcn.com, corp.example.com. 
* `subnet_id` - OCID of the customer subnet connected to private access channel. 
* `vcn_id` - OCID of the customer VCN peered with private access channel. 

