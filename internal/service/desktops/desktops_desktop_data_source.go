// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package desktops

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_desktops "github.com/oracle/oci-go-sdk/v65/desktops"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DesktopsDesktopDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDesktopsDesktop,
		Schema: map[string]*schema.Schema{
			"desktop_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"device_policy": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"audio_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cdm_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"clipboard_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_display_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_keyboard_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_pointer_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_printing_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
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
			"hosting_options": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"connect_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"image": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"image_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"image_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"pool_id": {
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
			"user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDesktopsDesktop(d *schema.ResourceData, m interface{}) error {
	sync := &DesktopsDesktopDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DesktopServiceClient()

	return tfresource.ReadResource(sync)
}

type DesktopsDesktopDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_desktops.DesktopServiceClient
	Res    *oci_desktops.GetDesktopResponse
}

func (s *DesktopsDesktopDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DesktopsDesktopDataSourceCrud) Get() error {
	request := oci_desktops.GetDesktopRequest{}

	if desktopId, ok := s.D.GetOkExists("desktop_id"); ok {
		tmp := desktopId.(string)
		request.DesktopId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "desktops")

	response, err := s.Client.GetDesktop(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DesktopsDesktopDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DevicePolicy != nil {
		s.D.Set("device_policy", []interface{}{DesktopDevicePolicyToMap(s.Res.DevicePolicy)})
	} else {
		s.D.Set("device_policy", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostingOptions != nil {
		s.D.Set("hosting_options", []interface{}{HostingOptionsToMap(s.Res.HostingOptions)})
	} else {
		s.D.Set("hosting_options", nil)
	}

	if s.Res.PoolId != nil {
		s.D.Set("pool_id", *s.Res.PoolId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.UserName != nil {
		s.D.Set("user_name", *s.Res.UserName)
	}

	return nil
}
