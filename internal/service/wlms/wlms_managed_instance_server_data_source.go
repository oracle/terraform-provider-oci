// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package wlms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_wlms "github.com/oracle/oci-go-sdk/v65/wlms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WlmsManagedInstanceServerDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularWlmsManagedInstanceServer,
		Schema: map[string]*schema.Schema{
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"host_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_admin": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"jdk_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"jdk_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"latest_patches_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"middleware_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"middleware_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patch_readiness_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"restart_order": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
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
			"weblogic_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wls_domain_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wls_domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wls_domain_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularWlmsManagedInstanceServer(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsManagedInstanceServerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsManagedInstanceServerDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.GetManagedInstanceServerResponse
}

func (s *WlmsManagedInstanceServerDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsManagedInstanceServerDataSourceCrud) Get() error {
	request := oci_wlms.GetManagedInstanceServerRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if serverId, ok := s.D.GetOkExists("server_id"); ok {
		tmp := serverId.(string)
		request.ServerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.GetManagedInstanceServer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WlmsManagedInstanceServerDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	if s.Res.IsAdmin != nil {
		s.D.Set("is_admin", *s.Res.IsAdmin)
	}

	if s.Res.JdkPath != nil {
		s.D.Set("jdk_path", *s.Res.JdkPath)
	}

	if s.Res.JdkVersion != nil {
		s.D.Set("jdk_version", *s.Res.JdkVersion)
	}

	s.D.Set("latest_patches_status", s.Res.LatestPatchesStatus)

	if s.Res.MiddlewarePath != nil {
		s.D.Set("middleware_path", *s.Res.MiddlewarePath)
	}

	if s.Res.MiddlewareType != nil {
		s.D.Set("middleware_type", *s.Res.MiddlewareType)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("patch_readiness_status", s.Res.PatchReadinessStatus)

	if s.Res.RestartOrder != nil {
		s.D.Set("restart_order", *s.Res.RestartOrder)
	}

	if s.Res.Status != nil {
		s.D.Set("status", *s.Res.Status)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.WeblogicVersion != nil {
		s.D.Set("weblogic_version", *s.Res.WeblogicVersion)
	}

	if s.Res.WlsDomainId != nil {
		s.D.Set("wls_domain_id", *s.Res.WlsDomainId)
	}

	if s.Res.WlsDomainName != nil {
		s.D.Set("wls_domain_name", *s.Res.WlsDomainName)
	}

	if s.Res.WlsDomainPath != nil {
		s.D.Set("wls_domain_path", *s.Res.WlsDomainPath)
	}

	return nil
}
