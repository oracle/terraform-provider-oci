// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceModelGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["model_group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatascienceModelGroupResource(), fieldMap, readSingularDatascienceModelGroup)
}

func readSingularDatascienceModelGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceModelGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetModelGroupResponse
}

func (s *DatascienceModelGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelGroupDataSourceCrud) Get() error {
	request := oci_datascience.GetModelGroupRequest{}

	if modelGroupId, ok := s.D.GetOkExists("model_group_id"); ok {
		tmp := modelGroupId.(string)
		request.ModelGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.GetModelGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatascienceModelGroupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreateType != nil {
		s.D.Set("create_type", *s.Res.CreateType)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MemberModelEntries != nil {
		s.D.Set("member_model_entries", []interface{}{MemberModelEntriesToMap(s.Res.MemberModelEntries)})
	} else {
		s.D.Set("member_model_entries", nil)
	}

	if s.Res.ModelGroupDetails != nil {
		modelGroupDetailsArray := []interface{}{}
		if modelGroupDetailsMap := ModelGroupDetailsToMap(&s.Res.ModelGroupDetails); modelGroupDetailsMap != nil {
			modelGroupDetailsArray = append(modelGroupDetailsArray, modelGroupDetailsMap)
		}
		s.D.Set("model_group_details", modelGroupDetailsArray)
	} else {
		s.D.Set("model_group_details", nil)
	}

	if s.Res.ModelGroupVersionHistoryId != nil {
		s.D.Set("model_group_version_history_id", *s.Res.ModelGroupVersionHistoryId)
	}

	if s.Res.ModelGroupVersionHistoryName != nil {
		s.D.Set("model_group_version_history_name", *s.Res.ModelGroupVersionHistoryName)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	if s.Res.SourceModelGroupId != nil {
		s.D.Set("source_model_group_id", *s.Res.SourceModelGroupId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VersionId != nil {
		s.D.Set("version_id", strconv.FormatInt(*s.Res.VersionId, 10))
	}

	if s.Res.VersionLabel != nil {
		s.D.Set("version_label", *s.Res.VersionLabel)
	}

	return nil
}
