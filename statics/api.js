
export function basename(p) {
    return p.split(/[\\/]/).pop();
}

export function pathbase() {
    return basename(window.location.pathname)
}

const addTagURL = "/api/addtag"
export async function addTag(path, name) {
    const response = await mustFetch(addTagURL, {
        method: "POST",
        headers: new Headers({"Content-Type": "application/json"}),
        body: JSON.stringify({videoID: path, tagname: name})
    })
    let newtag = await response.json()
    return newtag
}

async function mustFetch(url, request) {
    const response = await fetch(url, request)
    if (!response.ok) {
        throw new Error(await response.text())
    }
    return response
}

const deleteTagURL = "/api/deletetag"
export async function deleteTag(videoID, tagID) {
    const response = await mustFetch(deleteTagURL, {
        method: "POST",
        body: JSON.stringify({videoID: videoID, tagID: tagID})
    })
    return await response.text()
}

const tagsWithVideoURL = "/api/tagswithvideo"
export async function tagsWithVideo(videoID) {
    const response = await mustFetch(tagsWithVideoURL, {
        method: "POST",
        body: JSON.stringify({videoID})
    })
    return await response.json()
}

const updateVideoInfoURL = "/api/updatevideoinfo"
export async function updateVideoInfo(videoID, videoInfo) {
    const response = await mustFetch(updateVideoInfoURL, {
        method: "POST",
        body: JSON.stringify({videoID, videoInfo})
    })
    return response.text()
}