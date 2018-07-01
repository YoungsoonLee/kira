echo ">> 1. make test event data"
curl http://127.0.0.1:8080/event -d '{"text":"event number #0", "start_at": "2018-06-01T00:00:00.000Z", "end_at":"2018-06-10T00:00:00.000Z"}' -X POST
curl http://127.0.0.1:8080/event -d '{"text":"event number #1", "start_at": "2018-06-11T00:00:00.000Z", "end_at":"2018-06-20T00:00:00.000Z"}' -X POST
curl http://127.0.0.1:8080/event -d '{"text":"event number #2", "start_at": "2018-06-21T00:00:00.000Z", "end_at":"2018-06-30T00:00:00.000Z"}' -X POST
curl http://127.0.0.1:8080/event -d '{"text":"event number #3", "start_at": "2018-07-01T00:00:00.000Z", "end_at":"2018-07-10T00:00:00.000Z"}' -X POST
curl http://127.0.0.1:8080/event -d '{"text":"event number #4", "start_at": "2018-07-11T00:00:00.000Z", "end_at":"2018-07-20T00:00:00.000Z"}' -X POST
echo "\r\n"


echo ">> 2. test overlap data, add new event data start_at: 2018-06-04T00:00:00.000Z, end_at: 2018-06-22T00:00:00.000Z "
echo ">>    receive HTTP 400 BAD REQUEST Code and three overlaps data"
curl localhost:8080/event -d '{"text":"test overlap data", "start_at":"2018-06-04T00:00:00Z", "end_at": "2018-06-22T00:00:00Z"}' -X POST -v  
echo "\r\n"

echo ">> 3. test add a new event data start_at: 2018-10-01T00:00:00.000Z, end_at: 2018-10-10T00:00:00.000Z "
echo ">>    receive HTTP 200 SUCCESS Code and a new event data"
curl localhost:8080/event -d '{"text":"add a new event, no overlap", "start_at":"2018-10-01T00:00:00.000Z", "end_at": "2018-10-10T00:00:00.000Z"}' -X POST -v  
