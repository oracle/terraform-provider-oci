---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_database_group_group_member"
sidebar_current: "docs-oci-datasource-data_safe-target_database_group_group_member"
description: |-
  Provides details about a specific Target Database Group Group Member in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_target_database_group_group_member
This data source provides details about a specific Target Database Group Group Member resource in Oracle Cloud Infrastructure Data Safe service.

Retrieves the members of the target database group with the specified OCID.


## Example Usage

```hcl
data "oci_data_safe_target_database_group_group_member" "test_target_database_group_group_member" {
	#Required
	target_database_group_id = oci_data_safe_target_database_group.test_target_database_group.id

	#Optional
	target_database_id = oci_data_safe_target_database.test_target_database.id
}
```

## Argument Reference

The following arguments are supported:

* `target_database_group_id` - (Required) The OCID of the specified target database group.
* `target_database_id` - (Optional) A filter to return the target database only if it is a member of the target database group.


## Attributes Reference

The following attributes are exported:

* `target_databases` - List of the OCIDs of the target databases which are members of the target database group.

