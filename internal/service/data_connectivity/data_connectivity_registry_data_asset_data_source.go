// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_connectivity

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v60/dataconnectivity"
)

func DataConnectivityRegistryDataAssetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["data_asset_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["registry_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataConnectivityRegistryDataAssetResource(), fieldMap, readSingularDataConnectivityRegistryDataAsset)
}

func readSingularDataConnectivityRegistryDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryDataAssetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.ReadResource(sync)
}

type DataConnectivityRegistryDataAssetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_connectivity.DataConnectivityManagementClient
	Res    *oci_data_connectivity.GetDataAssetResponse
}

func (s *DataConnectivityRegistryDataAssetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataConnectivityRegistryDataAssetDataSourceCrud) Get() error {
	request := oci_data_connectivity.GetDataAssetRequest{}

	if dataAssetKey, ok := s.D.GetOkExists("data_asset_key"); ok {
		tmp := dataAssetKey.(string)
		request.DataAssetKey = &tmp
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_connectivity")

	response, err := s.Client.GetDataAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataConnectivityRegistryDataAssetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataConnectivityRegistryDataAssetDataSource-", DataConnectivityRegistryDataAssetDataSource(), s.D))

	s.D.Set("asset_properties", s.Res.AssetProperties)

	if s.Res.DefaultConnection != nil {
		s.D.Set("default_connection", []interface{}{DataConnectivityConnectionToMap(s.Res.DefaultConnection)})
	} else {
		s.D.Set("default_connection", nil)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.ExternalKey != nil {
		s.D.Set("external_key", *s.Res.ExternalKey)
	}

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
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

	if s.Res.NativeTypeSystem != nil {
		s.D.Set("native_type_system", []interface{}{TypeSystemToMap(s.Res.NativeTypeSystem)})
	} else {
		s.D.Set("native_type_system", nil)
	}

	if s.Res.ObjectStatus != nil {
		s.D.Set("object_status", *s.Res.ObjectStatus)
	}

	if s.Res.ObjectVersion != nil {
		s.D.Set("object_version", *s.Res.ObjectVersion)
	}

	s.D.Set("properties", tfresource.GenericMapToJsonMap(s.Res.Properties))

	if s.Res.RegistryMetadata != nil {
		s.D.Set("registry_metadata", []interface{}{DataConnectivityRegistryMetadataToMap(s.Res.RegistryMetadata)})
	} else {
		s.D.Set("registry_metadata", nil)
	}

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	return nil
}
