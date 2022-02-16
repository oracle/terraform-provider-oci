// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_dns "github.com/oracle/oci-go-sdk/v58/dns"
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
				ForceNew: true,
			},
			"view_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// Computed
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
		createZoneDetailsRequest.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

	return nil
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

func NameserverToMap(obj oci_dns.Nameserver) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
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
