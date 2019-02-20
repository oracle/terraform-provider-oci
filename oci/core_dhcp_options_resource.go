// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/hashicorp/terraform/helper/validation"

	"fmt"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CoreDhcpOptionsResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreDhcpOptions,
		Read:     readCoreDhcpOptions,
		Update:   updateCoreDhcpOptions,
		Delete:   deleteCoreDhcpOptions,
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
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DomainNameServer",
								"SearchDomain",
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
						"search_domain_names": {
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
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return CreateResource(d, sync)
}

func readCoreDhcpOptions(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDhcpOptionsResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

func updateCoreDhcpOptions(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDhcpOptionsResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return UpdateResource(d, sync)
}

func deleteCoreDhcpOptions(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDhcpOptionsResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreDhcpOptionsResourceCrud struct {
	BaseCrud
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
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "options", stateDataIndex)
			converted, err := s.mapToDhcpOption(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
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

func (s *CoreDhcpOptionsResourceCrud) Get() error {
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

func (s *CoreDhcpOptionsResourceCrud) Update() error {
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
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "options", stateDataIndex)
			converted, err := s.mapToDhcpOption(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
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

func (s *CoreDhcpOptionsResourceCrud) Delete() error {
	request := oci_core.DeleteDhcpOptionsRequest{}

	tmp := s.D.Id()
	request.DhcpId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteDhcpOptions(context.Background(), request)
	return err
}

func (s *CoreDhcpOptionsResourceCrud) SetData() error {
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
		if searchDomainNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "search_domain_names")); ok && len(searchDomainNames.([]interface{})) > 0 {
			return nil, fmt.Errorf("'search_domain_names' should not be specified for type DomainNameServer")
		}
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
		if customDnsServers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_dns_servers")); ok && len(customDnsServers.([]interface{})) > 0 {
			return nil, fmt.Errorf("'custom_dns_servers' should not be specified for type SearchDomain")
		}

		if serverType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "server_type")); ok && len(serverType.(string)) > 0 {
			return nil, fmt.Errorf("'server_type' should not be specified for type SearchDomain")
		}

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
