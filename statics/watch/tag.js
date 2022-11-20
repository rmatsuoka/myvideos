
import * as api from "../api.js"

let addTagForm = null;
let tagListElem = null;
let once = true;

window.addEventListener("load", () => {
    addTagForm = document.getElementById('addtagform')
    tagListElem = document.getElementById('taglist')

    buildTagList()

    addTagForm.addEventListener("submit", (e) => {
        e.preventDefault();
        addTag();
    })
})

function basename(p) {
    return p.split(/[\\/]/).pop();
}

function pathbase() {
    return basename(window.location.pathname)
}

function buildTagList() {
    api.tagsWithVideo(pathbase())
        .then(json => {
            for (let t of json) {
                appendTagList(t)
            }
        })
        .catch(err => console.error(err))
}

function addTag() {
    const e = addTagForm.elements;
    api.addTag(pathbase(), e["tagname"].value)
        .then(json => {
            appendTagList(json)
        })
        .catch(err => window.alert(err))
}

function deleteTag(elem, tagID) {
    api.deleteTag(pathbase(), tagID)
        .then(text => elem.remove())
        .catch(err => window.alert(err))
}

function appendTagList(tag) {
    const li = document.createElement("li");
    li.setAttribute("class", "tag");

    const a = document.createElement("a")
    a.setAttribute("href", "/tags/" + tag.ID);
    a.innerText = `#${tag.Name} (${tag.N})`;
    li.appendChild(a);

    const button = document.createElement("button");
    button.innerText = "delete";
    button.addEventListener("click", (e) => {
        e.preventDefault();
        deleteTag(li, tag.ID);
    })
    li.appendChild(button);

    tagListElem.appendChild(li);
}
