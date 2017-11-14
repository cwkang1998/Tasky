// Pending task initialization
function initPending(tasks) {
    //Set placeholder invisible and content visible
    let ph = document.getElementById("placeholder-pending");
    ph.style.display = "none";

    let content = document.getElementsByClassName("pending");
    for (let i = 0; i < content.length; i += 1) {
        content[i].style.display = "block";
    }

    //Population of the entries
    let fragment = document.createDocumentFragment();
    console.log(tasks.length);
    for (let i = 0; i < tasks.length; i++) {
        console.log(tasks[i]);
        let taskElements = createTask(tasks[i], "pending");
        fragment.appendChild(taskElements);
    }
    document.getElementById("pendingTasks").appendChild(fragment);
}

function showEmptyPending() {
    let ph = document.getElementById("placeholder-pending");
    ph.style.display = "inline";

    let content = document.getElementsByClassName("pending");
    for (let i = 0; i < content.length; i += 1) {
        content[i].style.display = "none";
    }
}

// For saving the tasks
function saveTasks(value) {
    let request = new XMLHttpRequest();
    request.onreadystatechange = function () {
        if (request.readyState == request.DONE) {
            if (request.status === 200) {
                reloadPending();
                showAlert("New Task Added.", "alert-success");
            } else {
                showAlert("Failed to Add New Task, Please try again later.", "alert-danger");
            }
        }
    };
    request.open("POST", ep_prefix + "/addTsk");
    request.setRequestHeader("Content-type", "application/json");
    request.send(JSON.stringify({ description: value }));
    console.log(request);
    console.log(JSON.stringify({ description: value }));
}

// Set task completed
function setCompleted(id) {

    makeGetRequest("GET", "/setTskStatus",
        function () {
            reloadAll();
            showAlert("Task Completed.", "alert-success");
        }
        , function () {
            showAlert("Action Failed. Please try again later", "alert-danger");
        }
        , "?task_id=" + encodeURIComponent(id) + "&status=1");

}

