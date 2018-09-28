// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"strings"
	"os"
	"text/template"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	oci_common "github.com/oracle/oci-go-sdk/common"
)

var tmpl template.Template = *template.New("tmpl")

// Applies values from a map to a string template
func apply(template string, values map[string]string) string {
	b := new(bytes.Buffer)
	t, _ := tmpl.Parse(template)
	if err := t.Execute(b, values); err != nil {
		log.Printf("[WARN] Unable to apply values to template: '%s'", err)
	}
	return b.String()
}

// Returns date-time formatted as a string, ex: 2017-10-12-000934-119299083"
func timestamp() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d-%02d%02d%02d-%d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
}

type TokenFn func(string, map[string]string) string

// Creates a form of "apply" above that will always supply the same value for {{.token}}
func tokenize() (string, TokenFn) {
	ts := timestamp()
	return ts, func(template string, values map[string]string) string {
		if values == nil {
			values = map[string]string{}
		}
		values["token"] = ts
		return apply(template, values)
	}
}

// custom TestCheckFunc helper, returns a value associated with a key from an instance in the current state
func fromInstanceState(s *terraform.State, name, key string) (string, error) {
	ms := s.RootModule()
	rs, ok := ms.Resources[name]
	if !ok {
		return "", fmt.Errorf("Not found: %s", name)
	}

	is := rs.Primary
	if is == nil {
		return "", fmt.Errorf("No primary instance: %s", name)
	}

	v, ok := is.Attributes[key]

	if ok {
		return v, nil
	} else {
		return "", fmt.Errorf("%s: Attribute '%s' not found", name, key)
	}
}

// TestCheckResourceAttributesEqual is a TestCheckFunc which ensures that the values of two
// attributes in two different resources are equal.
func TestCheckResourceAttributesEqual(name1, key1, name2, key2 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		val1, err := fromInstanceState(s, name1, key1)
		if err != nil {
			return err
		}

		val2, err := fromInstanceState(s, name2, key2)
		if err != nil {
			return err
		}

		if val1 != val2 {
			return fmt.Errorf(
				"%s: attribute '%s' value %#v not equal to '%s' attribute '%s' value %#v",
				name1, key1, val1, name2, key2, val2)
		}
		return nil
	}
}

// Function to be implemented by resources that wish to wait on a certain condition and this function is responsible for evaluating the specific condition for that resource
type WaitConditionFunc func(response oci_common.OCIOperationResponse) bool

// Function to be implemented by resources that wish to wait on a certain condition and this function is responsible for fetching the latest state using the resourceId
type FetchOperationFunc func(client *OracleClients, resourceId *string, retryPolicy *oci_common.RetryPolicy) error

// This function waits for the given time and retries the WaitConditionFunc and periodically invokes the FetchOperationFunc to fetch the latest response
func waitTillCondition(testAccProvider *schema.Provider, resourceId *string, condition WaitConditionFunc, timeout time.Duration,
	fetchOperationFunc FetchOperationFunc, service string, disableNotFoundRetries bool) func() {
	return func() {
		client := testAccProvider.Meta().(*OracleClients)
		log.Printf("[INFO] start of waitTillCondition for resource %s ", *resourceId)
		retryPolicy := getRetryPolicy(disableNotFoundRetries, service)
		retryPolicy.ShouldRetryOperation = conditionShouldRetry(timeout, condition, service, disableNotFoundRetries)

		err := fetchOperationFunc(client, resourceId, retryPolicy)
		if err != nil {
			log.Printf("[WARN] waitTillCondition failed for %s resource with error %v", *resourceId, err)
		} else {
			log.Printf("[INFO] end of waitTillCondition for resource %s ", *resourceId)
		}
	}
}

// This function is responsible for the actual check for WaitConditionFunc and the aborting
func conditionShouldRetry(timeout time.Duration, condition WaitConditionFunc, service string, disableNotFoundRetries bool) func(response oci_common.OCIOperationResponse) bool {
	stopTime := time.Now().Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		//Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		//Make sure we stop on default rules
		if shouldRetry(response, disableNotFoundRetries, service) {
			return true
		}

		return condition(response)
	}
}

type RepresentationType int

const (
	Required RepresentationType = iota + 1
	Optional
)

type RepresentationMode int

const (
	Create RepresentationMode = iota
	Update
)

type Representation struct {
	repType RepresentationType
	create  interface{}
	update  interface{}
}

type RepresentationGroup struct {
	repType RepresentationType
	group   map[string]interface{}
}

func cloneRepresentation(representations map[string]interface{}) map[string]interface{} {
	copyMap := map[string]interface{}{}

	for key, value := range representations {
		representation, ok := value.(Representation)
		if ok {
			copyMap[key] = Representation{representation.repType, representation.create, representation.update}
		}
		representationGroup, ok := value.(RepresentationGroup)
		if ok {
			copyMap[key] = RepresentationGroup{representationGroup.repType, cloneRepresentation(representationGroup.group)}
		}
	}

	return copyMap
}

