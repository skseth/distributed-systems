# HPCA - Georgia Tech

## Why architecture

Improve performance - speed, battery life, energy efficiency, size, weight, cost
Improve abilities - graphics, debugging support, security

## Moore's Law

Every 18-24 months, we will fit twice as many transistors on same chip area

=> if possible, double speed, halve energy / operation, memory capacity 2x every 18-24 month

## Memory Wall

Instructions / second => doubles every 2 years
Memory size => doubling every 2 years
Memory Latency => 1.1x every 2 years => solution - caches

## Speed, Power, Weight, Cost

Tradeoffs have to be made based on use case.

- embedded device
- mobile
- laptop
- server

Active Power - 0.5 C . V^2 . f . alpha (activity factor)

Voltage reduction reduces active power consumption, but increases static power consumption due to leakage
Higher frequency may need higher voltage, significantly increasing power consumption

Fabrication cost is higher for larger chips, and yield is lower, further increasing cost.

## Performance - Throughput and Latency

Throughput = 1/ Latency if there is just one step. So if it takes 4 hours to make a car, we can make 1/4 car per hour.

But if we have 20 steps, with a latency of 4 hours, say, and once we are done with one step we can pick up the next car for that step, then we can achieve a throughput of 5 cars / hour!

Speedup => increase in throughput, or decrease in latency

Benchmarks - Synthetic Benchmarks : 
TPC
EEMBC

SPEC
  gcc
  bwaves
  llvm
  perl
  cactus adm
  calculix

Different applications may have different speedups. To average, we use geometric mean

## Iron Law of Performance

