[![CircleCI](https://circleci.com/gh/broderickhyman/albiondata-client/tree/master.svg?style=svg)](https://circleci.com/gh/broderickhyman/albiondata-client/tree/master) [![Go Report Card](https://goreportcard.com/badge/github.com/broderickhyman/albiondata-client)](https://goreportcard.com/report/github.com/broderickhyman/albiondata-client)

# Albion Data - Client
Distributed client for the [Albion Online Data](https://www.albion-online-data.com/)
project.

A quick note on the legality of this application and if it
violates the Terms and Conditions for Albion Online. Here is
the response from SBI when asked if we are allowed to do
monitor network packets relating to Albion Online:
> Our position is quite simple. As long as you just look and
analyze we are ok with it. The moment you modify or manipulate
something or somehow interfere with our services we will react
(e.g. perma-ban, take legal action, whatever).

~ MadDave - Technical Lead for Albion Online

Source: https://forum.albiononline.com/index.php/Thread/51604-Is-it-allowed-to-scan-your-internet-trafic-and-pick-up-logs/?postID=512670#post512670

This client monitors local network traffic, identifies UDP packets
that contain relevant data for Albion Online, and ships the information
off to a central NATS server that anyone can subscribe to.

# Contributions
Many thanks to the original developers:
- [Regner](https://github.com/Regner)
- [pcdummy](https://github.com/pcdummy)
- [Ultraporing](https://github.com/Ultraporing)

# Downloads
Downloads can be found here: https://github.com/broderickhyman/albiondata-client/releases

# Related Projects
- [albiondata-deduper](https://github.com/BroderickHyman/albiondata-deduper)
- [albiondata-sql](https://github.com/BroderickHyman/albiondata-sql)
- [albiondata-api-dotNet](https://github.com/broderickhyman/albiondata-api-dotNet)

# Contact Us
The best way to get in touch with us is on the Albion Online Fansites Discord server in either the #proj-albiondata or the #developers channel. My username is MyPickle#2527. A permanent invite link can be found here: [https://discord.gg/8DDSjTs](https://discord.gg/8DDSjTs)

# Developer Setup
### Mac/Linux Setup
- Install [Dep](https://github.com/golang/dep)
  - Any OS: `go get -u github.com/golang/dep/cmd/dep`
  - Mac with Homebrew: `brew install dep`
- Install dependencies using `dep ensure`

### Windows Setup
[Windows Setup Guide](https://github.com/broderickhyman/albiondata-client/wiki/Building-in-Windows)

# License
This project, and all contributed code, are licensed under the MIT
License. A copy of the MIT License may be found in the repository.
