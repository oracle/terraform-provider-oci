// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
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
		Timeouts: DefaultTimeout,
		Create:   createObjectStorageObject,
		Read:     readObjectStorageObject,
		Update:   updateObjectStorageObject,
		Delete:   deleteObjectStorageObject,
		Schema: map[string]*schema.Schema{
			// @CODEGEN 2/2018:
			// Previous provider doesn't provide an Update method and sets all non-Computed fields to ForceNew.
			// This was done because the same PutObject() call can be used to create and modify existing objects.
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
				// @CODEGEN 06/2018: object renames are now supported
			},

			// Optional
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
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"content_length": {
				Type: schema.TypeString,
				// @CODEGEN 2/2018: this was generated as Required, we will compute the length from the 'content'
				Computed: true,
			},
			"content_md5": {
				Type: schema.TypeString,
				// @CODEGEN 2/2018: this was generated as Optional, we will set it from the service's response
				Computed: true,
			},
			"content_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"source": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"content", "source_uri_details"},
				StateFunc:     setSourceState,
				ValidateFunc:  validateSourceValue,
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
	sync.Client = m.(*OracleClients).objectStorageClient

	return CreateResource(d, sync)
}

func setSourceState(source interface{}) string {
	sourcePath := source.(string)
	sourceInfo, err := os.Stat(sourcePath)

	if err != nil {
		return sourcePath
	}

	return sourcePath + " " + sourceInfo.ModTime().String()
}

func (s *ObjectStorageObjectResourceCrud) createMultiPartObject() error {
	multipartUploadData := MultipartUploadData{}

	source, ok := s.D.GetOkExists("source")
	if !ok {
		return fmt.Errorf("the source is not specified to create multipart upload")
	}
	tmpSource := source.(string)
	sourceInfo, err := os.Stat(tmpSource)
	if err != nil {
		return fmt.Errorf("the specified source is not available: %q", err)
	}

	multipartUploadData.SourcePath = &tmpSource
	multipartUploadData.SourceInfo = &sourceInfo

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
	multipartUploadData.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

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
	}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		copyObjectRequest.DestinationBucket = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		copyObjectRequest.DestinationObjectMetadata = resourceObjectStorageMapToOPCMetadata(metadata.(map[string]interface{}))
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		copyObjectRequest.DestinationNamespace = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		copyObjectRequest.DestinationObjectName = &tmp
	}

	copyObjectRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

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
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest

	copyTimeout := *DefaultTimeout.Create
	err = copyObjectWaitForWorkRequest(&workRequestId, "object", copyTimeout, s.DisableNotFoundRetries, s.SourceRegionClient)

	if err != nil {
		// we are not able to verify the state of workRequest
		s.D.Set("state", string(oci_object_storage.WorkRequestStatusFailed))
		return err
	}

	s.D.Set("work_request_id", "")
	s.D.Set("state", string(oci_object_storage.WorkRequestStatusCompleted))
	id := getId(*copyObjectRequest.DestinationNamespace, *copyObjectRequest.DestinationBucket, *copyObjectRequest.DestinationObjectName)
	s.D.SetId(id)
	return s.Get()
}

func readObjectStorageObject(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return ReadResource(sync)
}

func updateObjectStorageObject(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return UpdateResource(d, sync)
}

func deleteObjectStorageObject(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
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
	BaseCrud
	Client                 *oci_object_storage.ObjectStorageClient
	SourceRegionClient     *oci_object_storage.ObjectStorageClient
	Res                    *ObjectStorageObject
	DisableNotFoundRetries bool
	WorkRequest            *oci_object_storage.WorkRequest
}

// @CODEGEN 2/2018: The existing provider returns a custom Id in following format:
// "tfobm-object-<namespace_name>/<bucket_name>/<object_name>"
func getId(namespaceName string, bucketName string, objectName string) string {
	return ObjIdPrefix + namespaceName + ObjIdDelim + bucketName + ObjIdDelim + objectName
}

func parseId(id string) (namespaceName string, bucketName string, objectName string, err error) {
	parts := strings.Split(strings.TrimPrefix(id, ObjIdPrefix), ObjIdDelim)
	if len(parts) < 3 {
		err = fmt.Errorf("Illegal id %s encountered", id)
	}
	namespaceName, bucketName, objectName = parts[0], parts[1], parts[2]

	// Sometimes, the delimiter is used in the object name, and we should use all of the remaining parts, rather than
	// first only
	if len(parts) > 3 {
		objectName = strings.Join(parts[2:], ObjIdDelim)
	}
	return
}

