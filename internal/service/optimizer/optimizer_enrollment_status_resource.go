// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package optimizer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_optimizer "github.com/oracle/oci-go-sdk/v58/optimizer"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OptimizerEnrollmentStatusResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOptimizerEnrollmentStatus,
		Read:     readOptimizerEnrollmentStatus,
		Update:   updateOptimizerEnrollmentStatus,
		Delete:   deleteOptimizerEnrollmentStatus,
		Schema: map[string]*schema.Schema{
			// Required
			"enrollment_status_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_reason": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOptimizerEnrollmentStatus(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerEnrollmentStatusResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.CreateResource(d, sync)
}

func readOptimizerEnrollmentStatus(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerEnrollmentStatusResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

func updateOptimizerEnrollmentStatus(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerEnrollmentStatusResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOptimizerEnrollmentStatus(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OptimizerEnrollmentStatusResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_optimizer.OptimizerClient
	Res                    *oci_optimizer.EnrollmentStatus
	DisableNotFoundRetries bool
}

func (s *OptimizerEnrollmentStatusResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OptimizerEnrollmentStatusResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_optimizer.LifecycleStateAttaching),
		string(oci_optimizer.LifecycleStateCreating),
	}
}

func (s *OptimizerEnrollmentStatusResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_optimizer.LifecycleStateActive),
	}
}

func (s *OptimizerEnrollmentStatusResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_optimizer.LifecycleStateDetaching),
		string(oci_optimizer.LifecycleStateDeleting),
	}
}

func (s *OptimizerEnrollmentStatusResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_optimizer.LifecycleStateDeleted),
	}
}

func (s *OptimizerEnrollmentStatusResourceCrud) Create() error {
	request := oci_optimizer.UpdateEnrollmentStatusRequest{}

	if enrollmentStatusId, ok := s.D.GetOkExists("enrollment_status_id"); ok {
		tmp := enrollmentStatusId.(string)
		request.EnrollmentStatusId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_optimizer.OptimizerEnrollmentStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "optimizer")

	response, err := s.Client.UpdateEnrollmentStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EnrollmentStatus
	return nil
}

func (s *OptimizerEnrollmentStatusResourceCrud) Get() error {
	request := oci_optimizer.GetEnrollmentStatusRequest{}

	tmp := s.D.Id()
	request.EnrollmentStatusId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "optimizer")

	response, err := s.Client.GetEnrollmentStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EnrollmentStatus
	return nil
}

func (s *OptimizerEnrollmentStatusResourceCrud) Update() error {
	request := oci_optimizer.UpdateEnrollmentStatusRequest{}

	tmp := s.D.Id()
	request.EnrollmentStatusId = &tmp

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_optimizer.OptimizerEnrollmentStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "optimizer")

	response, err := s.Client.UpdateEnrollmentStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EnrollmentStatus
	return nil
}

func (s *OptimizerEnrollmentStatusResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Id != nil {
		s.D.Set("enrollment_status_id", *s.Res.Id)
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

func EnrollmentStatusSummaryToMap(obj oci_optimizer.EnrollmentStatusSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.StatusReason != nil {
		result["status_reason"] = string(*obj.StatusReason)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
