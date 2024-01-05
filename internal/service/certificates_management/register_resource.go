// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_certificates_management_ca_bundle", CertificatesManagementCaBundleResource())
	tfresource.RegisterResource("oci_certificates_management_certificate", CertificatesManagementCertificateResource())
	tfresource.RegisterResource("oci_certificates_management_certificate_authority", CertificatesManagementCertificateAuthorityResource())
}
