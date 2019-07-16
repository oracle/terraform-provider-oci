---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_exadata_infrastructure"
sidebar_current: "docs-oci-resource-database-autonomous_exadata_infrastructure"
description: |-
  Provides the Autonomous Exadata Infrastructure resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_exadata_infrastructure
This resource provides the Autonomous Exadata Infrastructure resource in Oracle Cloud Infrastructure Database service.

Launches a new Autonomous Exadata Infrastructure in the specified compartment and availability domain.


## Example Usage

```hcl
resource "oci_database_autonomous_exadata_infrastructure" "test_autonomous_exadata_infrastructure" {
	#Required
	availability_domain = "${var.autonomous_exadata_infrastructure_availability_domain}"
	compartment_id = "${var.compartment_id}"
	shape = "${var.autonomous_exadata_infrastructure_shape}"
	subnet_id = "${oci_database_subnet.test_subnet.id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.autonomous_exadata_infrastructure_display_name}"
	domain = "${var.autonomous_exadata_infrastructure_domain}"
	freeform_tags = {"Department"= "Finance"}
	license_model = "${var.autonomous_exadata_infrastructure_license_model}"
	maintenance_window_details {
		#Required
		preference = "${var.autonomous_exadata_infrastructure_maintenance_window_details_preference}"

		#Optional
		days_of_week {
			#Required
			name = "${var.autonomous_exadata_infrastructure_maintenance_window_details_days_of_week_name}"
		}
		hours_of_day = "${var.autonomous_exadata_infrastructure_maintenance_window_details_hours_of_day}"
		months {
			#Required
			name = "${var.autonomous_exadata_infrastructure_maintenance_window_details_months_name}"
		}
		weeks_of_month = "${var.autonomous_exadata_infrastructure_maintenance_window_details_weeks_of_month}"
	}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain where the Autonomous Exadata Infrastructure is located.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the Autonomous Exadata Infrastructure belongs in.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) The user-friendly name for the Autonomous Exadata Infrastructure. It does not have to be unique.
* `domain` - (Optional) A domain name used for the Autonomous Exadata Infrastructure. If the Oracle-provided Internet and VCN Resolver is enabled for the specified subnet, the domain name for the subnet is used (don't provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `license_model` - (Optional) The Oracle license model that applies to all the databases in the Autonomous Exadata Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
* `maintenance_window_details` - (Optional) (Updatable) 
	* `days_of_week` - (Optional) (Updatable) Days during the week when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the day of the week.
	* `hours_of_day` - (Optional) (Updatable) The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `months` - (Optional) (Updatable) Months during the year when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the month of the year.
	* `preference` - (Required) (Updatable) The maintenance window scheduling preference.
	* `weeks_of_month` - (Optional) (Updatable) Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `shape` - (Required) The shape of the Autonomous Exadata Infrastructure. The shape determines resources allocated to the Autonomous Exadata Infrastructure (CPU cores, memory and storage). To get a list of shapes, use the ListDbSystemShapes operation.
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the Autonomous Exadata Infrastructure is associated with.

	**Subnet Restrictions:**
	* For Autonomous Exadata Infrastructures, do not use a subnet that overlaps with 192.168.128.0/20

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain that the Autonomous Exadata Infrastructure is located in.
* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the Autonomous Exadata Infrastructure.
* `domain` - The domain name for the Autonomous Exadata Infrastructure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - The host name for the Autonomous Exadata Infrastructure node.
* `id` - The OCID of the Autonomous Exadata Infrastructure.
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `license_model` - The Oracle license model that applies to all databases in the Autonomous Exadata Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
* `lifecycle_details` - Additional information about the current lifecycle state of the Autonomous Exadata Infrastructure.
* `maintenance_window` - 
	* `days_of_week` - Days during the week when maintenance should be performed.
		* `name` - Name of the day of the week.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `months` - Months during the year when maintenance should be performed.
		* `name` - Name of the month of the year.
	* `preference` - The maintenance window scheduling preference.
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `shape` - The shape of the Autonomous Exadata Infrastructure. The shape determines resources to allocate to the Autonomous Exadata Infrastructure (CPU cores, memory and storage).
* `state` - The current lifecycle state of the Autonomous Exadata Infrastructure.
* `subnet_id` - The OCID of the subnet the Autonomous Exadata Infrastructure is associated with.

	**Subnet Restrictions:**
	* For Autonomous Databases with Autonomous Exadata Infrastructure, do not use a subnet that overlaps with 192.168.128.0/20

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `time_created` - The date and time the Autonomous Exadata Infrastructure was created.

## Import

AutonomousExadataInfrastructures can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure "id"
```

