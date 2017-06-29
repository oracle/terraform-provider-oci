# baremetal\_load\_balancer\_backends

Provide a load balancer backends.

## Example Usage

```
data "baremetal_load_balancer_backends" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
  backendset_name  = "stub_backendset_name"
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The OCID of the load balancer.
* `backendset_name` - (Required) The name of the backend set.

## Attributes Reference
* `backends` - The list of backends

## Backend Reference
* `name` - A name to uniquely identify this backend server in the backend set.
* `ip_address` - The IP address of the backend server.
* `port` - The communication port for the backend server.
* `backup` - Whether the load balancer should treat this server as a backup unit.
* `drain` - Whether the load balancer should drain this server.
* `offline` - Whether the load balancer should treat this server as offline. 
* `weight` - The load balancing policy weight assigned to the server.

