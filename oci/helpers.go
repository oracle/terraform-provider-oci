// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"sort"
	"time"

	"strconv"

	"encoding/json"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

const (
	charset              = charsetWithoutDigits + "0123456789"
	charsetWithoutDigits = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	OciImageIdsVariable = `
variable "InstanceImageOCID" {
	type = "map"
	default = {
		// See https://docs.us-phoenix-1.oraclecloud.com/images/
		// Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaadjnj3da72bztpxinmqpih62c2woscbp6l3wjn36by2cvmdhjub6a"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaawufnve5jxze4xf7orejupw5iq3pms6cuadzjc7klojix6vmk42va"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaagbrvhganmn7awcr7plaaf5vhabmzhx763z5afiitswjwmzh7upna"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaajwtut4l7fo3cvyraate6erdkyf2wdk5vpk6fp6ycng3dv2y3ymvq"
	}
}
	// Gets a list of all Oracle Linux 7.5 images that support a given Instance shape
	data "oci_core_images" "supported_shape_images" {
		compartment_id   = "${var.tenancy_ocid}"
		shape            = "VM.Standard2.1"
		operating_system = "Oracle Linux"
	}

`
	FlexVmImageIdsVariable = `
	variable "FlexInstanceImageOCID" {
	  type = "map"
	  default = {
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaa6hooptnlbfwr5lwemqjbu3uqidntrlhnt45yihfj222zahe7p3wq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaa6tp7lhyrcokdtf7vrbmxyp2pctgg4uxvt4jz4vc47qoc2ec4anha"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaadvi77prh3vjijhwe5xbd6kjg3n5ndxjcpod6om6qaiqeu3csof7a"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaw5gvriwzjhzt2tnylrfnpanz5ndztyrv3zpwhlzxdbkqsjfkwxaq"
	  }
	}
	`
	MysqlConfigurationIdVariable = `
	variable "MysqlConfigurationOCID" {
	  type = "map"
	  default = {
		us-ashburn-1 = "ocid1.mysqlconfiguration.oc1..aaaaaaaah6o6qu3gdbxnqg6aw56amnosmnaycusttaa7abyq2tdgpgubvsgj"
	  }
	}
	`
	OciWindowsImageIdsVariable = `
	variable "InstanceImageOCID" {
		type = "map"
		default = {
			# The below Image OCIDs are for Windows-Server-2012-R2-Standard-Edition-VM-Gen2-2018.12.12-0
			# See https://docs.cloud.oracle.com/iaas/images/image/5e34cde5-6cef-4cc3-b8f1-c8fc3a088302/
			us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaarlo3ace3wq34aompwj3u2z2xteonboapg663woz6d2iovarowhja"
			us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaabzwak2haqxh3r7h6dajgu4enp7q7hcrreql45awryd5frjsd5l6a"
			eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaourcjktoe3gprvwfksxc36r4rxgbcjs5qvtrja6w6euivci635vq"
			uk-london-1  = "ocid1.image.oc1.uk-london-1.aaaaaaaadb4mg7ii73wkrntmiunr7x7qrh7ompczvy3xbggm27pkhotpgj2q"
		}
	}

`
	VolumeBackupPolicyDependency = `
data "oci_core_volume_backup_policies" "test_volume_backup_policies" {
	filter {
		name = "display_name"
		values = [ "silver" ]
	}
}
`
)

func literalTypeHashCodeForSets(m interface{}) int {
	return hashcode.String(fmt.Sprintf("%v", m))
}

func validateBoolInSlice(valid []bool) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(bool)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be bool", k))
			return
		}

		for _, str := range valid {
			if v == str {
				return
			}
		}

		es = append(es, fmt.Errorf("expected %s to be one of %v, got %t", k, valid, v))
		return
	}
}

func validateNotEmptyString() schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(string)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be string", k))
			return
		}
		if len(v) == 0 {
			es = append(es, fmt.Errorf("%s cannot be an empty string", k))
		}
		return
	}
}

