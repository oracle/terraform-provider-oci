# oci\_core\_drg

**API:** [Drg Reference][18ac8fea]

  [18ac8fea]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Drg/ "DrgReference"

Provide a Dynamic Routing Gateway (DRG) resource.

## Example Usage

```
data "oci_core_drgs" "t" {
  compartment_id = "compartment_id"
  limit = 1
  page = "page"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) The page to fetch.

## Attributes Reference

The following attributes are exported:

* `drgs` - The list of dynamic routing gateways.

## DRG Reference
* `compartment_id` - The OCID of the compartment containing the DRG.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The DRG's Oracle ID (OCID).
* `state` - The DRG's current state: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED].
* `time_created` - The date and time the DRG was created,  in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
