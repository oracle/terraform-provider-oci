# oci\_core\_drg

Provide a drg attachment resource.

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
* `compartment_id` - The OCID of the compartment containing the DRG.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - The DRG's Oracle ID (OCID).
* `state` - The DRG's current state: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED].
* `time_created` - The date and time the image was created.

