// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"
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
				func(ctx context.Context, old, new, meta interface{}) bool {
					return (old.(string) != new.(string)) && old.(string) != ""
				},
				func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
					//Image and image_digest are not used for PBFs. Image digest should be empty and should not be computed for PBF Functions.
					if isSourceTypePbf(d) {
						return nil
					}
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
				func(ctx context.Context, v, m interface{}) bool {
					// mark explicit requests for recomputation as "known after apply"
					return v.(string) == "" || v.(string) == requireRecompute
				},
				func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
					//Image and image_digest are not used for PBFs. Image digest should be empty and should not be computed for PBF Functions.
					if isSourceTypePbf(d) {
						return nil
					}
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
			"memory_in_mbs": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
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
			"image": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_digest": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,

				DefaultFunc: func() (interface{}, error) {
					return requireRecompute, nil
				},
			},
			"provisioned_concurrency_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"strategy": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CONSTANT",
								"NONE",
							}, true),
						},

						// Optional
						"count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"source_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"pbf_listing_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"source_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"PRE_BUILT_FUNCTIONS",
							}, true),
						},

						// Optional

						// Computed
					},
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
			"shape": {
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
func isSourceTypePbf(d *schema.ResourceDiff) bool {
	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
	sourceTypeRaw, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	if strings.EqualFold(sourceType, "PRE_BUILT_FUNCTIONS") {
		return true
	}
	return false
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
		request.Config = tfresource.ObjectMapToStringMap(config.(map[string]interface{}))
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

	if provisionedConcurrencyConfig, ok := s.D.GetOkExists("provisioned_concurrency_config"); ok {
		if tmpList := provisionedConcurrencyConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "provisioned_concurrency_config", 0)
			tmp, err := s.mapToFunctionProvisionedConcurrencyConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ProvisionedConcurrencyConfig = tmp
		}
	}

	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
			tmp, err := s.mapToFunctionSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SourceDetails = tmp
		}
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
		request.Config = tfresource.ObjectMapToStringMap(config.(map[string]interface{}))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

	if provisionedConcurrencyConfig, ok := s.D.GetOkExists("provisioned_concurrency_config"); ok {
		if tmpList := provisionedConcurrencyConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "provisioned_concurrency_config", 0)
			tmp, err := s.mapToFunctionProvisionedConcurrencyConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ProvisionedConcurrencyConfig = tmp
		}
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

	if s.Res.ProvisionedConcurrencyConfig != nil {
		provisionedConcurrencyConfigArray := []interface{}{}
		if provisionedConcurrencyConfigMap := FunctionProvisionedConcurrencyConfigToMap(&s.Res.ProvisionedConcurrencyConfig); provisionedConcurrencyConfigMap != nil {
			provisionedConcurrencyConfigArray = append(provisionedConcurrencyConfigArray, provisionedConcurrencyConfigMap)
		}
		s.D.Set("provisioned_concurrency_config", provisionedConcurrencyConfigArray)
	} else {
		s.D.Set("provisioned_concurrency_config", nil)
	}

	s.D.Set("shape", s.Res.Shape)

	if s.Res.SourceDetails != nil {
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := FunctionSourceDetailsToMap(&s.Res.SourceDetails); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
			// image_digest is computed field when only image is passed. For PBF image is not required
			//and hence image_digest will be not computed. Set the value to empty string to avoid showing as computed for PBF
			//s.D.Set("image_digest", "")
		}
		s.D.Set("source_details", sourceDetailsArray)
	} else {
		s.D.Set("source_details", nil)
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

func (s *FunctionsFunctionResourceCrud) mapToFunctionProvisionedConcurrencyConfig(fieldKeyFormat string) (oci_functions.FunctionProvisionedConcurrencyConfig, error) {
	var baseObject oci_functions.FunctionProvisionedConcurrencyConfig
	//discriminator
	strategyRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "strategy"))
	var strategy string
	if ok {
		strategy = strategyRaw.(string)
	} else {
		strategy = "" // default value
	}
	switch strings.ToLower(strategy) {
	case strings.ToLower("CONSTANT"):
		details := oci_functions.ConstantProvisionedConcurrencyConfig{}
		if count, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "count")); ok {
			tmp := count.(int)
			details.Count = &tmp
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_functions.NoneProvisionedConcurrencyConfig{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown strategy '%v' was specified", strategy)
	}
	return baseObject, nil
}

func FunctionProvisionedConcurrencyConfigToMap(obj *oci_functions.FunctionProvisionedConcurrencyConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_functions.ConstantProvisionedConcurrencyConfig:
		result["strategy"] = "CONSTANT"

		if v.Count != nil {
			result["count"] = int(*v.Count)
		}
	case oci_functions.NoneProvisionedConcurrencyConfig:
		result["strategy"] = "NONE"
	default:
		log.Printf("[WARN] Received 'strategy' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FunctionsFunctionResourceCrud) mapToFunctionSourceDetails(fieldKeyFormat string) (oci_functions.FunctionSourceDetails, error) {
	var baseObject oci_functions.FunctionSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("PRE_BUILT_FUNCTIONS"):
		details := oci_functions.PreBuiltFunctionSourceDetails{}
		if pbfListingId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pbf_listing_id")); ok {
			tmp := pbfListingId.(string)
			details.PbfListingId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func FunctionSourceDetailsToMap(obj *oci_functions.FunctionSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_functions.PreBuiltFunctionSourceDetails:
		result["source_type"] = "PRE_BUILT_FUNCTIONS"

		if v.PbfListingId != nil {
			result["pbf_listing_id"] = string(*v.PbfListingId)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
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
