
const PageId = {
	Home: 1, //"Home",
	Product: 2, //"Products and solutions",
	Subscriptions: 3, //"Subscription manager",
	About: 4, //"About us",
}

const numOfPages = Object.keys(PageId).length;
const pageClassName = "page"; // Make sure it's the same in index.html 
const pageIdPrefix = "page"; // Make sure it's the same in index.html

const defaultPageId = PageId.Home;
const namePageId = 'pageId';


function setPageId(pageId){
    localStorage.setItem(namePageId, pageId);
    updateView();
}

function getPageId(){
    return parseInt(localStorage.getItem(namePageId));
}

function resetPage(){
    setPageId(defaultPageId);
    hidePages();
    showPage();
}

function hidePages(){
    const elements = document.getElementsByClassName(pageClassName);

    for (let element of elements) {
        element.style.display = 'none';
    }
}

function showPage(pageId){
    const element = document.getElementById(pageIdPrefix+pageId);
    element.style.display = 'block';
}

function init(){
    if(getPageId()==null){
        setPageId(defaultPageId);
    }

    updateView();
}

function updateView(){
    updateViewHome(pageIdPrefix+PageId.Home);
}

document.addEventListener("DOMContentLoaded", function () {
    // Code inside this function will execute when the DOM is fully loaded
    init(); // Call the init function
});