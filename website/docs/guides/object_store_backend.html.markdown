---
layout: "oci"
page_title: "Object Store Backend"
sidebar_current: "docs-oci-guide-object_store_backend"
description: |-
  The Oracle Cloud Infrastructure provider. Object Store Backend
---

## Using the Object Store for Terraform State Files
You can store [Terraform state files](https://www.terraform.io/docs/state/index.html) in the 
Oracle Cloud Infrastructure Object Storage. Doing so requires that you configure a backend using one of the Terraform backend types.

Terraform supports various backend types to allow flexibility in how state files are loaded into Terraform. (For more 
information, see [Terraform Backend Types](https://www.terraform.io/docs/backends/types/index.html).) For our purposes, we address two of these approaches:

- Using an HTTP remote state backend
- Using an S3-compatible remote state backend

### Using an HTTP Backend

Using the [HTTP backend type](https://www.terraform.io/docs/backends/types/http.html) allows you to store state using a simple REST client. With the HTTP backend type, you can 
easily fetch, update, and purge state using the HTTP GET, POST, and DELETE methods.

To configure the HTTP backend to store your Oracle Cloud Infrastructure Terraform state files, do the following:


#### Create a Pre-Authenticated Request

Creating a pre-authenticated request in Oracle Object Storage enables accessing a bucket or object in the Oracle Cloud 
Infrastructure without needing to provide credentials. To do so, you must create a pre-authenticated request that has 
read/write permissions to the object store where you intend to save the Terraform state file. You can do so in any of 
three ways: by using the Console UI, by using the command line interface (CLI), or by using the REST APIs.

>    **Note**  
A state file must exist in the bucket before you create the pre-authenticated request. This file can be an existing state file, or an empty file for the initial state.

For guidance, see [Using Pre-Authenticated Requests](https://docs.cloud.oracle.com/iaas/Content/Object/Tasks/usingpreauthenticatedrequests.htm).


#### Upload Existing State

If you have an existing state file, you can upload it using Curl to make an HTTP Put request to the object store URL, as shown here:

```sh
curl -X PUT -H "Content-Type: text/plain" --data-binary "@path/to/local/tfstate" http://<prefix>/<my-access-uri>
```


#### Configure HTTP as a Terraform Backend

The [HTTP backend type](https://www.terraform.io/docs/backends/types/http.html) stores state using a simple REST client 
and allows you to easily fetch, update, and purge state using the HTTP GET, POST, and DELETE methods.

The access URI for addressing Oracle Cloud Infrastructure Terraform configurations must be of the form: 
https://objectstorage.us-phoenix-1.oraclecloud.com/my-access-uri (where region and access URI are specific to you).

For more example configuration and state files that reference code, and a summary of configuration variables, 
see [Standard Backends: HTTP](https://www.terraform.io/docs/backends/types/http.html).

Following is an example Terraform configuration. The region in the URL can be something other than the Phoenix region.

```hcl-terraform
terraform {
   backend "http" {
     address = "https://objectstorage.us-phoenix-1.oraclecloud.com/<my-access-uri>" update_method = "PUT" }
}
```


#### Reinitialize Terraform

Finally, you must reinitialize Terraform and then run the apply command, as shown following.

```sh
terraform init
terraform apply
```

After completing these steps, you are able to use Oracle Cloud Infrastructure as the backend for storing Terraform state files.


### Using an S3-Compatible Backend

Configuring the S3-compatible backend requires that the account be enabled with S3 authentication keys, which are set on a per-user basis.

1. In the Console, go to your user account and create a customer secret key. For more guidance on creating customer secret keys, 
see [Working with Amazon S3 Compatibility API Keys](https://docs.cloud.oracle.com/Content/Identity/Tasks/managingcredentials.htm#s3).

After generating the customer secret key, take note of the access key and secret key values displayed the Console.

2. Set the location for the credentials file. The default location is `~/.aws/credentials`. You can set an alternate location by using the S3 backend `shared_credentials_file` option. 
    
    > **Warning**  
    Never set the access_key and the secret_key attributes in the same Terraform backend configuration, since this creates a security risk.

3. Configure the `[default]` entry in the credentials file with the appropriate object storage credentials. 
The file can contain any number of credential profiles. If you provide a different profile name, you must also 
update the backend `profile` option in your Terraform configuration file.
    
    Following is an example of Object Storage credentials:
    
    ```
    [default]
    aws_access_key_id=81bc020dd274d7386a58852fc5081c231874a137
    aws_secret_access_key=mSTdaWhlbWj3ty4JZXlm0NUZV52xlImWjayJLJ6OH9A=
    ```
    
    Where `aws_access_key_id` and `aws_secret_access_key` are user-specific values given by the Console in step 1. 
    The key values provided in the example are not valid and provided as examples only.

4. Set the object storage endpoint value in the following format: `https://{namespace}.compat.objectstorage.{region}.oraclecloud.com`

    Where `{namespace}` is the namespace of your object storage bucket and `{region}` is the region where your object storage bucket is located.

Following is a full example of an Object Storage backend configuration:

```hcl-terraform
terraform {
  backend "s3" {
    bucket   = "terraform-state"
    key      = "terraform.tfstate"
    region   = "us-phoenix-1"
    endpoint = "https://mybucketnamespace.compat.objectstorage.us-phoenix-1.oraclecloud.com"

    skip_region_validation      = true
    skip_credentials_validation = true
    skip_requesting_account_id  = true
    skip_get_ec2_platforms      = true
    skip_metadata_api_check     = true
    force_path_style            = true
  }
}
```

The S3 backend configuration can also be used for the terraform_remote_state data source to enable sharing state across Terraform projects.

Once you have configured the backend, you must run `terraform init` to finish the setup. 
If you already have an existing `terraform.tfstate` file, then Terraform prompts you to confirm that the current state file is the one to upload to the remote state.


### For More Information

- [Using Pre-Authenticated Requests](https://docs.cloud.oracle.com/iaas/Content/Object/Tasks/usingpreauthenticatedrequests.htm)
- [State Files](https://www.terraform.io/docs/state/index.html)
- [Terraform Backend Types](https://www.terraform.io/docs/backends/types/index.html)
