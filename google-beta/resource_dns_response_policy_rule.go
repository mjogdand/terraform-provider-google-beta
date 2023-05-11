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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceDNSResponsePolicyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSResponsePolicyRuleCreate,
		Read:   resourceDNSResponsePolicyRuleRead,
		Update: resourceDNSResponsePolicyRuleUpdate,
		Delete: resourceDNSResponsePolicyRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDNSResponsePolicyRuleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dns_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.`,
			},
			"response_policy": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `Identifies the response policy addressed by this request.`,
			},
			"rule_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `An identifier for this rule. Must be unique with the ResponsePolicy.`,
			},
			"behavior": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   `Answer this query with a behavior rather than DNS data. Acceptable values are 'behaviorUnspecified', and 'bypassResponsePolicy'`,
				ConflictsWith: []string{"local_data"},
			},
			"local_data": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name;
in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_datas": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `All resource record sets for this selector, one per resource record type. The name must match the dns_name.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `For example, www.example.com.`,
									},
									"type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: verify.ValidateEnum([]string{"A", "AAAA", "CAA", "CNAME", "DNSKEY", "DS", "HTTPS", "IPSECVPNKEY", "MX", "NAPTR", "NS", "PTR", "SOA", "SPF", "SRV", "SSHFP", "SVCB", "TLSA", "TXT"}),
										Description:  `One of valid DNS resource types. Possible values: ["A", "AAAA", "CAA", "CNAME", "DNSKEY", "DS", "HTTPS", "IPSECVPNKEY", "MX", "NAPTR", "NS", "PTR", "SOA", "SPF", "SRV", "SSHFP", "SVCB", "TLSA", "TXT"]`,
									},
									"rrdatas": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1)`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ttl": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: `Number of seconds that this ResourceRecordSet can be cached by
