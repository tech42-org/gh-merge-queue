# gh-merge-queue

Sample Pets API used to exercise GitHub merge queues.

## Running

```sh
go run .
```

The server listens on `:8080`.

## Endpoints

- `GET /pets` — list all pets
  `curl localhost:8080/pets`
- `POST /pets` — create a pet
  `curl -X POST localhost:8080/pets -d '{"name":"Rex","species":"dog"}'`
- `GET /pets/{id}` — get a pet
  `curl localhost:8080/pets/1`
- `PUT /pets/{id}` — replace a pet
  `curl -X PUT localhost:8080/pets/1 -d '{"name":"Rexy","species":"dog"}'`
- `DELETE /pets/{id}` — delete a pet
  `curl -X DELETE localhost:8080/pets/1`

## Testing

```sh
go test ./...
```
