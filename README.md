# gh-merge-queue

Sample Pets API used to exercise GitHub merge queues.

## Running

```sh
go run .
```

The server listens on `:8080`.

## Endpoints

- `GET /pets` — list all pets
- `POST /pets` — create a pet
- `GET /pets/{id}` — get a pet
- `PUT /pets/{id}` — replace a pet
- `DELETE /pets/{id}` — delete a pet

## Testing

```sh
go test ./...
```
