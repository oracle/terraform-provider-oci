    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***

# MS SQL Always On

This Terraform template implements the infrastructure needed to fully deploy an MS SQL Always On cluster across availability domains.  The cluster can be spread across either 2 or 3 ADs, depending on your requirements, but cannot be located in a single AD.

This *ONLY* deploys the infrastructure needed to build MS SQL Always On, and *DOES NOT* perform the installation and configuration needed at the operating system.  That is covered as an exercise left to the reader in the [Deploying Microsoft SQL Server Always On Availability Groups](https://docs.us-phoenix-1.oraclecloud.com/Content/Resources/Assets/deploy-sql-server-availability-groups.pdf) white paper.  

The template is configured by modifying the variables contained in two files:

- configuration.tf - generalized configuration for the environment as a whole.
- sql.tf - configuration specific to the implementation of MS SQL Always On.

To configure, simply edit these two files, apply values that are appropriate to your environment, and run 'terraform plan; terraform apply'.  Before attempting to run the template, setup your environment by performing the following:

- Edit the env-vars file.
- Set the values indicated in the file
- Source the file into your existing environment by running:

. ./env-vars

Some items of note:

- This has only been tested using Windows 2012 R2 images.  The template and accompanying white paper has not been deployed using Windows 2016 Datacenter within the OCI environment.
- The image OCID listed in the configuration.tf file is the one used for testing.  You should update this with the image OCID of the latest revision of Windows 2012 R2.  For a list of image OCIDs, see https://docs.us-phoenix-1.oraclecloud.com/Content/Resources/Assets/OracleProvidedImageOCIDs.pdf
- If you choose *NOT* to specify an existing compartment in which to deploy MS SQL Always On, one will be created for you.  Understand that compartments, once created, *CANNOT* be destroyed.
- Do not edit any of the files other than configuration.tf and sql.tf. Look, but touch at your own peril.
- The template itself has been tested against Terraform version 0.11.3 and OCI Provider 2.0.6.  Running against older versions of both Terraform and the provider may provide unpredicable results.
- For this template to work, you must identify your OCI Home Region.  This may be different than the region in which you are deploying this template.  To identify your home region, do the following:
	- Log into the OCI console for the tenancy in which you are going to deploy.
	- In the upper left hand side of the console there is a hyperlink that highlights your tenancy name.  Click on the hyperlink.
	- On the left hand side of the resulting web page is a field labeled "Home Region".  This is the value to insert into the env-vars file.
- The infrastructure can be deployed in multiple compartments, or within the same compartment by specifying a unique, non-empty value in the 'label-prefix' variable contained in configuration.tf.  If deploying multiple times in the same compartment, this must be done in order to prevent confusion.

A sample ConfigurationFile.ini file has been included for SQL.  Use at your own discretion.  

Enjoy and happy SQL'ing!
