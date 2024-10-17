// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_dns "github.com/oracle/oci-go-sdk/v65/dns"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DnsZoneResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDnsZone,
		Read:     readDnsZone,
		Update:   updateDnsZone,
		Delete:   deleteDnsZone,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_dns.ZoneSummaryZoneTypePrimary),
					string(oci_dns.ZoneSummaryZoneTypeSecondary)}, false),
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"dnssec_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_downstreams": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"address": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tsig_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"external_masters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"address": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tsig_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"view_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// Computed
			"dnssec_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ksk_dnssec_key_versions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"algorithm": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ds_data": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"digest_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"rdata": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"key_tag": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"length_in_bytes": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"predecessor_dnssec_key_version_uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"successor_dnssec_key_version_uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_activated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_expired": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_inactivated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_promoted": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_published": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_unpublished": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"zsk_dnssec_key_versions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"algorithm": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_tag": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"length_in_bytes": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"predecessor_dnssec_key_version_uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"successor_dnssec_key_version_uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_activated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_expired": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_inactivated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_promoted": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_published": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_unpublished": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"is_protected": {
				Type:     schema.TypeBool,
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
				Type:     schema.TypeInt,
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
		},
	}
}

func createDnsZone(d *schema.ResourceData, m interface{}) error {
	sync := &DnsZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.CreateResource(d, sync)
}

func readDnsZone(d *schema.ResourceData, m interface{}) error {
	sync := &DnsZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

func updateDnsZone(d *schema.ResourceData, m interface{}) error {
	sync := &DnsZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDnsZone(d *schema.ResourceData, m interface{}) error {
	sync := &DnsZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()
	sync.DisableNotFoundRetries = true

	// Check if the zone is protected. If it is, attempting to delete it will fail,
	// so don't do anything.
	if isProtected := d.Get("is_protected"); isProtected != nil {
		if isProtectedBool, ok := isProtected.(bool); ok && isProtectedBool {
			log.Printf("[WARN] Not attempting to delete protected zone with ID %s", d.Id())
			sync.VoidState()
			return nil
		}
	}

	return tfresource.DeleteResource(d, sync)
}

type DnsZoneResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.Zone
	DisableNotFoundRetries bool
}

func (s *DnsZoneResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DnsZoneResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dns.ZoneLifecycleStateCreating),
	}
}

func (s *DnsZoneResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dns.ZoneLifecycleStateActive),
	}
}

func (s *DnsZoneResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dns.ZoneLifecycleStateDeleting),
	}
}

func (s *DnsZoneResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dns.ZoneLifecycleStateDeleted),
	}
}

func (s *DnsZoneResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_dns.ZoneLifecycleStateUpdating),
	}
}

func (s *DnsZoneResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_dns.ZoneLifecycleStateActive),
	}
}

