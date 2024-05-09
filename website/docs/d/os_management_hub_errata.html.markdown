---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_errata"
sidebar_current: "docs-oci-datasource-os_management_hub-errata"
description: |-
	Provides the list of Errata in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_errata
This data source provides the list of Errata in Oracle Cloud Infrastructure Os Management Hub service.

Lists all of the currently available errata. Filter the list against a variety of criteria including but not
limited to its name, classification type, advisory severity, and OS family.


## Example Usage

```hcl
data "oci_os_management_hub_errata" "test_errata" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	advisory_severity = var.erratum_advisory_severity
	advisory_type = var.erratum_advisory_type
	classification_type = var.erratum_classification_type
	name = var.erratum_name
	name_contains = var.erratum_name_contains
	os_family = var.erratum_os_family
	time_issue_date_end = var.erratum_time_issue_date_end
	time_issue_date_start = var.erratum_time_issue_date_start
}
```

## Argument Reference

The following arguments are supported:

* `advisory_severity` - (Optional) The advisory severity.
* `advisory_type` - (Optional) A filter to return only errata that match the given advisory types.
* `classification_type` - (Optional) A filter to return only packages that match the given update classification type.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This parameter is required and returns only resources contained within the specified compartment.
* `name` - (Optional) The assigned erratum name. It's unique and not changeable.  Example: `ELSA-2020-5804`
* `name_contains` - (Optional) A filter to return resources that may partially match the erratum name given.
* `os_family` - (Optional) A filter to return only resources that match the given operating system family.
* `time_issue_date_end` - (Optional) The issue date before which to list all errata, in ISO 8601 format  Example: 2017-07-14T02:40:00.000Z
* `time_issue_date_start` - (Optional) The issue date after which to list all errata, in ISO 8601 format  Example: 2017-07-14T02:40:00.000Z


## Attributes Reference

The following attributes are exported:

* `erratum_collection` - The list of erratum_collection.

### Erratum Reference

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

