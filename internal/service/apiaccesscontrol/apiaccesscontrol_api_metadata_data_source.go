// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apiaccesscontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apiaccesscontrol "github.com/oracle/oci-go-sdk/v65/apiaccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApiaccesscontrolApiMetadataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularApiaccesscontrolApiMetadata,
		Schema: map[string]*schema.Schema{
			"api_metadata_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"api_name": {
				Type:     schema.TypeString,
				Computed: true,
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
			"entity_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fields": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_name": {
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
			"time_deleted": {
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

func readSingularApiaccesscontrolApiMetadata(d *schema.ResourceData, m interface{}) error {
	sync := &ApiaccesscontrolApiMetadataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiMetadataClient()

	return tfresource.ReadResource(sync)
}

type ApiaccesscontrolApiMetadataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apiaccesscontrol.ApiMetadataClient
	Res    *oci_apiaccesscontrol.GetApiMetadataResponse
}

func (s *ApiaccesscontrolApiMetadataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApiaccesscontrolApiMetadataDataSourceCrud) Get() error {
	request := oci_apiaccesscontrol.GetApiMetadataRequest{}

	if apiMetadataId, ok := s.D.GetOkExists("api_metadata_id"); ok {
		tmp := apiMetadataId.(string)
		request.ApiMetadataId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apiaccesscontrol")

	response, err := s.Client.GetApiMetadata(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApiaccesscontrolApiMetadataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ApiName != nil {
		s.D.Set("api_name", *s.Res.ApiName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EntityType != nil {
		s.D.Set("entity_type", *s.Res.EntityType)
	}

	s.D.Set("fields", s.Res.Fields)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Path != nil {
		s.D.Set("path", *s.Res.Path)
	}

	if s.Res.ServiceName != nil {
		s.D.Set("service_name", *s.Res.ServiceName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDeleted != nil {
		s.D.Set("time_deleted", s.Res.TimeDeleted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
