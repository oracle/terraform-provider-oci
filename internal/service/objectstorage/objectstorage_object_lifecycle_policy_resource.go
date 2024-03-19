// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"
)

const (
	multipartUploads = "multipart-uploads"
)

func ObjectStorageObjectLifecyclePolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createObjectStorageObjectLifecyclePolicy,
		Read:     readObjectStorageObjectLifecyclePolicy,
		Update:   updateObjectStorageObjectLifecyclePolicy,
		Delete:   deleteObjectStorageObjectLifecyclePolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"rules": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      rulesHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"action": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"time_amount": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"time_unit": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(oci_object_storage.ObjectLifecycleRuleTimeUnitDays),
								string(oci_object_storage.ObjectLifecycleRuleTimeUnitYears),
							}, false),
						},

						// Optional
						"object_name_filter": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"exclusion_patterns": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Set:      tfresource.LiteralTypeHashCodeForSets,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"inclusion_patterns": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Set:      tfresource.LiteralTypeHashCodeForSets,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"inclusion_prefixes": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Set:      tfresource.LiteralTypeHashCodeForSets,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
						},
						"target": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "objects",
						},

						// Computed
					},
				},
			},

			// Computed
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createObjectStorageObjectLifecyclePolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectLifecyclePolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.CreateResource(d, sync)
}

func readObjectStorageObjectLifecyclePolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectLifecyclePolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.ReadResource(sync)
}

func updateObjectStorageObjectLifecyclePolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectLifecyclePolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteObjectStorageObjectLifecyclePolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectLifecyclePolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ObjectStorageObjectLifecyclePolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_object_storage.ObjectStorageClient
	Res                    *oci_object_storage.ObjectLifecyclePolicy
	DisableNotFoundRetries bool
}

func (s *ObjectStorageObjectLifecyclePolicyResourceCrud) ID() string {
	return GetObjectLifecyclePolicyCompositeId(s.D.Get("bucket").(string), s.D.Get("namespace").(string))
}

func (s *ObjectStorageObjectLifecyclePolicyResourceCrud) Create() error {
	request := oci_object_storage.PutObjectLifecyclePolicyRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	// Adding empty list by default to compensate for service behavior
	request.Items = []oci_object_storage.ObjectLifecycleRule{}
	if rules, ok := s.D.GetOkExists("rules"); ok {
		set := rules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_object_storage.ObjectLifecycleRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := rulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
			converted, err := s.mapToObjectLifecycleRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.Items = tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.PutObjectLifecyclePolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ObjectLifecyclePolicy
	return nil
}

func (s *ObjectStorageObjectLifecyclePolicyResourceCrud) Get() error {
	request := oci_object_storage.GetObjectLifecyclePolicyRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	bucket, namespace, err := parseObjectLifecyclePolicyCompositeId(s.D.Id())
	if err == nil {
		request.BucketName = &bucket
		request.NamespaceName = &namespace
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.GetObjectLifecyclePolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ObjectLifecyclePolicy
	return nil
}

func (s *ObjectStorageObjectLifecyclePolicyResourceCrud) Update() error {
	request := oci_object_storage.PutObjectLifecyclePolicyRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.Items = []oci_object_storage.ObjectLifecycleRule{}
	if rules, ok := s.D.GetOkExists("rules"); ok {
		set := rules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_object_storage.ObjectLifecycleRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := rulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
			converted, err := s.mapToObjectLifecycleRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.Items = tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.PutObjectLifecyclePolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ObjectLifecyclePolicy
	return nil
}

func (s *ObjectStorageObjectLifecyclePolicyResourceCrud) Delete() error {
	request := oci_object_storage.DeleteObjectLifecyclePolicyRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	_, err := s.Client.DeleteObjectLifecyclePolicy(context.Background(), request)
	return err
}

func (s *ObjectStorageObjectLifecyclePolicyResourceCrud) SetData() error {

	bucket, namespace, err := parseObjectLifecyclePolicyCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("bucket", &bucket)
		s.D.Set("namespace", &namespace)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	rules := []interface{}{}
	for _, item := range s.Res.Items {
		rules = append(rules, ObjectLifecycleRuleToMap(item))
	}
	s.D.Set("rules", schema.NewSet(rulesHashCodeForSets, rules))

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func GetObjectLifecyclePolicyCompositeId(bucket string, namespace string) string {
	bucket = url.PathEscape(bucket)
	namespace = url.PathEscape(namespace)
	compositeId := "n/" + namespace + "/b/" + bucket + "/l"
	return compositeId
}

func parseObjectLifecyclePolicyCompositeId(compositeId string) (bucket string, namespace string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("n/.*/b/.*/l", compositeId)
	if !match || len(parts) != 5 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	bucket, _ = url.PathUnescape(parts[3])

	return
}

func (s *ObjectStorageObjectLifecyclePolicyResourceCrud) mapToObjectLifecycleRule(fieldKeyFormat string) (oci_object_storage.ObjectLifecycleRule, error) {
	result := oci_object_storage.ObjectLifecycleRule{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		tmp := action.(string)
		result.Action = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	isNotAMultipartUpload := true

	if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
		tmp := target.(string)
		result.Target = &tmp
		isNotAMultipartUpload = *(result.Target) != multipartUploads
	}

	if objectNameFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_name_filter")); ok && isNotAMultipartUpload {
		if tmpList := objectNameFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_name_filter"), 0)
			tmp, err := s.mapToObjectNameFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert object_name_filter, encountered error: %v", err)
			}
			result.ObjectNameFilter = &tmp
		}
	}

	if timeAmount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_amount")); ok {
		tmp := timeAmount.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert timeAmount string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.TimeAmount = &tmpInt64
	}

	if timeUnit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_unit")); ok {
		result.TimeUnit = oci_object_storage.ObjectLifecycleRuleTimeUnitEnum(timeUnit.(string))
	}

	return result, nil
}

