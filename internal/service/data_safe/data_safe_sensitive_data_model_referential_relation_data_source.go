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

func DataSafeSensitiveDataModelReferentialRelationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["sensitive_data_model_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeSensitiveDataModelReferentialRelationResource(), fieldMap, readSingularDataSafeSensitiveDataModelReferentialRelation)
}

func readSingularDataSafeSensitiveDataModelReferentialRelation(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelReferentialRelationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSensitiveDataModelReferentialRelationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetReferentialRelationResponse
}

func (s *DataSafeSensitiveDataModelReferentialRelationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSensitiveDataModelReferentialRelationDataSourceCrud) Get() error {
	request := oci_data_safe.GetReferentialRelationRequest{}

	if referentialRelationKey, ok := s.D.GetOkExists("key"); ok {
		tmp := referentialRelationKey.(string)
		request.ReferentialRelationKey = &tmp
	}

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetReferentialRelation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSensitiveDataModelReferentialRelationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSensitiveDataModelReferentialRelationDataSource-", DataSafeSensitiveDataModelReferentialRelationDataSource(), s.D))

	if s.Res.Child != nil {
		s.D.Set("child", []interface{}{ColumnsInfoToMap(s.Res.Child)})
	} else {
		s.D.Set("child", nil)
	}

	if s.Res.IsSensitive != nil {
		s.D.Set("is_sensitive", *s.Res.IsSensitive)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Parent != nil {
		s.D.Set("parent", []interface{}{ColumnsInfoToMap(s.Res.Parent)})
	} else {
		s.D.Set("parent", nil)
	}

	s.D.Set("relation_type", s.Res.RelationType)

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
