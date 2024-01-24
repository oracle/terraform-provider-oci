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
	"time"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"
)

func ObjectStorageBucketResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createObjectStorageBucket,
		Read:     readObjectStorageBucket,
		Update:   updateObjectStorageBucket,
		Delete:   deleteObjectStorageBucket,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			"access_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  string(oci_object_storage.CreateBucketDetailsPublicAccessTypeNopublicaccess),
			},
			"auto_tiering": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"object_events_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"storage_tier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"retention_rules": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 1,
				Set:      retentionRulesHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: tfresource.ValidateNotEmptyString(),
						},
						"duration": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"time_amount": {
										Type:             schema.TypeString,
										Required:         true,
										ValidateFunc:     tfresource.ValidateInt64TypeString,
										DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
									},
									"time_unit": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"time_rule_locked": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"retention_rule_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"versioning": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"approximate_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"approximate_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bucket_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_read_only": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"object_lifecycle_policy_etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createObjectStorageBucket(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageBucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.CreateResource(d, sync)
}

func readObjectStorageBucket(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageBucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.ReadResource(sync)
}

func updateObjectStorageBucket(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageBucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteObjectStorageBucket(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageBucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ObjectStorageBucketResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_object_storage.ObjectStorageClient
	Res                    *oci_object_storage.Bucket
	DisableNotFoundRetries bool
	RetentionRuleRes       []*oci_object_storage.RetentionRule
}

func (s *ObjectStorageBucketResourceCrud) ID() string {
	return GetBucketCompositeId(s.D.Get("name").(string), s.D.Get("namespace").(string))
}

func (s *ObjectStorageBucketResourceCrud) mapToDuration(fieldKeyFormat string) (oci_object_storage.Duration, error) {
	result := oci_object_storage.Duration{}

	if timeAmount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_amount")); ok {
		tmp := timeAmount.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert timeAmount string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.TimeAmount = &tmpInt64
	}

	if timeUnit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_unit")); ok {
		result.TimeUnit = oci_object_storage.DurationTimeUnitEnum(timeUnit.(string))
	}

	return result, nil
}

