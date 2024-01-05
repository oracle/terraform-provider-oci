// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateConnectionAssignmentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["connection_assignment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GoldenGateConnectionAssignmentResource(), fieldMap, readSingularGoldenGateConnectionAssignment)
}

func readSingularGoldenGateConnectionAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateConnectionAssignmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateConnectionAssignmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.GetConnectionAssignmentResponse
}

func (s *GoldenGateConnectionAssignmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateConnectionAssignmentDataSourceCrud) Get() error {
	request := oci_golden_gate.GetConnectionAssignmentRequest{}

	if connectionAssignmentId, ok := s.D.GetOkExists("connection_assignment_id"); ok {
		tmp := connectionAssignmentId.(string)
		request.ConnectionAssignmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.GetConnectionAssignment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GoldenGateConnectionAssignmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AliasName != nil {
		s.D.Set("alias_name", *s.Res.AliasName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionId != nil {
		s.D.Set("connection_id", *s.Res.ConnectionId)
	}

	if s.Res.DeploymentId != nil {
		s.D.Set("deployment_id", *s.Res.DeploymentId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
