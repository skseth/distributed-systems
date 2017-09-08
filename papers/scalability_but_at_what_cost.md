# Scalability - But at what cost?

In many cases systems which don't scale well may provide much better performance over the range of expected configurations. Many benchmarks published for big data systems are based on algorithms which scale well, but do not compare the performance to algorithms which don't.

This paper corrects the record by calculating the COST (Configuration that outperforms a single thread) for various big data systems for various graph algorithms, using two datasets :

- Twitter (41 M nodes, 1.47 B edges, 5.7 GB data)
- uk-2007-05 (105 M nodes, 3.74 B edges, 14.72 GB data)

Comparisons have been made to 

- GraphChi : Disk based graph computation using parallel sliding window approach
- Stratosphere 
- X-Stream
- Spark
- Apache Giraph 
- Graphlab
- GraphX

Algoriths compared :

- PageRank

Single threaded implementation with SSD : 300s (Twitter) / 651s (uk)
Single threaded implementation (Hilbert order) with SSD : 242s (Twitter) / 256s (uk)
GraphLab : 249s (Twitter) on 128 cores / 833s (uk) on 128 cores was the best performer among the rest

- Connected Components 

Single threaded Label Propogation with SSD : 153s (Twitter) / 417s (uk)
Single threaded Union-Find with SSD : 1s (Twitter) / 15s (uk)
GraphLab : 242s (Twitter) / 714s (uk) on 128 cores

## Thoughts

https://news.ycombinator.com/item?id=11855594

- The author has chosen graph algorithms for the paper, which often have trouble scaling. Problems which are more naturally parallelizable will do much better

- If data is already distributed, distributed computing may improve I/O, which can be consideration in the decision to parallelize

- Human populations are not growing at the speed at which compute / storage has grown. So for many, many business problems distribution may not be needed

- As a rule of thumb, when faced with scaling a single-threaded implementation, 

- https://news.ycombinator.com/item?id=11856093  
I like the analysis, basically it says "hey you don't have big data" :-) but that requires a bit more explanation.
The only advantage of clustered systems like Spark, Hadoop, and others is aggregate bandwidth to disk and memory. We know that because Amdahl's law tells us that parallelizing something invariably adds overhead. So from a systems perspective that overhead has to be "paid for" by some other improvement, and we'll see that it is access to data.
If your task is to process a 2TB data set, on a single machine using a 6GBS SATA channel and 2TB of FLASH SSDs you can read in that dataset into memory in 3333 seconds (at 600MB/sec which is very optimistic for our system), process it, and lets say you write out a 200GB reduced data set for another 333 seconds. so, conveniently, an hour of I/O time.
Now you take that same dataset and distribute it evenly across a thousand nodes. Each one then has 2GB of the data on it. Each node can read in their portion of the data set in 3 seconds, process it and write out their reduction in .3 seconds.
You have "paid" for the overhead of parallelization by trading an I/O cost of an hour for an I/O cost of about 4 seconds.
That is when parallel data reduction architectures are better for a problem than a single threaded architecture. And that "betterness" is purely artificial in the sense that you would be better off with a single system that had 1,000 times the I/O bandwidth (cough mainframe cough) than 1,000 systems with the more limited bandwidth. However a 1,000 machines with one SSD it still cheaper buy than one mainframe of similar capability. So if, and its a big if, your algorithm can be expressed as a data map / reduce problem, and your data is large enough to push the cost of getting it into memory to have a look at significantly beyond cost of executing the program, then you can benefit positively by running it on a cluster rather than running it on a local machine.




* Naiad - Chandu Thekkath
  http://research.microsoft.com/Naiad/
  data parallel data flow computation, like dryad, but focused on low latency streaming and cyclic computations. Introduces "timely dataflow" - asynchronous messaging with lightweight coordination. Tries to move from map-reduce to higher level looping and streaming concepts.
