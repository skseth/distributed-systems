# Distributed Systems

## First rule of distribution - don't do it

* Scalability! But at what COST? - Frank McSherry, Michael Isard, Derek G. Murphy
  https://www.usenix.org/conference/hotos15/workshop-program/presentation/mcsherry

* Mosaic: Processing a Trillion-edge Graph on a Single Machine
  https://taesoo.gtisc.gatech.edu/pubs/2017/maass:mosaic.pdf  
  Paper from Georgia Tech on the same note as above. Trillion is Facebook-scale!

* LMAX Disruptor  
   A fast non-blocking inter-thread queue  
   See its architecture for how to extract huge performance out of Java.

* Mechanical Sympathy Blog. Martin Thompson's blog
  See his talks on infoq. This is for hardware oriented systems geeks like me. 

## Second rule of distribution - if you do it, you cannot hide it

Most distributed systems try to create abstractions to hide distribution and present a "single system image". However, all abstractions are leaky to some extent, and ultimately the abstraction will leak under some circumstances.

* Law of Leaky Abstraction - Joel Spolsky  
  https://www.joelonsoftware.com/2002/11/11/the-law-of-leaky-abstractions/

## Linearizability

Incoming concurrent requests are executed in a linear order which preserves temporal order.

3 properties

* External consistency - channels of communication between external actors can be used to verify effect e.g. user A writes and calls user B to check, outside of system. But if there is no other way for users to communicate (e.g. memory ), then sequential consistency is enough, linearizability is not necessary. Program order is still preserved. Atomicity is preserved.

Spanner - linearizable 

* Compositionality

Theorem - A linerizable system is compositional - allows splitting into further linearizable steps, and vice versa

* Non-Blocking

Lock order e.g. concurrent tree
Big issue with systems which are prone to failure - split brain scenario
Need progress guarantee

* Linearizability: A Correctness Condition for Concurrent Objects - Maurice Herlihy & Jeannette Wing  
  https://cs.brown.edu/~mph/HerlihyW90/p463-herlihy.pdf

## Replication

Primary-Backup
  Problem : Split Brain
  Solutions : Quorum on reads, Monitoring system

Quorum: R + W > N, W > N/2
Does this guarantee consistency 

Failover - avaoiding split brain :
Option 1 : 
HA - Pacemaker
HAProxy
Amazon RDS

Option 2: External Consensus systems
Chubby/ZK
Timed Keys/Lease - Lock + Time + Compare and Swap

Option 3


Cassandra : Quorum on reads, Uses real time as ordering mechanism

CAP theorem
  Consistency : Linearizability

Client-side-id
Google File System - 2 level kv store
Chubby
Atthiya & Welch Algorithm 
Should all clients need to agree on quorum numbers
Chubby paper
Raft paper - John Ousterhout

## Raft Edge Cases

Monotonic Reads - index must be higher 
e.g. Kafka - shared log

- **Time, Ordering and Events in a Distributed Systems - Leslie Lamport**
  http://amturing.acm.org/p558-lamport.pdf

- **End to End Arguments in System Design** 
  http://web.mit.edu/Saltzer/www/publications/endtoend/endtoend.pdf

- **Spanner**
  Spanner Paper

## FOLLOW-UP

- Anil Madhavapeddy - Unikernels
- http://www.seriouseats.com/the-food-lab#techniques