func (s *DnsZoneResourceCrud) Create() error {
	request := oci_dns.CreateZoneRequest{}
	createZoneDetailsRequest := oci_dns.CreateZoneDetails{}
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		createZoneDetailsRequest.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		createZoneDetailsRequest.DefinedTags = convertedDefinedTags
	}

	if dnssecState, ok := s.D.GetOkExists("dnssec_state"); ok {
		createZoneDetailsRequest.DnssecState = oci_dns.ZoneDnssecStateEnum(dnssecState.(string))
	}

	if externalDownstreams, ok := s.D.GetOkExists("external_downstreams"); ok {
		interfaces := externalDownstreams.([]interface{})
		tmp := make([]oci_dns.ExternalDownstream, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "external_downstreams", stateDataIndex)
			converted, err := s.mapToExternalDownstream(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("external_downstreams") {
			createZoneDetailsRequest.ExternalDownstreams = tmp
		}
	}

	if externalMasters, ok := s.D.GetOkExists("external_masters"); ok {
		interfaces := externalMasters.([]interface{})
		tmp := make([]oci_dns.ExternalMaster, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "external_masters", stateDataIndex)
			converted, err := s.mapToExternalMaster(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("external_masters") {
			createZoneDetailsRequest.ExternalMasters = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		createZoneDetailsRequest.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		createZoneDetailsRequest.Name = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.CreateZoneScopeEnum(scope.(string))
		createZoneDetailsRequest.Scope = oci_dns.ScopeEnum(scope.(string))
	}

	if viewId, ok := s.D.GetOkExists("view_id"); ok {
		tmp := viewId.(string)
		createZoneDetailsRequest.ViewId = &tmp
	}

	if zoneType, ok := s.D.GetOkExists("zone_type"); ok {
		createZoneDetailsRequest.ZoneType = oci_dns.CreateZoneDetailsZoneTypeEnum(zoneType.(string))
	}
	request.CreateZoneDetails = createZoneDetailsRequest

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.CreateZone(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Zone
	return nil
}

func (s *DnsZoneResourceCrud) Get() error {
	request := oci_dns.GetZoneRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.GetZoneScopeEnum(scope.(string))
	}

	if viewId, ok := s.D.GetOkExists("view_id"); ok {
		tmp := viewId.(string)
		request.ViewId = &tmp
	}

	zoneNameOrId, scope, viewId, err := parseZoneCompositeId(s.D.Id())
	if err == nil {
		request.ZoneNameOrId = &zoneNameOrId
		s.D.SetId(zoneNameOrId)
		request.Scope = oci_dns.GetZoneScopeEnum(scope)
		s.D.Set("scope", scope)
		if viewId != "" {
			request.ViewId = &viewId
			s.D.Set("view_id", viewId)
		}
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
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

func (s *DnsZoneResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dns.UpdateZoneRequest{}

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

	if dnssecState, ok := s.D.GetOkExists("dnssec_state"); ok {
		request.DnssecState = oci_dns.ZoneDnssecStateEnum(dnssecState.(string))
	}

	if externalDownstreams, ok := s.D.GetOkExists("external_downstreams"); ok {
		interfaces := externalDownstreams.([]interface{})
		tmp := make([]oci_dns.ExternalDownstream, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "external_downstreams", stateDataIndex)
			converted, err := s.mapToExternalDownstream(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("external_downstreams") {
			request.ExternalDownstreams = tmp
		}
	}

	if externalMasters, ok := s.D.GetOkExists("external_masters"); ok {
		interfaces := externalMasters.([]interface{})
		tmp := make([]oci_dns.ExternalMaster, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "external_masters", stateDataIndex)
			converted, err := s.mapToExternalMaster(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("external_masters") {
			request.ExternalMasters = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.UpdateZoneScopeEnum(scope.(string))
	}

	if viewId, ok := s.D.GetOkExists("view_id"); ok {
		tmp := viewId.(string)
		request.ViewId = &tmp
	}

	tmp := s.D.Id()
	request.ZoneNameOrId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.UpdateZone(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Zone
	return nil
}

func (s *DnsZoneResourceCrud) Delete() error {
	request := oci_dns.DeleteZoneRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.DeleteZoneScopeEnum(scope.(string))
	}

	if viewId, ok := s.D.GetOkExists("view_id"); ok {
		tmp := viewId.(string)
		request.ViewId = &tmp
	}

	tmp := s.D.Id()
	request.ZoneNameOrId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	_, err := s.Client.DeleteZone(context.Background(), request)
	return err
}

func (s *DnsZoneResourceCrud) SetData() error {
	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DnssecConfig != nil {
		s.D.Set("dnssec_config", []interface{}{DnssecConfigToMap(s.Res.DnssecConfig)})
	} else {
		s.D.Set("dnssec_config", nil)
	}

	s.D.Set("dnssec_state", s.Res.DnssecState)

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

	s.D.Set("scope", s.Res.Scope)

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

	if s.Res.Self != nil {
		s.D.Set("self", *s.Res.Self)
	}

	if s.Res.Serial != nil {
		s.D.Set("serial", *s.Res.Serial)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	s.D.Set("zone_type", s.Res.ZoneType)

	if s.Res.ViewId != nil {
		s.D.Set("view_id", *s.Res.ViewId)
	}

	if s.Res.IsProtected != nil {
		s.D.Set("is_protected", *s.Res.IsProtected)
	}

	zoneTransferServers := []interface{}{}
	for _, item := range s.Res.ZoneTransferServers {
		zoneTransferServers = append(zoneTransferServers, ZoneTransferServerToMap(item))
	}
	s.D.Set("zone_transfer_servers", zoneTransferServers)

	return nil
}

func DnssecConfigToMap(obj *oci_dns.DnssecConfig) map[string]interface{} {
	result := map[string]interface{}{}

	kskDnssecKeyVersions := []interface{}{}
	for _, item := range obj.KskDnssecKeyVersions {
		kskDnssecKeyVersions = append(kskDnssecKeyVersions, KskDnssecKeyVersionToMap(item))
	}
	result["ksk_dnssec_key_versions"] = kskDnssecKeyVersions

	zskDnssecKeyVersions := []interface{}{}
	for _, item := range obj.ZskDnssecKeyVersions {
		zskDnssecKeyVersions = append(zskDnssecKeyVersions, ZskDnssecKeyVersionToMap(item))
	}
	result["zsk_dnssec_key_versions"] = zskDnssecKeyVersions

	return result
}

func DnssecKeyVersionDsDataToMap(obj oci_dns.DnssecKeyVersionDsData) map[string]interface{} {
	result := map[string]interface{}{}

	result["digest_type"] = string(obj.DigestType)

	if obj.Rdata != nil {
		result["rdata"] = string(*obj.Rdata)
	}

	return result
}

func (s *DnsZoneResourceCrud) mapToExternalDownstream(fieldKeyFormat string) (oci_dns.ExternalDownstream, error) {
	result := oci_dns.ExternalDownstream{}

	if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
		tmp := address.(string)
		result.Address = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if tsigKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tsig_key_id")); ok {
		tmp := tsigKeyId.(string)
		result.TsigKeyId = &tmp
	}

	return result, nil
}

func ExternalDownstreamToMap(obj oci_dns.ExternalDownstream) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Address != nil {
		result["address"] = string(*obj.Address)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.TsigKeyId != nil {
		result["tsig_key_id"] = string(*obj.TsigKeyId)
	}

	return result
}

func (s *DnsZoneResourceCrud) mapToExternalMaster(fieldKeyFormat string) (oci_dns.ExternalMaster, error) {
	result := oci_dns.ExternalMaster{}

	if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
		tmp := address.(string)
		result.Address = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if tsigKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tsig_key_id")); ok {
		tmp := tsigKeyId.(string)
		result.TsigKeyId = &tmp
	}

	return result, nil
}

func ExternalMasterToMap(obj oci_dns.ExternalMaster) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Address != nil {
		result["address"] = string(*obj.Address)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.TsigKeyId != nil {
		result["tsig_key_id"] = string(*obj.TsigKeyId)
	}

	return result
}

func KskDnssecKeyVersionToMap(obj oci_dns.KskDnssecKeyVersion) map[string]interface{} {
	result := map[string]interface{}{}

	result["algorithm"] = string(obj.Algorithm)

	dsData := []interface{}{}
	for _, item := range obj.DsData {
		dsData = append(dsData, DnssecKeyVersionDsDataToMap(item))
	}
	result["ds_data"] = dsData

	if obj.KeyTag != nil {
		result["key_tag"] = int(*obj.KeyTag)
	}

	if obj.LengthInBytes != nil {
		result["length_in_bytes"] = int(*obj.LengthInBytes)
	}

	if obj.PredecessorDnssecKeyVersionUuid != nil {
		result["predecessor_dnssec_key_version_uuid"] = string(*obj.PredecessorDnssecKeyVersionUuid)
	}

	if obj.SuccessorDnssecKeyVersionUuid != nil {
		result["successor_dnssec_key_version_uuid"] = string(*obj.SuccessorDnssecKeyVersionUuid)
	}

	if obj.TimeActivated != nil {
		result["time_activated"] = obj.TimeActivated.String()
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeExpired != nil {
		result["time_expired"] = obj.TimeExpired.String()
	}

	if obj.TimeInactivated != nil {
		result["time_inactivated"] = obj.TimeInactivated.String()
	}

	if obj.TimePromoted != nil {
		result["time_promoted"] = obj.TimePromoted.String()
	}

	if obj.TimePublished != nil {
		result["time_published"] = obj.TimePublished.String()
	}

	if obj.TimeUnpublished != nil {
		result["time_unpublished"] = obj.TimeUnpublished.String()
	}

	if obj.Uuid != nil {
		result["uuid"] = string(*obj.Uuid)
	}

	return result
}

func NameserverToMap(obj oci_dns.Nameserver) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	return result
}

func ZoneTransferServerToMap(obj oci_dns.ZoneTransferServer) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Address != nil {
		result["address"] = string(*obj.Address)
	}

	if obj.IsTransferDestination != nil {
		result["is_transfer_destination"] = bool(*obj.IsTransferDestination)
	}

	if obj.IsTransferSource != nil {
		result["is_transfer_source"] = bool(*obj.IsTransferSource)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	return result
}

func ZskDnssecKeyVersionToMap(obj oci_dns.ZskDnssecKeyVersion) map[string]interface{} {
	result := map[string]interface{}{}

	result["algorithm"] = string(obj.Algorithm)

	if obj.KeyTag != nil {
		result["key_tag"] = int(*obj.KeyTag)
	}

	if obj.LengthInBytes != nil {
		result["length_in_bytes"] = int(*obj.LengthInBytes)
	}

	if obj.PredecessorDnssecKeyVersionUuid != nil {
		result["predecessor_dnssec_key_version_uuid"] = string(*obj.PredecessorDnssecKeyVersionUuid)
	}

	if obj.SuccessorDnssecKeyVersionUuid != nil {
		result["successor_dnssec_key_version_uuid"] = string(*obj.SuccessorDnssecKeyVersionUuid)
	}

	if obj.TimeActivated != nil {
		result["time_activated"] = obj.TimeActivated.String()
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeExpired != nil {
		result["time_expired"] = obj.TimeExpired.String()
	}

	if obj.TimeInactivated != nil {
		result["time_inactivated"] = obj.TimeInactivated.String()
	}

	if obj.TimePromoted != nil {
		result["time_promoted"] = obj.TimePromoted.String()
	}

	if obj.TimePublished != nil {
		result["time_published"] = obj.TimePublished.String()
	}

	if obj.TimeUnpublished != nil {
		result["time_unpublished"] = obj.TimeUnpublished.String()
	}

	if obj.Uuid != nil {
		result["uuid"] = string(*obj.Uuid)
	}

	return result
}

func (s *DnsZoneResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dns.ChangeZoneCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	if scope, ok := s.D.GetOkExists("scope"); ok {
		changeCompartmentRequest.Scope = oci_dns.ChangeZoneCompartmentScopeEnum(scope.(string))
	}

	idTmp := s.D.Id()
	changeCompartmentRequest.ZoneId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	_, err := s.Client.ChangeZoneCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func parseZoneCompositeId(compositeId string) (zoneNameOrId string, scope string, viewId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("zoneNameOrId/.*/scope/.*/viewId/.*", compositeId)

	if match && len(parts) == 6 {
		zoneNameOrId, _ = url.PathUnescape(parts[1])
		scope, _ = url.PathUnescape(parts[3])
		viewId, _ = url.PathUnescape(parts[5])
	} else {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}

	return
}
