// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
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

const (
	requireRecompute = "require-recompute"

	sourceTypeArchive        = "ARCHIVE"
	sourceTypeContainerImage = "CONTAINER_IMAGE"
	sourceTypePbf            = "PRE_BUILT_FUNCTIONS"
)

func FunctionsFunctionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:                       tfresource.DefaultTimeout,
		Create:                         createFunctionsFunction,
		Read:                           readFunctionsFunction,
		Update:                         updateFunctionsFunction,
		Delete:                         deleteFunctionsFunction,
		ValidateRawResourceConfigFuncs: []schema.ValidateRawResourceConfigFunc{warnOnExplicitLegacyImageFields},
		CustomizeDiff: customdiff.All(
			customdiff.IfValueChange("image",
				func(ctx context.Context, old, new, meta interface{}) bool {
					return (old.(string) != new.(string)) && old.(string) != ""
				},
				func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
					// If source_details.image is configured, the top-level image is ignored.
					if sourceDetailsStringFieldOverridesLegacyInDiff(d, "image", false) {
						return nil
					}
					//Image and image_digest are not used for PBFs. Image digest should be empty and should not be computed for PBF Functions.
					if isSourceTypePbf(d) {
						return nil
					}
					if sourceDetailsStringFieldOverridesLegacyInDiff(d, "image_digest", true) {
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
					// If source_details.image_digest is configured, the top-level image_digest is ignored.
					if sourceDetailsStringFieldOverridesLegacyInDiff(d, "image_digest", true) {
						return nil
					}
					//Image and image_digest are not used for PBFs. Image digest should be empty and should not be computed for PBF Functions.
					if isSourceTypePbf(d) {
						return nil
					}
					d.SetNewComputed("image_digest")
					return nil
				}),
			customdiff.IfValueChange("source_details.0.image",
				func(ctx context.Context, old, new, meta interface{}) bool {
					oldImage, oldOk := old.(string)
					newImage, newOk := new.(string)
					return oldOk && newOk && oldImage != newImage && oldImage != ""
				},
				func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
					if !isSourceTypeContainerImage(d) {
						return nil
					}
					o, n := d.GetChange("source_details.0.image_digest")
					if o == n || n == requireRecompute || n == "" {
						// The user's changing the canonical container image.
						// Mark nested image_digest as "known after apply" if there is no corresponding
						// explicit update to that field, so the control plane can recompute it.
						d.SetNewComputed("source_details.0.image_digest")
					}
					return nil
				}),
			customdiff.IfValue("source_details.0.image_digest",
				func(ctx context.Context, v, m interface{}) bool {
					value, ok := v.(string)
					return ok && (value == "" || value == requireRecompute)
				},
				func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
					if !isSourceTypeContainerImage(d) {
						return nil
					}
					d.SetNewComputed("source_details.0.image_digest")
					return nil
				}),
			validateLegacyVsSourceDetails,
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
			// optional for a year
			"source_details": {
				Type:     schema.TypeList,
				MaxItems: 1,
				MinItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"source_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								sourceTypeArchive,
								sourceTypeContainerImage,
								sourceTypePbf,
							}, true),
						},

						// Optional
						"archive_source_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"archive_source_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"DIRECT_ARCHIVE",
											"OBJECT_STORAGE_ARCHIVE",
										}, true),
									},

									// Optional
									"archive_file": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"object": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"object_version_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"handler": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						},
						"pbf_listing_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"runtime_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"functions_runtime_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"functions_runtime_version_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"runtime_config_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"FUNCTION_UPDATE",
											"MANUAL",
										}, true),
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
						"source_code_sha256": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
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
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"detached_mode_timeout_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"failure_destination": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"kind": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"NONE",
								"NOTIFICATION",
								"QUEUE",
								"STREAM",
							}, true),
						},

						// Optional
						"channel_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"queue_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stream_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"topic_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return sourceDetailsStringFieldOverridesLegacy(d, "image", false)
				},
				Deprecated: tfresource.FieldDeprecatedAndOverridenByAnother("image", "source_details.image"),
			},
			"image_digest": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return sourceDetailsStringFieldOverridesLegacy(d, "image_digest", true)
				},
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
			"success_destination": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"kind": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"NONE",
								"NOTIFICATION",
								"QUEUE",
								"STREAM",
							}, true),
						},

						// Optional
						"channel_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"queue_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stream_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"topic_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

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

// warnOnExplicitLegacyImageFields emits standard deprecation warnings only for
// legacy image fields present in user config, including unknown variable values.
func warnOnExplicitLegacyImageFields(ctx context.Context, req schema.ValidateResourceConfigFuncRequest, resp *schema.ValidateResourceConfigFuncResponse) {
	if image, configured := rawConfigAttribute(req.RawConfig, "image"); configured && !image.IsKnown() {
		appendLegacyImageFieldWarning(resp, "image", "source_details.image")
	}
	if rawConfigHasAttribute(req.RawConfig, "image_digest") {
		appendLegacyImageFieldWarning(resp, "image_digest", "source_details.image_digest")
	}
}

func appendLegacyImageFieldWarning(resp *schema.ValidateResourceConfigFuncResponse, deprecatedField string, newField string) {
	resp.Diagnostics = append(resp.Diagnostics, diag.Diagnostic{
		Severity:      diag.Warning,
		Summary:       "Argument is deprecated",
		Detail:        tfresource.FieldDeprecatedAndOverridenByAnother(deprecatedField, newField),
		AttributePath: cty.GetAttrPath(deprecatedField),
	})
}

// rawConfigHasAttribute checks whether a top-level attribute was explicitly set
// in config without relying on its resolved value.
func rawConfigHasAttribute(rawConfig cty.Value, attributeName string) bool {
	_, configured := rawConfigAttribute(rawConfig, attributeName)
	return configured
}

func rawConfigAttribute(rawConfig cty.Value, attributeName string) (cty.Value, bool) {
	if rawConfig.IsNull() || !rawConfig.Type().IsObjectType() || !rawConfig.Type().HasAttribute(attributeName) {
		return cty.NilVal, false
	}
	value := rawConfig.GetAttr(attributeName)
	return value, !value.IsNull()
}

