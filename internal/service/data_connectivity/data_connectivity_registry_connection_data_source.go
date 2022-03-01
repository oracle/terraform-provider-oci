// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_connectivity

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v60/dataconnectivity"
)

func DataConnectivityRegistryConnectionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["connection_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["registry_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataConnectivityRegistryConnectionResource(), fieldMap, readSingularDataConnectivityRegistryConnection)
}

func readSingularDataConnectivityRegistryConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryConnectionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.ReadResource(sync)
}

type DataConnectivityRegistryConnectionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_connectivity.DataConnectivityManagementClient
	Res    *oci_data_connectivity.GetConnectionResponse
}

func (s *DataConnectivityRegistryConnectionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataConnectivityRegistryConnectionDataSourceCrud) Get() error {
	request := oci_data_connectivity.GetConnectionRequest{}

	if connectionKey, ok := s.D.GetOkExists("connection_key"); ok {
		tmp := connectionKey.(string)
		request.ConnectionKey = &tmp
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_connectivity")

	response, err := s.Client.GetConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataConnectivityRegistryConnectionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataConnectivityRegistryConnectionDataSource-", DataConnectivityRegistryConnectionDataSource(), s.D))

	connectionProperties := []interface{}{}
	for _, item := range s.Res.ConnectionProperties {
		connectionProperties = append(connectionProperties, ConnectionPropertyToMap(item))
	}
	s.D.Set("connection_properties", connectionProperties)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
	}

	if s.Res.IsDefault != nil {
		s.D.Set("is_default", *s.Res.IsDefault)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{DataConnectivityObjectMetadataToMap(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.ModelType != nil {
		s.D.Set("model_type", *s.Res.ModelType)
	}

	if s.Res.ModelVersion != nil {
		s.D.Set("model_version", *s.Res.ModelVersion)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectStatus != nil {
		s.D.Set("object_status", *s.Res.ObjectStatus)
	}

	if s.Res.ObjectVersion != nil {
		s.D.Set("object_version", *s.Res.ObjectVersion)
	}

	if s.Res.PrimarySchema != nil {
		s.D.Set("primary_schema", []interface{}{DataConnectivitySchemaToMap(s.Res.PrimarySchema)})
	} else {
		s.D.Set("primary_schema", nil)
	}

	s.D.Set("properties", tfresource.GenericMapToJsonMap(s.Res.Properties))

	if s.Res.RegistryMetadata != nil {
		s.D.Set("registry_metadata", []interface{}{DataConnectivityDataConnectivityRegistryMetadataToMap(s.Res.RegistryMetadata)})
	} else {
		s.D.Set("registry_metadata", nil)
	}

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	return nil
}
