---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/contactcenterinsights/AnalysisRule.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Contact Center AI Insights"
description: |-
  The CCAI Insights project wide analysis rule.
---

# google_contact_center_insights_analysis_rule

The CCAI Insights project wide analysis rule.
This rule will be applied to all conversations that match the filter defined in the rule.
For a conversation matches the filter, the annotators specified in the rule will be run.
If a conversation matches multiple rules, a union of all the annotators will be run.
One project can have multiple analysis rules.


To get more information about AnalysisRule, see:

* [API documentation](https://cloud.google.com/contact-center/insights/docs/reference/rest/v1/projects.locations.analysisRules)
* How-to Guides
    * [Configure analysis rules using the API](https://cloud.google.com/contact-center/insights/docs/analysis-rule)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=contact_center_insights_analysis_rule_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Contact Center Insights Analysis Rule Basic


```hcl
resource "google_contact_center_insights_analysis_rule" "analysis_rule_basic" {
  location = "us-central1"
  display_name = "analysis-rule-display-name"
}
```
## Example Usage - Contact Center Insights Analysis Rule Full


```hcl
resource "google_contact_center_insights_analysis_rule" "analysis_rule_full" {
  location = "us-central1"
  display_name = "analysis-rule-display-name"
  conversation_filter = "agent_id = \"1\""
  annotator_selector {
    run_interruption_annotator = false
    issue_models    = ["projects/1111111111111/locations/us-central1/issueModels/some_issue_model_id"]
    phrase_matchers = ["projects/1111111111111/locations/us-central1/phraseMatchers/123"]
    qa_config {
      scorecard_list {
        qa_scorecard_revisions = ["projects/1111111111111/locations/us-central1/qaScorecards/*/revisions/some_scorecard_revision_id"]
      }
    }
    run_entity_annotator         = false
    run_intent_annotator         = false
    run_issue_model_annotator    = false
    run_phrase_matcher_annotator = false
    run_qa_annotator             = false
    run_sentiment_annotator      = false
    run_silence_annotator        = true
    run_summarization_annotator  = false
    summarization_config {
      summarization_model  = "BASELINE_MODEL"
    }
  }
  analysis_percentage = 0.5
  active    = true
}
```
## Example Usage - Contact Center Insights Analysis Rule Profile


```hcl
resource "google_contact_center_insights_analysis_rule" "analysis_rule_profile" {
  location = "us-central1"
  display_name = "analysis-rule-display-name"
  conversation_filter = "agent_id = \"1\""
  annotator_selector {
    run_interruption_annotator = false
    issue_models    = ["projects/1111111111111/locations/us-central1/issueModels/some_issue_model_id"]
    phrase_matchers = ["projects/1111111111111/locations/us-central1/phraseMatchers/123"]
    qa_config {
      scorecard_list {
        qa_scorecard_revisions = ["projects/1111111111111/locations/us-central1/qaScorecards/*/revisions/some_scorecard_revision_id"]
      }
    }
    run_entity_annotator         = false
    run_intent_annotator         = false
    run_issue_model_annotator    = false
    run_phrase_matcher_annotator = false
    run_qa_annotator             = false
    run_sentiment_annotator      = false
    run_silence_annotator        = true
    run_summarization_annotator  = false
    summarization_config {
      conversation_profile = "projects/1111111111111/locations/us-central1/conversationProfiles/some_conversation_profile"
    }
  }
  analysis_percentage = 0.5
  active    = true
}
```

## Argument Reference

The following arguments are supported:


* `location` -
  (Required)
  Location of the resource.


* `display_name` -
  (Optional)
  Display Name of the analysis rule.

* `conversation_filter` -
  (Optional)
  Filter for the conversations that should apply this analysis
  rule. An empty filter means this analysis rule applies to all
  conversations.
  Refer to https://cloud.google.com/contact-center/insights/docs/filtering
  for details.

* `annotator_selector` -
  (Optional)
  Selector of all available annotators and phrase matchers to run.
  Structure is [documented below](#nested_annotator_selector).

* `analysis_percentage` -
  (Optional)
  Percentage of conversations that we should apply this analysis setting
  automatically, between [0, 1]. For example, 0.1 means 10%. Conversations
  are sampled in a determenestic way. The original runtime_percentage &
  upload percentage will be replaced by defining filters on the conversation.

* `active` -
  (Optional)
  If true, apply this rule to conversations. Otherwise, this rule is
  inactive and saved as a draft.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_annotator_selector"></a>The `annotator_selector` block supports:

* `run_interruption_annotator` -
  (Optional)
  Whether to run the interruption annotator.

* `phrase_matchers` -
  (Optional)
  The list of phrase matchers to run. If not provided, all active phrase
  matchers will be used. If inactive phrase matchers are provided, they will
  not be used. Phrase matchers will be run only if
  run_phrase_matcher_annotator is set to true. Format:
  projects/{project}/locations/{location}/phraseMatchers/{phrase_matcher}

* `run_entity_annotator` -
  (Optional)
  Whether to run the entity annotator.

* `issue_models` -
  (Optional)
  The issue model to run. If not provided, the most recently deployed topic
  model will be used. The provided issue model will only be used for
  inference if the issue model is deployed and if run_issue_model_annotator
  is set to true. If more than one issue model is provided, only the first
  provided issue model will be used for inference.

* `run_qa_annotator` -
  (Optional)
  Whether to run the QA annotator.

* `run_silence_annotator` -
  (Optional)
  Whether to run the silence annotator.

* `run_phrase_matcher_annotator` -
  (Optional)
  Whether to run the active phrase matcher annotator(s).

* `run_sentiment_annotator` -
  (Optional)
  Whether to run the sentiment annotator.

* `run_intent_annotator` -
  (Optional)
  Whether to run the intent annotator.

* `run_issue_model_annotator` -
  (Optional)
  Whether to run the issue model annotator. A model should have already been
  deployed for this to take effect.

* `run_summarization_annotator` -
  (Optional)
  Whether to run the summarization annotator.

* `summarization_config` -
  (Optional)
  Configuration for summarization.
  Structure is [documented below](#nested_annotator_selector_summarization_config).

* `qa_config` -
  (Optional)
  Configuration for the QA feature.
  Structure is [documented below](#nested_annotator_selector_qa_config).


<a name="nested_annotator_selector_summarization_config"></a>The `summarization_config` block supports:

* `conversation_profile` -
  (Optional)
  Resource name of the Dialogflow conversation profile.
  Format:
  projects/{project}/locations/{location}/conversationProfiles/{conversation_profile}

* `summarization_model` -
  (Optional)
  Default summarization model to be used.
  Possible values:
  SUMMARIZATION_MODEL_UNSPECIFIED
  BASELINE_MODEL
  BASELINE_MODEL_V2_0
  Possible values are: `BASELINE_MODEL`, `BASELINE_MODEL_V2_0`.

<a name="nested_annotator_selector_qa_config"></a>The `qa_config` block supports:

* `scorecard_list` -
  (Optional)
  Container for a list of scorecards.
  Structure is [documented below](#nested_annotator_selector_qa_config_scorecard_list).


<a name="nested_annotator_selector_qa_config_scorecard_list"></a>The `scorecard_list` block supports:

* `qa_scorecard_revisions` -
  (Optional)
  List of QaScorecardRevisions.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/analysisRules/{{name}}`

* `name` -
  The resource name of the analysis rule. Randomly generated by Insights.

* `create_time` -
  Output only. The time at which this analysis rule was created.

* `update_time` -
  Output only. The most recent time at which this analysis rule was updated.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


AnalysisRule can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/analysisRules/{{name}}`
* `{{project}}/{{location}}/{{name}}`
* `{{location}}/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import AnalysisRule using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/analysisRules/{{name}}"
  to = google_contact_center_insights_analysis_rule.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), AnalysisRule can be imported using one of the formats above. For example:

```
$ terraform import google_contact_center_insights_analysis_rule.default projects/{{project}}/locations/{{location}}/analysisRules/{{name}}
$ terraform import google_contact_center_insights_analysis_rule.default {{project}}/{{location}}/{{name}}
$ terraform import google_contact_center_insights_analysis_rule.default {{location}}/{{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