// validateLegacyVsSourceDetails enforces the function source input contract at plan time.
// Users can either use legacy top-level image/image_digest fields or the canonical
// source_details block. For container-image functions, source_details values override
// matching legacy fields, and legacy values can fill source_details fields that are omitted.
func validateLegacyVsSourceDetails(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	legacyImageConfigured := isAttributeExplicitlyConfiguredInDiff(d, "image")
	legacyDigestConfigured := isAttributeExplicitlyConfiguredInDiff(d, "image_digest")
	legacyImageSet := isConfiguredNonEmptyStringInDiff(d, cty.Path{cty.GetAttrStep{Name: "image"}})
	legacyDigestSet := isLegacyImageDigestConfiguredInDiff(d)

	sourceDetailsConfigured := isSourceDetailsExplicitlyConfiguredInDiff(d)
	if sourceDetailsConfigured {
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
		sourceTypeRaw, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
		if !ok {
			return fmt.Errorf("'source_details.source_type' must be set")
		}
		sourceType := strings.ToUpper(strings.TrimSpace(sourceTypeRaw.(string)))
		if sourceType != sourceTypeContainerImage && (legacyImageConfigured || legacyDigestConfigured) {
			return fmt.Errorf("legacy 'image'/'image_digest' can only be set when 'source_details.source_type' is CONTAINER_IMAGE")
		}
		switch sourceType {
		case sourceTypeContainerImage:
			sourceImageSet := isConfiguredNonEmptyStringInDiff(d, cty.Path{
				cty.GetAttrStep{Name: "source_details"},
				cty.IndexStep{Key: cty.NumberIntVal(0)},
				cty.GetAttrStep{Name: "image"},
			})
			if !sourceImageSet && !legacyImageSet {
				return fmt.Errorf("'source_details.image' or legacy 'image' must be set when source_type is CONTAINER_IMAGE")
			}
		case sourceTypePbf:
			if pbfRaw, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pbf_listing_id")); !ok || strings.TrimSpace(pbfRaw.(string)) == "" {
				return fmt.Errorf("'source_details.pbf_listing_id' must be set when source_type is PRE_BUILT_FUNCTIONS")
			}
		case sourceTypeArchive:
			if archiveRaw, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "archive_source_details")); !ok || len(archiveRaw.([]interface{})) == 0 {
				return fmt.Errorf("'source_details.archive_source_details' must be set when source_type is ARCHIVE")
			}
			if runtimeRaw, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "runtime_config")); !ok || len(runtimeRaw.([]interface{})) == 0 {
				return fmt.Errorf("'source_details.runtime_config' must be set when source_type is ARCHIVE")
			}
		default:
			return fmt.Errorf("unknown source_type '%v' was specified", sourceType)
		}
		if runtimeRaw, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "runtime_config")); ok && len(runtimeRaw.([]interface{})) > 0 {
			runtimeKeyFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "runtime_config"), 0)
			runtimeTypeRaw, ok := d.GetOkExists(fmt.Sprintf(runtimeKeyFormat, "runtime_config_type"))
			if !ok {
				return fmt.Errorf("'source_details.runtime_config.runtime_config_type' must be set")
			}
			if strings.EqualFold(runtimeTypeRaw.(string), "MANUAL") {
				if v, ok := d.GetOkExists(fmt.Sprintf(runtimeKeyFormat, "functions_runtime_version_id")); !ok || strings.TrimSpace(v.(string)) == "" {
					return fmt.Errorf("'source_details.runtime_config.functions_runtime_version_id' must be set when runtime_config_type is MANUAL")
				}
			}
		}
		return nil
	}

	// Validate legacy-only fallback.
	if !legacyImageSet && !legacyDigestSet {
		if sourceTypeRaw, ok := d.GetOkExists("source_details.0.source_type"); ok && strings.TrimSpace(sourceTypeRaw.(string)) != "" {
			return nil
		}
		return fmt.Errorf("either 'source_details' or a legacy image field must be set")
	}
	if sourceTypeRaw, ok := d.GetOkExists("source_details.0.source_type"); ok && !strings.EqualFold(sourceTypeRaw.(string), sourceTypeContainerImage) {
		return fmt.Errorf("legacy 'image' cannot be set when 'source_details.source_type' is not CONTAINER_IMAGE")
	}
	return nil
}

// isConfiguredNonEmptyStringInDiff treats known non-empty strings and unknown
// variable values as configured, while empty strings remain unset.
func isConfiguredNonEmptyStringInDiff(d *schema.ResourceDiff, path cty.Path) bool {
	if d == nil {
		return false
	}
	value, diags := d.GetRawConfigAt(path)
	if diags.HasError() || value.IsNull() {
		return false
	}
	if !value.IsKnown() {
		return true
	}
	if value.Type().Equals(cty.String) {
		return strings.TrimSpace(value.AsString()) != ""
	}
	return true
}

// isLegacyImageDigestConfiguredInDiff treats the recompute sentinel as an omitted digest.
func isLegacyImageDigestConfiguredInDiff(d *schema.ResourceDiff) bool {
	if !isAttributeExplicitlyConfiguredInDiff(d, "image_digest") {
		return false
	}
	imageDigest, configured := getExplicitlyConfiguredStringInDiff(d, cty.Path{cty.GetAttrStep{Name: "image_digest"}})
	return !configured || imageDigest != requireRecompute
}

// getExplicitlyConfiguredStringInDiff returns a trimmed string only when raw config has a non-empty value.
func getExplicitlyConfiguredStringInDiff(d *schema.ResourceDiff, path cty.Path) (string, bool) {
	if d == nil {
		return "", false
	}
	value, diags := d.GetRawConfigAt(path)
	if diags.HasError() || value.IsNull() || !value.IsKnown() || !value.Type().Equals(cty.String) {
		return "", false
	}
	result := strings.TrimSpace(value.AsString())
	return result, result != ""
}

func isSourceDetailsExplicitlyConfigured(d *schema.ResourceData) bool {
	if d == nil {
		return false
	}
	value, diags := d.GetRawConfigAt(cty.Path{cty.GetAttrStep{Name: "source_details"}})
	if diags.HasError() || value.IsNull() {
		return false
	}
	if !value.IsKnown() {
		return true
	}
	if !value.CanIterateElements() {
		return true
	}
	return value.LengthInt() > 0
}

func isSourceDetailsExplicitlyConfiguredInDiff(d *schema.ResourceDiff) bool {
	return isAttributeExplicitlyConfiguredInDiff(d, "source_details")
}

func isAttributeExplicitlyConfiguredInDiff(d *schema.ResourceDiff, attributeName string) bool {
	if d == nil {
		return false
	}
	value, diags := d.GetRawConfigAt(cty.Path{cty.GetAttrStep{Name: attributeName}})
	if diags.HasError() || value.IsNull() {
		return false
	}
	if !value.IsKnown() {
		return true
	}
	if !value.CanIterateElements() {
		return true
	}
	return value.LengthInt() > 0
}

