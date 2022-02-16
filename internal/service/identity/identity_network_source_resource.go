// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"
)

func IdentityNetworkSourceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityNetworkSource,
		Read:     readIdentityNetworkSource,
		Update:   updateIdentityNetworkSource,
		Delete:   deleteIdentityNetworkSource,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"public_source_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"services": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"virtual_source_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 100,
				MinItems: 0,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vcn_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_ranges": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			// Computed
			"inactive_state": {
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

func createIdentityNetworkSource(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityNetworkSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityNetworkSource(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityNetworkSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentityNetworkSource(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityNetworkSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityNetworkSource(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityNetworkSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityNetworkSourceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.NetworkSources
	DisableNotFoundRetries bool
}

func (s *IdentityNetworkSourceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityNetworkSourceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.NetworkSourcesLifecycleStateCreating),
	}
}

func (s *IdentityNetworkSourceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.NetworkSourcesLifecycleStateActive),
	}
}

func (s *IdentityNetworkSourceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.NetworkSourcesLifecycleStateDeleting),
	}
}

func (s *IdentityNetworkSourceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.NetworkSourcesLifecycleStateDeleted),
	}
}

func (s *IdentityNetworkSourceResourceCrud) Create() error {
	request := oci_identity.CreateNetworkSourceRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if publicSourceList, ok := s.D.GetOkExists("public_source_list"); ok {
		interfaces := publicSourceList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("public_source_list") {
			request.PublicSourceList = tmp
		}
	}

	if services, ok := s.D.GetOkExists("services"); ok {
		interfaces := services.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("services") {
			request.Services = tmp
		}
	}

	if virtualSourceList, ok := s.D.GetOkExists("virtual_source_list"); ok {
		interfaces := virtualSourceList.([]interface{})
		tmp := make([]oci_identity.NetworkSourcesVirtualSourceList, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "virtual_source_list", stateDataIndex)
			converted, err := s.mapToNetworkSourcesVirtualSourceList(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("virtual_source_list") {
			request.VirtualSourceList = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateNetworkSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NetworkSources
	return nil
}

func (s *IdentityNetworkSourceResourceCrud) Get() error {
	request := oci_identity.GetNetworkSourceRequest{}

	tmp := s.D.Id()
	request.NetworkSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetNetworkSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NetworkSources
	return nil
}

func (s *IdentityNetworkSourceResourceCrud) Update() error {
	request := oci_identity.UpdateNetworkSourceRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.NetworkSourceId = &tmp

	if publicSourceList, ok := s.D.GetOkExists("public_source_list"); ok {
		interfaces := publicSourceList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("public_source_list") {
			request.PublicSourceList = tmp
		}
	}

	if services, ok := s.D.GetOkExists("services"); ok {
		interfaces := services.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("services") {
			request.Services = tmp
		}
	}

	if virtualSourceList, ok := s.D.GetOkExists("virtual_source_list"); ok {
		interfaces := virtualSourceList.([]interface{})
		tmp := make([]oci_identity.NetworkSourcesVirtualSourceList, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "virtual_source_list", stateDataIndex)
			converted, err := s.mapToNetworkSourcesVirtualSourceList(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("virtual_source_list") {
			request.VirtualSourceList = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateNetworkSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NetworkSources
	return nil
}

func (s *IdentityNetworkSourceResourceCrud) Delete() error {
	request := oci_identity.DeleteNetworkSourceRequest{}

	tmp := s.D.Id()
	request.NetworkSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteNetworkSource(context.Background(), request)
	return err
}

func (s *IdentityNetworkSourceResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("public_source_list", s.Res.PublicSourceList)

	s.D.Set("services", s.Res.Services)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	virtualSourceList := []interface{}{}
	for _, item := range s.Res.VirtualSourceList {
		virtualSourceList = append(virtualSourceList, networkSourcesVirtualSourceListToMap(item))
	}
	s.D.Set("virtual_source_list", virtualSourceList)

	return nil
}

func (s *IdentityNetworkSourceResourceCrud) mapToNetworkSourcesVirtualSourceList(fieldKeyFormat string) (oci_identity.NetworkSourcesVirtualSourceList, error) {
	result := oci_identity.NetworkSourcesVirtualSourceList{}

	if vcn_id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcn_id")); ok {
		tmp := vcn_id.(string)
		result.VcnId = &tmp
	}

	if ip_ranges, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ip_ranges")); ok {
		ranges := ip_ranges.([]interface{})
		tmp := make([]string, len(ranges))
		for i, ip_range := range ranges {
			tmp[i] = ip_range.(string)
		}
		result.IpRanges = tmp
	}

	return result, nil
}

func networkSourcesVirtualSourceListToMap(obj oci_identity.NetworkSourcesVirtualSourceList) map[string]interface{} {
	result := map[string]interface{}{}
	result["vcn_id"] = obj.VcnId
	result["ip_ranges"] = obj.IpRanges
	return result
}
