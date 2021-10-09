An Instagram API developed using GO-lang and mongoDB. API hosted locally on URL
```
localhost:12345/
```

**Features**

- [x] **Create new User**
```
localhost:12345/users/
```
Body: JSON Format - { "Name":"Roy", "Email":"roy@abc.com", "Password":"xyz"}

![Create User (Postman)](https://github.com/Swakshath/InstagramAPI/blob/0b14213b1f176adc21bd51ac890cb7237b61138e/CreateUser%20(Postman).JPG)

Note: Password hashed before storing in database

<br/><br/>
- [x] **Get a user using id**
```
localhost:12345/users/61617b7e883ab32afaa133f0
```
![Get User Using ID (Postman)](https://github.com/Swakshath/InstagramAPI/blob/826c47bf597a0c6f5addebfa755d159ebaeb20cb/GetUserUsingID%20(Postman).JPG)


<br/><br/>
- [x] **Create a Post**
```
localhost:12345/posts/
```

Body: JSON Format - {
    "Caption":"My Instagram Image",
    "ImageURL":"randomurl.com",
    "UserID":"61617b7e883ab32afaa133f0"
}

![Create Post (Postman)](https://github.com/Swakshath/InstagramAPI/blob/dca92934e4901ecb410136f4527d8bd8518f1429/CreatePost%20(Postman).JPG)
<br/><br/>
- [x] **Get a post using id**
```
localhost:12345/posts/61618017883ab32afaa133f3
```
![Get Post Using ID (Postman)](https://github.com/Swakshath/InstagramAPI/blob/e6754405238539d57ac1d7e600900bed41700c5b/GetPostUsingID%20(Postman).JPG)

<br/><br/>
- [x] **List all posts of a user**
```
localhost:12345/posts/users/61617b7e883ab32afaa133f0?page=1
```
![List All Posts of User (Postman)](https://github.com/Swakshath/InstagramAPI/blob/3c3b275d34d1ffe06f9560253b0f0bb71dfa7799/ListAllPosts%20(Postman).JPG)

Note: Two posts will be displayed per page (Pagination)

<br/><br/>
Refer Main_test.go for unit tests

```
go test
```

![Unit Tests](https://github.com/Swakshath/InstagramAPI/blob/cfbf2b50a169efea76d9824a474b8024121f239f/UnitTests.JPG)




