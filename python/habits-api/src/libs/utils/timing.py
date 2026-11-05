import functools
import time

def measure_time(func):
    """""
    Decorator to measure execution time
    """
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        start_time = time.perf_counter()
        result = func(*args, **kwargs)
        end_time = time.perf_counter()
        execution_time = (end_time - start_time) * 1000 # ms
        print(f"Function {func.__name__} took {execution_time:.2f} ms")
        return result
    return wrapper