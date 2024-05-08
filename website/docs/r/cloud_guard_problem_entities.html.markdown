---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_problem_entities"
sidebar_current: "docs-oci-datasource-cloud_guard-problem_entities"
description: |-
  Provides the list of Problem Entities in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_problem_entities
This data source provides the list of Problem Entities in Oracle Cloud Infrastructure Cloud Guard service.

Returns a list of entities for a CloudGuard Problem
Returns a list of entities for a problem.


## Example Usage

```hcl
data "oci_cloud_guard_problem_entities" "test_problem_entities" {
	#Required
	problem_id = oci_cloud_guard_problem.test_problem.id
}
```

## Argument Reference

The following arguments are supported:

* `problem_id` - (Required) OCID of the problem.


## Attributes Reference

The following attributes are exported:

* `problem_entity_collection` - The list of problem_entity_collection.

### ProblemEntity Reference

The following attributes are exported:

* `items` - List of entity details related to a data source
    * `entity_details` - List of entity details related to a data source
        * `display_name` - The display name of entity
        * `type` - Type of entity
        * `value` - The entity value
    * `problem_id` - Attached problem ID
    * `regions` - Data source problem entities region
    * `result_url` - Log result query URL for a data source query
    * `time_first_detected` - Data source problem entities first detected time
    * `time_last_detected` - Data source problem entities last detected time