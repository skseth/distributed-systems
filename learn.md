
[Network Programming in Go](http://jan.newmarch.name/go/)
[libchan]

[Course with Great Readings](http://www.andrew.cmu.edu/course/15-749/READINGS/)


[Accrual-Based Failure Detector]
http://ddg.jaist.ac.jp/pub/HDY+04.pdf
https://github.com/dgryski/go-failure

[Connection Pools](https://code.facebook.com/posts/1499322996995183/solving-the-mystery-of-link-imbalance-a-metastable-failure-state-at-scale/)
http://vladmihalcea.com/2014/04/17/the-anatomy-of-connection-pooling/
http://vladmihalcea.com/2014/04/25/flexy-pool-reactive-connection-pooling/
https://github.com/vladmihalcea/flexy-pool
http://vladmihalcea.com/2014/05/20/the-simple-scalability-equation/

http://repo.status302.com/doc/afMongo/src-ConnectionManagerPooled.fan.html

[Metrics]
[Coda Hale](https://dropwizard.github.io/metrics/3.1.0/)

[Distributed Hash Table](http://secondbit.org/blog/introducing-pastry/)

[Pastry Hash Table](http://www.wikipedia.org/wiki/Pastry_%28DHT%29)

[Leader Election](https://www.found.no/foundation/leader-election-in-general/)

http://www.cse.iitd.ac.in/~srsarangi/csl860/docs/pastry-lec.pdf
http://research.microsoft.com/en-us/um/people/antr/past/pastry.pdf

[Leslie Lamport](https://www.youtube.com/watch?v=JG2ESDGwHHY&list=PLvftSirbHo6tVk0IC7JvJiTQr7waUlTcE)

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

How cloud client systems work. music or video players, Shazzam

Network routing : BGP, high availability routers, future secure networks. Routing has a key role to play in the cloud. 

P2P file sharing and search : This latter functionality is often called a “key,value” store and is frequently implemented using a “distributed hash table”. We’ll study CAN and Chord, two widely known examples of DHTs.

First and second tier cloud services : Best-known examples include Memcached and Dynamo. We’ll also look briefly at GFS, BigTable, Chubby, Zookeeper, Sinfonia, MapReduce and other services and will see where each lives. 

Adaptive Overlays for the tier 2 cloud : Beehive, Pastry (a Plaxton-tree DHT structure that creates flexibility as to which nodes a given node should peer with) and Kelips (a gossip-based DHT with O(1) lookup times).

Torrents and Tit-for-Tat Incentives. BAR Gossip 

Anatomy of a cloud. CAP theorem 

"BASE methodology versus ACID model : 

Consistency and time. Logical and Vector clocks. Consistent cuts. 

2 and 3-phase commit 

Consensus and the FLP theorem : FLP asserts that fault-tolerant consensus is “impossible” in asynchronous systems. What do these terms mean? How is FLP proved? What does FLP mean in practice?

Paxos : This lecture will present the “basic” Paxos protocol. Paxos is a powerful protocol for managing replicated data and widely used in cloud computing systems. The Chubby lock service mentioned in lecture 7 runs on top of Paxos.

Atomic multicast with Virtual Synchrony guarantees (Isis2), Atomic multicast, "Atomic multicast: How much ordering is needed?

Transactional model, ACID 

Implementing transactional services 

Bimodal Multicast. Astrolabe. Notion of “information diameter” of a system 

Building complex overlay networks using Gossip 

The RealTime Cloud 

Information Retrieval in the Facebook cloud 

Security concerns for consolidated systems. Synthetic diversity. 

Future cloud I: Computing on encrypted data

Future cloud II: Cloud virtualization 

The cloud value proposition. 


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

[Posix - Linux Threading](http://www.cs.utexas.edu/~witchel/372/lectures/POSIX_Linux_Threading.pdf)

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
[Gul Agha's Thesis](https://www.cypherpunks.to/erights/history/actors/AITR-844.pdf)
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

http://www.quora.com/What-are-the-seminal-papers-in-distributed-systems-Why

[2PC](http://research.microsoft.com/en-us/um/people/gray/papers/DBOS.pdf)
[FLP](http://theory.lcs.mit.edu/tds/papers/Lynch/jacm85.pdf)
Paxos made live, Paxos for system builders, ZooKeeper Atomic Broadcast
http://labs.google.com/papers/pa...
http://www.cnds.jhu.edu/pub/pape...
http://research.yahoo.com/pub/3514

[Excellent Review of Key Papers in Distributed Systems](http://betathoughts.blogspot.in/2007/06/brief-history-of-consensus-2pc-and.html)

[Explanation of Paxos](http://research.microsoft.com/en-us/um/people/blampson/58-Consensus/Acrobat.pdf)

[Swami's List](http://scalingsystems.com/2011/09/07/reading-list-for-distributed-systems/)

[BFT - Yahoo Experience](http://www.sigops.org/sosp/sosp09/slides/song-slides-sosp09wip.pdf)