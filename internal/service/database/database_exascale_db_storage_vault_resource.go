// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseExascaleDbStorageVaultResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseExascaleDbStorageVault,
		Read:     readDatabaseExascaleDbStorageVault,
		Update:   updateDatabaseExascaleDbStorageVault,
		Delete:   deleteDatabaseExascaleDbStorageVault,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"high_capacity_database_storage": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"total_size_in_gbs": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Computed
						"available_size_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},

			// Optional
			"additional_flash_cache_in_percent": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"lifecycle_details": {
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
			"vm_cluster_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vm_cluster_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createDatabaseExascaleDbStorageVault(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExascaleDbStorageVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseExascaleDbStorageVault(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExascaleDbStorageVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseExascaleDbStorageVault(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExascaleDbStorageVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseExascaleDbStorageVault(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExascaleDbStorageVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseExascaleDbStorageVaultResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ExascaleDbStorageVault
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ExascaleDbStorageVaultLifecycleStateProvisioning),
	}
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ExascaleDbStorageVaultLifecycleStateAvailable),
	}
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ExascaleDbStorageVaultLifecycleStateTerminating),
	}
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExascaleDbStorageVaultLifecycleStateTerminated),
	}
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) UpdatePending() []string {
	return []string{
		string(oci_database.ExascaleDbStorageVaultLifecycleStateUpdating),
	}
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) UpdateTarget() []string {
	return []string{
		string(oci_database.ExascaleDbStorageVaultLifecycleStateAvailable),
	}
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) Create() error {
	request := oci_database.CreateExascaleDbStorageVaultRequest{}

	if additionalFlashCacheInPercent, ok := s.D.GetOkExists("additional_flash_cache_in_percent"); ok {
		tmp := additionalFlashCacheInPercent.(int)
		request.AdditionalFlashCacheInPercent = &tmp
	}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if highCapacityDatabaseStorage, ok := s.D.GetOkExists("high_capacity_database_storage"); ok {
		if tmpList := highCapacityDatabaseStorage.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "high_capacity_database_storage", 0)
			tmp, err := s.mapToExascaleDbStorageInputDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.HighCapacityDatabaseStorage = &tmp
		}
	}

	if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
		tmp := timeZone.(string)
		request.TimeZone = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateExascaleDbStorageVault(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.ExascaleDbStorageVault

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "exascaledbstoragevault", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) Get() error {
	request := oci_database.GetExascaleDbStorageVaultRequest{}

	tmp := s.D.Id()
	request.ExascaleDbStorageVaultId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetExascaleDbStorageVault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExascaleDbStorageVault
	return nil
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateExascaleDbStorageVaultRequest{}

	if additionalFlashCacheInPercent, ok := s.D.GetOkExists("additional_flash_cache_in_percent"); ok && s.D.HasChange("additional_flash_cache_in_percent") {
		tmp := additionalFlashCacheInPercent.(int)
		request.AdditionalFlashCacheInPercent = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok && s.D.HasChange("description") {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok && s.D.HasChange("display_name") {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.ExascaleDbStorageVaultId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if highCapacityDatabaseStorage, ok := s.D.GetOkExists("high_capacity_database_storage"); ok && s.D.HasChange("high_capacity_database_storage") {
		if tmpList := highCapacityDatabaseStorage.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "high_capacity_database_storage", 0)
			tmp, err := s.mapToExascaleDbStorageInputDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.HighCapacityDatabaseStorage = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateExascaleDbStorageVault(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "exascaledbstoragevault", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) Delete() error {
	request := oci_database.DeleteExascaleDbStorageVaultRequest{}

	tmp := s.D.Id()
	request.ExascaleDbStorageVaultId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteExascaleDbStorageVault(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "exascaledbstoragevault", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) SetData() error {
	if s.Res.AdditionalFlashCacheInPercent != nil {
		s.D.Set("additional_flash_cache_in_percent", *s.Res.AdditionalFlashCacheInPercent)
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HighCapacityDatabaseStorage != nil {
		s.D.Set("high_capacity_database_storage", []interface{}{ExascaleDbStorageDetailsToMap(s.Res.HighCapacityDatabaseStorage)})
	} else {
		s.D.Set("high_capacity_database_storage", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	if s.Res.VmClusterCount != nil {
		s.D.Set("vm_cluster_count", *s.Res.VmClusterCount)
	}

	s.D.Set("vm_cluster_ids", s.Res.VmClusterIds)

	return nil
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) mapToExascaleDbStorageInputDetails(fieldKeyFormat string) (oci_database.ExascaleDbStorageInputDetails, error) {
	result := oci_database.ExascaleDbStorageInputDetails{}

	if totalSizeInGbs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "total_size_in_gbs")); ok {
		tmp := totalSizeInGbs.(int)
		result.TotalSizeInGbs = &tmp
	}

	return result, nil
}

func ExascaleDbStorageDetailsToMap(obj *oci_database.ExascaleDbStorageDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailableSizeInGbs != nil {
		result["available_size_in_gbs"] = int(*obj.AvailableSizeInGbs)
	}

	if obj.TotalSizeInGbs != nil {
		result["total_size_in_gbs"] = int(*obj.TotalSizeInGbs)
	}

	return result
}

func (s *DatabaseExascaleDbStorageVaultResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeExascaleDbStorageVaultCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ExascaleDbStorageVaultId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeExascaleDbStorageVaultCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "exascaledbstoragevault", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}
