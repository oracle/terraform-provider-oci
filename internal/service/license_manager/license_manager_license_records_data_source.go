// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package license_manager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_license_manager "github.com/oracle/oci-go-sdk/v65/licensemanager"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LicenseManagerLicenseRecordsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLicenseManagerLicenseRecords,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"product_license_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"license_record_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(LicenseManagerLicenseRecordResource()),
						},
					},
				},
			},
		},
	}
}

func readLicenseManagerLicenseRecords(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerLicenseRecordsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.ReadResource(sync)
}

type LicenseManagerLicenseRecordsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_license_manager.LicenseManagerClient
	Res    *oci_license_manager.ListLicenseRecordsResponse
}

func (s *LicenseManagerLicenseRecordsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LicenseManagerLicenseRecordsDataSourceCrud) Get() error {
	request := oci_license_manager.ListLicenseRecordsRequest{}

	if productLicenseId, ok := s.D.GetOkExists("product_license_id"); ok {
		tmp := productLicenseId.(string)
		request.ProductLicenseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "license_manager")

	response, err := s.Client.ListLicenseRecords(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLicenseRecords(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LicenseManagerLicenseRecordsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LicenseManagerLicenseRecordsDataSource-", LicenseManagerLicenseRecordsDataSource(), s.D))
	resources := []map[string]interface{}{}
	licenseRecord := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LicenseRecordSummaryToMap(item))
	}
	licenseRecord["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LicenseManagerLicenseRecordsDataSource().Schema["license_record_collection"].Elem.(*schema.Resource).Schema)
		licenseRecord["items"] = items
	}

	resources = append(resources, licenseRecord)
	if err := s.D.Set("license_record_collection", resources); err != nil {
		return err
	}

	return nil
}
