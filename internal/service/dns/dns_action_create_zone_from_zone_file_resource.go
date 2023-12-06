// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"bytes"
	"context"
	"io/ioutil"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_dns "github.com/oracle/oci-go-sdk/v65/dns"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DnsActionCreateZoneFromZoneFileResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDnsActionCreateZoneFromZoneFile,
		Read:     readDnsActionCreateZoneFromZoneFile,
		Delete:   deleteDnsActionCreateZoneFromZoneFile,
		Schema: map[string]*schema.Schema{
			// Required
			"create_zone_from_zone_file_details": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"view_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"external_downstreams": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tsig_key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"external_masters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tsig_key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_protected": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nameservers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"self": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial": {
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
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone_transfer_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_transfer_destination": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_transfer_source": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"zone_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDnsActionCreateZoneFromZoneFile(d *schema.ResourceData, m interface{}) error {
	sync := &DnsActionCreateZoneFromZoneFileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.CreateResource(d, sync)
}

func readDnsActionCreateZoneFromZoneFile(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDnsActionCreateZoneFromZoneFile(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DnsActionCreateZoneFromZoneFileResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.Zone
	DisableNotFoundRetries bool
}

func (s *DnsActionCreateZoneFromZoneFileResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DnsActionCreateZoneFromZoneFileResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dns.ZoneLifecycleStateCreating),
	}
}

func (s *DnsActionCreateZoneFromZoneFileResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dns.ZoneLifecycleStateActive),
	}
}

func (s *DnsActionCreateZoneFromZoneFileResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dns.ZoneLifecycleStateDeleting),
	}
}

func (s *DnsActionCreateZoneFromZoneFileResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dns.ZoneLifecycleStateDeleted),
	}
}

func (s *DnsActionCreateZoneFromZoneFileResourceCrud) Create() error {
	request := oci_dns.CreateZoneFromZoneFileRequest{}

	if createZoneFromZoneFileDetails, ok := s.D.GetOkExists("create_zone_from_zone_file_details"); ok {
		tmp := []byte(createZoneFromZoneFileDetails.(string))
		request.CreateZoneFromZoneFileDetails = ioutil.NopCloser(bytes.NewReader(tmp))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.CreateZoneFromZoneFileScopeEnum(scope.(string))
	}

	if viewId, ok := s.D.GetOkExists("view_id"); ok {
		tmp := viewId.(string)
		request.ViewId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.CreateZoneFromZoneFile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Zone
	return nil
}

func (s *DnsActionCreateZoneFromZoneFileResourceCrud) Get() error {
	request := oci_dns.GetZoneRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.GetZoneScopeEnum(scope.(string))
	}

	tmp := s.D.Id()
	request.ZoneNameOrId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.GetZone(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Zone
	return nil
}

func (s *DnsActionCreateZoneFromZoneFileResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	externalDownstreams := []interface{}{}
	for _, item := range s.Res.ExternalDownstreams {
		externalDownstreams = append(externalDownstreams, ExternalDownstreamToMap(item))
	}
	s.D.Set("external_downstreams", externalDownstreams)

	externalMasters := []interface{}{}
	for _, item := range s.Res.ExternalMasters {
		externalMasters = append(externalMasters, ExternalMasterToMap(item))
	}
	s.D.Set("external_masters", externalMasters)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsProtected != nil {
		s.D.Set("is_protected", *s.Res.IsProtected)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	nameservers := []interface{}{}
	for _, item := range s.Res.Nameservers {
		nameservers = append(nameservers, NameserverToMap(item))
	}
	s.D.Set("nameservers", nameservers)

	s.D.Set("scope", s.Res.Scope)

	if s.Res.Self != nil {
		s.D.Set("self", *s.Res.Self)
	}

	if s.Res.Serial != nil {
		s.D.Set("serial", strconv.FormatInt(*s.Res.Serial, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	if s.Res.ViewId != nil {
		s.D.Set("view_id", *s.Res.ViewId)
	}

	zoneTransferServers := []interface{}{}
	for _, item := range s.Res.ZoneTransferServers {
		zoneTransferServers = append(zoneTransferServers, ZoneTransferServerToMap(item))
	}
	s.D.Set("zone_transfer_servers", zoneTransferServers)

	s.D.Set("zone_type", s.Res.ZoneType)

	return nil
}
