package convert

import (
	"fmt"

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/options"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// HalToKeel generates the Spinnaker keel config for the supplied config.Hal h
func HalToKeel(h *config.Hal, opts options.GenerateOptions) *config.Keel {
	if !opts.EnableKeel {
		return nil
	}

	var urlFormat = "http://%s.spinnaker:%d"
	services := map[string]struct {
		name string
		port int32
	}{
		"echo":        {name: "echo-worker", port: 8089},
		"front50":     {name: "front50", port: 8080},
		"igor":        {name: "igor", port: 8088},
		"clouddriver": {name: "clouddriver-rw", port: 7002},
		"orca":        {name: "orca", port: 8083},
	}

	return &config.Keel{
		Keel:         h.GetKeel(),
		Eureka:       getEureka(),
		Services:     getServices(h),
		Sql:          h.GetPersistentStorage().GetSql(),
		OkHttpClient: getHttpClient(),
		Server:       getServer(),
		Retrofit2:    getRetrofit(),
		Echo: &config.ServiceSettings{
			Enabled: wrapperspb.Bool(true),
			BaseUrl: fmt.Sprintf(urlFormat, services["echo"].name, services["echo"].port),
		},
		Igor: &config.ServiceSettings{
			Enabled: wrapperspb.Bool(true),
			BaseUrl: fmt.Sprintf(urlFormat, services["igor"].name, services["igor"].port),
		},
		Clouddriver: &config.ServiceSettings{
			Enabled: wrapperspb.Bool(true),
			BaseUrl: fmt.Sprintf(urlFormat, services["clouddriver"].name, services["clouddriver"].port),
		},
		Orca: &config.ServiceSettings{
			Enabled: wrapperspb.Bool(true),
			BaseUrl: fmt.Sprintf(urlFormat, services["orca"].name, services["orca"].port),
		},
		Front50: &config.ServiceSettings{
			Enabled: wrapperspb.Bool(true),
			BaseUrl: fmt.Sprintf(urlFormat, services["front50"].name, services["front50"].port),
		},
	}
}

func getEureka() *config.Keel_Eureka {
	return &config.Keel_Eureka{
		Enabled: wrapperspb.Bool(false),
	}
}

func getServices(h *config.Hal) *config.Keel_Services {
	return &config.Keel_Services{
		Fiat: &config.ServiceSettings{
			Enabled: wrapperspb.Bool(false),
		},
	}
}

func getHttpClient() *config.Keel_OkHttpClient {
	return &config.Keel_OkHttpClient{
		PropagateSpinnakerHeaders: wrapperspb.Bool(true),
		ConnectTimeoutMs:          30000,
		ReadTimeoutMs:             59000,
	}
}

func getServer() *config.Keel_Server {
	return &config.Keel_Server{
		Port: 7002,
	}
}

func getRetrofit() *config.Keel_Retrofit2 {
	return &config.Keel_Retrofit2{
		LogLevel: "BASIC",
	}
}
