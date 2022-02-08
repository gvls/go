##  bmap
a bucket for a go map
```shell
bmap		 _____________________________________
			 |        |        |        |        |
head node	 | bucket | bucket | bucket | bucket |	
			 |        |        |        |        |
			 -------------------------------------
				 |        |        |        |
				 v        v        v        v
link table	 o_bucket  o_bucket o_bucket 0_bucket	
				 |        |        |        |
				 v        v        v        v
		   	    ...      ...      ...      ...

when have hash冲突 in bucket, obucket will be create and obucket save new key-value
when obucket is overflow, next obucket will be create
```
