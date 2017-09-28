# oci\_load\_balancer\_backends

[Backend Reference][f5f4765f]

  [f5f4765f]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/Backend/ "BackendReference"

Provide a backend server that is a member of a load balancer backend set.

## Example Usage

```
data "oci_load_balancer_backends" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
  backendset_name  = "stub_backendset_name"
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The OCID of the load balancer.
* `backendset_name` - (Required) The name of the backend set.

## Attributes Reference
* `backends` - The list of backends.

## Backend Reference
* `name` - A name to uniquely identify this backend server in the backend set. Avoid entering confidential information.
* `ip_address` - The IP address of the backend server.
* `port` - The communication port for the backend server.
* `backup` - Whether the load balancer should treat this server as a backup unit. Example: `true`
* `drain` - Whether the load balancer should drain this server. Example: `true`
* `offline` - Whether the load balancer should treat this server as offline. Example: `true`
* `weight` - The load balancing policy weight assigned to the server.
