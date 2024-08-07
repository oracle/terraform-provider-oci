// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DevopsRepositoryProtectedBranchManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDevopsRepositoryProtectedBranchManagement,
		Read:     readDevopsRepositoryProtectedBranchManagement,
		Delete:   deleteDevopsRepositoryProtectedBranchManagement,
		Update:   createDevopsRepositoryProtectedBranchManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"branch_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"protection_levels": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			//"defined_tags": {
			//	Type:     schema.TypeMap,
			//	Computed: true,
			//	Elem:     schema.TypeString,
			//},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
		},
	}
}

func createDevopsRepositoryProtectedBranchManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryProtectedBranchManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.CreateResource(d, sync)
}

func readDevopsRepositoryProtectedBranchManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDevopsRepositoryProtectedBranchManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryProtectedBranchManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.DeleteResource(d, sync)
}

type DevopsRepositoryProtectedBranchManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.ProtectedBranch
	DisableNotFoundRetries bool
}

func (s *DevopsRepositoryProtectedBranchManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DevopsRepositoryProtectedBranchManagementResource-", DevopsRepositoryProtectedBranchManagementResource(), s.D)
}

func (s *DevopsRepositoryProtectedBranchManagementResourceCrud) Create() error {
	request := oci_devops.CreateOrUpdateProtectedBranchRequest{}

	if branchName, ok := s.D.GetOkExists("branch_name"); ok {
		tmp := branchName.(string)
		request.BranchName = &tmp
	}

	if protectionLevels, ok := s.D.GetOkExists("protection_levels"); ok {
		interfaces := protectionLevels.([]interface{})
		tmp := make([]oci_devops.ProtectionLevelEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_devops.ProtectionLevelEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("protection_levels") {
			request.ProtectionLevels = tmp
		}
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.CreateOrUpdateProtectedBranch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProtectedBranch
	return nil
}

func (s *DevopsRepositoryProtectedBranchManagementResourceCrud) SetData() error {
	if s.Res.BranchName != nil {
		s.D.Set("branch_name", *s.Res.BranchName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("protection_levels", s.Res.ProtectionLevels)

	return nil
}

func (s *DevopsRepositoryProtectedBranchManagementResourceCrud) Delete() error {
	request := oci_devops.DeleteProtectedBranchRequest{}

	if branchName, ok := s.D.GetOkExists("branch_name"); ok {
		tmp := branchName.(string)
		request.BranchName = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	_, err := s.Client.DeleteProtectedBranch(context.Background(), request)
	return err
}
