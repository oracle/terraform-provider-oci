// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v56/containerengine"
)

func ContainerengineMigrateToNativeVcnStatusDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularContainerengineMigrateToNativeVcnStatus,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_decommission_scheduled": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularContainerengineMigrateToNativeVcnStatus(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineMigrateToNativeVcnStatusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineMigrateToNativeVcnStatusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetClusterMigrateToNativeVcnStatusResponse
}

func (s *ContainerengineMigrateToNativeVcnStatusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineMigrateToNativeVcnStatusDataSourceCrud) Get() error {
	request := oci_containerengine.GetClusterMigrateToNativeVcnStatusRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.GetClusterMigrateToNativeVcnStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineMigrateToNativeVcnStatusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineMigrateToNativeVcnStatusDataSource-",
		ContainerengineMigrateToNativeVcnStatusDataSource(), s.D))

	s.D.Set("state", s.Res.State)

	if s.Res.TimeDecommissionScheduled != nil {
		s.D.Set("time_decommission_scheduled", s.Res.TimeDecommissionScheduled.String())
	}

	return nil
}