// Build SourceDetails for Create from either source_details or legacy fields.
func (s *FunctionsFunctionResourceCrud) buildCreateSourceDetails() (oci_functions.CreateFunctionSourceDetails, error) {
	// Prefer explicit source_details. State-computed source_details must not mask legacy image changes.
	if v, ok := s.D.GetOkExists("source_details"); ok && isSourceDetailsExplicitlyConfigured(s.D) {
		if lst := v.([]interface{}); len(lst) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
			return s.mapToCreateFunctionSourceDetails(fieldKeyFormat)
		}
	}

	// Legacy fallback: create CONTAINER_IMAGE sourceDetails from top-level image/image_digest.
	// Preserve the historical behavior where image and image_digest are independent inputs.
	image, imageDigest, ok := getLegacyContainerImageSourceValues(s.D)
	if !ok {
		return nil, fmt.Errorf("either source_details or a legacy image field must be set")
	}
	details := oci_functions.CreateContainerImageFunctionSourceDetails{}
	if image != nil {
		details.Image = image
	}
	if imageDigest != nil {
		details.ImageDigest = imageDigest
	}

	return details, nil
}

// Build SourceDetails for Update from either source_details or legacy fields.
func (s *FunctionsFunctionResourceCrud) buildUpdateSourceDetails() (oci_functions.UpdateFunctionSourceDetails, error) {
	// Prefer explicit source_details. State-computed source_details must not mask legacy image changes.
	if v, ok := s.D.GetOkExists("source_details"); ok && isSourceDetailsExplicitlyConfigured(s.D) {
		if lst := v.([]interface{}); len(lst) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
			return s.mapToUpdateFunctionSourceDetails(fieldKeyFormat)
		}
	}

	// Legacy fallback: update CONTAINER_IMAGE sourceDetails from top-level image/image_digest.
	// Preserve the historical behavior where image and image_digest are independent inputs.
	image, imageDigest, ok := getLegacyContainerImageSourceValues(s.D)
	if !ok {
		return nil, fmt.Errorf("either source_details or a legacy image field must be set")
	}
	details := oci_functions.UpdateContainerImageFunctionSourceDetails{}
	if image != nil {
		details.Image = image
	}
	if imageDigest != nil {
		details.ImageDigest = imageDigest
	}

	return details, nil
}

func getLegacyContainerImageSourceValues(d *schema.ResourceData) (*string, *string, bool) {
	var image *string
	var imageDigest *string

	if imgRaw, ok := d.GetOkExists("image"); ok {
		img := strings.TrimSpace(imgRaw.(string))
		if img != "" {
			image = &img
		}
	}
	if image == nil {
		if img, configured := getConfiguredStringInResourceData(d, cty.Path{cty.GetAttrStep{Name: "image"}}, "image"); configured && img != "" {
			image = &img
		}
	}

	if digRaw, ok := d.GetOkExists("image_digest"); ok {
		dig := strings.TrimSpace(digRaw.(string))
		if shouldSendImageDigest(dig) {
			imageDigest = &dig
		}
	}
	if imageDigest == nil {
		if dig, configured := getConfiguredStringInResourceData(d, cty.Path{cty.GetAttrStep{Name: "image_digest"}}, "image_digest"); configured && shouldSendImageDigest(dig) {
			imageDigest = &dig
		}
	}

	return image, imageDigest, image != nil || imageDigest != nil
}

// getConfiguredStringInResourceData reads a raw-config path and its resolved state key.
// The bool means the user configured the path, even if the resolved value is empty.
func getConfiguredStringInResourceData(d *schema.ResourceData, path cty.Path, stateKey string) (string, bool) {
	if d == nil {
		return "", false
	}
	rawValue, diags := d.GetRawConfigAt(path)
	if diags.HasError() || rawValue.IsNull() {
		return "", false
	}
	if value, ok := d.GetOkExists(stateKey); ok {
		return strings.TrimSpace(value.(string)), true
	}
	if rawValue.IsKnown() && rawValue.Type().Equals(cty.String) {
		return strings.TrimSpace(rawValue.AsString()), true
	}
	return "", true
}

func sourceDetailsStringFieldOverridesLegacyInDiff(d *schema.ResourceDiff, fieldName string, allowEmpty bool) bool {
	if d == nil {
		return false
	}
	return sourceDetailsStringFieldOverridesLegacyAtPath(d, fieldName, allowEmpty)
}

// sourceDetailsStringFieldOverridesLegacy reports whether a canonical
// source_details field should suppress the matching top-level legacy field.
func sourceDetailsStringFieldOverridesLegacy(d *schema.ResourceData, fieldName string, allowEmpty bool) bool {
	if d == nil {
		return false
	}
	return sourceDetailsStringFieldOverridesLegacyAtPath(d, fieldName, allowEmpty)
}

type rawConfigAtGetter interface {
	GetRawConfigAt(cty.Path) (cty.Value, diag.Diagnostics)
}

func sourceDetailsStringFieldOverridesLegacyAtPath(rawConfig rawConfigAtGetter, fieldName string, allowEmpty bool) bool {
	if sourceType, configured := sourceDetailsSourceType(rawConfig); configured && !strings.EqualFold(sourceType, sourceTypeContainerImage) {
		return true
	}
	path := cty.Path{
		cty.GetAttrStep{Name: "source_details"},
		cty.IndexStep{Key: cty.NumberIntVal(0)},
		cty.GetAttrStep{Name: fieldName},
	}
	rawValue, diags := rawConfig.GetRawConfigAt(path)
	if diags.HasError() || rawValue.IsNull() {
		return false
	}
	if !rawValue.IsKnown() {
		return true
	}
	if rawValue.Type().Equals(cty.String) {
		return allowEmpty || strings.TrimSpace(rawValue.AsString()) != ""
	}
	return true
}

func sourceDetailsSourceType(rawConfig rawConfigAtGetter) (string, bool) {
	path := cty.Path{
		cty.GetAttrStep{Name: "source_details"},
		cty.IndexStep{Key: cty.NumberIntVal(0)},
		cty.GetAttrStep{Name: "source_type"},
	}
	rawValue, diags := rawConfig.GetRawConfigAt(path)
	if diags.HasError() || rawValue.IsNull() || !rawValue.IsKnown() || !rawValue.Type().Equals(cty.String) {
		return "", false
	}
	return strings.ToUpper(strings.TrimSpace(rawValue.AsString())), true
}

