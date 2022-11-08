import concurrent.futures
#import multiprocessing
from multiprocessing.dummy import freeze_support
import time


def do_something(seconds):
    print(f'Sleeping {seconds} second(s)...')
    time.sleep(seconds)
    # print('Done sleeping...')
    return f'Done sleeping {seconds}...'


if __name__ == '__main__':

    freeze_support()
    start = time.perf_counter()

    with concurrent.futures.ProcessPoolExecutor() as executor:
        secs = [5, 4, 3, 2, 1]
        
        # f1 = executor.submit(do_something, 1)
        # print(f1.result())
        
        # random finish, in this case backwards since faster finish earlier
        # results = [executor.submit(do_something, s) for s in secs]
        # for f in concurrent.futures.as_completed(results):
        #     print(f.result())
        
        # with map, same finish order as sent sort of FIFO 
        results = executor.map(do_something, secs)
        for result in results:
            print(result)
       
       

    end = time.perf_counter()
    print(f'Time taken : {end - start}')
    # old way
    # processes = []

    # for _ in range(10):
    #     p = multiprocessing.Process(target=do_something, args=[1.5])
    #     p.start()
    #     processes.append(p)

    # for process in processes:
    #     p.join()
