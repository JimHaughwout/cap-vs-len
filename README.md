# cap-vs-len
Exploring cap vs. len in golang to learn more myself. Inspired by this [excellent post](https://stackoverflow.com/questions/41668053/cap-vs-len-of-slice-in-golang) in Stack Overflow.

What is most interesting to me is to see how capacity allocations change on my machine vs. that in the SO post:

When `append_size` is 5, capacity increments by 3:
```
Base size is 3, incrementally appending 5 elements
cap 3, len 0, addr 0xc0000b6020, slice []
cap 3, len 1, addr 0xc0000b6020, slice [0]
cap 3, len 2, addr 0xc0000b6020, slice [0 1]
cap 3, len 3, addr 0xc0000b6020, slice [0 1 2]
cap 6, len 4, addr 0xc0000b8030, slice [0 1 2 3]
cap 6, len 5, addr 0xc0000b8030, slice [0 1 2 3 4]
```

When `append_size` is 10, we increase capacity geometrically, first by 3, then by 6:
```
Base size is 3, incrementally appending 10 elements
cap 3, len 0, addr 0xc0000b6020, slice []
cap 3, len 1, addr 0xc0000b6020, slice [0]
cap 3, len 2, addr 0xc0000b6020, slice [0 1]
cap 3, len 3, addr 0xc0000b6020, slice [0 1 2]
cap 6, len 4, addr 0xc0000b8030, slice [0 1 2 3]
cap 6, len 5, addr 0xc0000b8030, slice [0 1 2 3 4]
cap 6, len 6, addr 0xc0000b8030, slice [0 1 2 3 4 5]
cap 12, len 7, addr 0xc00008c0c0, slice [0 1 2 3 4 5 6]
cap 12, len 8, addr 0xc00008c0c0, slice [0 1 2 3 4 5 6 7]
cap 12, len 9, addr 0xc00008c0c0, slice [0 1 2 3 4 5 6 7 8]
cap 12, len 10, addr 0xc00008c0c0, slice [0 1 2 3 4 5 6 7 8 9]
```

This geometric pattern continues when `append_size` is 20. Capacity increases by 3, then 6, then 12:
```
Base size is 3, incrementally appending 20 elements
cap 3, len 0, addr 0xc0000b6020, slice []
cap 3, len 1, addr 0xc0000b6020, slice [0]
cap 3, len 2, addr 0xc0000b6020, slice [0 1]
cap 3, len 3, addr 0xc0000b6020, slice [0 1 2]
cap 6, len 4, addr 0xc0000b8030, slice [0 1 2 3]
cap 6, len 5, addr 0xc0000b8030, slice [0 1 2 3 4]
cap 6, len 6, addr 0xc0000b8030, slice [0 1 2 3 4 5]
cap 12, len 7, addr 0xc00008c0c0, slice [0 1 2 3 4 5 6]
cap 12, len 8, addr 0xc00008c0c0, slice [0 1 2 3 4 5 6 7]
cap 12, len 9, addr 0xc00008c0c0, slice [0 1 2 3 4 5 6 7 8]
cap 12, len 10, addr 0xc00008c0c0, slice [0 1 2 3 4 5 6 7 8 9]
cap 12, len 11, addr 0xc00008c0c0, slice [0 1 2 3 4 5 6 7 8 9 10]
cap 12, len 12, addr 0xc00008c0c0, slice [0 1 2 3 4 5 6 7 8 9 10 11]
cap 24, len 13, addr 0xc0000ba000, slice [0 1 2 3 4 5 6 7 8 9 10 11 12]
cap 24, len 14, addr 0xc0000ba000, slice [0 1 2 3 4 5 6 7 8 9 10 11 12 13]
cap 24, len 15, addr 0xc0000ba000, slice [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]
cap 24, len 16, addr 0xc0000ba000, slice [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
cap 24, len 17, addr 0xc0000ba000, slice [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
cap 24, len 18, addr 0xc0000ba000, slice [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17]
cap 24, len 19, addr 0xc0000ba000, slice [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18]
cap 24, len 20, addr 0xc0000ba000, slice [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
```

