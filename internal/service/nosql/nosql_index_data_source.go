// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package nosql

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_nosql "github.com/oracle/oci-go-sdk/v65/nosql"
)

func NosqlIndexDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["index_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["table_name_or_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NosqlIndexResource(), fieldMap, readSingularNosqlIndex)
}

func readSingularNosqlIndex(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlIndexDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NosqlClient()

	return tfresource.ReadResource(sync)
}

type NosqlIndexDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_nosql.NosqlClient
	Res    *oci_nosql.GetIndexResponse
}

func (s *NosqlIndexDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NosqlIndexDataSourceCrud) Get() error {
	request := oci_nosql.GetIndexRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if indexName, ok := s.D.GetOkExists("index_name"); ok {
		tmp := indexName.(string)
		request.IndexName = &tmp
	}

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		request.TableNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "nosql")

	response, err := s.Client.GetIndex(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NosqlIndexDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NosqlIndexDataSource-", NosqlIndexDataSource(), s.D))

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	keys := []interface{}{}
	for _, item := range s.Res.Keys {
		keys = append(keys, IndexKeyToMap(item))
	}
	s.D.Set("keys", keys)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TableId != nil {
		s.D.Set("table_id", *s.Res.TableId)
	}

	if s.Res.TableName != nil {
		s.D.Set("table_name", *s.Res.TableName)
	}

	return nil
}
