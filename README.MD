Tag v1.6.xx
===========

Handling Error in golang
------------------------
1. Panic 
    - shutdowning program when called/used
    - use panic when program need to shutdown if some procces fail.
2. Log
    - Standart eror Logging
        * log.print
        * log.fatal
        * log.panic
3. Fmt 
    - fmt package printing error use Errorf
4. special code based on the error
    - http status code

-------
* func Is(err, target error) bool
    - comparing error to other error var / new costum error

* Wrapping erors using new struct contain detail of eror happend
    - best practice is add method Unwrap() error, then caller can check the type or error are returned by specify package by comparing using Is()

* Or creating detailed wrapped error massage in method side using fmt.Erorrf() error, then method caller can use .erorr() string

* func As(err, target any) bool
    - comparing error to other interface erorr. AS are unwraping and search a error
