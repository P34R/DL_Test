# Distributed Lab Test Task
### Important notes
I slightly modified the initial data (the file with them).  
The original first line looked like this:  
`907;1909; 1929;1790.87;00:50:00;20:00:00`  
Here you can see the "extra" space before the number 1929.I removed it manually.  
The last line looked like this:  
`9101;1902;1921;160.56;05:15:00;20:50:00}`  
Here you can see the symbol "}" at the end of the line. I also deleted it manually.

In the task, it was required to find the best travel options for all stations according to price and time.  
The "time" in my solution is "journey time" (I don't take into account the waiting time between trains). For example, we have 3 stations A, B and C. Time A->B 12:00 pm -> 13:00 pm (departure time -> arrival time).   
B->C 19:00 -> 20:00. This is the best option (for example) and its time will be 2 hours (actually 8 hours if we take into account the time we spend waiting for the second train, but we don't).  

Also, I must say that I display the path by station (station A -> station B -> station C, etc.), 
but there are many "same" trains in the test data. 
For example, trains with numbers 907 and 908 have exactly the same parameters (except for the number). 
This means that "our path" can be traveled by different trains, the cost will be the same.

### Solution
Initially, I analyzed the data (with the help of various functions that were redesigned to be more "beautiful").  
There are 6 stations in the test data (If you change the file, everything should work the same probably).  
The algorithm can be called "brute force". By increasing the number of stations, the algorithm may run for a very long time.  
Initially, I go through the data and build a matrix (let's call it arr) of size K * K (K = number of stations)  
arr[i][j] - the cheapest way (in cost or journey time) from station i to station j (which station is "i" and "j" we can find out from stationsMap).  
After that, I iterate over all possible options (recursively) to go through all the stations (if the path A -> B does not exist, this path is folded back) and calculate the cost of this path.  
When the algorithm has reached the last unvisited station, it enters this path into "path" array and it's cost into "pathCosts" array.  
After that, I go through all the paths found and look for the minimum cost. Then, I again go through all possible paths and output to the console all that have a minimum cost.  
Same repeats for "time" minimum.  
Also, I should say that the time is displayed in seconds.  

### How to launch 

Firstly, clone the repository using Git:  
`git clone https://github.com/P34R/DL_Test`  

After that, enter created folder using  
`cd DL_Test`  

Then  
`go run main.go logic.go`  

You should see something like that:  
```
--------COST---------

path [ A->B->C->... ] cost FLOAT

--------TIME (in seconds)---------

path [ A->B->C->... ] cost INT
```
where A->B->C->... is our path (A,B,C are stations), 
FLOAT is a floating point number that means the cost of this path (price), 
INT is the cost of the path in seconds (time)

If you want to change test data, name it `test_task_data.csv` and put into folder or change line 4 of main.go to  
`k := ParseCSV("FILENAME.csv")`  
where FILENAME.csv - name of your .csv file (data must be in the same format).  
