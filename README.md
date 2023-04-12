# Before start

```shell
make init
go install github.com/vektra/mockery/v2@v2.20.0
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.52.2
GOBIN=/usr/local/bin go install github.com/bufbuild/buf/cmd/buf@v1.17.0
```

# Assumptions

- Application will evolve beyond a simple CRUD
- UNLOC refers to the United Nations Code for Trade and Transport Locations (UN/LOCODE). Which is a [five-character code system](https://uncefact.unece.org/pages/viewpage.action?pageId=17830748) that provides a coded representation for the names of ports, airports, inland clearance depots, inland freight terminals and other transport related locations which are used for the movement of goods for trade. ([source](https://unece.org/trade/uncefact/unlocode)) 
- Timezones must exist in the IANA Database
- UNLOCS list must contain the UNLOC used as field name in data json