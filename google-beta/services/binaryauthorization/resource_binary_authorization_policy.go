// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/binaryauthorization/Policy.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package binaryauthorization

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func DefaultBinaryAuthorizationPolicy(project string) map[string]interface{} {
	return map[string]interface{}{
		"name": fmt.Sprintf("projects/%s/policy", project),
		"admissionWhitelistPatterns": []interface{}{
			map[string]interface{}{
				"namePattern": "gcr.io/google_containers/*",
			},
		},
		"defaultAdmissionRule": map[string]interface{}{
			"evaluationMode":  "ALWAYS_ALLOW",
			"enforcementMode": "ENFORCED_BLOCK_AND_AUDIT_LOG",
		},
	}
}

func ResourceBinaryAuthorizationPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceBinaryAuthorizationPolicyCreate,
		Read:   resourceBinaryAuthorizationPolicyRead,
		Update: resourceBinaryAuthorizationPolicyUpdate,
		Delete: resourceBinaryAuthorizationPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBinaryAuthorizationPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"default_admission_rule": {
				Type:     schema.TypeList,
				Required: true,
				Description: `Default admission rule for a cluster without a per-cluster admission
rule.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enforcement_mode": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: verify.ValidateEnum([]string{"ENFORCED_BLOCK_AND_AUDIT_LOG", "DRYRUN_AUDIT_LOG_ONLY"}),
							Description:  `The action when a pod creation is denied by the admission rule. Possible values: ["ENFORCED_BLOCK_AND_AUDIT_LOG", "DRYRUN_AUDIT_LOG_ONLY"]`,
						},
						"evaluation_mode": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: verify.ValidateEnum([]string{"ALWAYS_ALLOW", "REQUIRE_ATTESTATION", "ALWAYS_DENY"}),
							Description:  `How this admission rule will be evaluated. Possible values: ["ALWAYS_ALLOW", "REQUIRE_ATTESTATION", "ALWAYS_DENY"]`,
						},
						"require_attestations_by": {
							Type:             schema.TypeSet,
							Optional:         true,
							DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
							Description: `The resource names of the attestors that must attest to a
container image. If the attestor is in a different project from the
policy, it should be specified in the format 'projects/*/attestors/*'.
Each attestor must exist before a policy can reference it. To add an
attestor to a policy the principal issuing the policy change
request must be able to read the attestor resource.

Note: this field must be non-empty when the evaluation_mode field
specifies REQUIRE_ATTESTATION, otherwise it must be empty.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: tpgresource.SelfLinkNameHash,
						},
					},
				},
			},
			"admission_whitelist_patterns": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A whitelist of image patterns to exclude from admission rules. If an
