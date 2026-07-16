function createUser(username){


fetch(
"http://localhost:1196/api/users/create",
{

method:"POST",

headers:{
"Content-Type":"application/json"
},

body:JSON.stringify({

username:username

})

}

)

.then(
response=>response.json()
)

.then(
data=>{

console.log(data);

}

);


}