func (s *ObjectStorageObjectResourceCrud) ID() string {
	return getId(s.Res.NamespaceName, s.Res.BucketName, s.Res.ObjectName)
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

	if contentEncoding, ok := s.D.GetOkExists("content_encoding"); ok {
		tmp := contentEncoding.(string)
		request.ContentEncoding = &tmp
	}

	if contentLanguage, ok := s.D.GetOkExists("content_language"); ok {
		tmp := contentLanguage.(string)
		request.ContentLanguage = &tmp
	}

	// @CODEGEN 2/2018: Generator code allow you to set the content_length and
	// content_md5 fields from the schema. These are omitted from the existing provider
	// resource because they can be computed.

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
		request.PutObjectBody = ioutil.NopCloser(bytes.NewBuffer(tmp))
	} else {
		tmp := int64(0)
		request.ContentLength = &tmp
		request.PutObjectBody = ioutil.NopCloser(bytes.NewBuffer([]byte{}))
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	s.D.Set("work_request_id", "")
	s.D.Set("state", oci_object_storage.WorkRequestStatusInProgress)

	_, err := s.Client.PutObject(context.Background(), request)
	if err != nil {
		return err
	}

	id := getId(*request.NamespaceName, *request.BucketName, *request.ObjectName)
	s.D.SetId(id)

	// @CODEGEN 2/2018: PutObject() call doesn't return an object. Instead, use existing
	// Get() implementation to retrieve the state of the object and set its state to completed
	s.D.Set("state", oci_object_storage.WorkRequestStatusCompleted)

	return s.Get()
}

func (s *ObjectStorageObjectResourceCrud) getObjectHead() error {

	headObjectRequest := &oci_object_storage.HeadObjectRequest{}

	namespaceName, bucketName, objectName, err := parseId(s.D.Id())
	if err != nil {
		return err
	}

	headObjectRequest.NamespaceName = &namespaceName
	headObjectRequest.BucketName = &bucketName
	headObjectRequest.ObjectName = &objectName

	if headObjectRequest.NamespaceName == nil || headObjectRequest.BucketName == nil || headObjectRequest.ObjectName == nil {
		return fmt.Errorf("'namespace', 'bucket', or 'object' identifiers are missing")
	}

	headObjectRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

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
				retryPolicy := getRetryPolicy(s.DisableNotFoundRetries, "object_storage")
				copyTimeout := DefaultTimeout.Create
				retryPolicy.ShouldRetryOperation = objectStorageWorkRequestShouldRetryFunc(*copyTimeout)

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

	namespaceName, bucketName, objectName, err := parseId(s.D.Id())
	if err != nil {
		return err
	}

	request.NamespaceName = &namespaceName
	request.BucketName = &bucketName
	request.ObjectName = &objectName

	if request.NamespaceName == nil || request.BucketName == nil || request.ObjectName == nil {
		return fmt.Errorf("'namespace', 'bucket', or 'object' identifiers are missing")
	}

	// TODO: May be better to use HeadObject() to retrieve status of the object. For large content, doesn't make sense
	// to call Get() all the time
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

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
	id := s.D.Id()
	namespaceName, bucketName, objectName, err := parseId(id)
	if err != nil {
		return err
	}

	// @CODEGEN 06/2018: Update is only supported for the change in name - all others are a forceNew
	if !s.D.HasChange("object") {
		return fmt.Errorf("unexpected change encountered")
	}
	request := oci_object_storage.RenameObjectRequest{}
	request.NamespaceName = &namespaceName
	request.BucketName = &bucketName
	request.SourceName = &objectName
	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.NewName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")
	_, err = s.Client.RenameObject(context.Background(), request)
	if err != nil {
		return err
	}

	updatedId := getId(namespaceName, bucketName, *request.NewName)
	s.D.SetId(updatedId)
	return s.Get()
}

func (s *ObjectStorageObjectResourceCrud) Delete() error {
	request := oci_object_storage.DeleteObjectRequest{}

	namespaceName, bucketName, objectName, err := parseId(s.D.Id())
	if err != nil {
		return err
	}

	request.NamespaceName = &namespaceName
	request.BucketName = &bucketName
	request.ObjectName = &objectName

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	_, err = s.Client.DeleteObject(context.Background(), request)
	return err
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

	if response.OpcMeta != nil {
		if err := s.D.Set("metadata", response.OpcMeta); err != nil {
			log.Printf("Unable to set 'metadata'. Error: %q", err)
		}
	}

	return nil
}

func (s *ObjectStorageObjectResourceCrud) setDataObject() error {
	s.D.Set("namespace", s.Res.NamespaceName)
	s.D.Set("bucket", s.Res.BucketName)
	s.D.Set("object", s.Res.ObjectName)

	contentReader := s.Res.ObjectResponse.Content
	contentArray, err := ioutil.ReadAll(contentReader)
	if err != nil {
		log.Printf("Unable to read 'content' from response. Error: %q", err)
		return err
	}
	s.D.Set("content", contentArray)

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

	if s.Res.ObjectResponse.OpcMeta != nil {
		// Note: regardless of what we sent to the SDK, the keys we get back from OpcMeta will always be
		// converted to lower case
		if err := s.D.Set("metadata", s.Res.ObjectResponse.OpcMeta); err != nil {
			log.Printf("Unable to set 'metadata'. Error: %q", err)
		}
	}

	return nil
}

// @CODEGEN 2/2018: Remove generated mapToObjectSummary as it's not being called

func ObjectSummaryToMap(obj oci_object_storage.ObjectSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Md5 != nil {
		result["md5"] = string(*obj.Md5)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Size != nil {
		result["size"] = strconv.FormatInt(*obj.Size, 10)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
