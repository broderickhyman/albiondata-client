# Albion Market - Client
Distributed client for the [Albion Market](https://albion-market.com/)
website.

This client monitors local network traffic, identifies UDP packets
that contain market data for Albion Online, and ships the raw JSON
from the market packets off to an ingest server.

The client is not meant to parse the packets any more than what is
required to pull the JSON out of them. Processing the JSON is done
on the ingest server so that if an update is required to parse the
JSON it does not require updating the distributed client.

# TODO
- [ ] When installed auto start on boot
- [ ] Ignore market packets for the test server
- [ ] Identify where the character is located and set that location
- [ ] Deal with multiple running clients??
- [ ] Release clients for Mac and Linux

# Related Projects
- [Albion Market - Backend](https://github.com/Regner/albionmarket-backend/)
- [Albion Market - Frontend](https://github.com/Regner/albionmarket-frontend/)

# License
This project, and all contributed code, are licensed under the MIT
License. A copy of the MIT License may be found in the repository.
