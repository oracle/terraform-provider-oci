// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	oci_functions "github.com/oracle/oci-go-sdk/functions"
)

func FunctionsFunctionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createFunctionsFunction,
		Read:     readFunctionsFunction,
		Update:   updateFunctionsFunction,
		Delete:   deleteFunctionsFunction,
		Schema: map[string]*schema.Schema{
			// Required
			"application_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image": {
				Type:     schema.TypeString,
				Required: true,
			},
			"memory_in_mbs": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     validateInt64TypeString,
				DiffSuppressFunc: int64StringDiffSuppressFunction,
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
			"image_digest": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"invoke_endpoint": {
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

func createFunctionsFunction(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsFunctionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).functionsManagementClient

	return CreateResource(d, sync)
}

func readFunctionsFunction(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsFunctionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).functionsManagementClient

	return ReadResource(sync)
}

func updateFunctionsFunction(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsFunctionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).functionsManagementClient

	return UpdateResource(d, sync)
}

func deleteFunctionsFunction(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsFunctionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).functionsManagementClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type FunctionsFunctionResourceCrud struct {
	BaseCrud
	Client                 *oci_functions.FunctionsManagementClient
	Res                    *oci_functions.Function
	DisableNotFoundRetries bool
}

func (s *FunctionsFunctionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FunctionsFunctionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_functions.FunctionLifecycleStateCreating),
	}
}

func (s *FunctionsFunctionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_functions.FunctionLifecycleStateActive),
	}
}

func (s *FunctionsFunctionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_functions.FunctionLifecycleStateDeleting),
	}
}

func (s *FunctionsFunctionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_functions.FunctionLifecycleStateDeleted),
	}
}

func (s *FunctionsFunctionResourceCrud) Create() error {
	request := oci_functions.CreateFunctionRequest{}

	if applicationId, ok := s.D.GetOkExists("application_id"); ok {
		tmp := applicationId.(string)
		request.ApplicationId = &tmp
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

	if image, ok := s.D.GetOkExists("image"); ok {
		tmp := image.(string)
		request.Image = &tmp
	}

	if imageDigest, ok := s.D.GetOkExists("image_digest"); ok {
		tmp := imageDigest.(string)
		request.ImageDigest = &tmp
	}

	if memoryInMBs, ok := s.D.GetOkExists("memory_in_mbs"); ok {
		tmp := memoryInMBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert memoryInMBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.MemoryInMBs = &tmpInt64
	}

	if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
		tmp := timeoutInSeconds.(int)
		request.TimeoutInSeconds = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.CreateFunction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Function
	return nil
}

func (s *FunctionsFunctionResourceCrud) Get() error {
	request := oci_functions.GetFunctionRequest{}

	tmp := s.D.Id()
	request.FunctionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.GetFunction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Function
	return nil
}

func (s *FunctionsFunctionResourceCrud) Update() error {
	request := oci_functions.UpdateFunctionRequest{}

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

	tmp := s.D.Id()
	request.FunctionId = &tmp

	if image, ok := s.D.GetOkExists("image"); ok {
		tmp := image.(string)
		request.Image = &tmp
	}

	if imageDigest, ok := s.D.GetOkExists("image_digest"); ok {
		tmp := imageDigest.(string)
		request.ImageDigest = &tmp
	}

	if memoryInMBs, ok := s.D.GetOkExists("memory_in_mbs"); ok {
		tmp := memoryInMBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert memoryInMBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.MemoryInMBs = &tmpInt64
	}

	if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
		tmp := timeoutInSeconds.(int)
		request.TimeoutInSeconds = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.UpdateFunction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Function
	return nil
}

func (s *FunctionsFunctionResourceCrud) Delete() error {
	request := oci_functions.DeleteFunctionRequest{}

	tmp := s.D.Id()
	request.FunctionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "functions")

	_, err := s.Client.DeleteFunction(context.Background(), request)
	return err
}

func (s *FunctionsFunctionResourceCrud) SetData() error {
	if s.Res.ApplicationId != nil {
		s.D.Set("application_id", *s.Res.ApplicationId)
	}

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

	if s.Res.Image != nil {
		s.D.Set("image", *s.Res.Image)
	}

	if s.Res.ImageDigest != nil {
		s.D.Set("image_digest", *s.Res.ImageDigest)
	}

	if s.Res.InvokeEndpoint != nil {
		s.D.Set("invoke_endpoint", *s.Res.InvokeEndpoint)
	}

	if s.Res.MemoryInMBs != nil {
		s.D.Set("memory_in_mbs", strconv.FormatInt(*s.Res.MemoryInMBs, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeoutInSeconds != nil {
		s.D.Set("timeout_in_seconds", *s.Res.TimeoutInSeconds)
	}

	return nil
}

func (s *FunctionsFunctionResourceCrud) ExtraWaitPostDelete() time.Duration {
	if httpreplay.ShouldRetryImmediately() {
		return time.Duration(1 * time.Second)
	}
	log.Printf("[DEBUG] Waiting for 40 minutes post destroy of function resource due to known service issue")
	return time.Duration(40 * time.Minute)
}
