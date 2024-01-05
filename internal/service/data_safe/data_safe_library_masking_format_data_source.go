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

func DataSafeLibraryMaskingFormatDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["library_masking_format_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeLibraryMaskingFormatResource(), fieldMap, readSingularDataSafeLibraryMaskingFormat)
}

func readSingularDataSafeLibraryMaskingFormat(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeLibraryMaskingFormatDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeLibraryMaskingFormatDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetLibraryMaskingFormatResponse
}

func (s *DataSafeLibraryMaskingFormatDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeLibraryMaskingFormatDataSourceCrud) Get() error {
	request := oci_data_safe.GetLibraryMaskingFormatRequest{}

	if libraryMaskingFormatId, ok := s.D.GetOkExists("library_masking_format_id"); ok {
		tmp := libraryMaskingFormatId.(string)
		request.LibraryMaskingFormatId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetLibraryMaskingFormat(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeLibraryMaskingFormatDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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

	formatEntries := []interface{}{}
	for _, item := range s.Res.FormatEntries {
		formatEntries = append(formatEntries, FormatEntryToMap(item))
	}
	s.D.Set("format_entries", formatEntries)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("sensitive_type_ids", s.Res.SensitiveTypeIds)

	s.D.Set("source", s.Res.Source)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
