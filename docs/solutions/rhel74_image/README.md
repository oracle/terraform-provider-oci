This example provides a method to generate a RHEL 7.4 image for use by both VM and BM shapes.
There are several prerequisites for using this process:

1. You MUST have a valid RedHat account with subscriptions available.  The TF template needs a 
   RH Username and Password to allow you to temporarily subscribe the instance that is building the image 
   and get access to the various RH repos.
2. The template expects pre-configured VCNs and Subnets.  
3. You need to have a bucket in the OCI object store that contains the RHEL 7.4 iso.  The bucket must be in 
   your tenancy, and must be accessible by the user executing the build.
4. The template uses filters that expect unique Compartment, VCN and Subnet names.
5. The following must be specified in your shell environment (prefixed with TF_VAR_ of course:
    - tenancy_ocid
    - user_ocid
    - fingerprint
    - private_key_path
    - private_key_password (if required)
    - ssh_public_key (the actual public key, not the file)
    - region
    
    There is a tool included in this directory, settfenv.py, that will generate a file that can be sourced into
    your shell, containing all these pieces of information.  The tool assumes you have previously setup
    the OCI python libraries (preferred) or OCI CLI locally.

Using this template is simple:

1. Set your environment variables
2. Open the configuration.tf file and substitute the values in each of the sections appropriate to your environment
3. Execute 'terraform plan; terraform apply'
4. Get coffee or favorite beverage...
5. After your image is created, execute 'terraform destroy -force' (there will not be a resource to actually kill,
   so force is required).

What happens in the background:
The template generates a script that embeds all the configuration files needed to build the iPXE server, extract the ISO
boot the instance used to load RHEL, causes RHEL to load, builds the image, destroys the build instance, and finally destroys the iPXE server.  You are left with a custom image named "RHEL_74" in your environment.

NOTE: The source configuration files for the iPXE server are included here.  It is *STRONGLY* recommended that they not be 
      altered.

Enjoy.
