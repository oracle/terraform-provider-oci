// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"
	"fmt"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_functions "github.com/oracle/oci-go-sdk/v56/functions"
)

// The Functions API uses `imageDigest` as an optional I/O parameter. If unspecified, the controlplane
// will compute the appropriate digest and utilise that. However, if the caller specifies `imageDigest`, that
// digest value will be preferred.
// This doesn't play well with Terraform's notion of what constitutes a change: in particular, it may supply
// the old digest value for an image from its state. This can prevent users from updating their functions
// to a newer tag (since the old digest may still be available in their image repo).
// We apply some heuristics here to determine when we should pass through the current image_digest value,
// or omit it from API calls.
// Additionally, we explicitly support the behaviour of setting
//    image_digest = ""
// in an Update to *force* the controlplane-side resolution of the image coordinates.

// In summary:
// - same image, leaving the digest unspecified -> won't force an Update
// - changing the image, leaving the digest unspecified -> works, updates the digest to correspond to the image
// - same image, digest explicitly empty -> works, forces the controlplane to supply a new value

const requireRecompute = "require-recompute"

func FunctionsFunctionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFunctionsFunction,
		Read:     readFunctionsFunction,
		Update:   updateFunctionsFunction,
		Delete:   deleteFunctionsFunction,
		CustomizeDiff: customdiff.All(
			customdiff.IfValueChange("image",
				func(old, new, meta interface{}) bool {
					return (old.(string) != new.(string)) && old.(string) != ""
				},
				func(d *schema.ResourceDiff, meta interface{}) error {
					o, n := d.GetChange("image_digest")

					if o == n || n == requireRecompute || n == "" {
						// The user's changing the image.
						// Mark image_digest as "known after apply" if there is no corresponding
						// explicit Update to that field - either a supplied value or a demand for
						// controlplane-side recalculation.
						d.SetNewComputed("image_digest")
					}
					return nil
				}),
			customdiff.IfValue("image_digest",
				func(v, m interface{}) bool {
					// mark explicit requests for recomputation as "known after apply"
					return v.(string) == "" || v.(string) == requireRecompute
				},
				func(d *schema.ResourceDiff, meta interface{}) error {
					d.SetNewComputed("image_digest")
					return nil
				}),
		),
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
				ValidateFunc:     utils.ValidateInt64TypeString,
				DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
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
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
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

				DefaultFunc: func() (interface{}, error) {
					return requireRecompute, nil
				},
			},
			"timeout_in_seconds": {
				Type:     schema.TypeInt,
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
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readFunctionsFunction(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsFunctionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateFunctionsFunction(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsFunctionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFunctionsFunction(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsFunctionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FunctionsFunctionResourceCrud struct {
	tfresource.BaseCrud
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
		request.Config = utils.ObjectMapToStringMap(config.(map[string]interface{}))
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

	if image, ok := s.D.GetOkExists("image"); ok {
		tmp := image.(string)
		request.Image = &tmp
	}

	// This is important: we might receive the sentinel value during a Create. If we do, do *not* pass that
	// through to the API.
	if imageDigest, ok := s.D.GetOkExists("image_digest"); ok {
		tmp := imageDigest.(string)
		if tmp != "" && tmp != requireRecompute {
			request.ImageDigest = &tmp
		}
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

	if traceConfig, ok := s.D.GetOkExists("trace_config"); ok {
		if tmpList := traceConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "trace_config", 0)
			tmp, err := s.mapToFunctionTraceConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TraceConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

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
		request.Config = utils.ObjectMapToStringMap(config.(map[string]interface{}))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.FunctionId = &tmp

	if image, ok := s.D.GetOkExists("image"); ok {
		tmp := image.(string)
		request.Image = &tmp
	}

	// Again, during an Update we must detect the special sentinel value and avoid passing it to the API.
	if imageDigest, ok := s.D.GetOkExists("image_digest"); ok {
		tmp := imageDigest.(string)
		if tmp != "" && tmp != requireRecompute {
			request.ImageDigest = &tmp
		}
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

	if traceConfig, ok := s.D.GetOkExists("trace_config"); ok {
		if tmpList := traceConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "trace_config", 0)
			tmp, err := s.mapToFunctionTraceConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TraceConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

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
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
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

	if s.Res.TraceConfig != nil {
		s.D.Set("trace_config", []interface{}{FunctionTraceConfigToMap(s.Res.TraceConfig)})
	} else {
		s.D.Set("trace_config", nil)
	}

	return nil
}

func (s *FunctionsFunctionResourceCrud) mapToFunctionTraceConfig(fieldKeyFormat string) (oci_functions.FunctionTraceConfig, error) {
	result := oci_functions.FunctionTraceConfig{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func FunctionTraceConfigToMap(obj *oci_functions.FunctionTraceConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}
