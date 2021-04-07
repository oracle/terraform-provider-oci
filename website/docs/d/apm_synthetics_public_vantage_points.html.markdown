---
subcategory: "Apm Synthetics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_synthetics_public_vantage_points"
sidebar_current: "docs-oci-datasource-apm_synthetics-public_vantage_points"
description: |-
  Provides the list of Public Vantage Points in Oracle Cloud Infrastructure Apm Synthetics service
---

# Data Source: oci_apm_synthetics_public_vantage_points
This data source provides the list of Public Vantage Points in Oracle Cloud Infrastructure Apm Synthetics service.

Returns a list of public vantage points.


## Example Usage

```hcl
data "oci_apm_synthetics_public_vantage_points" "test_public_vantage_points" {
	#Required
	apm_domain_id = oci_apm_synthetics_apm_domain.test_apm_domain.id

	#Optional
	display_name = var.public_vantage_point_display_name
	name = var.public_vantage_point_name
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM domain ID the request is intended for. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `name` - (Optional) A filter to return only resources that match the entire name given.


## Attributes Reference

The following attributes are exported:

* `public_vantage_point_collection` - The list of public_vantage_point_collection.

### PublicVantagePoint Reference

The following attributes are exported:

* `items` - List of PublicVantagePointSummary items.
	* `display_name` - Unique name that can be edited. The name should not contain any confidential information.
	* `geo` - Geographic summary about a vantage point.
		* `admin_div_code` - The ISO 3166-2 code for this location's first-level administrative division, either a US state or Canadian province. Only included for locations in the US or Canada. For a list of codes, see Country Codes. 
		* `city_name` - Common English-language name for the city.
		* `country_code` - The ISO 3166-1 alpha-2 country code. For a list of codes, see Country Codes.
		* `country_name` - The common English-language name for the country.
		* `latitude` - Degrees north of the Equator.
		* `longitude` - Degrees east of the prime meridian.
	* `name` - Unique permanent name of the vantage point.

