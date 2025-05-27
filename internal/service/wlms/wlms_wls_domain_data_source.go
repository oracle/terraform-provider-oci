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

func WlmsWlsDomainDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularWlmsWlsDomain,
		Schema: map[string]*schema.Schema{
			"wls_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"admin_server_control_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"admin_server_start_script_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"admin_server_stop_script_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_patch_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_rollback_on_failure": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"managed_server_control_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"managed_server_start_script_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"managed_server_stop_script_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"servers_shutdown_timeout": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_accepted_terms_and_conditions": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"middleware_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patch_readiness_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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
		},
	}
}

func readSingularWlmsWlsDomain(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsWlsDomainDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsWlsDomainDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.GetWlsDomainResponse
}

func (s *WlmsWlsDomainDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsWlsDomainDataSourceCrud) Get() error {
	request := oci_wlms.GetWlsDomainRequest{}

	if wlsDomainId, ok := s.D.GetOkExists("wls_domain_id"); ok {
		tmp := wlsDomainId.(string)
		request.WlsDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.GetWlsDomain(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WlmsWlsDomainDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Configuration != nil {
		s.D.Set("configuration", []interface{}{WlsDomainConfigurationToMap(s.Res.Configuration)})
	} else {
		s.D.Set("configuration", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAcceptedTermsAndConditions != nil {
		s.D.Set("is_accepted_terms_and_conditions", *s.Res.IsAcceptedTermsAndConditions)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MiddlewareType != nil {
		s.D.Set("middleware_type", *s.Res.MiddlewareType)
	}

	s.D.Set("patch_readiness_status", s.Res.PatchReadinessStatus)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
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

	return nil
}
