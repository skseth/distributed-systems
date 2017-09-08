# Raft

https://raft.github.io/raft.pdf

## Key Features

Raft allows a set of distributed, replicated state machines to compute the same state, from point of view of a client. Each state machine is deterministic i.e. given a sequence of commands and the same initial state, all machines end up in the same final state.

The goal of the system is to ensure that all state machines execute commands in the same order, such that they appear as a single, consistent state machine to an external observer. This should happen even under various kinds of failure. To break this up the system must ensure :

* Safety : Under network delays, partitions, packet loss, duplication and reordering of messages, no incorrect response should be returned.

* Availability : System should be available as long as a simple majority of servers are operational and can communicate with each other. Failure may be shutdown or pause. Failed servers may rejoin. Failure may also be a slow server - one slow server cannot make system unavailable.

* No need for agreed time : System must not depend on agreed timing between distributed nodes

* We should be able to correctly add or remove machines from the cluster

## Raft Consensus Algorithm

In order to achieve consensus among the machines, each machine executes commands from a persistent log which contains command. Now the problem is to achieve consensus on the commands in the log, which can then be executed against the state machine. 

To achieve consensus on the log, Raft has a distinguished **strong leader**. The leader has the sole responsibility for managing the log. The algorithm provides the following key safety guarantee : If any machine has applied a command at a particular log index to its state machine, no other machine may apply a different command at the same index.

### Structure of the state machines

A raft cluster contains several servers - eg five servers would allow 2 failures. Each server may be either a leader, follower or candidate (to become a leader). A leader election is held to elect a leader for a "term". 

The machines communicate via RPCs. During normal operation, the only RPC sent is a AppendEntries RPC message which is sent by the leader for log replication and heartbeat. All other machines are passive. During elections, an additional RPC, RequestVote may be sent by candidates. Followers are passive.

If a leader fails, a new leader is elected for a new term - this is to ensure that old leaders do not create an issue when they come back to life for any reason. Election of a leader requires a majority of machines to agree on a leader - sometimes there is a split vote. In that case the term closes with no leader and a new term is started and a new election is held. Every server holds the current term number, and this is guaranteed to increase monotonically. Every message exchanged in the cluster contains the term. If a request has a stale term number, the receiver rejects it. If a candidate or leader detects it has a stale term, it reverts to follower state.

Each log entry is provided associated with a log index, and a term. Raft guarantees that any entry that is committed will always appear in the log of all future leaders (Election Safety). Note that committed means that the entry is written to a majority of servers. This impacts how entries are committed and how leaders are elected.

### Leader Election

The AppendEntries message is also used as a heartbeat - when a follower node receives no communication for a specific period (election timeout), it then makes itself a candidate, increments its term, and starts an election.

An election involves voting for itself, and sending out a RequestVote to each machine. 

If a majority vote yes, then the candidate becomes the new leader. To ensure no repeat votes are possible in a term,votes are persisted along with the term. Thus if a machine receives a RequestVote and has already voted, it refuses to vote for the candidate.

Alternatively the candidate may receive an AppendEntries message from another server with a same or higher term, and revert to being a follower. 

Finally, the votes may split and there may be no winner. In that case the term is incremented, and a randomized timeout (between 150-300ms) is used by each candidate before a fresh election. This ensures that a leader has sufficient time to gather majority votes before another candidate wakes up.

One question is : When should a follower vote for a candidate? Given the safety guarantee, it is necessary that any leader has the latest committed log entries in its logs, so no node should be elected leader whose log does not contain the committed entries, because these committed entries may have been executed in a state machine and then a response sent to a client. 

To maintain this, the RequestVote carries the (index+term) of the last entry of the follower. If the term of the follower is greater, or it is equal but the index of its last entry is lower, then it votes for the candidate. This ensures that the leader having the highest index of the highest term is selected leader.

## Log Replication

Once a leader is elected, it starts processing client commands, and also maintains a heartbeat with other servers. 

When a new command is received, it is added to the log of the leader with the current term of the leader, and the log index incremented. The leader sends AppendEntries RPC to all other machines. The AppendEntries message contains 1) the last confirmed committed index 2) the term, index of the previous entry and 3) the term, index and command for the latest entry.

A follower receiving the message does the following : If the follower does not find the previous (index+term) entry in its log, it rejects the message. If it does,it stores the message and returns ok. It also executes all messages upto the latest committed message in the AppendEntries call. 

If the leader gets an acceptance from a majority of followers, it commits i.e. executes the state machine and replies to the client. Note that all messages from the earlier committed index to the latest one are considered committed.

An important property of the log replication is that when an entry is marked as committed, it also commits all previous entries in the leaders logs, because we can confirm that a majority of servers agree on this and all previous entries (Log Matching Property). 

Followers may have different logs from leader. For this a leader maintains a nextindex for each follower, and decrements it whenever an AppendEntries with that index fails, until it gets to a matching entry and then the entries start filling up.

1) Straggler : A node is at an earlier point, even in earlier term. The nextindex reduces till the matching entry and then the log starts filling up.

2) Extra entries : a node has entries not present in the leader. The (index + term) should match, and the extra entries are removed and the new one inserted.

3) Divergent - the index+term of the previous value in AppendEntries does not match the log. The node rejects the message, forcing same behaviour as straggler, until a matching index+term is found.

One important point is that a leader will only mark messages from earlier messages as committed, after at least one message from its current term has been committed. This is necessary to ensure that the safety property is maintained.












