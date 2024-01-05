// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package license_manager

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_license_manager "github.com/oracle/oci-go-sdk/v65/licensemanager"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LicenseManagerLicenseRecordResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLicenseManagerLicenseRecord,
		Read:     readLicenseManagerLicenseRecord,
		Update:   updateLicenseManagerLicenseRecord,
		Delete:   deleteLicenseManagerLicenseRecord,
		Schema: map[string]*schema.Schema{
			// Required
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_perpetual": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"is_unlimited": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"product_license_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"expiration_date": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"license_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"product_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"support_end_date": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"license_unit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_license": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLicenseManagerLicenseRecord(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerLicenseRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.CreateResource(d, sync)
}

func readLicenseManagerLicenseRecord(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerLicenseRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.ReadResource(sync)
}

func updateLicenseManagerLicenseRecord(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerLicenseRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLicenseManagerLicenseRecord(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerLicenseRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LicenseManagerLicenseRecordResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_license_manager.LicenseManagerClient
	Res                    *oci_license_manager.LicenseRecord
	DisableNotFoundRetries bool
}

func (s *LicenseManagerLicenseRecordResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *LicenseManagerLicenseRecordResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *LicenseManagerLicenseRecordResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_license_manager.LifeCycleStateActive),
	}
}

func (s *LicenseManagerLicenseRecordResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *LicenseManagerLicenseRecordResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_license_manager.LifeCycleStateDeleted),
	}
}

func (s *LicenseManagerLicenseRecordResourceCrud) Create() error {
	request := oci_license_manager.CreateLicenseRecordRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if expirationDate, ok := s.D.GetOkExists("expiration_date"); ok {
		tmp, err := time.Parse(time.RFC3339, expirationDate.(string))
		if err != nil {
			return err
		}
		request.ExpirationDate = &oci_common.SDKTime{Time: tmp}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isPerpetual, ok := s.D.GetOkExists("is_perpetual"); ok {
		tmp := isPerpetual.(bool)
		request.IsPerpetual = &tmp
	}

	if isUnlimited, ok := s.D.GetOkExists("is_unlimited"); ok {
		tmp := isUnlimited.(bool)
		request.IsUnlimited = &tmp
	}

	if licenseCount, ok := s.D.GetOkExists("license_count"); ok {
		tmp := licenseCount.(int)
		request.LicenseCount = &tmp
	}

	if productId, ok := s.D.GetOkExists("product_id"); ok {
		tmp := productId.(string)
		request.ProductId = &tmp
	}

	if productLicenseId, ok := s.D.GetOkExists("product_license_id"); ok {
		tmp := productLicenseId.(string)
		request.ProductLicenseId = &tmp
	}

	if supportEndDate, ok := s.D.GetOkExists("support_end_date"); ok {
		tmp, err := time.Parse(time.RFC3339, supportEndDate.(string))
		if err != nil {
			return err
		}
		request.SupportEndDate = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "license_manager")

	response, err := s.Client.CreateLicenseRecord(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LicenseRecord
	return nil
}

func (s *LicenseManagerLicenseRecordResourceCrud) Get() error {
	request := oci_license_manager.GetLicenseRecordRequest{}

	tmp := s.D.Id()
	request.LicenseRecordId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "license_manager")

	response, err := s.Client.GetLicenseRecord(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LicenseRecord
	return nil
}

func (s *LicenseManagerLicenseRecordResourceCrud) Update() error {
	request := oci_license_manager.UpdateLicenseRecordRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if expirationDate, ok := s.D.GetOkExists("expiration_date"); ok {
		tmp, err := time.Parse(time.RFC3339, expirationDate.(string))
		if err != nil {
			return err
		}
		request.ExpirationDate = &oci_common.SDKTime{Time: tmp}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isPerpetual, ok := s.D.GetOkExists("is_perpetual"); ok {
		tmp := isPerpetual.(bool)
		request.IsPerpetual = &tmp
	}

	if isUnlimited, ok := s.D.GetOkExists("is_unlimited"); ok {
		tmp := isUnlimited.(bool)
		request.IsUnlimited = &tmp
	}

	if licenseCount, ok := s.D.GetOkExists("license_count"); ok {
		tmp := licenseCount.(int)
		request.LicenseCount = &tmp
	}

	tmp := s.D.Id()
	request.LicenseRecordId = &tmp

	if productId, ok := s.D.GetOkExists("product_id"); ok {
		tmp := productId.(string)
		request.ProductId = &tmp
	}

	if supportEndDate, ok := s.D.GetOkExists("support_end_date"); ok {
		tmp, err := time.Parse(time.RFC3339, supportEndDate.(string))
		if err != nil {
			return err
		}
		request.SupportEndDate = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "license_manager")

	response, err := s.Client.UpdateLicenseRecord(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LicenseRecord
	return nil
}

func (s *LicenseManagerLicenseRecordResourceCrud) Delete() error {
	request := oci_license_manager.DeleteLicenseRecordRequest{}

	tmp := s.D.Id()
	request.LicenseRecordId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "license_manager")

	_, err := s.Client.DeleteLicenseRecord(context.Background(), request)
	if err != nil {
		return err
	}

	return nil
}

func (s *LicenseManagerLicenseRecordResourceCrud) SetData() error {
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

func LicenseRecordSummaryToMap(obj oci_license_manager.LicenseRecordSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExpirationDate != nil {
		result["expiration_date"] = obj.ExpirationDate.String()
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsPerpetual != nil {
		result["is_perpetual"] = bool(*obj.IsPerpetual)
	}

	if obj.IsUnlimited != nil {
		result["is_unlimited"] = bool(*obj.IsUnlimited)
	}

	if obj.LicenseCount != nil {
		result["license_count"] = int(*obj.LicenseCount)
	}

	result["license_unit"] = string(obj.LicenseUnit)

	if obj.ProductId != nil {
		result["product_id"] = string(*obj.ProductId)
	}

	if obj.ProductLicense != nil {
		result["product_license"] = string(*obj.ProductLicense)
	}

	if obj.ProductLicenseId != nil {
		result["product_license_id"] = string(*obj.ProductLicenseId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SupportEndDate != nil {
		result["support_end_date"] = obj.SupportEndDate.String()
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
