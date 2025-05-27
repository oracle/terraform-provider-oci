---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_compliance_record_counts"
sidebar_current: "docs-oci-datasource-fleet_apps_management-compliance_record_counts"
description: |-
  Provides the list of Compliance Record Counts in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_compliance_record_counts
This data source provides the list of Compliance Record Counts in Oracle Cloud Infrastructure Fleet Apps Management service.

Retrieve  aggregated summary information of ComplianceRecords within a Compartment.


## Example Usage

```hcl
data "oci_fleet_apps_management_compliance_record_counts" "test_compliance_record_counts" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.compliance_record_count_compartment_id_in_subtree
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources. Empty only if the resource OCID query param is not specified. 
* `compartment_id_in_subtree` - (Optional) If set to true, resources will be returned for not only the provided compartment, but all compartments which descend from it. Which resources are returned and their field contents depends on the value of accessLevel. 


## Attributes Reference

The following attributes are exported:

* `compliance_record_aggregation_collection` - The list of compliance_record_aggregation_collection.

### ComplianceRecordCount Reference

The following attributes are exported:

* `items` - List of ComplianceRecordAggregation objects.
	* `compliance_record_count_count` - Count of compliance records in a compartment.
	* `dimensions` - Aggregated summary information for ComplianceRecord
		* `compliance_level` - Level at which the compliance is calculated.
		* `compliance_state` - Last known compliance state.

