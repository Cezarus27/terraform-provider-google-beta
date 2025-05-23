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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/servicedirectory/Service.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package servicedirectory

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceServiceDirectoryService() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDirectoryServiceCreate,
		Read:   resourceServiceDirectoryServiceRead,
		Update: resourceServiceDirectoryServiceUpdate,
		Delete: resourceServiceDirectoryServiceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceServiceDirectoryServiceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The resource name of the namespace this service will belong to.`,
			},
			"service_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateRFC1035Name(2, 63),
				Description: `The Resource ID must be 1-63 characters long, including digits,
lowercase letters or the hyphen character.`,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Metadata for the service. This data can be consumed
by service clients. The entire metadata dictionary may contain
up to 2000 characters, spread across all key-value pairs.
Metadata that goes beyond any these limits will be rejected.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name for the service in the
format 'projects/*/locations/*/namespaces/*/services/*'.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceServiceDirectoryServiceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	metadataProp, err := expandServiceDirectoryServiceMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !tpgresource.IsEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ServiceDirectoryBasePath}}{{namespace}}/services?serviceId={{service_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Service: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating Service: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	err = resourceServiceDirectoryServicePostCreateSetComputedFields(d, meta, res)
	if err != nil {
		return fmt.Errorf("setting computed ID format fields: %w", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Service %q: %#v", d.Id(), res)

	return resourceServiceDirectoryServiceRead(d, meta)
}

func resourceServiceDirectoryServiceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ServiceDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ServiceDirectoryService %q", d.Id()))
	}

	if err := d.Set("name", flattenServiceDirectoryServiceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}
	if err := d.Set("metadata", flattenServiceDirectoryServiceMetadata(res["metadata"], d, config)); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}

	return nil
}

func resourceServiceDirectoryServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	metadataProp, err := expandServiceDirectoryServiceMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ServiceDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Service %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("metadata") {
		updateMask = append(updateMask, "metadata")
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

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating Service %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Service %q: %#v", d.Id(), res)
		}

	}

	return resourceServiceDirectoryServiceRead(d, meta)
}

func resourceServiceDirectoryServiceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ServiceDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Service %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Service")
	}

	log.Printf("[DEBUG] Finished deleting Service %q: %#v", d.Id(), res)
	return nil
}

func resourceServiceDirectoryServiceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	nameParts := strings.Split(d.Get("name").(string), "/")
	if len(nameParts) == 8 {
		// `projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}`
		if err := d.Set("namespace", fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", nameParts[1], nameParts[3], nameParts[5])); err != nil {
			return nil, fmt.Errorf("Error setting namespace: %s", err)
		}
		if err := d.Set("service_id", nameParts[7]); err != nil {
			return nil, fmt.Errorf("Error setting service_id: %s", err)
		}
	} else if len(nameParts) == 4 {
		// `{{project}}/{{location}}/{{namespace_id}}/{{service_id}}`
		if err := d.Set("namespace", fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", nameParts[0], nameParts[1], nameParts[2])); err != nil {
			return nil, fmt.Errorf("Error setting namespace: %s", err)
		}
		if err := d.Set("service_id", nameParts[3]); err != nil {
			return nil, fmt.Errorf("Error setting service_id: %s", err)
		}
		id := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", nameParts[0], nameParts[1], nameParts[2], nameParts[3])
		if err := d.Set("name", id); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
		d.SetId(id)
	} else if len(nameParts) == 3 {
		// `{{location}}/{{namespace_id}}/{{service_id}}`
		project, err := tpgresource.GetProject(d, config)
		if err != nil {
			return nil, err
		}
		if err := d.Set("namespace", fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", project, nameParts[0], nameParts[1])); err != nil {
			return nil, fmt.Errorf("Error setting namespace: %s", err)
		}
		if err := d.Set("service_id", nameParts[2]); err != nil {
			return nil, fmt.Errorf("Error setting service_id: %s", err)
		}
		id := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", project, nameParts[0], nameParts[1], nameParts[2])
		if err := d.Set("name", id); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
		d.SetId(id)
	} else {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s, %s or %s",
			d.Get("name"),
			"projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}",
			"{{project}}/{{location}}/{{namespace_id}}/{{service_id}}",
			"{{location}}/{{namespace_id}}/{{service_id}}")
	}
	return []*schema.ResourceData{d}, nil

}

func flattenServiceDirectoryServiceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenServiceDirectoryServiceMetadata(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandServiceDirectoryServiceMetadata(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func resourceServiceDirectoryServicePostCreateSetComputedFields(d *schema.ResourceData, meta interface{}, res map[string]interface{}) error {
	config := meta.(*transport_tpg.Config)
	if err := d.Set("name", flattenServiceDirectoryServiceName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}
	return nil
}
