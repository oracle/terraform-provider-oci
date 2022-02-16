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

func DnsResolverEndpointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDnsResolverEndpoint,
		Read:     readDnsResolverEndpoint,
		Update:   updateDnsResolverEndpoint,
		Delete:   deleteDnsResolverEndpoint,
		Schema: map[string]*schema.Schema{
			// Required
			"is_forwarding": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"is_listening": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resolver_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"endpoint_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"VNIC",
				}, true),
			},
			"forwarding_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"listening_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Set:      utils.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"compartment_id": {
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

func createDnsResolverEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DnsResolverEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.CreateResource(d, sync)
}

func readDnsResolverEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DnsResolverEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

func updateDnsResolverEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DnsResolverEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDnsResolverEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DnsResolverEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DnsResolverEndpointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.ResolverEndpoint
	DisableNotFoundRetries bool
}

func (s *DnsResolverEndpointResourceCrud) ID() string {
	return getResolverEndpointCompositeId(s.D.Get("name").(string), s.D.Get("resolver_id").(string))
}

func (s *DnsResolverEndpointResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dns.ResolverEndpointLifecycleStateCreating),
	}
}

func (s *DnsResolverEndpointResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dns.ResolverEndpointLifecycleStateActive),
	}
}

func (s *DnsResolverEndpointResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dns.ResolverEndpointLifecycleStateDeleting),
	}
}

func (s *DnsResolverEndpointResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dns.ResolverEndpointLifecycleStateDeleted),
	}
}

func (s *DnsResolverEndpointResourceCrud) Create() error {
	request := oci_dns.CreateResolverEndpointRequest{}
	err := s.populateTopLevelPolymorphicCreateResolverEndpointRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.CreateResolverEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ResolverEndpoint
	return nil
}

