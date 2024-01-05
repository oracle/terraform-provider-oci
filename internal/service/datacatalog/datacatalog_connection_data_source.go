// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacatalog

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v65/datacatalog"
)

func DatacatalogConnectionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["catalog_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["connection_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["data_asset_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["fields"] = &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Set:      tfresource.LiteralTypeHashCodeForSets,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	return tfresource.GetSingularDataSourceItemSchema(DatacatalogConnectionResource(), fieldMap, readSingularDatacatalogConnection)
}

func readSingularDatacatalogConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogConnectionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.ReadResource(sync)
}

type DatacatalogConnectionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacatalog.DataCatalogClient
	Res    *oci_datacatalog.GetConnectionResponse
}

func (s *DatacatalogConnectionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatacatalogConnectionDataSourceCrud) Get() error {
	request := oci_datacatalog.GetConnectionRequest{}

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	if connectionKey, ok := s.D.GetOkExists("connection_key"); ok {
		tmp := connectionKey.(string)
		request.ConnectionKey = &tmp
	}

	if dataAssetKey, ok := s.D.GetOkExists("data_asset_key"); ok {
		tmp := dataAssetKey.(string)
		request.DataAssetKey = &tmp
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.([]interface{})
		tmp := make([]oci_datacatalog.GetConnectionFieldsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_datacatalog.GetConnectionFieldsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datacatalog")

	response, err := s.Client.GetConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatacatalogConnectionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatacatalogConnectionDataSource-", DatacatalogConnectionDataSource(), s.D))

	if s.Res.CreatedById != nil {
		s.D.Set("created_by_id", *s.Res.CreatedById)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalKey != nil {
		s.D.Set("external_key", *s.Res.ExternalKey)
	}

	if s.Res.IsDefault != nil {
		s.D.Set("is_default", *s.Res.IsDefault)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Properties != nil {
		s.D.Set("properties", propertiesToMap(s.Res.Properties))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeStatusUpdated != nil {
		s.D.Set("time_status_updated", s.Res.TimeStatusUpdated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TypeKey != nil {
		s.D.Set("type_key", *s.Res.TypeKey)
	}

	if s.Res.UpdatedById != nil {
		s.D.Set("updated_by_id", *s.Res.UpdatedById)
	}

	if s.Res.Uri != nil {
		s.D.Set("uri", *s.Res.Uri)
	}

	return nil
}
