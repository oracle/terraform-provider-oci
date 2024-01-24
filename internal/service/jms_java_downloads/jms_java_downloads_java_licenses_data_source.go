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

func JmsJavaDownloadsJavaLicensesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsJavaDownloadsJavaLicenses,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"java_license_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"license_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"license_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readJmsJavaDownloadsJavaLicenses(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaLicensesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaDownloadsJavaLicensesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_java_downloads.JavaDownloadClient
	Res    *oci_jms_java_downloads.ListJavaLicensesResponse
}

func (s *JmsJavaDownloadsJavaLicensesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaDownloadsJavaLicensesDataSourceCrud) Get() error {
	request := oci_jms_java_downloads.ListJavaLicensesRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if licenseType, ok := s.D.GetOkExists("license_type"); ok {
		request.LicenseType = oci_jms_java_downloads.ListJavaLicensesLicenseTypeEnum(licenseType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_java_downloads")

	response, err := s.Client.ListJavaLicenses(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListJavaLicenses(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsJavaDownloadsJavaLicensesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsJavaDownloadsJavaLicensesDataSource-", JmsJavaDownloadsJavaLicensesDataSource(), s.D))
	resources := []map[string]interface{}{}
	javaLicense := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, JavaLicenseSummaryToMap(item))
	}
	javaLicense["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsJavaDownloadsJavaLicensesDataSource().Schema["java_license_collection"].Elem.(*schema.Resource).Schema)
		javaLicense["items"] = items
	}

	resources = append(resources, javaLicense)
	if err := s.D.Set("java_license_collection", resources); err != nil {
		return err
	}

	return nil
}

func JavaLicenseSummaryToMap(obj oci_jms_java_downloads.JavaLicenseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["license_type"] = string(obj.LicenseType)

	if obj.LicenseUrl != nil {
		result["license_url"] = string(*obj.LicenseUrl)
	}

	return result
}
