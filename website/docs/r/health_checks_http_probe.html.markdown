---
subcategory: "Health Checks"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_health_checks_http_probe"
sidebar_current: "docs-oci-resource-health_checks-http_probe"
description: |-
  Provides the Http Probe resource in Oracle Cloud Infrastructure Health Checks service
---

# oci_health_checks_http_probe
This resource provides the Http Probe resource in Oracle Cloud Infrastructure Health Checks service.

Creates an on-demand HTTP probe. The location response header contains the URL for
fetching the probe results.

*Note:* On-demand probe configurations are not saved.


## Example Usage

```hcl
resource "oci_health_checks_http_probe" "test_http_probe" {
	#Required
	compartment_id = var.compartment_id
	protocol = var.http_probe_protocol
	targets = var.http_probe_targets

	#Optional
	headers = var.http_probe_headers
	method = var.http_probe_method
	path = var.http_probe_path
	port = var.http_probe_port
	timeout_in_seconds = var.http_probe_timeout_in_seconds
	vantage_point_names = var.http_probe_vantage_point_names
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `headers` - (Optional) A dictionary of HTTP request headers.

	*Note:* Monitors and probes do not support the use of the `Authorization` HTTP header. 
* `method` - (Optional) The supported HTTP methods available for probes.
* `path` - (Optional) The optional URL path to probe, including query parameters.
* `port` - (Optional) The port on which to probe endpoints. If unspecified, probes will use the default port of their protocol. 
* `protocol` - (Required) The supported protocols available for HTTP probes.
* `targets` - (Required) A list of targets (hostnames or IP addresses) of the probe.
* `timeout_in_seconds` - (Optional) The probe timeout in seconds. Valid values: 10, 20, 30, and 60. The probe timeout must be less than or equal to `intervalInSeconds` for monitors. 
* `vantage_point_names` - (Optional) A list of names of vantage points from which to execute the probe.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `headers` - A dictionary of HTTP request headers.

	*Note:* Monitors and probes do not support the use of the `Authorization` HTTP header. 
* `home_region` - The region where updates must be made and where results must be fetched from. 
* `id` - The OCID of the resource.
* `method` - The supported HTTP methods available for probes.
* `path` - The optional URL path to probe, including query parameters.
* `port` - The port on which to probe endpoints. If unspecified, probes will use the default port of their protocol. 
* `protocol` - The supported protocols available for HTTP probes.
* `results_url` - A URL for fetching the probe results.
* `targets` - A list of targets (hostnames or IP addresses) of the probe.
* `time_created` - The RFC 3339-formatted creation date and time of the probe. 
* `timeout_in_seconds` - The probe timeout in seconds. Valid values: 10, 20, 30, and 60. The probe timeout must be less than or equal to `intervalInSeconds` for monitors. 
* `vantage_point_names` - A list of names of vantage points from which to execute the probe.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Http Probe
	* `update` - (Defaults to 20 minutes), when updating the Http Probe
	* `delete` - (Defaults to 20 minutes), when destroying the Http Probe


## Import

HttpProbes can be imported using the `id`, e.g.

```
$ terraform import oci_health_checks_http_probe.test_http_probe "id"
```

