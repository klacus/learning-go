Error handling
Lesson 49 - 50

Go does not have exception...s. Instead, it uses a multi-value return to indicate an error. The first return value is the result, and the second return value is an error. If the error is nil, then the result is valid. If the error is not nil, then the result is invalid.