func (s *DnsResolverEndpointResourceCrud) Get() error {
	request := oci_dns.GetResolverEndpointRequest{}

	if resolverEndpointName, ok := s.D.GetOkExists("name"); ok {
		tmp := resolverEndpointName.(string)
		request.ResolverEndpointName = &tmp
	}

	if resolverId, ok := s.D.GetOkExists("resolver_id"); ok {
		tmp := resolverId.(string)
		request.ResolverId = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.GetResolverEndpointScopeEnum(scope.(string))
	}

	resolverEndpointName, resolverId, scope, err := parseResolverEndpointCompositeId(s.D.Id())
	if err == nil {
		request.ResolverEndpointName = &resolverEndpointName
		request.ResolverId = &resolverId
		if scope != "" {
			request.Scope = oci_dns.GetResolverEndpointScopeEnum(scope)
		}
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.GetResolverEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ResolverEndpoint
	return nil
}

func (s *DnsResolverEndpointResourceCrud) Update() error {
	request := oci_dns.UpdateResolverEndpointRequest{}
	err := s.populateTopLevelPolymorphicUpdateResolverEndpointRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.UpdateResolverEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ResolverEndpoint
	return nil
}

func (s *DnsResolverEndpointResourceCrud) Delete() error {
	request := oci_dns.DeleteResolverEndpointRequest{}

	if resolverEndpointName, ok := s.D.GetOkExists("name"); ok {
		tmp := resolverEndpointName.(string)
		request.ResolverEndpointName = &tmp
	}

	if resolverId, ok := s.D.GetOkExists("resolver_id"); ok {
		tmp := resolverId.(string)
		request.ResolverId = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.DeleteResolverEndpointScopeEnum(scope.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	_, err := s.Client.DeleteResolverEndpoint(context.Background(), request)
	return err
}

func (s *DnsResolverEndpointResourceCrud) SetData() error {

	resolverEndpointName, resolverId, scope, err := parseResolverEndpointCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("name", &resolverEndpointName)
		s.D.Set("resolver_id", &resolverId)
		s.D.SetId(getResolverEndpointCompositeId(resolverEndpointName, resolverId))
		if scope != "" {
			s.D.Set("scope", scope)
		}
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_dns.ResolverVnicEndpoint:
		s.D.Set("endpoint_type", "VNIC")

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", schema.NewSet(utils.LiteralTypeHashCodeForSets, nsgIds))

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.ForwardingAddress != nil {
			s.D.Set("forwarding_address", *v.ForwardingAddress)
		}

		if v.IsForwarding != nil {
			s.D.Set("is_forwarding", *v.IsForwarding)
		}

		if v.IsListening != nil {
			s.D.Set("is_listening", *v.IsListening)
		}

		if v.ListeningAddress != nil {
			s.D.Set("listening_address", *v.ListeningAddress)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.Self != nil {
			s.D.Set("self", *v.Self)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'endpoint_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func getResolverEndpointCompositeId(resolverEndpointName string, resolverId string) string {
	resolverEndpointName = url.PathEscape(resolverEndpointName)
	resolverId = url.PathEscape(resolverId)
	compositeId := "resolverId/" + resolverId + "/name/" + resolverEndpointName
	return compositeId
}

func parseResolverEndpointCompositeId(compositeId string) (resolverEndpointName string, resolverId string, scope string, err error) {
	parts := strings.Split(compositeId, "/")
	match1, _ := regexp.MatchString("resolverId/.*/name/.*", compositeId)
	match2, _ := regexp.MatchString("resolverId/.*/name/.*/scope/.*", compositeId)
	if match1 && len(parts) == 4 {
		resolverId, _ = url.PathUnescape(parts[1])
		resolverEndpointName, _ = url.PathUnescape(parts[3])
	} else if match2 && len(parts) == 6 {
		resolverId, _ = url.PathUnescape(parts[1])
		resolverEndpointName, _ = url.PathUnescape(parts[3])
		scope, _ = url.PathUnescape(parts[5])
	} else {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}

	return
}

func (s *DnsResolverEndpointResourceCrud) populateTopLevelPolymorphicCreateResolverEndpointRequest(request *oci_dns.CreateResolverEndpointRequest) error {
	//discriminator
	endpointTypeRaw, ok := s.D.GetOkExists("endpoint_type")
	var endpointType string
	if ok {
		endpointType = endpointTypeRaw.(string)
	} else {
		endpointType = "VNIC" // default value
	}
	switch strings.ToLower(endpointType) {
	case strings.ToLower("VNIC"):
		details := oci_dns.CreateResolverVnicEndpointDetails{}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if forwardingAddress, ok := s.D.GetOkExists("forwarding_address"); ok {
			tmp := forwardingAddress.(string)
			details.ForwardingAddress = &tmp
		}
		if isForwarding, ok := s.D.GetOkExists("is_forwarding"); ok {
			tmp := isForwarding.(bool)
			details.IsForwarding = &tmp
		}
		if isListening, ok := s.D.GetOkExists("is_listening"); ok {
			tmp := isListening.(bool)
			details.IsListening = &tmp
		}
		if listeningAddress, ok := s.D.GetOkExists("listening_address"); ok {
			tmp := listeningAddress.(string)
			details.ListeningAddress = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if resolverId, ok := s.D.GetOkExists("resolver_id"); ok {
			tmp := resolverId.(string)
			request.ResolverId = &tmp
		}
		if scope, ok := s.D.GetOkExists("scope"); ok {
			request.Scope = oci_dns.CreateResolverEndpointScopeEnum(scope.(string))
		}
		request.CreateResolverEndpointDetails = details
	default:
		return fmt.Errorf("unknown endpoint_type '%v' was specified", endpointType)
	}
	return nil
}

func (s *DnsResolverEndpointResourceCrud) populateTopLevelPolymorphicUpdateResolverEndpointRequest(request *oci_dns.UpdateResolverEndpointRequest) error {
	//discriminator
	endpointTypeRaw, ok := s.D.GetOkExists("endpoint_type")
	var endpointType string
	if ok {
		endpointType = endpointTypeRaw.(string)
	} else {
		endpointType = "VNIC" // default value
	}
	switch strings.ToLower(endpointType) {
	case strings.ToLower("VNIC"):
		details := oci_dns.UpdateResolverVnicEndpointDetails{}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if resolverEndpointName, ok := s.D.GetOkExists("name"); ok {
			tmp := resolverEndpointName.(string)
			request.ResolverEndpointName = &tmp
		}
		if resolverId, ok := s.D.GetOkExists("resolver_id"); ok {
			tmp := resolverId.(string)
			request.ResolverId = &tmp
		}
		if scope, ok := s.D.GetOkExists("scope"); ok {
			request.Scope = oci_dns.UpdateResolverEndpointScopeEnum(scope.(string))
		}
		request.UpdateResolverEndpointDetails = details
	default:
		return fmt.Errorf("unknown endpoint_type '%v' was specified", endpointType)
	}
	return nil
}
