import * as api from '../api.js'

function updateDescription() {
    const descrption = ""
    api.updateVideoInfo({descrption})
        .catch(err => window.alert(err))
}

class Title {
    constructor() {
        this.titleElem = document.querySelector("#title");
        this.buttonElem = document.querySelector("#updatetitle");
        this.buttonElem.addEventListener("click", (e) =>{
            e.preventDefault()
            this.update()
        })
    }
    update() {
        const newTitle = window.prompt("update the title", this.titleElem.innerText)
        api.updateVideoInfo(api.pathbase(), {title: newTitle})
            .then(() => {this.titleElem.innerText = newTitle})
            .catch(err => window.alert(err))
    }
}

new Title();