# dwaf

dwaf (dredteam WAF) is a modern HTTP Web Application Firewall that acts as a reverse proxy; dwaf is written in Go.

❗**dwaf is not ready for use yet and is still in active development phase.**

## Features (planned)

* Updates its signatures in background
* Focuses on blocking requests and not responses
* Uses centralized KV storage (consul) for everything
* Ships in docker

## Limitations by design

* does not support TLS
* does not support acting as an edge HTTP server
* does not support explicit deny (default allow)

## Roadmap

TBA

## Contributing

TBA

## License

    dwaf (dredteam WAF) is a modern HTTP Web Application Firewall that acts as a reverse proxy
    Copyright (C) 2020  DREDTEAM S.R.L.

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
