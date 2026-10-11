# kv-store: a small Raft-style key-value store (learning project)

Language: Go. Transport: gRPC. Goal: understand Raft by building it, not to ship it.

## Scope (in)
- 3 nodes, fixed membership, config passed by flags.
- Leader election (terms, votes, randomized timeouts).
- Log replication with AppendEntries and commit index.
- In-memory log and in-memory KV state machine.
- Client API: Get, Put, Delete. Writes go through the leader.
- Followers redirect clients to the leader.

## Scope (out)
- Snapshots and log compaction.
- Membership changes.
- Persistence to disk (add later only if wanted).
- Linearizable reads (reads are served by the leader without ReadIndex).
- TLS, auth, metrics.

## Layout
- cmd/kvnode        main: flags, start gRPC server, start raft.
- proto/            raft.proto (RequestVote, AppendEntries), kv.proto (Get, Put, Delete).
- internal/raft     state, election, replication, timers. No gRPC types here.
- internal/kv       state machine: apply(log entry) -> map.
- internal/transport gRPC adapters that call raft and implement the proto services.
- test/             multi-node tests using in-process gRPC (bufconn).

## Milestones
1. Skeleton: go mod, protos, a node that starts and answers Get with empty results.
2. Election: nodes elect one leader. Kill leader, a new one wins.
3. Replication: Put on leader shows up on all nodes after commit.
4. Client redirect: follower returns leader address; a tiny CLI client follows it.
5. Fault tests: partition one node, send writes, heal, check logs converge.
6. Optional: persist term, vote, and log to a file.

## Rules for myself
- Keep raft package free of networking so it can be unit tested with fakes.
- Read the Raft paper section 5 before milestone 2 and 3.
