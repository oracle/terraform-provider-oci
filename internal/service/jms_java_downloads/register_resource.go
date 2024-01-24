// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_java_downloads

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_jms_java_downloads_java_download_report", JmsJavaDownloadsJavaDownloadReportResource())
	tfresource.RegisterResource("oci_jms_java_downloads_java_download_token", JmsJavaDownloadsJavaDownloadTokenResource())
	tfresource.RegisterResource("oci_jms_java_downloads_java_license_acceptance_record", JmsJavaDownloadsJavaLicenseAcceptanceRecordResource())
}
