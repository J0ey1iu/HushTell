// note function
(function() {
    const noteFunction = document.getElementsByClassName("functionBox")[0].getElementsByClassName("function")[0];
    noteFunction.addEventListener("focus",function() {
        this.placeholder = "";
    },false)
    noteFunction.addEventListener("blur",function() {
        this.placeholder = "Write your note here...";
    },false)
})();

// setting options
(function() {
    const settingsBtn = document.getElementsByClassName("settingsBtn")[0].getElementsByTagName("button")[0];
    const settingsForm = document.getElementsByClassName("settingsForm")[0];
    settingsBtn.addEventListener("click",function(){
        if (settingsForm.classList.contains("hidden")){
            settingsForm.classList.remove("hidden");
            settingsBtn.innerHTML = `<i class="iconfontIndex">&#xe607;</i>&nbspHidden options`;
            console.log("1");
        }else{
            settingsForm.classList.add("hidden");
            settingsBtn.innerHTML = `<i class="iconfontIndex">&#xe607;</i>&nbspShow options`;
            console.log("2");
        }
    },false)
})();

// 