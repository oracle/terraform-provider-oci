// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	oci_functions "github.com/oracle/oci-go-sdk/v46/functions"
)

func init() {
	RegisterResource("oci_functions_application", FunctionsApplicationResource())
}

func FunctionsApplicationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createFunctionsApplication,
		Read:     readFunctionsApplication,
		Update:   updateFunctionsApplication,
		Delete:   deleteFunctionsApplication,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_ids": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"config": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"syslog_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trace_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"domain_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFunctionsApplication(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).functionsManagementClient()

	return CreateResource(d, sync)
}

func readFunctionsApplication(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).functionsManagementClient()

	return ReadResource(sync)
}

func updateFunctionsApplication(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).functionsManagementClient()

	return UpdateResource(d, sync)
}

func deleteFunctionsApplication(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).functionsManagementClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type FunctionsApplicationResourceCrud struct {
	BaseCrud
	Client                 *oci_functions.FunctionsManagementClient
	Res                    *oci_functions.Application
	DisableNotFoundRetries bool
}

func (s *FunctionsApplicationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FunctionsApplicationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_functions.ApplicationLifecycleStateCreating),
	}
}

func (s *FunctionsApplicationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_functions.ApplicationLifecycleStateActive),
	}
}

func (s *FunctionsApplicationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_functions.ApplicationLifecycleStateDeleting),
	}
}

func (s *FunctionsApplicationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_functions.ApplicationLifecycleStateDeleted),
	}
}

func (s *FunctionsApplicationResourceCrud) Create() error {
	request := oci_functions.CreateApplicationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if config, ok := s.D.GetOkExists("config"); ok {
		request.Config = objectMapToStringMap(config.(map[string]interface{}))
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

	if subnetIds, ok := s.D.GetOkExists("subnet_ids"); ok {
		interfaces := subnetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("subnet_ids") {
			request.SubnetIds = tmp
		}
	}

	if syslogUrl, ok := s.D.GetOkExists("syslog_url"); ok {
		tmp := syslogUrl.(string)
		request.SyslogUrl = &tmp
	}

	if traceConfig, ok := s.D.GetOkExists("trace_config"); ok {
		if tmpList := traceConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "trace_config", 0)
			tmp, err := s.mapToApplicationTraceConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TraceConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.CreateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *FunctionsApplicationResourceCrud) Get() error {
	request := oci_functions.GetApplicationRequest{}

	tmp := s.D.Id()
	request.ApplicationId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.GetApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *FunctionsApplicationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_functions.UpdateApplicationRequest{}

	tmp := s.D.Id()
	request.ApplicationId = &tmp

	if config, ok := s.D.GetOkExists("config"); ok {
		request.Config = objectMapToStringMap(config.(map[string]interface{}))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if syslogUrl, ok := s.D.GetOkExists("syslog_url"); ok {
		tmp := syslogUrl.(string)
		request.SyslogUrl = &tmp
	}

	if traceConfig, ok := s.D.GetOkExists("trace_config"); ok && s.D.HasChange("trace_config") {
		if tmpList := traceConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "trace_config", 0)
			tmp, err := s.mapToApplicationTraceConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TraceConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.UpdateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *FunctionsApplicationResourceCrud) Delete() error {
	request := oci_functions.DeleteApplicationRequest{}

	tmp := s.D.Id()
	request.ApplicationId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "functions")

	_, err := s.Client.DeleteApplication(context.Background(), request)
	return err
}

func (s *FunctionsApplicationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("config", s.Res.Config)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subnet_ids", s.Res.SubnetIds)

	if s.Res.SyslogUrl != nil {
		s.D.Set("syslog_url", *s.Res.SyslogUrl)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TraceConfig != nil {
		s.D.Set("trace_config", []interface{}{ApplicationTraceConfigToMap(s.Res.TraceConfig)})
	} else {
		s.D.Set("trace_config", nil)
	}

	return nil
}

func (s *FunctionsApplicationResourceCrud) mapToApplicationTraceConfig(fieldKeyFormat string) (oci_functions.ApplicationTraceConfig, error) {
	result := oci_functions.ApplicationTraceConfig{}

	if domainId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain_id")); ok {
		tmp := domainId.(string)
		result.DomainId = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func ApplicationTraceConfigToMap(obj *oci_functions.ApplicationTraceConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DomainId != nil {
		result["domain_id"] = string(*obj.DomainId)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func (s *FunctionsApplicationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_functions.ChangeApplicationCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.ApplicationId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "functions")

	_, err := s.Client.ChangeApplicationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := waitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *FunctionsApplicationResourceCrud) ExtraWaitPostDelete() time.Duration {
	if httpreplay.ShouldRetryImmediately() {
		return time.Duration(1 * time.Second)
	}
	log.Printf("[DEBUG] Waiting for 5 minutes post destroy of application resource due to known service issue")
	return time.Duration(5 * time.Minute)
}
