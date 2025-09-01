---
subcategory: "Jms Utils"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_utils_analyze_applications_configuration"
sidebar_current: "docs-oci-resource-jms_utils-analyze_applications_configuration"
description: |-
  Provides the Analyze Applications Configuration resource in Oracle Cloud Infrastructure Jms Utils service
---

# oci_jms_utils_analyze_applications_configuration
This resource provides the Analyze Applications Configuration resource in Oracle Cloud Infrastructure Jms Utils service.

Updates the configuration for analyze application.

## Example Usage

```hcl
resource "oci_jms_utils_analyze_applications_configuration" "test_analyze_applications_configuration" {

	#Optional
	bucket = var.analyze_applications_configuration_bucket
	compartment_id = var.compartment_id
	namespace = var.analyze_applications_configuration_namespace
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Optional) (Updatable) The name of the bucket used for analyzing applications.
* `compartment_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `namespace` - (Optional) (Updatable) The Object Storage namespace used for analyzing applications.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bucket` - The name of the bucket used for analyzing applications.
* `namespace` - The Object Storage namespace used for analyzing applications.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Analyze Applications Configuration
	* `update` - (Defaults to 20 minutes), when updating the Analyze Applications Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Analyze Applications Configuration


## Import

AnalyzeApplicationsConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_jms_utils_analyze_applications_configuration.test_analyze_applications_configuration "id"
```

