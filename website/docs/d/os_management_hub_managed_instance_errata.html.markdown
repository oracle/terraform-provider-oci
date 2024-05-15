---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_errata"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_errata"
description: |-
  Provides the list of Managed Instance Errata in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_errata
This data source provides the list of Managed Instance Errata in Oracle Cloud Infrastructure Os Management Hub service.

Returns a list of applicable errata on the managed instance.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_errata" "test_managed_instance_errata" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

	#Optional
	classification_type = var.managed_instance_errata_classification_type
	compartment_id = var.compartment_id
	name = var.managed_instance_errata_name
	name_contains = var.managed_instance_errata_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `classification_type` - (Optional) A filter to return only packages that match the given update classification type.
* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `name` - (Optional) The assigned erratum name. It's unique and not changeable.  Example: `ELSA-2020-5804` 
* `name_contains` - (Optional) A filter to return resources that may partially match the erratum name given.


## Attributes Reference

The following attributes are exported:

* `managed_instance_erratum_summary_collection` - The list of managed_instance_erratum_summary_collection.

### ManagedInstanceErrata Reference

The following attributes are exported:

* `items` - List of errata.
	* `advisory_type` - The advisory type of the erratum.
	* `name` - The identifier of the erratum.
	* `packages` - The list of packages affected by this erratum.
		* `architecture` - The CPU architecture type for which this package was built.
		* `display_name` - Full package name in NERVA format. This value should be unique.
		* `name` - The name of the software package.
		* `type` - Type of the package.
		* `version` - The version of the software package.
	* `related_cves` - The list of CVEs applicable to this erratum.
	* `synopsis` - A summary description of the erratum.
	* `time_issued` - The date and time the package was issued by a providing erratum (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 

