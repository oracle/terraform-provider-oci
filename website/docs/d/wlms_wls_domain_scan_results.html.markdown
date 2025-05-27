---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_wls_domain_scan_results"
sidebar_current: "docs-oci-datasource-wlms-wls_domain_scan_results"
description: |-
  Provides the list of Wls Domain Scan Results in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_wls_domain_scan_results
This data source provides the list of Wls Domain Scan Results in Oracle Cloud Infrastructure Wlms service.

Get all scan results for a server in a specific WebLogic domain.


## Example Usage

```hcl
data "oci_wlms_wls_domain_scan_results" "test_wls_domain_scan_results" {
	#Required
	wls_domain_id = oci_wlms_wls_domain.test_wls_domain.id

	#Optional
	server_name = var.wls_domain_scan_result_server_name
}
```

## Argument Reference

The following arguments are supported:

* `server_name` - (Optional) The name of the server.
* `wls_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.


## Attributes Reference

The following attributes are exported:

* `scan_result_collection` - The list of scan_result_collection.

### WlsDomainScanResult Reference

The following attributes are exported:

* `items` - List of scan results.
	* `server_check_name` - The name of the check performed.
	* `server_check_result` - The result of the server check.
	* `server_check_result_id` - The identifier of the the server check result.
	* `server_check_status` - The status of the server check which is OK, FAILURE, or WARNING.
	* `server_name` - The name of the WebLogic server to which the server check belongs.
	* `time_of_server_check` - The date when the WebLogic server health check is performed (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).  Example: `2016-08-25T21:10:29.600Z` 
	* `wls_domain_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.

