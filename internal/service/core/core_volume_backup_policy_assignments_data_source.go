// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreVolumeBackupPolicyAssignmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVolumeBackupPolicyAssignments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"asset_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"volume_backup_policy_assignments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreVolumeBackupPolicyAssignmentResource()),
			},
		},
	}
}

func readCoreVolumeBackupPolicyAssignments(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupPolicyAssignmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

type CoreVolumeBackupPolicyAssignmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.GetVolumeBackupPolicyAssetAssignmentResponse
}

func (s *CoreVolumeBackupPolicyAssignmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVolumeBackupPolicyAssignmentsDataSourceCrud) Get() error {
	request := oci_core.GetVolumeBackupPolicyAssetAssignmentRequest{}

	if assetId, ok := s.D.GetOkExists("asset_id"); ok {
		tmp := assetId.(string)
		request.AssetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetVolumeBackupPolicyAssetAssignment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.GetVolumeBackupPolicyAssetAssignment(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreVolumeBackupPolicyAssignmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreVolumeBackupPolicyAssignmentsDataSource-", CoreVolumeBackupPolicyAssignmentsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		volumeBackupPolicyAssignment := map[string]interface{}{
			"asset_id": *r.AssetId,
		}

		if r.Id != nil {
			volumeBackupPolicyAssignment["id"] = *r.Id
		}

		if r.PolicyId != nil {
			volumeBackupPolicyAssignment["policy_id"] = *r.PolicyId
		}

		if r.TimeCreated != nil {
			volumeBackupPolicyAssignment["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, volumeBackupPolicyAssignment)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreVolumeBackupPolicyAssignmentsDataSource().Schema["volume_backup_policy_assignments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("volume_backup_policy_assignments", resources); err != nil {
		return err
	}

	return nil
}