// getContainerImageSourceValues merges container-image fields with per-field precedence:
// source_details.image beats image, source_details.image_digest beats image_digest,
// and legacy top-level values fill only the source_details fields that were omitted.
func getContainerImageSourceValues(d *schema.ResourceData, fieldKeyFormat string) (*string, *string, bool) {
	image, imageDigest, _ := getLegacyContainerImageSourceValues(d)

	sourceImagePath := cty.Path{
		cty.GetAttrStep{Name: "source_details"},
		cty.IndexStep{Key: cty.NumberIntVal(0)},
		cty.GetAttrStep{Name: "image"},
	}
	if sourceImage, configured := getConfiguredStringInResourceData(d, sourceImagePath, fmt.Sprintf(fieldKeyFormat, "image")); configured && sourceImage != "" {
		image = &sourceImage
	}

	sourceDigestPath := cty.Path{
		cty.GetAttrStep{Name: "source_details"},
		cty.IndexStep{Key: cty.NumberIntVal(0)},
		cty.GetAttrStep{Name: "image_digest"},
	}
	if sourceDigest, configured := getConfiguredStringInResourceData(d, sourceDigestPath, fmt.Sprintf(fieldKeyFormat, "image_digest")); configured {
		imageDigest = nil
		if shouldSendImageDigest(sourceDigest) {
			imageDigest = &sourceDigest
		}
	}

	return image, imageDigest, image != nil || imageDigest != nil
}

// Source details validation helpers
// isSourceTypePbf checks whether the planned source_details block is for a pre-built function.
func isSourceTypePbf(d *schema.ResourceDiff) bool {
	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
	sourceTypeRaw, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	if strings.EqualFold(sourceType, sourceTypePbf) {
		return true
	}
	return false
}

// isSourceTypeContainerImage checks whether the planned source_details block is container-image based.
func isSourceTypeContainerImage(d *schema.ResourceDiff) bool {
	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
	sourceTypeRaw, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	if !ok {
		return false
	}
	return strings.EqualFold(sourceTypeRaw.(string), sourceTypeContainerImage)
}

// shouldSendImageDigest returns false for values that mean "let OCI recompute the digest".
func shouldSendImageDigest(imageDigest string) bool {
	return imageDigest != "" && imageDigest != requireRecompute
}

func createFunctionsFunction(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsFunctionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()
	sync.WorkRequestClient = m.(*client.OracleClients).FunctionsWorkRequestManagementClient()

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
	sync.WorkRequestClient = m.(*client.OracleClients).FunctionsWorkRequestManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFunctionsFunction(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsFunctionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).FunctionsWorkRequestManagementClient()

	return tfresource.DeleteResource(d, sync)
}

type FunctionsFunctionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_functions.FunctionsManagementClient
	Res                    *oci_functions.Function
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_functions.WorkRequestManagementClient
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

	if detachedModeTimeoutInSeconds, ok := s.D.GetOkExists("detached_mode_timeout_in_seconds"); ok {
		tmp := detachedModeTimeoutInSeconds.(int)
		request.DetachedModeTimeoutInSeconds = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if failureDestination, ok := s.D.GetOkExists("failure_destination"); ok {
		if tmpList := failureDestination.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "failure_destination", 0)
			tmp, err := s.mapToFailureDestinationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FailureDestination = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	// IMPORTANT: do not set request.Image / request.ImageDigest anymore.
	// Those are legacy API fields and will break when removed server-side.

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

	// Always populate SourceDetails (canonical). If source_details not provided, derive from legacy image fields.
	sd, err := s.buildCreateSourceDetails()
	if err != nil {
		return err
	}
	request.SourceDetails = sd

	if successDestination, ok := s.D.GetOkExists("success_destination"); ok {
		if tmpList := successDestination.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "success_destination", 0)
			tmp, err := s.mapToSuccessDestinationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SuccessDestination = tmp
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

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getFunctionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions"), oci_functions.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

// getFunctionFromWorkRequest waits for an async function operation and refreshes Terraform state.
func (s *FunctionsFunctionResourceCrud) getFunctionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_functions.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	functionId, err := functionWaitForWorkRequest(workId, "function",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, functionId)
		_, cancelErr := s.WorkRequestClient.CancelWorkRequest(context.Background(),
			oci_functions.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*functionId)

	return s.Get()
}

// functionWorkRequestShouldRetryFunc keeps polling until the work request finishes or times out.
func functionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "functions", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_functions.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

