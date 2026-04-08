---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_errata"
sidebar_current: "docs-oci-datasource-os_management_hub-errata"
description: |-
  Provides details about a specific Errata in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_errata
This data source provides details about a specific Errata resource in Oracle Cloud Infrastructure Os Management Hub service.

Returns information about the specified erratum based on its advisory name.


## Example Usage

```hcl
data "oci_os_management_hub_errata" "test_errata" {
	#Required
	compartment_id = var.compartment_id
	name = var.errata_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This parameter is required and returns only resources contained within the specified compartment.
* `name` - (Required) The erratum name (such as ELSA-2023-34678).


## Attributes Reference

The following attributes are exported:

* `advisory_severity` - The severity for a security advisory, otherwise, null.
* `advisory_type` - The advisory type of the erratum.
* `classification_type` - Type of the erratum. This property is deprecated and it will be removed in a future API release. Please refer to the advisoryType property instead.
* `description` - Details describing the erratum.
* `from` - Information specifying from where the erratum was release.
* `name` - Advisory name.
* `os_families` - List of affected OS families.
* `packages` - List of packages affected by this erratum.
	* `architecture` - The architecture for which this software was built.
	* `checksum` - Checksum of the package.
	* `checksum_type` - Type of the checksum.
	* `display_name` - Package name.
	* `is_latest` - Indicates whether this package is the latest version.
	* `last_modified_date` - The date and time the package was last modified (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
	* `name` - Unique identifier for the package. Note that this is not an OCID.
	* `os_families` - The OS families the package belongs to.
	* `software_sources` - List of software sources that provide the software package. This property is deprecated and it will be removed in a future API release.
		* `description` - Software source description.
		* `display_name` - Software source name.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
		* `is_mandatory_for_autonomous_linux` - Indicates whether this is a required software source for Autonomous Linux instances. If true, the user can't unselect it.
		* `software_source_type` - Type of the software source.
	* `type` - Type of the package.
	* `version` - Version of the package.
* `references` - Information describing how to find more information about. the erratum.
* `related_cves` - List of CVEs applicable to this erratum.
* `repositories` - List of repository identifiers.
* `solution` - Information describing how the erratum can be resolved.
* `synopsis` - Summary description of the erratum.
* `time_issued` - The date and time the erratum was issued (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
* `time_updated` - The date and time the erratum was updated (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 

