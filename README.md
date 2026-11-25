## Go module for MaxMind GeoIP web service access
[![GoDoc](https://godoc.org/github.com/tdrn-org/go-maxmind-geoip?status.svg)](https://godoc.org/github.com/tdrn-org/go-maxmind-geoip)
[![Build](https://github.com/tdrn-org/go-maxmind-geoip/actions/workflows/build.yml/badge.svg)](https://github.com/tdrn-org/go-maxmind-geoip/actions/workflows/build.yml)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=tdrn-org_go-maxmind-geoip&metric=coverage)](https://sonarcloud.io/summary/new_code?id=tdrn-org_go-maxmind-geoip)

### MaxMind GeoIP web service
The [MaxMind GeoIP web service]https://www.maxmind.com/en/geoip-api-web-services) provides a [REST API](https://github.com/maxmind/openapi) to query GeoIP information. This library provides the Go wrapper for this API.

## License & Attribution
This project is subject to the the MIT License (see the [LICENSE](./LICENSE) file for details).

#### Third-Party Notices
This project is based on the official OpenAPI specification from **[MaxMind](https://github.com/maxmind/openapi)**.

* **Component:** MaxMind web service OpenAPI Specification (`api/geoip.yaml`)
* **Copyright:** See https://github.com/maxmind/openapi
* **License:** [Apache License, Version 2.0](https://raw.githubusercontent.com/maxmind/openapi/refs/heads/main/LICENSE-APACHE) or [MIT License](https://raw.githubusercontent.com/maxmind/openapi/refs/heads/main/LICENSE-MIT)

The OpenAPI specification file remains under the terms of its license (see above). This module's Go code is subject to the MIT License.
