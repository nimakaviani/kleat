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

package convert

import (
	"github.com/spinnaker/kleat/api/client"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/options"
	"github.com/spinnaker/kleat/pkg/version"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// HalToEcho generates the echo config for the supplied config.Hal h.
func HalToEcho(h *config.Hal, opts options.GenerateOptions) *config.Echo {
	return &config.Echo{
		Slack:        h.GetNotifications().GetSlack(),
		Twilio:       h.GetNotifications().GetTwilio(),
		GithubStatus: h.GetNotifications().GetGithubStatus(),
		Artifacts:    h.GetArtifacts(),
		Pubsub:       h.GetPubsub(),
		Gcb:          h.GetCi().GetGcb(),
		Stats:        getEchoStats(h),
		Scheduler:    getEchoScheduler(h),
		Services:     getEchoServices(opts),
	}
}

func getEchoServices(opts options.GenerateOptions) *config.Echo_Services {
	if !opts.EnableKeel {
		return nil
	}

	return &config.Echo_Services{
		Keel: &config.ServiceSettings{
			Enabled: wrapperspb.Bool(opts.EnableKeel),
		},
	}
}

func getEchoScheduler(h *config.Hal) *config.Echo_Scheduler {
	if h.GetTimezone() == "" {
		return nil
	}
	return &config.Echo_Scheduler{
		Cron: &config.Echo_Scheduler_Cron{
			Timezone: h.GetTimezone(),
		},
	}
}

func getEchoStats(h *config.Hal) *config.Echo_Stats {
	return &config.Echo_Stats{
		Enabled:          h.GetStats().GetEnabled(),
		InstanceId:       h.GetStats().GetInstanceId(),
		Endpoint:         h.GetStats().GetEndpoint(),
		DeploymentMethod: getDeploymentMethod(),
		SpinnakerVersion: h.GetVersion(),
	}
}

func getDeploymentMethod() *client.DeploymentMethod {
	return &client.DeploymentMethod{Type: "kleat", Version: version.Version}
}