func representationCopyWithNewProperties(representations map[string]interface{}, newProperties map[string]interface{}) map[string]interface{} {
	representationsCopy := cloneRepresentation(representations)
	for propName, value := range newProperties {
		representationsCopy[propName] = value
	}
	return representationsCopy
}

func getUpdatedRepresentationCopy(propertyNameStr string, newValue interface{}, representations map[string]interface{}) map[string]interface{} {
	propertyNames := strings.Split(propertyNameStr, ".")
	return updateNestedRepresentation(0, propertyNames, newValue, cloneRepresentation(representations))
}

func updateNestedRepresentation(currIndex int, propertyNames []string, newValue interface{}, representations map[string]interface{}) map[string]interface{} {
	//recursively search the property to update
	for prop := range representations {
		if prop == propertyNames[currIndex] {
			representationGroup, ok := representations[prop].(RepresentationGroup)
			if ok && currIndex+1 < len(propertyNames) {
				updateNestedRepresentation(currIndex+1, propertyNames, newValue, representationGroup.group)
			} else {
				representations[prop] = newValue
			}
			return representations
		}
	}

	return nil
}

func generateDataSourceFromRepresentationMap(resourceType string, resourceName string, representationType RepresentationType, representationMode RepresentationMode, representations map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(`data "%s" "%s" %s`, resourceType, resourceName, generateResourceFromMap(representationType, representationMode, representations)))
	return buffer.String()
}

func generateResourceFromRepresentationMap(resourceType string, resourceName string, representationType RepresentationType, representationMode RepresentationMode, representations map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(`resource "%s" "%s" %s`, resourceType, resourceName, generateResourceFromMap(representationType, representationMode, representations)))
	return buffer.String()
}

func generateResourceFromMap(representationType RepresentationType, representationMode RepresentationMode, representations map[string]interface{}) string {
	var buffer bytes.Buffer
	var lineSeparator = "\n"
	buffer.WriteString("{" + lineSeparator)

	sortedRepresentations := make([]string, 0, len(representations))
	for key := range representations {
		sortedRepresentations = append(sortedRepresentations, key)
	}
	sort.Strings(sortedRepresentations)

	for _, prop := range sortedRepresentations {
		representation, ok := representations[prop].(Representation)
		if ok && representation.repType <= representationType {

			representationValue := representation.create
			if representationMode == Update && representation.update != nil {
				representationValue = representation.update
			}

			repStrValue, strRep := representationValue.(string)
			if strRep {
				buffer.WriteString(fmt.Sprintf(`%s = "%s"%s`, prop, repStrValue, lineSeparator))
			}

			repArrayStrValue, arrayRep := representationValue.([]string)
			if arrayRep {
				var repArrayStrEscValue []string
				for _, arrayValue := range repArrayStrValue {
					repArrayStrEscValue = append(repArrayStrEscValue, fmt.Sprintf(`"%s"`, arrayValue))
				}
				buffer.WriteString(fmt.Sprintf(`%s = [%s]%s`, prop, strings.Join(repArrayStrEscValue, ", "), lineSeparator))
			}

			repMapStrValue, mapRep := representationValue.(map[string]string)
			if mapRep {
				sortedKeys := make([]string, 0, len(repMapStrValue))
				for key := range repMapStrValue {
					sortedKeys = append(sortedKeys, key)
				}
				sort.Strings(sortedKeys)

				var repMapStrEscValue []string
				for _, key := range sortedKeys {
					repMapStrEscValue = append(repMapStrEscValue, fmt.Sprintf(`"%s" = "%s"`, key, repMapStrValue[key]))
				}
				buffer.WriteString(fmt.Sprintf("%s = {\n%s\n}%s", prop, strings.Join(repMapStrEscValue, lineSeparator), lineSeparator))

			}

		}
		representationGroup, ok := representations[prop].(RepresentationGroup)
		if ok && representationGroup.repType <= representationType {
			buffer.WriteString(fmt.Sprintf("%s %s", prop, generateResourceFromMap(representationType, representationMode, representationGroup.group)))
		}
		representationGroupArray, ok := representations[prop].([]RepresentationGroup)
		if ok {
			for _, representationGroupInArray := range representationGroupArray {
				if representationGroupInArray.repType <= representationType {
					buffer.WriteString(fmt.Sprintf("%s %s", prop, generateResourceFromMap(representationType, representationMode, representationGroupInArray.group)))
				}
			}
		}
	}
	buffer.WriteString(fmt.Sprintf("}%s", lineSeparator))
	return buffer.String()
}

// GenerateTestResourceName generates a name for the resource based on the prefix and current time stamp.
func GenerateTestResourceName(prefix string, maxLength int) string {
	return (prefix + time.Now().UTC().Format("20060102150405"))[:maxLength]
}

func setEnvSetting(s, v string) error {
	error := os.Setenv(s, v)
	if error != nil {
		return fmt.Errorf("Failed to set env setting '%s', encountered error: %v", s, error)
	}
	return nil
}
