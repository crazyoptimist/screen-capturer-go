# Screen Capturer

This is a client-server application designed for monitoring computer screens on a local area network (LAN).

### Technical Outline

**Client:**

A client is the computer that monitors registered computers on a LAN.

- Server Registration: The administrator manually registers servers on the client by providing a unique name for each server.
- MDNS Discovery: The client uses Multicast DNS (MDNS) to discover the local IP address of registered servers. IP address scan is performed at a hard-coded frequency (for example, once per minute, for now).
- Screenshot Requests: The client periodically requests screenshots from all registered servers at a configurable frequency.
- Screenshot Storage: Received screenshots are stored in named subdirectories (corresponding to the server names) within a specified directory on the client.

**Server:**

A server is the computer that is being monitored.

- Registration: The server application will be given a registered name by the administrator.
- Screenshot Capture: The server captures its current screen (including multiple extended screens) using screen APIs or graphics drivers. The captured image is saved as a [file format, e.g., JPEG or PNG] file.
- HTTP Serving: The server serves the captured screenshot file via HTTP when requested by the client.

**Additional Notes:**

* The client uses a SQLite database to store persistent information about registered servers.
* The client and server communicate using the HTTP protocol over the network.

### Development

### TODO (Documentation)

- Development Guide
- Build Guide
- Cross-Platform Installation Guide

### Author

[crazyoptimist](https://crazyoptimist.net)
