---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_exadata_infrastructure"
sidebar_current: "docs-oci-datasource-database-autonomous_exadata_infrastructure"
description: |-
  Provides details about a specific Autonomous Exadata Infrastructure in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_exadata_infrastructure
This data source provides details about a specific Autonomous Exadata Infrastructure resource in Oracle Cloud Infrastructure Database service.

**Deprecated.** Use the [GetCloudExadataInfrastructure](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudExadataInfrastructure/GetCloudExadataInfrastructure) operation to get details of an Exadata Infrastructure resource and the [GetCloudAutonomousVmCluster](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudAutonomousVmCluster/GetCloudAutonomousVmCluster) operation to get details of an Autonomous Exadata VM cluster.


## Example Usage

```hcl
data "oci_database_autonomous_exadata_infrastructure" "test_autonomous_exadata_infrastructure" {
	#Required
	autonomous_exadata_infrastructure_id = oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_exadata_infrastructure_id` - (Required) The Autonomous Exadata Infrastructure  [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain that the Autonomous Exadata Infrastructure is located in.
* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Autonomous Exadata Infrastructure.
* `domain` - The domain name for the Autonomous Exadata Infrastructure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - The host name for the Autonomous Exadata Infrastructure node.
* `id` - The OCID of the Autonomous Exadata Infrastructure.
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `license_model` - The Oracle license model that applies to all databases in the Autonomous Exadata Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
* `lifecycle_details` - Additional information about the current lifecycle state of the Autonomous Exadata Infrastructure.
* `maintenance_window` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `days_of_week` - Days during the week when maintenance should be performed.
		* `name` - Name of the day of the week.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `lead_time_in_weeks` - Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - Months during the year when maintenance should be performed.
		* `name` - Name of the month of the year.
	* `preference` - The maintenance window scheduling preference.
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty. 
* `scan_dns_name` - The FQDN of the DNS record for the SCAN IP addresses that are associated with the Autonomous Exadata Infrastructure. 
* `shape` - The shape of the Autonomous Exadata Infrastructure. The shape determines resources to allocate to the Autonomous Exadata Infrastructure (CPU cores, memory and storage).
* `state` - The current lifecycle state of the Autonomous Exadata Infrastructure.
* `subnet_id` - The OCID of the subnet the Autonomous Exadata Infrastructure is associated with.

	**Subnet Restrictions:**
	* For Autonomous Databases with Autonomous Exadata Infrastructure, do not use a subnet that overlaps with 192.168.128.0/20

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `time_created` - The date and time the Autonomous Exadata Infrastructure was created.
* `zone_id` - The OCID of the zone the Autonomous Exadata Infrastructure is associated with. 

