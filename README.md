# Maritime

## Before start

Make sure you have all the dependencies installed in the expected version by running the command `make init`

## Development commands

- Run tests: `make tests`
- Run lints: `make lints`
- Run server: `make run`
- Example query: `grpcurl -plaintext -d '{"filename":"ports.json"}' localhost:8080 port.v1.PortDomainService/CreateOrUpdateFromPortsDataFile`
- Generate mocks: `go generate ./...`
- Generate proto stubs: `buf generate port`

## Assumptions

- Application will evolve beyond a simple CRUD
- UNLOC refers to the United Nations Code for Trade and Transport Locations (UN/LOCODE). Which is a [five-character code system](https://uncefact.unece.org/pages/viewpage.action?pageId=17830748) that provides a coded representation for the names of ports, airports, inland clearance depots, inland freight terminals and other transport related locations which are used for the movement of goods for trade. ([source](https://unece.org/trade/uncefact/unlocode)) 
- Timezones must exist in the IANA Database
- UNLOCS list must contain the UNLOC used as field name in data json