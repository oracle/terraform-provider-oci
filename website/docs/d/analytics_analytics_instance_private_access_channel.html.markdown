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

Retrieve private access channel for the specified Analytics Instance.


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

* `analytics_instance_id` - (Required) The OCID of the Analytics instance. 
* `private_access_channel_key` - (Required) The unique identifier key of the private access channel. 


## Attributes Reference

The following attributes are exported:

* `display_name` - Display name of the private access channel. 
* `egress_source_ip_addresses` - List of IP addresses from the customer subnet connected to the private access channel, used as a source IP by the private access channel for network traffic from the Analytics instance to the private sources. 
* `ip_address` - IP address of the private access channel. 
* `key` - Private access channel unique identifier key. 
* `network_security_group_ids` - Network Security Group OCIDs for the Analytics instance. 
* `private_source_dns_zones` - List of private source DNS zones registered with the private access channel. The datasource hostnames from these DNS zones / domains will be resolved in the peered VCN for access from  the Analytics instance. Minimum 1 private source is required. Maximum 30 private source DNS zones can be registered. 
	* `description` - Description of the private source DNS zone. 
	* `dns_zone` - Private source DNS zone. For example: example-vcn.oraclevcn.com, corp.example.com. 
* `private_source_scan_hosts` - List of private source database SCAN hosts registered with the private access channel for access from the Analytics instance. 
	* `description` - Description of private source SCAN host zone. 
	* `scan_hostname` - Private source SCAN hostname. For example: db01-scan.corp.example.com, prd-db01-scan.mycompany.com. 
	* `scan_port` - Private source SCAN host port. This is the source port where the SCAN protocol connects (for example, 1521). 
* `subnet_id` - OCID of the customer subnet connected to the private access channel. 
* `vcn_id` - OCID of the customer VCN peered with the private access channel. 

