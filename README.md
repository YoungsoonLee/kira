# Skill stacks.  
    Go, Docker, Docker-compose , mongoDB, gorilla/mux, 

# Solution.  
    * Check. whether there is overlaps event or not, when a new event data insert through the query.
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

    for add a new event data> 
    $ curl localhost:8080/event -d '{"text":"Awesome Kira", "start_at":"2019-01-01T00:00:00Z", "end_at": "2019-01-10T00:00:00Z"}' -X POST  
    (* you can see 200 success when there is no overlap data)  


    (* if you try abow curl again, you can receive 400 Bad request(there is overlap data) and overlap data.)  

    for test> 
        $ docker exec -it api /bin/bash
        $ go test ./...  
        * make test event data like that  


