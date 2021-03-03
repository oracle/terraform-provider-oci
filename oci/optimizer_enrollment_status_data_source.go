// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_optimizer "github.com/oracle/oci-go-sdk/v36/optimizer"
)

func init() {
	RegisterDatasource("oci_optimizer_enrollment_status", OptimizerEnrollmentStatusDataSource())
}

func OptimizerEnrollmentStatusDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["enrollment_status_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(OptimizerEnrollmentStatusResource(), fieldMap, readSingularOptimizerEnrollmentStatus)
}

func readSingularOptimizerEnrollmentStatus(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerEnrollmentStatusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).optimizerClient()

	return ReadResource(sync)
}

type OptimizerEnrollmentStatusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_optimizer.OptimizerClient
	Res    *oci_optimizer.GetEnrollmentStatusResponse
}

func (s *OptimizerEnrollmentStatusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OptimizerEnrollmentStatusDataSourceCrud) Get() error {
	request := oci_optimizer.GetEnrollmentStatusRequest{}

	if enrollmentStatusId, ok := s.D.GetOkExists("enrollment_status_id"); ok {
		tmp := enrollmentStatusId.(string)
		request.EnrollmentStatusId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "optimizer")

	response, err := s.Client.GetEnrollmentStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OptimizerEnrollmentStatusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.StatusReason != nil {
		s.D.Set("status_reason", *s.Res.StatusReason)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
