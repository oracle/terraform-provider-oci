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

func LicenseManagerProductLicensesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLicenseManagerProductLicenses,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"product_license_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(LicenseManagerProductLicenseResource()),
						},
					},
				},
			},
		},
	}
}

func readLicenseManagerProductLicenses(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerProductLicensesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.ReadResource(sync)
}

type LicenseManagerProductLicensesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_license_manager.LicenseManagerClient
	Res    *oci_license_manager.ListProductLicensesResponse
}

func (s *LicenseManagerProductLicensesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LicenseManagerProductLicensesDataSourceCrud) Get() error {
	request := oci_license_manager.ListProductLicensesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isCompartmentIdInSubtree, ok := s.D.GetOkExists("is_compartment_id_in_subtree"); ok {
		tmp := isCompartmentIdInSubtree.(bool)
		request.IsCompartmentIdInSubtree = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "license_manager")

	response, err := s.Client.ListProductLicenses(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProductLicenses(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LicenseManagerProductLicensesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LicenseManagerProductLicensesDataSource-", LicenseManagerProductLicensesDataSource(), s.D))
	resources := []map[string]interface{}{}
	productLicense := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProductLicenseSummaryToMap(item))
	}
	productLicense["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LicenseManagerProductLicensesDataSource().Schema["product_license_collection"].Elem.(*schema.Resource).Schema)
		productLicense["items"] = items
	}

	resources = append(resources, productLicense)
	if err := s.D.Set("product_license_collection", resources); err != nil {
		return err
	}

	return nil
}
