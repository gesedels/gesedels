# Gesedels 

**Gesedels** is an obsessively over-engineered key-value API, written in [Go 1.24][go] by [Stephen Malone][sm].

- See [`changes.md`][ch] for the full changelog.
- See [`license.md`][li] for the open-source license (BSD-3).

## Command-Line Usage

Flag    | Default     | Description 
------- | ----------- | -----------
`-addr` | `:8080`     | The server address to listen on.
`-cost` | `12`        | The cost to generate a bcrypt hash.
`-dbse` | `./bolt.db` | The database file to connect to. 
`-rate` | `100`       | The number of requests a user can make per hour.

## Database Structure

All data is stored in a single [Bolt database][db]. Key-value pairs are stored as buckets with the key as the bucket name and the value data as byte fields in a map:

```json
"foo": {
  "body": "Bar.",
  "init": "1751435396",
  "hash": "8e2f3a93aeb2dff313fbb6e5b915261f36a8eca426fa7f8bd385f19c2ba287ae",
  "pass": "$2a$12$Bb1CsGvg7FP33U3XCse7tu5Z4VHP8sevkD7cKi8RQ.uyzGLYXxz76"
}    
```

Field  | Description
------ | -----------
`body` | The key's raw whitespace-trimmed value body.
`init` | The key's creation timestamp in unix UTC.
`hash` | A SHA256 hash of the key's `body` field.
`pass` | A bcrypt hash of the key's password (if provided).

## Endpoints

### Design

- All content is encoded in UTF-8 unicode.
- All endpoints (except for `GET /`) return [JSend][js]-formatted JSON data.

### `GET /`

Return [`readme.md`][rm] as a `text/plain` Markdown page.

```text
$ GET /
> 200 "Gesedels is..."
```

### `GET /{key}`

Return the value (the `body` field) of an existing key.

> [!WARNING]
> All keys are public. If a user knows a key name, they can access it.

```text
$ GET /foo
> 200 {"status": "success", "data": "Bar."}
```

[ch]: https://github.com/gesedels/gesedels/blob/main/changes.md
[db]: https://github.com/etcd-io/bbolt
[js]: https://github.com/omniti-labs/jsend
[li]: https://github.com/gesedels/gesedels/blob/main/license.md
[go]: https://go.dev/doc/go1.24
[rm]: https://github.com/gesedels/gesedels/blob/main/readme.md
[sm]: https://github.com/gesedels
