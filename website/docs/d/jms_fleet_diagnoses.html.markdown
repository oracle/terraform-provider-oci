---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_diagnoses"
sidebar_current: "docs-oci-datasource-jms-fleet_diagnoses"
description: |-
  Provides the list of Fleet Diagnoses in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_diagnoses
This data source provides the list of Fleet Diagnoses in Oracle Cloud Infrastructure Jms service.

List potential diagnoses that would put a fleet into FAILED or NEEDS_ATTENTION lifecycle state.


## Example Usage

```hcl
data "oci_jms_fleet_diagnoses" "test_fleet_diagnoses" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.


## Attributes Reference

The following attributes are exported:

* `fleet_diagnosis_collection` - The list of fleet_diagnosis_collection.

### FleetDiagnose Reference

The following attributes are exported:

* `items` - A list of the fleet resource diagnosis.
	* `resource_diagnosis` - The diagnosis message.
	* `resource_id` - The OCID of the external resouce needed by the fleet.
	* `resource_state` - The state of the resource. The resource state is ACTIVE when it works properly for the fleet. In case it would cause an issue for the fleet function, the state is INACTIVE. When JMS can't locate the resource, the state is NOT_FOUND. OTHER covers other cases, such as a temporarily network issue that prevents JMS from detecting the resource. Check the resourceDiagnosis for details. 
	* `resource_type` - The type of the resource needed by the fleet. This is the role of a resource in the fleet. Use the OCID to determine the actual Oracle Cloud Infrastructure resource type such as log group or log. 

