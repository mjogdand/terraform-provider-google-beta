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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccDataLossPreventionInspectTemplate_dlpInspectTemplateBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionInspectTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionInspectTemplate_dlpInspectTemplateBasicExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_inspect_template.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionInspectTemplate_dlpInspectTemplateBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "My description"
	display_name = "display_name"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
		info_types {
			name = "PERSON_NAME"
		}
		info_types {
			name = "LAST_NAME"
		}
		info_types {
			name = "DOMAIN_NAME"
		}
		info_types {
			name = "PHONE_NUMBER"
		}
		info_types {
			name = "FIRST_NAME"
		}

		min_likelihood = "UNLIKELY"
		rule_set {
			info_types {
				name = "EMAIL_ADDRESS"
			}
			rules {
				exclusion_rule {
					regex {
						pattern = ".+@example.com"
					}
					matching_type = "MATCHING_TYPE_FULL_MATCH"
				}
			}
		}
		rule_set {
			info_types {
				name = "EMAIL_ADDRESS"
			}
			info_types {
				name = "DOMAIN_NAME"
			}
			info_types {
				name = "PHONE_NUMBER"
			}
			info_types {
				name = "PERSON_NAME"
			}
			info_types {
				name = "FIRST_NAME"
			}
			rules {
				exclusion_rule {
					dictionary {
						word_list {
							words = ["TEST"]
						}
					}
					matching_type = "MATCHING_TYPE_PARTIAL_MATCH"
				}
			}
		}

		rule_set {
			info_types {
				name = "PERSON_NAME"
			}
			rules {
				hotword_rule {
					hotword_regex {
						pattern = "patient"
					}
					proximity {
						window_before = 50
					}
					likelihood_adjustment {
						fixed_likelihood = "VERY_LIKELY"
					}
				}
			}
		}

		limits {
			max_findings_per_item    = 10
			max_findings_per_request = 50
			max_findings_per_info_type {
				max_findings = "75"
				info_type {
					name = "PERSON_NAME"
				}
			}
			max_findings_per_info_type {
				max_findings = "80"
				info_type {
					name = "LAST_NAME"
				}
			}
		}
	}
}
`, context)
}

func TestAccDataLossPreventionInspectTemplate_dlpInspectTemplateCustomTypeExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionInspectTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionInspectTemplate_dlpInspectTemplateCustomTypeExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_inspect_template.custom",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionInspectTemplate_dlpInspectTemplateCustomTypeExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_inspect_template" "custom" {
	parent = "projects/%{project}"
	description = "My description"
	display_name = "display_name"

	inspect_config {
		custom_info_types {
			info_type {
				name = "MY_CUSTOM_TYPE"
			}

			likelihood = "UNLIKELY"

			regex {
				pattern = "test*"
			}
		}

		info_types {
			name = "EMAIL_ADDRESS"
		}

		min_likelihood = "UNLIKELY"
		rule_set {
			info_types {
				name = "EMAIL_ADDRESS"
			}
			rules {
				exclusion_rule {
					regex {
						pattern = ".+@example.com"
					}
					matching_type = "MATCHING_TYPE_FULL_MATCH"
				}
			}
		}

		rule_set {
			info_types {
				name = "MY_CUSTOM_TYPE"
			}
			rules {
				hotword_rule {
					hotword_regex {
						pattern = "example*"
					}
					proximity {
						window_before = 50
					}
					likelihood_adjustment {
						fixed_likelihood = "VERY_LIKELY"
					}
				}
			}
		}

		limits {
			max_findings_per_item    = 10
			max_findings_per_request = 50
		}
	}
}
`, context)
}

func TestAccDataLossPreventionInspectTemplate_dlpInspectTemplateCustomTypeSurrogateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionInspectTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionInspectTemplate_dlpInspectTemplateCustomTypeSurrogateExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_inspect_template.custom_type_surrogate",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionInspectTemplate_dlpInspectTemplateCustomTypeSurrogateExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_inspect_template" "custom_type_surrogate" {
  parent = "projects/%{project}"
  description = "My description"
  display_name = "display_name"

  inspect_config {
    custom_info_types {
      info_type {
        name = "MY_CUSTOM_TYPE"
      }

      likelihood = "UNLIKELY"

      surrogate_type {}
    }

    info_types {
      name = "EMAIL_ADDRESS"
    }

    min_likelihood = "UNLIKELY"
    rule_set {
      info_types {
        name = "EMAIL_ADDRESS"
      }
      rules {
        exclusion_rule {
          regex {
            pattern = ".+@example.com"
          }
          matching_type = "MATCHING_TYPE_FULL_MATCH"
        }
      }
    }

    rule_set {
      info_types {
        name = "MY_CUSTOM_TYPE"
      }
      rules {
        hotword_rule {
          hotword_regex {
            pattern = "example*"
          }
          proximity {
            window_before = 50
          }
          likelihood_adjustment {
            fixed_likelihood = "VERY_LIKELY"
          }
        }
      }
    }

    limits {
      max_findings_per_item    = 10
      max_findings_per_request = 50
    }
  }
}
`, context)
}

func testAccCheckDataLossPreventionInspectTemplateDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_loss_prevention_inspect_template" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataLossPreventionBasePath}}{{parent}}/inspectTemplates/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("DataLossPreventionInspectTemplate still exists at %s", url)
			}
		}

		return nil
	}
}
