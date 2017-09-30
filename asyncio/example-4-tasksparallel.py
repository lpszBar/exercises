import asyncio


async def my_coroutine(task_name, seconds_to_sleep):
    print('{0} sleeping for: {1} seconds'.format(task_name, seconds_to_sleep))
    await asyncio.sleep(seconds_to_sleep)
    print('{0} is finished'.format(task_name))


loop = asyncio.get_event_loop()
tasks = [
    my_coroutine('task1', 4),
    my_coroutine('task2', 3),
    my_coroutine('task3', 2)]
print("------!-------")
loop.run_until_complete(asyncio.wait(tasks))
print("------^-------")
loop.close()