// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"
)

const (
	ObjIdDelim  = "/"
	ObjIdPrefix = "tfobm-object-"
)

func ObjectStorageObjectResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createObjectStorageObject,
		Read:     readObjectStorageObject,
		Update:   updateObjectStorageObject,
		Delete:   deleteObjectStorageObject,
		Schema: map[string]*schema.Schema{
			// @CODEGEN 2/2018:
			// Previous provider doesn't provide an Update method and sets all non-Computed fields to ForceNew.
			// This was done because the same PutObject() call can be used to Create and modify existing objects.
			//
			// For generated code, we removed the ForceNew attribute from non-Computed fields and added
			// an Update method which calls the Create() method. This should have the same effect as the
			// previous behavior; while minimizing diffs between this and the generated code.

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
			"object": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"cache_control": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"content_disposition": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"content": {
				Type: schema.TypeString,
				// @CODEGEN 2/2018: content is optional and stored as checksum to avoid bloating the state file
				// Generator was setting it as required.
				Optional: true,
				ForceNew: true,
				StateFunc: func(body interface{}) string {
					v := body.(string)
					if v == "" {
						return ""
					}
					h := md5.Sum([]byte(v))
					return hex.EncodeToString(h[:])
				},
				ConflictsWith: []string{"source", "source_uri_details"},
			},
			"content_encoding": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"content_language": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"content_length": {
				Type: schema.TypeString,
				// @CODEGEN 2/2018: this was generated as Required, we will compute the length from the 'content'
				Computed: true,
			},
			"content_md5": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if !tfresource.IsHex(new) {
						return old == new
					}

					base64, err := tfresource.HexToB64(new)
					if err != nil {
						return false
					}
					return old == *base64
				},
			},
			"content_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"delete_all_object_versions": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"metadata": {
				// @CODEGEN 2/2018: This should be a map[string]string. Spec doesn't specify this correctly and
				// generates it as a TypeString
				Type:         schema.TypeMap,
				Elem:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateLowerCaseKeysInMetadata,
			},
			"storage_tier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"opc_sse_kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"content", "source_uri_details"},
				StateFunc:     tfresource.GetSourceFileState,
				ValidateFunc:  tfresource.ValidateSourceValue,
			},
			"source_uri_details": {
				Type:          schema.TypeList,
				Optional:      true,
				ForceNew:      true,
				MaxItems:      1,
				MinItems:      1,
				ConflictsWith: []string{"content", "source"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"region": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						// Required
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						// Required
						"bucket": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						// Required
						"object": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"source_object_if_match_etag": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"destination_object_if_match_etag": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"destination_object_if_none_match_etag": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"source_version_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},

			// Computed
			// @CODEGEN 12/20/2018 - Even though Object resource is not stateful for content and multi-part variations
			// making those variations stateful to match the logic for copy case to ensure that provider does not fail during state polling due to missing state property
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"work_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createObjectStorageObject(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.CreateResource(d, sync)
}

func (s *ObjectStorageObjectResourceCrud) createMultiPartObject() error {
	multipartUploadData := MultipartUploadData{}

	source, ok := s.D.GetOkExists("source")
	if !ok {
		return fmt.Errorf("the source is not specified to Create multipart upload")
	}
	tmpSource := source.(string)
	sourceInfo, err := os.Stat(tmpSource)
	if err != nil {
		return fmt.Errorf("the specified source is not available: %q", err)
	}

	multipartUploadData.SourcePath = &tmpSource
	multipartUploadData.SourceInfo = &sourceInfo

	if cacheControl, ok := s.D.GetOkExists("cache_control"); ok {
		tmp := cacheControl.(string)
		multipartUploadData.CacheControl = &tmp
	}

	if contentDisposition, ok := s.D.GetOkExists("content_disposition"); ok {
		tmp := contentDisposition.(string)
		multipartUploadData.ContentDisposition = &tmp
	}

	if contentEncoding, ok := s.D.GetOkExists("content_encoding"); ok {
		tmp := contentEncoding.(string)
		multipartUploadData.ContentEncoding = &tmp
	}

	if contentLanguage, ok := s.D.GetOkExists("content_language"); ok {
		tmp := contentLanguage.(string)
		multipartUploadData.ContentLanguage = &tmp
	}

	if contentType, ok := s.D.GetOkExists("content_type"); ok {
		tmp := contentType.(string)
		multipartUploadData.ContentType = &tmp
	}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		multipartUploadData.BucketName = &tmp
	}

	if storageTier, ok := s.D.GetOkExists("storage_tier"); ok {
		tmp := storageTier.(string)
		multipartUploadData.StorageTier = StorageTierEnumFromString(tmp)
	}

	if opcSseKmsKeyId, ok := s.D.GetOkExists("opc_sse_kms_key_id"); ok {
		tmp := opcSseKmsKeyId.(string)
		multipartUploadData.OpcSseKmsKeyId = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		multipartUploadData.Metadata = metadata.(map[string]interface{})
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		multipartUploadData.NamespaceName = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		multipartUploadData.ObjectName = &tmp
	}

	multipartUploadData.ObjectStorageClient = s.Client
	multipartUploadData.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	s.D.Set("work_request_id", "")
	s.D.Set("state", oci_object_storage.WorkRequestStatusInProgress)

	id, multipartInitErr := MultiPartUpload(multipartUploadData)
	if multipartInitErr != nil {
		return multipartInitErr
	}

	s.D.SetId(id)
	s.D.Set("state", oci_object_storage.WorkRequestStatusCompleted)

	return s.Get()
}