func (s *ObjectStorageBucketResourceCrud) Create() error {
	request := oci_object_storage.CreateBucketRequest{}

	if accessType, ok := s.D.GetOkExists("access_type"); ok {
		request.PublicAccessType = oci_object_storage.CreateBucketDetailsPublicAccessTypeEnum(accessType.(string))
	}

	if autoTiering, ok := s.D.GetOkExists("auto_tiering"); ok {
		request.AutoTiering = oci_object_storage.BucketAutoTieringEnum(autoTiering.(string))
	}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = resourceObjectStorageMapToMetadata(metadata.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if objectEventsEnabled, ok := s.D.GetOkExists("object_events_enabled"); ok {
		tmp := objectEventsEnabled.(bool)
		request.ObjectEventsEnabled = &tmp
	}

	if storageTier, ok := s.D.GetOkExists("storage_tier"); ok {
		request.StorageTier = oci_object_storage.CreateBucketDetailsStorageTierEnum(storageTier.(string))
	}

	if versioning, ok := s.D.GetOkExists("versioning"); ok {
		request.Versioning = oci_object_storage.CreateBucketDetailsVersioningEnum(versioning.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.CreateBucket(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Bucket

	if err := s.handleRetentionRules(); err != nil {
		log.Printf("[ERROR] Error in retention rule Create: '%v'", err)
		return err
	}

	return nil
}

func (s *ObjectStorageBucketResourceCrud) Get() error {
	request := oci_object_storage.GetBucketRequest{}
	listRetentionRulesRequest := oci_object_storage.ListRetentionRulesRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.BucketName = &tmp
		listRetentionRulesRequest.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
		listRetentionRulesRequest.NamespaceName = &tmp
	}

	bucket, namespace, err := parseBucketCompositeId(s.D.Id())
	if err == nil {
		request.BucketName = &bucket
		request.NamespaceName = &namespace
		listRetentionRulesRequest.BucketName = &bucket
		listRetentionRulesRequest.NamespaceName = &namespace
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.Fields = oci_object_storage.GetGetBucketFieldsEnumValues()
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.GetBucket(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Bucket

	// using list call as summary and get response is same for a retention rule
	listRetentionRulesRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")
	listRetentionRulesResponse, e := s.Client.ListRetentionRules(context.Background(), listRetentionRulesRequest)

	if e != nil {
		return e
	}

	s.RetentionRuleRes = listResponseToRetentionRuleRes(listRetentionRulesResponse)

	return nil
}

func listResponseToRetentionRuleRes(listRetentionRulesResponse oci_object_storage.ListRetentionRulesResponse) []*oci_object_storage.RetentionRule {
	RetentionRuleRes := make([]*oci_object_storage.RetentionRule, len(listRetentionRulesResponse.Items))

	for i, item := range listRetentionRulesResponse.Items {
		tmp := oci_object_storage.RetentionRule{}
		tmp.Id = item.Id
		tmp.DisplayName = item.DisplayName
		tmp.Duration = item.Duration
		tmp.TimeRuleLocked = item.TimeRuleLocked
		tmp.TimeModified = item.TimeModified
		tmp.TimeCreated = item.TimeCreated
		RetentionRuleRes[i] = &tmp
	}

	return RetentionRuleRes
}

func (s *ObjectStorageBucketResourceCrud) Update() error {
	request := oci_object_storage.UpdateBucketRequest{}

	if accessType, ok := s.D.GetOkExists("access_type"); ok {
		request.PublicAccessType = oci_object_storage.UpdateBucketDetailsPublicAccessTypeEnum(accessType.(string))
	}

	if autoTiering, ok := s.D.GetOkExists("auto_tiering"); ok {
		request.AutoTiering = oci_object_storage.BucketAutoTieringEnum(autoTiering.(string))
	}

	if bucket, ok := s.D.GetOkExists("name"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = resourceObjectStorageMapToMetadata(metadata.(map[string]interface{}))
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if objectEventsEnabled, ok := s.D.GetOkExists("object_events_enabled"); ok {
		tmp := objectEventsEnabled.(bool)
		request.ObjectEventsEnabled = &tmp
	}

	// @CODEGEN 2/2018: This should be used to change the name of a bucket, but the "namespace" field
	// is already being used to identify the bucket. Should have a new field for this.
	// Existing provider omits this, so we will omit it for now to avoid a potential breaking change.
	//if namespace, ok := s.D.GetOkExists("namespace"); ok {
	//	tmp := namespace.(string)
	//	request.Namespace = &tmp
	//}

	if versioning, ok := s.D.GetOkExists("versioning"); ok && s.D.HasChange("versioning") {
		request.Versioning = oci_object_storage.UpdateBucketDetailsVersioningEnum(versioning.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.UpdateBucket(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Bucket

	if err := s.handleRetentionRules(); err != nil {
		log.Printf("[ERROR] Error in retention rule handling: '%v'", err)
		return err
	}

	return nil
}

func (s *ObjectStorageBucketResourceCrud) Delete() error {
	request := oci_object_storage.DeleteBucketRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	_, err := s.Client.DeleteBucket(context.Background(), request)
	return err
}

func (s *ObjectStorageBucketResourceCrud) SetData() error {

	s.D.Set("bucket_id", *s.Res.Id)

	_, namespace, err := parseBucketCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("namespace", &namespace)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("access_type", s.Res.PublicAccessType)

	if s.Res.ApproximateCount != nil {
		s.D.Set("approximate_count", strconv.FormatInt(*s.Res.ApproximateCount, 10))
	}

	if s.Res.ApproximateSize != nil {
		s.D.Set("approximate_size", strconv.FormatInt(*s.Res.ApproximateSize, 10))
	}

	s.D.Set("auto_tiering", s.Res.AutoTiering)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Etag != nil {
		s.D.Set("etag", *s.Res.Etag)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsReadOnly != nil {
		s.D.Set("is_read_only", *s.Res.IsReadOnly)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", s.Res.Metadata)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.ObjectEventsEnabled != nil {
		s.D.Set("object_events_enabled", *s.Res.ObjectEventsEnabled)
	}

	if s.Res.ObjectLifecyclePolicyEtag != nil {
		s.D.Set("object_lifecycle_policy_etag", *s.Res.ObjectLifecyclePolicyEtag)
	}

	if s.Res.ReplicationEnabled != nil {
		s.D.Set("replication_enabled", *s.Res.ReplicationEnabled)
	}

	s.D.Set("storage_tier", s.Res.StorageTier)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("retention_rules", retentionRulesResToSet(s.RetentionRuleRes, false))

	s.D.Set("versioning", s.Res.Versioning)

	return nil
}

func retentionRulesResToSet(retentionRuleRes []*oci_object_storage.RetentionRule, datasource bool) interface{} {
	if retentionRuleRes != nil && len(retentionRuleRes) > 0 {
		retentionRules := []interface{}{}
		for _, item := range retentionRuleRes {
			retentionRules = append(retentionRules, RetentionRuleToMap(*item))
		}
		if datasource {
			return retentionRules
		} else {
			return schema.NewSet(retentionRulesHashCodeForSets, retentionRules)
		}
	} else {
		return nil
	}
}

func (s *ObjectStorageBucketResourceCrud) createRetentionRulesHelper(newRetentionRulesToCreate []interface{}) error {
	if len(newRetentionRulesToCreate) == 0 {
		return nil
	}

	createDetails := make([]oci_object_storage.RetentionRule, len(newRetentionRulesToCreate))
	for i, item := range newRetentionRulesToCreate {
		details, err := s.mapToRetentionRule(item.(map[string]interface{}))
		if err != nil {
			return err
		}
		createDetails[i] = details
	}

	createRequest := oci_object_storage.CreateRetentionRuleRequest{}
	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		createRequest.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		createRequest.NamespaceName = &tmp
	}

	responseList := []*oci_object_storage.RetentionRule{}

	for _, item := range createDetails {
		createRequest.DisplayName = item.DisplayName
		createRequest.Duration = item.Duration
		createRequest.TimeRuleLocked = item.TimeRuleLocked
		createRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

		createResponse, err := s.Client.CreateRetentionRule(context.Background(), createRequest)
		if err != nil {
			log.Printf("[ERROR] Failed to Create the retention rule '%s' : %v", *createRequest.DisplayName, err)
			s.RetentionRuleRes = responseList
			return err
		}
		s.RetentionRuleRes = append(s.RetentionRuleRes, &createResponse.RetentionRule)
	}

	return nil
}

func (s *ObjectStorageBucketResourceCrud) handleRetentionRules() error {

	defer func() {
		if err := s.Get(); err != nil {
			log.Printf("[ERROR] unable to invoke Get() '%v'", err)
		}
		if err := s.SetData(); err != nil {
			log.Printf("[ERROR] unable to invoke SetData() '%v'", err)
		}
	}()

	var newRetentionRulesToCreate, retentionRulesToUpdate, oldRetentionRulesToDelete []interface{}
	var err error = nil
	if newRetentionRulesToCreate, retentionRulesToUpdate, oldRetentionRulesToDelete, err = s.categorizeRetentionRules(); err != nil {
		log.Printf("[ERROR] Failed to categorize the retention rules: %v", err)
		return err
	}

	if err := s.createRetentionRulesHelper(newRetentionRulesToCreate); err != nil {
		return err
	}

	if err := s.updateRetentionRulesHelper(retentionRulesToUpdate); err != nil {
		return err
	}

	if err := s.deleteRetentionRulesHelper(oldRetentionRulesToDelete); err != nil {
		return err
	}

	return nil
}

func (s *ObjectStorageBucketResourceCrud) categorizeRetentionRules() ([]interface{}, []interface{}, []interface{}, error) {
	o, n := s.D.GetChange("retention_rules")
	if o == nil {
		o = new(schema.Set)
	}
	if n == nil {
		n = new(schema.Set)
	}

	orr := o.(*schema.Set).List()
	nrr := n.(*schema.Set).List()

	newDisplayNameToItemMap := map[string]interface{}{}
	oldDisplayNameToItemMap := map[string]interface{}{}

	for _, item := range nrr {
		tmp := item.(map[string]interface{})
		if displayName, ok := tmp["display_name"]; ok {
			if _, ok := newDisplayNameToItemMap[displayName.(string)]; ok {
				return nil, nil, nil, fmt.Errorf("[ERROR] display_name %s already taken", displayName.(string))
			}
			newDisplayNameToItemMap[displayName.(string)] = item
		}
	}

	for _, item := range orr {
		tmp := item.(map[string]interface{})
		if displayName, ok := tmp["display_name"]; ok {
			oldDisplayNameToItemMap[displayName.(string)] = item
		}
	}

	newRetentionRulesToCreate, retentionRulesToUpdate, oldRetentionRulesToDelete := []interface{}{}, []interface{}{}, []interface{}{}

	for displayName, itemNew := range newDisplayNameToItemMap {
		if itemOld, ok := oldDisplayNameToItemMap[displayName]; ok {
			hcNew := retentionRulesHashCodeForSets(itemNew)
			hcOld := retentionRulesHashCodeForSets(itemOld)
			if hcNew != hcOld {
				tmpOrr := itemOld.(map[string]interface{})
				tmpNrr := itemNew.(map[string]interface{})

				if tmp, ok := tmpOrr["retention_rule_id"]; ok {
					tmpNrr["retention_rule_id"] = tmp.(string)
				}

				retentionRulesToUpdate = append(retentionRulesToUpdate, itemNew)
			} else {
				unchangedRetentionRule, _ := s.mapToRetentionRule(itemOld.(map[string]interface{}))
				s.RetentionRuleRes = append(s.RetentionRuleRes, &unchangedRetentionRule)
			}
		} else {
			newRetentionRulesToCreate = append(newRetentionRulesToCreate, itemNew)
		}
	}

	for k := range oldDisplayNameToItemMap {
		if _, ok := newDisplayNameToItemMap[k]; !ok {
			oldRetentionRulesToDelete = append(oldRetentionRulesToDelete, oldDisplayNameToItemMap[k])
		}
	}

	return newRetentionRulesToCreate, retentionRulesToUpdate, oldRetentionRulesToDelete, nil
}

func (s *ObjectStorageBucketResourceCrud) updateRetentionRulesHelper(retentionRulesToUpdate []interface{}) error {
	if len(retentionRulesToUpdate) == 0 {
		return nil
	}

	updateDetails := make([]oci_object_storage.RetentionRule, len(retentionRulesToUpdate))
	for i, item := range retentionRulesToUpdate {
		details, err := s.mapToRetentionRule(item.(map[string]interface{}))
		if err != nil {
			return err
		}
		updateDetails[i] = details
	}

	updateRequest := oci_object_storage.UpdateRetentionRuleRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		updateRequest.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		updateRequest.NamespaceName = &tmp
	}

	responseList := []*oci_object_storage.RetentionRule{}

	for _, item := range updateDetails {
		updateRequest.DisplayName = item.DisplayName
		updateRequest.Duration = item.Duration
		updateRequest.TimeRuleLocked = item.TimeRuleLocked
		updateRequest.RetentionRuleId = item.Id

		updateResponse, err := s.Client.UpdateRetentionRule(context.Background(), updateRequest)
		if err != nil {
			log.Printf("[ERROR] Failed to Update the retention rule '%s' : %v", *updateRequest.DisplayName, err)
			s.RetentionRuleRes = responseList
			return err
		}
		s.RetentionRuleRes = append(s.RetentionRuleRes, &updateResponse.RetentionRule)
	}

	return nil
}

func (s *ObjectStorageBucketResourceCrud) deleteRetentionRulesHelper(oldRetentionRulesToDelete []interface{}) error {
	if len(oldRetentionRulesToDelete) == 0 {
		return nil
	}

	deleteIds := make([]string, len(oldRetentionRulesToDelete))
	for i, item := range oldRetentionRulesToDelete {
		tmp := item.(map[string]interface{})
		if retentionRuleId, ok := tmp["retention_rule_id"]; ok {
			deleteIds[i] = retentionRuleId.(string)
		}
	}

	deleteRequest := oci_object_storage.DeleteRetentionRuleRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		deleteRequest.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		deleteRequest.NamespaceName = &tmp
	}

	for _, retentionRuleId := range deleteIds {
		deleteRequest.RetentionRuleId = &retentionRuleId

		_, err := s.Client.DeleteRetentionRule(context.Background(), deleteRequest)
		if err != nil {
			log.Printf("Failed to delete the retention rule '%s' : %v", *deleteRequest.RetentionRuleId, err)
			e := tfresource.ReadResource(s)
			if e != nil {
				log.Printf("Failed to store the retention rules in the state file: %v", e)
			}
			return err
		}
	}

	return nil
}

func RetentionRuleToMap(obj oci_object_storage.RetentionRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["retention_rule_id"] = string(*obj.Id)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeModified != nil {
		result["time_modified"] = obj.TimeModified.String()
	}

	if obj.TimeRuleLocked != nil {
		result["time_rule_locked"] = obj.TimeRuleLocked.Format(time.RFC3339Nano)
	}

	if obj.Duration != nil {
		result["duration"] = []interface{}{DurationToMap(obj.Duration)}
	} else {
		result["duration"] = []interface{}{}
	}

	return result
}

func DurationToMap(obj *oci_object_storage.Duration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TimeAmount != nil {
		result["time_amount"] = strconv.FormatInt(*obj.TimeAmount, 10)
	}

	result["time_unit"] = string(obj.TimeUnit)

	return result
}

func GetBucketCompositeId(bucket string, namespace string) string {
	bucket = url.PathEscape(bucket)
	namespace = url.PathEscape(namespace)
	compositeId := "n/" + namespace + "/b/" + bucket
	return compositeId
}

func parseBucketCompositeId(compositeId string) (bucket string, namespace string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("n/.*/b/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	bucket, _ = url.PathUnescape(parts[3])

	return
}

func (s *ObjectStorageBucketResourceCrud) mapToRetentionRule(retentionRule map[string]interface{}) (oci_object_storage.RetentionRule, error) {
	details := oci_object_storage.RetentionRule{}

	if displayName, ok := retentionRule["display_name"]; ok {
		tmp := displayName.(string)
		details.DisplayName = &tmp
	}

	if duration, ok := retentionRule["duration"]; ok {
		durationReq := oci_object_storage.Duration{}

		if tmpList := duration.([]interface{}); len(tmpList) > 0 {
			durationRaw := tmpList[0].(map[string]interface{})
			if timeAmount, ok := durationRaw["time_amount"]; ok && timeAmount != "" {
				tmp := timeAmount.(string)
				tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
				if err != nil {
					return details, fmt.Errorf("unable to convert timeAmount string: %s to an int64 and encountered error: %v", tmp, err)
				}
				durationReq.TimeAmount = &tmpInt64
			}

			if timeUnit, ok := durationRaw["time_unit"]; ok && timeUnit != "" {
				durationReq.TimeUnit = oci_object_storage.DurationTimeUnitEnum(timeUnit.(string))
			}
		}
		details.Duration = &durationReq
	}

	if timeRuleLocked, ok := retentionRule["time_rule_locked"]; ok && timeRuleLocked != "" {
		tmp, err := time.Parse(time.RFC3339, timeRuleLocked.(string))
		if err != nil {
			return details, err
		}
		details.TimeRuleLocked = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreated, ok := retentionRule["time_created"]; ok && timeCreated != "" {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return details, err
		}
		details.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if timeModified, ok := retentionRule["time_modified"]; ok && timeModified != "" {
		tmp, err := time.Parse(time.RFC3339, timeModified.(string))
		if err != nil {
			return details, err
		}
		details.TimeModified = &oci_common.SDKTime{Time: tmp}
	}

	if tmp, ok := retentionRule["retention_rule_id"]; ok {
		retenionRuleId := tmp.(string)
		details.Id = &retenionRuleId
	}

	return details, nil
}

func retentionRulesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	if displayName, ok := m["display_name"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", displayName))
	}

	if duration, ok := m["duration"]; ok {
		if tmpList := duration.([]interface{}); len(tmpList) > 0 {
			buf.WriteString("duration-")
			durationRaw := tmpList[0].(map[string]interface{})
			if type_, ok := durationRaw["time_amount"]; ok && type_ != "" {
				buf.WriteString(fmt.Sprintf("%v-", type_))
			}
			if type_, ok := durationRaw["time_unit"]; ok && type_ != "" {
				buf.WriteString(fmt.Sprintf("%v-", type_))
			}
		}
	}

	if timeRuleLocked, ok := m["time_rule_locked"]; ok && timeRuleLocked != "" {
		buf.WriteString(fmt.Sprintf("%v-", timeRuleLocked))
	}

	return utils.GetStringHashcode(buf.String())
}
