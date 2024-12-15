fetch("http://localhost:8000/movies",{
    method:"GET",
    headers:{
        'Content-Type': 'application/json'
    }
})
.then(response=>{
    return response.json()
}).then(data=>{
    console.log(data)
})