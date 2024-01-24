// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package appmgmt_control

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_appmgmt_control "github.com/oracle/oci-go-sdk/v65/appmgmtcontrol"
)

func AppmgmtControlMonitoredInstanceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularAppmgmtControlMonitoredInstance,
		Schema: map[string]*schema.Schema{
			"monitored_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_agent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"monitoring_state": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularAppmgmtControlMonitoredInstance(d *schema.ResourceData, m interface{}) error {
	sync := &AppmgmtControlMonitoredInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AppmgmtControlClient()

	return tfresource.ReadResource(sync)
}

type AppmgmtControlMonitoredInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_appmgmt_control.AppmgmtControlClient
	Res    *oci_appmgmt_control.GetMonitoredInstanceResponse
}

func (s *AppmgmtControlMonitoredInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AppmgmtControlMonitoredInstanceDataSourceCrud) Get() error {
	request := oci_appmgmt_control.GetMonitoredInstanceRequest{}

	if monitoredInstanceId, ok := s.D.GetOkExists("monitored_instance_id"); ok {
		tmp := monitoredInstanceId.(string)
		request.MonitoredInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "appmgmt_control")

	response, err := s.Client.GetMonitoredInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AppmgmtControlMonitoredInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AppmgmtControlMonitoredInstanceDataSource-", AppmgmtControlMonitoredInstanceDataSource(), s.D))

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.InstanceId != nil {
		s.D.Set("instance_id", *s.Res.InstanceId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ManagementAgentId != nil {
		s.D.Set("management_agent_id", *s.Res.ManagementAgentId)
	}

	s.D.Set("monitoring_state", s.Res.MonitoringState)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