// functionWaitForWorkRequest polls a work request and returns the affected function identifier.
func functionWaitForWorkRequest(wId *string, entityType string, action oci_functions.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_functions.WorkRequestManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "functions")
	retryPolicy.ShouldRetryOperation = functionWorkRequestShouldRetryFunc(timeout)

	response := oci_functions.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_functions.OperationStatusInProgress),
			string(oci_functions.OperationStatusAccepted),
			string(oci_functions.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_functions.OperationStatusSucceeded),
			string(oci_functions.OperationStatusFailed),
			string(oci_functions.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_functions.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_functions.OperationStatusFailed || response.Status == oci_functions.OperationStatusCanceled {
		return nil, getErrorFromFunctionsFunctionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

// getErrorFromFunctionsFunctionWorkRequest converts OCI work request errors into a Terraform error.
func getErrorFromFunctionsFunctionWorkRequest(client *oci_functions.WorkRequestManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_functions.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_functions.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
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

	if detachedModeTimeoutInSeconds, ok := s.D.GetOkExists("detached_mode_timeout_in_seconds"); ok {
		tmp := detachedModeTimeoutInSeconds.(int)
		request.DetachedModeTimeoutInSeconds = &tmp
	}

	if failureDestination, ok := s.D.GetOkExists("failure_destination"); ok {
		if tmpList := failureDestination.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "failure_destination", 0)
			tmp, err := s.mapToFailureDestinationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FailureDestination = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.FunctionId = &tmp

	if isSourceDetailsExplicitlyConfigured(s.D) || s.D.HasChange("image") || s.D.HasChange("image_digest") || s.D.HasChange("source_details") {
		sd, err := s.buildUpdateSourceDetails()
		if err != nil {
			return err
		}
		if sd != nil {
			request.SourceDetails = sd
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

	if successDestination, ok := s.D.GetOkExists("success_destination"); ok {
		if tmpList := successDestination.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "success_destination", 0)
			tmp, err := s.mapToSuccessDestinationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SuccessDestination = tmp
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

	workId := response.OpcWorkRequestId
	return s.getFunctionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions"), oci_functions.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FunctionsFunctionResourceCrud) Delete() error {
	request := oci_functions.DeleteFunctionRequest{}

	tmp := s.D.Id()
	request.FunctionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.DeleteFunction(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := functionWaitForWorkRequest(workId, "function",
		oci_functions.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
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

	if s.Res.DetachedModeTimeoutInSeconds != nil {
		s.D.Set("detached_mode_timeout_in_seconds", *s.Res.DetachedModeTimeoutInSeconds)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FailureDestination != nil {
		failureDestinationArray := []interface{}{}
		if failureDestinationMap := FailureDestinationDetailsToMap(&s.Res.FailureDestination); failureDestinationMap != nil {
			failureDestinationArray = append(failureDestinationArray, failureDestinationMap)
		}
		s.D.Set("failure_destination", failureDestinationArray)
	} else {
		s.D.Set("failure_destination", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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
			preserveDirectArchiveFileFromState(s.D, sourceDetailsMap)
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
			// image_digest is computed field when only image is passed. For PBF image is not required
			//and hence image_digest will be not computed. Set the value to empty string to avoid showing as computed for PBF
			//s.D.Set("image_digest", "")
		}
		s.D.Set("source_details", sourceDetailsArray)
	} else {
		s.D.Set("source_details", nil)
	}

	if s.Res.SourceDetails != nil {
		switch v := (s.Res.SourceDetails).(type) {
		case oci_functions.ContainerImageFunctionSourceDetails:
			setLegacyContainerImageFields(s.D, v.Image, v.ImageDigest)
		default:
			clearLegacyContainerImageFields(s.D)
		}
	} else {
		clearLegacyContainerImageFields(s.D)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SuccessDestination != nil {
		successDestinationArray := []interface{}{}
		if successDestinationMap := SuccessDestinationDetailsToMap(&s.Res.SuccessDestination); successDestinationMap != nil {
			successDestinationArray = append(successDestinationArray, successDestinationMap)
		}
		s.D.Set("success_destination", successDestinationArray)
	} else {
		s.D.Set("success_destination", nil)
	}

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

// preserveDirectArchiveFileFromState keeps direct-upload archive bytes in state when OCI does not echo them.
func preserveDirectArchiveFileFromState(d *schema.ResourceData, sourceDetailsMap map[string]interface{}) {
	if sourceDetailsMap == nil || !strings.EqualFold(fmt.Sprintf("%v", sourceDetailsMap["source_type"]), sourceTypeArchive) {
		return
	}
	archiveSourceDetailsRaw, ok := sourceDetailsMap["archive_source_details"]
	if !ok {
		return
	}
	archiveSourceDetails, ok := archiveSourceDetailsRaw.([]interface{})
	if !ok || len(archiveSourceDetails) == 0 {
		return
	}
	archiveSourceDetailsMap, ok := archiveSourceDetails[0].(map[string]interface{})
	if !ok || !strings.EqualFold(fmt.Sprintf("%v", archiveSourceDetailsMap["archive_source_type"]), "DIRECT_ARCHIVE") {
		return
	}
	if _, ok := archiveSourceDetailsMap["archive_file"]; ok {
		return
	}
	if archiveFile, ok := d.GetOkExists("source_details.0.archive_source_details.0.archive_file"); ok {
		if archiveFileValue, ok := archiveFile.(string); ok && archiveFileValue != "" {
			archiveSourceDetailsMap["archive_file"] = archiveFileValue
		}
	}
}

// setLegacyContainerImageFields mirrors container-image source_details into deprecated top-level fields.
func setLegacyContainerImageFields(d *schema.ResourceData, image *string, imageDigest *string) {
	if image != nil {
		_ = d.Set("image", *image)
	}
	if imageDigest != nil {
		_ = d.Set("image_digest", *imageDigest)
	} else {
		// Align state with the default sentinel to avoid diffs for legacy configs that omitted image_digest.
		_ = d.Set("image_digest", requireRecompute)
	}
}

// clearLegacyContainerImageFields clears deprecated image fields for non-container-image sources.
func clearLegacyContainerImageFields(d *schema.ResourceData) {
	_ = d.Set("image", "")
	_ = d.Set("image_digest", "")
}

// addLegacyContainerImageFieldsToMap adds compatibility image fields to function data source results.
func addLegacyContainerImageFieldsToMap(result map[string]interface{}, sourceDetails oci_functions.FunctionSourceDetails) {
	switch v := sourceDetails.(type) {
	case oci_functions.ContainerImageFunctionSourceDetails:
		if v.Image != nil {
			result["image"] = *v.Image
		}
		if v.ImageDigest != nil {
			result["image_digest"] = *v.ImageDigest
		} else {
			result["image_digest"] = requireRecompute
		}
	case oci_functions.UpdateContainerImageFunctionSourceDetails:
		if v.Image != nil {
			result["image"] = *v.Image
		}
		if v.ImageDigest != nil {
			result["image_digest"] = *v.ImageDigest
		} else {
			result["image_digest"] = requireRecompute
		}
	}
}

// Archive source details mappers
// mapToCreateArchiveSourceDetails converts Terraform archive source config into the create SDK model.
func (s *FunctionsFunctionResourceCrud) mapToCreateArchiveSourceDetails(fieldKeyFormat string) (oci_functions.CreateArchiveSourceDetails, error) {
	var baseObject oci_functions.CreateArchiveSourceDetails
	//discriminator
	archiveSourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "archive_source_type"))
	var archiveSourceType string
	if ok {
		archiveSourceType = archiveSourceTypeRaw.(string)
	} else {
		archiveSourceType = "" // default value
	}
	switch strings.ToLower(archiveSourceType) {
	case strings.ToLower("DIRECT_ARCHIVE"):
		details := oci_functions.CreateDirectArchiveSourceDetails{}
		if archiveFile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "archive_file")); ok {
			tmp := archiveFile.(string)
			decoded, err := base64.StdEncoding.DecodeString(tmp)
			if err != nil {
				return nil, fmt.Errorf("unable to base64 decode archive_file: %v", err)
			}
			details.ArchiveFile = decoded
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE_ARCHIVE"):
		details := oci_functions.CreateObjectStorageArchiveSourceDetails{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
			tmp := object.(string)
			details.ObjectName = &tmp
		}
		if objectVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_version_id")); ok {
			tmp := objectVersionId.(string)
			details.ObjectVersionId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown archive_source_type '%v' was specified", archiveSourceType)
	}
	return baseObject, nil
}

// mapToUpdateArchiveSourceDetails converts Terraform archive source config into the update SDK model.
func (s *FunctionsFunctionResourceCrud) mapToUpdateArchiveSourceDetails(fieldKeyFormat string) (oci_functions.UpdateArchiveSourceDetails, error) {
	var baseObject oci_functions.UpdateArchiveSourceDetails
	//discriminator
	archiveSourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "archive_source_type"))
	var archiveSourceType string
	if ok {
		archiveSourceType = archiveSourceTypeRaw.(string)
	} else {
		archiveSourceType = "" // default value
	}
	switch strings.ToLower(archiveSourceType) {
	case strings.ToLower("DIRECT_ARCHIVE"):
		details := oci_functions.UpdateDirectArchiveSourceDetails{}
		if archiveFile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "archive_file")); ok {
			tmp := archiveFile.(string)
			decoded, err := base64.StdEncoding.DecodeString(tmp)
			if err != nil {
				return nil, fmt.Errorf("unable to base64 decode archive_file: %v", err)
			}
			details.ArchiveFile = decoded
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE_ARCHIVE"):
		details := oci_functions.UpdateObjectStorageArchiveSourceDetails{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
			tmp := object.(string)
			details.ObjectName = &tmp
		}
		if objectVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_version_id")); ok {
			tmp := objectVersionId.(string)
			details.ObjectVersionId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown archive_source_type '%v' was specified", archiveSourceType)
	}
	return baseObject, nil
}

// ArchiveSourceDetailsToMap converts OCI archive source details back into Terraform state.
func ArchiveSourceDetailsToMap(obj *oci_functions.ArchiveSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_functions.DirectArchiveSourceDetails:
		result["archive_source_type"] = "DIRECT_ARCHIVE"
	case oci_functions.ObjectStorageArchiveSourceDetails:
		result["archive_source_type"] = "OBJECT_STORAGE_ARCHIVE"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}

		if v.ObjectVersionId != nil {
			result["object_version_id"] = string(*v.ObjectVersionId)
		}
	case oci_functions.UpdateDirectArchiveSourceDetails:
		result["archive_source_type"] = "DIRECT_ARCHIVE"
		if len(v.ArchiveFile) > 0 {
			result["archive_file"] = base64.StdEncoding.EncodeToString(v.ArchiveFile)
		}
	case oci_functions.UpdateObjectStorageArchiveSourceDetails:
		result["archive_source_type"] = "OBJECT_STORAGE_ARCHIVE"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}

		if v.ObjectVersionId != nil {
			result["object_version_id"] = string(*v.ObjectVersionId)
		}
	default:
		log.Printf("[WARN] Received 'archive_source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

// mapToCreateFunctionSourceDetails builds the correct create source_details model by source_type.
func (s *FunctionsFunctionResourceCrud) mapToCreateFunctionSourceDetails(fieldKeyFormat string) (oci_functions.CreateFunctionSourceDetails, error) {
	var baseObject oci_functions.CreateFunctionSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower(sourceTypeArchive):
		details := oci_functions.CreateArchiveFunctionSourceDetails{}
		if archiveSourceDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "archive_source_details")); ok {
			if tmpList := archiveSourceDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "archive_source_details"), 0)
				tmp, err := s.mapToCreateArchiveSourceDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert archive_source_details, encountered error: %v", err)
				}
				details.ArchiveSourceDetails = tmp
			}
		}
		if handler, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "handler")); ok {
			tmp := handler.(string)
			details.Handler = &tmp
		}
		if runtimeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "runtime_config")); ok {
			if tmpList := runtimeConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "runtime_config"), 0)
				tmp, err := s.mapToCreateRuntimeConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert runtime_config, encountered error: %v", err)
				}
				details.RuntimeConfig = tmp
			}
		}
		baseObject = details
	case strings.ToLower(sourceTypeContainerImage):
		details := oci_functions.CreateContainerImageFunctionSourceDetails{}
		image, imageDigest, ok := getContainerImageSourceValues(s.D, fieldKeyFormat)
		if !ok {
			return nil, fmt.Errorf("either 'source_details.image' or legacy 'image' must be set when source_type is CONTAINER_IMAGE")
		}
		if image != nil {
			details.Image = image
		}
		if imageDigest != nil {
			details.ImageDigest = imageDigest
		}
		baseObject = details
	case strings.ToLower(sourceTypePbf):
		details := oci_functions.CreatePreBuiltFunctionSourceDetails{}
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

