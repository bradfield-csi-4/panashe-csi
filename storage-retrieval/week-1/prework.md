## Reading

Read Chapter 3 (Storage and Retrieval) of _Designing Data-Intensive Applications_. This will give a broad overview of the ideas we'll be discussing in this module.

I know it's 42 pages long, but you're not expected to perfectly understand everything—for example, in previous cohorts it's taken multiple sessions worth of implementation exercises and discussion for LSM-trees to really "click" with students, so don't worry if you haven't gotten it after just reading the chapter.

I've included a copy of the chapter if you don't have a copy of the book yet.

## Implementation

**If you have limited time, please prioritize the reading.** The implementation exercise is meant to be a chance to explore / investigate some of the ideas from the reading, but for the first session it's more important for you to get a broad overview than it is to spend a lot of time getting your implementation working.

We're going to revisit Exercise 4.12 (on page 113) from _The Go Programming Language_, which you might've already done as part of the prep phase for CSI:

The popular web comic _xkcd_ has a JSON interface. For example, a request to [https://xkcd.com/571/info.0.json](https://xkcd.com/571/info.0.json) produces a detailed description of comic 571, one of any favorites. Download each URL (once!) and build an offline index. Write a tool `xkcd` that, using this index, prints the URL and transcript of each comic that matches a search term provided on the command line.

However, this time we'll make the problem a bit more interesting and relevant to the module by adding additional goals:

- **Build an on-disk index** instead of just storing everything in-memory, so that you don't have to rebuild the index every time you restart the tool
- If possible, **make it easy to switch out the data storage layer**, e.g. replacing SQLite with LevelDB
- For fun, consider how well your approach would continue to work if you significantly increase the size of the dataset being indexed

There are many possible choices here. It might be interesting if people make different choices that we can discuss together, and it might also be interesting for one person to try multiple approaches (depending on time/interest):

- Design the index format yourself
- Use an existing system such as:
	- SQLite
	- LevelDB / RocksDB
	- PostgreSQL / MySQL
	- Cassandra
	- Lucene / Solr / Elastic Search
	- Any other system you're familiar with, or want an excuse to try

If you use an existing system instead of designing your own index format, try to track down what actually gets stored on disk, inspect the file using a tool such as `xxd`, and identify which approach from the reading (if any) is being used.