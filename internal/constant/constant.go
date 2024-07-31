package constant

const SERVER_WEB_PORT = 8949
const CLIENT_WEB_PORT = 9999
const FIND_PC_TIMEOUT_IN_SECONDS = 2

// Assuming we can't find out any computers in the LAN,
// and the number of registered computer is 20,
// then you might want to set the scan interval to
// FIND_PC_TIMEOUT_IN_SECONDS * 20
const SCAN_NETWORK_INTERVAL_IN_SECONDS = 60
