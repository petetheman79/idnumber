-- Case Study attempted by Peter Solomon 08 November 2015<br/><br/>
No prior experience in the GO language.<br/>
Used http://geekswithblogs.net/willemf/archive/2005/10/30/58561.aspx as a reference for the SA ID number validation.
<br />
<br />
<br />
Original test requirements below:
<br />
The objective of the test is to display capabilities in the following areas:
<ul>
<li>Test Driven Development</li>
<li>Performant coding</li>
<li>Design and planning</li>
<li>Efficiency</li>
</ul>
<br /><br />
Scenario 1
<ul>
<li>The project requires 2 components ...  a server and a GUI client.</li>
<li>The client GUI must have the facility to input a single  ID number as well as upload a file containing a list of ID Numbers separated by a new line</li>
<li>We want the server-side to be a REST API (using MVC) and write to a persistant layer</li>
<li>Something small and compact is preferred.</li>
</ul>
<br /><br /?
Please use the below rules to process the process the ID Numbers and display results in grid format
<br /><br />
You are required to parse the 13 digit number, and supply the following:
<ul>
<li>Validity</li>
<li>Date of birth</li>
<li>Gender</li>
<li>Citizenship</li>
</ul>
<br /><br />
You are required to process the entire list concurrently, and write the following results:
<ul>
<li>List of valid id numbers, with the breakdown as per the above validation, for example: 7808215185082,21 Aug 1978,Male,SA.</li>
<li>List of invalid id number, without the breakdown.</li>
</ul>
<br /><br />
Rules:
<br>
<ul>
<li>You have to do the validation, then the breakdown, then write to file</li>
<li>You are only allowed to have a single writer to the valid result file.</li>
<li>You are only allowed to have a single writer to the invalid result file.</li>
<li>Code must be written in GO</li>
</ul>