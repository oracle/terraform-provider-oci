---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy_report_role_grant_paths"
sidebar_current: "docs-oci-datasource-data_safe-security_policy_report_role_grant_paths"
description: |-
  Provides the list of Security Policy Report Role Grant Paths in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_policy_report_role_grant_paths
This data source provides the list of Security Policy Report Role Grant Paths in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of all role grant paths for a particular user.

The ListRoleGrantPaths operation returns only the role grant paths for the specified security policy report.


## Example Usage

```hcl
data "oci_data_safe_security_policy_report_role_grant_paths" "test_security_policy_report_role_grant_paths" {
	#Required
	granted_role = var.security_policy_report_role_grant_path_granted_role
	grantee = var.security_policy_report_role_grant_path_grantee
	security_policy_report_id = oci_data_safe_security_policy_report.test_security_policy_report.id
}
```

## Argument Reference

The following arguments are supported:

* `granted_role` - (Required) A filter to return only items that match the specified role.
* `grantee` - (Required) A filter to return only items that match the specified grantee.
* `security_policy_report_id` - (Required) The OCID of the security policy report resource.


## Attributes Reference

The following attributes are exported:

* `role_grant_path_collection` - The list of role_grant_path_collection.

### SecurityPolicyReportRoleGrantPath Reference

The following attributes are exported:

* `items` - An array of grant path summary objects.
	* `depth_level` - The grant depth level of the indirect grant. An indirectly granted role/privilege is granted to the user through another role. The depth level indicates how deep a privilege is within the grant hierarchy. 
	* `granted_role` - The name of the role.
	* `grantee` - Grantee is the user who can access the table.
	* `key` - The unique key of a role grant.

