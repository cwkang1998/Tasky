//Makes a get Request
function makeGetRequest(method, endpoints, successHandler, failureHandler, param) {
    let request = new XMLHttpRequest();
    request.onreadystatechange = function () {
        if (request.readyState == request.DONE && request.responseText != "") {
            let responseItems = JSON.parse(request.responseText);
            if (request.status == 200 && responseItems.length > 0) {
                successHandler(responseItems);
            } else {
                failureHandler();
            }
        } else if (request.readyState == request.DONE) {
            if (request.status == 200) {
                successHandler();
            } else {
                failureHandler();
            }
        }
    };
    request.open(method, ep_prefix + endpoints + param);
    request.setRequestHeader("Content-type", "application/json");
    request.send(param);
}

function createTask(task, identifyingClass) {
    //Parent task element
    let parent = document.createElement("div");
    parent.className = identifyingClass + " tasks-element shadow-dim-post container-fluid";
    parent.id = task.task_id;
    if (identifyingClass === "pending") {
        parent.onclick = function () { setCompleted(parent.id); };
    } else if (identifyingClass === "completed") {
        parent.onclick = attachShowOptions(parent.id);
    }

    //task element
    let topDiv = document.createElement("div");
    topDiv.className = "row";
    let colDiv = document.createElement("div");
    colDiv.className = "col-xs-12";
    let date = document.createElement("span");
    if (identifyingClass === "pending") {
        date.className = "label label-warning";
    } else if (identifyingClass === "completed") {
        date.className = "label label-info";
    }
    date.innerText = task.time;
    let taskName = document.createElement("span");
    taskName.className = "tasks-element-descp";
    taskName.innerText = task.description;
    colDiv.appendChild(date);
    colDiv.appendChild(taskName);
    topDiv.appendChild(colDiv);

    //Append all those elements to the parent
    parent.appendChild(colDiv);

    //returning back
    return parent;
}


function reloadPending() {
    let content = document.getElementsByClassName("pending");
    console.log(content.length);
    while (content.length > 0) {
        content[0].parentNode.removeChild(content[0]);
    }
    makeGetRequest("GET", "/getTsks", initPending, showEmptyPending, "?status=0");
}


function reloadCompleted() {
    let content = document.getElementsByClassName("completed");
    while (content.length > 0) {
        content[0].parentNode.removeChild(content[0]);
    }
    makeGetRequest("GET", "/getTsks", initCompleted, showEmptyCompleted, "?status=1");
}

function reloadAll() {
    reloadPending();
    reloadCompleted();
}

function showAlert(message, alertType) {
    $('#alert-pop-up')
        .append("<div id='alert-div' class='fade in alert alert-dismissible " +
        alertType + "'><span class='close' data-dismiss='alert' aria-label='close'>&times;</span><span>" +
        message + "</span></div>");

    setTimeout(function () {
        $("#alert-div").alert("close");
    }, 4000);
}