// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseVmClusterPatchDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseVmClusterPatch,
		Schema: map[string]*schema.Schema{
			"patch_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"available_actions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_released": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseVmClusterPatch(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterPatchDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseVmClusterPatchDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetVmClusterPatchResponse
}

func (s *DatabaseVmClusterPatchDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseVmClusterPatchDataSourceCrud) Get() error {
	request := oci_database.GetVmClusterPatchRequest{}

	if patchId, ok := s.D.GetOkExists("patch_id"); ok {
		tmp := patchId.(string)
		request.PatchId = &tmp
	}

	if vmClusterId, ok := s.D.GetOkExists("vm_cluster_id"); ok {
		tmp := vmClusterId.(string)
		request.VmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetVmClusterPatch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseVmClusterPatchDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("available_actions", s.Res.AvailableActions)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("last_action", s.Res.LastAction)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeReleased != nil {
		s.D.Set("time_released", s.Res.TimeReleased.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
