import '../css/showLink.css'

//clipboard
import ClipboardJS from 'api/js/clipboard.min.js'

// clipboard
(function(){
    const copyBtn = document.getElementById("copyBtn");
    new ClipboardJS(copyBtn);
})();