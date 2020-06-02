/*
 * Copyright The Spinnaker Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package parse_hal

import (
	"github.com/spinnaker/kleat/api/client/cloudprovider"
	"github.com/spinnaker/kleat/api/client/config"
)

// TODO: Actually get Gate's URL (which may be http or https)
// Should Kleat output the service discovery file?
const gateUrl = "https://gate.spinnaker:8084"

func HalToDeck(h *config.Hal) *config.Deck {
	return &config.Deck{
		GateUrl:         gateUrl,
		AuthEnabled:     h.GetSecurity().GetAuthn().GetEnabled(),
		AuthEndpoint:    gateUrl + "/auth/user",
		BakeryDetailUrl: gateUrl + "/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
		Canary:          getDeckCanaryConfig(h),
		Changelog:       getDeckChangelogConfig(h),
		DefaultTimeZone: h.GetTimezone(),
		Features:        getDeckFeaturesConfig(h),
		Notifications:   getDeckNotificationsConfig(h),
		Providers:       getDeckProvidersConfig(h),
	}
}

func getDeckCanaryConfig(h *config.Hal) *config.Deck_Canary {
	return &config.Deck_Canary{
		DefaultJudge:       h.GetCanary().GetDefaultJudge(),
		FeatureDisabled:    !h.GetCanary().GetEnabled(),
		MetricsAccountName: h.GetCanary().GetDefaultMetricsAccount(),
		MetricStore:        h.GetCanary().GetDefaultMetricsStore(),
		ShowAllConfigs:     h.GetCanary().GetShowAllConfigsEnabled(),
		StagesEnabled:      h.GetCanary().GetStagesEnabled(),
		StorageAccountName: h.GetCanary().GetStorageAccountName(),
		TemplatesEnabled:   h.GetCanary().GetTemplatesEnabled(),
	}
}

// TODO: Actually get changelog config
func getDeckChangelogConfig(h *config.Hal) *config.Deck_Changelog {
	return &config.Deck_Changelog{
		FileName: "",
		GistId:   "",
	}
}

func getDeckFeaturesConfig(h *config.Hal) *config.Deck_Features {
	return &config.Deck_Features{
		PipelineTemplates: h.GetFeatures().GetPipelineTemplates(),
		Canary:            h.GetFeatures().GetMineCanary(),
		ChaosMonkey:       h.GetFeatures().GetChaos(),
		FiatEnabled:       h.GetSecurity().GetAuthz().GetEnabled(),
		RoscoMode:         h.GetFeatures().GetRoscoMode(),
	}
}

func getDeckNotificationsConfig(h *config.Hal) *config.Deck_Notifications {
	return &config.Deck_Notifications{
		Bearychat:    h.GetNotifications().GetBearychat(),
		Email:        h.GetNotifications().GetEmail(),
		GithubStatus: h.GetNotifications().GetGithubStatus(),
		GoogleChat:   h.GetNotifications().GetGooglechat(),
		Pubsub:       h.GetNotifications().GetPubsub(),
		Slack:        h.GetNotifications().GetSlack(),
		Sms:          h.GetNotifications().GetTwilio(),
	}
}

// TODO: find more elegant go-native way to reduce duplicated logic for:
// 1. Finding account config for primaryAccount
// 2. Finding first configured region for a primaryAccount
func getDeckProvidersConfig(h *config.Hal) *config.Deck_Providers {
	providers := &config.Deck_Providers{}

	if h.GetProviders().GetAppengine().GetEnabled() {
		providers.Appengine = &config.Deck_Providers_Appengine{
			Defaults: &config.Deck_Providers_Appengine_Defaults{
				Account: h.GetProviders().GetAppengine().GetPrimaryAccount(),
			},
		}
	}

	if h.GetProviders().GetAws().GetEnabled() {
		defaultAccountName := h.GetProviders().GetAws().GetPrimaryAccount()
		var defaultAccount *cloudprovider.AwsAccount
		for _, a := range h.GetProviders().GetAws().GetAccounts() {
			if a.GetName() == defaultAccountName {
				defaultAccount = a
			}
		}

		defaultRegion := ""
		if defaultAccount != nil && len(defaultAccount.GetRegions()) > 0 {
			defaultRegion = defaultAccount.GetRegions()[0].GetName()
		}

		providers.Aws = &config.Deck_Providers_Aws{
			Defaults: &config.Deck_Providers_Aws_Defaults{
				Account: defaultAccountName,
				Region:  defaultRegion,
			},
		}
	}

	if h.GetProviders().GetAzure().GetEnabled() {
		defaultAccountName := h.GetProviders().GetAzure().GetPrimaryAccount()
		var defaultAccount *cloudprovider.AzureAccount
		for _, a := range h.GetProviders().GetAzure().GetAccounts() {
			if a.GetName() == defaultAccountName {
				defaultAccount = a
			}
		}

		defaultRegion := ""
		if defaultAccount != nil && len(defaultAccount.GetRegions()) > 0 {
			defaultRegion = defaultAccount.GetRegions()[0]
		}

		providers.Azure = &config.Deck_Providers_Azure{
			Defaults: &config.Deck_Providers_Azure_Defaults{
				Account: defaultAccountName,
				Region:  defaultRegion,
			},
		}
	}

	if h.GetProviders().GetCloudfoundry().GetEnabled() {
		providers.Cloudfoundry = &config.Deck_Providers_Cloudfoundry{
			Defaults: &config.Deck_Providers_Cloudfoundry_Defaults{
				Account: h.GetProviders().GetCloudfoundry().GetPrimaryAccount(),
			},
		}
	}

	if h.GetProviders().GetDcos().GetEnabled() {
		providers.Dcos = &config.Deck_Providers_Dcos{
			Defaults: &config.Deck_Providers_Dcos_Defaults{
				Account: h.GetProviders().GetDcos().GetPrimaryAccount(),
			},
		}
	}

	if h.GetProviders().GetEcs().GetEnabled() {
		providers.Ecs = &config.Deck_Providers_Ecs{
			Defaults: &config.Deck_Providers_Ecs_Defaults{
				Account: h.GetProviders().GetEcs().GetPrimaryAccount(),
			},
		}
	}

	if h.GetProviders().GetGoogle().GetEnabled() {
		defaultAccountName := h.GetProviders().GetGoogle().GetPrimaryAccount()
		var defaultAccount *cloudprovider.GoogleComputeEngineAccount
		for _, a := range h.GetProviders().GetGoogle().GetAccounts() {
			if a.GetName() == defaultAccountName {
				defaultAccount = a
			}
		}

		defaultRegion := ""
		if defaultAccount != nil && len(defaultAccount.GetRegions()) > 0 {
			defaultRegion = defaultAccount.GetRegions()[0]
		}

		providers.Gce = &config.Deck_Providers_Gce{
			Defaults: &config.Deck_Providers_Gce_Defaults{
				Account: defaultAccountName,
				Region:  defaultRegion,
			},
		}
	}

	if h.GetProviders().GetHuaweicloud().GetEnabled() {
		defaultAccountName := h.GetProviders().GetHuaweicloud().GetPrimaryAccount()
		var defaultAccount *cloudprovider.HuaweiCloudAccount
		for _, a := range h.GetProviders().GetHuaweicloud().GetAccounts() {
			if a.GetName() == defaultAccountName {
				defaultAccount = a
			}
		}

		defaultRegion := ""
		if defaultAccount != nil && len(defaultAccount.GetRegions()) > 0 {
			defaultRegion = defaultAccount.GetRegions()[0]
		}

		providers.Huaweicloud = &config.Deck_Providers_HuaweiCloud{
			Defaults: &config.Deck_Providers_HuaweiCloud_Defaults{
				Account: defaultAccountName,
				Region:  defaultRegion,
			},
		}
	}

	return providers
}
