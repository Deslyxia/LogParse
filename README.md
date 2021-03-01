# LogParse
Simple csv log parser


This is a simple log parser. The goal of this was simple:

Read in a csv file
Marshall it out into an array of structs
Answer a few questions about the parsed logs


Approach:

Reading the csv file is pretty straightforward. Go has a built in os library for that. Once I had the file I like to marshall into Structs 
because it is easier/cleaner to work with inside the code itself rather than parsing the csv directly into a 2d array that needs to be managed.
The first couple of questions were easy:

1) How many uninque users are represented in the logfile?

      The approach here was fairly simple. I used a map where username was the key. As I iterated over the array of Log structs I looked at the username 
      field and checked to see if that key already existed in the map. If it did I added to the count in the value column. If not I added the key and moved
      to question 2. 

      Interesting byproduct here is that using this method I could tell you how many unique users there were AND how many times each of them did an operation. 
      
2) How many uploads were greater than 50kb?

      Also a simple approach. as I parsed through the array I simply checked if the operation type was 'upload' and the size was > 50kb
      
3) How many times did user 'jeff22' log in on Apr 15th?
 
      The approach here was straight forward though slightly more in depth than the first two. Using the Time library I just made a check time at EOD on the 14th 
      and the Stard of Day on the 16th and user the time.Before() time.After() in the check to find the answer. 



Additional considerations:

All in all this was a pretty quick exercise. Mainly the time was spent in becomming reaquainted with the nuances of the Time library. As a simple one off this exercise
is a good one that in a short period of time dove into various aspects of the language. With more time and with some sort of production problem to solve around this 
basic issue I could see a few things to expand upon. 

I would expect there to be no shortage of log files. As such I would split this up into pieces. Piece 1 would be a fetcher of sorts that actively grabbed log files off
of the various systems for processing. thinking about it very quickly maybe this fetcher grabs the csv files, parses them and drops the logs into a queue. Piece 2 would 
be the processor. Its job would be to pick up the logs off of the queue and "process" them. For sake of argument maybe there is some sanitization or dedup work that needs
to be done before pushing the logs to something like Kibana. Piece 3 is Kibana or something like it where an end user could dictate any questions he/she wanted against 
the logs. 
    