// mapToUpdateFunctionSourceDetails builds the correct update source_details model by source_type.
func (s *FunctionsFunctionResourceCrud) mapToUpdateFunctionSourceDetails(fieldKeyFormat string) (oci_functions.UpdateFunctionSourceDetails, error) {
	var baseObject oci_functions.UpdateFunctionSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower(sourceTypeArchive):
		details := oci_functions.UpdateArchiveFunctionSourceDetails{}
		if archiveSourceDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "archive_source_details")); ok {
			if tmpList := archiveSourceDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "archive_source_details"), 0)
				tmp, err := s.mapToUpdateArchiveSourceDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert archive_source_details, encountered error: %v", err)
				}
				details.ArchiveSourceDetails = tmp
			}
		}
		if handler, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "handler")); ok {
			tmp := handler.(string)
			details.Handler = &tmp
		}
		if runtimeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "runtime_config")); ok {
			if tmpList := runtimeConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "runtime_config"), 0)
				tmp, err := s.mapToUpdateRuntimeConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert runtime_config, encountered error: %v", err)
				}
				details.RuntimeConfig = tmp
			}
		}
		baseObject = details
	case strings.ToLower(sourceTypeContainerImage):
		details := oci_functions.UpdateContainerImageFunctionSourceDetails{}
		image, imageDigest, ok := getContainerImageSourceValues(s.D, fieldKeyFormat)
		if !ok {
			return nil, fmt.Errorf("either 'source_details.image' or legacy 'image' must be set when source_type is CONTAINER_IMAGE")
		}
		if image != nil {
			details.Image = image
		}
		if imageDigest != nil {
			details.ImageDigest = imageDigest
		}
		baseObject = details
	case strings.ToLower(sourceTypePbf):
		// The Functions update API does not support PRE_BUILT_FUNCTIONS source_details.
		// Omit source_details for updates to PBF-backed functions so other attributes can still be updated.
		baseObject = nil
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

