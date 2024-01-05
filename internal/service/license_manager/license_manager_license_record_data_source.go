// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package license_manager

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_license_manager "github.com/oracle/oci-go-sdk/v65/licensemanager"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LicenseManagerLicenseRecordDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["license_record_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(LicenseManagerLicenseRecordResource(), fieldMap, readSingularLicenseManagerLicenseRecord)
}

func readSingularLicenseManagerLicenseRecord(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerLicenseRecordDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.ReadResource(sync)
}

type LicenseManagerLicenseRecordDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_license_manager.LicenseManagerClient
	Res    *oci_license_manager.GetLicenseRecordResponse
}

func (s *LicenseManagerLicenseRecordDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LicenseManagerLicenseRecordDataSourceCrud) Get() error {
	request := oci_license_manager.GetLicenseRecordRequest{}

	if licenseRecordId, ok := s.D.GetOkExists("license_record_id"); ok {
		tmp := licenseRecordId.(string)
		request.LicenseRecordId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "license_manager")

	response, err := s.Client.GetLicenseRecord(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LicenseManagerLicenseRecordDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExpirationDate != nil {
		s.D.Set("expiration_date", s.Res.ExpirationDate.Format(time.RFC3339Nano))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsPerpetual != nil {
		s.D.Set("is_perpetual", *s.Res.IsPerpetual)
	}

	if s.Res.IsUnlimited != nil {
		s.D.Set("is_unlimited", *s.Res.IsUnlimited)
	}

	if s.Res.LicenseCount != nil {
		s.D.Set("license_count", *s.Res.LicenseCount)
	}

	s.D.Set("license_unit", s.Res.LicenseUnit)

	if s.Res.ProductId != nil {
		s.D.Set("product_id", *s.Res.ProductId)
	}

	if s.Res.ProductLicense != nil {
		s.D.Set("product_license", *s.Res.ProductLicense)
	}

	if s.Res.ProductLicenseId != nil {
		s.D.Set("product_license_id", *s.Res.ProductLicenseId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SupportEndDate != nil {
		s.D.Set("support_end_date", s.Res.SupportEndDate.Format(time.RFC3339Nano))
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

	return nil
}
