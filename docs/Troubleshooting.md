### Not Authenticated Error when configuring Terraform

_If the Terraform CLI gives an error message like:_

`* oci_core_vcn.resource1: Service error:NotAuthenticated. The required information to complete authentication was not provided or was incorrect.. http status code: 401`

* Verify you have properly set `user_ocid`, `tenancy_ocid`, `fingerprint` and `private_key_path` 
* Verify your `private_key_path` is pointing to your private key and not the corresponding public key
* Verify you have added the corresponding public key to the user account you have specified with `user_ocid`
* Verify the public/private key pairs you are using are of the correct format
  * see: [Required Keys](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/apisigningkey.htm) for details on the correct format and how to generate keys
* Verify the user account is part of a group with the appropriate permissions to perform the actions in the plan you are executing
* Verify your Tenancy has been subscribed to the Region you are targeting in your plan


### Error message after upgrading OCI Terraform Provider

_If the Terraform CLI gives an error message like:_

```
Error asking for user input: 1 error(s) occurred:
 
* provider.oci: dial unix /var/folders/6r/8fk5dmbj4_z3sl0mc_y_fhjw0000gn/T/plugin811254328|netrpc: connect: no such file or directory
```

You are likely using a version of the OCI Terraform Provider that is not compatible with the Terraform binary you have 
installed. For OCI Provider versions v3.x.x and above, a minimum Terraform version of v.0.10.1 is required. 
