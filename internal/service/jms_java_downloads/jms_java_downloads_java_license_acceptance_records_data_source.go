// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_java_downloads

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaDownloadsJavaLicenseAcceptanceRecordsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsJavaDownloadsJavaLicenseAcceptanceRecords,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"search_by_user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"java_license_acceptance_record_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(JmsJavaDownloadsJavaLicenseAcceptanceRecordResource()),
						},
					},
				},
			},
		},
	}
}

func readJmsJavaDownloadsJavaLicenseAcceptanceRecords(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaLicenseAcceptanceRecordsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaDownloadsJavaLicenseAcceptanceRecordsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_java_downloads.JavaDownloadClient
	Res    *oci_jms_java_downloads.ListJavaLicenseAcceptanceRecordsResponse
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordsDataSourceCrud) Get() error {
	request := oci_jms_java_downloads.ListJavaLicenseAcceptanceRecordsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if licenseType, ok := s.D.GetOkExists("license_type"); ok {
		request.LicenseType = oci_jms_java_downloads.ListJavaLicenseAcceptanceRecordsLicenseTypeEnum(licenseType.(string))
	}

	if searchByUser, ok := s.D.GetOkExists("search_by_user"); ok {
		tmp := searchByUser.(string)
		request.SearchByUser = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_jms_java_downloads.ListJavaLicenseAcceptanceRecordsStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_java_downloads")

	response, err := s.Client.ListJavaLicenseAcceptanceRecords(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListJavaLicenseAcceptanceRecords(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsJavaDownloadsJavaLicenseAcceptanceRecordsDataSource-", JmsJavaDownloadsJavaLicenseAcceptanceRecordsDataSource(), s.D))
	resources := []map[string]interface{}{}
	javaLicenseAcceptanceRecord := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, JavaLicenseAcceptanceRecordSummaryToMap(item))
	}
	javaLicenseAcceptanceRecord["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsJavaDownloadsJavaLicenseAcceptanceRecordsDataSource().Schema["java_license_acceptance_record_collection"].Elem.(*schema.Resource).Schema)
		javaLicenseAcceptanceRecord["items"] = items
	}

	resources = append(resources, javaLicenseAcceptanceRecord)
	if err := s.D.Set("java_license_acceptance_record_collection", resources); err != nil {
		return err
	}

	return nil
}
