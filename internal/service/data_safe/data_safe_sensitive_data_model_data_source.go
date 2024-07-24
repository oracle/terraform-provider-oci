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

func DataSafeSensitiveDataModelDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["sensitive_data_model_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeSensitiveDataModelResource(), fieldMap, readSingularDataSafeSensitiveDataModel)
}

func readSingularDataSafeSensitiveDataModel(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSensitiveDataModelDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetSensitiveDataModelResponse
}

func (s *DataSafeSensitiveDataModelDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSensitiveDataModelDataSourceCrud) Get() error {
	request := oci_data_safe.GetSensitiveDataModelRequest{}

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetSensitiveDataModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSensitiveDataModelDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AppSuiteName != nil {
		s.D.Set("app_suite_name", *s.Res.AppSuiteName)
	}

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

	if s.Res.IsAppDefinedRelationDiscoveryEnabled != nil {
		s.D.Set("is_app_defined_relation_discovery_enabled", *s.Res.IsAppDefinedRelationDiscoveryEnabled)
	}

	if s.Res.IsIncludeAllSchemas != nil {
		s.D.Set("is_include_all_schemas", *s.Res.IsIncludeAllSchemas)
	}

	if s.Res.IsIncludeAllSensitiveTypes != nil {
		s.D.Set("is_include_all_sensitive_types", *s.Res.IsIncludeAllSensitiveTypes)
	}

	if s.Res.IsSampleDataCollectionEnabled != nil {
		s.D.Set("is_sample_data_collection_enabled", *s.Res.IsSampleDataCollectionEnabled)
	}

	s.D.Set("schemas_for_discovery", s.Res.SchemasForDiscovery)

	s.D.Set("sensitive_type_ids_for_discovery", s.Res.SensitiveTypeIdsForDiscovery)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	tablesForDiscovery := []interface{}{}
	for _, item := range s.Res.TablesForDiscovery {
		tablesForDiscovery = append(tablesForDiscovery, TablesForDiscoveryToMap(item))
	}
	s.D.Set("tables_for_discovery", tablesForDiscovery)

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
