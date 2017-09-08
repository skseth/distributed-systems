# What Every Programmer Should Know About Memory

http://futuretech.blinkenlights.nl/misc/cpumemory.pdf

## Terminology

FSB - Front Side Bus
Northbridge
Memory Controller
RAM
DRAM
SDRAM
DDR2
FB-DRAM
Southbridge
PCI
PCIE
SATA
USB
DMA
CSI - Common System Interconnect
Nehalem
Bisection Network
NUMA - non-uniform memory architecture
1 GBit = 10^9 bits or nearly 2^30 bits

## CPU - Memory Interaction

Shared memory architecture broadly fall into 2 types :

Uniform Memory Access (UMA) : Here all CPUs interact with entire memory uniformly. Usually a "Front-Side Bus" is used to provide this common access.

Non-Uniform Memory Access (NUMA) : CPUs are assigned individual chunks of memory, which they can access independently. An "interconnect" between the CPUs allows access to another CPUs memory with some delay - this means the time to access memory is not uniform.

## SRAM vs DRAM  

SRAM - Static RAM - is expensive, but very fast memory used for in-processor caches
DRAM - Dynamic RAM - is much cheaper, but slow

Dynamic RAM stores charge in a capacitor and loses some charge on each read. Thus it needs to be continuously refreshed, and also charging / discharging a capacitor is slower.

But DRAM chips need much less real estate for same amount of storage, and has a simpler design which results in much lower costs. Most RAM in general purpose computers uses dynamic RAM

## Memory Addressability - Synchronous

Synchronous memory access refers to the fact that memory access is driven by the clock of the FSB - the memory controller puts out an address on the bus at specific point by clock, and the response is provided after a known number of cycles.

DRAM is typically organized in rows and columns. To get a memory address, the memory controller sets a Row Address Selector (RAS) and a Column Address Selector (CAS), triggering the memory read or write.

Since the entire 64-bit memory is not stored in a single chip, this is carried out in parallel in multiple DRAM chips to read or write an entire "word" of memory.

The reason for this way of addressing DRAM is the size of the data stored. With 1 GBit chip one would need 30 bits of address and 2^30 select lines to take away the output - which is not practical. By splitting the address into row and column addresses, and by allowing only one value in a column to be read at a time, the number of lines required reduce tremendously.

- **Understanding System Memory and CPU Speeds**  
  https://web.archive.org/web/20160118134028/http://www.directron.com:80/fsbguide.html#part3

## Types of DRAM

Typically, the clock speed of the FSB would determine the rate at which a single memory read or write can happen. However, processors have tried to improve the situation by being able to read upto twice or even 4 times per cycle (double or quad-pumped memory), which allows an effective increase in speed at which RAM is accessed.

Transfer Rate = size of word * pumping * FSB-frequency e.g. for a quad-pumped 64-bit bus at 200 MHz, the transfer rate is 8 * 4 * 200 MHz = 6.4 GB / sec. But this is only the maximum achievable transfer rate.

SDRAM - Synchronous DRAM - responds on only rising edge of clock signal
DDR RAM - Double Data Rate RAM - responds on rising and falling edge. We have DDR, DDR2, DDR3 and DDR4 RAM

DDR - while transferring data, will transmit 2 words per cycle instead of one

- **Secrets of PC Memory**  
  http://www.bit-tech.net/reviews/tech/memory/the_secrets_of_pc_memory_part_1/1/

## Read Access

CL - CAS latency
tRCD - Row-Command or Row-Column Delay
tRP - RAS Precharge
tRAS - RAS Active Time



- **Everything You wanted to know about SDRAM**  
  http://www.anandtech.com/show/3851/everything-you-always-wanted-to-know-about-sdram-memory-but-were-afraid-to-ask


## DIMMs - Dual In-line Memory Modules

- introduced with Pentium P5
- pins on both sides with 64-bit bus width


Registered / Unregistered
Ranking 
Error Correction
Frequency


CAS Latency


- **Wikipedia article on DIMMs**  
  https://en.wikipedia.org/wiki/DIMM
