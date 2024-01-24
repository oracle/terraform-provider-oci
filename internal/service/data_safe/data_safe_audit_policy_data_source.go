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

func DataSafeAuditPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["audit_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeAuditPolicyResource(), fieldMap, readSingularDataSafeAuditPolicy)
}

func readSingularDataSafeAuditPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetAuditPolicyResponse
}

func (s *DataSafeAuditPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditPolicyDataSourceCrud) Get() error {
	request := oci_data_safe.GetAuditPolicyRequest{}

	if auditPolicyId, ok := s.D.GetOkExists("audit_policy_id"); ok {
		tmp := auditPolicyId.(string)
		request.AuditPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetAuditPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeAuditPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	auditConditions := []interface{}{}
	for _, item := range s.Res.AuditConditions {
		auditConditions = append(auditConditions, AuditConditionsToMap(item))
	}
	s.D.Set("audit_conditions", auditConditions)

	auditSpecifications := []interface{}{}
	for _, item := range s.Res.AuditSpecifications {
		auditSpecifications = append(auditSpecifications, AuditSpecificationToMap(item))
	}
	s.D.Set("audit_specifications", auditSpecifications)

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

	if s.Res.IsDataSafeServiceAccountExcluded != nil {
		s.D.Set("is_data_safe_service_account_excluded", *s.Res.IsDataSafeServiceAccountExcluded)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastProvisioned != nil {
		s.D.Set("time_last_provisioned", s.Res.TimeLastProvisioned.String())
	}

	if s.Res.TimeLastRetrieved != nil {
		s.D.Set("time_last_retrieved", s.Res.TimeLastRetrieved.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
