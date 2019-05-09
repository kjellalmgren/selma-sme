# self signed certificate

## openssl installed
	
	$ which openssl
	Ska svara med en path "/usr/bin/openssl" för OSX, om inte openssl är inte installerat
	
## Create certificate

	# use 127.0.0.1:8443 when register
	$ openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
	...

## Generate SSL certificate

The self-signed SSL certificate is generated from the server.key private key and server.csr files.

	$ openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -out server.crt

## Node js example

	$ openssl genrsa -des3 -out server.key 1024
	$ openssl req -new -key server.key -out server.csr
	$ openssl x509 -req -days 3650 -in server.csr -signkey server.key -out server.crt
	$ cp server.key server.key.copy
	$ openssl rsa -in server.key.copy -out server.key
	$ rm server.key.copy
