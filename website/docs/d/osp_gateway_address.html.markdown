---
subcategory: "Osp Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osp_gateway_address"
sidebar_current: "docs-oci-datasource-osp_gateway-address"
description: |-
  Provides details about a specific Address in Oracle Cloud Infrastructure Osp Gateway service
---

# Data Source: oci_osp_gateway_address
This data source provides details about a specific Address resource in Oracle Cloud Infrastructure Osp Gateway service.

Get the address by id for the compartment

## Example Usage

```hcl
data "oci_osp_gateway_address" "test_address" {
	#Required
	address_id = oci_osp_gateway_addres.test_addres.id
	compartment_id = var.compartment_id
	osp_home_region = var.address_osp_home_region
}
```

## Argument Reference

The following arguments are supported:

* `address_id` - (Required) The identifier of the address.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `osp_home_region` - (Required) The home region's public name of the logged in user. 


## Attributes Reference

The following attributes are exported:

* `address_key` - Address identifier.
* `city` - Name of the city.
* `company_name` - Name of the customer company.
* `contributor_class` - Contributor class of the customer company.
* `country` - Country of the address.
* `county` - County of the address.
* `department_name` - Department name of the customer company.
* `email_address` - Contact person email address.
* `first_name` - First name of the contact person.
* `internal_number` - Internal number of the customer company.
* `job_title` - Job title of the contact person.
* `last_name` - Last name of the contact person.
* `line1` - Address line 1.
* `line2` - Address line 2.
* `line3` - Address line 3.
* `line4` - Address line 4.
* `middle_name` - Middle name of the contact person.
* `municipal_inscription` - Municipal Inscription.
* `phone_country_code` - Phone country code of the contact person.
* `phone_number` - Phone number of the contact person.
* `postal_code` - Post code of the address.
* `province` - Province of the address.
* `state` - State of the address.
* `state_inscription` - State Inscription.
* `street_name` - Street name of the address.
* `street_number` - Street number of the address.

