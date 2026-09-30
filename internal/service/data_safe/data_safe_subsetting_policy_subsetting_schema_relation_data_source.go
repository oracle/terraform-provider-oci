// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingPolicySubsettingSchemaRelationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["subsetting_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["subsetting_schema_relation_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DataSafeSubsettingPolicySubsettingSchemaRelationResource(), fieldMap, readSingularDataSafeSubsettingPolicySubsettingSchemaRelationWithContext)
}

func readSingularDataSafeSubsettingPolicySubsettingSchemaRelationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingSchemaRelationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingPolicySubsettingSchemaRelationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetSubsettingSchemaRelationResponse
}

func (s *DataSafeSubsettingPolicySubsettingSchemaRelationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingPolicySubsettingSchemaRelationDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetSubsettingSchemaRelationRequest{}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if subsettingSchemaRelationKey, ok := s.D.GetOkExists("subsetting_schema_relation_key"); ok {
		tmp := subsettingSchemaRelationKey.(string)
		request.SubsettingSchemaRelationKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetSubsettingSchemaRelation(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSubsettingPolicySubsettingSchemaRelationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSubsettingPolicySubsettingSchemaRelationDataSource-", DataSafeSubsettingPolicySubsettingSchemaRelationDataSource(), s.D))

	s.D.Set("child_columns", s.Res.ChildColumns)

	if s.Res.ChildObjectKey != nil {
		s.D.Set("child_object_key", *s.Res.ChildObjectKey)
	}

	if s.Res.ChildObjectName != nil {
		s.D.Set("child_object_name", *s.Res.ChildObjectName)
	}

	if s.Res.ChildSchemaName != nil {
		s.D.Set("child_schema_name", *s.Res.ChildSchemaName)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("parent_columns", s.Res.ParentColumns)

	if s.Res.ParentObjectKey != nil {
		s.D.Set("parent_object_key", *s.Res.ParentObjectKey)
	}

	if s.Res.ParentObjectName != nil {
		s.D.Set("parent_object_name", *s.Res.ParentObjectName)
	}

	if s.Res.ParentSchemaName != nil {
		s.D.Set("parent_schema_name", *s.Res.ParentSchemaName)
	}

	s.D.Set("relation_type", s.Res.RelationType)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
