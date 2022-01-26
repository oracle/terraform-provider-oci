// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v56/database"
)

func DatabaseDatabaseSoftwareImageResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("2h"),
			Update: tfresource.GetTimeoutDuration("30m"),
			Delete: tfresource.GetTimeoutDuration("30m"),
		},
		Create: createDatabaseDatabaseSoftwareImage,
		Read:   readDatabaseDatabaseSoftwareImage,
		Update: updateDatabaseDatabaseSoftwareImage,
		Delete: deleteDatabaseDatabaseSoftwareImage,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"database_software_image_one_off_patches": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"database_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
			"image_shape_family": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"image_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ls_inventory": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"patch_set": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"source_db_home_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"database_software_image_included_patches": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"included_patches_summary": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_upgrade_supported": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
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

func createDatabaseDatabaseSoftwareImage(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseSoftwareImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseDatabaseSoftwareImage(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseSoftwareImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseDatabaseSoftwareImage(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseSoftwareImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseDatabaseSoftwareImage(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseSoftwareImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseDatabaseSoftwareImageResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.DatabaseSoftwareImage
	DisableNotFoundRetries bool
}

func (s *DatabaseDatabaseSoftwareImageResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDatabaseSoftwareImageResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.DatabaseSoftwareImageLifecycleStateProvisioning),
	}
}

func (s *DatabaseDatabaseSoftwareImageResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.DatabaseSoftwareImageLifecycleStateAvailable),
	}
}

func (s *DatabaseDatabaseSoftwareImageResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.DatabaseSoftwareImageLifecycleStateDeleting),
		string(oci_database.DatabaseSoftwareImageLifecycleStateTerminating),
	}
}

func (s *DatabaseDatabaseSoftwareImageResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.DatabaseSoftwareImageLifecycleStateDeleted),
		string(oci_database.DatabaseSoftwareImageLifecycleStateTerminated),
	}
}

func (s *DatabaseDatabaseSoftwareImageResourceCrud) Create() error {
	request := oci_database.CreateDatabaseSoftwareImageRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if databaseSoftwareImageOneOffPatches, ok := s.D.GetOkExists("database_software_image_one_off_patches"); ok {
		interfaces := databaseSoftwareImageOneOffPatches.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("database_software_image_one_off_patches") {
			request.DatabaseSoftwareImageOneOffPatches = tmp
		}
	}

	if databaseVersion, ok := s.D.GetOkExists("database_version"); ok {
		tmp := databaseVersion.(string)
		request.DatabaseVersion = &tmp
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if imageShapeFamily, ok := s.D.GetOkExists("image_shape_family"); ok {
		request.ImageShapeFamily = oci_database.CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum(imageShapeFamily.(string))
	}

	if imageType, ok := s.D.GetOkExists("image_type"); ok {
		request.ImageType = oci_database.CreateDatabaseSoftwareImageDetailsImageTypeEnum(imageType.(string))
	}

	if lsInventory, ok := s.D.GetOkExists("ls_inventory"); ok {
		tmp := lsInventory.(string)
		request.LsInventory = &tmp
	}

	if patchSet, ok := s.D.GetOkExists("patch_set"); ok {
		tmp := patchSet.(string)
		request.PatchSet = &tmp
	}

	if sourceDbHomeId, ok := s.D.GetOkExists("source_db_home_id"); ok {
		tmp := sourceDbHomeId.(string)
		request.SourceDbHomeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateDatabaseSoftwareImage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseSoftwareImage

	if waitErr := tfresource.WaitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatabaseDatabaseSoftwareImageResourceCrud) Get() error {
	request := oci_database.GetDatabaseSoftwareImageRequest{}

	tmp := s.D.Id()
	request.DatabaseSoftwareImageId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDatabaseSoftwareImage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseSoftwareImage
	return nil
}

func (s *DatabaseDatabaseSoftwareImageResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateDatabaseSoftwareImageRequest{}

	tmp := s.D.Id()
	request.DatabaseSoftwareImageId = &tmp

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateDatabaseSoftwareImage(context.Background(), request)
	if err != nil {
		return err
	}

	// This Update does not support work-request
	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	s.Res = &response.DatabaseSoftwareImage
	return nil
}

func (s *DatabaseDatabaseSoftwareImageResourceCrud) Delete() error {
	request := oci_database.DeleteDatabaseSoftwareImageRequest{}

	tmp := s.D.Id()
	request.DatabaseSoftwareImageId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteDatabaseSoftwareImage(context.Background(), request)
	return err
}

func (s *DatabaseDatabaseSoftwareImageResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("database_software_image_included_patches", s.Res.DatabaseSoftwareImageIncludedPatches)

	s.D.Set("database_software_image_one_off_patches", s.Res.DatabaseSoftwareImageOneOffPatches)

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

	s.D.Set("image_type", s.Res.ImageType)

	if s.Res.IncludedPatchesSummary != nil {
		s.D.Set("included_patches_summary", *s.Res.IncludedPatchesSummary)
	}

	if s.Res.IsUpgradeSupported != nil {
		s.D.Set("is_upgrade_supported", *s.Res.IsUpgradeSupported)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LsInventory != nil {
		s.D.Set("ls_inventory", *s.Res.LsInventory)
	}

	if s.Res.PatchSet != nil {
		s.D.Set("patch_set", *s.Res.PatchSet)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *DatabaseDatabaseSoftwareImageResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeDatabaseSoftwareImageCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatabaseSoftwareImageId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeDatabaseSoftwareImageCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
