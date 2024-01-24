// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package license_manager

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_license_manager "github.com/oracle/oci-go-sdk/v65/licensemanager"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LicenseManagerProductLicenseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLicenseManagerProductLicense,
		Read:     readLicenseManagerProductLicense,
		Update:   updateLicenseManagerProductLicense,
		Delete:   deleteLicenseManagerProductLicense,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_vendor_oracle": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"license_unit": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"images": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"listing_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"package_version": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"listing_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"publisher": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"vendor_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"active_license_record_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_over_subscribed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_unlimited": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_description": {
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
			"total_active_license_unit_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_license_record_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_license_units_consumed": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func createLicenseManagerProductLicense(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerProductLicenseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.CreateResource(d, sync)
}

func readLicenseManagerProductLicense(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerProductLicenseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.ReadResource(sync)
}

func updateLicenseManagerProductLicense(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerProductLicenseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLicenseManagerProductLicense(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerProductLicenseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LicenseManagerProductLicenseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_license_manager.LicenseManagerClient
	Res                    *oci_license_manager.ProductLicense
	DisableNotFoundRetries bool
}

func (s *LicenseManagerProductLicenseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *LicenseManagerProductLicenseResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *LicenseManagerProductLicenseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_license_manager.LifeCycleStateActive),
	}
}

func (s *LicenseManagerProductLicenseResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *LicenseManagerProductLicenseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_license_manager.LifeCycleStateDeleted),
	}
}

func (s *LicenseManagerProductLicenseResourceCrud) Create() error {
	request := oci_license_manager.CreateProductLicenseRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if images, ok := s.D.GetOkExists("images"); ok {
		interfaces := images.([]interface{})
		tmp := make([]oci_license_manager.ImageDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "images", stateDataIndex)
			converted, err := s.mapToImageDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("images") {
			request.Images = tmp
		}
	}

	if isVendorOracle, ok := s.D.GetOkExists("is_vendor_oracle"); ok {
		tmp := isVendorOracle.(bool)
		request.IsVendorOracle = &tmp
	}

	if licenseUnit, ok := s.D.GetOkExists("license_unit"); ok {
		request.LicenseUnit = oci_license_manager.LicenseUnitEnum(licenseUnit.(string))
	}

	if vendorName, ok := s.D.GetOkExists("vendor_name"); ok {
		tmp := vendorName.(string)
		request.VendorName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "license_manager")

	response, err := s.Client.CreateProductLicense(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProductLicense
	return nil
}

func (s *LicenseManagerProductLicenseResourceCrud) Get() error {
	request := oci_license_manager.GetProductLicenseRequest{}

	tmp := s.D.Id()
	request.ProductLicenseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "license_manager")

	response, err := s.Client.GetProductLicense(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProductLicense
	return nil
}

func (s *LicenseManagerProductLicenseResourceCrud) Update() error {
	request := oci_license_manager.UpdateProductLicenseRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if images, ok := s.D.GetOkExists("images"); ok {
		interfaces := images.([]interface{})
		tmp := make([]oci_license_manager.ImageDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "images", stateDataIndex)
			converted, err := s.mapToImageDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("images") {
			request.Images = tmp
		}
	}

	tmp := s.D.Id()
	request.ProductLicenseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "license_manager")

	response, err := s.Client.UpdateProductLicense(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProductLicense
	return nil
}

func (s *LicenseManagerProductLicenseResourceCrud) Delete() error {
	request := oci_license_manager.DeleteProductLicenseRequest{}

	tmp := s.D.Id()
	request.ProductLicenseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "license_manager")

	_, err := s.Client.DeleteProductLicense(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func (s *LicenseManagerProductLicenseResourceCrud) SetData() error {
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

func (s *LicenseManagerProductLicenseResourceCrud) mapToImageDetails(fieldKeyFormat string) (oci_license_manager.ImageDetails, error) {
	result := oci_license_manager.ImageDetails{}

	if listingId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "listing_id")); ok {
		tmp := listingId.(string)
		result.ListingId = &tmp
	}

	if packageVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "package_version")); ok {
		tmp := packageVersion.(string)
		result.PackageVersion = &tmp
	}

	return result, nil
}

func ImageResponseToMap(obj oci_license_manager.ImageResponse) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ListingId != nil {
		result["listing_id"] = string(*obj.ListingId)
	}

	if obj.ListingName != nil {
		result["listing_name"] = string(*obj.ListingName)
	}

	if obj.PackageVersion != nil {
		result["package_version"] = string(*obj.PackageVersion)
	}

	if obj.Publisher != nil {
		result["publisher"] = string(*obj.Publisher)
	}

	return result
}

func ProductLicenseSummaryToMap(obj oci_license_manager.ProductLicenseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActiveLicenseRecordCount != nil {
		result["active_license_record_count"] = int(*obj.ActiveLicenseRecordCount)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	images := []interface{}{}
	for _, item := range obj.Images {
		images = append(images, ImageResponseToMap(item))
	}
	result["images"] = images

	if obj.IsOverSubscribed != nil {
		result["is_over_subscribed"] = bool(*obj.IsOverSubscribed)
	}

	if obj.IsUnlimited != nil {
		result["is_unlimited"] = bool(*obj.IsUnlimited)
	}

	if obj.IsVendorOracle != nil {
		result["is_vendor_oracle"] = bool(*obj.IsVendorOracle)
	}

	result["license_unit"] = string(obj.LicenseUnit)

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.StatusDescription != nil {
		result["status_description"] = string(*obj.StatusDescription)
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

	if obj.TotalActiveLicenseUnitCount != nil {
		result["total_active_license_unit_count"] = int(*obj.TotalActiveLicenseUnitCount)
	}

	if obj.TotalLicenseRecordCount != nil {
		result["total_license_record_count"] = int(*obj.TotalLicenseRecordCount)
	}

	if obj.TotalLicenseUnitsConsumed != nil {
		result["total_license_units_consumed"] = float64(*obj.TotalLicenseUnitsConsumed)
	}

	if obj.VendorName != nil {
		result["vendor_name"] = string(*obj.VendorName)
	}

	return result
}
