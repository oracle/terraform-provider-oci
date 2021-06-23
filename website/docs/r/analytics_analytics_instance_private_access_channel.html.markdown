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
}
```

## Argument Reference

The following arguments are supported:

* `analytics_instance_id` - (Required) The OCID of the AnalyticsInstance. 
* `display_name` - (Required) (Updatable) Display Name of the Private Access Channel. 
* `private_source_dns_zones` - (Required) (Updatable) List of Private Source DNS zones registered with Private Access Channel, where datasource hostnames from these dns zones / domains will be resolved in the peered VCN for access from Analytics Instance. Min of 1 is required and Max of 30 Private Source DNS zones can be registered. 
	* `description` - (Optional) (Updatable) Description of private source dns zone. 
	* `dns_zone` - (Required) (Updatable) Private Source DNS Zone. Ex: example-vcn.oraclevcn.com, corp.example.com. 
* `subnet_id` - (Required) (Updatable) OCID of the customer subnet connected to private access channel. 
* `vcn_id` - (Required) (Updatable) OCID of the customer VCN peered with private access channel. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Analytics Instance Private Access Channel
	* `update` - (Defaults to 20 minutes), when updating the Analytics Instance Private Access Channel
	* `delete` - (Defaults to 20 minutes), when destroying the Analytics Instance Private Access Channel


## Import

AnalyticsInstancePrivateAccessChannels can be imported using the `id`, e.g.

```
$ terraform import oci_analytics_analytics_instance_private_access_channel.test_analytics_instance_private_access_channel "analyticsInstances/{analyticsInstanceId}/privateAccessChannels/{privateAccessChannelKey}" 
```

