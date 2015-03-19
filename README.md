
[Network Programming in Go](http://jan.newmarch.name/go/)
[libchan]

[Course with Great Readings](http://www.andrew.cmu.edu/course/15-749/READINGS/)

[Distributed Systems Course](http://pdos.csail.mit.edu/6.824-2013/schedule.html)

- MapReduce
- RPC
- Fault Tolerance via Primary / Backup Replication
- Primary / Backup Replication : Flat Datacenter Storage
- Fault Tolerance : Paxos
- Using Paxos in a real service - Replicated File System
- Egalitarian Paxos
- Spanner
- Distributed Transactions - 2pc vs lynx, transaction chains
- Distributed Shared Memory & Sequential Consistency
- DSM and consistency continued
- Resilient Distributed Datasets - Spark
- Optimism, Causality, Vector Timestamps
- Eventual Consistency
- Relaxed Consistency - PNUTS
- Dynamo - Amazons Store eventually consistent
- P2P - BitTorrent, DHTs - Distributed Hash Tables
- Akamai
- P2P : Bitcoin

[CD425 - Prof Indranil Gupta](https://courses.engr.illinois.edu/cs425/fa2014/lectures.html)
- Mapreduce / Hadoop
- Failure Detectors
- Time and Ordering
- Snapshots
- Multicast Communications
- Gossipping
- P2P systems
- Leader Election
- Mutual Exclusion
- Consensus
- Networking and Routing
- RPCs and Marshaling
- Concurrency Control
- Replication Control, Paxos
- Key-Value Stores, NoSQL - Cassandra, Mongo
- Stream Processing and Graph Processing - Storm, Spark, Pregel
- Measurements and Characteristics of Real Distributed Systems
- Sensor Networks, Security
- Distributed Shared Memory
- Self-Stabilization
- Distributed File Systems
- Data center disasters - case studies

http://product.hubspot.com/blog/how-we-deploy-300-times-a-day

[Ken Birmans Course](http://www.cs.cornell.edu/Courses/CS5412/2014sp/Syllabus.htm)

How it works : (2) How cloud client systems work. We’ll look at some typical cloud applications (music or video players, Shazzam) and ask how they do it.

Network routing : (3) BGP, high availability routers, future secure networks. Routing has a key role to play in the cloud. But how does the Internet do routing? What options exist for controlling routing from within a cloud application? We’ll look mostly at BGP but will also discuss the resilient overlay network idea (RON), from MIT. Then we’ll ask how one might make the Internet robust in other dimensions such as security, and how it could support cloud-based routing control.

P2P file sharing and search : (4) Peer-to-peer protocols have transformed the Internet by making it easy for people who don’t know one another to still cooperate to share music or other kinds of files. Underlying this are two aspects: the Internet file transfer protocol (which is trivial) and the ability to search for things in a network where you only know about a few of the members. This latter functionality is often called a “key,value” store and is frequently implemented using a “distributed hash table”. We’ll study CAN and Chord, two widely known examples of DHTs.

First and second tier cloud services : (5) We use the term “tier 2” to talk about services that are nearly as elastic and scalable as the ones in tier 1, but have roles focused on caching. Best-known examples include Memcached and Dynamo. We’ll also look briefly at GFS, BigTable, Chubby, Zookeeper, Sinfonia, MapReduce and other services and will see where each lives. We’ll only get detailed for a few of these (it would take too long to cover all of them), but we will try and identify common themes. Interestingly, although these systems run inside the cloud (not in the network, like Chord), many use the same key-value structure as Chord and in fact many evolved from ideas that originated in the P2P networking community! On the other hand, Chord itself is viewed as too slow for use in the tier 2 cloud (for a small DHT with 1000 members, Chord might need 8 or 9 routing hops to find an item, which is very costly). Another issue is that access distributions might be skewed: there could be hot spots.

Adaptive Overlays for the tier 2 cloud : (6) Motivated by the desire to use a DHT in tier2 to support technologies like the ones mentioned in lecture 5, we'll revisit the pure DHT concept and explore ways of using dynamic adaptation to compensate for skewed access distributions, skewed node-to-node latencies, etc. We’ll look at Beehive (a modified version of Chord that achieves O(c) lookup times for constant c), Pastry (a Plaxton-tree DHT structure that creates flexibility as to which nodes a given node should peer with) and Kelips (a gossip-based DHT with O(1) lookup times).

Torrents and Tit-for-Tat Incentives. BAR Gossip : (7) BitTorrent is widely used both in the P2P network world and inside cloud computing data centers. In this lecture, we’ll look closely at BitTorrent, to see exactly how its “tit for tat” concept of fairness works. This will lead us to look at an idea from UT Austin for a scheme they call “BAR” gossip

 : February break: February 15-19

Anatomy of a cloud. CAP theorem : (8) At every level of the cloud we struggle with deep tradeoffs: rapid response and scalability tend to fight against techniques that provide strong guarantees. Berkeley’s Eric Brewer captured one of these in his CAP theorem, which focuses on tier1 and tier2 and says you can have just two of Consistency, Availability, Fault or Partition Tolerance. We’ll see how CAP applies and how it was proved by Gilbert and Lynch. Although the proof applies only in WAN settings, we’ll also find that CAP has implications in the first tier of a typical cloud computing data center.

"BASE methodology versus ACID model : (9) Many cloud systems are rejecting the database ACID guarantees in favor of what has come to be known as the BASE methodology, which weakens consistency to enhance responsiveness, but can make mistakes because of the weaker properties that must later (“eventually”) be repaired.

Guest lecture by Theo and Stavros."

Consistency and time. Logical and Vector clocks. Consistent cuts. : (10) To reason about stronger forms of consistency, we’ll need some abstractions. One tool that turns be especially useful was created by Leslie Lamport and offers a way to model time in terms of causal relationships. We’ll see how this leads to notions of logical clocks that sometimes make more sense that real clocks.

2 and 3-phase commit : (11)_We’ll look at the famous two and three phase commit protocols, which play a wide range of roles in many kinds of systems. These are used as building blocks, but are also for their own sake. Oddly, it turns out that neither protocol is able to guarantee that progress will occur, although 3PC does better than 2PC. This motivates looking more closely at just when fault-tolerance can be achieved.

Consensus and the FLP theorem : (12) We’ll look at a theoretical result most people refer to as the “FLP theorem”. FLP asserts that fault-tolerant consensus is “impossible” in asynchronous systems. What do these terms mean? How is FLP proved? What does FLP mean in practice?

Paxos : (13) This lecture will present the “basic” Paxos protocol. Paxos is a powerful protocol for managing replicated data and widely used in cloud computing systems. The Chubby lock service mentioned in lecture 7 runs on top of Paxos.

Atomic multicast with Virtual Synchrony guarantees (Isis2) : (14) We’ll review the virtual synchrony model, a less stringent way to implement replication that brings complexity not seen in Paxos, but also gains performance advantages in doing so. Isis2 employs the virtual synchrony model, but also supports Paxos (called SafeSend within the system).

Atomic multicast: How much durability is needed? : (15) We’ll look at the question of weak versus strong durability as it arises when applying Isis2 to replicate data in tier-1 services. Here, we’ll look at the question of how durable a multicast needs to be if the corresponding data is managed in-memory and there isn’t any way to preserve state across node failures and restarts (i.e. there is only soft-state).

"Atomic multicast: How much ordering is needed?

Also: a quick look at a few other features of Isis2 : (16) We’ll take a short guided tour of the Isis2 system. As a starting point we'll see that in addition to various durability properties, Isis2 offers many ordering options. Not every system needs the strongest and most costly ordering guarantees, and by picking the right protocol to match a use case, one can gain big speedups: you ideally want to pay for what you need. In the time remaining, we'll then review many of the other features of the system. All of this adds up to. our theme: ""too many options""! But we'll also see that you can use the system without using more than a tiny subset of the features. Isis2 is a substantial system, like an operating system for distributed computing, and the fancy features are really for advanced users."

Transactional model, ACID : (17) We’ll look more closely at the ACID model mentioned in Lecture 10 and how it works. ACID is the standard way to think about database consistency.

Implementing transactional services : (18) How can ACID-based systems be implemented? We’ll look at some basic mechanisms and will also touch upon snapshot isolation, an approach popular in cloud settings.

"Bimodal Multicast.

Astrolabe. Notion of “information diameter” of a system : (19) This starts a series of “single topic” lectures (the numbering skips because we covered lecture 19 early). Bimodal Multicast is a reliable multicast built using gossip. Astrolabe is a technology used at Amazon.com for scalable monitoring, data mining and aggregation in cloud settings. An amusing bug that was introduced by a well-intentioned optimization leads us to think about the “information distance” between nodes in the system and sheds light on how gossip mechanisms work when scaled up."

Building complex overlay networks using Gossip : (20) We'll look at the work on the T-Man system at the University of Bologna.

The RealTime Cloud : (21) In this lecture we'll talk about cloud computing systems that have real-time goals and the notion of a real-time QoS guarantee for a data distribution service that has associated state. Our focus will be on work IBM did in 1995 for a new air traffic control system that was being created in the US at that time (in fact the project failed, but the reasons are very instructive)

Information Retrieval in the Facebook cloud : (22) This lecture will focus on technologies used within the Facebook cloud. We'll look at the ways that Facebook uses caching to accelerate worldwide photo retrieval, and then we'll dive deeper and look at the TAO subsystem it uses to track relationships and the Haystack storage backend system.

Security concerns for consolidated systems. Synthetic diversity. : (23) Cloud security worries are a big barrier to adoption. The paper we’ll discuss was written by Professors Schneider and Birman as a report on a study looking at these tradeoffs.

Future cloud I: Computing on encrypted data : (24) One very new idea, still very much considered research, imagines that we might upload data to the cloud in an encrypted form and then compute on the data without actually decrypting it. We'll look at a cutting-edge idea that tackles this question, with surprisingly successful results (but also some limitations).

Future cloud II: Cloud virtualization : (25) We all know that virtualization plays a huge role in modern cloud settings but that virtual machine migration isn't really exploited very much. We'll talk about work by Hakim Weatherspoon and a team at IBM in which cloud virtualization is combined with migration to to let an application access sensitive files that you don’t trust the cloud to host. Hakim's approach also allows you to create a virtual image that includes resources like devices that actually live on remote machines.

The cloud value proposition. : (26) The cloud "Value Proposition". We'll discuss some of the major business models for cloud enterprises to try and understand what determines value in the cloud. This turns out to be a much trickier question than you might have expected.


[15-440](http://www.cs.cmu.edu/~dga/15-440/F12/index.html)
"Communication 1, the Internet in a Day, day 1 
"Communication 2, the Internet in a Day, day 2 
"Consistency - Classical synchronization + Go-style synchronization 
"Distributed Filesystems 
"Remote Procedure Calls 
"Distributed Filesystems 2 - AFS, Coda, callbacks 
"Time and Synchronization 
"Distributed Mutual Exclusion 
"Fault Tolerance 1 - Detecting and Correcting Local Faults 
"RAID 
"Concurrency Control 
"Logging and Crash Recovery 
"Consistent hashing and name-by-hash 
"Distributed Replication 
Distributed Replication 2
"Data-Intensive Computing and MapReduce/Hadoop
"Distributed Filesystems for MapReduce / HDFS 
"DNS and Content Delivery Networks 
"Peer-to-Peer 
"Virtual Machines 
"Byzantine Fault Tolerance 
"Security Protocols 
Case Study - Anonymous Routing and TOR
Causally Consistent Wide-Area Replication (TBD)


[Principles of Distributed Computing](http://dcg.ethz.ch/lectures/podc_allstars/)

[Paxos Made Simple](http://research.microsoft.com/en-us/um/people/lamport/pubs/paxos-simple.pdf)

[Part-time Parliament](http://research.microsoft.com/en-us/um/people/lamport/pubs/lamport-paxos.pdf)

[Replication in the Harp File System](http://pdos.csail.mit.edu/6.824-2004/papers/harp.pdf)

[Time, Clocks and the Ordering of Events in A Distributed System](http://research.microsoft.com/en-us/um/people/lamport/pubs/time-clocks.pdf)

[Leslie Lamports Papers](http://research.microsoft.com/en-us/um/people/lamport/pubs/pubs.html)


[Another One](http://www.cs.cmu.edu/~dga/15-440/F12/syllabus.html)

[Sophomoric Parallelism and Concurrency](http://homes.cs.washington.edu/~djg/teachingMaterials/spac/sophomoricParallelismAndConcurrency.pdf)

[Spark Paper](http://www.eecs.berkeley.edu/Pubs/TechRpts/2012/EECS-2012-259.pdf)

[Armstrong Thesis](https://www.sics.se/~joe/thesis/armstrong_thesis_2003.pdf)
[Gul Agha's Thesis](file:///Users/skseth/Downloads/AITR-844.pdf)
[Why computers stop and what can be done about it?](http://www.hpl.hp.com/techreports/tandem/TR-85.7.pdf)
[There's Just No Getting Around It](http://delivery.acm.org/10.1145/2490000/2482856/p30-cavage.pdf?ip=117.192.230.87&id=2482856&acc=OPEN&key=4D4702B0C3E38B35%2E4D4702B0C3E38B35%2E4D4702B0C3E38B35%2E6D218144511F3437&CFID=645522560&CFTOKEN=48606563&__acm__=1426661459_828c606a5b7cf3f44c9530455cbd5e35)
[Programming Erlang](Programming Erlang - Software for a Concurrent World)

http://jonasboner.com/2007/12/19/hotswap-code-using-scala-and-actors/

http://www.infoq.com/news/2008/06/scala-vs-erlang

[Compare of Scala actors vs Go routines?](https://groups.google.com/forum/#!topic/golang-nuts/aAgIQiHVNq4)

https://savanne.be/articles/concurrency-in-erlang-scala/

[Event Based Programming without inversion of Control](http://lampwww.epfl.ch/~odersky/papers/jmlc06.pdf)

[Actors that unify threads and events](http://www.ist-palcom.org/publications/files/haller07actorsunify.pdf)

http://infoscience.epfl.ch/record/98468/files/MatchingObjectsWithPatterns-TR.pdf

[Viewing Control Structures as Patterns of Passing Messages](http://dspace.mit.edu/handle/1721.1/6272)

https://atilanevesoncode.wordpress.com/2013/12/05/go-vs-d-vs-erlang-vs-c-in-real-life-mqtt-broker-implementation-shootout/

[Algorithm + Strategy = Parallelism](https://www.google.co.in/url?sa=t&rct=j&q=&esrc=s&source=web&cd=1&cad=rja&uact=8&ved=0CCIQFjAA&url=http%3A%2F%2Fresearch.microsoft.com%2Fpubs%2F67082%2Fstrategies.ps.gz&ei=G1IKVc-kFJC9uAT1voCYAw&usg=AFQjCNFufdPxntpeSPV3dsUDYw7_zBRUwA&sig2=DO0pHtNqhV-9oOnsLiQaXA&bvm=bv.88528373,d.c2E)

http://www.hpcs.cs.tsukuba.ac.jp/~tatebe/lecture/h23/dsys/dsd-tutorial.html

[BASE](http://delivery.acm.org/10.1145/1400000/1394128/p48-pritchett.pdf?ip=117.192.226.132&id=1394128&acc=OPEN&key=4D4702B0C3E38B35%2E4D4702B0C3E38B35%2E4D4702B0C3E38B35%2E6D218144511F3437&CFID=645522560&CFTOKEN=48606563&__acm__=1426759469_3f3f609ce545656e14b2c08583910f58)

[Myth of Perfect Design](http://www.artima.com/intv/perfect.html)