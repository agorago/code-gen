Feature: Test the hello proxy and service

Scenario: Invoking Uppercase with valid token
When I invoke Uppercase with "hello,world" and a secure token with value "passpass"
Then I must get back an upper case return value of "HELLO,WORLD"

Scenario: Invoking Uppercase with invalid token
When I invoke Uppercase with "hello,world" and a secure token with value "fail"
Then I must get back an error with http error code 403

Scenario: Invoking Uppercase with no token
When I invoke Uppercase with "hello,world" without secure token
Then I must get back an error with http error code 403

Scenario: Invoking Count
When I invoke Count with "hello,world" 
Then I must get back a count return value of 11

Scenario: Invoking Count with string "Count" which is treated in a special way by the proxy interceptor
When I invoke Count with "Count"
Then I must get back a count return value of 8

Scenario: Invoking AddNumbers
When I invoke AddNumbers with 78 and 85 
Then I must get back an addnumbers return value of 163
