// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	"time"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func IpSecConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Create: createIpSecConnection,
		Read:   readIpSecConnection,
		Update: updateIpSecConnection,
		Delete: deleteIpSecConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cpe_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"static_routes": {
				Type: schema.TypeList,
				// @CODEGEN 1/2018: Existing provider allows static_routes to be empty.
				// Avoid breaking change by keeping this optional, even though spec says it's
				// required.
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"id": {
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
		},
	}
}

func createIpSecConnection(d *schema.ResourceData, m interface{}) error {
	sync := &IpSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readIpSecConnection(d *schema.ResourceData, m interface{}) error {
	sync := &IpSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateIpSecConnection(d *schema.ResourceData, m interface{}) error {
	sync := &IpSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteIpSecConnection(d *schema.ResourceData, m interface{}) error {
	sync := &IpSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type IpSecConnectionResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.IpSecConnection
	DisableNotFoundRetries bool
}

func (s *IpSecConnectionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IpSecConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateProvisioning),
	}
}

func (s *IpSecConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateAvailable),
	}
}

func (s *IpSecConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateTerminating),
	}
}

func (s *IpSecConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateTerminated),
	}
}

func (s *IpSecConnectionResourceCrud) Create() error {
	request := oci_core.CreateIPSecConnectionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpeId, ok := s.D.GetOkExists("cpe_id"); ok {
		tmp := cpeId.(string)
		request.CpeId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.StaticRoutes = []string{}
	if staticRoutes, ok := s.D.GetOkExists("static_routes"); ok {
		interfaces := staticRoutes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = toBeConverted.(string)
		}
		request.StaticRoutes = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateIPSecConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IpSecConnection
	return nil
}

func (s *IpSecConnectionResourceCrud) Get() error {
	request := oci_core.GetIPSecConnectionRequest{}

	tmp := s.D.Id()
	request.IpscId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetIPSecConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IpSecConnection
	return nil
}

func (s *IpSecConnectionResourceCrud) Update() error {
	request := oci_core.UpdateIPSecConnectionRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.IpscId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateIPSecConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IpSecConnection
	return nil
}

func (s *IpSecConnectionResourceCrud) Delete() error {
	request := oci_core.DeleteIPSecConnectionRequest{}

	tmp := s.D.Id()
	request.IpscId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteIPSecConnection(context.Background(), request)
	return err
}

func (s *IpSecConnectionResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpeId != nil {
		s.D.Set("cpe_id", *s.Res.CpeId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DrgId != nil {
		s.D.Set("drg_id", *s.Res.DrgId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("static_routes", s.Res.StaticRoutes)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
