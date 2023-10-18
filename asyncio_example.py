import asyncio
import timeit

# Asynchronous function that simulates a task
async def task(number):
    await asyncio.sleep(2)  # Simulate a time-consuming task (2 seconds)
    return f"Task {number} result"

# Asynchronous function to run 1000 tasks concurrently
async def main():
    tasks = []
    # Create and run 1000 tasks concurrently
    for i in range(1, 10000001):
        task_instance = asyncio.create_task(task(i))
        tasks.append(task_instance)

    # Wait for all tasks to complete concurrently
    results = await asyncio.gather(*tasks)

    print("All tasks completed.")

# Run the event loop to execute the asynchronous tasks
if __name__ == "__main__":
    execution_time = timeit.timeit(lambda: asyncio.run(main()), number=1)
    print(f"Execution time: {execution_time:.2f} seconds")
