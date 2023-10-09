// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_java_downloads

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_jms_java_downloads_java_download_records", JmsJavaDownloadsJavaDownloadRecordsDataSource())
	tfresource.RegisterDatasource("oci_jms_java_downloads_java_download_report", JmsJavaDownloadsJavaDownloadReportDataSource())
	tfresource.RegisterDatasource("oci_jms_java_downloads_java_download_report_content", JmsJavaDownloadsJavaDownloadReportContentDataSource())
	tfresource.RegisterDatasource("oci_jms_java_downloads_java_download_reports", JmsJavaDownloadsJavaDownloadReportsDataSource())
	tfresource.RegisterDatasource("oci_jms_java_downloads_java_download_token", JmsJavaDownloadsJavaDownloadTokenDataSource())
	tfresource.RegisterDatasource("oci_jms_java_downloads_java_download_tokens", JmsJavaDownloadsJavaDownloadTokensDataSource())
	tfresource.RegisterDatasource("oci_jms_java_downloads_java_license", JmsJavaDownloadsJavaLicenseDataSource())
	tfresource.RegisterDatasource("oci_jms_java_downloads_java_license_acceptance_record", JmsJavaDownloadsJavaLicenseAcceptanceRecordDataSource())
	tfresource.RegisterDatasource("oci_jms_java_downloads_java_license_acceptance_records", JmsJavaDownloadsJavaLicenseAcceptanceRecordsDataSource())
	tfresource.RegisterDatasource("oci_jms_java_downloads_java_licenses", JmsJavaDownloadsJavaLicensesDataSource())
}
