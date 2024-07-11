An interesting property of arrays is that the size is encoded in its type. 
If you try to pass an [4]int into a function that expects [5]int, it won't compile.

They are different types so it's just the same as trying to pass a string into a function that wants an int.