// FunctionSourceDetailsToMap converts OCI function source details back into Terraform state.
func FunctionSourceDetailsToMap(obj *oci_functions.FunctionSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_functions.ArchiveFunctionSourceDetails:
		result["source_type"] = "ARCHIVE"

		if v.ArchiveSourceDetails != nil {
			archiveSourceDetailsArray := []interface{}{}
			if archiveSourceDetailsMap := ArchiveSourceDetailsToMap(&v.ArchiveSourceDetails); archiveSourceDetailsMap != nil {
				archiveSourceDetailsArray = append(archiveSourceDetailsArray, archiveSourceDetailsMap)
			}
			result["archive_source_details"] = archiveSourceDetailsArray
		}

		if v.Handler != nil {
			result["handler"] = string(*v.Handler)
		}

		if v.RuntimeConfig != nil {
			runtimeConfigArray := []interface{}{}
			if runtimeConfigMap := RuntimeConfigToMap(&v.RuntimeConfig); runtimeConfigMap != nil {
				runtimeConfigArray = append(runtimeConfigArray, runtimeConfigMap)
			}
			result["runtime_config"] = runtimeConfigArray
		}

		if v.SourceCodeSha256 != nil {
			result["source_code_sha256"] = string(*v.SourceCodeSha256)
		}
	case oci_functions.ContainerImageFunctionSourceDetails:
		result["source_type"] = "CONTAINER_IMAGE"

		if v.Image != nil {
			result["image"] = string(*v.Image)
		}

		if v.ImageDigest != nil {
			result["image_digest"] = string(*v.ImageDigest)
		}
	case oci_functions.PreBuiltFunctionSourceDetails:
		result["source_type"] = "PRE_BUILT_FUNCTIONS"

		if v.PbfListingId != nil {
			result["pbf_listing_id"] = string(*v.PbfListingId)
		}
	case oci_functions.UpdateArchiveFunctionSourceDetails:
		result["source_type"] = "ARCHIVE"

		if v.ArchiveSourceDetails != nil {
			archiveSourceDetailsArray := []interface{}{}
			if archiveSourceDetailsMap := UpdateArchiveSourceDetailsToMap(&v.ArchiveSourceDetails); archiveSourceDetailsMap != nil {
				archiveSourceDetailsArray = append(archiveSourceDetailsArray, archiveSourceDetailsMap)
			}
			result["archive_source_details"] = archiveSourceDetailsArray
		}

		if v.Handler != nil {
			result["handler"] = string(*v.Handler)
		}

		if v.RuntimeConfig != nil {
			runtimeConfigArray := []interface{}{}
			if runtimeConfigMap := UpdateRuntimeConfigToMap(&v.RuntimeConfig); runtimeConfigMap != nil {
				runtimeConfigArray = append(runtimeConfigArray, runtimeConfigMap)
			}
			result["runtime_config"] = runtimeConfigArray
		}
	case oci_functions.UpdateContainerImageFunctionSourceDetails:
		result["source_type"] = "CONTAINER_IMAGE"

		if v.Image != nil {
			result["image"] = string(*v.Image)
		}

		if v.ImageDigest != nil {
			result["image_digest"] = string(*v.ImageDigest)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

// Runtime config mappers
// mapToCreateRuntimeConfig converts Terraform runtime_config into the create SDK model.
func (s *FunctionsFunctionResourceCrud) mapToCreateRuntimeConfig(fieldKeyFormat string) (oci_functions.CreateRuntimeConfig, error) {
	var baseObject oci_functions.CreateRuntimeConfig
	//discriminator
	runtimeConfigTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "runtime_config_type"))
	var runtimeConfigType string
	if ok {
		runtimeConfigType = runtimeConfigTypeRaw.(string)
	} else {
		runtimeConfigType = "" // default value
	}
	switch strings.ToLower(runtimeConfigType) {
	case strings.ToLower("FUNCTION_UPDATE"):
		details := oci_functions.CreateFunctionUpdateRuntimeConfig{}
		if functionsRuntimeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "functions_runtime_name")); ok {
			tmp := functionsRuntimeName.(string)
			details.FunctionsRuntimeName = &tmp
		}
		baseObject = details
	case strings.ToLower("MANUAL"):
		details := oci_functions.CreateManualRuntimeConfig{}
		if functionsRuntimeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "functions_runtime_name")); ok {
			tmp := functionsRuntimeName.(string)
			details.FunctionsRuntimeName = &tmp
		}
		if functionsRuntimeVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "functions_runtime_version_id")); ok {
			tmp := functionsRuntimeVersionId.(string)
			details.FunctionsRuntimeVersionId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown runtime_config_type '%v' was specified", runtimeConfigType)
	}
	return baseObject, nil
}

// mapToUpdateRuntimeConfig converts Terraform runtime_config into the update SDK model.
func (s *FunctionsFunctionResourceCrud) mapToUpdateRuntimeConfig(fieldKeyFormat string) (oci_functions.UpdateRuntimeConfig, error) {
	var baseObject oci_functions.UpdateRuntimeConfig
	//discriminator
	runtimeConfigTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "runtime_config_type"))
	var runtimeConfigType string
	if ok {
		runtimeConfigType = runtimeConfigTypeRaw.(string)
	} else {
		runtimeConfigType = "" // default value
	}
	switch strings.ToLower(runtimeConfigType) {
	case strings.ToLower("FUNCTION_UPDATE"):
		details := oci_functions.UpdateFunctionUpdateRuntimeConfig{}
		if functionsRuntimeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "functions_runtime_name")); ok {
			tmp := functionsRuntimeName.(string)
			details.FunctionsRuntimeName = &tmp
		}
		baseObject = details
	case strings.ToLower("MANUAL"):
		details := oci_functions.UpdateManualRuntimeConfig{}
		if functionsRuntimeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "functions_runtime_name")); ok {
			tmp := functionsRuntimeName.(string)
			details.FunctionsRuntimeName = &tmp
		}
		if functionsRuntimeVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "functions_runtime_version_id")); ok {
			tmp := functionsRuntimeVersionId.(string)
			details.FunctionsRuntimeVersionId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown runtime_config_type '%v' was specified", runtimeConfigType)
	}
	return baseObject, nil
}

