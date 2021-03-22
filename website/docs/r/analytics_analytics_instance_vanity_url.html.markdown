---
subcategory: "Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_analytics_analytics_instance_vanity_url"
sidebar_current: "docs-oci-resource-analytics-analytics_instance_vanity_url"
description: |-
  Provides the Analytics Instance Vanity Url resource in Oracle Cloud Infrastructure Analytics service
---

# oci_analytics_analytics_instance_vanity_url
This resource provides the Analytics Instance Vanity Url resource in Oracle Cloud Infrastructure Analytics service.

Allows specifying a custom host name to be used to access the analytics instance.  This requires prior setup of DNS entry and certificate
for this host.


## Example Usage

```hcl
resource "oci_analytics_analytics_instance_vanity_url" "test_analytics_instance_vanity_url" {
	#Required
	analytics_instance_id = oci_analytics_analytics_instance.test_analytics_instance.id
	ca_certificate = var.analytics_instance_vanity_url_ca_certificate
	hosts = var.analytics_instance_vanity_url_hosts
	private_key = var.analytics_instance_vanity_url_private_key
	public_certificate = var.analytics_instance_vanity_url_public_certificate

	#Optional
	description = var.analytics_instance_vanity_url_description
	passphrase = var.analytics_instance_vanity_url_passphrase
}
```

## Argument Reference

The following arguments are supported:

* `analytics_instance_id` - (Required) The OCID of the AnalyticsInstance. 
* `ca_certificate` - (Required) (Updatable) PEM CA certificate(s) for HTTPS connections. This may include multiple PEM certificates. 
* `description` - (Optional) Optional description. 
* `hosts` - (Required) List of fully qualified hostnames supported by this vanity URL definition (max of 3). 
* `passphrase` - (Optional) (Updatable) Passphrase for the PEM Private key (if any). 
* `private_key` - (Required) (Updatable) PEM Private key for HTTPS connections. 
* `public_certificate` - (Required) (Updatable) PEM certificate for HTTPS connections. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Analytics Instance Vanity Url
	* `update` - (Defaults to 20 minutes), when updating the Analytics Instance Vanity Url
	* `delete` - (Defaults to 20 minutes), when destroying the Analytics Instance Vanity Url


## Import

AnalyticsInstanceVanityUrls can be imported using the `id`, e.g.

```
$ terraform import oci_analytics_analytics_instance_vanity_url.test_analytics_instance_vanity_url "analyticsInstances/{analyticsInstanceId}/vanityUrls/{vanityUrlKey}" 
```

