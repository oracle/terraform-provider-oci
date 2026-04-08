// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubDynamicSetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["dynamic_set_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(OsManagementHubDynamicSetResource(), fieldMap, readSingularOsManagementHubDynamicSetWithContext)
}

func readSingularOsManagementHubDynamicSetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OsManagementHubDynamicSetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DynamicSetClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type OsManagementHubDynamicSetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.DynamicSetClient
	Res    *oci_os_management_hub.GetDynamicSetResponse
}

func (s *OsManagementHubDynamicSetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubDynamicSetDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_os_management_hub.GetDynamicSetRequest{}

	if dynamicSetId, ok := s.D.GetOkExists("dynamic_set_id"); ok {
		tmp := dynamicSetId.(string)
		request.DynamicSetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetDynamicSet(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubDynamicSetDataSourceCrud) SetData() error {
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

	s.D.Set("match_type", s.Res.MatchType)

	if s.Res.MatchingRule != nil {
		s.D.Set("matching_rule", []interface{}{MatchingRuleToMap(s.Res.MatchingRule)})
	} else {
		s.D.Set("matching_rule", nil)
	}

	if s.Res.ScheduledJobCount != nil {
		s.D.Set("scheduled_job_count", *s.Res.ScheduledJobCount)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	targetCompartments := []interface{}{}
	for _, item := range s.Res.TargetCompartments {
		targetCompartments = append(targetCompartments, TargetCompartmentDetailsToMap(item))
	}
	s.D.Set("target_compartments", targetCompartments)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
