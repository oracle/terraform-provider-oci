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

func DatabaseOneoffPatchResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseOneoffPatch,
		Read:     readDatabaseOneoffPatch,
		Update:   updateDatabaseOneoffPatch,
		Delete:   deleteDatabaseOneoffPatch,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"release_update": {
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
			"one_off_patches": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"download_oneoff_patch_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sha256sum": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_kbs": {
				Type:     schema.TypeFloat,
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
			"time_of_expiration": {
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

func createDatabaseOneoffPatch(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseOneoffPatchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("download_oneoff_patch_trigger"); ok {
		err := sync.DownloadOneoffPatch()
		if err != nil {
			return err
		}
	}
	return nil

}

func readDatabaseOneoffPatch(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseOneoffPatchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseOneoffPatch(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseOneoffPatchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	if _, ok := sync.D.GetOkExists("download_oneoff_patch_trigger"); ok && sync.D.HasChange("download_oneoff_patch_trigger") {
		oldRaw, newRaw := sync.D.GetChange("download_oneoff_patch_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.DownloadOneoffPatch()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("download_oneoff_patch_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteDatabaseOneoffPatch(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseOneoffPatchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseOneoffPatchResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.OneoffPatch
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseOneoffPatchResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseOneoffPatchResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.OneoffPatchLifecycleStateCreating),
	}
}

func (s *DatabaseOneoffPatchResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.OneoffPatchLifecycleStateAvailable),
	}
}

func (s *DatabaseOneoffPatchResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.OneoffPatchLifecycleStateDeleting),
		string(oci_database.OneoffPatchLifecycleStateTerminating),
	}
}

func (s *DatabaseOneoffPatchResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.OneoffPatchLifecycleStateDeleted),
		string(oci_database.OneoffPatchLifecycleStateTerminated),
	}
}

func (s *DatabaseOneoffPatchResourceCrud) Create() error {
	request := oci_database.CreateOneoffPatchRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
		tmp := dbVersion.(string)
		request.DbVersion = &tmp
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

	if oneOffPatches, ok := s.D.GetOkExists("one_off_patches"); ok {
		interfaces := oneOffPatches.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("one_off_patches") {
			request.OneOffPatches = tmp
		}
	}

	if releaseUpdate, ok := s.D.GetOkExists("release_update"); ok {
		tmp := releaseUpdate.(string)
		request.ReleaseUpdate = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateOneoffPatch(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.OneoffPatch

	if workId != nil {
		identifier, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "oneoffPatch", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabaseOneoffPatchResourceCrud) Get() error {
	request := oci_database.GetOneoffPatchRequest{}

	tmp := s.D.Id()
	request.OneoffPatchId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetOneoffPatch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OneoffPatch
	return nil
}

func (s *DatabaseOneoffPatchResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists("compartmentId"); ok && s.D.HasChange("compartmentId") {
		err := s.ChangeOneoffPatchCompartment()
		if err != nil {
			return err
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateOneoffPatchRequest{}

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

	tmp := s.D.Id()
	request.OneoffPatchId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateOneoffPatch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OneoffPatch
	return nil
}

func (s *DatabaseOneoffPatchResourceCrud) Delete() error {
	request := oci_database.DeleteOneoffPatchRequest{}

	tmp := s.D.Id()
	request.OneoffPatchId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteOneoffPatch(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "oneoffPatch", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseOneoffPatchResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("one_off_patches", s.Res.OneOffPatches)
	s.D.Set("one_off_patches", s.Res.OneOffPatches)

	if s.Res.ReleaseUpdate != nil {
		s.D.Set("release_update", *s.Res.ReleaseUpdate)
	}

	if s.Res.Sha256Sum != nil {
		s.D.Set("sha256sum", *s.Res.Sha256Sum)
	}

	if s.Res.SizeInKBs != nil {
		s.D.Set("size_in_kbs", *s.Res.SizeInKBs)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfExpiration != nil {
		s.D.Set("time_of_expiration", s.Res.TimeOfExpiration.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DatabaseOneoffPatchResourceCrud) DownloadOneoffPatch() error {
	request := oci_database.DownloadOneoffPatchRequest{}

	idTmp := s.D.Id()
	request.OneoffPatchId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DownloadOneoffPatch(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("download_oneoff_patch_trigger")
	s.D.Set("download_oneoff_patch_trigger", val)

	return s.Get()
}

func (s *DatabaseOneoffPatchResourceCrud) ChangeOneoffPatchCompartment() error {
	request := oci_database.ChangeOneoffPatchCompartmentRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	idTmp := s.D.Id()
	request.OneoffPatchId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeOneoffPatchCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatabaseOneoffPatchResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeOneoffPatchCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OneoffPatchId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeOneoffPatchCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "oneoffPatch", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}
