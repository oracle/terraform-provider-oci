// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_ocvp_cluster", OcvpClusterDataSource())
	tfresource.RegisterDatasource("oci_ocvp_clusters", OcvpClustersDataSource())
	tfresource.RegisterDatasource("oci_ocvp_esxi_host", OcvpEsxiHostDataSource())
	tfresource.RegisterDatasource("oci_ocvp_esxi_hosts", OcvpEsxiHostsDataSource())
	tfresource.RegisterDatasource("oci_ocvp_retrieve_password", OcvpRetrievePasswordDataSource())
	tfresource.RegisterDatasource("oci_ocvp_sddc", OcvpSddcDataSource())
	tfresource.RegisterDatasource("oci_ocvp_sddcs", OcvpSddcsDataSource())
	tfresource.RegisterDatasource("oci_ocvp_supported_commitments", OcvpSupportedCommitmentsDataSource())
	tfresource.RegisterDatasource("oci_ocvp_supported_host_shapes", OcvpSupportedHostShapesDataSource())
	tfresource.RegisterDatasource("oci_ocvp_supported_skus", OcvpSupportedSkusDataSource())
	tfresource.RegisterDatasource("oci_ocvp_supported_vmware_software_versions", OcvpSupportedVmwareSoftwareVersionsDataSource())
}
