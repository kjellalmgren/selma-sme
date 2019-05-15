# API Authorization Guide

This guide shows you how to get a user’s authorization to access private data through the API.

Permission
Requests to the Roaring API require authorization; that is, the user must have granted permission for an application to access the requested data. To prove that the user has granted permission, the request header sent by the application must include a valid access token.

As the first step towards authorization, you will need to register your application on the developer page and have that application subscribe to the api that you will be calling. That will give you a unique Consumer Key and Consumer Secret to use in the authorization header.

1. Your application requests access tokens

Using your unique Consumer Key and Consumer Secret you call the token service to retrieve an access token. The call is made towards the /token endpoint:

	POST https://api.roaring.io/token

The body of this POST request must contain the following parameters:

REQUEST BODY PARAMETER	VALUE
grant_type	Required. Shall be set to "client_credentials"
The header of this POST request must contain the following parameter:

HEADER PARAMETER	VALUE
Authorization	Required. Base 64 encoded string that contains the consumer key and consumer secret. The field must have the format: Authorization: Basic
2. The tokens are returned to the application

On success, the response from the Roaring Accounts service has the status code 200 OK in the response header, and the following JSON data in the response body:

	KEY	VALUE DESCRIPTION
	access_token
	string	An access token to be used in subsequent calls to the Roaring API.
	token_type 
	string	How the access_token may be used, always "Bearer".
	scope 
	string	A space-separated list of scopes which have been granted for this 	access_token
	expires_in 
	int	The time period (in seconds) for which the access token is valid.
	An example request and response to the token endpoint will look something like this:
	$ curl -H "Authorization: Basic XXX...zzz" -d grant_type=client_credentials https://api.roaring.io/token 
	{
  		"access_token": "asdfg...xzy", 
  		"token_type": "Bearer", 
  		"scope": "am_application_scope default", 
  		"expires_in": 3600
	}
3. Use the access token to access the Roaring API

The access token allows you to make requests to the Roaring API.

An example of how access token allows you to make requests to the Roaring API.

	$ curl -H "Authorization: Bearer XXXX...zzzzz" "https://api.roaring.io/person/1.0/	person?personalNumber=193604139208"
	{
  		"posts": [
    	{
      		"nationalRegistryChangeDate": "2011-03-15T00:00",
      		"personalNumber": "193604139208",
      		"hasHistory": true,
      		"secrecyChangeDate": "2010-02-02T00:00",
      		"secrecyMarked": false,
      		"details": [
        	{
          	"dateFrom": "2011-03-15T00:00",
          	"dateTo": "9999-12-31T00:00",
          	"firstName": "Carina",
          	"surName": "Efternamn1301",
          	"gender": "F",
          	"birthDate": "1936-04-13T00:00",
          	"deRegistrationDate": "2011-02-02",
          	"deRegistrationReason": "A"
        	}
      	],
      	"address": {
        "nationalRegistrationAddress": [
          {
            "dateFrom": "2015-12-18T00:00",
            "dateTo": "9999-12-31T00:00",
            "registrationDate": "2002-09-01T00:00",
            "careOf": "CO-NAMN",
            "deliveryAddress2": "Gatan177 2",
            "postalNumber": "17890",
            "city": "EKERÖ",
            "districtCode": "215002",
            "communeCode": "25",
            "countyCode": "01"
          }
        ]
      }
    }
  	]
	}
Authentication

All requests towards the API's require authentication. This is achieved by using the Consumer key and Consumer secret received from the application registered as subscriber on the developer site and calling the token service:

	Use the following example to generate an access token using the client_credentials grant type.
	curl request
	curl -k -d "grant_type=client_credentials" \
                    -H "Authorization: Basic Base64(consumer-key:consumer-secret)" \
                     https://api.roaring.io/token
	csharp request
	var client = new RestClient("https://api.roaring.io/token");
	var request = new RestRequest(Method.POST);
	request.AddHeader("cache-control", "no-cache");
	request.AddHeader("content-type", "application/x-www-form-urlencoded");
	request.AddHeader("authorization", "Basic Base64(consumer-key:consumer-secret)");
	request.AddParameter("application/x-www-form-urlencoded", 	"grant_type=client_credentials", ParameterType.RequestBody);
	IRestResponse response = client.Execute(request);                        
	java request
	OkHttpClient client = new OkHttpClient();
	MediaType mediaType = MediaType.parse("application/x-www-form-urlencoded");
	RequestBody body = RequestBody.create(mediaType, "grant_type=client_credentials");
	Request request = new Request.Builder()
  		.url("https://api.roaring.io/token")
  		.post(body)
  		.addHeader("authorization", "Basic Base64(consumer-key:consumer-secret)")
  		.addHeader("content-type", "application/x-www-form-urlencoded")
  		.addHeader("cache-control", "no-cache")
  		.build();

	Response response = client.newCall(request).execute();
	php request
	<?php

	$request = new HttpRequest();
	$request->setUrl('https://api.roaring.io/token');
	$request->setMethod(HTTP_METH_POST);

	$request->setHeaders(array(
  	'Cache-Control' => 'no-cache',
  	'Content-Type' => 'application/x-www-form-urlencoded',
  	'authorization' => 'Basic Base64(consumer-key:consumer-secret)'
	));

	$request->setContentType('application/x-www-form-urlencoded');
	$request->setPostFields(array(
  	'scope' => 'PRODUCTION',
  	'grant_type' => 'client_credentials'
	));

	try {
  		$response = $request->send();

  		echo $response->getBody();
	} catch (HttpException $ex) {
  		echo $ex;
	}
	                           
	Where Base64(consumer-key:consumer-secret) is a placeholder for the actual base64 	encoded version of the string "consumer-key" + "consumer-secret". This will return 	an OAuth token with a validity time according to the settings you have made, a 	response example:

	Example response
	{
   		"access_token": "d4c7d9a8-eb5f-34fb-b7b6-2acaad1cdd86",
   		"scope": "am_application_scope default",
   		"token_type": "Bearer",
   		"expires_in": 3600
	}
	
#Landshypotek test-bench

	consumer-key:secret-key HVD7kM1uBFFzSsVKDmpy7YjRfeIa:buXOYxxSujXJW0lAOaUaxqtpqyUa
	
	$ curl -k -d "grant_type=client_credentials" \
                    -H "Authorization: Basic Base64(HVD7kM1uBFFzSsVKDmpy7YjRfeIa:buXOYxxSujXJW0lAOaUaxqtpqyUa)" \
                     https://api.roaring.io/token
                     
    $ curl -k -d "grant_type=client_credentials" \
                    -H "Authorization: Basic SFZEN2tNMXVCRkZ6U3NWS0RtcHk3WWpSZmVJYTpidVhPWXh4U3VqWEpXMGxBT2FVYXhxdHBxeVVh" \
                     https://api.roaring.io/token
                     
                     
                     
                 