// RuntimeConfigToMap converts OCI runtime config back into Terraform state.
func RuntimeConfigToMap(obj *oci_functions.RuntimeConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_functions.FunctionUpdateRuntimeConfig:
		result["runtime_config_type"] = "FUNCTION_UPDATE"

		if v.FunctionsRuntimeName != nil {
			result["functions_runtime_name"] = string(*v.FunctionsRuntimeName)
		}

		if v.FunctionsRuntimeVersionId != nil {
			result["functions_runtime_version_id"] = string(*v.FunctionsRuntimeVersionId)
		}
	case oci_functions.ManualRuntimeConfig:
		result["runtime_config_type"] = "MANUAL"

		if v.FunctionsRuntimeName != nil {
			result["functions_runtime_name"] = string(*v.FunctionsRuntimeName)
		}

		if v.FunctionsRuntimeVersionId != nil {
			result["functions_runtime_version_id"] = string(*v.FunctionsRuntimeVersionId)
		}
	case oci_functions.UpdateFunctionUpdateRuntimeConfig:
		result["runtime_config_type"] = "FUNCTION_UPDATE"

		if v.FunctionsRuntimeName != nil {
			result["functions_runtime_name"] = string(*v.FunctionsRuntimeName)
		}
	case oci_functions.UpdateManualRuntimeConfig:
		result["runtime_config_type"] = "MANUAL"

		if v.FunctionsRuntimeName != nil {
			result["functions_runtime_name"] = string(*v.FunctionsRuntimeName)
		}

		if v.FunctionsRuntimeVersionId != nil {
			result["functions_runtime_version_id"] = string(*v.FunctionsRuntimeVersionId)
		}
	default:
		log.Printf("[WARN] Received 'runtime_config_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FunctionsFunctionResourceCrud) mapToFailureDestinationDetails(fieldKeyFormat string) (oci_functions.FailureDestinationDetails, error) {
	var baseObject oci_functions.FailureDestinationDetails
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("NONE"):
		details := oci_functions.NoneFailureDestinationDetails{}
		baseObject = details
	case strings.ToLower("NOTIFICATION"):
		details := oci_functions.NotificationFailureDestinationDetails{}
		if topicId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "topic_id")); ok {
			tmp := topicId.(string)
			details.TopicId = &tmp
		}
		baseObject = details
	case strings.ToLower("QUEUE"):
		details := oci_functions.QueueFailureDestinationDetails{}
		if channelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "channel_id")); ok {
			tmp := channelId.(string)
			details.ChannelId = &tmp
		}
		if queueId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "queue_id")); ok {
			tmp := queueId.(string)
			details.QueueId = &tmp
		}
		baseObject = details
	case strings.ToLower("STREAM"):
		details := oci_functions.StreamFailureDestinationDetails{}
		if streamId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stream_id")); ok {
			tmp := streamId.(string)
			details.StreamId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func FailureDestinationDetailsToMap(obj *oci_functions.FailureDestinationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_functions.NoneFailureDestinationDetails:
		result["kind"] = "NONE"
	case oci_functions.NotificationFailureDestinationDetails:
		result["kind"] = "NOTIFICATION"

		if v.TopicId != nil {
			result["topic_id"] = string(*v.TopicId)
		}
	case oci_functions.QueueFailureDestinationDetails:
		result["kind"] = "QUEUE"

		if v.ChannelId != nil {
			result["channel_id"] = string(*v.ChannelId)
		}

		if v.QueueId != nil {
			result["queue_id"] = string(*v.QueueId)
		}
	case oci_functions.StreamFailureDestinationDetails:
		result["kind"] = "STREAM"

		if v.StreamId != nil {
			result["stream_id"] = string(*v.StreamId)
		}
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", *obj)
		return nil
	}

	return result
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

func (s *FunctionsFunctionResourceCrud) mapToSuccessDestinationDetails(fieldKeyFormat string) (oci_functions.SuccessDestinationDetails, error) {
	var baseObject oci_functions.SuccessDestinationDetails
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("NONE"):
		details := oci_functions.NoneSuccessDestinationDetails{}
		baseObject = details
	case strings.ToLower("NOTIFICATION"):
		details := oci_functions.NotificationSuccessDestinationDetails{}
		if topicId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "topic_id")); ok {
			tmp := topicId.(string)
			details.TopicId = &tmp
		}
		baseObject = details
	case strings.ToLower("QUEUE"):
		details := oci_functions.QueueSuccessDestinationDetails{}
		if channelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "channel_id")); ok {
			tmp := channelId.(string)
			details.ChannelId = &tmp
		}
		if queueId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "queue_id")); ok {
			tmp := queueId.(string)
			details.QueueId = &tmp
		}
		baseObject = details
	case strings.ToLower("STREAM"):
		details := oci_functions.StreamSuccessDestinationDetails{}
		if streamId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stream_id")); ok {
			tmp := streamId.(string)
			details.StreamId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func SuccessDestinationDetailsToMap(obj *oci_functions.SuccessDestinationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_functions.NoneSuccessDestinationDetails:
		result["kind"] = "NONE"
	case oci_functions.NotificationSuccessDestinationDetails:
		result["kind"] = "NOTIFICATION"

		if v.TopicId != nil {
			result["topic_id"] = string(*v.TopicId)
		}
	case oci_functions.QueueSuccessDestinationDetails:
		result["kind"] = "QUEUE"

		if v.ChannelId != nil {
			result["channel_id"] = string(*v.ChannelId)
		}

		if v.QueueId != nil {
			result["queue_id"] = string(*v.QueueId)
		}
	case oci_functions.StreamSuccessDestinationDetails:
		result["kind"] = "STREAM"

		if v.StreamId != nil {
			result["stream_id"] = string(*v.StreamId)
		}
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", *obj)
		return nil
	}

	return result
}

// UpdateArchiveSourceDetailsToMap converts update archive source models back into Terraform state.
func UpdateArchiveSourceDetailsToMap(obj *oci_functions.UpdateArchiveSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_functions.UpdateDirectArchiveSourceDetails:
		result["archive_source_type"] = "DIRECT_ARCHIVE"
		if len(v.ArchiveFile) > 0 {
			// Persist as base64 string to match schema expectation
			result["archive_file"] = base64.StdEncoding.EncodeToString(v.ArchiveFile)
		}
	case oci_functions.UpdateObjectStorageArchiveSourceDetails:
		result["archive_source_type"] = "OBJECT_STORAGE_ARCHIVE"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}

		if v.ObjectVersionId != nil {
			result["object_version_id"] = string(*v.ObjectVersionId)
		}
	default:
		log.Printf("[WARN] Received 'archive_source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

// UpdateRuntimeConfigToMap converts update runtime config models back into Terraform state.
func UpdateRuntimeConfigToMap(obj *oci_functions.UpdateRuntimeConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_functions.UpdateFunctionUpdateRuntimeConfig:
		result["runtime_config_type"] = "FUNCTION_UPDATE"

		if v.FunctionsRuntimeName != nil {
			result["functions_runtime_name"] = string(*v.FunctionsRuntimeName)
		}
	case oci_functions.UpdateManualRuntimeConfig:
		result["runtime_config_type"] = "MANUAL"

		if v.FunctionsRuntimeName != nil {
			result["functions_runtime_name"] = string(*v.FunctionsRuntimeName)
		}

		if v.FunctionsRuntimeVersionId != nil {
			result["functions_runtime_version_id"] = string(*v.FunctionsRuntimeVersionId)
		}
	default:
		log.Printf("[WARN] Received 'runtime_config_type' of unknown type %v", *obj)
		return nil
	}

	return result
}
