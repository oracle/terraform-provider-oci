// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreVolumeBackupPolicyAssignmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreVolumeBackupPolicyAssignment,
		Read:     readCoreVolumeBackupPolicyAssignment,
		Delete:   deleteCoreVolumeBackupPolicyAssignment,
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
			"xrc_kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreVolumeBackupPolicyAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupPolicyAssignmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreVolumeBackupPolicyAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupPolicyAssignmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

func deleteCoreVolumeBackupPolicyAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupPolicyAssignmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreVolumeBackupPolicyAssignmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.BlockstorageClient
	Res                    *oci_core.VolumeBackupPolicyAssignment
	DisableNotFoundRetries bool
}

func (s *CoreVolumeBackupPolicyAssignmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreVolumeBackupPolicyAssignmentResourceCrud) Create() error {
	request := oci_core.CreateVolumeBackupPolicyAssignmentRequest{}

	if assetId, ok := s.D.GetOkExists("asset_id"); ok {
		tmp := assetId.(string)
		request.AssetId = &tmp
	}

	if policyId, ok := s.D.GetOkExists("policy_id"); ok {
		tmp := policyId.(string)
		request.PolicyId = &tmp
	}

	if xrcKmsKeyId, ok := s.D.GetOkExists("xrc_kms_key_id"); ok {
		tmp := xrcKmsKeyId.(string)
		request.XrcKmsKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVolumeBackupPolicyAssignment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeBackupPolicyAssignment
	return nil
}

func (s *CoreVolumeBackupPolicyAssignmentResourceCrud) Get() error {
	request := oci_core.GetVolumeBackupPolicyAssignmentRequest{}

	tmp := s.D.Id()
	request.PolicyAssignmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVolumeBackupPolicyAssignment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeBackupPolicyAssignment
	return nil
}

func (s *CoreVolumeBackupPolicyAssignmentResourceCrud) Delete() error {
	request := oci_core.DeleteVolumeBackupPolicyAssignmentRequest{}

	tmp := s.D.Id()
	request.PolicyAssignmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVolumeBackupPolicyAssignment(context.Background(), request)
	return err
}

func (s *CoreVolumeBackupPolicyAssignmentResourceCrud) SetData() error {
	if s.Res.AssetId != nil {
		s.D.Set("asset_id", *s.Res.AssetId)
	}

	if s.Res.PolicyId != nil {
		s.D.Set("policy_id", *s.Res.PolicyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.XrcKmsKeyId != nil {
		s.D.Set("xrc_kms_key_id", *s.Res.XrcKmsKeyId)
	}

	return nil
}
