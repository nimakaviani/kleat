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
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/options"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// HalToFront50 generates the front50 config for the supplied config.Hal h.
func HalToFront50(h *config.Hal, opts options.GenerateOptions) *config.Front50 {
	return &config.Front50{
		Spinnaker: &config.Front50_Spinnaker{
			Gcs:      h.GetPersistentStorage().GetGcs(),
			Azs:      h.GetPersistentStorage().GetAzs(),
			Oracle:   h.GetPersistentStorage().GetOracle(),
			S3:       h.GetPersistentStorage().GetS3(),
			Redis:    h.GetPersistentStorage().GetRedis(),
			Delivery: getDelivery(h, opts),
		},
		StorageService: getStorageService(opts),
	}
}

func getDelivery(h *config.Hal, opts options.GenerateOptions) *config.Front50_Delivery {
	if !opts.EnableKeel {
		return nil
	}

	return &config.Front50_Delivery{
		Enabled: wrapperspb.Bool(false),
	}
}

func getStorageService(opts options.GenerateOptions) *config.Front50_StorageService {
	if !opts.EnableKeel {
		return nil
	}

	return &config.Front50_StorageService{
		DeliveryConfig: &config.Front50_StorageService_DeliveryConfig{
			RefreshMs: 60,
		},
	}
}
