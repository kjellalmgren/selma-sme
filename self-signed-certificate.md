# self signed certificate

## openssl installed
	
	$ which openssl
	Ska svara med en path "/usr/bin/openssl" för OSX, om inte openssl är inte installerat
	
## Create certificate

	$ openssl genrsa -des3 -passout pass:x -out server.pass.key 2048
	...
	$ openssl rsa -passin pass:x -in server.pass.key -out server.key
	# writing RSA key
	$ rm server.pass.key
	$ openssl req -new -key server.key -out server.csr
	...
	Country Name (2 letter code) [AU]:US
	State or Province Name (full name) [Some-State]:California
	...
	A challenge password []:
	...

## Generate SSL certificate

The self-signed SSL certificate is generated from the server.key private key and server.csr files.

	$ openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -out server.crt