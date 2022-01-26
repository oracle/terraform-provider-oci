// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v56/datasafe"
)

func DataSafeTargetDatabaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["target_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeTargetDatabaseResource(), fieldMap, readSingularDataSafeTargetDatabase)
}

func readSingularDataSafeTargetDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeTargetDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetTargetDatabaseResponse
}

func (s *DataSafeTargetDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeTargetDatabaseDataSourceCrud) Get() error {
	request := oci_data_safe.GetTargetDatabaseRequest{}

	if targetDatabaseId, ok := s.D.GetOkExists("target_database_id"); ok {
		tmp := targetDatabaseId.(string)
		request.TargetDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetTargetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeTargetDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionOption != nil {
		connectionOptionArray := []interface{}{}
		if connectionOptionMap := ConnectionOptionToMap(&s.Res.ConnectionOption); connectionOptionMap != nil {
			connectionOptionArray = append(connectionOptionArray, connectionOptionMap)
		}
		s.D.Set("connection_option", connectionOptionArray)
	} else {
		s.D.Set("connection_option", nil)
	}

	if s.Res.Credentials != nil {
		s.D.Set("credentials", []interface{}{CredentialsToMap(s.Res.Credentials)})
	} else {
		s.D.Set("credentials", nil)
	}

	if s.Res.DatabaseDetails != nil {
		databaseDetailsArray := []interface{}{}
		if databaseDetailsMap := DatabaseDetailsToMap(&s.Res.DatabaseDetails); databaseDetailsMap != nil {
			databaseDetailsArray = append(databaseDetailsArray, databaseDetailsMap)
		}
		s.D.Set("database_details", databaseDetailsArray)
	} else {
		s.D.Set("database_details", nil)
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

	if s.Res.TlsConfig != nil {
		s.D.Set("tls_config", []interface{}{TlsConfigToMap(s.Res.TlsConfig)})
	} else {
		s.D.Set("tls_config", nil)
	}

	return nil
}
