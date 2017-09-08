# Distributed Systems & Computer Architecture

## Papers & Articles

* Scalability! But at what COST? - Frank McSherry, Michael Isard, Derek G. Murphy  
  https://www.usenix.org/conference/hotos15/workshop-program/presentation/mcsherry

* In Search of an Understandable Consensus Algorithm  
  https://raft.github.io/raft.pdf

### In Progress

* What Every Programmer Should Know About Memory
  http://futuretech.blinkenlights.nl/misc/cpumemory.pdf

### TODO

* Law of Leaky Abstraction - Joel Spolsky  
  https://www.joelonsoftware.com/2002/11/11/the-law-of-leaky-abstractions/

* Go's Memory Model - Russ Cox  
  http://nil.csail.mit.edu/6.824/2016/notes/gomem.pdf

* Mosaic: Processing a Trillion-edge Graph on a Single Machine
  https://taesoo.gtisc.gatech.edu/pubs/2017/maass:mosaic.pdf  
  Paper from Georgia Tech on the same note as above. Trillion is Facebook-scale!

* LMAX Disruptor  
   A fast non-blocking inter-thread queue  
   See its architecture for how to extract huge performance out of Java.

* Time, Ordering and Events in a Distributed Systems - Leslie Lamport**
  http://amturing.acm.org/p558-lamport.pdf

* End to End Arguments in System Design   
  http://web.mit.edu/Saltzer/www/publications/endtoend/endtoend.pdf

* Spanner
  https://research.google.com/archive/spanner.html

* Communicating Sequential Processes - C.A.R. Hoare
  http://lass.cs.umass.edu/~shenoy/courses/677/readings/Hoare.pdf

* Wait-Free Synchronization  
  https://cs.brown.edu/~mph/Herlihy91/p124-herlihy.pdf

* Stack Overflow - The Architecture  
  https://nickcraver.com/blog/2016/02/17/stack-overflow-the-architecture-2016-edition/

* The Anatomy of a Large-Scale Hypertextual Web Search Engine
  http://infolab.stanford.edu/~backrub/google.html

* Scalable commutativity rule  
  https://people.csail.mit.edu/nickolai/papers/clements-sc.pdf

* Advanced Topics in Programming Languages: A Lock-Free Hash Table - Cliff Click  
  https://web.stanford.edu/class/ee380/Abstracts/070221_LockFreeHash.pdf
  https://www.youtube.com/watch?v=HJ-719EGIts

* A pragmatic implementation of non-blocking linked lists  
  https://timharris.uk/papers/2001-disc.pdf

* Solution of a Problem in Concurrent Programming Control - E. W. DIJXSTRA  
  http://www.di.ens.fr/~pouzet/cours/systeme/bib/dijkstra.pdf

* Shared Memory Consistency Models - A Tutorial - Sarita Adve, K. Gharachorloo  
  http://www.hpl.hp.com/techreports/Compaq-DEC/WRL-95-7.pdf

* Linearizability: A Correctness Condition for Concurrent Objects - Maurice Herlihy & Jeannette Wing  
  https://cs.brown.edu/~mph/HerlihyW90/p463-herlihy.pdf

* Specifying Systems  
  https://www.microsoft.com/en-us/research/wp-content/uploads/2002/12/The-TLA-Language-and-Tools-for-Hardware-and-Software-Engineers-Book.pdf  
  https://link.springer.com/chapter/10.1007/978-3-662-43652-3_3

## Blogs & Sites

* **Mechanical Sympathy Blog**  
  https://mechanical-sympathy.blogspot.in/

* **High Scalabity**  
  https://highscalability.com

## Algorithms to Implement

* TODO - Raft

## Interesting Links

* [Memory Barriers are Like Source Control Operations](http://preshing.com/20120710/memory-barriers-are-like-source-control-operations/)

* [Taming Complexity by Irreversibility](https://martinfowler.com/articles/zaninotto.pdf)

* [Optimizing Web Servers for High Throughput and Low Latency](https://blogs.dropbox.com/tech/2017/09/optimizing-web-servers-for-high-throughput-and-low-latency/)

* [Linux Performance](http://www.brendangregg.com/linuxperf.html)

* [High Performance Browser Networking](https://hpbn.co/)

## Cleanup TODOs
- Operational Transform
- Amdahl's Law
- concurrent trees - locking issues, lock free implementation, compare-and-swap or test-and-set, non-monotonic updates, CRDT, optimistic concurrency with exponential back-off,djikstra mutex algo
-  Command Query Responsibility Segregation
Google File System - 2 level kv store
Atthiya & Welch Algorithm 
Chubby papers
merge learn.md file