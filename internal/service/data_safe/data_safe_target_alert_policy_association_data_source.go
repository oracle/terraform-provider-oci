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

func DataSafeTargetAlertPolicyAssociationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["target_alert_policy_association_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeTargetAlertPolicyAssociationResource(), fieldMap, readSingularDataSafeTargetAlertPolicyAssociation)
}

func readSingularDataSafeTargetAlertPolicyAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetAlertPolicyAssociationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeTargetAlertPolicyAssociationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetTargetAlertPolicyAssociationResponse
}

func (s *DataSafeTargetAlertPolicyAssociationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeTargetAlertPolicyAssociationDataSourceCrud) Get() error {
	request := oci_data_safe.GetTargetAlertPolicyAssociationRequest{}

	if targetAlertPolicyAssociationId, ok := s.D.GetOkExists("target_alert_policy_association_id"); ok {
		tmp := targetAlertPolicyAssociationId.(string)
		request.TargetAlertPolicyAssociationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetTargetAlertPolicyAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeTargetAlertPolicyAssociationDataSourceCrud) SetData() error {
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

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.PolicyId != nil {
		s.D.Set("policy_id", *s.Res.PolicyId)
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

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
