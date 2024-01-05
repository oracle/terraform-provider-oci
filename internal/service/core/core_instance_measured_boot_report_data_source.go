// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreInstanceMeasuredBootReportDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreInstanceMeasuredBootReport,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"is_policy_verification_successful": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"measurements": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"actual": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"hash_algorithm": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pcr_index": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"policy": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"hash_algorithm": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pcr_index": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readSingularCoreInstanceMeasuredBootReport(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceMeasuredBootReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreInstanceMeasuredBootReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetMeasuredBootReportResponse
}

func (s *CoreInstanceMeasuredBootReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstanceMeasuredBootReportDataSourceCrud) Get() error {
	request := oci_core.GetMeasuredBootReportRequest{}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetMeasuredBootReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreInstanceMeasuredBootReportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreInstanceMeasuredBootReportDataSource-", CoreInstanceMeasuredBootReportDataSource(), s.D))

	if s.Res.IsPolicyVerificationSuccessful != nil {
		s.D.Set("is_policy_verification_successful", *s.Res.IsPolicyVerificationSuccessful)
	}

	if s.Res.Measurements != nil {
		s.D.Set("measurements", []interface{}{MeasuredBootReportMeasurementsToMap(s.Res.Measurements)})
	} else {
		s.D.Set("measurements", nil)
	}

	return nil
}

func MeasuredBootEntryToMap(obj oci_core.MeasuredBootEntry) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HashAlgorithm != nil {
		result["hash_algorithm"] = string(*obj.HashAlgorithm)
	}

	if obj.PcrIndex != nil {
		result["pcr_index"] = string(*obj.PcrIndex)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func MeasuredBootReportMeasurementsToMap(obj *oci_core.MeasuredBootReportMeasurements) map[string]interface{} {
	result := map[string]interface{}{}

	actual := []interface{}{}
	for _, item := range obj.Actual {
		actual = append(actual, MeasuredBootEntryToMap(item))
	}
	result["actual"] = actual

	policy := []interface{}{}
	for _, item := range obj.Policy {
		policy = append(policy, MeasuredBootEntryToMap(item))
	}
	result["policy"] = policy

	return result
}
