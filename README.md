# Skill stacks.  
    Go, Docker, Docker-compose , mongoDB, gorilla/mux, 

# Solution.  
    * Check. whether there is overlaps event data or not, when a new event data insert through the query.
    * The meaning of overlaps data is as follows.
        - If the start date and end date of the event data to be newly input exists  
          between the start date and end date of the previously input data.  
        Case 1: start_at of a new event data <= start_at of existing the event data <= end_at of a new event data  
        Case 2: start_at of a new event data <= end_at of existing the event data <= end_at of a new event data  
        Case 3: start_at of existing the event data  <= start_at, end_at of a new event data <= end_at of existing the event data  
        Case 4: start_at of a new event data <= start_at, end_at of existing the event data <= end_at of a new event data  

    * There is main logic in routes > event.go for the solution.  

# How to Useage.  
    $ git clone https://github.com/YoungsoonLee/kira.git  
    $ cd kira
    $ docker-compose up -d --build
    $ docker ps  (*check run this project well)
    ![]()  
    
    for add a new event data> 
    ex) $ curl localhost:8080/event -d '{"text":"Awesome Kira", "start_at":"2019-01-01T00:00:00Z", "end_at": "2019-01-10T00:00:00Z"}' -X POST -v  

    * you can receive HTTP 200 code success when there is no overlap data
    ![]()  

    * if you try abow curl again, you can receive HTTP 400 code Bad request(there is overlap data) and overlap data.
    ![]()  

    for test> 
        * use testwithCurl.sh(like live test)
            $ chmod 0700 testwithCurl.sh && ./testwithCurl.sh 

        * in api contanior(like local test)
            $ docker exec -it api /bin/bash
            $ go test ./...  
                - delete all collections and then make test events data in DB 

        * all test make test event data like below.  
        {
            "text" : "event number #0",
            "start_at" : ISODate("2018-06-01T00:00:00.000Z"),
            "end_at" : ISODate("2018-06-10T00:00:00.000Z")
        },
        {
            "text" : "event number #1",
            "start_at" : ISODate("2018-06-11T00:00:00.000Z"),
            "end_at" : ISODate("2018-06-20T00:00:00.000Z")
        },
        {
            "text" : "event number #2",
            "start_at" : ISODate("2018-06-21T00:00:00.000Z"),
            "end_at" : ISODate("2018-06-30T00:00:00.000Z")
        },
        {
            "text" : "event number #3",
            "start_at" : ISODate("2018-07-01T00:00:00.000Z"),
            "end_at" : ISODate("2018-07-10T00:00:00.000Z")
        },
        {
            "text" : "event number #4",
            "start_at" : ISODate("2018-07-11T00:00:00.000Z"),
            "end_at" : ISODate("2018-07-20T00:00:00.000Z")
        },
        {
            "text" : "add a new event, no overlap",
            "start_at" : ISODate("2018-10-01T00:00:00.000Z"),
            "end_at" : ISODate("2018-10-10T00:00:00.000Z")
        }


