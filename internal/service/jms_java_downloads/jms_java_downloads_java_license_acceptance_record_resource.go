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

func JmsJavaDownloadsJavaLicenseAcceptanceRecordResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createJmsJavaDownloadsJavaLicenseAcceptanceRecord,
		Read:     readJmsJavaDownloadsJavaLicenseAcceptanceRecord,
		Update:   updateJmsJavaDownloadsJavaLicenseAcceptanceRecord,
		Delete:   deleteJmsJavaDownloadsJavaLicenseAcceptanceRecord,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"license_acceptance_status": {
				Type:     schema.TypeString,
				Required: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"created_by": {
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
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"last_updated_by": {
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
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"time_accepted": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createJmsJavaDownloadsJavaLicenseAcceptanceRecord(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.CreateResource(d, sync)
}

func readJmsJavaDownloadsJavaLicenseAcceptanceRecord(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.ReadResource(sync)
}

func updateJmsJavaDownloadsJavaLicenseAcceptanceRecord(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteJmsJavaDownloadsJavaLicenseAcceptanceRecord(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_jms_java_downloads.JavaDownloadClient
	Res                    *oci_jms_java_downloads.JavaLicenseAcceptanceRecord
	DisableNotFoundRetries bool
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_jms_java_downloads.LifecycleStateCreating),
	}
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_jms_java_downloads.LifecycleStateActive),
		string(oci_jms_java_downloads.LifecycleStateNeedsAttention),
	}
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_jms_java_downloads.LifecycleStateDeleting),
	}
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_jms_java_downloads.LifecycleStateDeleted),
	}
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud) Create() error {
	request := oci_jms_java_downloads.CreateJavaLicenseAcceptanceRecordRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if licenseAcceptanceStatus, ok := s.D.GetOkExists("license_acceptance_status"); ok {
		request.LicenseAcceptanceStatus = oci_jms_java_downloads.LicenseAcceptanceStatusEnum(licenseAcceptanceStatus.(string))
	}

	if licenseType, ok := s.D.GetOkExists("license_type"); ok {
		request.LicenseType = oci_jms_java_downloads.LicenseTypeEnum(licenseType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads")

	response, err := s.Client.CreateJavaLicenseAcceptanceRecord(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.JavaLicenseAcceptanceRecord
	return nil
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud) Get() error {
	request := oci_jms_java_downloads.GetJavaLicenseAcceptanceRecordRequest{}

	tmp := s.D.Id()
	request.JavaLicenseAcceptanceRecordId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads")

	response, err := s.Client.GetJavaLicenseAcceptanceRecord(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.JavaLicenseAcceptanceRecord
	return nil
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud) Update() error {
	request := oci_jms_java_downloads.UpdateJavaLicenseAcceptanceRecordRequest{}

	tmp := s.D.Id()
	request.JavaLicenseAcceptanceRecordId = &tmp

	if licenseAcceptanceStatus, ok := s.D.GetOkExists("license_acceptance_status"); ok {
		request.LicenseAcceptanceStatus = oci_jms_java_downloads.LicenseAcceptanceStatusEnum(licenseAcceptanceStatus.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads")

	response, err := s.Client.UpdateJavaLicenseAcceptanceRecord(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.JavaLicenseAcceptanceRecord
	return nil
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud) Delete() error {
	request := oci_jms_java_downloads.DeleteJavaLicenseAcceptanceRecordRequest{}

	tmp := s.D.Id()
	request.JavaLicenseAcceptanceRecordId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads")

	_, err := s.Client.DeleteJavaLicenseAcceptanceRecord(context.Background(), request)
	return err
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", []interface{}{PrincipalToMap(s.Res.CreatedBy)})
	} else {
		s.D.Set("created_by", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LastUpdatedBy != nil {
		s.D.Set("last_updated_by", []interface{}{PrincipalToMap(s.Res.LastUpdatedBy)})
	} else {
		s.D.Set("last_updated_by", nil)
	}

	s.D.Set("license_acceptance_status", s.Res.LicenseAcceptanceStatus)

	s.D.Set("license_type", s.Res.LicenseType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeLastUpdated != nil {
		s.D.Set("time_last_updated", s.Res.TimeLastUpdated.String())
	}

	return nil
}

func JavaLicenseAcceptanceRecordSummaryToMap(obj oci_jms_java_downloads.JavaLicenseAcceptanceRecordSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CreatedBy != nil {
		result["created_by"] = []interface{}{PrincipalToMap(obj.CreatedBy)}
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LastUpdatedBy != nil {
		result["last_updated_by"] = []interface{}{PrincipalToMap(obj.LastUpdatedBy)}
	}

	result["license_acceptance_status"] = string(obj.LicenseAcceptanceStatus)

	result["license_type"] = string(obj.LicenseType)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeAccepted != nil {
		result["time_accepted"] = obj.TimeAccepted.String()
	}

	if obj.TimeLastUpdated != nil {
		result["time_last_updated"] = obj.TimeLastUpdated.String()
	}
	return result
}
