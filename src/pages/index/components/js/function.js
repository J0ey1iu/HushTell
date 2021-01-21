import Dropzone from 'api/dropzone-5.7.0/dist/dropzone.js'
import axios from 'api/js/axios.min.js'
import { options } from '../../../../api/dropzone-5.7.0/dist/dropzone';
import * as URL from '../config.js';

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
        }else{
            settingsForm.classList.add("hidden");
            settingsBtn.innerHTML = `<i class="iconfontIndex">&#xe607;</i>&nbspShow options`;
        }
    },false)
})();

// dropzone settings and submit
(function (){
    // dropzone settings
    Dropzone.options.myAwesomeDropzone = false;
    Dropzone.autoDiscover = false;
    const dropzoneEl = document.getElementById("my-awesome-dropzone");
    let myDropzone = new Dropzone(dropzoneEl, {
        paramName: "file", // The name that will be used to transfer the file
        maxFiles: 1, 
        addRemoveLinks: true,
        autoProcessQueue:false,
        dictDefaultMessage: "Drop files here </br> or </br> <span>Click here</span>",
        dictRemoveFile: "remove",
        dictFallbackMessage: "Your browser does not support drag'n'drop file uploads.",
        dictFallbackText: "Please use the fallback form below to upload your files like in the olden days.",
        dictInvalidFileType: "You can't upload files of this type.",
        dictResponseError: "Server responded with {{statusCode}} code.",
        dictMaxFilesExceeded: "You can not upload any more files.",
        previewTemplate: "<div class=\"dz-preview dz-file-preview\">\n  <div class=\"dz-image\"><img data-dz-thumbnail /></div>\n  <div class=\"dz-details\">\n    <div class=\"dz-size\"><span data-dz-size></span></div>\n    <div class=\"dz-filename\"><span data-dz-name></span></div>\n  </div>\n  <div class=\"dz-progress\"><span class=\"dz-upload\" data-dz-uploadprogress></span></div>\n  <div class=\"dz-error-message\"><span data-dz-errormessage></span></div>\n  <div class=\"dz-success-mark\">\n    <svg width=\"54px\" height=\"54px\" viewBox=\"0 0 54 54\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\">\n      <title>Check</title>\n      <g stroke=\"none\" stroke-width=\"1\" fill=\"none\" fill-rule=\"evenodd\">\n        <path d=\"M23.5,31.8431458 L17.5852419,25.9283877 C16.0248253,24.3679711 13.4910294,24.366835 11.9289322,25.9289322 C10.3700136,27.4878508 10.3665912,30.0234455 11.9283877,31.5852419 L20.4147581,40.0716123 C20.5133999,40.1702541 20.6159315,40.2626649 20.7218615,40.3488435 C22.2835669,41.8725651 24.794234,41.8626202 26.3461564,40.3106978 L43.3106978,23.3461564 C44.8771021,21.7797521 44.8758057,19.2483887 43.3137085,17.6862915 C41.7547899,16.1273729 39.2176035,16.1255422 37.6538436,17.6893022 L23.5,31.8431458 Z M27,53 C41.3594035,53 53,41.3594035 53,27 C53,12.6405965 41.3594035,1 27,1 C12.6405965,1 1,12.6405965 1,27 C1,41.3594035 12.6405965,53 27,53 Z\" stroke-opacity=\"0.198794158\" stroke=\"#747474\" fill-opacity=\"0.816519475\" fill=\"#FFFFFF\"></path>\n      </g>\n    </svg>\n  </div>\n  <div class=\"dz-error-mark\">\n    <svg width=\"54px\" height=\"54px\" viewBox=\"0 0 54 54\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\">\n      <title>Error</title>\n      <g stroke=\"none\" stroke-width=\"1\" fill=\"none\" fill-rule=\"evenodd\">\n        <g stroke=\"#747474\" stroke-opacity=\"0.198794158\" fill=\"#FFFFFF\" fill-opacity=\"0.816519475\">\n          <path d=\"M32.6568542,29 L38.3106978,23.3461564 C39.8771021,21.7797521 39.8758057,19.2483887 38.3137085,17.6862915 C36.7547899,16.1273729 34.2176035,16.1255422 32.6538436,17.6893022 L27,23.3431458 L21.3461564,17.6893022 C19.7823965,16.1255422 17.2452101,16.1273729 15.6862915,17.6862915 C14.1241943,19.2483887 14.1228979,21.7797521 15.6893022,23.3461564 L21.3431458,29 L15.6893022,34.6538436 C14.1228979,36.2202479 14.1241943,38.7516113 15.6862915,40.3137085 C17.2452101,41.8726271 19.7823965,41.8744578 21.3461564,40.3106978 L27,34.6568542 L32.6538436,40.3106978 C34.2176035,41.8744578 36.7547899,41.8726271 38.3137085,40.3137085 C39.8758057,38.7516113 39.8771021,36.2202479 38.3106978,34.6538436 L32.6568542,29 Z M27,53 C41.3594035,53 53,41.3594035 53,27 C53,12.6405965 41.3594035,1 27,1 C12.6405965,1 1,12.6405965 1,27 C1,41.3594035 12.6405965,53 27,53 Z\"></path>\n        </g>\n      </g>\n    </svg>\n  </div>\n</div>",
    });

    // click the btn to submit file and user's options
    const submitBtn = document.getElementsByClassName("submitBtn")[0].getElementsByTagName("button")[0];
    const optionsForm = document.getElementById("optionsForm");
    const originalNote = document.getElementById("originalNote");
    submitBtn.addEventListener("click",function(){
        if (!originalNote.value && !myDropzone.files.length){
            submitBtn.style.background = "red";
            submitBtn.innerHTML = "Pls Add Notes or Files";
            setTimeout(function(){
                submitBtn.style.background = "linear-gradient( 83deg, rgb(67,191,102) 0%, rgb(116,220,123) 100%)";
                submitBtn.innerHTML = "CREATE NOTE";              
            },3000)
        }else if(originalNote.value && myDropzone.files.length){
            submitBtn.style.background = "red";
            submitBtn.innerHTML = "Pls Only Notes or Only File";
            setTimeout(function(){
                submitBtn.style.background = "linear-gradient( 83deg, rgb(67,191,102) 0%, rgb(116,220,123) 100%)";
                submitBtn.innerHTML = "CREATE NOTE";              
            },3000)            
        }else{
            let params = new FormData();
            params.append('mytext', originalNote.value);
            params.append('myfile', myDropzone.files[0]);
            params.append('options', {
                "eamil": optionsForm.email.value,
                "pwd": optionsForm.pwd.value
            })
            console.log(URL["UPLOADURL"]);
            // POST
            axios.post(URL["UPLOADURL"],params,{
                headers: {'Content-Type': 'multipart/form-data'}
                }
            ).then(function (response) {
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            });;
            }
    },false);
})()
