# Maritime

## Before start

Make sure you have all the dependencies installed in the expected version by running the command `make init`


## Assumptions

- Application will evolve beyond a simple CRUD
- UNLOC refers to the United Nations Code for Trade and Transport Locations (UN/LOCODE). Which is a [five-character code system](https://uncefact.unece.org/pages/viewpage.action?pageId=17830748) that provides a coded representation for the names of ports, airports, inland clearance depots, inland freight terminals and other transport related locations which are used for the movement of goods for trade. ([source](https://unece.org/trade/uncefact/unlocode))
- Timezones must exist in the IANA Database
- UNLOCS list must contain the UNLOC used as field name in data json

## Development commands

- Run tests: `make tests`
- Run lints: `make lints`
- Generate mocks: `make generate-mocks`
- Generate proto stubs: `make generate-protos`
- Docker build: `make docker-build`

## Running exercise

- Run server locally: `make run` or dockerized: `make run-docker`
- Example query: `grpcurl -plaintext -d '{"filename":"data/ports.json"}' localhost:8080 port.v1.PortDomainService/CreateOrUpdateFromPortsDataFile`

Execution should return an error as the assumption is that timezones are based on IANA timezone database and America/Argentina is not fully accurate to IANA.

```json
{
  "portsProcessed": 1632,
  "portsSucceeded": 1631,
  "portsFailed": 1,
  "errors": [
    "[InvalidTimezone] - timezone provided doesn't exists in IANA timezone database | {\"UNLOC\":\"ARRIC\",\"Name\":\"Rio Cullen\",\"City\":\"Rio Cullen\",\"Country\":\"Argentina\",\"Alias\":[],\"Regions\":[],\"CoordinatesLongitude\":-68.3523021,\"CoordinatesLatitude\":-52.8955609,\"Province\":\"Zaire\",\"Timezone\":\"America/Argentina\",\"UNLOCS\":[\"ARRIC\"],\"Code\":\"35700\"}"
  ]
}

```

## Running tests

Unit tests: 
- `make tests`

Integration tests:
- `docker-compose up`
- `make integration-tests`
