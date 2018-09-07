# Cloudera on OCI automation with Terraform
Included here are Terraform templates for deploying a fully configured Cloudera Enterprise Data Hub (EDH) instance or cluster on OCI. 

## Sandbox

This is a great starting point for customers wanting to explore the power and functionality of EDH on OCI.  This deployment consists of a single instance running the Cloudera Docker container.  This is a good fit for individuals who want to explore Cloudera on OCI, while maintaining a cost-effective bottom line.  This is not a good fit for multiple users, development efforts, or large datasets.

Minimum Instance: VM.Standard1.8

Suggested Instance: VM.Standard2.8

## Development
For small implementations, this is the next step up for running EDH on OCI.  This deployment consists of five instances - a bastion host, utility Host, and three workers.  This environment provides a much higher HDFS storage capacity, along with a compute and memory resources for use with a variety of big-data workloads.   This environment is not a good fit for customers who want highly available services, as the reduced infrastructure footprint does not support high availability.

| Minimum Worker Instance | Minimum Bastion Instance | Minimum Utility Instance | 
| :---------------------: |  :---------------------: |  :---------------------: |
| BM.Standard1.16 with (3) 700GB Block Storage devices per worker | VM.Standard1.4 | VM.Standard1.8  |                   


| Suggested Worker Instance | Suggested Bastion Instance | Suggested Utility Instance | 
| :---------------------: |  :---------------------: |  :---------------------: |
| BM.Standard2.24 with (3) 1TB Block Storage devices per worker | VM.Standard2.4 | VM.Standard2.8 |

## Production Starter
This is the most powerful pre-configured option, providing high density and performance for EDH on OCI. This environment provides high availability, and is an appropriate entry point for scaling up a production Big Data practice. For larger scale deployments, see Custom (N-Node).

| Minimum Worker Instance | Minimum Bastion Instance | Minimum Utility & Master Instance | 
| :---------------------: |  :---------------------: |  :---------------------: |
| BM.DenseIO1.36 | VM.Standard1.4 | VM.Standard1.8 |                                

| Suggested Worker Instance | Suggested Bastion Instance | Suggested Utility & Master Instance | 
| :---------------------: |  :---------------------: |  :---------------------: |
| BM.DenseIO2.52 | VM.Standard2.4 | VM.Standard2.8 |                                   

## Custom (N-Node)
OCI also supports N-Node EDH implementations for customers whose needs may exceed the performance or capacity limitations of the largest pre-set cluster configuration.   Please contact OCI for more information.  We are happy to provide guidance on optimizing cluster deployment, and have an automated solution for dynamic cluster scaling into the thousands of nodes.

| Minimum Worker Instance | Minimum Bastion Instance | Minimum Utility & Master Instance |
| :---------------------: |  :---------------------: |  :---------------------: |
| BM.DenseIO1.36 | VM.Standard1.4 | VM.Standard1.8 |    

| Suggested Worker Instance | Suggested Bastion Instance | Suggested Utility & Master Instance |
| :---------------------: |  :---------------------: |  :---------------------: |
| BM.DenseIO2.52 | VM.Standard2.4 | VM.Standard2.16 |  

## Scripts

This is a top level directory hosting shared scripts used by AD-Spanning, Development, N-Node, and Production templates.   As such it should be placed in the same directory heirarchy found here, otherwise Terraform remote-execution references will fail to find the files for deployment.

# How to use these templates
In addition to an active tenancy on OCI, you will need a functional installation of Terraform, and an API key for a privileged user in the tenancy.  See these documentation links for more information:

[Getting Started with Terraform on OCI](https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/terraformgetstarted.htm)

[How to Generate an API Signing Key](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/apisigningkey.htm#How)

Once the pre-requisites are in place, you will need to copy the templates from this repository to where you have Terraform installed.  Refer to the README.md for each template for additional deployment instructions.

