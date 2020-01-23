# Mutex and Channel basics

### What is an atomic operation?
> An atomic operation is an operation that is completely independent from other processes. It happens "instantly" from the viewpoint of the other processes

### What is a semaphore?
> A shared variable signaling access to a resource

### What is a mutex?
> A lock that only one thread can hold at a time, only the thread holding the lock can interract with stuff.

### What is the difference between a mutex and a binary semaphore?
> Mutexes are opened/closed by the thread itself, a semaphore can be opened by other threads

### What is a critical section?
> The section of code that needs to be protected from concurrent access from several threads

### What is the difference between race conditions and data races?
 > Race conditions are situations where the end result depends on the unpredictable timing of events, (eg. current propogation)

 > A data race is when two or more threads try to access the same data at the same time.

### List some advantages of using message passing over lock-based synchronization primitives.
> Easier to implement, no conflicts that need to be avoided.

### List some advantages of using lock-based synchronization primitives over message passing.
> Allows for higher speed as it doesn't require kernel intervention to queue things.