func ObjectLifecycleRuleToMap(obj oci_object_storage.ObjectLifecycleRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Action != nil {
		result["action"] = string(*obj.Action)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectNameFilter != nil {
		result["object_name_filter"] = []interface{}{ObjectNameFilterToMap(obj.ObjectNameFilter)}
	}

	if obj.Target != nil {
		result["target"] = string(*obj.Target)
	}

	if obj.TimeAmount != nil {
		result["time_amount"] = strconv.FormatInt(*obj.TimeAmount, 10)
	}

	result["time_unit"] = string(obj.TimeUnit)

	return result
}

func (s *ObjectStorageObjectLifecyclePolicyResourceCrud) mapToObjectNameFilter(fieldKeyFormat string) (oci_object_storage.ObjectNameFilter, error) {
	result := oci_object_storage.ObjectNameFilter{}

	if exclusionPatterns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclusion_patterns")); ok {
		set := exclusionPatterns.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exclusion_patterns")) {
			result.ExclusionPatterns = tmp
		}
	}

	if inclusionPatterns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "inclusion_patterns")); ok {
		set := inclusionPatterns.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "inclusion_patterns")) {
			result.InclusionPatterns = tmp
		}
	}

	result.InclusionPrefixes = []string{}
	if inclusionPrefixes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "inclusion_prefixes")); ok {
		set := inclusionPrefixes.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		result.InclusionPrefixes = tmp
	}

	return result, nil
}

func ObjectNameFilterToMap(obj *oci_object_storage.ObjectNameFilter) map[string]interface{} {
	result := map[string]interface{}{}

	exclusionPatterns := []interface{}{}
	for _, item := range obj.ExclusionPatterns {
		exclusionPatterns = append(exclusionPatterns, item)
	}
	result["exclusion_patterns"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, exclusionPatterns)

	inclusionPatterns := []interface{}{}
	for _, item := range obj.InclusionPatterns {
		inclusionPatterns = append(inclusionPatterns, item)
	}
	result["inclusion_patterns"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, inclusionPatterns)

	inclusionPrefixes := []interface{}{}
	for _, item := range obj.InclusionPrefixes {
		inclusionPrefixes = append(inclusionPrefixes, item)
	}
	result["inclusion_prefixes"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, inclusionPrefixes)
	return result
}

func rulesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if action, ok := m["action"]; ok && action != "" {
		buf.WriteString(fmt.Sprintf("%v-", action))
	}
	if isEnabled, ok := m["is_enabled"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", isEnabled))
	}
	if name, ok := m["name"]; ok && name != "" {
		buf.WriteString(fmt.Sprintf("%v-", name))
	}
	if objectNameFilter, ok := m["object_name_filter"]; ok {
		if tmpList := objectNameFilter.([]interface{}); len(tmpList) > 0 {
			buf.WriteString("object_name_filter-")
			objectNameFilterRaw := tmpList[0].(map[string]interface{})
			if exclusionPatterns, ok := objectNameFilterRaw["exclusion_patterns"]; ok && exclusionPatterns != "" {
				buf.WriteString(fmt.Sprintf("exclusion_patterns-"))
				set := exclusionPatterns.(*schema.Set)
				exclusionPatternsArr := set.List()
				for _, exclusionPattern := range exclusionPatternsArr {
					buf.WriteString(fmt.Sprintf("%v-", exclusionPattern))
				}
			}
			if inclusionPatterns, ok := objectNameFilterRaw["inclusion_patterns"]; ok && inclusionPatterns != "" {
				buf.WriteString(fmt.Sprintf("inclusion_patterns-"))
				set := inclusionPatterns.(*schema.Set)
				inclusionPatternsArr := set.List()
				for _, inclusionPattern := range inclusionPatternsArr {
					buf.WriteString(fmt.Sprintf("%v-", inclusionPattern))
				}
			}
			if inclusionPrefixes, ok := objectNameFilterRaw["inclusion_prefixes"]; ok && inclusionPrefixes != "" {
				buf.WriteString(fmt.Sprintf("inclusionPrefix-"))
				set := inclusionPrefixes.(*schema.Set)
				inclusionPrefixesArr := set.List()
				for _, inclusionPrefix := range inclusionPrefixesArr {
					buf.WriteString(fmt.Sprintf("%v-", inclusionPrefix))
				}
			}
		}
	}
	if target, ok := m["target"]; ok && target != "" {
		buf.WriteString(fmt.Sprintf("%v-", target))
	}
	if timeAmount, ok := m["time_amount"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", timeAmount))
	}
	if timeUnit, ok := m["time_unit"]; ok && timeUnit != "" {
		buf.WriteString(fmt.Sprintf("%v-", timeUnit))
	}
	return utils.GetStringHashcode(buf.String())
}
