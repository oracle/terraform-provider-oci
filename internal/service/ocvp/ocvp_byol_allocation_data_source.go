// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpByolAllocationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["byol_allocation_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(OcvpByolAllocationResource(), fieldMap, readSingularOcvpByolAllocationWithContext)
}

func readSingularOcvpByolAllocationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OcvpByolAllocationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ByolAllocationClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type OcvpByolAllocationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.ByolAllocationClient
	Res    *oci_ocvp.GetByolAllocationResponse
}

func (s *OcvpByolAllocationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpByolAllocationDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_ocvp.GetByolAllocationRequest{}

	if byolAllocationId, ok := s.D.GetOkExists("byol_allocation_id"); ok {
		tmp := byolAllocationId.(string)
		request.ByolAllocationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.GetByolAllocation(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OcvpByolAllocationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AllocatedUnits != nil {
		s.D.Set("allocated_units", *s.Res.AllocatedUnits)
	}

	if s.Res.AvailableUnits != nil {
		s.D.Set("available_units", *s.Res.AvailableUnits)
	}

	if s.Res.ByolId != nil {
		s.D.Set("byol_id", *s.Res.ByolId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EntitlementKey != nil {
		s.D.Set("entitlement_key", *s.Res.EntitlementKey)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("software_type", s.Res.SoftwareType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeTermEnd != nil {
		s.D.Set("time_term_end", s.Res.TimeTermEnd.String())
	}

	if s.Res.TimeTermStart != nil {
		s.D.Set("time_term_start", s.Res.TimeTermStart.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