func (s *ObjectStorageObjectResourceCrud) createCopyObject() error {

	copyObjectRequest := oci_object_storage.CopyObjectRequest{}

	configProvider := *s.Client.ConfigurationProvider()
	if configProvider == nil {
		return fmt.Errorf("cannot access ConfigurationProvider")
	}
	currentRegion, error := configProvider.Region()
	if error != nil {
		return fmt.Errorf("cannot access Region for the current ConfigurationProvider")
	}

	if sourceURI, ok := s.D.GetOkExists("source_uri_details"); ok && sourceURI != nil {
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_uri_details", 0)

		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			err := s.createSourceRegionClient(tmp)
			if err != nil {
				return err
			}
		}
		copyObjectRequest.DestinationRegion = &currentRegion

		if sourceNamespaceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := sourceNamespaceName.(string)
			copyObjectRequest.NamespaceName = &tmp
		}

		if sourceBucketName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := sourceBucketName.(string)
			copyObjectRequest.BucketName = &tmp
		}

		if sourceObjectName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
			tmp := sourceObjectName.(string)
			copyObjectRequest.SourceObjectName = &tmp
		}

		if sourceObjectIfMatchETag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_object_if_match_etag")); ok {
			tmp := sourceObjectIfMatchETag.(string)
			copyObjectRequest.SourceObjectIfMatchETag = &tmp
		}

		if destinationObjectIfMatchETag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_object_if_match_etag")); ok {
			tmp := destinationObjectIfMatchETag.(string)
			copyObjectRequest.DestinationObjectIfMatchETag = &tmp
		}

		if destinationObjectIfNoneMatchETag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_object_if_none_match_etag")); ok {
			tmp := destinationObjectIfNoneMatchETag.(string)
			copyObjectRequest.DestinationObjectIfNoneMatchETag = &tmp
		}

		if sourceVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_version_id")); ok {
			tmp := sourceVersionId.(string)
			copyObjectRequest.SourceVersionId = &tmp
		}
	}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		copyObjectRequest.DestinationBucket = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		copyObjectRequest.DestinationObjectMetadata = resourceObjectStorageMapToOPCMetadata(metadata.(map[string]interface{}))
	}

	if storageTier, ok := s.D.GetOkExists("storage_tier"); ok {
		tmp := storageTier.(string)
		copyObjectRequest.DestinationObjectStorageTier = StorageTierEnumFromString(tmp)
	}

	if opcSseKmsKeyId, ok := s.D.GetOkExists("opc_sse_kms_key_id"); ok {
		tmp := opcSseKmsKeyId.(string)
		copyObjectRequest.OpcSseKmsKeyId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		copyObjectRequest.DestinationNamespace = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		copyObjectRequest.DestinationObjectName = &tmp
	}

	copyObjectRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	var workRequestId = ""
	if state, ok := s.D.GetOkExists("state"); ok {
		if state == oci_object_storage.WorkRequestStatusInProgress {
			workRequestIdStateValue := s.D.Get("work_request_id")
			workRequestId = workRequestIdStateValue.(string)
		}
	}

	if workRequestId == "" {
		copyObjectResponse, err := s.SourceRegionClient.CopyObject(context.Background(), copyObjectRequest)
		if err != nil {
			s.D.Set("state", string(oci_object_storage.WorkRequestStatusCanceled))
			return err
		}
		workRequestId = *copyObjectResponse.OpcWorkRequestId
	}

	s.D.Set("work_request_id", workRequestId)
	s.D.Set("state", string(oci_object_storage.WorkRequestStatusInProgress))

	getWorkRequestRequest := oci_object_storage.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = &workRequestId
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")
	workRequestResponse, err := s.SourceRegionClient.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest

	copyTimeout := s.D.Timeout(schema.TimeoutCreate)
	err = copyObjectWaitForWorkRequest(&workRequestId, "object", copyTimeout, s.DisableNotFoundRetries, s.SourceRegionClient)

	if err != nil {
		// we are not able to verify the state of workRequest
		s.D.Set("state", string(oci_object_storage.WorkRequestStatusFailed))
		return err
	}

	s.D.Set("work_request_id", "")
	s.D.Set("state", string(oci_object_storage.WorkRequestStatusCompleted))
	id := GetObjectCompositeId(*copyObjectRequest.DestinationBucket, *copyObjectRequest.DestinationNamespace, *copyObjectRequest.DestinationObjectName)
	s.D.SetId(id)
	return s.Get()
}

