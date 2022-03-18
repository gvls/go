##  insert and update
1. use hash0 and hashFuc to hash key to get hashcode
2. use **low bit** hash to location bucket
3. use **high bit** hash to location cell of bucket
4. If key exist, do update. If key not exist and have free space, do insert
5. If key not exitst and not have free space, judge if need expand cap
6. If expand cap, do pre stey again
7. If not expand cap, create obucket


