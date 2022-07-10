this project dont use DB, please add some data using API POST before use API GET

Note :
New User [POST]
http://localhost:8080/newuser
Request : id[string],name[string],password[string],retypepassword[string]
response : result [data], status [just message]
---------------------------------------------------------------------------------------

NOTE : before use "Login User", please use "New user" first
Login User [POST]
http://localhost:8080/login
Request : id[string],password[string]
response : Token,id,name,password
---------------------------------------------------------------------------------------

List Product [GET]
http://localhost:8080/product
Request : 
response : result [data]
---------------------------------------------------------------------------------------

Order [POST]
http://localhost:8080/newuser
Request : username[string],id[int],nameitem[string],price[int]
response : result [data], status [just message]
---------------------------------------------------------------------------------------

NOTE : before use "List Order", please use "Order" first
List Order [GET]
http://localhost:8080/order/{iduser}
Request : 
response : result [data]