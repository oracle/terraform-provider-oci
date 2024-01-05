// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FusionAppsFusionEnvironmentScheduledActivityDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularFusionAppsFusionEnvironmentScheduledActivity,
		Schema: map[string]*schema.Schema{
			"fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scheduled_activity_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"actions": {
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
						"artifact": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"qualifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reference_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"delay_in_hours": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"run_cycle": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_availability": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_expected_finish": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_scheduled_start": {
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

func readSingularFusionAppsFusionEnvironmentScheduledActivity(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentScheduledActivityDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentScheduledActivityDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.GetScheduledActivityResponse
}

func (s *FusionAppsFusionEnvironmentScheduledActivityDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentScheduledActivityDataSourceCrud) Get() error {
	request := oci_fusion_apps.GetScheduledActivityRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	if scheduledActivityId, ok := s.D.GetOkExists("scheduled_activity_id"); ok {
		tmp := scheduledActivityId.(string)
		request.ScheduledActivityId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.GetScheduledActivity(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FusionAppsFusionEnvironmentScheduledActivityDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	actions := []interface{}{}
	for _, item := range s.Res.Actions {
		actions = append(actions, ActionToMap(item))
	}
	s.D.Set("actions", actions)

	if s.Res.DelayInHours != nil {
		s.D.Set("delay_in_hours", *s.Res.DelayInHours)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	s.D.Set("run_cycle", s.Res.RunCycle)

	s.D.Set("service_availability", s.Res.ServiceAvailability)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExpectedFinish != nil {
		s.D.Set("time_expected_finish", s.Res.TimeExpectedFinish.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeScheduledStart != nil {
		s.D.Set("time_scheduled_start", s.Res.TimeScheduledStart.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
