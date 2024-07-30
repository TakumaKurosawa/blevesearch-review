# blevesearch-review

## Search result

```
-------- keyword: a1001 ----------
2024/07/30 15:33:53 user1: &searchstore.User{EmpNum:"A1001", FullName:"黒澤\u3000あきら", FullNameKana:"ｸﾛｻﾜ\u3000ｱｷﾗ", Email:"kurosawa_akira@example.com", Score:0.27563438421101444}
-------- keyword: a1001 ----------

-------- keyword: kurosawa ----------
2024/07/30 15:33:53 user1: &searchstore.User{EmpNum:"A1000", FullName:"黒澤\u3000拓磨", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾏ", Email:"kurosawa_takuma@example.com", Score:0.05409637803962386}
2024/07/30 15:33:53 user2: &searchstore.User{EmpNum:"A1003", FullName:"黒澤\u3000拓実", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾐ", Email:"kurosawa_takumi@example.com", Score:0.05409637803962386}
2024/07/30 15:33:53 user3: &searchstore.User{EmpNum:"A1001", FullName:"黒澤\u3000あきら", FullNameKana:"ｸﾛｻﾜ\u3000ｱｷﾗ", Email:"kurosawa_akira@example.com", Score:0.05304586197404865}
-------- keyword: kurosawa ----------

-------- keyword: kurosawa_takuma ----------
2024/07/30 15:33:53 user1: &searchstore.User{EmpNum:"A1000", FullName:"黒澤\u3000拓磨", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾏ", Email:"kurosawa_takuma@example.com", Score:0.2810930258102453}
-------- keyword: kurosawa_takuma ----------

-------- string query keyword: ｸﾛｻﾜ ----------
2024/07/30 15:33:53 user1: &searchstore.User{EmpNum:"A1000", FullName:"黒澤\u3000拓磨", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾏ", Email:"kurosawa_takuma@example.com", Score:0.14246358763251668}
2024/07/30 15:33:53 user2: &searchstore.User{EmpNum:"A1003", FullName:"黒澤\u3000拓実", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾐ", Email:"kurosawa_takumi@example.com", Score:0.14246358763251668}
2024/07/30 15:33:53 user3: &searchstore.User{EmpNum:"A1001", FullName:"黒澤\u3000あきら", FullNameKana:"ｸﾛｻﾜ\u3000ｱｷﾗ", Email:"kurosawa_akira@example.com", Score:0.13969703850314946}
-------- string query keyword: ｸﾛｻﾜ ----------

-------- match phrase query keyword: ｸﾛｻﾜ ----------
2024/07/30 15:33:53 user1: &searchstore.User{EmpNum:"A1000", FullName:"黒澤\u3000拓磨", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾏ", Email:"kurosawa_takuma@example.com", Score:0.14246358763251668}
2024/07/30 15:33:53 user2: &searchstore.User{EmpNum:"A1003", FullName:"黒澤\u3000拓実", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾐ", Email:"kurosawa_takumi@example.com", Score:0.14246358763251668}
2024/07/30 15:33:53 user3: &searchstore.User{EmpNum:"A1001", FullName:"黒澤\u3000あきら", FullNameKana:"ｸﾛｻﾜ\u3000ｱｷﾗ", Email:"kurosawa_akira@example.com", Score:0.13969703850314946}
-------- match phrase query keyword: ｸﾛｻﾜ ----------

-------- string query keyword: ｸﾛｻﾜ ﾀｸﾏ ----------
2024/07/30 15:33:53 user1: &searchstore.User{EmpNum:"A1000", FullName:"黒澤\u3000拓磨", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾏ", Email:"kurosawa_takuma@example.com", Score:0.31513356368417345}
2024/07/30 15:33:53 user2: &searchstore.User{EmpNum:"A1003", FullName:"黒澤\u3000拓実", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾐ", Email:"kurosawa_takumi@example.com", Score:0.0322020186676597}
2024/07/30 15:33:53 user3: &searchstore.User{EmpNum:"A1001", FullName:"黒澤\u3000あきら", FullNameKana:"ｸﾛｻﾜ\u3000ｱｷﾗ", Email:"kurosawa_akira@example.com", Score:0.031576676654382005}
-------- string query keyword: ｸﾛｻﾜ ﾀｸﾏ ----------

-------- match phrase query keyword: ｸﾛｻﾜ ﾀｸﾏ ----------
2024/07/30 15:33:53 user1: &searchstore.User{EmpNum:"A1000", FullName:"黒澤\u3000拓磨", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾏ", Email:"kurosawa_takuma@example.com", Score:0.31513356368417345}
-------- match phrase query keyword: ｸﾛｻﾜ ﾀｸﾏ ----------

-------- string query keyword: 拓磨 ----------
2024/07/30 15:33:53 user1: &searchstore.User{EmpNum:"A1000", FullName:"黒澤\u3000拓磨", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾏ", Email:"kurosawa_takuma@example.com", Score:0.3449830290771593}
2024/07/30 15:33:53 user2: &searchstore.User{EmpNum:"A1003", FullName:"黒澤\u3000拓実", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾐ", Email:"kurosawa_takumi@example.com", Score:0.05797386801764451}
-------- string query keyword: 拓磨 ----------

-------- match phrase query keyword: 拓磨 ----------
2024/07/30 15:33:53 user1: &searchstore.User{EmpNum:"A1000", FullName:"黒澤\u3000拓磨", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾏ", Email:"kurosawa_takuma@example.com", Score:0.3449830290771593}
-------- match phrase query keyword: 拓磨 ----------

-------- string query keyword: 黒澤　拓磨 ----------
2024/07/30 15:33:53 user1: &searchstore.User{EmpNum:"A1000", FullName:"黒澤\u3000拓磨", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾏ", Email:"kurosawa_takuma@example.com", Score:0.39950599238748297}
2024/07/30 15:33:53 user2: &searchstore.User{EmpNum:"A1003", FullName:"黒澤\u3000拓実", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾐ", Email:"kurosawa_takumi@example.com", Score:0.15129638290165254}
2024/07/30 15:33:53 user3: &searchstore.User{EmpNum:"A1001", FullName:"黒澤\u3000あきら", FullNameKana:"ｸﾛｻﾜ\u3000ｱｷﾗ", Email:"kurosawa_akira@example.com", Score:0.049815876772866224}
-------- string query keyword: 黒澤　拓磨 ----------

-------- match phrase query keyword: 黒澤　拓磨 ----------
2024/07/30 15:33:53 user1: &searchstore.User{EmpNum:"A1000", FullName:"黒澤\u3000拓磨", FullNameKana:"ｸﾛｻﾜ\u3000ﾀｸﾏ", Email:"kurosawa_takuma@example.com", Score:0.39950599238748297}
-------- match phrase query keyword: 黒澤　拓磨 ----------
```