CPU Time => (# of intructions / second) x (cycles / instruction) x (clock cycle time)

Levers to improve time :

No of intructions / second => better algorithms or compiler, instruction set design
No of cycles / instruction => instruction set design, processor design
clock cycle time => processor design, circuit design etc

## Amdahl's Law

orig time (for some process) = t

support for some part which takes F % of the time, we are able to speedup by a factor of s. How much is the overall speedup S ?

new time = t[(1 - F) + F/s]

S = orig time / new time = 1 / [(1 - F) + F/s]

If s = infinity,

S = 1 / (1 - F) = maximum possible overall speedup.

Thus, if we improved performance for 25% of the execution (by time), max overall speedup cannot ever exceed 1 / .75 = 1.33 or 33% speedup, even if we cut this 25% to zero. Thus small speedup of a large part of the execution may have a larger impact, then a large speedup on a small part of the execution time => Make common case fast.

But conversely, a small speedup of the common case which results in a large slowdown of an uncommon case may wipe out all gains => Don't slowdown the uncommon case too much.

Finally, there is the law of diminishing returns. Suppose you find a way to speed up some sub-process x which takes half the overall time, by 2 times. This will give an overall 33% speedup. A further speedup of 2 for x will result in only a 20% speedup, and so on. What's more it is likely that the second speedup is costlier than the first - which means we are paying more for smaller returns => Don't over-optimize one part of the execution, as you get diminishing returns.

## Pipelining

Pipelining can significantly increase throughput in a processor. If we take a typical processor sequence of : Fetch -> Decode -> Execute -> Memory Write -> Wr Register

If the overall sequence takes 5 cycles and we process instructions sequentially, we will take 5 cycles per instruction. Suppose we can execute instructions in a pipeline, we can start approaching nearly one cycle per instruction - an increase in throughput of 5 times!

### Pipeline Hazards - Flushes, Stalls and Forwards

There are some reasons why we cannot always achieve a CPI of 1 :

- Initial fill time (usually negligible in steady state)
- Pipeline stalls caused by branches (control dependence), or dependencies between instructions (data dependence)

When branches are taken, deeper the pipeline, higher the penalty. If we have just 25% of branches taken, and a branch is decided in the sixth stage, we can move from a CPI of 1 to (1 + 0.25*5) = 2.25!
Control dependencies may result in flushing the pipeline.

Data Dependence can be of many types : RAW - Read after write, WAR, WAW, RAR.

WAR, WAW - false dependencies or name dependencies
RAW - true dependency
RAR - not a dependency

RAW can result in a pipelined instruction getting a stale value. However, even for RAW dependencies, if there is are enough instructions between the 2 instructions, the dependency may not create any problem.

A pipeline hazard is a case where we have to stall or do other optimization to the pipeline to avoid avoid issues caused by data dependencies. But unlike control dependencies, we do not have to flush the pipeline. One kind of optimization we can do is to forward a needed value from a future stage in a pipeline to process an earlier step. Finally we can also consider doing out-of-order processing to reduce impact of stalls.

### Pipeline Stages

Increases CPI due to increase impact of hazards. Decrease cycle time as each stage does less, so we can speed up the clock. 

Execution time = #instructions x CPI x cycle time. 

For modern processors, optimal performance can allow 30-40 stages. But increased cycle time also increases power - so in actual processors we are usually at around 10-15 stages.

e.g. Core Series II (i3,i5,i7)- 14 stages

- **The Optimum Pipeline Depth for a Microprocessor**  
  https://www.researchgate.net/profile/Thomas_Puzak/publication/220771139_The_Optimum_Pipeline_Depth_for_a_Microprocessor/links/558ca1c008ae1f30aa80aede/The-Optimum-Pipeline-Depth-for-a-Microprocessor.pdf

### Prediction 

CPI = 1 + mispreds/instruction * penalty per mispred

penalty per mispred increases with deeper pipelines => improving branch prediction becomes more important

penalty of misprediction is also higher for superscalar processors which process multiple instructions per cycle

predicting always wins over not predicting => like a goalkeeper who chooses a side

predict not-taken predictor still does ok

1-bit predictor : branch-target-buffer - store history of pcnew given pcnow, storing address when taken. Can work very well for loops with more iterations.

2-bit predictor, or 2-bit counter. bit 1 - prediction bit, bit 0 - hyterisis bit or conviction bit

pshare predictor : private history, shared history

gshare : global history, shared history

tournament predictor - use both pshare, gshare + 2 bc to break tie

hierarchical predictor 

RAS predictor - return address stack, predecoding in cache to know a return or branch

### Predication

Prediction is good for loops & returns

conditional move removes hard to predict branches, needs compiler support

full predication - eg itanium - make every instruction conditional

## ILP & IPC - Instruction Level Parallelization & Instructions per Cycle

ILP - property of program. A program's ILP is the inherent amount of instructions per cycle that can be executed on an ideal processor that have no physical limit of parallelizability and which does perfect branch prediction and elimination of named dependencies (WAW, WAR) by (for eg) renaming registers.

IPC - property of a processor, for a program. It is the actual instructions per cycle that the processor actually achieves for a program. IPC can be improved by :
- Control dependencies => Better branch prediction 
- WAR and WAW dependencies => using register renaming
- Handle RAW depencies => out of order execution
- Structural dependencies => wider issue - more hardware basically

IPC <= ILP

In general, for narrow-issue processors (i.e processors which can execute fewer, say less than 3 instructions in parallel), out-of-order execution is less of a problem. But for wide-issue processors out-of-order execution is imperative and has to be supported.

### Tomasulo's Algorithm

Developed for IBM 360/91 in 1960s. Largely survives even today. Uses register renaming, determines which instructions have inputs ready.

Modern processors have extended this :
- extended from floating point to all instructions
- current processors may look at "window" of hundreds of instructions ahead
- exception handling was less of an issue, unlike today
- may reorder loads and stores also

Issue : Instruction moved from instruction queue to reservation station
Dispatch : Instruction sent to execution unit
Write Result / Broadcast : Post-execution broadcast of executed value of instruction

Tomasulo's algorithm issues in-order, may dispatch out-of-order and write to registers out-of-order. Loads and stores are always in-order. A RAT (register alias table) is used to perform register renaming.

Tomasulo's algorithm does not handle exceptions and branch mispredictions. In order to handle this modern processors incorporate a Reorder Buffer (ROB) to ensure actual writes to registers happens in-order.

- **Tomasulo's Algorithm - Georgia Tech**  
  https://www.cc.gatech.edu/~milos/Teaching/CS6290F07/4_Tomasulo.pdf  
