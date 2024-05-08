```


//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present
---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_resource_vulnerabilities"
sidebar_current: "docs-oci-datasource-cloud_guard-resource_vulnerabilities"
description: |-
  Provides the list of Resource Vulnerabilities in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_resource_vulnerabilities
This data source provides the list of Resource Vulnerabilities in Oracle Cloud Infrastructure Cloud Guard service.

Returns the list of vulnerabilities associated with the resourceId where resource is an instance

## Example Usage

```hcl
data "oci_cloud_guard_resource_vulnerabilities" "test_resource_vulnerabilities" {
	#Required
	resource_id = oci_cloud_guard_resource.test_resource.id

	#Optional
	cve_id = oci_cloud_guard_cve.test_cve.id
	risk_level = var.resource_vulnerability_risk_level
}

## Argument Reference

The following arguments are supported:

* `cve_id` - (Optional) CVE ID associated with the resource.
* `resource_id` - (Required) CloudGuard resource OCID
* `risk_level` - (Optional) Risk level of the problem.


## Attributes Reference

The following attributes are exported:

* `resource_vulnerability_collection` - The list of resource_vulnerability_collection.

### ResourceVulnerability Reference

The following attributes are exported:

* `cvss_score` - cvssScore of CVE
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `description` - The description of the vulnerability
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

  Avoid entering confidential information.
* `id` - The cve id of the vulnerability
* `package_details` - list for packages causing vulnerability
    * `cause` - cause of the vulnerability in the package
    * `location` - location of the package
    * `name` - name of the package
    * `package_type` - type of the package
    * `remediation` - remediation for vulnerability
    * `version` - version of the package
* `risk_level` - The Risk Level
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_first_detected` - First detected time
* `time_last_detected` - First detected time
* `time_last_modified` - Time the vulnerability was last modified
* `time_published` - Time the vulnerability was published
* `url` - URL of the CVE
```