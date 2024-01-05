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

func LicenseManagerTopUtilizedProductLicensesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLicenseManagerTopUtilizedProductLicenses,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_unlimited": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"product_license_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"product_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_license_unit_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"total_units_consumed": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"unit_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readLicenseManagerTopUtilizedProductLicenses(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerTopUtilizedProductLicensesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.ReadResource(sync)
}

type LicenseManagerTopUtilizedProductLicensesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_license_manager.LicenseManagerClient
	Res    *oci_license_manager.ListTopUtilizedProductLicensesResponse
}

func (s *LicenseManagerTopUtilizedProductLicensesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LicenseManagerTopUtilizedProductLicensesDataSourceCrud) Get() error {
	request := oci_license_manager.ListTopUtilizedProductLicensesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isCompartmentIdInSubtree, ok := s.D.GetOkExists("is_compartment_id_in_subtree"); ok {
		tmp := isCompartmentIdInSubtree.(bool)
		request.IsCompartmentIdInSubtree = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "license_manager")

	response, err := s.Client.ListTopUtilizedProductLicenses(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LicenseManagerTopUtilizedProductLicensesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LicenseManagerTopUtilizedProductLicensesDataSource-", LicenseManagerTopUtilizedProductLicensesDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TopUtilizedProductLicensesSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func TopUtilizedProductLicensesSummaryToMap(obj oci_license_manager.TopUtilizedProductLicenseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsUnlimited != nil {
		result["is_unlimited"] = bool(*obj.IsUnlimited)
	}

	if obj.ProductLicenseId != nil {
		result["product_license_id"] = string(*obj.ProductLicenseId)
	}

	if obj.ProductType != nil {
		result["product_type"] = string(*obj.ProductType)
	}

	result["status"] = string(obj.Status)

	if obj.TotalLicenseUnitCount != nil {
		result["total_license_unit_count"] = int(*obj.TotalLicenseUnitCount)
	}

	if obj.TotalUnitsConsumed != nil {
		result["total_units_consumed"] = float64(*obj.TotalUnitsConsumed)
	}

	result["unit_type"] = string(obj.UnitType)

	return result
}
