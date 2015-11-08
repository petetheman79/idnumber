# idnumber
-- Test attempted by Peter Solomon 08 November 2015

* Database persistance layer not yet implemented
* Tests not yet implemented
* No prior experience in the GO language
* Used http://geekswithblogs.net/willemf/archive/2005/10/30/58561.aspx as a reference for the SA ID number validation.


Original test requirements below:

The objective of the test is to display capabilities in the following areas:
•	Test Driven Development
•	Performant coding
•	Design and planning
•	Efficiency

Scenario 1
•	The project requires 2 components ...  a server and a GUI client.
•	The client GUI must have the facility to input a single  ID number as well as upload a file containing a list of ID Numbers separated by a new line
•	We want the server-side to be a REST API (using MVC) and write to a persistant layer
•	Something small and compact is preferred.

Please use the below rules to process the process the ID Numbers and display results in grid format

You are required to parse the 13 digit number, and supply the following:
•	Validity
•	Date of birth
•	Gender
•	Citizenship

You are required to process the entire list concurrently, and write the following results:
•	List of valid id numbers, with the breakdown as per the above validation, for example: 7808215185082,21 Aug 1978,Male,SA.
•	List of invalid id number, without the breakdown.

Rules:
•	You have to do the validation, then the breakdown, then write to file
•	You are only allowed to have a single writer to the valid result file.
•	You are only allowed to have a single writer to the invalid result file.
•	Code must be written in GO
