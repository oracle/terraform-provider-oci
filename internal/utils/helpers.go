// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	"strconv"

	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

const (
	Charset                       = CharsetWithoutDigits + "0123456789"
	CharsetWithoutDigits          = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetLowerCaseWithoutDigits = "abcdefghijklmnopqrstuvwxyz"

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
	DefinedShieldedImageOCIDs = `
      variable "InstanceImageOCIDShieldedCompatible" {
	  type = "map"
	  default = {
		// See https://docs.us-phoenix-1.oraclecloud.com/images/
		// Oracle-provided image "Oracle-Linux-8.4-2021.07.27-0"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaalw6v7wxgiuedh36jzy5ilbnfjezsxey2glgg3jtlodwzltxregba"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaay7qb3q2bttzhvzzacdpweqo2mvj43tfkm5b4j46xf6pzazspz6aq"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaasehd3xu75nmbxp6lbaaontfgmowlszrx5c72mw4kks5f75euj7gq"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaa7fgs4dpcjkkeemyfzyo3yo5lezqfskac45dblmgnfq5az4jmgcza"
	  }
	}`

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
		us-phoenix-1 = "ocid1.mysqlconfiguration.oc1..aaaaaaaah6o6qu3gdbxnqg6aw56amnosmnaycusttaa7abyq2tdgpgubvsgj"
	  }
	}
    `
	MysqlHAConfigurationIdVariable = `
	variable "MysqlHAConfigurationOCID" {
		type = "map"
		default = {
			us-ashburn-1 = "ocid1.mysqlconfiguration.oc1..aaaaaaaalwzc2a22xqm56fwjwfymixnulmbq3v77p5v4lcbb6qhkftxf2trq"
			us-phoenix-1 = "ocid1.mysqlconfiguration.oc1..aaaaaaaantprksu6phqfgr5xvyut46wdfesdszonbclybfwvahgysfjbrb4q"
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
	OsManagedImageIdsVariable = `
	variable "OsManagedImageOCID" {
	  type = "map"
	  default = {
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaavxm3s4jskx5rcoi63rekg54e3a27v2b7tiuuumnx5owzhkul6ufq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaqhzgbezuoq5fz7haq5p7uyydfipffclz6w7fwyzge7tcxbbloz3q"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaulz7xiht632iidvdm4iezy33fofulmerq2nkllwnkjy335qkswza"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaayt6ppuyj6q4dwb4pkkyy3llrhxntywewfk4ssd365d4cn22i6yxa"
	  }
	}
	`
	TfEnvPrefix  = "TF_VAR_"
	OciEnvPrefix = "OCI_"
)

func LiteralTypeHashCodeForSets(m interface{}) int {
	return hashcode.String(fmt.Sprintf("%v", m))
}

func ValidateBoolInSlice(valid []bool) schema.SchemaValidateFunc {
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

func ValidateNotEmptyString() schema.SchemaValidateFunc {
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

func ObjectMapToStringMap(rm map[string]interface{}) map[string]string {
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

// Borrowed from https://mijailovic.net/2017/05/09/error-handling-patterns-in-go/
func SafeClose(c io.Closer, err *error) {
	if cerr := c.Close(); cerr != nil && *err == nil {
		*err = cerr
	}
}

func StringMapToObjectMap(sm map[string]string) map[string]interface{} {
	var result = make(map[string]interface{})
	if len(sm) > 0 {
		for types, v := range sm {
			result[types] = v
		}
	}
	return result
}

// Returns date-time formatted as a string, ex: 2017-10-12-000934-119299083"
func Timestamp() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d-%02d%02d%02d-%d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
}

func ValidateInt64TypeString(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	_, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q (%q) must be a 64-bit integer", k, v))
	}
	return
}

func Int64StringDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
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
func MonetaryDiffSuppress(key string, old string, new string, d *schema.ResourceData) bool {
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

func TimeDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
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

func ConvertMapOfStringSlicesToMapOfStrings(rm map[string][]string) (map[string]string, error) {
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

func RandomString(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomStringOrHttpReplayValue(length int, charset string, httpReplayValue string) string {
	if httpreplay.ModeRecordReplay() {
		return httpReplayValue
	}
	return RandomString(length, charset)
}

// Set the state for the input source file using the file path and last modification time
// this information helps us to identify if the file has changed.
func GetSourceFileState(source interface{}) string {
	sourcePath := source.(string)
	sourceInfo, err := os.Stat(sourcePath)

	if err != nil {
		return sourcePath
	}

	return sourcePath + " " + sourceInfo.ModTime().String()
}

// Returns a slice of keys from the given map in alphabetical order
func GetSortedKeys(source map[string]interface{}) []string {
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
func JsonStringDiffSuppressFunction(key, old, new string, d *schema.ResourceData) bool {
	var oldVal, newVal interface{}

	if err := json.Unmarshal([]byte(old), &oldVal); err != nil {
		return false
	}

	if err := json.Unmarshal([]byte(new), &newVal); err != nil {
		return false
	}

	return reflect.DeepEqual(oldVal, newVal)
}

func GetMd5Hash(source interface{}) string {
	if source == nil {
		return ""
	}
	data := source.(string)
	hexSum := md5.Sum([]byte(data))
	return hex.EncodeToString(hexSum[:])
}

func HexToB64(hexEncoded string) (*string, error) {
	decoded, err := hex.DecodeString(hexEncoded)
	if err != nil {
		return nil, err
	}

	b64Encoded := base64.StdEncoding.EncodeToString(decoded)
	return &b64Encoded, nil
}

func IsHex(content string) bool {
	_, err := hex.DecodeString(content)
	return err == nil
}

// Get obo token from file
func GetTokenFromFile(path string) (string, error) {
	token, err := ioutil.ReadFile(path)
	return string(token), err
}

// multierror with \t does not show up on Team City logs,
// replacing \t with 4 blank spaces
func CustomErrorFormat(es []error) string {
	if len(es) == 1 {
		return fmt.Sprintf("1 error occurred:\n    * %s\n\n", es[0])
	}

	points := make([]string, len(es))
	for i, err := range es {
		points[i] = fmt.Sprintf("* %s", err)
	}

	return fmt.Sprintf("%d errors occurred:\n    %s\n\n", len(es), strings.Join(points, "\n    "))
}

// Added for resource discovery AUTH
func GetProviderEnvSettingWithDefault(s string, dv string) string {
	v := os.Getenv(TfEnvPrefix + s)
	if v != "" {
		return v
	}
	v = os.Getenv(OciEnvPrefix + strings.ToUpper(s))
	if v != "" {
		return v
	}
	v = os.Getenv(s)
	if v != "" {
		return v
	}
	return dv
}

func GetEnvSettingWithBlankDefault(s string) string {
	return GetEnvSettingWithDefault(s, "")
}

func GetEnvSettingWithDefault(s string, dv string) string {
	v := os.Getenv(TfEnvPrefix + s)
	if v != "" {
		return v
	}
	v = os.Getenv(OciEnvPrefix + s)
	if v != "" {
		return v
	}
	v = os.Getenv(s)
	if v != "" {
		return v
	}
	return dv
}

// Deprecated: There should be only no need to panic individually
func GetRequiredEnvSetting(s string) string {
	v := GetEnvSettingWithBlankDefault(s)
	if v == "" {
		panic(fmt.Sprintf("Required env setting %s is missing", s))
	}
	return v
}

func GetHomeFolder() string {
	if os.Getenv("TF_HOME_OVERRIDE") != "" {
		return os.Getenv("TF_HOME_OVERRIDE")
	}
	current, e := user.Current()
	if e != nil {
		//Give up and try to return something sensible
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return current.HomeDir
}

// cleans and expands the path if it contains a tilde , returns the expanded path or the input path as is if not expansion
// was performed
func ExpandPath(filepath string) string {
	if strings.HasPrefix(filepath, fmt.Sprintf("~%c", os.PathSeparator)) {
		filepath = path.Join(GetHomeFolder(), filepath[2:])
	}
	return path.Clean(filepath)
}

func CheckProfile(profile string, path string) (err error) {
	var profileRegex = regexp.MustCompile(`^\[(.*)\]`)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(data)
	splitContent := strings.Split(content, "\n")
	for _, line := range splitContent {
		if match := profileRegex.FindStringSubmatch(line); match != nil && len(match) > 1 && match[1] == profile {
			return nil
		}
	}

	return fmt.Errorf("configuration file did not contain profile: %s", profile)
}

func CheckIncompatibleAttrsForApiKeyAuth(d *schema.ResourceData, apiKeyConfigAttributes [5]string) ([]string, bool) {
	var apiKeyConfigAttributesToUnset []string
	for _, apiKeyConfigAttribute := range apiKeyConfigAttributes {
		apiKeyConfigAttributeValue, hasConfigVariable := d.GetOkExists(apiKeyConfigAttribute)
		if (hasConfigVariable && apiKeyConfigAttributeValue != "") || GetEnvSettingWithBlankDefault(apiKeyConfigAttribute) != "" {
			apiKeyConfigAttributesToUnset = append(apiKeyConfigAttributesToUnset, apiKeyConfigAttribute)
		}
	}
	return apiKeyConfigAttributesToUnset, len(apiKeyConfigAttributesToUnset) == 0
}

func GetCertificateFileBytes(certificateFileFullPath string) (pemRaw []byte, err error) {
	absFile, err := filepath.Abs(certificateFileFullPath)
	if err != nil {
		return nil, fmt.Errorf("can't form absolute path of %s: %v", certificateFileFullPath, err)
	}

	if pemRaw, err = ioutil.ReadFile(absFile); err != nil {
		return nil, fmt.Errorf("can't read %s: %v", certificateFileFullPath, err)
	}
	return
}

func RemoveFile(file string) error {
	return os.Remove(file)
}

func WriteTempFile(data string, originFileName string) (err error) {
	f, err := os.OpenFile(originFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		_, _ = f.WriteString(data)
	}
	return err
}

func ValidateSourceValue(i interface{}, k string) (s []string, es []error) {
	v, ok := i.(string)
	if !ok {
		es = append(es, fmt.Errorf("expected type of %s to be string", k))
		return
	}
	info, err := os.Stat(v)
	if err != nil {
		es = append(es, fmt.Errorf("cannot get file information for the specified source: %s", v))
		return
	}
	if info.Size() > 10000*50*1024*1024*1024 {
		es = append(es, fmt.Errorf("the specified source: %s file is too large", v))
	}
	return
}