func readObjectStorageObject(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectResourceCrud{}
	// For backward compatibility with CompositeId change
	log.Printf("[DEBUG] readObjectStorageObject() Resource Id in state: %s", d.Id())
	_, _, _, err := parseObjectCompositeId(d.Id())

	if err != nil {
		bucket, bOk := d.GetOkExists("bucket")
		namespace, nOk := d.GetOkExists("namespace")
		object, oOk := d.GetOkExists("object")

		if bOk && nOk && oOk {
			compositeId := GetObjectCompositeId(bucket.(string), namespace.(string), object.(string))
			d.SetId(compositeId)
		}
	}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.ReadResource(sync)
}

func updateObjectStorageObject(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteObjectStorageObject(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

// There's no struct to represent this in SDK, so we define our own including a fake LifecycleState
type ObjectStorageObject struct {
	NamespaceName      string
	BucketName         string
	ObjectName         string
	HeadObjectResponse oci_object_storage.HeadObjectResponse
	ObjectResponse     oci_object_storage.GetObjectResponse
	LifecycleState     string
}

type ObjectStorageObjectResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_object_storage.ObjectStorageClient
	SourceRegionClient     *oci_object_storage.ObjectStorageClient
	Res                    *ObjectStorageObject
	DisableNotFoundRetries bool
	WorkRequest            *oci_object_storage.WorkRequest
}

func (s *ObjectStorageObjectResourceCrud) ID() string {
	return GetObjectCompositeId(s.D.Get("bucket").(string), s.D.Get("namespace").(string), s.D.Get("object").(string))
}

func (s *ObjectStorageObjectResourceCrud) Create() error {

	if s.isCopyCreate() {
		return s.createCopyObject()
	}

	if s.isMultiPartCreate() {
		return s.createMultiPartObject()
	}

	return s.createContentObject()
}

func (s *ObjectStorageObjectResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_object_storage.WorkRequestStatusAccepted),
		string(oci_object_storage.WorkRequestStatusInProgress),
		string(oci_object_storage.WorkRequestStatusCanceling),
	}
}

