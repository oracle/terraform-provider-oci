// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_certificates_management_association", CertificatesManagementAssociationDataSource())
	tfresource.RegisterDatasource("oci_certificates_management_associations", CertificatesManagementAssociationsDataSource())
	tfresource.RegisterDatasource("oci_certificates_management_ca_bundle", CertificatesManagementCaBundleDataSource())
	tfresource.RegisterDatasource("oci_certificates_management_ca_bundles", CertificatesManagementCaBundlesDataSource())
	tfresource.RegisterDatasource("oci_certificates_management_certificate", CertificatesManagementCertificateDataSource())
	tfresource.RegisterDatasource("oci_certificates_management_certificate_authorities", CertificatesManagementCertificateAuthoritiesDataSource())
	tfresource.RegisterDatasource("oci_certificates_management_certificate_authority", CertificatesManagementCertificateAuthorityDataSource())
	tfresource.RegisterDatasource("oci_certificates_management_certificate_authority_version", CertificatesManagementCertificateAuthorityVersionDataSource())
	tfresource.RegisterDatasource("oci_certificates_management_certificate_authority_versions", CertificatesManagementCertificateAuthorityVersionsDataSource())
	tfresource.RegisterDatasource("oci_certificates_management_certificate_version", CertificatesManagementCertificateVersionDataSource())
	tfresource.RegisterDatasource("oci_certificates_management_certificate_versions", CertificatesManagementCertificateVersionsDataSource())
	tfresource.RegisterDatasource("oci_certificates_management_certificates", CertificatesManagementCertificatesDataSource())
}
