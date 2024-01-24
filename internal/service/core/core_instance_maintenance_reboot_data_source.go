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

func CoreInstanceMaintenanceRebootDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreInstanceMaintenanceReboot,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"time_maintenance_reboot_due_max": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreInstanceMaintenanceReboot(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceMaintenanceRebootDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreInstanceMaintenanceRebootDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetInstanceMaintenanceRebootResponse
}

func (s *CoreInstanceMaintenanceRebootDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstanceMaintenanceRebootDataSourceCrud) Get() error {
	request := oci_core.GetInstanceMaintenanceRebootRequest{}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetInstanceMaintenanceReboot(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreInstanceMaintenanceRebootDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreInstanceMaintenanceRebootDataSource-", CoreInstanceMaintenanceRebootDataSource(), s.D))

	if s.Res.TimeMaintenanceRebootDueMax != nil {
		s.D.Set("time_maintenance_reboot_due_max", s.Res.TimeMaintenanceRebootDueMax.String())
	}

	return nil
}
