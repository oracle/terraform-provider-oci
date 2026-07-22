---
subcategory: "Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_analytics_analytics_instance_private_access_channel"
sidebar_current: "docs-oci-resource-analytics-analytics_instance_private_access_channel"
description: |-
  Provides the Analytics Instance Private Access Channel resource in Oracle Cloud Infrastructure Analytics service
---

# oci_analytics_analytics_instance_private_access_channel
This resource provides the Analytics Instance Private Access Channel resource in Oracle Cloud Infrastructure Analytics service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/analytics/latest/AnalyticsInstancePrivateAccessChannel

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/analytics

Create a Private access Channel for the Analytics instance. The operation is long-running
and creates a new WorkRequest.


## Example Usage

```hcl
resource "oci_analytics_analytics_instance_private_access_channel" "test_analytics_instance_private_access_channel" {
	#Required
	analytics_instance_id = oci_analytics_analytics_instance.test_analytics_instance.id
	display_name = var.analytics_instance_private_access_channel_display_name
	private_source_dns_zones {
		#Required
		dns_zone = var.analytics_instance_private_access_channel_private_source_dns_zones_dns_zone

		#Optional
		description = var.analytics_instance_private_access_channel_private_source_dns_zones_description
	}
	subnet_id = oci_core_subnet.test_subnet.id
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	network_security_group_ids = var.analytics_instance_private_access_channel_network_security_group_ids
	private_source_scan_hosts {
		#Required
		scan_hostname = var.analytics_instance_private_access_channel_private_source_scan_hosts_scan_hostname
		scan_port = var.analytics_instance_private_access_channel_private_source_scan_hosts_scan_port

		#Optional
		description = var.analytics_instance_private_access_channel_private_source_scan_hosts_description
	}
}
```

## Argument Reference

The following arguments are supported:

* `analytics_instance_id` - (Required) The OCID of the Analytics instance. 
* `display_name` - (Required) (Updatable) Display name of the private access channel. 
* `network_security_group_ids` - (Optional) (Updatable) Network Security Group OCIDs for the Analytics instance. 
* `private_source_dns_zones` - (Required) (Updatable) List of private source DNS zones registered with the private access channel. The datasource hostnames from these DNS zones / domains will be resolved in the peered VCN for access from  the Analytics instance. Minimum 1 private source is required. Maximum 30 private source DNS zones can be registered. 
	* `description` - (Optional) (Updatable) Description of the private source DNS zone. 
	* `dns_zone` - (Required) (Updatable) Private source DNS zone. For example: example-vcn.oraclevcn.com, corp.example.com. 
* `private_source_scan_hosts` - (Optional) (Updatable) List of private source database SCAN hosts registered with the private access channel for access from the Analytics instance. 
	* `description` - (Optional) (Updatable) Description of private source SCAN host zone. 
	* `scan_hostname` - (Required) (Updatable) Private source SCAN hostname. For example: db01-scan.corp.example.com, prd-db01-scan.mycompany.com. 
	* `scan_port` - (Required) (Updatable) Private source SCAN host port. This is the source port where the SCAN protocol connects (for example, 1521). 
* `subnet_id` - (Required) (Updatable) OCID of the customer subnet connected to the private access channel. 
* `vcn_id` - (Required) (Updatable) OCID of the customer VCN peered with the private access channel. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Analytics Instance Private Access Channel
	* `update` - (Defaults to 20 minutes), when updating the Analytics Instance Private Access Channel
	* `delete` - (Defaults to 20 minutes), when destroying the Analytics Instance Private Access Channel


## Import

AnalyticsInstancePrivateAccessChannels can be imported using the `id`, e.g.

```
$ terraform import oci_analytics_analytics_instance_private_access_channel.test_analytics_instance_private_access_channel "analyticsInstances/{analyticsInstanceId}/privateAccessChannels/{privateAccessChannelKey}" 
```

