Section 4 Pointers

Pointers store value addresses instead of values.

Lessons 63 - 72

Use the & sign before the variable to get its address, when passing it to the function.

When defining the type of a pointer for a function parameter, use the * sign before the variable name.
Also use the * when you want to get the value not the address of the variable. It is called dereferencing.

When passing it to a function, use the & sign before the variable to pass its address or pass a pointerType variable directly without the & sign.

Watch out for unexpected behavior when using pointers with functions.
Like add(&x,z) does not return the combined value, instead it adds the second value and adds it to the first variable modifying it.

Some functions like fmt.Scan expects a pointer, dereferences it internally.

Advantages:
- no unnecessary value copies
- can directly mutate the value
