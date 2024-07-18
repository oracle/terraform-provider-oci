// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeDiscoveryJobDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["discovery_job_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeDiscoveryJobResource(), fieldMap, readSingularDataSafeDiscoveryJob)
}

func readSingularDataSafeDiscoveryJob(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDiscoveryJobDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeDiscoveryJobDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetDiscoveryJobResponse
}

func (s *DataSafeDiscoveryJobDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeDiscoveryJobDataSourceCrud) Get() error {
	request := oci_data_safe.GetDiscoveryJobRequest{}

	if discoveryJobId, ok := s.D.GetOkExists("discovery_job_id"); ok {
		tmp := discoveryJobId.(string)
		request.DiscoveryJobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetDiscoveryJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeDiscoveryJobDataSourceCrud) SetData() error {
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

	s.D.Set("discovery_type", s.Res.DiscoveryType)

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

	if s.Res.SensitiveDataModelId != nil {
		s.D.Set("sensitive_data_model_id", *s.Res.SensitiveDataModelId)
	}

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

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TotalColumnsScanned != nil {
		s.D.Set("total_columns_scanned", strconv.FormatInt(*s.Res.TotalColumnsScanned, 10))
	}

	if s.Res.TotalDeletedSensitiveColumns != nil {
		s.D.Set("total_deleted_sensitive_columns", strconv.FormatInt(*s.Res.TotalDeletedSensitiveColumns, 10))
	}

	if s.Res.TotalModifiedSensitiveColumns != nil {
		s.D.Set("total_modified_sensitive_columns", strconv.FormatInt(*s.Res.TotalModifiedSensitiveColumns, 10))
	}

	if s.Res.TotalNewSensitiveColumns != nil {
		s.D.Set("total_new_sensitive_columns", strconv.FormatInt(*s.Res.TotalNewSensitiveColumns, 10))
	}

	if s.Res.TotalObjectsScanned != nil {
		s.D.Set("total_objects_scanned", strconv.FormatInt(*s.Res.TotalObjectsScanned, 10))
	}

	if s.Res.TotalSchemasScanned != nil {
		s.D.Set("total_schemas_scanned", strconv.FormatInt(*s.Res.TotalSchemasScanned, 10))
	}

	return nil
}
