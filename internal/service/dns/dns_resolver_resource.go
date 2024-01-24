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

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_dns "github.com/oracle/oci-go-sdk/v65/dns"
)

func DnsResolverResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDnsResolver,
		Read:     readDnsResolver,
		Update:   updateDnsResolver,
		Delete:   deleteDnsResolver,
		Schema: map[string]*schema.Schema{
			// Required
			"resolver_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"attached_views": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"view_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"action": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"FORWARD",
							}, true),
						},
						"destination_addresses": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"source_endpoint_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"client_address_conditions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"qname_cover_conditions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// Computed
			"attached_vcn_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_view_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"endpoint_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"forwarding_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_forwarding": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_listening": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"listening_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"self": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_protected": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"self": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDnsResolver(d *schema.ResourceData, m interface{}) error {
	sync := &DnsResolverResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()
	compartment, ok := sync.D.GetOkExists("compartment_id")

	err := tfresource.CreateResource(d, sync)
	if err != nil {
		return err
	}

	if ok && compartment != *sync.Res.CompartmentId {
		err = sync.updateCompartment(compartment)
		if err != nil {
			return err
		}
		tmp := compartment.(string)
		sync.Res.CompartmentId = &tmp
		err := sync.Get()
		if err != nil {
			log.Printf("error doing a Get() after compartment Update: %v", err)
		}
		err = sync.SetData()
		if err != nil {
			log.Printf("error doing a SetData() after compartment Update: %v", err)
		}
	}
	return nil
}

func readDnsResolver(d *schema.ResourceData, m interface{}) error {
	sync := &DnsResolverResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

func updateDnsResolver(d *schema.ResourceData, m interface{}) error {
	sync := &DnsResolverResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDnsResolver(d *schema.ResourceData, m interface{}) error {
	sync := &DnsResolverResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()
	// This will call the Delete function which internally calls the updateResolver dns API operation in order to clear properties of the resolver.
	// DeleteResolver is not a public facing dns operation. Resolvers are deleted when their corresponding VCN is deleted.
	return sync.Delete()
}

type DnsResolverResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.Resolver
	DisableNotFoundRetries bool
}

func (s *DnsResolverResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DnsResolverResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dns.ResolverLifecycleStateCreating),
	}
}

func (s *DnsResolverResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dns.ResolverLifecycleStateActive),
	}
}

func (s *DnsResolverResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_dns.ResolverLifecycleStateUpdating),
	}
}

func (s *DnsResolverResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_dns.ResolverLifecycleStateActive),
	}
}

func (s *DnsResolverResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dns.ResolverLifecycleStateDeleting),
	}
}

func (s *DnsResolverResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dns.ResolverLifecycleStateDeleted),
	}
}

