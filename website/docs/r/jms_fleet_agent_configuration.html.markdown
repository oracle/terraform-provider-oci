---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_agent_configuration"
sidebar_current: "docs-oci-resource-jms-fleet_agent_configuration"
description: |-
  Provides the Fleet Agent Configuration resource in Oracle Cloud Infrastructure Jms service
---

# oci_jms_fleet_agent_configuration
This resource provides the Fleet Agent Configuration resource in Oracle Cloud Infrastructure Jms service.

Update the Fleet Agent Configuration for the specified Fleet.

## Example Usage

```hcl
resource "oci_jms_fleet_agent_configuration" "test_fleet_agent_configuration" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id

	#Optional
  agent_polling_interval_in_minutes                  = 10
  is_capturing_ip_address_and_fqdn_enabled           = false
  is_collecting_managed_instance_metrics_enabled     = false
  is_collecting_usernames_enabled                    = false
  is_libraries_scan_enabled                          = false
  java_usage_tracker_processing_frequency_in_minutes = 10
  jre_scan_frequency_in_minutes                      = 180  # must be >= 180
  linux_configuration {
    #Required
    exclude_paths = ["/user/private1", "/opt/private1"]
    include_paths = ["/user", "/opt"]
  }
  mac_os_configuration {
    #Required
    exclude_paths = ["/home/private1"]
    include_paths = ["/home"]
  }
  windows_configuration {
    #Required
    exclude_paths = ["c:\\windows\\private1", "d:\\data\\private1"]
    include_paths = ["c:\\windows", "d:\\data"]
  }
  work_request_validity_period_in_days = 10
}
```

## Argument Reference

The following arguments are supported:

* `agent_polling_interval_in_minutes` - (Optional) (Updatable) Agent polling interval in minutes 
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `is_capturing_ip_address_and_fqdn_enabled` - (Optional) (Updatable) Collect network addresses of managed instances in the fleet. 
* `is_collecting_managed_instance_metrics_enabled` - (Optional) (Updatable) Collect JMS agent metrics on all managed instances in the fleet. 
* `is_collecting_usernames_enabled` - (Optional) (Updatable) Collect username for application invocations for all managed instances in the fleet. 
* `is_libraries_scan_enabled` - (Optional) (Updatable) Enable libraries scan on all managed instances in the fleet. 
* `java_usage_tracker_processing_frequency_in_minutes` - (Optional) (Updatable) The frequency (in minutes) of Java Usage Tracker processing. (That is, how often should JMS process data from the Java Usage Tracker.) 
* `jre_scan_frequency_in_minutes` - (Optional) (Updatable) The frequency (in minutes) of JRE scanning. (That is, how often should JMS scan for JRE installations.) 
* `linux_configuration` - (Optional) (Updatable) Management Agent Configuration for list of include/exclude file system paths (specific to operating system). 
	* `exclude_paths` - (Required) (Updatable) An array of file system paths (environment variables supported). 
	* `include_paths` - (Required) (Updatable) An array of file system paths (environment variables supported). 
* `mac_os_configuration` - (Optional) (Updatable) Management Agent Configuration for list of include/exclude file system paths (specific to operating system). 
	* `exclude_paths` - (Required) (Updatable) An array of file system paths (environment variables supported). 
	* `include_paths` - (Required) (Updatable) An array of file system paths (environment variables supported). 
* `windows_configuration` - (Optional) (Updatable) Management Agent Configuration for list of include/exclude file system paths (specific to operating system). 
	* `exclude_paths` - (Required) (Updatable) An array of file system paths (environment variables supported). 
	* `include_paths` - (Required) (Updatable) An array of file system paths (environment variables supported). 
* `work_request_validity_period_in_days` - (Optional) (Updatable) The validity period in days for work requests. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fleet Agent Configuration
	* `update` - (Defaults to 20 minutes), when updating the Fleet Agent Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Fleet Agent Configuration


## Import

FleetAgentConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_jms_fleet_agent_configuration.test_fleet_agent_configuration "fleets/{fleetId}/agentConfiguration" 
```

