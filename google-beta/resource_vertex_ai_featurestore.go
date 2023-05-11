// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceVertexAIFeaturestore() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAIFeaturestoreCreate,
		Read:   resourceVertexAIFeaturestoreRead,
		Update: resourceVertexAIFeaturestoreUpdate,
		Delete: resourceVertexAIFeaturestoreDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAIFeaturestoreImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"encryption_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `If set, both of the online and offline data storage will be secured by this key.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource. Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the compute resource is created.`,
						},
					},
				},
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `A set of key/value label pairs to assign to this Featurestore.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.`,
			},
			"online_serving_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Config for online serving resources.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fixed_node_count": {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  `The number of nodes for each cluster. The number of nodes will not scale automatically but can be scaled manually by providing different values when updating.`,
							ExactlyOneOf: []string{"online_serving_config.0.fixed_node_count", "online_serving_config.0.scaling"},
						},
						"scaling": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Online serving scaling configuration. Only one of fixedNodeCount and scaling can be set. Setting one will reset the other.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_node_count": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: `The maximum number of nodes to scale up to. Must be greater than minNodeCount, and less than or equal to 10 times of 'minNodeCount'.`,
									},
									"min_node_count": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: `The minimum number of nodes to scale down to. Must be greater than or equal to 1.`,
									},
								},
							},
							ExactlyOneOf: []string{"online_serving_config.0.fixed_node_count", "online_serving_config.0.scaling"},
						},
					},
				},
			},
			"online_storage_ttl_days": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: `TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a featurestore. If not set, default to 4000 days`,
				Default:     4000,
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the dataset. eg us-central1`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Used to perform consistent read-modify-write updates.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"force_destroy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: `If set to true, any EntityTypes and Features for this Featurestore will also be deleted`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceVertexAIFeaturestoreCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandVertexAIFeaturestoreLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	onlineServingConfigProp, err := expandVertexAIFeaturestoreOnlineServingConfig(d.Get("online_serving_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("online_serving_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(onlineServingConfigProp)) && (ok || !reflect.DeepEqual(v, onlineServingConfigProp)) {
		obj["onlineServingConfig"] = onlineServingConfigProp
	}
	onlineStorageTtlDaysProp, err := expandVertexAIFeaturestoreOnlineStorageTtlDays(d.Get("online_storage_ttl_days"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("online_storage_ttl_days"); !tpgresource.IsEmptyValue(reflect.ValueOf(onlineStorageTtlDaysProp)) && (ok || !reflect.DeepEqual(v, onlineStorageTtlDaysProp)) {
		obj["onlineStorageTtlDays"] = onlineStorageTtlDaysProp
	}
	encryptionSpecProp, err := expandVertexAIFeaturestoreEncryptionSpec(d.Get("encryption_spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("encryption_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptionSpecProp)) && (ok || !reflect.DeepEqual(v, encryptionSpecProp)) {
		obj["encryptionSpec"] = encryptionSpecProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featurestores?featurestoreId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Featurestore: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Featurestore: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Featurestore: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/featurestores/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = VertexAIOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Featurestore", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Featurestore: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/featurestores/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Featurestore %q: %#v", d.Id(), res)

	return resourceVertexAIFeaturestoreRead(d, meta)
}

func resourceVertexAIFeaturestoreRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featurestores/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Featurestore: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("VertexAIFeaturestore %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("force_destroy"); !ok {
		if err := d.Set("force_destroy", false); err != nil {
			return fmt.Errorf("Error setting force_destroy: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Featurestore: %s", err)
	}

	if err := d.Set("create_time", flattenVertexAIFeaturestoreCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Featurestore: %s", err)
	}
	if err := d.Set("update_time", flattenVertexAIFeaturestoreUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Featurestore: %s", err)
	}
	if err := d.Set("labels", flattenVertexAIFeaturestoreLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Featurestore: %s", err)
	}
	if err := d.Set("online_serving_config", flattenVertexAIFeaturestoreOnlineServingConfig(res["onlineServingConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Featurestore: %s", err)
	}
	if err := d.Set("online_storage_ttl_days", flattenVertexAIFeaturestoreOnlineStorageTtlDays(res["onlineStorageTtlDays"], d, config)); err != nil {
		return fmt.Errorf("Error reading Featurestore: %s", err)
	}
	if err := d.Set("encryption_spec", flattenVertexAIFeaturestoreEncryptionSpec(res["encryptionSpec"], d, config)); err != nil {
		return fmt.Errorf("Error reading Featurestore: %s", err)
	}

	return nil
}

func resourceVertexAIFeaturestoreUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Featurestore: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandVertexAIFeaturestoreLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	onlineServingConfigProp, err := expandVertexAIFeaturestoreOnlineServingConfig(d.Get("online_serving_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("online_serving_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, onlineServingConfigProp)) {
		obj["onlineServingConfig"] = onlineServingConfigProp
	}
	onlineStorageTtlDaysProp, err := expandVertexAIFeaturestoreOnlineStorageTtlDays(d.Get("online_storage_ttl_days"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("online_storage_ttl_days"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, onlineStorageTtlDaysProp)) {
		obj["onlineStorageTtlDays"] = onlineStorageTtlDaysProp
	}
	encryptionSpecProp, err := expandVertexAIFeaturestoreEncryptionSpec(d.Get("encryption_spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("encryption_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, encryptionSpecProp)) {
		obj["encryptionSpec"] = encryptionSpecProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featurestores/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Featurestore %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("online_serving_config") {
		updateMask = append(updateMask, "onlineServingConfig")
	}

	if d.HasChange("online_storage_ttl_days") {
		updateMask = append(updateMask, "onlineStorageTtlDays")
	}

	if d.HasChange("encryption_spec") {
		updateMask = append(updateMask, "encryptionSpec")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Featurestore %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Featurestore %q: %#v", d.Id(), res)
	}

	err = VertexAIOperationWaitTime(
		config, res, project, "Updating Featurestore", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceVertexAIFeaturestoreRead(d, meta)
}

func resourceVertexAIFeaturestoreDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Featurestore: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featurestores/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	if v, ok := d.GetOk("force_destroy"); ok {
		url, err = transport_tpg.AddQueryParams(url, map[string]string{"force": fmt.Sprintf("%v", v)})
		if err != nil {
			return err
		}
	}
	log.Printf("[DEBUG] Deleting Featurestore %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Featurestore")
	}

	err = VertexAIOperationWaitTime(
		config, res, project, "Deleting Featurestore", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Featurestore %q: %#v", d.Id(), res)
	return nil
}

func resourceVertexAIFeaturestoreImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/featurestores/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/featurestores/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("force_destroy", false); err != nil {
		return nil, fmt.Errorf("Error setting force_destroy: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenVertexAIFeaturestoreCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreOnlineServingConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["fixed_node_count"] =
		flattenVertexAIFeaturestoreOnlineServingConfigFixedNodeCount(original["fixedNodeCount"], d, config)
	transformed["scaling"] =
		flattenVertexAIFeaturestoreOnlineServingConfigScaling(original["scaling"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeaturestoreOnlineServingConfigFixedNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenVertexAIFeaturestoreOnlineServingConfigScaling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["min_node_count"] =
		flattenVertexAIFeaturestoreOnlineServingConfigScalingMinNodeCount(original["minNodeCount"], d, config)
	transformed["max_node_count"] =
		flattenVertexAIFeaturestoreOnlineServingConfigScalingMaxNodeCount(original["maxNodeCount"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeaturestoreOnlineServingConfigScalingMinNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenVertexAIFeaturestoreOnlineServingConfigScalingMaxNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenVertexAIFeaturestoreOnlineStorageTtlDays(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenVertexAIFeaturestoreEncryptionSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_name"] =
		flattenVertexAIFeaturestoreEncryptionSpecKmsKeyName(original["kmsKeyName"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeaturestoreEncryptionSpecKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandVertexAIFeaturestoreLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandVertexAIFeaturestoreOnlineServingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFixedNodeCount, err := expandVertexAIFeaturestoreOnlineServingConfigFixedNodeCount(original["fixed_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFixedNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fixedNodeCount"] = transformedFixedNodeCount
	}

	transformedScaling, err := expandVertexAIFeaturestoreOnlineServingConfigScaling(original["scaling"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaling); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scaling"] = transformedScaling
	}

	return transformed, nil
}

func expandVertexAIFeaturestoreOnlineServingConfigFixedNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreOnlineServingConfigScaling(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinNodeCount, err := expandVertexAIFeaturestoreOnlineServingConfigScalingMinNodeCount(original["min_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minNodeCount"] = transformedMinNodeCount
	}

	transformedMaxNodeCount, err := expandVertexAIFeaturestoreOnlineServingConfigScalingMaxNodeCount(original["max_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxNodeCount"] = transformedMaxNodeCount
	}

	return transformed, nil
}

func expandVertexAIFeaturestoreOnlineServingConfigScalingMinNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreOnlineServingConfigScalingMaxNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreOnlineStorageTtlDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreEncryptionSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandVertexAIFeaturestoreEncryptionSpecKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandVertexAIFeaturestoreEncryptionSpecKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
