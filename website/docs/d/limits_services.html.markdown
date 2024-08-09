---
subcategory: "Limits"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_limits_services"
sidebar_current: "docs-oci-datasource-limits-services"
description: |-
  Provides the list of Services in Oracle Cloud Infrastructure Limits service
---

# Data Source: oci_limits_services
This data source provides the list of Services in Oracle Cloud Infrastructure Limits service.

Returns the list of supported services. If subscription ID is provided then only services supported by subscription will be returned.
This includes the programmatic service name, along with the friendly service name.


## Example Usage

```hcl
data "oci_limits_services" "test_services" {
	#Required
	compartment_id = var.tenancy_ocid

	#Optional
    subscription_id = var.subscription_ocid
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the parent compartment (remember that the tenancy is simply the root compartment). 
* `subscription_id` - (Optional) The OCID of the subscription assigned to tenant 


## Attributes Reference

The following attributes are exported:

* `services` - The list of services.

### Service Reference

The following attributes are exported:

* `description` - The friendly service name.
* `name` - The service name. Use this when calling other APIs.
* `supported_subscriptions` - An array of subscription types supported by the service. e,g The type of subscription, such as 'SAAS', 'ERP', 'CRM'. 

