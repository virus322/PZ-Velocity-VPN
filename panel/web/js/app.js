const API = "http://localhost:1196";


function loadStatus(){


fetch(API + "/api/status")

.then(response => response.json())

.then(data => {


document.getElementById("status")
.innerHTML =
data.status;


});


}



function loadUsers(){


fetch(API + "/api/users")

.then(response => response.json())

.then(data => {


document.getElementById("users")
.innerHTML =
data.length;


});


}



loadStatus();

loadUsers();
