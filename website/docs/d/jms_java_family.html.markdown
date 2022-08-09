---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_family"
sidebar_current: "docs-oci-datasource-jms-java_family"
description: |-
  Provides details about a specific Java Family in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_java_family
This data source provides details about a specific Java Family resource in Oracle Cloud Infrastructure Jms service.

Returns details of a Java release family based on specified version.


## Example Usage

```hcl
data "oci_jms_java_family" "test_java_family" {
	#Required
	family_version = var.java_family_family_version
}
```

## Argument Reference

The following arguments are supported:

* `family_version` - (Required) Unique Java family version identifier.


## Attributes Reference

The following attributes are exported:

* `display_name` - The display name of the release family.
* `doc_url` - Link to access the documentation for the release.
* `end_of_support_life_date` - The End of Support Life (EOSL) date of the Java release family (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 
* `family_version` - The Java release family identifier.
* `support_type` - This indicates the support category for the Java release family.

