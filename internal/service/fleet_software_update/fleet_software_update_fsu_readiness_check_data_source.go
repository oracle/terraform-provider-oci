// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_software_update

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_software_update "github.com/oracle/oci-go-sdk/v65/fleetsoftwareupdate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetSoftwareUpdateFsuReadinessCheckDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fsu_readiness_check_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(FleetSoftwareUpdateFsuReadinessCheckResource(), fieldMap, readSingularFleetSoftwareUpdateFsuReadinessCheckWithContext)
}

func readSingularFleetSoftwareUpdateFsuReadinessCheckWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FleetSoftwareUpdateFsuReadinessCheckDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type FleetSoftwareUpdateFsuReadinessCheckDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_software_update.FleetSoftwareUpdateClient
	Res    *oci_fleet_software_update.GetFsuReadinessCheckResponse
}

func (s *FleetSoftwareUpdateFsuReadinessCheckDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetSoftwareUpdateFsuReadinessCheckDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_fleet_software_update.GetFsuReadinessCheckRequest{}

	if fsuReadinessCheckId, ok := s.D.GetOkExists("fsu_readiness_check_id"); ok {
		tmp := fsuReadinessCheckId.(string)
		request.FsuReadinessCheckId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_software_update")

	response, err := s.Client.GetFsuReadinessCheck(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetSoftwareUpdateFsuReadinessCheckDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.FsuReadinessCheck).(type) {
	case oci_fleet_software_update.TargetFsuReadinessCheck:
		s.D.Set("type", "TARGET")

		targets := []interface{}{}
		for _, item := range v.Targets {
			targets = append(targets, ReadinessCheckTargetEntryToMap(item))
		}
		s.D.Set("targets", targets)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.IssueCount != nil {
			s.D.Set("issue_count", *v.IssueCount)
		}

		issues := []interface{}{}
		for _, item := range v.Issues {
			issues = append(issues, PatchingIssueEntryToMap(item))
		}
		s.D.Set("issues", issues)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeFinished != nil {
			s.D.Set("time_finished", v.TimeFinished.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.FsuReadinessCheck)
		return nil
	}

	return nil
}
