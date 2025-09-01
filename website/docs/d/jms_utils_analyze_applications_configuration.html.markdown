---
subcategory: "Jms Utils"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_utils_analyze_applications_configuration"
sidebar_current: "docs-oci-datasource-jms_utils-analyze_applications_configuration"
description: |-
  Provides details about a specific Analyze Applications Configuration in Oracle Cloud Infrastructure Jms Utils service
---

# Data Source: oci_jms_utils_analyze_applications_configuration
This data source provides details about a specific Analyze Applications Configuration resource in Oracle Cloud Infrastructure Jms Utils service.

Returns the configuration for analyzing applications.


## Example Usage

```hcl
data "oci_jms_utils_analyze_applications_configuration" "test_analyze_applications_configuration" {

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `bucket` - The name of the bucket used for analyzing applications.
* `namespace` - The Object Storage namespace used for analyzing applications.

