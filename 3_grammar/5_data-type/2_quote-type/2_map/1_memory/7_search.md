##  search
0. If in expand cap status, first search oldbuckets
1. use hash0 and hashFuc to hash key to get hashcode
2. use **low bit** hash to location bucket
3. use **high bit** hash to location cell of bucket
4. compare hashcode 
5. If key is equal, return elem
6. If can't find, search obucket
7. If can't find in obucket return default value
