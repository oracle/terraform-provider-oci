---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_list_jre_usage"
sidebar_current: "docs-oci-datasource-jms-list_jre_usage"
description: |-
  Provides details about a specific List Jre Usage in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_list_jre_usage
This data source provides details about a specific List Jre Usage resource in Oracle Cloud Infrastructure Jms service.

List Java Runtime usage in a specified host filtered by query parameters.

## Example Usage

```hcl
data "oci_jms_list_jre_usage" "test_list_jre_usage" {

	#Optional
	application_id = oci_dataflow_application.test_application.id
	application_name = oci_dataflow_application.test_application.name
	compartment_id = var.compartment_id
	host_id = oci_jms_host.test_host.id
	time_end = var.list_jre_usage_time_end
	time_start = var.list_jre_usage_time_start
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The Fleet-unique identifier of the application.
* `application_name` - (Optional) The name of the application.
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. 
* `host_id` - (Optional) The host [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `time_end` - (Optional) The end of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_start` - (Optional) The start of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).


## Attributes Reference

The following attributes are exported:

* `items` - A list of Java Runtimes.
	* `approximate_application_count` - The approximate count of the applications running on this Java Runtime.
	* `approximate_installation_count` - The approximate count of installations that are installations of this Java Runtime.
	* `approximate_managed_instance_count` - The approximate count of the managed instances that report this Java Runtime.
	* `distribution` - The distribution of a Java Runtime is the name of the lineage of product to which it belongs, for example _Java(TM) SE Runtime Environment_.
	* `end_of_support_life_date` - The End of Support Life (EOSL) date of the Java Runtime (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
	* `fleet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related fleet.  This property value is present only for /actions/listJreUsage.
	* `id` - The internal identifier of the Java Runtime.
	* `managed_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance. This property value is present only for /actions/listJreUsage.
	* `operating_systems` - The operating systems that have this Java Runtime installed.
		* `architecture` - The architecture of the operating system as provided by the Java system property os.arch.
		* `family` - The operating system type, such as Windows or Linux
		* `name` - The name of the operating system as provided by the Java system property os.name.
		* `version` - The version of the operating system as provided by the Java system property os.version.
	* `release_date` - The release date of the Java Runtime (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
	* `security_status` - The security status of the Java Runtime.
	* `time_end` - Upper bound of the specified time period filter.
	* `time_first_seen` - The date and time the resource was _first_ reported to JMS. This is potentially _before_ the specified time period provided by the filters. For example, a resource can be first reported to JMS before the start of a specified time period, if it is also reported during the time period. 
	* `time_last_seen` - The date and time the resource was _last_ reported to JMS. This is potentially _after_ the specified time period provided by the filters. For example, a resource can be last reported to JMS before the start of a specified time period, if it is also reported during the time period. 
	* `time_start` - Lower bound of the specified time period filter.
	* `vendor` - The vendor of the Java Runtime.
	* `version` - The version of the Java Runtime.

