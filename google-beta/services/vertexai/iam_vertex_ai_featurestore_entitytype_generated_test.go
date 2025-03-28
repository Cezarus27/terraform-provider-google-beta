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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/vertexai/FeaturestoreEntitytype.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/examples/base_configs/iam_test_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package vertexai_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccVertexAIFeaturestoreEntitytypeIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"role":            "roles/viewer",
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),

		"kms_key_name": acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeaturestoreEntitytypeIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_entitytype_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("%s/entityTypes/%s roles/viewer", fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccVertexAIFeaturestoreEntitytypeIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_entitytype_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("%s/entityTypes/%s roles/viewer", fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVertexAIFeaturestoreEntitytypeIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"role":            "roles/viewer",
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),

		"kms_key_name": acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccVertexAIFeaturestoreEntitytypeIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_entitytype_iam_member.foo",
				ImportStateId:     fmt.Sprintf("%s/entityTypes/%s roles/viewer user:admin@hashicorptest.com", fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVertexAIFeaturestoreEntitytypeIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"role":            "roles/viewer",
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),

		"kms_key_name": acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeaturestoreEntitytypeIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_vertex_ai_featurestore_entitytype_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_entitytype_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("%s/entityTypes/%s", fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVertexAIFeaturestoreEntitytypeIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_entitytype_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("%s/entityTypes/%s", fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVertexAIFeaturestoreEntitytypeIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
}

resource "google_vertex_ai_featurestore_entitytype" "entity" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  description = "test description"
  featurestore = google_vertex_ai_featurestore.featurestore.id
  monitoring_config {
    snapshot_analysis {
      disabled = false
      monitoring_interval_days = 1
      staleness_days = 21
    }
    numerical_threshold_config {
      value = 0.8
    }
    categorical_threshold_config {
      value = 10.0
    }
    import_features_analysis {
      state = "ENABLED"
      anomaly_detection_baseline = "PREVIOUS_IMPORT_FEATURES_STATS"
    }
  }
}

resource "google_vertex_ai_featurestore_entitytype_iam_member" "foo" {
  featurestore = google_vertex_ai_featurestore_entitytype.entity.featurestore
  entitytype = google_vertex_ai_featurestore_entitytype.entity.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccVertexAIFeaturestoreEntitytypeIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
}

resource "google_vertex_ai_featurestore_entitytype" "entity" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  description = "test description"
  featurestore = google_vertex_ai_featurestore.featurestore.id
  monitoring_config {
    snapshot_analysis {
      disabled = false
      monitoring_interval_days = 1
      staleness_days = 21
    }
    numerical_threshold_config {
      value = 0.8
    }
    categorical_threshold_config {
      value = 10.0
    }
    import_features_analysis {
      state = "ENABLED"
      anomaly_detection_baseline = "PREVIOUS_IMPORT_FEATURES_STATS"
    }
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_vertex_ai_featurestore_entitytype_iam_policy" "foo" {
  featurestore = google_vertex_ai_featurestore_entitytype.entity.featurestore
  entitytype = google_vertex_ai_featurestore_entitytype.entity.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_vertex_ai_featurestore_entitytype_iam_policy" "foo" {
  featurestore = google_vertex_ai_featurestore_entitytype.entity.featurestore
  entitytype = google_vertex_ai_featurestore_entitytype.entity.name
  depends_on = [
    google_vertex_ai_featurestore_entitytype_iam_policy.foo
  ]
}
`, context)
}

func testAccVertexAIFeaturestoreEntitytypeIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
}

resource "google_vertex_ai_featurestore_entitytype" "entity" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  description = "test description"
  featurestore = google_vertex_ai_featurestore.featurestore.id
  monitoring_config {
    snapshot_analysis {
      disabled = false
      monitoring_interval_days = 1
      staleness_days = 21
    }
    numerical_threshold_config {
      value = 0.8
    }
    categorical_threshold_config {
      value = 10.0
    }
    import_features_analysis {
      state = "ENABLED"
      anomaly_detection_baseline = "PREVIOUS_IMPORT_FEATURES_STATS"
    }
  }
}

data "google_iam_policy" "foo" {
}

resource "google_vertex_ai_featurestore_entitytype_iam_policy" "foo" {
  featurestore = google_vertex_ai_featurestore_entitytype.entity.featurestore
  entitytype = google_vertex_ai_featurestore_entitytype.entity.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccVertexAIFeaturestoreEntitytypeIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
}

resource "google_vertex_ai_featurestore_entitytype" "entity" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  description = "test description"
  featurestore = google_vertex_ai_featurestore.featurestore.id
  monitoring_config {
    snapshot_analysis {
      disabled = false
      monitoring_interval_days = 1
      staleness_days = 21
    }
    numerical_threshold_config {
      value = 0.8
    }
    categorical_threshold_config {
      value = 10.0
    }
    import_features_analysis {
      state = "ENABLED"
      anomaly_detection_baseline = "PREVIOUS_IMPORT_FEATURES_STATS"
    }
  }
}

resource "google_vertex_ai_featurestore_entitytype_iam_binding" "foo" {
  featurestore = google_vertex_ai_featurestore_entitytype.entity.featurestore
  entitytype = google_vertex_ai_featurestore_entitytype.entity.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccVertexAIFeaturestoreEntitytypeIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
}

resource "google_vertex_ai_featurestore_entitytype" "entity" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  description = "test description"
  featurestore = google_vertex_ai_featurestore.featurestore.id
  monitoring_config {
    snapshot_analysis {
      disabled = false
      monitoring_interval_days = 1
      staleness_days = 21
    }
    numerical_threshold_config {
      value = 0.8
    }
    categorical_threshold_config {
      value = 10.0
    }
    import_features_analysis {
      state = "ENABLED"
      anomaly_detection_baseline = "PREVIOUS_IMPORT_FEATURES_STATS"
    }
  }
}

resource "google_vertex_ai_featurestore_entitytype_iam_binding" "foo" {
  featurestore = google_vertex_ai_featurestore_entitytype.entity.featurestore
  entitytype = google_vertex_ai_featurestore_entitytype.entity.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