func (s *ObjectStorageObjectResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_object_storage.WorkRequestSummaryStatusCompleted),
		string(oci_object_storage.WorkRequestSummaryStatusCanceled),
		string(oci_object_storage.WorkRequestStatusFailed),
	}
}

func (s *ObjectStorageObjectResourceCrud) createContentObject() error {
	request := oci_object_storage.PutObjectRequest{}

	if cacheControl, ok := s.D.GetOkExists("cache_control"); ok {
		tmp := cacheControl.(string)
		request.CacheControl = &tmp
	}

	if contentDisposition, ok := s.D.GetOkExists("content_disposition"); ok {
		tmp := contentDisposition.(string)
		request.ContentDisposition = &tmp
	}

	if contentEncoding, ok := s.D.GetOkExists("content_encoding"); ok {
		tmp := contentEncoding.(string)
		request.ContentEncoding = &tmp
	}

	if contentLanguage, ok := s.D.GetOkExists("content_language"); ok {
		tmp := contentLanguage.(string)
		request.ContentLanguage = &tmp
	}

	// @CODEGEN 2/2018: Generator code allow you to set the content_length
	// from the schema. These are omitted from the existing provider
	// resource because they can be computed.

	if contentMd5, ok := s.D.GetOkExists("content_md5"); ok {
		tmp := contentMd5.(string)

		if tfresource.IsHex(tmp) {
			contentMd5B64, err := tfresource.HexToB64(tmp)
			if err != nil {
				return err
			}
			request.ContentMD5 = contentMd5B64
		} else {
			request.ContentMD5 = &tmp
		}
	}

	if contentType, ok := s.D.GetOkExists("content_type"); ok {
		tmp := contentType.(string)
		request.ContentType = &tmp
	}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if content, ok := s.D.GetOkExists("content"); ok {
		// @CODEGEN 2/2018: The generator doesn't yet handle strings that should be converted to byte arrays.
		tmp := []byte(content.(string))
		tmpLength := int64(len(tmp))
		request.ContentLength = &tmpLength
		if tmpLength == 0 {
			request.PutObjectBody = http.NoBody
		} else {
			request.PutObjectBody = ioutil.NopCloser(bytes.NewBuffer(tmp))
		}
	} else {
		tmp := int64(0)
		request.ContentLength = &tmp
		request.PutObjectBody = http.NoBody
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.OpcMeta = resourceObjectStorageMapToMetadata(metadata.(map[string]interface{}))
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	if storageTier, ok := s.D.GetOkExists("storage_tier"); ok {
		tmp := storageTier.(string)
		request.StorageTier = PutObjectStorageTierEnumFromString(tmp)
	}

	if opcSseKmsKeyId, ok := s.D.GetOkExists("opc_sse_kms_key_id"); ok {
		tmp := opcSseKmsKeyId.(string)
		request.OpcSseKmsKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	s.D.Set("work_request_id", "")
	s.D.Set("state", oci_object_storage.WorkRequestStatusInProgress)

	_, err := s.Client.PutObject(context.Background(), request)
	if err != nil {
		return err
	}

	id := GetObjectCompositeId(*request.BucketName, *request.NamespaceName, *request.ObjectName)
	s.D.SetId(id)

	// @CODEGEN 2/2018: PutObject() call doesn't return an object. Instead, use existing
	// Get() implementation to retrieve the state of the object and set its state to completed
	s.D.Set("state", oci_object_storage.WorkRequestStatusCompleted)

	return s.Get()
}

func (s *ObjectStorageObjectResourceCrud) getObjectHead() error {

	headObjectRequest := &oci_object_storage.HeadObjectRequest{}

	bucket, namespace, object, err := parseObjectCompositeId(s.D.Id())

	if err == nil {
		headObjectRequest.BucketName = &bucket
		headObjectRequest.NamespaceName = &namespace
		headObjectRequest.ObjectName = &object
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	headObjectRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	headObjectResponse, err := s.Client.HeadObject(context.Background(), *headObjectRequest)
	if err != nil {
		return err
	}

	s.Res = &ObjectStorageObject{
		NamespaceName:      *headObjectRequest.NamespaceName,
		BucketName:         *headObjectRequest.BucketName,
		ObjectName:         *headObjectRequest.ObjectName,
		HeadObjectResponse: headObjectResponse,
		LifecycleState:     s.D.Get("state").(string),
	}

	return nil
}

func (s *ObjectStorageObjectResourceCrud) updateState() (bool, error) {
	if state, ok := s.D.GetOkExists("state"); ok {
		if state == oci_object_storage.WorkRequestStatusInProgress {

			if wrid, ok := s.D.GetOkExists("work_request_id"); ok {
				retryPolicy := tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")
				retryPolicy.ShouldRetryOperation = objectStorageWorkRequestShouldRetryFunc(s.D.Timeout(schema.TimeoutCreate))

				getWorkRequestRequest := oci_object_storage.GetWorkRequestRequest{}
				wridStr := wrid.(string)
				getWorkRequestRequest.WorkRequestId = &wridStr
				getWorkRequestRequest.RequestMetadata.RetryPolicy = retryPolicy

				if sourceURI, ok := s.D.GetOkExists("source_uri_details"); ok && sourceURI != nil {
					fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_uri_details", 0)
					// the region should exist for source_uri_details
					if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
						tmp := region.(string)
						err := s.createSourceRegionClient(tmp)
						if err != nil {
							return false, err
						}
					}
				} else {
					return false, fmt.Errorf("no source_uri_details specified to verify copy state by WorkRequest")
				}

				workRequestResponse, err := s.SourceRegionClient.GetWorkRequest(context.Background(), getWorkRequestRequest)
				if err != nil {
					return false, err
				}

				wr := &workRequestResponse.WorkRequest
				s.D.Set("state", string(wr.Status))

				if wr.Status == oci_object_storage.WorkRequestStatusInProgress {
					return false, nil
				}

				s.D.Set("work_request_id", "")
				return true, nil

			}

			return false, fmt.Errorf("the state is incorrect. no work_request_id found for the InProgress State")
		}

		return true, nil
	}

	return true, nil
}

func (s *ObjectStorageObjectResourceCrud) shouldUseObjectHeadForGet() bool {
	content, _ := s.D.GetOkExists("content")
	return content == ""
}

func (s *ObjectStorageObjectResourceCrud) isMultiPartCreate() bool {
	source, _ := s.D.GetOkExists("source")
	return source != ""
}

func (s *ObjectStorageObjectResourceCrud) isCopyCreate() bool {
	if sourceURI, ok := s.D.GetOkExists("source_uri_details"); ok {
		if tmpList := sourceURI.([]interface{}); len(tmpList) > 0 {
			return true
		}
	}
	return false
}

func (s *ObjectStorageObjectResourceCrud) Get() error {

	workRequestFinished, err := s.updateState()
	if err != nil {
		return err
	}

	if !workRequestFinished {
		// old object can exist
		return nil
	}

	if s.shouldUseObjectHeadForGet() {
		return s.getObjectHead()
	}

	return s.getObject()
}

func (s *ObjectStorageObjectResourceCrud) getObject() error {
	request := oci_object_storage.GetObjectRequest{}

	bucketName, namespaceName, objectName, err := parseObjectCompositeId(s.D.Id())

	if err == nil {
		request.NamespaceName = &namespaceName
		request.BucketName = &bucketName
		request.ObjectName = &objectName
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	// TODO: May be better to use HeadObject() to retrieve status of the object. For large content, doesn't make sense
	// to call Get() all the time
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.GetObject(context.Background(), request)
	if err != nil {
		return err
	}

	// @CODEGEN 2/2018: We must store the response along with the identifiers that aren't
	// returned in the GetResponse.
	s.Res = &ObjectStorageObject{
		ObjectResponse: response,
		NamespaceName:  *request.NamespaceName,
		BucketName:     *request.BucketName,
		ObjectName:     *request.ObjectName,
		LifecycleState: s.D.Get("state").(string),
	}

	return nil
}

func (s *ObjectStorageObjectResourceCrud) Update() error {

	// @CODEGEN 06/2018: Update is only supported for the change in name - all others are a forceNew
	if s.D.HasChange("object") {

		request := oci_object_storage.RenameObjectRequest{}

		if bucket, ok := s.D.GetOkExists("bucket"); ok {
			tmp := bucket.(string)
			request.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			tmp := namespace.(string)
			request.NamespaceName = &tmp
		}
		oldRaw, newRaw := s.D.GetChange("object")
		sourceName := oldRaw.(string)
		request.SourceName = &sourceName

		newName := newRaw.(string)
		request.NewName = &newName

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")
		_, err := s.Client.RenameObject(context.Background(), request)
		if err != nil {
			return err
		}

		updatedId := GetObjectCompositeId(*request.BucketName, *request.NamespaceName, *request.NewName)
		s.D.SetId(updatedId)
	}
	return s.Get()
}

func (s *ObjectStorageObjectResourceCrud) Delete() error {
	request := oci_object_storage.DeleteObjectRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	if deleteAllObjectVersions, ok := s.D.GetOkExists("delete_all_object_versions"); ok && deleteAllObjectVersions.(bool) {
		return DeleteAllObjectVersions(s.Client, *request.BucketName, *request.NamespaceName, *request.ObjectName)
	} else {
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

		_, err := s.Client.DeleteObject(context.Background(), request)
		return err
	}
}

func (s *ObjectStorageObjectResourceCrud) SetData() error {
	if s.shouldUseObjectHeadForGet() {
		return s.setDataObjectHead()
	}

	return s.setDataObject()
}

func (s *ObjectStorageObjectResourceCrud) setDataObjectHead() error {
	s.D.Set("namespace", s.Res.NamespaceName)
	s.D.Set("bucket", s.Res.BucketName)
	s.D.Set("object", s.Res.ObjectName)

	response := s.Res.HeadObjectResponse

	if response.CacheControl != nil {
		s.D.Set("cache_control", *response.CacheControl)
	}

	if response.ContentDisposition != nil {
		s.D.Set("content_disposition", *response.ContentDisposition)
	}

	if response.ContentEncoding != nil {
		s.D.Set("content_encoding", *response.ContentEncoding)
	}

	if response.ContentLanguage != nil {
		s.D.Set("content_language", *response.ContentLanguage)
	}

	if response.ContentLength != nil {
		s.D.Set("content_length", strconv.FormatInt(*response.ContentLength, 10))
	}

	if response.OpcMultipartMd5 != nil {
		s.D.Set("content_md5", *response.OpcMultipartMd5)
	}

	if response.ContentMd5 != nil {
		s.D.Set("content_md5", *response.ContentMd5)
	}

	if response.ContentType != nil {
		s.D.Set("content_type", *response.ContentType)
	}

	if response.VersionId != nil {
		s.D.Set("version_id", *response.VersionId)
	}

	s.D.Set("storage_tier", string(response.StorageTier))

	if response.OpcMeta != nil {
		if err := s.D.Set("metadata", response.OpcMeta); err != nil {
			log.Printf("Unable to set 'metadata'. Error: %q", err)
		}
	}

	return nil
}

// @CODEGEN 2/2018: The existing provider returns a custom Id in following format:
// "tfobm-object-<namespace_name>/<bucket_name>/<object_name>"
// Update - Id format updated to "n/tfobm-object-<namespace_name>/b/<bucket_name>/o/<object_name>"
func GetObjectCompositeId(bucket string, namespace string, object string) string {
	bucket = url.PathEscape(bucket)
	namespace = url.PathEscape(namespace)
	object = url.PathEscape(object)
	compositeId := "n/" + namespace + "/b/" + bucket + "/o/" + object
	return compositeId
}

func parseObjectCompositeId(compositeId string) (bucket string, namespace string, object string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("n/.*/b/.*/o/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	bucket, _ = url.PathUnescape(parts[3])
	object, _ = url.PathUnescape(parts[5])

	return
}

func (s *ObjectStorageObjectResourceCrud) setDataObject() error {
	s.D.Set("namespace", s.Res.NamespaceName)
	s.D.Set("bucket", s.Res.BucketName)
	s.D.Set("object", s.Res.ObjectName)

	contentReader := s.Res.ObjectResponse.Content
	if contentReader != nil {
		contentArray, err := ioutil.ReadAll(contentReader)
		if err != nil {
			log.Printf("Unable to read 'content' from response. Error: %q", err)
			return err
		}
		h := md5.Sum(contentArray)
		s.D.Set("content", hex.EncodeToString(h[:]))
	}

	if s.Res.ObjectResponse.CacheControl != nil {
		s.D.Set("cache_control", *s.Res.ObjectResponse.CacheControl)
	}

	if s.Res.ObjectResponse.ContentDisposition != nil {
		s.D.Set("content_disposition", *s.Res.ObjectResponse.ContentDisposition)
	}

	if s.Res.ObjectResponse.ContentEncoding != nil {
		s.D.Set("content_encoding", *s.Res.ObjectResponse.ContentEncoding)
	}

	if s.Res.ObjectResponse.ContentLanguage != nil {
		s.D.Set("content_language", *s.Res.ObjectResponse.ContentLanguage)
	}

	if s.Res.ObjectResponse.ContentLength != nil {
		s.D.Set("content_length", strconv.FormatInt(*s.Res.ObjectResponse.ContentLength, 10))
	}

	if s.Res.ObjectResponse.ContentMd5 != nil {
		s.D.Set("content_md5", *s.Res.ObjectResponse.ContentMd5)
	}

	if s.Res.ObjectResponse.ContentType != nil {
		s.D.Set("content_type", *s.Res.ObjectResponse.ContentType)
	}

	if s.Res.ObjectResponse.VersionId != nil {
		s.D.Set("version_id", *s.Res.ObjectResponse.VersionId)
	}

	s.D.Set("storage_tier", string(s.Res.ObjectResponse.StorageTier))

	if s.Res.ObjectResponse.OpcMeta != nil {
		// Note: regardless of what we sent to the SDK, the keys we get back from OpcMeta will always be
		// converted to lower case
		if err := s.D.Set("metadata", s.Res.ObjectResponse.OpcMeta); err != nil {
			log.Printf("Unable to set 'metadata'. Error: %q", err)
		}
	}

	return nil
}

func ObjectSummaryToMap(obj oci_object_storage.ObjectSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["archival_state"] = string(obj.ArchivalState)

	if obj.Etag != nil {
		result["etag"] = string(*obj.Etag)
	}

	if obj.Md5 != nil {
		result["md5"] = string(*obj.Md5)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Size != nil {
		result["size"] = strconv.FormatInt(*obj.Size, 10)
	}

	result["storage_tier"] = string(obj.StorageTier)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeModified != nil {
		result["time_modified"] = obj.TimeModified.String()
	}

	return result
}
