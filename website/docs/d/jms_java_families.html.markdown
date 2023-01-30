---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_families"
sidebar_current: "docs-oci-datasource-jms-java_families"
description: |-
  Provides the list of Java Families in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_java_families
This data source provides the list of Java Families in Oracle Cloud Infrastructure Jms service.

Returns a list of the Java release family information.
A Java release family is typically a major version in the Java version identifier.


## Example Usage

```hcl
data "oci_jms_java_families" "test_java_families" {

	#Optional
	display_name = var.java_family_display_name
	family_version = var.java_family_family_version
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) The display name for the Java family.
* `family_version` - (Optional) The version identifier for the Java family.


## Attributes Reference

The following attributes are exported:

* `java_family_collection` - The list of java_family_collection.

### JavaFamily Reference

The following attributes are exported:

* `display_name` - The display name of the release family.
* `doc_url` - Link to access the documentation for the release.
* `end_of_support_life_date` - The End of Support Life (EOSL) date of the Java release family (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 
* `family_version` - The Java release family identifier.
* `support_type` - This indicates the support category for the Java release family.

