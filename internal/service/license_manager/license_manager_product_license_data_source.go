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

func LicenseManagerProductLicenseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["product_license_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(LicenseManagerProductLicenseResource(), fieldMap, readSingularLicenseManagerProductLicense)
}

func readSingularLicenseManagerProductLicense(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerProductLicenseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.ReadResource(sync)
}

type LicenseManagerProductLicenseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_license_manager.LicenseManagerClient
	Res    *oci_license_manager.GetProductLicenseResponse
}

func (s *LicenseManagerProductLicenseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LicenseManagerProductLicenseDataSourceCrud) Get() error {
	request := oci_license_manager.GetProductLicenseRequest{}

	if productLicenseId, ok := s.D.GetOkExists("product_license_id"); ok {
		tmp := productLicenseId.(string)
		request.ProductLicenseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "license_manager")

	response, err := s.Client.GetProductLicense(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LicenseManagerProductLicenseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ActiveLicenseRecordCount != nil {
		s.D.Set("active_license_record_count", *s.Res.ActiveLicenseRecordCount)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	images := []interface{}{}
	for _, item := range s.Res.Images {
		images = append(images, ImageResponseToMap(item))
	}
	s.D.Set("images", images)

	if s.Res.IsOverSubscribed != nil {
		s.D.Set("is_over_subscribed", *s.Res.IsOverSubscribed)
	}

	if s.Res.IsUnlimited != nil {
		s.D.Set("is_unlimited", *s.Res.IsUnlimited)
	}

	if s.Res.IsVendorOracle != nil {
		s.D.Set("is_vendor_oracle", *s.Res.IsVendorOracle)
	}

	s.D.Set("license_unit", s.Res.LicenseUnit)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.StatusDescription != nil {
		s.D.Set("status_description", *s.Res.StatusDescription)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalActiveLicenseUnitCount != nil {
		s.D.Set("total_active_license_unit_count", *s.Res.TotalActiveLicenseUnitCount)
	}

	if s.Res.TotalLicenseRecordCount != nil {
		s.D.Set("total_license_record_count", *s.Res.TotalLicenseRecordCount)
	}

	if s.Res.TotalLicenseUnitsConsumed != nil {
		s.D.Set("total_license_units_consumed", *s.Res.TotalLicenseUnitsConsumed)
	}

	if s.Res.VendorName != nil {
		s.D.Set("vendor_name", *s.Res.VendorName)
	}

	return nil
}
