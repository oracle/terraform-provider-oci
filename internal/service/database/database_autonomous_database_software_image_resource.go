// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousDatabaseSoftwareImageResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseAutonomousDatabaseSoftwareImage,
		Read:     readDatabaseAutonomousDatabaseSoftwareImage,
		Update:   updateDatabaseAutonomousDatabaseSoftwareImage,
		Delete:   deleteDatabaseAutonomousDatabaseSoftwareImage,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image_shape_family": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_cdb_id": {
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

			// Computed
			"autonomous_dsi_one_off_patches": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"database_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"release_update": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseAutonomousDatabaseSoftwareImage(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseSoftwareImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseAutonomousDatabaseSoftwareImage(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseSoftwareImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseAutonomousDatabaseSoftwareImage(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseSoftwareImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseAutonomousDatabaseSoftwareImage(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseSoftwareImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseAutonomousDatabaseSoftwareImageResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousDatabaseSoftwareImage
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseAutonomousDatabaseSoftwareImageResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseAutonomousDatabaseSoftwareImageResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseSoftwareImageLifecycleStateProvisioning),
	}
}

func (s *DatabaseAutonomousDatabaseSoftwareImageResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseSoftwareImageLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousDatabaseSoftwareImageResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseSoftwareImageLifecycleStateTerminating),
	}
}

func (s *DatabaseAutonomousDatabaseSoftwareImageResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseSoftwareImageLifecycleStateTerminated),
	}
}

func (s *DatabaseAutonomousDatabaseSoftwareImageResourceCrud) Create() error {
	request := oci_database.CreateAutonomousDatabaseSoftwareImageRequest{}

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

	if imageShapeFamily, ok := s.D.GetOkExists("image_shape_family"); ok {
		request.ImageShapeFamily = oci_database.CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum(imageShapeFamily.(string))
	}

	if sourceCdbId, ok := s.D.GetOkExists("source_cdb_id"); ok {
		tmp := sourceCdbId.(string)
		request.SourceCdbId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateAutonomousDatabaseSoftwareImage(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.AutonomousDatabaseSoftwareImage

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomousdatabasesoftwareimage", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabaseAutonomousDatabaseSoftwareImageResourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseSoftwareImageRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseSoftwareImageId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousDatabaseSoftwareImage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabaseSoftwareImage
	return nil
}

func (s *DatabaseAutonomousDatabaseSoftwareImageResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateAutonomousDatabaseSoftwareImageRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseSoftwareImageId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousDatabaseSoftwareImage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabaseSoftwareImage
	return nil
}

func (s *DatabaseAutonomousDatabaseSoftwareImageResourceCrud) Delete() error {
	request := oci_database.DeleteAutonomousDatabaseSoftwareImageRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseSoftwareImageId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteAutonomousDatabaseSoftwareImage(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomousdatabasesoftwareimage", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseSoftwareImageResourceCrud) SetData() error {
	s.D.Set("autonomous_dsi_one_off_patches", s.Res.AutonomousDsiOneOffPatches)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("image_shape_family", s.Res.ImageShapeFamily)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ReleaseUpdate != nil {
		s.D.Set("release_update", *s.Res.ReleaseUpdate)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func AutonomousDatabaseSoftwareImageSummaryToMap(obj oci_database.AutonomousDatabaseSoftwareImageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["autonomous_dsi_one_off_patches"] = obj.AutonomousDsiOneOffPatches

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DatabaseVersion != nil {
		result["database_version"] = string(*obj.DatabaseVersion)
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

	result["image_shape_family"] = string(obj.ImageShapeFamily)

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ReleaseUpdate != nil {
		result["release_update"] = string(*obj.ReleaseUpdate)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *DatabaseAutonomousDatabaseSoftwareImageResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeAutonomousDatabaseSoftwareImageCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AutonomousDatabaseSoftwareImageId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeAutonomousDatabaseSoftwareImageCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomousdatabasesoftwareimage", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}
