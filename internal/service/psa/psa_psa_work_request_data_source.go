// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psa "github.com/oracle/oci-go-sdk/v65/psa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsaPsaWorkRequestDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularPsaPsaWorkRequestWithContext,
		Schema: map[string]*schema.Schema{
			"work_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operation_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"percent_complete": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"action_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entity_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entity_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_accepted": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
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

func readSingularPsaPsaWorkRequestWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &PsaPsaWorkRequestDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivateServiceAccessClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type PsaPsaWorkRequestDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psa.PrivateServiceAccessClient
	Res    *oci_psa.GetPsaWorkRequestResponse
}

func (s *PsaPsaWorkRequestDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsaPsaWorkRequestDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_psa.GetPsaWorkRequestRequest{}

	if workRequestId, ok := s.D.GetOkExists("work_request_id"); ok {
		tmp := workRequestId.(string)
		request.WorkRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psa")

	response, err := s.Client.GetPsaWorkRequest(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *PsaPsaWorkRequestDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("operation_type", s.Res.OperationType)

	if s.Res.PercentComplete != nil {
		s.D.Set("percent_complete", *s.Res.PercentComplete)
	}

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, WorkRequestResourceToMap(item))
	}
	s.D.Set("resources", resources)

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
