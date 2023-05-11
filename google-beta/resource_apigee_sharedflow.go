// ----------------------------------------------------------------------------
//
//     This file is partially automatically generated by Magic Modules and with manual
//     changes to resourceApigeeSharedFlowCreate
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"google.golang.org/api/googleapi"
)

func ResourceApigeeSharedFlow() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeSharedFlowCreate,
		Read:   resourceApigeeSharedFlowRead,
		Update: resourceApigeeSharedFlowUpdate,
		Delete: resourceApigeeSharedFlowDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeSharedFlowImport,
		},

		CustomizeDiff: customdiff.All(
			/*
				If any of the config_bundle, detect_md5hash or md5hash is changed,
				then an update is expected, so we tell Terraform core to expect update on meta_data,
				latest_revision_id and revision
			*/

			customdiff.ComputedIf("meta_data", apigeeSharedflowDetectBundleUpdate),
			customdiff.ComputedIf("latest_revision_id", apigeeSharedflowDetectBundleUpdate),
			customdiff.ComputedIf("revision", apigeeSharedflowDetectBundleUpdate),
		),

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the shared flow.`,
			},
			"org_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The Apigee Organization name associated with the Apigee instance.`,
			},
			"latest_revision_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The id of the most recently created revision for this shared flow.`,
			},
			"meta_data": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Metadata describing the shared flow.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Time at which the API proxy was created, in milliseconds since epoch.`,
						},
						"last_modified_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Time at which the API proxy was most recently modified, in milliseconds since epoch.`,
						},
						"sub_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The type of entity described`,
						},
					},
				},
			},
			"revision": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `A list of revisions of this shared flow.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"config_bundle": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Path to the config zip bundle`,
			},
			"md5hash": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Base 64 MD5 hash of the uploaded config bundle.`,
			},
			"detect_md5hash": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Different Hash",
				Description: `A hash of local config bundle in string, user needs to use a Terraform Hash function of their choice. A change in hash will trigger an update.`,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					localMd5Hash := ""
					if config_bundle, ok := d.GetOkExists("config_bundle"); ok {
						localMd5Hash = getFileMd5Hash(config_bundle.(string))
					}
					if localMd5Hash == "" {
						return false
					}

					// `old` is the md5 hash we speculated from server responses,
					// when apply responded with succeed, hash is set to the hash of uploaded bundle
					if old != localMd5Hash {
						return false
					}

					return true
				},
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeSharedFlowCreate(d *schema.ResourceData, meta interface{}) error {
	ctx := context.TODO()
	tflog.Info(ctx, "resourceApigeeSharedFlowCreate")
	log.Printf("[DEBUG] resourceApigeeSharedFlowCreate")

	log.Printf("[DEBUG] resourceApigeeSharedFlowCreate, name=			 	%s", d.Get("name").(string))
	log.Printf("[DEBUG] resourceApigeeSharedFlowCreate, org_id=, 			%s", d.Get("org_id").(string))
	log.Printf("[DEBUG] resourceApigeeSharedFlowCreate, config_bundle=, 	%s", d.Get("config_bundle").(string))

	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	var file *os.File
	var localMd5Hash string
	if configBundlePath, ok := d.GetOk("config_bundle"); ok {
		var err error
		file, err = os.Open(configBundlePath.(string))
		if err != nil {
			return err
		}
		localMd5Hash = getFileMd5Hash(configBundlePath.(string))
	} else {
		return fmt.Errorf("Error, \"config_bundle\" must be specified")
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org_id}}/sharedflows?name={{name}}&action=import")
	if err != nil {
		return err
	}
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	log.Printf("[DEBUG] resourceApigeeSharedFlowCreate, url=, 	%s", url)
	res, err := sendRequestRawBodyWithTimeout(config, "POST", billingProject, url, userAgent, file, "application/octet-stream", d.Timeout(schema.TimeoutCreate))

	log.Printf("[DEBUG] sendRequestRawBodyWithTimeout Done")
	if err != nil {
		return fmt.Errorf("Error creating SharedFlow: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{org_id}}/sharedflows/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	log.Printf("[DEBUG] create d.SetId done, id = %s", id)

	log.Printf("[DEBUG] Finished creating SharedFlow %q: %#v", d.Id(), res)

	if resourceApigeeSharedFlowRead(d, meta) != nil {
		return fmt.Errorf("Error reading SharedFlow at end of Create: %s", err)
	}
	d.Set("md5hash", localMd5Hash)
	d.Set("detect_md5hash", localMd5Hash)
	return nil
}

func resourceApigeeSharedFlowUpdate(d *schema.ResourceData, meta interface{}) error {
	//For how sharedflow api is implemented, just treat an update as create, when the name is same, it will create a new revision
	return resourceApigeeSharedFlowCreate(d, meta)
}

func resourceApigeeSharedFlowRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org_id}}/sharedflows/{{name}}")
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] sharedflow read url is: %s", url)

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	log.Printf("[DEBUG] resourceApigeeSharedFlowRead sendRequest")
	log.Printf("[DEBUG] resourceApigeeSharedFlowRead, url=, 	%s", url)
	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ApigeeSharedFlow %q", d.Id()))
	}
	log.Printf("[DEBUG] resourceApigeeSharedFlowRead sendRequest completed")
	previousLastModifiedAt := getApigeeSharedFlowLastModifiedAt(d)
	if err := d.Set("meta_data", flattenApigeeSharedFlowMetaData(res["metaData"], d, config)); err != nil {
		return fmt.Errorf("Error reading SharedFlow: %s", err)
	}
	currentLastModifiedAt := getApigeeSharedFlowLastModifiedAt(d)
	if err := d.Set("name", flattenApigeeSharedFlowName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading SharedFlow: %s", err)
	}
	if err := d.Set("revision", flattenApigeeSharedFlowRevision(res["revision"], d, config)); err != nil {
		return fmt.Errorf("Error reading SharedFlow: %s", err)
	}
	if err := d.Set("latest_revision_id", flattenApigeeSharedFlowLatestRevisionId(res["latestRevisionId"], d, config)); err != nil {
		return fmt.Errorf("Error reading SharedFlow: %s", err)
	}

	//setting hash to suggest update
	if previousLastModifiedAt != currentLastModifiedAt {
		d.Set("md5hash", "UNKNOWN")
		d.Set("detect_md5hash", "UNKNOWN")
	}
	return nil
}

func getApigeeSharedFlowLastModifiedAt(d *schema.ResourceData) string {

	metaDataRaw := d.Get("meta_data").([]interface{})
	if len(metaDataRaw) != 1 {
		//in Terraform Schema, a nest in object is implemented as an array of length one, even if it's technically an object
		return "UNKNOWN"
	}
	metaData := metaDataRaw[0].(map[string]interface{})
	if metaData == nil {
		return "UNKNOWN"
	}
	lastModifiedAt := metaData["last_modified_at"].(string)
	if lastModifiedAt == "" {
		return "UNKNOWN"
	}
	return lastModifiedAt
}

func resourceApigeeSharedFlowDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceApigeeSharedFlowDelete")
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org_id}}/sharedflows/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting SharedFlow %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "SharedFlow")
	}

	log.Printf("[DEBUG] Finished deleting SharedFlow %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeSharedFlowImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"organizations/(?P<org_id>[^/]+)/sharedflows/(?P<name>[^/]+)",
		"(?P<org_id>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{org_id}}/sharedflows/{{name}}")

	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	log.Printf("[DEBUG] resourceApigeeSharedFlowImport, id=			 	%s", id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeSharedFlowMetaData(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["created_at"] =
		flattenApigeeSharedFlowMetaDataCreatedAt(original["createdAt"], d, config)
	transformed["last_modified_at"] =
		flattenApigeeSharedFlowMetaDataLastModifiedAt(original["lastModifiedAt"], d, config)
	transformed["sub_type"] =
		flattenApigeeSharedFlowMetaDataSubType(original["subType"], d, config)
	return []interface{}{transformed}
}
func flattenApigeeSharedFlowMetaDataCreatedAt(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeSharedFlowMetaDataLastModifiedAt(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeSharedFlowMetaDataSubType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeSharedFlowName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeSharedFlowRevision(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeSharedFlowLatestRevisionId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApigeeSharedFlowName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

// sendRequestRawBodyWithTimeout is derived from sendRequestWithTimeout with direct pass through of request body
func sendRequestRawBodyWithTimeout(config *transport_tpg.Config, method, project, rawurl, userAgent string, body io.Reader, contentType string, timeout time.Duration, errorRetryPredicates ...transport_tpg.RetryErrorPredicateFunc) (map[string]interface{}, error) {
	log.Printf("[DEBUG] sendRequestRawBodyWithTimeout start")
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", userAgent)
	reqHeaders.Set("Content-Type", contentType)

	if config.UserProjectOverride && project != "" {
		// Pass the project into this fn instead of parsing it from the URL because
		// both project names and URLs can have colons in them.
		reqHeaders.Set("X-Goog-User-Project", project)
	}

	if timeout == 0 {
		timeout = time.Duration(1) * time.Minute
	}

	var res *http.Response

	log.Printf("[DEBUG] sendRequestRawBodyWithTimeout sending request")

	err := transport_tpg.RetryTimeDuration(
		func() error {
			req, err := http.NewRequest(method, rawurl, body)
			if err != nil {
				return err
			}

			req.Header = reqHeaders
			res, err = config.Client.Do(req)
			if err != nil {
				return err
			}

			if err := googleapi.CheckResponse(res); err != nil {
				googleapi.CloseBody(res)
				return err
			}

			return nil
		},
		timeout,
		errorRetryPredicates...,
	)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, fmt.Errorf("Unable to parse server response. This is most likely a terraform problem, please file a bug at https://github.com/hashicorp/terraform-provider-google/issues.")
	}

	// The defer call must be made outside of the retryFunc otherwise it's closed too soon.
	defer googleapi.CloseBody(res)

	// 204 responses will have no body, so we're going to error with "EOF" if we
	// try to parse it. Instead, we can just return nil.
	if res.StatusCode == 204 {
		return nil, nil
	}
	result := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] sendRequestRawBodyWithTimeout returning")
	return result, nil
}

func apigeeSharedflowDetectBundleUpdate(_ context.Context, diff *schema.ResourceDiff, v interface{}) bool {
	tmp, _ := diff.GetChange("detect_md5hash")
	oldBundleHash := tmp.(string)
	currentBundleHash := ""
	if config_bundle, ok := diff.GetOkExists("config_bundle"); ok {
		currentBundleHash = getFileMd5Hash(config_bundle.(string))
	}
	log.Printf("[DEBUG] apigeeSharedflowDetectUpdate detect_md5hash: %s -> %s", oldBundleHash, currentBundleHash)

	if oldBundleHash != currentBundleHash {
		return true
	}
	return diff.HasChange("config_bundle") || diff.HasChange("md5hash")
}
