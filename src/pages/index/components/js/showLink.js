import '../css/showLink.css'

//clipboard
import ClipboardJS from 'api/js/clipboard.min.js'

// clipboard
(function(){
    const copyBtn = document.getElementById("copyBtn");
    new ClipboardJS(copyBtn);
})();

// show createLink
(function (){
    const returnBtn = document.getElementsByClassName("returnBtn")[0];
    const showLinkBox =document.getElementsByClassName("showLinkBox")[0];
    const createLinkBox = document.getElementsByClassName("createLinkBox")[0];
    returnBtn.addEventListener("click",function(){
        showLinkBox.style.display = "none";
        createLinkBox.style.display = "block";
    },false)
})();