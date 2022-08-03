# redis-clone
A [Redis](https://redis.io/) clone written in [Go](https://go.dev/) for [CodeCrafters Redis Challenge](https://app.codecrafters.io/courses/redis/overview)

## Supported Data Types
- [Simple String](https://redis.io/docs/reference/protocol-spec/#resp-simple-strings) (+)
- [Error](https://redis.io/docs/reference/protocol-spec/#resp-errors) (-)
- [Integer](https://redis.io/docs/reference/protocol-spec/#resp-integers) (:)
- [Bulk String](https://redis.io/docs/reference/protocol-spec/#resp-bulk-strings) ($)
- [Array](https://redis.io/docs/reference/protocol-spec/#resp-arrays) (*)

## Supported Commands
- [Ping](https://redis.io/commands/ping/)
- [Echo](https://redis.io/commands/echo/)
- [Set](https://redis.io/commands/set/) (PX)
- [Get](https://redis.io/commands/get/)
