(window.webpackJsonp=window.webpackJsonp||[]).push([[12],{368:function(e,t,_){"use strict";_.r(t);var v=_(42),d=Object(v.a)({},(function(){var e=this,t=e.$createElement,_=e._self._c||t;return _("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[_("h1",{attrs:{id:"coap-gateway"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#coap-gateway"}},[e._v("#")]),e._v(" COAP gateway")]),e._v(" "),_("h2",{attrs:{id:"description"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#description"}},[e._v("#")]),e._v(" Description")]),e._v(" "),_("p",[e._v("OCF Servers / Clients communicate over TCP / UDP using the CoAP application protocol. Communication within the OCF Native Cloud shouldn't be restricted to the CoAP protocol, implementation should allow the use of whatever protocol might be introduced in the future. That's why the gateway is the access point for CoAP over TCP, and further communication is OCF Native Cloud specific.")]),e._v(" "),_("p",[e._v("TCP connection to the OCF Native Cloud is by its nature stateful. The OCF CoAP Gateway is therefore also stateful, keeping open connections to the OCF Servers / Clients.  The goal of the Gateway is to translate between the OCF Servers / Clients (CoAP) and the protocol of the OCF Native Cloud and communicate in an asynchronous way.")]),e._v(" "),_("h3",{attrs:{id:"validation"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#validation"}},[e._v("#")]),e._v(" Validation")]),e._v(" "),_("ul",[_("li",[e._v("OCF CoAP Gateway can accept requests from the OCF Client / Server only after a successful sign-in")]),e._v(" "),_("li",[e._v("OCF CoAP Gateway can forward requests to the OCF Client / Server only after successful sign-in")]),e._v(" "),_("li",[e._v("If sign-in was not issued within the configured amount of time or sign-in request failed, OCF Native Cloud will forcibly close the TCP connection")]),e._v(" "),_("li",[e._v("OCF CoAP Gateway sends command to update device core resource with its status.\n"),_("ul",[_("li",[e._v("Online when the device was successfully signed-in and communication lock released")]),e._v(" "),_("li",[e._v("Offline when the device was disconnected or signed-out")])])]),e._v(" "),_("li",[e._v("Access Token from a successful sign-in must be locally persisted in the OCF CoAP Gateway and linked with an opened TCP channel")]),e._v(" "),_("li",[e._v("Access Token linked with the opened TCP channel has to be included in each command issued to other OCF Native Cloud components")]),e._v(" "),_("li",[e._v("OCF CoAP Gateway processes only those commands, which are designated for a device which the Gateway has an opened TCP channel to")]),e._v(" "),_("li",[e._v("OCF CoAP Gateway is observing each resource published to the resource directory and publishes an event for every change")]),e._v(" "),_("li",[e._v("OCF CoAP Gateway retrieves each published resource and updates Resources")]),e._v(" "),_("li",[e._v("OCF CoAP Gateway has to expose the coap ping-pong + retry count configuration, which can be configured during the deployment")]),e._v(" "),_("li",[e._v("OCF CoAP Gateway has to ping the device in the configured time, if pong is not received after the configured number of retries, then the connection with the device is closed and device is set as offline")]),e._v(" "),_("li",[e._v("OCF CoAP Gateway processes events from Resources, by issuing a proper CoAP request to the device and raising an event with the response")]),e._v(" "),_("li",[e._v("OCF CoAP Gateway has to process a waiting request within the configured time, or set the device as offline")])]),e._v(" "),_("h2",{attrs:{id:"docker-image"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#docker-image"}},[e._v("#")]),e._v(" Docker Image")]),e._v(" "),_("div",{staticClass:"language-bash extra-class"},[_("pre",{pre:!0,attrs:{class:"language-bash"}},[_("code",[e._v("docker pull plgd/coap-gateway:vnext\n")])])]),_("h2",{attrs:{id:"api"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#api"}},[e._v("#")]),e._v(" API")]),e._v(" "),_("p",[e._v("Follow "),_("a",{attrs:{href:"https://openconnectivity.org/specs/OCF_Device_To_Cloud_Services_Specification_v2.2.0.pdf",target:"_blank",rel:"noopener noreferrer"}},[e._v("OCF Device To Cloud Services Specification"),_("OutboundLink")],1)]),e._v(" "),_("h3",{attrs:{id:"commands"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#commands"}},[e._v("#")]),e._v(" Commands")]),e._v(" "),_("ul",[_("li",[e._v("POST /oic/sec/account - sign up the device with authorization code")]),e._v(" "),_("li",[e._v("DELETE /oic/sec/account - sign off the device with access token")]),e._v(" "),_("li",[e._v("POST /oic/sec/tokenrefresh - refresh access token with refresh token")]),e._v(" "),_("li",[e._v("POST /oic/sec/session - sign in the device with access token and with login true")]),e._v(" "),_("li",[e._v("POST /oic/sec/session - sign out the device with access token and with login false")]),e._v(" "),_("li",[e._v("POST /oic/rd - publish resources from the signed device")]),e._v(" "),_("li",[e._v("DELETE /oic/rd - unpublish resources from the signed device")]),e._v(" "),_("li",[e._v("GET /oic/res - discover all cloud devices resources from the signed device")]),e._v(" "),_("li",[e._v("GET /oic/route/{deviceID}/{href} - get/observe resource of the cloud device from signed device")]),e._v(" "),_("li",[e._v("POST /oic/route/{deviceID}/{href} - update resource of the cloud device from signed device")])]),e._v(" "),_("h2",{attrs:{id:"configuration"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#configuration"}},[e._v("#")]),e._v(" Configuration")]),e._v(" "),_("table",[_("thead",[_("tr",[_("th",[e._v("Option")]),e._v(" "),_("th",[e._v("ENV variable")]),e._v(" "),_("th",[e._v("Type")]),e._v(" "),_("th",[e._v("Description")]),e._v(" "),_("th",[e._v("Default")])])]),e._v(" "),_("tbody",[_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("ADDRESS")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("listen address")])]),e._v(" "),_("td",[_("code",[e._v('"0.0.0.0:5684"')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("EXTERNAL_PORT")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("used to fill discovery hrefs")])]),e._v(" "),_("td",[_("code",[e._v('"0.0.0.0:5684"')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("FQDN")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("used to fill discovery")])]),e._v(" "),_("td",[_("code",[e._v('"coapgw.ocf.cloud"')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("AUTH_SERVER_ADDRESS")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("authoriztion server address")])]),e._v(" "),_("td",[_("code",[e._v('"127.0.0.1:9100"')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("RESOURCE_AGGREGATE_ADDRESS")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("resource aggregate address")])]),e._v(" "),_("td",[_("code",[e._v('"127.0.0.1:9100"')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("RESOURCE_DIRECTORY_ADDRESS")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("resource directory address")])]),e._v(" "),_("td",[_("code",[e._v('"127.0.0.1:9100"')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("REQUEST_TIMEOUT")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("wait for update/retrieve resource")])]),e._v(" "),_("td",[_("code",[e._v("10s")])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("KEEPALIVE_ENABLE")])]),e._v(" "),_("td",[e._v("bool")]),e._v(" "),_("td",[_("code",[e._v("check devices connection")])]),e._v(" "),_("td",[e._v("true")])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("KEEPALIVE_TIMEOUT_CONNECTION")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("close inactive connection after limit")])]),e._v(" "),_("td",[_("code",[e._v('"20s"')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DISABLE_BLOCKWISE_TRANSFER")])]),e._v(" "),_("td",[e._v("bool")]),e._v(" "),_("td",[_("code",[e._v("disable blockwise transfer")])]),e._v(" "),_("td",[_("code",[e._v("false")])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("BLOCKWISE_TRANSFER_SZX")])]),e._v(" "),_("td",[e._v("int")]),e._v(" "),_("td",[_("code",[e._v("size of blockwise transfer block")])]),e._v(" "),_("td",[_("code",[e._v("1024")])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DISABLE_TCP_SIGNAL_MESSAGE_CSM")])]),e._v(" "),_("td",[e._v("bool")]),e._v(" "),_("td",[_("code",[e._v("disable send CSM when connection was established")])]),e._v(" "),_("td",[_("code",[e._v("false")])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DISABLE_PEER_TCP_SIGNAL_MESSAGE_CSMS")])]),e._v(" "),_("td",[e._v("bool")]),e._v(" "),_("td",[_("code",[e._v("disable process peer CSM")])]),e._v(" "),_("td",[_("code",[e._v("false")])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("ERROR_IN_RESPONSE")])]),e._v(" "),_("td",[e._v("bool")]),e._v(" "),_("td",[_("code",[e._v("send text error message in response")])]),e._v(" "),_("td",[_("code",[e._v("true")])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("SERVICE_OAUTH_ENDPOINT_TOKEN_URL")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("url to get service access token via OAUTH client credential flow")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("SERVICE_OAUTH_CLIENT_ID")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("client id for authentication to get access token")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("SERVICE_OAUTH_CLIENT_SECRET")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("secrest for authentication to get access token")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("SERVICE_OAUTH_AUDIENCE")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("refer to the resource servers that should accept the token")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("HEARTBEAT")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("defines check of live service")])]),e._v(" "),_("td",[_("code",[e._v('"4s"')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("MAX_MESSAGE_SIZE")])]),e._v(" "),_("td",[e._v("int")]),e._v(" "),_("td",[_("code",[e._v("defines max message size which can be send/receive via coap")])]),e._v(" "),_("td",[_("code",[e._v("262144")])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DIAL_TYPE")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("defines how to obtain dial TLS certificates - options: acme|file")])]),e._v(" "),_("td",[_("code",[e._v('"acme"')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DIAL_ACME_CA_POOL")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("path to pem file of CAs")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DIAL_ACME_DIRECTORY_URL")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("url of acme directory")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DIAL_ACME_DOMAINS")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("list of domains for which will be in certificate provided from acme")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DIAL_ACME_REGISTRATION_EMAIL")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("registration email for acme")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DIAL_ACME_TICK_FREQUENCY")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("interval of validate certificate")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DIAL_ACME_USE_SYSTEM_CERTIFICATION_POOL")])]),e._v(" "),_("td",[e._v("bool")]),e._v(" "),_("td",[_("code",[e._v("load CAs from system")])]),e._v(" "),_("td",[_("code",[e._v("false")])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DIAL_FILE_CA_POOL")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("path to pem file of CAs")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DIAL_FILE_CERT_KEY_NAME")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("name of pem certificate key file")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DIAL_FILE_CERT_DIR_PATH")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("path to directory which contains DIAL_FILE_CERT_KEY_NAME and DIAL_FILE_CERT_NAME")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DIAL_FILE_CERT_NAME")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("name of pem certificate file")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("DIAL_FILE_USE_SYSTEM_CERTIFICATION_POOL")])]),e._v(" "),_("td",[e._v("bool")]),e._v(" "),_("td",[_("code",[e._v("load CAs from system")])]),e._v(" "),_("td",[_("code",[e._v("false")])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_TYPE")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("defines how to obtain listen TLS certificates - options: acme|file")])]),e._v(" "),_("td",[_("code",[e._v('"acme"')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_ACME_CA_POOL")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("path to pem file of CAs")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_ACME_DIRECTORY_URL")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("url of acme directory")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_ACME_DOMAINS")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("list of domains for which will be in certificate provided from acme")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_ACME_REGISTRATION_EMAIL")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("registration email for acme")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_ACME_TICK_FREQUENCY")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("interval of validate certificate")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_ACME_USE_SYSTEM_CERTIFICATION_POOL")])]),e._v(" "),_("td",[e._v("bool")]),e._v(" "),_("td",[_("code",[e._v("load CAs from system")])]),e._v(" "),_("td",[_("code",[e._v("false")])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_ACME_DEVICE_ID")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("deviceID for OCF Identity Certificate")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_FILE_CA_POOL")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("path to pem file of CAs")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_FILE_CERT_KEY_NAME")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("name of pem certificate key file")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_FILE_CERT_DIR_PATH")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("path to directory which contains LISTEN_FILE_CERT_KEY_NAME and LISTEN_FILE_CERT_NAME")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_FILE_CERT_NAME")])]),e._v(" "),_("td",[e._v("string")]),e._v(" "),_("td",[_("code",[e._v("name of pem certificate file")])]),e._v(" "),_("td",[_("code",[e._v('""')])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_FILE_USE_SYSTEM_CERTIFICATION_POOL")])]),e._v(" "),_("td",[e._v("bool")]),e._v(" "),_("td",[_("code",[e._v("load CAs from system")])]),e._v(" "),_("td",[_("code",[e._v("false")])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LISTEN_WITHOUT_TLS")])]),e._v(" "),_("td",[e._v("bool")]),e._v(" "),_("td",[_("code",[e._v("listen without TLS")])]),e._v(" "),_("td",[_("code",[e._v("false")])])]),e._v(" "),_("tr",[_("td",[_("code",[e._v("-")])]),e._v(" "),_("td",[_("code",[e._v("LOG_ENABLE_DEBUG")])]),e._v(" "),_("td",[e._v("bool")]),e._v(" "),_("td",[_("code",[e._v("debug logging")])]),e._v(" "),_("td",[_("code",[e._v("false")])])])])])])}),[],!1,null,null,null);t.default=d.exports}}]);