resolvers.`,
									},
								},
							},
						},
					},
				},
				ConflictsWith: []string{"behavior"},
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

func resourceDNSResponsePolicyRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	ruleNameProp, err := expandDNSResponsePolicyRuleRuleName(d.Get("rule_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rule_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(ruleNameProp)) && (ok || !reflect.DeepEqual(v, ruleNameProp)) {
		obj["ruleName"] = ruleNameProp
	}
	dnsNameProp, err := expandDNSResponsePolicyRuleDnsName(d.Get("dns_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dns_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(dnsNameProp)) && (ok || !reflect.DeepEqual(v, dnsNameProp)) {
		obj["dnsName"] = dnsNameProp
	}
	localDataProp, err := expandDNSResponsePolicyRuleLocalData(d.Get("local_data"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("local_data"); !tpgresource.IsEmptyValue(reflect.ValueOf(localDataProp)) && (ok || !reflect.DeepEqual(v, localDataProp)) {
		obj["localData"] = localDataProp
	}
	behaviorProp, err := expandDNSResponsePolicyRuleBehavior(d.Get("behavior"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("behavior"); !tpgresource.IsEmptyValue(reflect.ValueOf(behaviorProp)) && (ok || !reflect.DeepEqual(v, behaviorProp)) {
		obj["behavior"] = behaviorProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/responsePolicies/{{response_policy}}/rules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ResponsePolicyRule: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ResponsePolicyRule: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ResponsePolicyRule: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/responsePolicies/{{response_policy}}/rules/{{rule_name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ResponsePolicyRule %q: %#v", d.Id(), res)

	return resourceDNSResponsePolicyRuleRead(d, meta)
}

func resourceDNSResponsePolicyRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/responsePolicies/{{response_policy}}/rules/{{rule_name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ResponsePolicyRule: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DNSResponsePolicyRule %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ResponsePolicyRule: %s", err)
	}

	if err := d.Set("rule_name", flattenDNSResponsePolicyRuleRuleName(res["ruleName"], d, config)); err != nil {
		return fmt.Errorf("Error reading ResponsePolicyRule: %s", err)
	}
	if err := d.Set("dns_name", flattenDNSResponsePolicyRuleDnsName(res["dnsName"], d, config)); err != nil {
		return fmt.Errorf("Error reading ResponsePolicyRule: %s", err)
	}
	if err := d.Set("local_data", flattenDNSResponsePolicyRuleLocalData(res["localData"], d, config)); err != nil {
		return fmt.Errorf("Error reading ResponsePolicyRule: %s", err)
	}
	if err := d.Set("behavior", flattenDNSResponsePolicyRuleBehavior(res["behavior"], d, config)); err != nil {
		return fmt.Errorf("Error reading ResponsePolicyRule: %s", err)
	}

	return nil
}

func resourceDNSResponsePolicyRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ResponsePolicyRule: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	dnsNameProp, err := expandDNSResponsePolicyRuleDnsName(d.Get("dns_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dns_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, dnsNameProp)) {
		obj["dnsName"] = dnsNameProp
	}
	localDataProp, err := expandDNSResponsePolicyRuleLocalData(d.Get("local_data"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("local_data"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, localDataProp)) {
		obj["localData"] = localDataProp
	}
	behaviorProp, err := expandDNSResponsePolicyRuleBehavior(d.Get("behavior"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("behavior"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, behaviorProp)) {
		obj["behavior"] = behaviorProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/responsePolicies/{{response_policy}}/rules/{{rule_name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ResponsePolicyRule %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ResponsePolicyRule %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ResponsePolicyRule %q: %#v", d.Id(), res)
	}

	return resourceDNSResponsePolicyRuleRead(d, meta)
}

func resourceDNSResponsePolicyRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ResponsePolicyRule: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/responsePolicies/{{response_policy}}/rules/{{rule_name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ResponsePolicyRule %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ResponsePolicyRule")
	}

	log.Printf("[DEBUG] Finished deleting ResponsePolicyRule %q: %#v", d.Id(), res)
	return nil
}

func resourceDNSResponsePolicyRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/responsePolicies/(?P<response_policy>[^/]+)/rules/(?P<rule_name>[^/]+)",
		"(?P<project>[^/]+)/(?P<response_policy>[^/]+)/(?P<rule_name>[^/]+)",
		"(?P<response_policy>[^/]+)/(?P<rule_name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/responsePolicies/{{response_policy}}/rules/{{rule_name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDNSResponsePolicyRuleRuleName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSResponsePolicyRuleDnsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSResponsePolicyRuleLocalData(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["local_datas"] =
		flattenDNSResponsePolicyRuleLocalDataLocalDatas(original["localDatas"], d, config)
	return []interface{}{transformed}
}
func flattenDNSResponsePolicyRuleLocalDataLocalDatas(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name":    flattenDNSResponsePolicyRuleLocalDataLocalDatasName(original["name"], d, config),
			"type":    flattenDNSResponsePolicyRuleLocalDataLocalDatasType(original["type"], d, config),
			"ttl":     flattenDNSResponsePolicyRuleLocalDataLocalDatasTtl(original["ttl"], d, config),
			"rrdatas": flattenDNSResponsePolicyRuleLocalDataLocalDatasRrdatas(original["rrdatas"], d, config),
		})
	}
	return transformed
}
func flattenDNSResponsePolicyRuleLocalDataLocalDatasName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSResponsePolicyRuleLocalDataLocalDatasType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSResponsePolicyRuleLocalDataLocalDatasTtl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDNSResponsePolicyRuleLocalDataLocalDatasRrdatas(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSResponsePolicyRuleBehavior(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDNSResponsePolicyRuleRuleName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSResponsePolicyRuleDnsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSResponsePolicyRuleLocalData(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLocalDatas, err := expandDNSResponsePolicyRuleLocalDataLocalDatas(original["local_datas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocalDatas); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["localDatas"] = transformedLocalDatas
	}

	return transformed, nil
}

func expandDNSResponsePolicyRuleLocalDataLocalDatas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDNSResponsePolicyRuleLocalDataLocalDatasName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedType, err := expandDNSResponsePolicyRuleLocalDataLocalDatasType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		transformedTtl, err := expandDNSResponsePolicyRuleLocalDataLocalDatasTtl(original["ttl"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTtl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ttl"] = transformedTtl
		}

		transformedRrdatas, err := expandDNSResponsePolicyRuleLocalDataLocalDatasRrdatas(original["rrdatas"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRrdatas); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["rrdatas"] = transformedRrdatas
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSResponsePolicyRuleLocalDataLocalDatasName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSResponsePolicyRuleLocalDataLocalDatasType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSResponsePolicyRuleLocalDataLocalDatasTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSResponsePolicyRuleLocalDataLocalDatasRrdatas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSResponsePolicyRuleBehavior(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
