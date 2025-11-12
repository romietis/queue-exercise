## API Endpoints

### Queue Service (localhost:8000)
- `POST /add-item` - Add an item to the queue
- `GET /get-item` - Retrieve and remove the first item from the queue

### Worker Service (localhost: 8080)
- `POST /send` - Send multiline file to the queue service (splits by newline and adds each line as separate item)
- `GET /receive` - Retreive on eitem from the queue and append to `output.text`

## Running the application

### Start the Queue service

```bash
go run cmd/queue/queue.go
```

### Start the Worker Service

```bash
go run cmd/worker/worker.go
```


### Send file via the Worker

```bash
curl -X POST --data-binary @input.txt "http://localhost:8080/send"
```

it is important to send as a binary so curl doesn't wrap request with additional fields

### Receive data via the Worker

```bash
curl "http://localhost:8080/receive"
```

And the file `output.txt` should be written line by line as many curl requests you execute

## Testing

```bash
go test -v -coverprofile=cover.out  ./...  
```

Visual code coverage
```bash
go tool cover -html=cover.out
```

## TODO
- [ ] Potect Queue from simultanues access by mutex
- [ ] Write dockerfiles for separte services using multi-stage build
- [ ] Restructuring the code to avoid tight coupling - http handlers are direclty in main
- [ ] Using web framework like Gin or Chi
