# golang_database
golang homework5

#### Task

Implement a http based db service(storing user) by using "database/sql" or "gorm" library and do ab test

#### Result

* use postman to test the api

  1. http://localhost/service/userinfo?username=weimumu&password=weimumu POST ![post](/images/adduser.jpg)
  2. http://localhost/service/userinfo?userid=2 GET![get](/images/getuser.jpg)
  3. http://localhost/getallusers GET ![get](/images/getallusers.jpg)

* Use ab to test performance

  1. test gorm/mysql ![1](/images/1.png) ![2](/images/2.png)

  2. test database/sql ![3](/images/3.png)

     ![4](/images/4.png)

* We can see the "Time per request" of gorm is smaller than that of the database/sql, so it is said the perfomance of the gorm/mysql is better than database/sql

#### Thanks

