// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSensitiveTypeDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["sensitive_type_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeSensitiveTypeResource(), fieldMap, readSingularDataSafeSensitiveType)
}

func readSingularDataSafeSensitiveType(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveTypeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSensitiveTypeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetSensitiveTypeResponse
}

func (s *DataSafeSensitiveTypeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSensitiveTypeDataSourceCrud) Get() error {
	request := oci_data_safe.GetSensitiveTypeRequest{}

	if sensitiveTypeId, ok := s.D.GetOkExists("sensitive_type_id"); ok {
		tmp := sensitiveTypeId.(string)
		request.SensitiveTypeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetSensitiveType(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSensitiveTypeDataSourceCrud) SetData() error {

	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.SensitiveType.GetId())

	if s.Res.SensitiveType.GetCompartmentId() != nil {
		s.D.Set("compartment_id", *s.Res.SensitiveType.GetCompartmentId())
	}

	if s.Res.SensitiveType.GetDefinedTags() != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.SensitiveType.GetDefinedTags()))
	}

	if s.Res.SensitiveType.GetDescription() != nil {
		s.D.Set("description", *s.Res.SensitiveType.GetDescription())
	}

	if s.Res.SensitiveType.GetDisplayName() != nil {
		s.D.Set("display_name", *s.Res.SensitiveType.GetDisplayName())
	}

	s.D.Set("freeform_tags", s.Res.SensitiveType.GetFreeformTags())

	if s.Res.SensitiveType.GetIsCommon() != nil {
		s.D.Set("is_common", *s.Res.SensitiveType.GetIsCommon())
	}

	if s.Res.SensitiveType.GetIsCommon() != nil {
		s.D.Set("is_common", *s.Res.SensitiveType.GetIsCommon())
	}

	if s.Res.SensitiveType.GetParentCategoryId() != nil {
		s.D.Set("parent_category_id", *s.Res.SensitiveType.GetParentCategoryId())
	}

	if s.Res.SensitiveType.GetShortName() != nil {
		s.D.Set("short_name", *s.Res.SensitiveType.GetShortName())
	}

	s.D.Set("source", s.Res.SensitiveType.GetSource())

	s.D.Set("state", s.Res.SensitiveType.GetLifecycleState())

	if s.Res.SensitiveType.GetSystemTags() != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SensitiveType.GetSystemTags()))
	}

	if s.Res.SensitiveType.GetTimeCreated() != nil {
		s.D.Set("time_created", s.Res.SensitiveType.GetTimeCreated().String())
	}

	if s.Res.SensitiveType.GetTimeUpdated() != nil {
		s.D.Set("time_updated", s.Res.SensitiveType.GetTimeUpdated().String())
	}

	switch v := (s.Res.SensitiveType).(type) {

	case oci_data_safe.SensitiveTypePattern:
		s.D.Set("entity_type", "SENSITIVE_TYPE")

		if v.CommentPattern != nil {
			s.D.Set("comment_pattern", *v.CommentPattern)
		}
	}

	return nil
}
