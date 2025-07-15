---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_containers"
sidebar_current: "docs-oci-datasource-jms-fleet_containers"
description: |-
  Provides the list of Fleet Containers in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_containers
This data source provides the list of Fleet Containers in Oracle Cloud Infrastructure Jms service.

List containers in a fleet filtered by query parameters.

## Example Usage

```hcl
data "oci_jms_fleet_containers" "test_fleet_containers" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id

	#Optional
	application_name = oci_dataflow_application.test_application.name
	display_name = var.fleet_container_display_name
	jre_security_status = var.fleet_container_jre_security_status
	jre_version = var.fleet_container_jre_version
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
	time_started_greater_than_or_equal_to = var.fleet_container_time_started_greater_than_or_equal_to
	time_started_less_than_or_equal_to = var.fleet_container_time_started_less_than_or_equal_to
}
```

## Argument Reference

The following arguments are supported:

* `application_name` - (Optional) The name of the application.
* `display_name` - (Optional) The display name.
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `jre_security_status` - (Optional) The security status of the Java Runtime.
* `jre_version` - (Optional) The version of the related Java Runtime.
* `managed_instance_id` - (Optional) The Fleet-unique identifier of the managed instance.
* `time_started_greater_than_or_equal_to` - (Optional) If specified, only containers with a start time later than or equal to this parameter will be included in the response (formatted according to RFC3339).
* `time_started_less_than_or_equal_to` - (Optional) If specified, only containers with a start time earlier than or equal to this parameter will be included in the response (formatted according to RFC3339).


## Attributes Reference

The following attributes are exported:

* `container_collection` - The list of container_collection.

### FleetContainer Reference

The following attributes are exported:

* `items` - A list of the container summaries.
	* `application_key` - Unique key that identifies the application running in the container.
	* `application_name` - The name of the application running in the container.
	* `container_key` - Unique identifier for the container.
	* `display_name` - The name of the container.
	* `image_name` - The container image name.
	* `java_version` - The Java runtime used to run the application in the container.
	* `jre_key` - Unique key that identifies the Java runtime used to run the application in the container.
	* `jre_security_status` - The security status of the Java runtime used to run the application in the container.
	* `managed_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated managed instance of type OCMA.
	* `namespace` - The namespace of the container.
	* `node_name` - The name of the node associated with the pod running this container.
	* `pod_name` - The name of the pod running this container.
	* `time_started` - The start time of the container.

