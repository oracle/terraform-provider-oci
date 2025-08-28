// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAttributeSetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["attribute_set_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeAttributeSetResource(), fieldMap, readSingularDataSafeAttributeSet)
}

func readSingularDataSafeAttributeSet(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAttributeSetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAttributeSetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetAttributeSetResponse
}

func (s *DataSafeAttributeSetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAttributeSetDataSourceCrud) Get() error {
	request := oci_data_safe.GetAttributeSetRequest{}

	if attributeSetId, ok := s.D.GetOkExists("attribute_set_id"); ok {
		tmp := attributeSetId.(string)
		request.AttributeSetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetAttributeSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeAttributeSetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("attribute_set_type", s.Res.AttributeSetType)

	s.D.Set("attribute_set_values", s.Res.AttributeSetValues)

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

	s.D.Set("in_use", s.Res.InUse)

	if s.Res.IsUserDefined != nil {
		s.D.Set("is_user_defined", *s.Res.IsUserDefined)
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

	return nil
}
