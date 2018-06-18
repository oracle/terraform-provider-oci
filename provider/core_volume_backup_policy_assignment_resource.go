// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func VolumeBackupPolicyAssignmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createVolumeBackupPolicyAssignment,
		Read:     readVolumeBackupPolicyAssignment,
		Delete:   deleteVolumeBackupPolicyAssignment,
		Schema: map[string]*schema.Schema{
			// Required
			"asset_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"id": {
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

func createVolumeBackupPolicyAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeBackupPolicyAssignmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return crud.CreateResource(d, sync)
}

func readVolumeBackupPolicyAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeBackupPolicyAssignmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return crud.ReadResource(sync)
}

func deleteVolumeBackupPolicyAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeBackupPolicyAssignmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type VolumeBackupPolicyAssignmentResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.BlockstorageClient
	Res                    *oci_core.VolumeBackupPolicyAssignment
	DisableNotFoundRetries bool
}

func (s *VolumeBackupPolicyAssignmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VolumeBackupPolicyAssignmentResourceCrud) Create() error {
	request := oci_core.CreateVolumeBackupPolicyAssignmentRequest{}

	if assetId, ok := s.D.GetOkExists("asset_id"); ok {
		tmp := assetId.(string)
		request.AssetId = &tmp
	}

	if policyId, ok := s.D.GetOkExists("policy_id"); ok {
		tmp := policyId.(string)
		request.PolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVolumeBackupPolicyAssignment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeBackupPolicyAssignment
	return nil
}

func (s *VolumeBackupPolicyAssignmentResourceCrud) Get() error {
	request := oci_core.GetVolumeBackupPolicyAssignmentRequest{}

	tmp := s.D.Id()
	request.PolicyAssignmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVolumeBackupPolicyAssignment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeBackupPolicyAssignment
	return nil
}

func (s *VolumeBackupPolicyAssignmentResourceCrud) Delete() error {
	request := oci_core.DeleteVolumeBackupPolicyAssignmentRequest{}

	tmp := s.D.Id()
	request.PolicyAssignmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVolumeBackupPolicyAssignment(context.Background(), request)
	return err
}

func (s *VolumeBackupPolicyAssignmentResourceCrud) SetData() {
	if s.Res.AssetId != nil {
		s.D.Set("asset_id", *s.Res.AssetId)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.PolicyId != nil {
		s.D.Set("policy_id", *s.Res.PolicyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
