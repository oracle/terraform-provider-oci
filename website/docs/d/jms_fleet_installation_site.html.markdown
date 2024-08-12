---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_installation_site"
sidebar_current: "docs-oci-datasource-jms-fleet_installation_site"
description: |-
  Provides details about a specific Fleet Installation Site in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_installation_site
This data source provides details about a specific Fleet Installation Site resource in Oracle Cloud Infrastructure Jms service.

List Java installation sites in a Fleet filtered by query parameters.

## Example Usage

```hcl
data "oci_jms_fleet_installation_site" "test_fleet_installation_site" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id

	#Optional
	application_id = var.fleet_installation_site_application_id
	installation_path = var.fleet_installation_site_installation_path
	jre_distribution = var.fleet_installation_site_jre_distribution
	jre_security_status = var.fleet_installation_site_jre_security_status
	jre_vendor = var.fleet_installation_site_jre_vendor
	jre_version = var.fleet_installation_site_jre_version
	managed_instance_id = var.fleet_installation_site_managed_instance_id
	os_family = var.fleet_installation_site_os_family
	path_contains = var.fleet_installation_site_path_contains
	time_end = var.fleet_installation_site_time_end
	time_start = var.fleet_installation_site_time_start
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The Fleet-unique identifier of the related application.
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `installation_path` - (Optional) The file system path of the installation.
* `jre_distribution` - (Optional) The distribution of the related Java Runtime.
* `jre_security_status` - (Optional) The security status of the Java Runtime.
* `jre_vendor` - (Optional) The vendor of the related Java Runtime.
* `jre_version` - (Optional) The version of the related Java Runtime.
* `managed_instance_id` - (Optional) The Fleet-unique identifier of the related managed instance.
* `os_family` - (Optional) The operating system type.
* `path_contains` - (Optional) Filter the list with path contains the given value. 
* `time_end` - (Optional) The end of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_start` - (Optional) The start of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).


## Attributes Reference

The following attributes are exported:

* `items` - A list of Java installation sites.
	* `approximate_application_count` - The approximate count of applications running on this installation
	* `blocklist` - The list of operations that are blocklisted.
		* `operation` - The operation type.
		* `reason` - The reason why the operation is blocklisted.
	* `installation_key` - The unique identifier for the installation of Java Runtime at a specific path on a specific operating system.
	* `jre` - The essential properties to identify a Java Runtime.
		* `distribution` - The distribution of a Java Runtime is the name of the lineage of product to which it belongs, for example _Java(TM) SE Runtime Environment_.
		* `jre_key` - The unique identifier for a Java Runtime.
		* `vendor` - The vendor of the Java Runtime.
		* `version` - The version of the Java Runtime.
	* `managed_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance. 
	* `operating_system` - Operating System of the platform on which the Java Runtime was reported. 
		* `architecture` - The architecture of the operating system as provided by the Java system property os.arch.
		* `family` - The operating system type, such as Windows or Linux
		* `managed_instance_count` - Number of instances running the operating system
		* `name` - The name of the operating system as provided by the Java system property os.name.
		* `version` - The version of the operating system as provided by the Java system property os.version.
	* `path` - The file system path of the installation.
	* `security_status` - The security status of the Java Runtime.
	* `state` - The lifecycle state of the installation site.
	* `time_last_seen` - The date and time the resource was _last_ reported to JMS. This is potentially _after_ the specified time period provided by the filters. For example, a resource can be last reported to JMS before the start of a specified time period, if it is also reported during the time period. 

