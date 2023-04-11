# Before start

- `go install github.com/vektra/mockery/v2@v2.20.0`

# Assumptions

- Application will evolve beyond a simple CRUD
- UNLOC refers to the United Nations Code for Trade and Transport Locations (UN/LOCODE). Which is a [five-character code system](https://uncefact.unece.org/pages/viewpage.action?pageId=17830748) that provides a coded representation for the names of ports, airports, inland clearance depots, inland freight terminals and other transport related locations which are used for the movement of goods for trade. ([source](https://unece.org/trade/uncefact/unlocode)) 
- Timezones must exist in the IANA Database
- UNLOCS list must contain the UNLOC used as field name in data json