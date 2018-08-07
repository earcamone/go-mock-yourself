
# Go Mock Yourself

## Summary

Go Mock Yourself is an attempt to build a set of flexible, robust and easy to use packages that will allow you to change
different Go native packages behaviour at run-time without having to change a single line of code (either from Go's
native packages or your production code). 

## Go Mock Yourself HTTP Requests

Go Mock Yourself HTTP package will allow you (hopefully) to easily mock any complex HTTP communication flow due to its
flexible and robust schemes without having to change a single line of code of your application. Go Mock Yourself HTTP
package replaces at runtime Go's native http package functions in order to make it behave like you wish, taking the 
overhead of having to develop complex test code to mimik your production environment. 

Matching incoming HTTP requests to serve your Mock HTTP responses can be either accomplished using regular expressions 
on every aspect of an incoming HTTP request (method, url, headers, body and more) or you can even register what we call 
"Dynamic Matching Functions" that will allow you at run-time to determine if an incoming request should match or not
a specific mock (you could for example connect to a database and based on its information mock or not the request).

Mock HTTP Responses can also be built dynamically at run-time or you can even make any HTTP call (for example http.Post)
just fail with an also dynamically generated (or not) Go error.

Additionally, Go Mock Yourself HTTP package allows you to do lots of other cool things such as (among many others):

 * Pausing/Playing Go Mock Yourself HTTP Mocking Scheme at run-time
 
 * Forcing HTTP requests to last a specific duration (no matter if you decide to mock the response or allow it to 
   reach the target server)
   
 * Just plug it to your application and log specific communication flows in a nice human readable format to debug
   complex scenarios or even build mocks for a different mocking solution ;)
   
 * Conditionally (and dynamically) serve a Mock Response or allow the request to be sent to the remote server
 
For a detailed explanation kindly visit [Go Mock Yourself HTTP package](./http).

## Go Mock Yourself SQL Statements

Coming soon..

## DISCLAIMER

If you take the time to inspect all Go Mock Yourself packages you will notice many places where Go's military syntax
enforcement is not respected. Don't get me wrong, i think Go is a wonderful language that is here to stay and will
probably replace in the near future massive languages such as Python.. BUT, on the other hand, and this is something
i notice in big companies too, people like us, that tend to write open source stuff and started coding when we were
kids, do what we do because we love doing it and its fun! and when someone comes and tells you how you must write your
code, even though i understand the argument that enforces this procedure, it takes the fun out of it. 

I personally believe that taking the fun out of coding for people who is passionate about what they do is NEVER AN 
ACCEPTABLE TRADE OFF. Hopefully, people behind Go (and big companies enforcing linters and whatsoever) will some day 
understand that, if they hire decent coders, there is no need to tell them how to write their code, they just need to 
leave them keep having fun with what they love doing.

Good or bad, Go Mock Yourself code has a zero tolerance linters policy, thus no linter was used whatsoever in this project.