func objectMapToStringMap(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		switch assertedValue := v.(type) {
		case string:
			result[k] = assertedValue
		default:
			// Make a best effort to coerce into a string, even if underlying type is not a string
			log.Printf("[DEBUG] non-string value encountered for key '%s' while converting object map to string map", k)
			result[k] = fmt.Sprintf("%v", assertedValue)
		}
	}
	return result
}

func validateInt64TypeString(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	_, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q (%q) must be a 64-bit integer", k, v))
	}
	return
}

func int64StringDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	// We may get interpolation syntax in this function call as well; so be sure to check for errors.
	oldIntVal, err := strconv.ParseInt(old, 10, 64)
	if err != nil {
		return false
	}

	newIntVal, err := strconv.ParseInt(new, 10, 64)
	if err != nil {
		return false
	}
	return oldIntVal == newIntVal
}

// Ignore differences in floating point numbers after the second decimal place, ex: 1.001 == 1.002
func monetaryDiffSuppress(key string, old string, new string, d *schema.ResourceData) bool {
	oldVal, err := strconv.ParseFloat(old, 10)
	if err != nil {
		return false
	}

	newVal, err := strconv.ParseFloat(new, 10)
	if err != nil {
		return false
	}
	return fmt.Sprintf("%.2f", oldVal) == fmt.Sprintf("%.2f", newVal)
}

func timeDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	oldTime, err := time.Parse(time.RFC3339Nano, old)
	if err != nil {
		return false
	}
	newTime, err := time.Parse(time.RFC3339Nano, new)
	if err != nil {
		return false
	}
	return oldTime.Equal(newTime)
}

func convertMapOfStringSlicesToMapOfStrings(rm map[string][]string) (map[string]string, error) {
	result := map[string]string{}
	for k, v := range rm {
		val, err := json.Marshal(v)
		if err == nil {
			result[k] = string(val)
		} else {
			return nil, err
		}
	}
	return result, nil
}

func randomString(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func randomStringOrHttpReplayValue(length int, charset string, httpReplayValue string) string {
	if httpreplay.ModeRecordReplay() {
		return httpReplayValue
	}
	return randomString(length, charset)
}

// Set the state for the input source file using the file path and last modification time
// this information helps us to identify if the file has changed.
func getSourceFileState(source interface{}) string {
	sourcePath := source.(string)
	sourceInfo, err := os.Stat(sourcePath)

	if err != nil {
		return sourcePath
	}

	return sourcePath + " " + sourceInfo.ModTime().String()
}

// Returns a slice of keys from the given map in alphabetical order
func getSortedKeys(source map[string]interface{}) []string {
	sortedKeys := make([]string, len(source))
	cnt := 0
	for k := range source {
		sortedKeys[cnt] = k
		cnt++
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

// Diff suppression function to make sure that any change in ordering of attributes in JSON objects don't result in diffs.
// For example, the config may have created this:
//  extended_metadata = {
//    nested_object       = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
//  }
//
// But we use json.Marshal to convert the service value to string before storing in state.
// The marshalling doesn't guarantee the same ordering as our config, and so the value in state may look like:
//
//  extended_metadata = {
//    nested_object       = "{\"object\": {\"some_string\": \"stringC\"}, \"some_string\": \"stringB\"}"
//  }
//
// These are the same JSON objects and should be treated as such.
func jsonStringDiffSuppressFunction(key, old, new string, d *schema.ResourceData) bool {
	var oldVal, newVal interface{}

	if err := json.Unmarshal([]byte(old), &oldVal); err != nil {
		return false
	}

	if err := json.Unmarshal([]byte(new), &newVal); err != nil {
		return false
	}

	return reflect.DeepEqual(oldVal, newVal)
}

func getMd5Hash(source interface{}) string {
	if source == nil {
		return ""
	}
	data := source.(string)
	hexSum := md5.Sum([]byte(data))
	return hex.EncodeToString(hexSum[:])
}
