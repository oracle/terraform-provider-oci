---
subcategory: "Health Checks"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_health_checks_vantage_points"
sidebar_current: "docs-oci-datasource-health_checks-vantage_points"
description: |-
  Provides the list of Vantage Points in Oracle Cloud Infrastructure Health Checks service
---

# Data Source: oci_health_checks_vantage_points
This data source provides the list of Vantage Points in Oracle Cloud Infrastructure Health Checks service.

Gets information about all vantage points available to the user.


## Example Usage

```hcl
data "oci_health_checks_vantage_points" "test_vantage_points" {

	#Optional
	display_name = var.vantage_point_display_name
	name = var.vantage_point_name
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) Filters results that exactly match the `displayName` field.
* `name` - (Optional) Filters results that exactly match the `name` field.


## Attributes Reference

The following attributes are exported:

* `health_checks_vantage_points` - The list of health_checks_vantage_points.

### VantagePoint Reference

The following attributes are exported:

* `display_name` - The display name for the vantage point. Display names are determined by the best information available and may change over time. 
* `geo` - Geographic information about a vantage point.
	* `admin_div_code` - The ISO 3166-2 code for this location's first-level administrative division, either a US state or Canadian province. Only included for locations in the US or Canada. For a list of codes, see [Country Codes](https://www.iso.org/obp/ui/#search). 
	* `city_name` - Common English-language name for the city. 
	* `country_code` - The ISO 3166-1 alpha-2 country code. For a list of codes, see [Country Codes](https://www.iso.org/obp/ui/#search). 
	* `country_name` - The common English-language name for the country. 
	* `geo_key` - An opaque identifier for the geographic location of the vantage point.
	* `latitude` - Degrees north of the Equator. 
	* `longitude` - Degrees east of the prime meridian. 
* `name` - The unique, permanent name for the vantage point.
* `provider_name` - The organization on whose infrastructure this vantage point resides. Provider names are not unique, as Oracle Cloud Infrastructure maintains many vantage points in each major provider. 
* `routing` - An array of objects that describe how traffic to this vantage point is routed, including which prefixes and ASNs connect it to the internet.

	The addresses are sorted from the most-specific to least-specific prefix (the smallest network to largest network). When a prefix has multiple origin ASNs (MOAS routing), they are sorted by weight (highest to lowest). Weight is determined by the total percentage of peers observing the prefix originating from an ASN. Only present if `fields` includes `routing`. The field will be null if the address's routing information is unknown. 
	* `as_label` - The registry label for `asn`, usually the name of the organization that owns the ASN. May be omitted or null. 
	* `asn` - The Autonomous System Number (ASN) identifying the organization responsible for routing packets to `prefix`. 
	* `prefix` - An IP prefix (CIDR syntax) that is less specific than `address`, through which `address` is routed. 
	* `weight` - An integer between 0 and 100 used to select between multiple origin ASNs when routing to `prefix`. Most prefixes have exactly one origin ASN, in which case `weight` will be 100. 

