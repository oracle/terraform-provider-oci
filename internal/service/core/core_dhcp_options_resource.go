// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreDhcpOptionsResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreDhcpOptions,
		Read:     readCoreDhcpOptions,
		Update:   updateCoreDhcpOptions,
		Delete:   deleteCoreDhcpOptions,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"options": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      optionsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DomainNameServer",
								"SearchDomain",
							}, true),
						},

						// Optional
						"custom_dns_servers": {
							Type:     schema.TypeList,
							Optional: true,
							// remove `computed` because it prevents unsetting the `custom_dns_servers` list when
							// changing dhcp_options-> options-> server_type from `CustomDnsServer` to `VcnLocalPlusInternet`
							//Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"search_domain_names": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"server_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(oci_core.DhcpDnsOptionServerTypeCustomdnsserver),
								string(oci_core.DhcpDnsOptionServerTypeVcnlocal),
								string(oci_core.DhcpDnsOptionServerTypeVcnlocalplusinternet),
							}, false),
						},

						// Computed
					},
				},
			},
			"vcn_id": {
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name_type": {
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

func createCoreDhcpOptions(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDhcpOptionsResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreDhcpOptions(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDhcpOptionsResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreDhcpOptions(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDhcpOptionsResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreDhcpOptions(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDhcpOptionsResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreDhcpOptionsResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.DhcpOptions
	DisableNotFoundRetries bool
}

func (s *CoreDhcpOptionsResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreDhcpOptionsResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.DhcpOptionsLifecycleStateProvisioning),
	}
}

func (s *CoreDhcpOptionsResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.DhcpOptionsLifecycleStateAvailable),
	}
}

func (s *CoreDhcpOptionsResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.DhcpOptionsLifecycleStateTerminating),
	}
}

func (s *CoreDhcpOptionsResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.DhcpOptionsLifecycleStateTerminated),
	}
}

func (s *CoreDhcpOptionsResourceCrud) Create() error {
	request := oci_core.CreateDhcpOptionsRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if domainNameType, ok := s.D.GetOkExists("domain_name_type"); ok {
		request.DomainNameType = oci_core.CreateDhcpDetailsDomainNameTypeEnum(domainNameType.(string))
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if options, ok := s.D.GetOkExists("options"); ok {
		set := options.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.DhcpOption, len(interfaces))
		for i := range interfaces {
			stateDataIndex := optionsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "options", stateDataIndex)
			converted, err := s.mapToDhcpOption(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("options") {
			request.Options = tmp
		}
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateDhcpOptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DhcpOptions

	// this is needed to make the infrastructure match what is on the config as in some cases during the Create the service adds an option for SearchDomain by default if the user doesn't provide it.
	updateRequest := oci_core.UpdateDhcpOptionsRequest{}
	updateRequest.DhcpId = s.Res.Id
	updateRequest.Options = request.Options
	updateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
	updateResponse, err := s.Client.UpdateDhcpOptions(context.Background(), updateRequest)
	if err != nil {
		log.Printf("[ERROR] Could not perform an Update right after the Create of the dhcpOptions: %v", err)
	}
	s.Res = &updateResponse.DhcpOptions

	return nil
}

func (s *CoreDhcpOptionsResourceCrud) Get() error {
	request := oci_core.GetDhcpOptionsRequest{}

	tmp := s.D.Id()
	request.DhcpId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetDhcpOptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DhcpOptions
	return nil
}

func (s *CoreDhcpOptionsResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateDhcpOptionsRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.DhcpId = &tmp

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if domainNameType, ok := s.D.GetOkExists("domain_name_type"); ok {
		request.DomainNameType = oci_core.UpdateDhcpDetailsDomainNameTypeEnum(domainNameType.(string))
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if options, ok := s.D.GetOkExists("options"); ok {
		set := options.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.DhcpOption, len(interfaces))
		for i := range interfaces {
			stateDataIndex := optionsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "options", stateDataIndex)
			converted, err := s.mapToDhcpOption(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("options") {
			request.Options = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDhcpOptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DhcpOptions
	return nil
}

func (s *CoreDhcpOptionsResourceCrud) Delete() error {
	request := oci_core.DeleteDhcpOptionsRequest{}

	tmp := s.D.Id()
	request.DhcpId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteDhcpOptions(context.Background(), request)
	return err
}

func (s *CoreDhcpOptionsResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("domain_name_type", s.Res.DomainNameType)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	options := []interface{}{}
	for _, item := range s.Res.Options {
		options = append(options, DhcpOptionToMap(item))
	}
	s.D.Set("options", schema.NewSet(optionsHashCodeForSets, options))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func (s *CoreDhcpOptionsResourceCrud) mapToDhcpOption(fieldKeyFormat string) (oci_core.DhcpOption, error) {
	var baseObject oci_core.DhcpOption
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DomainNameServer"):
		details := oci_core.DhcpDnsOption{}
		details.CustomDnsServers = []string{}
		if customDnsServers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_dns_servers")); ok {
			interfaces := customDnsServers.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.CustomDnsServers = tmp

		}
		if serverType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "server_type")); ok {
			details.ServerType = oci_core.DhcpDnsOptionServerTypeEnum(serverType.(string))
		}
		baseObject = details
	case strings.ToLower("SearchDomain"):
		details := oci_core.DhcpSearchDomainOption{}
		details.SearchDomainNames = []string{}
		if searchDomainNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "search_domain_names")); ok {
			interfaces := searchDomainNames.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.SearchDomainNames = tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func DhcpOptionToMap(obj oci_core.DhcpOption) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_core.DhcpDnsOption:
		result["type"] = "DomainNameServer"

		result["custom_dns_servers"] = v.CustomDnsServers

		result["server_type"] = string(v.ServerType)
	case oci_core.DhcpSearchDomainOption:
		result["type"] = "SearchDomain"

		result["search_domain_names"] = v.SearchDomainNames
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func optionsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if customDnsServers, ok := m["custom_dns_servers"]; ok && customDnsServers != "" {
		if tmpList, ok := customDnsServers.([]interface{}); ok && len(tmpList) > 0 && tmpList[0] != "" {
			buf.WriteString("custom_dns_servers-")
			for _, customDnsServer := range tmpList {
				buf.WriteString(fmt.Sprintf("%v-", customDnsServer))
			}
		}
	}
	if searchDomainNames, ok := m["search_domain_names"]; ok && searchDomainNames != "" {
		if tmpList, ok := searchDomainNames.([]interface{}); ok && len(tmpList) > 0 && tmpList[0] != "" {
			buf.WriteString("search_domain_names-")
			for _, searchDomainName := range tmpList {
				buf.WriteString(fmt.Sprintf("%v-", searchDomainName))
			}
		}
	}
	if serverType, ok := m["server_type"]; ok && serverType != "" {
		buf.WriteString(fmt.Sprintf("%v-", serverType))
	}
	if type_, ok := m["type"]; ok && type_ != "" {
		buf.WriteString(fmt.Sprintf("%v-", strings.ToLower(type_.(string))))
	}
	return hashcode.String(buf.String())
}
func (s *CoreDhcpOptionsResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeDhcpOptionsCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DhcpId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeDhcpOptionsCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
