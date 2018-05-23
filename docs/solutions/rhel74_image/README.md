This example provides a method to generate a RHEL 7.4 image for use by both VM and BM shapes.

UPDATED 23 May 2018!  Please read below!

What is fixed/changed in this release:

- Fixed issue with OCI-CLI and pip that was preventing the process from working.  The original 
  process attempted to install OCI-CLI as a global resource in the global python repo...this broke
  after some changes in pip.  So now we install pip locally to each user.
- Changed our blank image OCID to use a new global value.  This makes the process *more* portable across
  regions (but not completely...see next bullet).
- Updated the ipxe server image list to include the LON region.  LON is now fully supported on this process.
- *** IMPORTANT CHANGE *** Updated process to use Instance Principal authorization method for OCI-CLI.  This
  removes the requirement to upload the user private key into the ipxe server.  *HOWEVER*, it does REQUIRE 
  the creation of a Dynamic Group for the Compartment this is being executed in.  This is included in the
  prerequisites section below.  The new process WILL NOT WORK without an Instance Principal (also known
  as a Dynamic Group).
  
There are several prerequisites for using this process:

1. *** UPDATE *** You MUST setup a Dynamic Group for the Compartment in which you are going to run this process.  The
   Dynamic Group allows the ipxe instance itself to authenticate to OCI so that no user configuration is 
   needed to create the image.  
   
   Information on how to create a Dynamic Group can be found here:
   https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Tasks/managingdynamicgroups.htm?Highlight=Dynamic%20Group
   
   In short, from the console:
   - Get the Compartment OCID for the Compartment you will be using.
   - In Identity, select Dynamic Groups
   - Click on the Create Dynamic Group box
   - Specify a name for the group
   - Click on the link labeled "Launch Rule Builder"
   - Select 'in Compartment ID' as the Resource Attribute
   - Enter the Compartment OCID in the Value box
   - Click on the Add Rule button
   - Click on the Create Dynamic Group button.
   
   The compartment is now enabled for Instance Principals.  If you do not want this after the image is
   created, simply delete the Dynamic Group AFTER the image is created.  Deletion of the DG will not
   affect the usability of the image.
2. You MUST have a valid RedHat account with subscriptions available.  The TF template needs a 
   RH Username and Password to allow you to temporarily subscribe the instance that is building the image 
   and get access to the various RH repos.
2. The template expects pre-configured VCNs and Subnets.  
3. You need to provide a URL that points to the RHEL 7.4 ISO.  This URL must contain the name of the ISO, 
   with an '.iso' extension.  An OCI Pre-Authenticated Request (PAR) works well for this operation.  How to create
   OCI PARs can be found here: https://docs.us-phoenix-1.oraclecloud.com/Content/Object/Tasks/managingobjects.htm#par.
4. The template uses filters that expect unique Compartment, VCN and Subnet names.
	NOTE: The root compartment CANNOT be used for this process.
5. The following must be specified in your shell environment (prefixed with TF_VAR_ of course):
    - tenancy_ocid
    - user_ocid
    - fingerprint
    - private_key_path
    - private_key_password (if required)
    - ssh_public_key (the actual public key, not the file)
    - region
6. The subnet to be used must have the following configuration:
	- Port 80 TCP must be allowed on the subnet
	- All ICMP traffic must be allowed on the subnet (ICMP All)
7. *** UPDATE *** Ensure that either uuencode/uudecode is available in your shell or that the sharutils package has been
   installed on the OS where this terraform template will be executed.  Uuencode is used to dynamically load
   the configuration files needed for the RHEL build process.

NOTE: A template env-vars file is provided as part of this example.  Simply complete the items inside the template and source the result into your shell by using:

. ./env-vars    

Using this template is simple:

1. Set your environment variables
2. Open the configuration.tf file and substitute the values in each of the sections appropriate to your environment
	NOTE: The AD is specified as either 'AD-x' or 'ad-x' where x is the AD number you wish to use for the process.
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
