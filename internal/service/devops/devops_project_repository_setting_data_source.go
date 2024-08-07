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

func DevopsProjectRepositorySettingDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["project_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DevopsProjectRepositorySettingResource(), fieldMap, readSingularDevopsProjectRepositorySetting)
}

func readSingularDevopsProjectRepositorySetting(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsProjectRepositorySettingDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsProjectRepositorySettingDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetProjectRepositorySettingsResponse
}

func (s *DevopsProjectRepositorySettingDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsProjectRepositorySettingDataSourceCrud) Get() error {
	request := oci_devops.GetProjectRepositorySettingsRequest{}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetProjectRepositorySettings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsProjectRepositorySettingDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsProjectRepositorySettingDataSource-", DevopsProjectRepositorySettingDataSource(), s.D))

	if s.Res.ApprovalRules != nil {
		s.D.Set("approval_rules", []interface{}{ApprovalRuleCollectionToMap(s.Res.ApprovalRules)})
	} else {
		s.D.Set("approval_rules", nil)
	}

	if s.Res.MergeSettings != nil {
		s.D.Set("merge_settings", []interface{}{MergeSettingsToMap(s.Res.MergeSettings)})
	} else {
		s.D.Set("merge_settings", nil)
	}

	return nil
}
