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

	oci_dns "github.com/oracle/oci-go-sdk/v58/dns"
)

func DnsViewResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDnsView,
		Read:     readDnsView,
		Update:   updateDnsView,
		Delete:   deleteDnsView,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scope": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
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

func createDnsView(d *schema.ResourceData, m interface{}) error {
	sync := &DnsViewResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.CreateResource(d, sync)
}

func readDnsView(d *schema.ResourceData, m interface{}) error {
	sync := &DnsViewResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

func updateDnsView(d *schema.ResourceData, m interface{}) error {
	sync := &DnsViewResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDnsView(d *schema.ResourceData, m interface{}) error {
	sync := &DnsViewResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DnsViewResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.View
	DisableNotFoundRetries bool
}

func (s *DnsViewResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DnsViewResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DnsViewResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dns.ViewLifecycleStateActive),
	}
}

func (s *DnsViewResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dns.ViewLifecycleStateDeleting),
	}
}

func (s *DnsViewResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dns.ViewLifecycleStateDeleted),
	}
}

func (s *DnsViewResourceCrud) Create() error {
	request := oci_dns.CreateViewRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.CreateViewScopeEnum(scope.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.CreateView(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.View
	return nil
}

func (s *DnsViewResourceCrud) Get() error {
	request := oci_dns.GetViewRequest{}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.GetViewScopeEnum(scope.(string))
	}

	scope, viewId, err := parseDnsViewCompositeId(s.D.Id())
	if err == nil {
		request.ViewId = &viewId
		s.D.SetId(viewId)
		request.Scope = oci_dns.GetViewScopeEnum(scope)
		s.D.Set("scope", scope)
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	tmp := s.D.Id()
	request.ViewId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.GetView(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.View
	return nil
}

func (s *DnsViewResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dns.UpdateViewRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.UpdateViewScopeEnum(scope.(string))
	}

	tmp := s.D.Id()
	request.ViewId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.UpdateView(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.View
	return nil
}

func (s *DnsViewResourceCrud) Delete() error {
	request := oci_dns.DeleteViewRequest{}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.DeleteViewScopeEnum(scope.(string))
	}

	tmp := s.D.Id()
	request.ViewId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	_, err := s.Client.DeleteView(context.Background(), request)
	return err
}

func (s *DnsViewResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsProtected != nil {
		s.D.Set("is_protected", *s.Res.IsProtected)
	}

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

func (s *DnsViewResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dns.ChangeViewCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	if scope, ok := s.D.GetOkExists("scope"); ok {
		changeCompartmentRequest.Scope = oci_dns.ChangeViewCompartmentScopeEnum(scope.(string))
	}

	idTmp := s.D.Id()
	changeCompartmentRequest.ViewId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	_, err := s.Client.ChangeViewCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func parseDnsViewCompositeId(compositeId string) (scope string, viewId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("viewId/.*/scope/.*", compositeId)

	if match && len(parts) == 4 {
		viewId, _ = url.PathUnescape(parts[1])
		scope, _ = url.PathUnescape(parts[3])
	} else {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}

	return
}
