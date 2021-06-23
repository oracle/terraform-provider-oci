---
subcategory: "Health Checks"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_health_checks_ping_probe"
sidebar_current: "docs-oci-resource-health_checks-ping_probe"
description: |-
  Provides the Ping Probe resource in Oracle Cloud Infrastructure Health Checks service
---

# oci_health_checks_ping_probe
This resource provides the Ping Probe resource in Oracle Cloud Infrastructure Health Checks service.

Creates an on-demand ping probe. The location response header contains the URL for
fetching probe results.

*Note:* The on-demand probe configuration is not saved.


## Example Usage

```hcl
resource "oci_health_checks_ping_probe" "test_ping_probe" {
	#Required
	compartment_id = var.compartment_id
	protocol = var.ping_probe_protocol
	targets = var.ping_probe_targets

	#Optional
	port = var.ping_probe_port
	timeout_in_seconds = var.ping_probe_timeout_in_seconds
	vantage_point_names = var.ping_probe_vantage_point_names
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `port` - (Optional) The port on which to probe endpoints. If unspecified, probes will use the default port of their protocol. 
* `protocol` - (Required) The protocols for ping probes.
* `targets` - (Required) A list of targets (hostnames or IP addresses) of the probe.
* `timeout_in_seconds` - (Optional) The probe timeout in seconds. Valid values: 10, 20, 30, and 60. The probe timeout must be less than or equal to `intervalInSeconds` for monitors. 
* `vantage_point_names` - (Optional) A list of names of vantage points from which to execute the probe.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `home_region` - The region where updates must be made and where results must be fetched from. 
* `id` - The OCID of the resource.
* `port` - The port on which to probe endpoints. If unspecified, probes will use the default port of their protocol. 
* `protocol` - The protocols for ping probes.
* `results_url` - A URL for fetching the probe results.
* `targets` - A list of targets (hostnames or IP addresses) of the probe.
* `time_created` - The RFC 3339-formatted creation date and time of the probe. 
* `timeout_in_seconds` - The probe timeout in seconds. Valid values: 10, 20, 30, and 60. The probe timeout must be less than or equal to `intervalInSeconds` for monitors. 
* `vantage_point_names` - A list of names of vantage points from which to execute the probe.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ping Probe
	* `update` - (Defaults to 20 minutes), when updating the Ping Probe
	* `delete` - (Defaults to 20 minutes), when destroying the Ping Probe


## Import

PingProbes can be imported using the `id`, e.g.

```
$ terraform import oci_health_checks_ping_probe.test_ping_probe "id"
```

