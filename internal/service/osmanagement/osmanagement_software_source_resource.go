// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_osmanagement "github.com/oracle/oci-go-sdk/v65/osmanagement"
)

func OsmanagementSoftwareSourceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsmanagementSoftwareSource,
		Read:     readOsmanagementSoftwareSource,
		Update:   updateOsmanagementSoftwareSource,
		Delete:   deleteOsmanagementSoftwareSource,
		Schema: map[string]*schema.Schema{
			// Required
			"arch_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"checksum_type": {
				Type:     schema.TypeString,
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
			"maintainer_email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maintainer_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maintainer_phone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"associated_managed_instances": {
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
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"gpg_key_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gpg_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gpg_key_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"packages": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"parent_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repo_type": {
				Type:     schema.TypeString,
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
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOsmanagementSoftwareSource(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementSoftwareSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readOsmanagementSoftwareSource(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementSoftwareSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateOsmanagementSoftwareSource(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementSoftwareSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOsmanagementSoftwareSource(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementSoftwareSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OsmanagementSoftwareSourceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_osmanagement.OsManagementClient
	Res                    *oci_osmanagement.SoftwareSource
	DisableNotFoundRetries bool
}

func (s *OsmanagementSoftwareSourceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsmanagementSoftwareSourceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_osmanagement.LifecycleStatesCreating),
	}
}

func (s *OsmanagementSoftwareSourceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_osmanagement.LifecycleStatesActive),
	}
}

func (s *OsmanagementSoftwareSourceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_osmanagement.LifecycleStatesDeleting),
	}
}

func (s *OsmanagementSoftwareSourceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_osmanagement.LifecycleStatesDeleted),
	}
}

func (s *OsmanagementSoftwareSourceResourceCrud) Create() error {
	request := oci_osmanagement.CreateSoftwareSourceRequest{}

	if archType, ok := s.D.GetOkExists("arch_type"); ok {
		request.ArchType = oci_osmanagement.ArchTypesEnum(archType.(string))
	}

	if checksumType, ok := s.D.GetOkExists("checksum_type"); ok {
		request.ChecksumType = oci_osmanagement.ChecksumTypesEnum(checksumType.(string))
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

	if maintainerEmail, ok := s.D.GetOkExists("maintainer_email"); ok {
		tmp := maintainerEmail.(string)
		request.MaintainerEmail = &tmp
	}

	if maintainerName, ok := s.D.GetOkExists("maintainer_name"); ok {
		tmp := maintainerName.(string)
		request.MaintainerName = &tmp
	}

	if maintainerPhone, ok := s.D.GetOkExists("maintainer_phone"); ok {
		tmp := maintainerPhone.(string)
		request.MaintainerPhone = &tmp
	}

	if parentId, ok := s.D.GetOkExists("parent_id"); ok {
		tmp := parentId.(string)
		request.ParentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	response, err := s.Client.CreateSoftwareSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SoftwareSource
	return nil
}

func (s *OsmanagementSoftwareSourceResourceCrud) Get() error {
	request := oci_osmanagement.GetSoftwareSourceRequest{}

	tmp := s.D.Id()
	request.SoftwareSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	response, err := s.Client.GetSoftwareSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SoftwareSource
	return nil
}

func (s *OsmanagementSoftwareSourceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_osmanagement.UpdateSoftwareSourceRequest{}

	if checksumType, ok := s.D.GetOkExists("checksum_type"); ok {
		request.ChecksumType = oci_osmanagement.ChecksumTypesEnum(checksumType.(string))
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

	if maintainerEmail, ok := s.D.GetOkExists("maintainer_email"); ok {
		tmp := maintainerEmail.(string)
		request.MaintainerEmail = &tmp
	}

	if maintainerName, ok := s.D.GetOkExists("maintainer_name"); ok {
		tmp := maintainerName.(string)
		request.MaintainerName = &tmp
	}

	if maintainerPhone, ok := s.D.GetOkExists("maintainer_phone"); ok {
		tmp := maintainerPhone.(string)
		request.MaintainerPhone = &tmp
	}

	tmp := s.D.Id()
	request.SoftwareSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	response, err := s.Client.UpdateSoftwareSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SoftwareSource
	return nil
}

func (s *OsmanagementSoftwareSourceResourceCrud) Delete() error {
	request := oci_osmanagement.DeleteSoftwareSourceRequest{}

	tmp := s.D.Id()
	request.SoftwareSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	_, err := s.Client.DeleteSoftwareSource(context.Background(), request)
	return err
}

func (s *OsmanagementSoftwareSourceResourceCrud) SetData() error {
	s.D.Set("arch_type", s.Res.ArchType)

	associatedManagedInstances := []interface{}{}
	for _, item := range s.Res.AssociatedManagedInstances {
		associatedManagedInstances = append(associatedManagedInstances, IdToMap(item))
	}
	s.D.Set("associated_managed_instances", associatedManagedInstances)

	s.D.Set("checksum_type", s.Res.ChecksumType)

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

	if s.Res.GpgKeyFingerprint != nil {
		s.D.Set("gpg_key_fingerprint", *s.Res.GpgKeyFingerprint)
	}

	if s.Res.GpgKeyId != nil {
		s.D.Set("gpg_key_id", *s.Res.GpgKeyId)
	}

	if s.Res.GpgKeyUrl != nil {
		s.D.Set("gpg_key_url", *s.Res.GpgKeyUrl)
	}

	if s.Res.MaintainerEmail != nil {
		s.D.Set("maintainer_email", *s.Res.MaintainerEmail)
	}

	if s.Res.MaintainerName != nil {
		s.D.Set("maintainer_name", *s.Res.MaintainerName)
	}

	if s.Res.MaintainerPhone != nil {
		s.D.Set("maintainer_phone", *s.Res.MaintainerPhone)
	}

	if s.Res.Packages != nil {
		s.D.Set("packages", *s.Res.Packages)
	}

	if s.Res.ParentId != nil {
		s.D.Set("parent_id", *s.Res.ParentId)
	}

	if s.Res.ParentName != nil {
		s.D.Set("parent_name", *s.Res.ParentName)
	}

	if s.Res.RepoType != nil {
		s.D.Set("repo_type", *s.Res.RepoType)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.Url != nil {
		s.D.Set("url", *s.Res.Url)
	}

	return nil
}

func (s *OsmanagementSoftwareSourceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_osmanagement.ChangeSoftwareSourceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SoftwareSourceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	_, err := s.Client.ChangeSoftwareSourceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
