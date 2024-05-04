function setHomeList(){
    let homeTestListDiv = document.getElementById("home-test-list");

    homeTestListDiv.innerHTML = "<li>Item 1</li>";
    homeTestListDiv.innerHTML += "<li>Item 2</li>";
}

function setHome(homeInsertPlace){
    let homeDiv = document.getElementById(homeInsertPlace);
    homeDiv.innerHTML = createHome();
}

function updateViewHome(homeInsertPlace){
    setHome(homeInsertPlace); 
    setHomeList();
}