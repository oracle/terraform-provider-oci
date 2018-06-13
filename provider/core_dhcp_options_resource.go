// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	"github.com/hashicorp/terraform/helper/validation"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	DhcpOptionTypeDomainNameServer = "DomainNameServer"
	DhcpOptionTypeSearchDomain     = "SearchDomain"
)

func DhcpOptionsResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createDhcpOptions,
		Read:     readDhcpOptions,
		Update:   updateDhcpOptions,
		Delete:   deleteDhcpOptions,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"options": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					// Polymorphic type.
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: crud.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								DhcpOptionTypeDomainNameServer,
								DhcpOptionTypeSearchDomain,
							}, true),
						},

						// Optional
						"custom_dns_servers": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"server_type": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(oci_core.DhcpDnsOptionServerTypeCustomdnsserver),
								string(oci_core.DhcpDnsOptionServerTypeVcnlocal),
								string(oci_core.DhcpDnsOptionServerTypeVcnlocalplusinternet),
							}, false),
						},
						"search_domain_names": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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

func createDhcpOptions(d *schema.ResourceData, m interface{}) error {
	sync := &DhcpOptionsResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readDhcpOptions(d *schema.ResourceData, m interface{}) error {
	sync := &DhcpOptionsResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateDhcpOptions(d *schema.ResourceData, m interface{}) error {
	sync := &DhcpOptionsResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteDhcpOptions(d *schema.ResourceData, m interface{}) error {
	sync := &DhcpOptionsResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type DhcpOptionsResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.DhcpOptions
	DisableNotFoundRetries bool
}

func (s *DhcpOptionsResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DhcpOptionsResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.DhcpOptionsLifecycleStateProvisioning),
	}
}

func (s *DhcpOptionsResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.DhcpOptionsLifecycleStateAvailable),
	}
}

func (s *DhcpOptionsResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.DhcpOptionsLifecycleStateTerminating),
	}
}

func (s *DhcpOptionsResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.DhcpOptionsLifecycleStateTerminated),
	}
}

func (s *DhcpOptionsResourceCrud) Create() error {
	request := oci_core.CreateDhcpOptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.Options = []oci_core.DhcpOption{}
	if options, ok := s.D.GetOkExists("options"); ok {
		interfaces := options.([]interface{})
		tmp := make([]oci_core.DhcpOption, len(interfaces))
		for i, toBeConverted := range interfaces {
			var conversionError error
			tmp[i], conversionError = mapToDhcpOption(toBeConverted.(map[string]interface{}))
			if conversionError != nil {
				return conversionError
			}
		}
		request.Options = tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateDhcpOptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DhcpOptions
	return nil
}

func (s *DhcpOptionsResourceCrud) Get() error {
	request := oci_core.GetDhcpOptionsRequest{}

	tmp := s.D.Id()
	request.DhcpId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetDhcpOptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DhcpOptions
	return nil
}

func (s *DhcpOptionsResourceCrud) Update() error {
	request := oci_core.UpdateDhcpOptionsRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.Options = []oci_core.DhcpOption{}
	if options, ok := s.D.GetOkExists("options"); ok {
		interfaces := options.([]interface{})
		tmp := make([]oci_core.DhcpOption, len(interfaces))
		for i, toBeConverted := range interfaces {
			var conversionError error
			tmp[i], conversionError = mapToDhcpOption(toBeConverted.(map[string]interface{}))
			if conversionError != nil {
				return conversionError
			}
		}
		request.Options = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDhcpOptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DhcpOptions
	return nil
}

func (s *DhcpOptionsResourceCrud) Delete() error {
	request := oci_core.DeleteDhcpOptionsRequest{}

	tmp := s.D.Id()
	request.DhcpId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteDhcpOptions(context.Background(), request)
	return err
}

func (s *DhcpOptionsResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	options := []interface{}{}
	for _, item := range s.Res.Options {
		options = append(options, DhcpOptionToMap(item))
	}
	s.D.Set("options", options)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

}

func mapToDhcpOption(raw map[string]interface{}) (oci_core.DhcpOption, error) {
	dhcpOptionType := ""
	if type_, ok := raw["type"]; ok {
		dhcpOptionType = strings.ToLower(type_.(string))
	}

	searchDomainNames := []string{}
	for _, toBeConverted := range raw["search_domain_names"].([]interface{}) {
		searchDomainNames = append(searchDomainNames, toBeConverted.(string))
	}

	customDnsServers := []string{}
	for _, toBeConverted := range raw["custom_dns_servers"].([]interface{}) {
		customDnsServers = append(customDnsServers, toBeConverted.(string))
	}

	serverType := oci_core.DhcpDnsOptionServerTypeEnum(raw["server_type"].(string))

	// TODO: Ideally, the following validations for invalid fields based on polymorphic type should be done through a ValidateFunc
	// for the 'options' field. Terraform doesn't currently support ValidateFunc for TypeList fields, so do it here instead.
	// Move this to the schema once it's supported.
	var result oci_core.DhcpOption
	if dhcpOptionType == strings.ToLower(DhcpOptionTypeDomainNameServer) {
		if len(searchDomainNames) > 0 {
			return nil, fmt.Errorf("'search_domain_names' should not be specified for type %s", DhcpOptionTypeDomainNameServer)
		}

		result = oci_core.DhcpDnsOption{
			CustomDnsServers: customDnsServers,
			ServerType:       serverType,
		}
	} else if dhcpOptionType == strings.ToLower(DhcpOptionTypeSearchDomain) {
		if len(customDnsServers) > 0 {
			return nil, fmt.Errorf("'custom_dns_servers' should not be specified for type %s", DhcpOptionTypeSearchDomain)
		}

		if len(serverType) > 0 {
			return nil, fmt.Errorf("'server_type' should not be specified for type %s", DhcpOptionTypeSearchDomain)
		}

		result = oci_core.DhcpSearchDomainOption{
			SearchDomainNames: searchDomainNames,
		}
	} else {
		return nil, fmt.Errorf("Unknown Dhcp option type: %s", dhcpOptionType)
	}

	return result, nil
}

func DhcpOptionToMap(obj oci_core.DhcpOption) map[string]interface{} {
	result := map[string]interface{}{}

	if dhcpDnsOption, ok := obj.(oci_core.DhcpDnsOption); ok {
		result["type"] = DhcpOptionTypeDomainNameServer
		result["custom_dns_servers"] = dhcpDnsOption.CustomDnsServers
		result["server_type"] = string(dhcpDnsOption.ServerType)
	} else if dhcpSearchDomainOption, ok := obj.(oci_core.DhcpSearchDomainOption); ok {
		result["type"] = DhcpOptionTypeSearchDomain
		result["search_domain_names"] = dhcpSearchDomainOption.SearchDomainNames
	} else {
		return nil
	}

	return result
}
