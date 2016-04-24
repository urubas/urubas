# Type System

## Basic Scalar Values

* bool
* int8
* int16
* int32
* int64
* uint8
* uint16
* uint32
* uint64
* float
* double
* string

## Arrays

Example notation:
	x [length]string

Arrays have fixed length.

## Slices
    x []string

Example notation:
    

## Structs

Example notation:
    x struct {
        field string
    }

## Channels

Analogous to golang channels.

Example notation:
    x chan string


### Casting

Given a channel of any type, a scalar value of same type can always be casted to the equivalent channel type whom this value will be the single input.

Also, any array can be casted to an array of it's equivalent channel type whom all items of the array will serve as input for the channel.