func (s *DnsResolverResourceCrud) Create() error {
	request := oci_dns.UpdateResolverRequest{}

	if attachedViews, ok := s.D.GetOkExists("attached_views"); ok {
		interfaces := attachedViews.([]interface{})
		tmp := make([]oci_dns.AttachedViewDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "attached_views", stateDataIndex)
			converted, err := s.mapToAttachedViewDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("attached_views") {
			request.AttachedViews = tmp
		}
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if resolverId, ok := s.D.GetOkExists("resolver_id"); ok {
		tmp := resolverId.(string)
		request.ResolverId = &tmp
	}

	if rules, ok := s.D.GetOkExists("rules"); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_dns.ResolverRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
			converted, err := s.mapToResolverRuleDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("rules") {
			request.Rules = tmp
		}
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.UpdateResolverScopeEnum(scope.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.UpdateResolver(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Resolver
	s.D.SetId(s.ID())

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DnsResolverResourceCrud) Get() error {
	request := oci_dns.GetResolverRequest{}

	tmp := s.D.Id()
	request.ResolverId = &tmp

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.GetResolverScopeEnum(scope.(string))
	}

	resolverId, scope, err := parseDnsResolverCompositeId(s.D.Id())
	if err == nil {
		request.ResolverId = &resolverId
		s.D.SetId(resolverId)
		s.D.Set("resolver_id", resolverId)
		request.Scope = oci_dns.GetResolverScopeEnum(scope)
		s.D.Set("scope", scope)
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.GetResolver(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Resolver
	return nil
}

func (s *DnsResolverResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dns.UpdateResolverRequest{}

	if attachedViews, ok := s.D.GetOkExists("attached_views"); ok {
		interfaces := attachedViews.([]interface{})
		tmp := make([]oci_dns.AttachedViewDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "attached_views", stateDataIndex)
			converted, err := s.mapToAttachedViewDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("attached_views") {
			request.AttachedViews = tmp
		}
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ResolverId = &tmp

	if rules, ok := s.D.GetOkExists("rules"); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_dns.ResolverRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
			converted, err := s.mapToResolverRuleDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("rules") {
			request.Rules = tmp
		}
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.UpdateResolverScopeEnum(scope.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.UpdateResolver(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Resolver

	return nil
}

// The Delete function below checks if existing resolver has attachedViews, rules, defined tags and/or freeform tags. If any of these property is present then
// the function calls dns updateResolver API and drops the AttachedViews and Rules and clears defined tags and freeform tags.
func (s *DnsResolverResourceCrud) Delete() error {
	var hasAttachedViews, hasRules, hasDefinedtags, hasFreeformTags bool

	// ...check if the existing resolver has attachedViews
	if attachedViews, ok := s.D.GetOkExists("attached_views"); ok {
		interfaces := attachedViews.([]interface{})
		if len(interfaces) != 0 {
			hasAttachedViews = true
		}
	}

	// ...check if the existing resolver has rules
	if rules, ok := s.D.GetOkExists("rules"); ok {
		interfaces := rules.([]interface{})
		if len(interfaces) != 0 {
			hasRules = true
		}
	}

	// ...check if the existing resolver has defined tags
	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		if len(definedTags.(map[string]interface{})) != 0 {
			hasDefinedtags = true
		}
	}

	// ...check if the existing resolver has freeform tags
	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		if len(freeformTags.(map[string]interface{})) != 0 {
			hasFreeformTags = true
		}
	}

	// ...if any of the above property is present then make a call to updateResolver dns API
	if hasAttachedViews || hasRules || hasDefinedtags || hasFreeformTags {
		request := oci_dns.UpdateResolverRequest{}

		// Setting AttachedViews to an empty list
		tmpViewList := make([]oci_dns.AttachedViewDetails, 0)
		request.AttachedViews = tmpViewList

		// Setting Rules to an empty list
		tmpRulesList := make([]oci_dns.ResolverRuleDetails, 0)
		request.Rules = tmpRulesList

		// Setting defined tags as empty
		definedTags := make(map[string]map[string]interface{})
		request.DefinedTags = definedTags

		// Setting defined tags as empty
		freeformTags := map[string]string{}
		request.FreeformTags = freeformTags

		tmp := s.D.Id()
		request.ResolverId = &tmp

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

		response, err := s.Client.UpdateResolver(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res = &response.Resolver
		return nil
	}
	return nil
}

func (s *DnsResolverResourceCrud) SetData() error {
	if s.Res.Id != nil {
		s.D.Set("resolver_id", *s.Res.Id)
	}

	if s.Res.AttachedVcnId != nil {
		s.D.Set("attached_vcn_id", *s.Res.AttachedVcnId)
	}

	attachedViews := []interface{}{}
	for _, item := range s.Res.AttachedViews {
		attachedViews = append(attachedViews, AttachedViewToMap(item))
	}
	s.D.Set("attached_views", attachedViews)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefaultViewId != nil {
		s.D.Set("default_view_id", *s.Res.DefaultViewId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	endpoints := []interface{}{}
	for _, item := range s.Res.Endpoints {
		endpoints = append(endpoints, ResolverEndpointSummaryToMap(item))
	}
	s.D.Set("endpoints", endpoints)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsProtected != nil {
		s.D.Set("is_protected", *s.Res.IsProtected)
	}

	rules := []interface{}{}
	for _, item := range s.Res.Rules {
		rules = append(rules, ResolverRuleToMap(item))
	}
	s.D.Set("rules", rules)

	if s.Res.Self != nil {
		s.D.Set("self", *s.Res.Self)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DnsResolverResourceCrud) mapToAttachedViewDetails(fieldKeyFormat string) (oci_dns.AttachedViewDetails, error) {
	result := oci_dns.AttachedViewDetails{}

	if viewId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "view_id")); ok {
		tmp := viewId.(string)
		result.ViewId = &tmp
	}

	return result, nil
}

func AttachedViewToMap(obj oci_dns.AttachedView) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ViewId != nil {
		result["view_id"] = string(*obj.ViewId)
	}

	return result
}

func ResolverEndpointSummaryToMap(obj oci_dns.ResolverEndpointSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_dns.ResolverVnicEndpointSummary:
		result["endpoint_type"] = "VNIC"

		if v.SubnetId != nil {
			result["subnet_id"] = string(*v.SubnetId)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.ForwardingAddress != nil {
			result["forwarding_address"] = string(*v.ForwardingAddress)
		}

		if v.IsForwarding != nil {
			result["is_forwarding"] = bool(*v.IsForwarding)
		}

		if v.IsListening != nil {
			result["is_listening"] = bool(*v.IsListening)
		}

		if v.ListeningAddress != nil {
			result["listening_address"] = string(*v.ListeningAddress)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.Self != nil {
			result["self"] = string(*v.Self)
		}

		result["state"] = string(v.LifecycleState)

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}
	default:
		log.Printf("[WARN] Received 'endpoint_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DnsResolverResourceCrud) mapToResolverRuleDetails(fieldKeyFormat string) (oci_dns.ResolverRuleDetails, error) {
	var baseObject oci_dns.ResolverRuleDetails
	//discriminator
	actionRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action"))
	var action string
	if ok {
		action = actionRaw.(string)
	} else {
		action = "" // default value
	}
	switch strings.ToLower(action) {
	case strings.ToLower("FORWARD"):
		details := oci_dns.ResolverForwardRuleDetails{}
		if destinationAddresses, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_addresses")); ok {
			interfaces := destinationAddresses.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "destination_addresses")) {
				details.DestinationAddresses = tmp
			}
		}
		if sourceEndpointName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_endpoint_name")); ok {
			tmp := sourceEndpointName.(string)
			details.SourceEndpointName = &tmp
		}
		if clientAddressConditions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_address_conditions")); ok {
			interfaces := clientAddressConditions.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "client_address_conditions")) {
				details.ClientAddressConditions = tmp
			}
		}
		if qnameCoverConditions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "qname_cover_conditions")); ok {
			interfaces := qnameCoverConditions.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "qname_cover_conditions")) {
				details.QnameCoverConditions = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown action '%v' was specified", action)
	}
	return baseObject, nil
}

func ResolverRuleToMap(obj oci_dns.ResolverRule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_dns.ResolverForwardRule:
		result["action"] = "FORWARD"

		result["destination_addresses"] = v.DestinationAddresses

		if v.SourceEndpointName != nil {
			result["source_endpoint_name"] = string(*v.SourceEndpointName)
		}

		result["client_address_conditions"] = v.ClientAddressConditions

		result["qname_cover_conditions"] = v.QnameCoverConditions
	default:
		log.Printf("[WARN] Received 'action' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DnsResolverResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dns.ChangeResolverCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ResolverId = &idTmp

	if scope, ok := s.D.GetOkExists("scope"); ok {
		changeCompartmentRequest.Scope = oci_dns.ChangeResolverCompartmentScopeEnum(scope.(string))
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	_, err := s.Client.ChangeResolverCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func parseDnsResolverCompositeId(compositeId string) (resolverId string, scope string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("resolverId/.*/scope/.*", compositeId)

	if match && len(parts) == 4 {
		resolverId, _ = url.PathUnescape(parts[1])
		scope, _ = url.PathUnescape(parts[3])
	} else {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}

	return
}