image's name matches a whitelist pattern, the image's admission
requests will always be permitted regardless of your admission rules.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name_pattern": {
							Type:     schema.TypeString,
							Required: true,
							Description: `An image name pattern to whitelist, in the form
'registry/path/to/image'. This supports a trailing * as a
wildcard, but this is allowed only in text after the registry/
part.`,
						},
					},
				},
			},
			"cluster_admission_rules": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `Per-cluster admission rules. An admission rule specifies either that
all container images used in a pod creation request must be attested
to by one or more attestors, that all pod creations will be allowed,
or that all pod creations will be denied. There can be at most one
admission rule per cluster spec.


Identifier format: '{{location}}.{{clusterId}}'.
A location is either a compute zone (e.g. 'us-central1-a') or a region
(e.g. 'us-central1').`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster": {
							Type:     schema.TypeString,
							Required: true,
						},
						"enforcement_mode": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: verify.ValidateEnum([]string{"ENFORCED_BLOCK_AND_AUDIT_LOG", "DRYRUN_AUDIT_LOG_ONLY"}),
							Description:  `The action when a pod creation is denied by the admission rule. Possible values: ["ENFORCED_BLOCK_AND_AUDIT_LOG", "DRYRUN_AUDIT_LOG_ONLY"]`,
						},
						"evaluation_mode": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: verify.ValidateEnum([]string{"ALWAYS_ALLOW", "REQUIRE_ATTESTATION", "ALWAYS_DENY"}),
							Description:  `How this admission rule will be evaluated. Possible values: ["ALWAYS_ALLOW", "REQUIRE_ATTESTATION", "ALWAYS_DENY"]`,
						},
						"require_attestations_by": {
							Type:             schema.TypeSet,
							Optional:         true,
							DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
							Description: `The resource names of the attestors that must attest to a
container image. If the attestor is in a different project from the
policy, it should be specified in the format 'projects/*/attestors/*'.
Each attestor must exist before a policy can reference it. To add an
attestor to a policy the principal issuing the policy change
request must be able to read the attestor resource.

Note: this field must be non-empty when the evaluation_mode field
specifies REQUIRE_ATTESTATION, otherwise it must be empty.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: tpgresource.SelfLinkNameHash,
						},
					},
				},
				Set: func(v interface{}) int {
					// require_attestations_by is a set of strings that can have the format
					// projects/{project}/attestors/{attestor} or {attestor}. We diffsuppress
					// and hash that set on the name, but now we need to make sure that the
					// overall hash here respects that so changing the attestor format doesn't
					// change the hash code of cluster_admission_rules.
					raw := v.(map[string]interface{})

					// modifying raw actually modifies the values passed to the provider.
					// Use a copy to avoid that.
					copy := make((map[string]interface{}))
					for key, value := range raw {
						copy[key] = value
					}
					at := copy["require_attestations_by"].(*schema.Set)
					if at != nil {
						t := tpgresource.ConvertAndMapStringArr(at.List(), tpgresource.GetResourceNameFromSelfLink)
						copy["require_attestations_by"] = schema.NewSet(tpgresource.SelfLinkNameHash, tpgresource.ConvertStringArrToInterface(t))
					}
					var buf bytes.Buffer
					schema.SerializeResourceForHash(&buf, copy, ResourceBinaryAuthorizationPolicy().Schema["cluster_admission_rules"].Elem.(*schema.Resource))
					return tpgresource.Hashcode(buf.String())
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A descriptive comment.`,
			},
			"global_policy_evaluation_mode": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ENABLE", "DISABLE", ""}),
				Description: `Controls the evaluation of a Google-maintained global admission policy
for common system-level images. Images not covered by the global
policy will be subject to the project admission policy. Possible values: ["ENABLE", "DISABLE"]`,
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

func resourceBinaryAuthorizationPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandBinaryAuthorizationPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	globalPolicyEvaluationModeProp, err := expandBinaryAuthorizationPolicyGlobalPolicyEvaluationMode(d.Get("global_policy_evaluation_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("global_policy_evaluation_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(globalPolicyEvaluationModeProp)) && (ok || !reflect.DeepEqual(v, globalPolicyEvaluationModeProp)) {
		obj["globalPolicyEvaluationMode"] = globalPolicyEvaluationModeProp
	}
	admissionWhitelistPatternsProp, err := expandBinaryAuthorizationPolicyAdmissionWhitelistPatterns(d.Get("admission_whitelist_patterns"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("admission_whitelist_patterns"); !tpgresource.IsEmptyValue(reflect.ValueOf(admissionWhitelistPatternsProp)) && (ok || !reflect.DeepEqual(v, admissionWhitelistPatternsProp)) {
		obj["admissionWhitelistPatterns"] = admissionWhitelistPatternsProp
	}
	clusterAdmissionRulesProp, err := expandBinaryAuthorizationPolicyClusterAdmissionRules(d.Get("cluster_admission_rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cluster_admission_rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(clusterAdmissionRulesProp)) && (ok || !reflect.DeepEqual(v, clusterAdmissionRulesProp)) {
		obj["clusterAdmissionRules"] = clusterAdmissionRulesProp
	}
	defaultAdmissionRuleProp, err := expandBinaryAuthorizationPolicyDefaultAdmissionRule(d.Get("default_admission_rule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_admission_rule"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultAdmissionRuleProp)) && (ok || !reflect.DeepEqual(v, defaultAdmissionRuleProp)) {
		obj["defaultAdmissionRule"] = defaultAdmissionRuleProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/policy")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Policy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Policy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating Policy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Policy %q: %#v", d.Id(), res)

	return resourceBinaryAuthorizationPolicyRead(d, meta)
}

func resourceBinaryAuthorizationPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/policy")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Policy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BinaryAuthorizationPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}

	if err := d.Set("description", flattenBinaryAuthorizationPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("global_policy_evaluation_mode", flattenBinaryAuthorizationPolicyGlobalPolicyEvaluationMode(res["globalPolicyEvaluationMode"], d, config)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("admission_whitelist_patterns", flattenBinaryAuthorizationPolicyAdmissionWhitelistPatterns(res["admissionWhitelistPatterns"], d, config)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("cluster_admission_rules", flattenBinaryAuthorizationPolicyClusterAdmissionRules(res["clusterAdmissionRules"], d, config)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("default_admission_rule", flattenBinaryAuthorizationPolicyDefaultAdmissionRule(res["defaultAdmissionRule"], d, config)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}

	return nil
}

func resourceBinaryAuthorizationPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Policy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandBinaryAuthorizationPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	globalPolicyEvaluationModeProp, err := expandBinaryAuthorizationPolicyGlobalPolicyEvaluationMode(d.Get("global_policy_evaluation_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("global_policy_evaluation_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, globalPolicyEvaluationModeProp)) {
		obj["globalPolicyEvaluationMode"] = globalPolicyEvaluationModeProp
	}
	admissionWhitelistPatternsProp, err := expandBinaryAuthorizationPolicyAdmissionWhitelistPatterns(d.Get("admission_whitelist_patterns"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("admission_whitelist_patterns"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, admissionWhitelistPatternsProp)) {
		obj["admissionWhitelistPatterns"] = admissionWhitelistPatternsProp
	}
	clusterAdmissionRulesProp, err := expandBinaryAuthorizationPolicyClusterAdmissionRules(d.Get("cluster_admission_rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cluster_admission_rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clusterAdmissionRulesProp)) {
		obj["clusterAdmissionRules"] = clusterAdmissionRulesProp
	}
	defaultAdmissionRuleProp, err := expandBinaryAuthorizationPolicyDefaultAdmissionRule(d.Get("default_admission_rule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_admission_rule"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, defaultAdmissionRuleProp)) {
		obj["defaultAdmissionRule"] = defaultAdmissionRuleProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/policy")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Policy %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating Policy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Policy %q: %#v", d.Id(), res)
	}

	return resourceBinaryAuthorizationPolicyRead(d, meta)
}

func resourceBinaryAuthorizationPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Policy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/policy")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	obj = DefaultBinaryAuthorizationPolicy(d.Get("project").(string))

	log.Printf("[DEBUG] Deleting Policy %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Policy")
	}

	log.Printf("[DEBUG] Finished deleting Policy %q: %#v", d.Id(), res)
	return nil
}

func resourceBinaryAuthorizationPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)$",
		"^(?P<project>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBinaryAuthorizationPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBinaryAuthorizationPolicyGlobalPolicyEvaluationMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBinaryAuthorizationPolicyAdmissionWhitelistPatterns(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"name_pattern": flattenBinaryAuthorizationPolicyAdmissionWhitelistPatternsNamePattern(original["namePattern"], d, config),
		})
	}
	return transformed
}
func flattenBinaryAuthorizationPolicyAdmissionWhitelistPatternsNamePattern(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBinaryAuthorizationPolicyClusterAdmissionRules(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(l))
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"cluster":                 k,
			"evaluation_mode":         flattenBinaryAuthorizationPolicyClusterAdmissionRulesEvaluationMode(original["evaluationMode"], d, config),
			"require_attestations_by": flattenBinaryAuthorizationPolicyClusterAdmissionRulesRequireAttestationsBy(original["requireAttestationsBy"], d, config),
			"enforcement_mode":        flattenBinaryAuthorizationPolicyClusterAdmissionRulesEnforcementMode(original["enforcementMode"], d, config),
		})
	}
	return transformed
}
func flattenBinaryAuthorizationPolicyClusterAdmissionRulesEvaluationMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBinaryAuthorizationPolicyClusterAdmissionRulesRequireAttestationsBy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(tpgresource.SelfLinkNameHash, v.([]interface{}))
}

func flattenBinaryAuthorizationPolicyClusterAdmissionRulesEnforcementMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBinaryAuthorizationPolicyDefaultAdmissionRule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["evaluation_mode"] =
		flattenBinaryAuthorizationPolicyDefaultAdmissionRuleEvaluationMode(original["evaluationMode"], d, config)
	transformed["require_attestations_by"] =
		flattenBinaryAuthorizationPolicyDefaultAdmissionRuleRequireAttestationsBy(original["requireAttestationsBy"], d, config)
	transformed["enforcement_mode"] =
		flattenBinaryAuthorizationPolicyDefaultAdmissionRuleEnforcementMode(original["enforcementMode"], d, config)
	return []interface{}{transformed}
}
func flattenBinaryAuthorizationPolicyDefaultAdmissionRuleEvaluationMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBinaryAuthorizationPolicyDefaultAdmissionRuleRequireAttestationsBy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(tpgresource.SelfLinkNameHash, v.([]interface{}))
}

