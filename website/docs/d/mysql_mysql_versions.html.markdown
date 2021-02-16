---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_mysql_versions"
sidebar_current: "docs-oci-datasource-mysql-mysql_versions"
description: |-
  Provides the list of Mysql Versions in Oracle Cloud Infrastructure MySQL Database service
---

# Data Source: oci_mysql_mysql_versions
This data source provides the list of Mysql Versions in Oracle Cloud Infrastructure MySQL Database service.

Get a list of supported and available MySQL database major versions.

The list is sorted by version family.


## Example Usage

```hcl
data "oci_mysql_mysql_versions" "test_mysql_versions" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `versions` - The list of versions.

### MysqlVersion Reference

The following attributes are exported:

* `version_family` - A descriptive summary of a group of versions.
* `versions` - The list of supported MySQL Versions.
	* `description` - A link to a page describing the version.
	* `version` - The specific version identifier

