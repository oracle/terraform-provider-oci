---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_agent_configuration"
sidebar_current: "docs-oci-datasource-jms-fleet_agent_configuration"
description: |-
  Provides details about a specific Fleet Agent Configuration in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_agent_configuration
This data source provides details about a specific Fleet Agent Configuration resource in Oracle Cloud Infrastructure Jms service.

Retrieve a Fleet Agent Configuration for the specified Fleet.

## Example Usage

```hcl
data "oci_jms_fleet_agent_configuration" "test_fleet_agent_configuration" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.


## Attributes Reference

The following attributes are exported:

* `agent_polling_interval_in_minutes` - Agent polling interval in minutes 
* `is_capturing_ip_address_and_fqdn_enabled` - Collect network addresses of managed instances in the fleet. 
* `is_collecting_managed_instance_metrics_enabled` - Collect JMS agent metrics on all managed instances in the fleet. 
* `is_collecting_usernames_enabled` - Collect username for application invocations for all managed instances in the fleet. 
* `is_libraries_scan_enabled` - Enable libraries scan on all managed instances in the fleet. 
* `java_usage_tracker_processing_frequency_in_minutes` - The frequency (in minutes) of Java Usage Tracker processing. (That is, how often should JMS process data from the Java Usage Tracker.) 
* `jre_scan_frequency_in_minutes` - The frequency (in minutes) of JRE scanning. (That is, how often should JMS scan for JRE installations.) 
* `linux_configuration` - Management Agent Configuration for list of include/exclude file system paths (specific to operating system). 
	* `exclude_paths` - An array of file system paths (environment variables supported). 
	* `include_paths` - An array of file system paths (environment variables supported). 
* `mac_os_configuration` - Management Agent Configuration for list of include/exclude file system paths (specific to operating system). 
	* `exclude_paths` - An array of file system paths (environment variables supported). 
	* `include_paths` - An array of file system paths (environment variables supported). 
* `time_last_modified` - The date and time of the last modification to the Fleet Agent Configuration (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 
* `windows_configuration` - Management Agent Configuration for list of include/exclude file system paths (specific to operating system). 
	* `exclude_paths` - An array of file system paths (environment variables supported). 
	* `include_paths` - An array of file system paths (environment variables supported). 
* `work_request_validity_period_in_days` - The validity period in days for work requests. 

