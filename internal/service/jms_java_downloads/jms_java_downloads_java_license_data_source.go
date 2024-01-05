// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_java_downloads

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaDownloadsJavaLicenseDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsJavaDownloadsJavaLicense,
		Schema: map[string]*schema.Schema{
			"license_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"license_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularJmsJavaDownloadsJavaLicense(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaLicenseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaDownloadsJavaLicenseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_java_downloads.JavaDownloadClient
	Res    *oci_jms_java_downloads.GetJavaLicenseResponse
}

func (s *JmsJavaDownloadsJavaLicenseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaDownloadsJavaLicenseDataSourceCrud) Get() error {
	request := oci_jms_java_downloads.GetJavaLicenseRequest{}

	if licenseType, ok := s.D.GetOkExists("license_type"); ok {
		request.LicenseType = oci_jms_java_downloads.GetJavaLicenseLicenseTypeEnum(licenseType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_java_downloads")

	response, err := s.Client.GetJavaLicense(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsJavaDownloadsJavaLicenseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsJavaDownloadsJavaLicenseDataSource-", JmsJavaDownloadsJavaLicenseDataSource(), s.D))

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.LicenseUrl != nil {
		s.D.Set("license_url", *s.Res.LicenseUrl)
	}

	return nil
}
