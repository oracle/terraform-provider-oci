---
subcategory: "Osp Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osp_gateway_address_action_verification"
sidebar_current: "docs-oci-resource-osp_gateway-address_action_verification"
description: |-
  Provides the Address Action Verification resource in Oracle Cloud Infrastructure Osp Gateway service
---

# oci_osp_gateway_address_action_verification
This resource provides the Address Action Verification resource in Oracle Cloud Infrastructure Osp Gateway service.

Verify address

## Example Usage

```hcl
resource "oci_osp_gateway_address_action_verification" "test_address_action_verification" {
	#Required
	compartment_id = var.compartment_id
	osp_home_region = var.address_action_verification_osp_home_region

	#Optional
	address_key = var.address_action_verification_address_key
	city = var.address_action_verification_city
	company_name = var.address_action_verification_company_name
	contributor_class = var.address_action_verification_contributor_class
	country = var.address_action_verification_country
	county = var.address_action_verification_county
	department_name = var.address_action_verification_department_name
	email_address = var.address_action_verification_email_address
	first_name = var.address_action_verification_first_name
	internal_number = var.address_action_verification_internal_number
	job_title = var.address_action_verification_job_title
	last_name = var.address_action_verification_last_name
	line1 = var.address_action_verification_line1
	line2 = var.address_action_verification_line2
	line3 = var.address_action_verification_line3
	line4 = var.address_action_verification_line4
	middle_name = var.address_action_verification_middle_name
	municipal_inscription = var.address_action_verification_municipal_inscription
	phone_country_code = var.address_action_verification_phone_country_code
	phone_number = var.address_action_verification_phone_number
	postal_code = var.address_action_verification_postal_code
	province = var.address_action_verification_province
	state = var.address_action_verification_state
	state_inscription = var.address_action_verification_state_inscription
	street_name = var.address_action_verification_street_name
	street_number = var.address_action_verification_street_number
}
```

## Argument Reference

The following arguments are supported:

* `address_key` - (Optional) Address identifier.
* `city` - (Optional) Name of the city.
* `company_name` - (Optional) Name of the customer company.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `contributor_class` - (Optional) Contributor class of the customer company.
* `country` - (Optional) Country of the address.
* `county` - (Optional) County of the address.
* `department_name` - (Optional) Department name of the customer company.
* `email_address` - (Optional) Contact person email address.
* `first_name` - (Optional) First name of the contact person.
* `internal_number` - (Optional) Internal number of the customer company.
* `job_title` - (Optional) Job title of the contact person.
* `last_name` - (Optional) Last name of the contact person.
* `line1` - (Optional) Address line 1.
* `line2` - (Optional) Address line 2.
* `line3` - (Optional) Address line 3.
* `line4` - (Optional) Address line 4.
* `middle_name` - (Optional) Middle name of the contact person.
* `municipal_inscription` - (Optional) Municipal Inscription.
* `osp_home_region` - (Required) The home region's public name of the logged in user. 
* `phone_country_code` - (Optional) Phone country code of the contact person.
* `phone_number` - (Optional) Phone number of the contact person.
* `postal_code` - (Optional) Post code of the address.
* `province` - (Optional) Province of the address.
* `state` - (Optional) State of the address.
* `state_inscription` - (Optional) State Inscription.
* `street_name` - (Optional) Street name of the address.
* `street_number` - (Optional) Street number of the address.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `address` - Address details model.
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
* `quality` - Address quality type.
* `verification_code` - Address verification code.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Address Action Verification
	* `update` - (Defaults to 20 minutes), when updating the Address Action Verification
	* `delete` - (Defaults to 20 minutes), when destroying the Address Action Verification


## Import

AddressActionVerifications can be imported using the `id`, e.g.

```
$ terraform import oci_osp_gateway_address_action_verification.test_address_action_verification "id"
```

