package identity_domains

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetIdentityDomainsCompositeId(idcsEndpoint string, resourceName string, resId string) string {
	// e.g. idcsEndpoint/https://something.com/groups/{groupId}
	return fmt.Sprintf("idcsEndpoint/%s/%s/%s", idcsEndpoint, resourceName, resId)
}

func getIdcsEndpoint(d *schema.ResourceData) (string, error) {
	endpoint, ok := d.GetOkExists("idcs_endpoint")
	if !ok {
		return "", fmt.Errorf("[ERROR] IdcsEndpointHelper: idcs_endpoint missing for resource. OCID:%s ID:%s", d.Get("ocid"), d.Id())
	}
	return endpoint.(string), nil
}

func getIdcsEndpointForRead(d *schema.ResourceData, resourceName string) (string, error) {
	if endpoint, err := getIdcsEndpoint(d); err == nil {
		return endpoint, nil
	}

	// if failed, check if it's Import use case
	var err error
	var endpoint string
	id := d.Id()
	// compositeId format: idcsEndpoint/{idcsEndpoint}/{resource}/{groupId}
	regex, _ := regexp.Compile("^idcsEndpoint/(.*)/" + resourceName + "/(.*)$")
	tokens := regex.FindStringSubmatch(id)

	if len(tokens) == 3 {
		endpoint = tokens[1]
		// set resource idcs_endpoint and id
		d.Set("idcs_endpoint", tokens[1])
		d.SetId(tokens[2])
		err = nil
	} else {
		err = fmt.Errorf("IdcsEndpointHelperForRead: idcs_endpoint missing. Format of id might be wrong. id: %s", id)
	}

	if err != nil {
		return "", err
	}

	return endpoint, nil
}

const identityDomainsCompositeIdRegexPattern = "^idcsEndpoint/(.*)/users/(.*)$"

func tryMatchUserSubResFilter(value string) string {
	var userId string
	regex, _ := regexp.Compile(identityDomainsCompositeIdRegexPattern)
	tokens := regex.FindStringSubmatch(value)
	if len(tokens) == 3 {
		userId = tokens[2]
		return fmt.Sprintf("user.value eq \"%s\"", userId)
	}

	return value
}

func IsEmptyValued(rawValue interface{}) bool {
	return rawValue == "" || rawValue == 0
}

func IsOptionalField(resource *schema.Resource, fieldPath []string) bool {
	pathLen := len(fieldPath)
	targetField := fieldPath[pathLen-1]
	currentSchemaMap := resource.Schema

	for i := 0; i < pathLen-1; i++ {
		part := fieldPath[i]
		if part == "0" {
			continue
		}
		currentSchemaMap = currentSchemaMap[part].Elem.(*schema.Resource).Schema
	}

	return currentSchemaMap[targetField].Optional
}

func IsOptionalAndEmpty(resource *schema.Resource, D *schema.ResourceData, fieldKey string) bool {
	return IsOptionalField(resource, strings.Split(fieldKey, ".")) && IsEmptyValued(D.Get(fieldKey))
}
