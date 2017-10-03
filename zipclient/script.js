var form = document.getElementById("form");

form.addEventListener("submit", (evt) => {  
    evt.preventDefault();
    var name = document.getElementById("nameInput").value;
    console.log(name);
    displayAlloc('http://localhost:4000/Seattle');
    //displayAlloc('http://localhost:4000/memory');
    //displayName('http://localhost:4000/hello?name=' + name);
    

})

function displayAlloc(input) {
    fetch(input) 
    .then(function(response) {
        console.log(response);
        return response.json();
    })
    .then(function(data) {
        var alloc = data['Alloc'];
        showAlloc(alloc);
    });
}

function displayName(name) {
    fetch(name) 
    .then(function(response) {
        console.log(response);
        return response.text();
    })
    .then(function(data) {
        console.log(data);
        var name = document.createElement('p');
        name.textContent = data;
        var head = document.querySelector('h1');
        head.appendChild(name);
    })
}

function showAlloc(data) {
    var alloc = document.createElement('p');
    alloc.textContent = data;
    console.log(data);
    var header = document.querySelector('h1');
    header.appendChild(alloc);
}