func flattenBinaryAuthorizationPolicyDefaultAdmissionRuleEnforcementMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandBinaryAuthorizationPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyGlobalPolicyEvaluationMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyAdmissionWhitelistPatterns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNamePattern, err := expandBinaryAuthorizationPolicyAdmissionWhitelistPatternsNamePattern(original["name_pattern"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNamePattern); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["namePattern"] = transformedNamePattern
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBinaryAuthorizationPolicyAdmissionWhitelistPatternsNamePattern(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyClusterAdmissionRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedEvaluationMode, err := expandBinaryAuthorizationPolicyClusterAdmissionRulesEvaluationMode(original["evaluation_mode"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEvaluationMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["evaluationMode"] = transformedEvaluationMode
		}

		transformedRequireAttestationsBy, err := expandBinaryAuthorizationPolicyClusterAdmissionRulesRequireAttestationsBy(original["require_attestations_by"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRequireAttestationsBy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["requireAttestationsBy"] = transformedRequireAttestationsBy
		}

		transformedEnforcementMode, err := expandBinaryAuthorizationPolicyClusterAdmissionRulesEnforcementMode(original["enforcement_mode"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnforcementMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["enforcementMode"] = transformedEnforcementMode
		}

		transformedCluster, err := tpgresource.ExpandString(original["cluster"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedCluster] = transformed
	}
	return m, nil
}

func expandBinaryAuthorizationPolicyClusterAdmissionRulesEvaluationMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyClusterAdmissionRulesRequireAttestationsBy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	r := regexp.MustCompile("projects/(.+)/attestors/(.+)")

	// It's possible that all entries in the list will specify a project, in
	// which case the user wouldn't necessarily have to specify a provider
	// project.
	var project string
	var err error
	for _, s := range v.(*schema.Set).List() {
		if !r.MatchString(s.(string)) {
			project, err = tpgresource.GetProject(d, config)
			if err != nil {
				return []interface{}{}, err
			}
			break
		}
	}

	return tpgresource.ConvertAndMapStringArr(v.(*schema.Set).List(), func(s string) string {
		if r.MatchString(s) {
			return s
		}

		return fmt.Sprintf("projects/%s/attestors/%s", project, s)
	}), nil
}

func expandBinaryAuthorizationPolicyClusterAdmissionRulesEnforcementMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyDefaultAdmissionRule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEvaluationMode, err := expandBinaryAuthorizationPolicyDefaultAdmissionRuleEvaluationMode(original["evaluation_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEvaluationMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["evaluationMode"] = transformedEvaluationMode
	}

	transformedRequireAttestationsBy, err := expandBinaryAuthorizationPolicyDefaultAdmissionRuleRequireAttestationsBy(original["require_attestations_by"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireAttestationsBy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["requireAttestationsBy"] = transformedRequireAttestationsBy
	}

	transformedEnforcementMode, err := expandBinaryAuthorizationPolicyDefaultAdmissionRuleEnforcementMode(original["enforcement_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnforcementMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enforcementMode"] = transformedEnforcementMode
	}

	return transformed, nil
}

func expandBinaryAuthorizationPolicyDefaultAdmissionRuleEvaluationMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyDefaultAdmissionRuleRequireAttestationsBy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	r := regexp.MustCompile("projects/(.+)/attestors/(.+)")

	// It's possible that all entries in the list will specify a project, in
	// which case the user wouldn't necessarily have to specify a provider
	// project.
	var project string
	var err error
	for _, s := range v.(*schema.Set).List() {
		if !r.MatchString(s.(string)) {
			project, err = tpgresource.GetProject(d, config)
			if err != nil {
				return []interface{}{}, err
			}
			break
		}
	}

	return tpgresource.ConvertAndMapStringArr(v.(*schema.Set).List(), func(s string) string {
		if r.MatchString(s) {
			return s
		}

		return fmt.Sprintf("projects/%s/attestors/%s", project, s)
	}), nil
}

func expandBinaryAuthorizationPolicyDefaultAdmissionRuleEnforcementMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
