### pool

java.util.concurrent.Executors 

| 类型                     | 方法                                     | 描述                         |
| ------------------------ | ---------------------------------------- | ---------------------------- |
| ExecutorService          | newFixedThreadPool(int nThreads)         | 固定线程数量的线程池         |
| ExecutorService          | newCachedThreadPool()                    | 缓存线程池，线程数量动态变化 |
| ScheduledExecutorService | newScheduledThreadPool(int corePoolSize) | 定时固定线程数量的线程池     |
| ExecutorService          | newSingleThreadExecutor()                | 单线程线程池                 |

### Executors

| Modifier and Type     | Method                                                       | Description                                                  |
| --------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `boolean`             | `awaitTermination(long timeout,                TimeUnit unit)` | Blocks until all tasks have completed execution after a shutdown request, or the timeout occurs, or the current thread is interrupted, whichever happens first. |
| `<T> List<Future<T>>` | `invokeAll(Collection<? extends Callable<T>> tasks)`         | Executes the given tasks, returning a list of Futures holding their status and results when all complete. |
| `<T> List<Future<T>>` | `invokeAll(Collection<? extends Callable<T>> tasks,         long timeout,         TimeUnit unit)` | Executes the given tasks, returning a list of Futures holding their status and results when all complete or the timeout expires, whichever happens first. |
| `<T> T`               | `invokeAny(Collection<? extends Callable<T>> tasks)`         | Executes the given tasks, returning the result of one that has completed successfully (i.e., without throwing an exception), if any do. |
| `<T> T`               | `invokeAny(Collection<? extends Callable<T>> tasks,         long timeout,         TimeUnit unit)` | Executes the given tasks, returning the result of one that has completed successfully (i.e., without throwing an exception), if any do before the given timeout elapses. |
| `boolean`             | `isShutdown()`                                               | Returns `true` if this executor has been shut down.          |
| `boolean`             | `isTerminated()`                                             | Returns `true` if all tasks have completed following shut down. |
| `void`                | `shutdown()`                                                 | Initiates an orderly shutdown in which previously submitted tasks are executed, but no new tasks will be accepted. |
| `List<Runnable>`      | `shutdownNow()`                                              | Attempts to stop all actively executing tasks, halts the processing of waiting tasks, and returns a list of the tasks that were awaiting execution. |
| `<T> Future<T>`       | `submit(Callable<T> task)`                                   | Submits a value-returning task for execution and returns a Future representing the pending results of the task. |
| `Future<?>`           | `submit(Runnable task)`                                      | Submits a Runnable task for execution and returns a Future representing that task. |
| `<T> Future<T>`       | `submit(Runnable task,      T result)`                       | Submits a Runnable task for execution and returns a Future representing that task. |

