fetch("/api/status")

.then(
response=>response.json()
)

.then(
data=>{

document.getElementById("status")
.innerHTML=data.status;


}
);