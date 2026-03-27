// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateDeploymentDisasterRecoveryPrecheckReportDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularGoldenGateDeploymentDisasterRecoveryPrecheckReportWithContext,
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"deployment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fault_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"checks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"corrective_action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"related_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"related_resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"precheck_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_precheck_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_precheck_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularGoldenGateDeploymentDisasterRecoveryPrecheckReportWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GoldenGateDeploymentDisasterRecoveryPrecheckReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GoldenGateDeploymentDisasterRecoveryPrecheckReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.GetDisasterRecoveryPrecheckReportResponse
}

func (s *GoldenGateDeploymentDisasterRecoveryPrecheckReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateDeploymentDisasterRecoveryPrecheckReportDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_golden_gate.GetDisasterRecoveryPrecheckReportRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
		tmp := deploymentId.(string)
		request.DeploymentId = &tmp
	}

	if faultDomain, ok := s.D.GetOkExists("fault_domain"); ok {
		tmp := faultDomain.(string)
		request.FaultDomain = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.GetDisasterRecoveryPrecheckReport(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GoldenGateDeploymentDisasterRecoveryPrecheckReportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateDeploymentDisasterRecoveryPrecheckReportDataSource-", GoldenGateDeploymentDisasterRecoveryPrecheckReportDataSource(), s.D))

	checks := []interface{}{}
	for _, item := range s.Res.Checks {
		checks = append(checks, DisasterRecoveryPrecheckResultToMap(item))
	}
	s.D.Set("checks", checks)

	s.D.Set("precheck_status", s.Res.PrecheckStatus)

	if s.Res.TimePrecheckFinished != nil {
		s.D.Set("time_precheck_finished", s.Res.TimePrecheckFinished.String())
	}

	if s.Res.TimePrecheckStarted != nil {
		s.D.Set("time_precheck_started", s.Res.TimePrecheckStarted.String())
	}

	return nil
}

func DisasterRecoveryPrecheckResultToMap(obj oci_golden_gate.DisasterRecoveryPrecheckResult) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Code != nil {
		result["code"] = string(*obj.Code)
	}

	if obj.CorrectiveAction != nil {
		result["corrective_action"] = string(*obj.CorrectiveAction)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.RelatedResourceId != nil {
		result["related_resource_id"] = string(*obj.RelatedResourceId)
	}

	result["related_resource_type"] = string(obj.RelatedResourceType)

	result["status"] = string(obj.Status)

	return result
}
