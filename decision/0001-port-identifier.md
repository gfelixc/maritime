# 1. Port identifier persistence

Author: George Felix
Date: 2022-04-11

## Context

According to UNECE the UNLOC is a [five-character code system](https://uncefact.unece.org/pages/viewpage.action?pageId=17830748)
that provides a coded representation for the names of ports, airports, inland clearance depots, inland freight terminals
and other transport related locations which are used for the movement of goods for trade. ([source](https://unece.org/trade/uncefact/unlocode)).
So UNLOC must be the port identifier. The problem using UNLOC as the port identifier comes during the creation of a new port, 
since the identifier is delegated to the clients, this could lead into tricky situation where the creation will create or 
update the port in case of it exists.

## Alternatives

### A. Surrogate Id

Using a surrogate id (created either by the application or database) + index UNLOC as unique index will help to prevent
undesired overwrites in the creation meanwhile grants the uniqueness of UNLOC.

✅Prevent undesired overwrites by creation
❌Limits the Database choose
❌Performance overhead derived by unique index
❌Pollute Port model with surrogate id

### B. UNLOC as port identifier + separated Create/Update-like actions

In order to prevent undesired updates by the Create action but guarantee the uniqueness of UNLOC, would require locking
the PortRepository, search the given UNLOC in the repository then either create the Port or fail with a "UNLOC already exists"
and finally unlock the PortRepository.

✅Prevent undesired overwrites by creation
❌Locking PortRepository impact other write ops
❌Limits the Database choose

### C. UNLOC as port identifier + CreateUpdate action

Using an explicit naming as CreateUpdate will be an accurate reflect of the action to perform.

✅Port model accuracy with business
❌Evolving the application beyond CRUD will require to replace this solution

## Decision

In the context of this exercise, the decision is to use the alternative C as it fit in the exercise requirements
