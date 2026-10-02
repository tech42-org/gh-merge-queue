# gh-merge-queue

Sample Pets API used to exercise GitHub merge queues.

## Running

```sh
go run .
```

The server listens on `:8080`.

## Endpoints

| Method   | Path         | Description     | Success |
|----------|--------------|-----------------|---------|
| `GET`    | `/pets`      | List all pets   | 200     |
| `POST`   | `/pets`      | Create a pet    | 201     |
| `GET`    | `/pets/{id}` | Get a pet       | 200     |
| `PUT`    | `/pets/{id}` | Replace a pet   | 200     |
| `DELETE` | `/pets/{id}` | Delete a pet    | 204     |

## Testing

```sh
go test ./...
```
