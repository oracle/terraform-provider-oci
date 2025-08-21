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

func DataSafeUnifiedAuditPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["unified_audit_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeUnifiedAuditPolicyResource(), fieldMap, readSingularDataSafeUnifiedAuditPolicy)
}

func readSingularDataSafeUnifiedAuditPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUnifiedAuditPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeUnifiedAuditPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetUnifiedAuditPolicyResponse
}

func (s *DataSafeUnifiedAuditPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeUnifiedAuditPolicyDataSourceCrud) Get() error {
	request := oci_data_safe.GetUnifiedAuditPolicyRequest{}

	if unifiedAuditPolicyId, ok := s.D.GetOkExists("unified_audit_policy_id"); ok {
		tmp := unifiedAuditPolicyId.(string)
		request.UnifiedAuditPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetUnifiedAuditPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeUnifiedAuditPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	conditions := []interface{}{}
	for _, item := range s.Res.Conditions {
		conditions = append(conditions, PolicyConditionToMap(item))
	}
	s.D.Set("conditions", conditions)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("enabled_entities", s.Res.EnabledEntities)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsSeeded != nil {
		s.D.Set("is_seeded", *s.Res.IsSeeded)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SecurityPolicyId != nil {
		s.D.Set("security_policy_id", *s.Res.SecurityPolicyId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.UnifiedAuditPolicyDefinitionId != nil {
		s.D.Set("unified_audit_policy_definition_id", *s.Res.UnifiedAuditPolicyDefinitionId)
	}

	return nil
}
