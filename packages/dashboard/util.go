package dashboard

import (
	"fmt"
	"strings"
	"time"

	"github.com/iotaledger/wasp/packages/parameters"
)

func args(args ...interface{}) []interface{} {
	return args
}

func formatTimestamp(ts interface{}) string {
	t, ok := ts.(time.Time)
	if !ok {
		t = time.Unix(0, ts.(int64))
	}
	return t.UTC().Format(time.RFC3339)
}

func exploreAddressUrl(baseUrl string) func(address fmt.Stringer) string {
	return func(address fmt.Stringer) string {
		return baseUrl + "/" + address.String()
	}
}

func exploreAddressBaseUrl() string {
	baseUrl := parameters.GetString(parameters.DashboardExploreAddressUrl)
	if baseUrl != "" {
		return baseUrl
	}
	return exploreAddressUrlFromGoshimmerUri(parameters.GetString(parameters.NodeAddress))
}

func exploreAddressUrlFromGoshimmerUri(uri string) string {
	url := strings.Split(uri, ":")[0] + ":8081/explorer/address"
	if !strings.HasPrefix(url, "http") {
		return "http://" + url
	}
	return url
}
