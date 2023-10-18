# concurrency_python_golang
Asyncio in Python vs GoRoutines concurrency speed test

I have two similar methods added in Python and Golang. Tried to simulate 1000s of concurrent task execution and measure their performance. 

Test results:
1000 tasks
Python with Asyncio took 2.02 seconds
golang took 2.00 seconds

10M tasks
Python with Asyncio took 148 seconds.
golang took 16 seconds. 
